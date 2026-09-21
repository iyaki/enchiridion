// Command enchiridion mirrors a Notion knowledge base as greppable markdown.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/iyaki/enchiridion/internal/notion"
	"github.com/iyaki/enchiridion/internal/pull"
	"github.com/iyaki/enchiridion/internal/sync"
)

// version is injected at build time via -ldflags (see .goreleaser.yml).
var version = "dev"

// runEngine is the sync entry point; a variable so the command wiring stays
// testable without network (specs/architecture.md — testability).
var runEngine = sync.Run

const (
	envToken   = "NOTION_TOKEN"
	envSource  = "KNOWLEDGE_BASE_DATASOURCE_ID"
	envHome    = "ENCHIRIDION_HOME"
	envXDGData = "XDG_DATA_HOME"

	envGithubToken = "GITHUB_TOKEN" // #nosec G101 -- environment variable name, not a credential
	envRepo        = "ENCHIRIDION_REPO"

	dirPerm = 0o755

	exitOK    = 0
	exitError = 1
	exitUsage = 2
)

func main() {
	os.Exit(run(os.Args[1:]))
}

// run dispatches a command and returns the process exit code.
func run(args []string) int {
	if len(args) > 0 {
		switch args[0] {
		case "version":
			fmt.Printf("enchiridion %s\n", version)

			return exitOK
		case "sync":
			return runSync(args[1:])
		case "pull":
			return runPull(args[1:])
		}
	}

	usage()

	return exitUsage
}

// usage prints the command summary to stderr.
func usage() {
	fmt.Fprint(os.Stderr, `enchiridion mirrors a Notion knowledge base as greppable markdown.

usage:
  enchiridion version
  enchiridion sync [--full]
  enchiridion pull [--out DIR]

configuration (environment):
  NOTION_TOKEN                   Notion integration token (sync)
  KNOWLEDGE_BASE_DATASOURCE_ID   data source to mirror (sync)
  ENCHIRIDION_HOME               cache root (default: ~/.local/share/enchiridion)
  GITHUB_TOKEN                   token that can read the distribution repo (pull)
  ENCHIRIDION_REPO               distribution repo (default: `+pull.DefaultRepo+`)

sync chooses its mode automatically (full or incremental); --full forces a
full sync, the only mode that propagates page deletions.

pull vendors the published mirror into the current project (--out, default
"data"): no Notion credentials and no shared machine are required (ADR-18);
commit the result so every checkout of the project carries the mirror.
`)
}

// runSync mirrors the knowledge base into the local cache. Configuration
// problems fail before any API contact (specs/architecture.md — error
// behavior); per-page failures are logged by the engine and make the run
// exit non-zero.
func runSync(args []string) int {
	forceFull, err := parseSyncArgs(args)
	if err != nil {
		usage()

		return exitUsage
	}

	token, dataSourceID, home, err := resolveConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "enchiridion: %v\n", err)

		return exitError
	}
	if err := os.MkdirAll(home, dirPerm); err != nil {
		fmt.Fprintf(os.Stderr, "enchiridion: %v\n", err)

		return exitError
	}

	stats, err := runEngine(notion.NewClient(token), sync.Options{
		DataSourceID: dataSourceID,
		Home:         home,
		ForceFull:    forceFull,
		Now:          time.Now(),
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "enchiridion: %v\n", err)

		return exitError
	}

	fmt.Printf("sync complete: mode=%s kept=%d written=%d removed=%d failed=%d\n",
		stats.Mode, stats.Kept, stats.Written, stats.Removed, stats.Failed)

	return exitOK
}

// runPull vendors the published mirror into the current project: it needs a
// GitHub token that can read the distribution repo, never Notion credentials
// (ADR-18). The result is meant to be committed by the consumer project.
func runPull(args []string) int {
	out, err := parsePullArgs(args)
	if err != nil {
		usage()

		return exitUsage
	}

	token := os.Getenv(envGithubToken)
	if token == "" {
		fmt.Fprintf(os.Stderr, "enchiridion: missing required environment variable: %s\n", envGithubToken)

		return exitError
	}
	repo := os.Getenv(envRepo)
	if repo == "" {
		repo = pull.DefaultRepo
	}

	stats, err := pull.NewClient(token).Pull(repo, out)
	if err != nil {
		fmt.Fprintf(os.Stderr, "enchiridion: %v\n", err)

		return exitError
	}

	fmt.Printf("pull complete: knowledge=%d tools=%d out=%s\n", stats.Knowledge, stats.Tools, out)

	return exitOK
}

// parsePullArgs parses pull flags; only --out is allowed.
func parsePullArgs(args []string) (out string, err error) {
	fs := flag.NewFlagSet("pull", flag.ContinueOnError)
	fs.SetOutput(io.Discard) // usage() explains the contract on any error
	fs.StringVar(&out, "out", "data", "target directory for knowledge/ and tools/")
	if err := fs.Parse(args); err != nil {
		return "", err
	}
	if fs.NArg() > 0 {
		return "", fmt.Errorf("unexpected argument: %s", fs.Arg(0))
	}

	return out, nil
}

// parseSyncArgs parses sync flags; only --full is allowed.
func parseSyncArgs(args []string) (forceFull bool, err error) {
	fs := flag.NewFlagSet("sync", flag.ContinueOnError)
	fs.SetOutput(io.Discard) // usage() explains the contract on any error
	fs.BoolVar(&forceFull, "full", false, "force a full sync")
	if err := fs.Parse(args); err != nil {
		return false, err
	}
	if fs.NArg() > 0 {
		return false, fmt.Errorf("unexpected argument: %s", fs.Arg(0))
	}

	return forceFull, nil
}

// resolveConfig gathers credentials and the cache root without contacting the
// API; missing variables are all reported at once.
func resolveConfig() (token, dataSourceID, home string, err error) {
	token = os.Getenv(envToken)
	dataSourceID = os.Getenv(envSource)
	var missing []string
	if token == "" {
		missing = append(missing, envToken)
	}
	if dataSourceID == "" {
		missing = append(missing, envSource)
	}
	if len(missing) > 0 {
		return "", "", "", fmt.Errorf("missing required environment variables: %s", strings.Join(missing, ", "))
	}

	home, err = resolveHome()
	if err != nil {
		return "", "", "", err
	}

	return token, dataSourceID, home, nil
}

// resolveHome returns the cache root: $ENCHIRIDION_HOME, else
// $XDG_DATA_HOME/enchiridion, else ~/.local/share/enchiridion
// (specs/integration.md).
func resolveHome() (string, error) {
	if home := os.Getenv(envHome); home != "" {
		return home, nil
	}
	if xdg := os.Getenv(envXDGData); xdg != "" {
		return filepath.Join(xdg, "enchiridion"), nil
	}

	base, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("resolve cache home: %w", err)
	}

	return filepath.Join(base, ".local", "share", "enchiridion"), nil
}
