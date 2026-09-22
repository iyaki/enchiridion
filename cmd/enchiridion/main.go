// Command enchiridion mirrors a Notion knowledge base as greppable markdown.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/iyaki/enchiridion/internal/notion"
	"github.com/iyaki/enchiridion/internal/pull"
	"github.com/iyaki/enchiridion/internal/search"
	"github.com/iyaki/enchiridion/internal/sync"
)

// version is injected at build time via -ldflags (see .goreleaser.yml).
var version = "dev"

// runEngine is the sync entry point; a variable so the command wiring stays
// testable without network (specs/architecture.md — testability).
var runEngine = sync.Run

// doctor seams: variables so doctor's connectivity checks run against fakes
// in tests (same pattern as runEngine).
type pinger interface {
	Ping(dataSourceID string) error
}

type repoChecker interface {
	Check(repo string) error
}

var newPinger = func(token string) pinger { return notion.NewClient(token) }

var newRepoChecker = func(token string) repoChecker { return pull.NewClient(token) }

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
		case "help", "-h", "--help":
			printUsage(os.Stdout)

			return exitOK
		case "version":
			fmt.Printf("enchiridion %s\n", version)

			return exitOK
		case "sync":
			return runSync(args[1:])
		case "pull":
			return runPull(args[1:])
		case "doctor":
			return runDoctor(args[1:])
		case "search":
			return runSearch(args[1:])
		}
	}

	printUsage(os.Stderr)

	return exitUsage
}

// printUsage prints the command summary to w (stderr on errors, stdout for
// explicit help requests).
func printUsage(w io.Writer) {
	// ponytail: best-effort write; usage errors are not actionable
	_, _ = fmt.Fprint(w, `enchiridion mirrors a Notion knowledge base as greppable markdown.

usage:
  enchiridion help
  enchiridion version
  enchiridion sync [--full] [--quiet]
  enchiridion pull [--out DIR]
  enchiridion doctor
  enchiridion search [--dir DIR] TERM [TERM...]

configuration (environment):
  NOTION_TOKEN                   Notion integration token (sync)
  KNOWLEDGE_BASE_DATASOURCE_ID   data source to mirror (sync)
  ENCHIRIDION_HOME               cache root (default: ~/.local/share/enchiridion)
  GITHUB_TOKEN                   token that can read the distribution repo (pull)
  ENCHIRIDION_REPO               distribution repo (default: `+pull.DefaultRepo+`)

sync chooses its mode automatically (full or incremental); --full forces a
full sync, the only mode that propagates page deletions; --quiet suppresses
progress output. "enchiridion <command> --help" explains each command.

pull vendors the published mirror into the current project (--out, default
"data"): no Notion credentials and no shared machine are required (ADR-18);
commit the result so every checkout of the project carries the mirror.

doctor checks the local configuration and connectivity without syncing.
`)
}

// wantsHelp reports whether the argument list is a help request, checked
// before flag parsing so --help always wins.
func wantsHelp(args []string) bool {
	return slices.Contains(args, "-h") || slices.Contains(args, "--help")
}

// syncUsage prints the sync command contract to w.
func syncUsage(w io.Writer) {
	_, _ = fmt.Fprint(w, `enchiridion sync mirrors the knowledge base into the local cache.

usage:
  enchiridion sync [--full] [--quiet]

flags:
  --full     force a full sync (the only mode that propagates page deletions)
  --quiet    suppress progress output on stderr; only the final summary prints

configuration (environment):
  NOTION_TOKEN                   Notion integration token
  KNOWLEDGE_BASE_DATASOURCE_ID   data source to mirror
  ENCHIRIDION_HOME               cache root (default: ~/.local/share/enchiridion)

sync chooses its mode automatically (full or incremental); --full forces a
full sync, the only mode that propagates page deletions. Progress is written
to stderr; stdout carries a single summary line.
`)
}

// pullUsage prints the pull command contract to w.
func pullUsage(w io.Writer) {
	_, _ = fmt.Fprint(w, `enchiridion pull vendors the published mirror into the current project.

usage:
  enchiridion pull [--out DIR]

flags:
  --out DIR    target directory for knowledge/ and tools/ (default "data")

configuration (environment):
  GITHUB_TOKEN      token that can read the distribution repo
  ENCHIRIDION_REPO  distribution repo (default: `+pull.DefaultRepo+`)

The download completes before anything on disk is touched, so a failed pull
never damages an existing mirror; commit the result so every checkout of the
project carries it (ADR-18).
`)
}

// searchUsage prints the search command contract to w.
func searchUsage(w io.Writer) {
	_, _ = fmt.Fprint(w, `enchiridion search searches the mirror for files matching every term.

usage:
  enchiridion search [--dir DIR] TERM [TERM...]

flags:
  --dir DIR    mirror root to search (default: the sync cache home;
               pass a vendored data/ directory in consumer projects)

Matching is case-insensitive substring over title, tags, filename and body
(the title ranks highest). Output: one "path — title" line per hit, best
first; exit 1 when nothing matches.
`)
}

// runSync mirrors the knowledge base into the local cache. Configuration
// problems fail before any API contact (specs/architecture.md — error
// behavior); per-page failures are logged by the engine and make the run
// exit non-zero.
func runSync(args []string) int {
	if wantsHelp(args) {
		syncUsage(os.Stdout)

		return exitOK
	}

	forceFull, quiet, err := parseSyncArgs(args)
	if err != nil {
		syncUsage(os.Stderr)

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

	// Progress goes to stderr; stdout carries only the final summary.
	// A TTY gets one updating line, CI gets a heartbeat every 100 pages.
	progress := newProgress(quiet)
	stats, err := runEngine(notion.NewClient(token), sync.Options{
		DataSourceID: dataSourceID,
		Home:         home,
		ForceFull:    forceFull,
		Now:          time.Now(),
		Progress:     progress.report,
	})
	progress.close()
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
	if wantsHelp(args) {
		pullUsage(os.Stdout)

		return exitOK
	}

	out, err := parsePullArgs(args)
	if err != nil {
		pullUsage(os.Stderr)

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

// runSearch ranks mirror files by term matches and prints them; the caller
// judges relevance. Read-only: it never contacts the API nor takes the lock.
func runSearch(args []string) int {
	if wantsHelp(args) {
		searchUsage(os.Stdout)

		return exitOK
	}

	dir, terms, err := parseSearchArgs(args)
	if err != nil {
		searchUsage(os.Stderr)

		return exitUsage
	}
	root := dir
	if root == "" {
		if root, err = resolveHome(); err != nil {
			fmt.Fprintf(os.Stderr, "enchiridion: %v\n", err)

			return exitError
		}
	}
	hits, err := search.Run(root, terms)
	if err != nil {
		fmt.Fprintf(os.Stderr, "enchiridion: %v\n", err)

		return exitError
	}
	for _, hit := range hits {
		fmt.Printf("%s — %s\n", hit.Path, hit.Title)
	}
	if len(hits) == 0 {
		return exitError
	}

	return exitOK
}

// runDoctor checks the local configuration and connectivity without syncing:
// never Notion content, never the cache lock. Exit 0 means the sync surface
// is ready; warnings (optional surfaces) never fail the run.
func runDoctor(args []string) int {
	if wantsHelp(args) {
		doctorUsage(os.Stdout)

		return exitOK
	}

	token, dataSourceID := os.Getenv(envToken), os.Getenv(envSource)
	tokenSet := checkEnv(envToken, "NOTION_TOKEN set", "FAIL")
	idSet := checkEnv(envSource, "KNOWLEDGE_BASE_DATASOURCE_ID set", "FAIL")

	failed := !(tokenSet && idSet)
	failed = doctorHome() || failed
	failed = doctorNotion(token, dataSourceID, tokenSet && idSet) || failed
	failed = doctorState() || failed
	doctorPull()

	if failed {
		return exitError
	}

	return exitOK
}

// printCheck prints one doctor check line to stdout.
func printCheck(tag, label string) {
	fmt.Printf("%-4s %s\n", tag, label)
}

// checkEnv reports whether the variable is set, tagging its line FAIL
// (required) or warn (optional) when it is not.
func checkEnv(name, label, missingTag string) bool {
	if value := os.Getenv(name); value != "" {
		printCheck("ok", label)

		return true
	}
	printCheck(missingTag, label)

	return false
}

// doctorHome verifies the cache root resolves and reports whether it exists.
func doctorHome() bool {
	home, err := resolveHome()
	if err != nil {
		printCheck("FAIL", "cache home: "+err.Error())

		return true
	}
	if _, err := os.Stat(home); err != nil {
		printCheck("warn", "cache home: "+home+" (created on first sync)")

		return false
	}
	printCheck("ok", "cache home: "+home)

	return false
}

// doctorNotion verifies credentials and data source reachability; only a
// reachable Notion counts as pass, and failure fails the run.
func doctorNotion(token, dataSourceID string, configured bool) bool {
	if !configured {
		return false
	}
	if err := newPinger(token).Ping(dataSourceID); err != nil {
		printCheck("FAIL", "notion: credentials and data source reachable: "+err.Error())

		return true
	}
	printCheck("ok", "notion: credentials and data source reachable")

	return false
}

// doctorState reports the last full sync watermark.
func doctorState() bool {
	home, err := resolveHome()
	if err != nil {
		return false // already reported by doctorHome
	}
	state, err := sync.LoadState(home)
	switch {
	case err != nil:
		printCheck("FAIL", "last full sync: "+err.Error())

		return true
	case state == nil:
		printCheck("warn", "last full sync: no watermark yet (first sync pending)")
	default:
		printCheck("ok", "last full sync: "+state.LastFullAt.Format(time.RFC3339))
	}

	return false
}

// doctorPull verifies the optional pull surface against the distribution
// repo; every outcome here is at most a warning.
func doctorPull() {
	if !checkEnv(envGithubToken, "GITHUB_TOKEN set (pull)", "warn") {
		return
	}
	repo := os.Getenv(envRepo)
	if repo == "" {
		repo = pull.DefaultRepo
	}
	if err := newRepoChecker(os.Getenv(envGithubToken)).Check(repo); err != nil {
		printCheck("warn", "pull: GITHUB_TOKEN can read "+repo+": "+err.Error())

		return
	}
	printCheck("ok", "pull: GITHUB_TOKEN can read "+repo)
}

// doctorUsage prints the doctor command contract to w.
func doctorUsage(w io.Writer) {
	_, _ = fmt.Fprint(w, `enchiridion doctor checks the local configuration and connectivity.

usage:
  enchiridion doctor

Read-only: it never syncs and never takes the cache lock. It verifies the
Notion credentials and data source (the sync surface), the cache home and
watermark, and — when GITHUB_TOKEN is set — access to the distribution repo
for pull. Exit 0 means sync is ready to run; warnings never fail.
`)
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

// parseSearchArgs parses search flags; only --dir is allowed, and at least
// one term is required.
func parseSearchArgs(args []string) (dir string, terms []string, err error) {
	fs := flag.NewFlagSet("search", flag.ContinueOnError)
	fs.SetOutput(io.Discard) // searchUsage explains the contract on any error
	fs.StringVar(&dir, "dir", "", "mirror root to search (default: the sync cache home)")
	if err := fs.Parse(args); err != nil {
		return "", nil, err
	}
	if fs.NArg() == 0 {
		return "", nil, fmt.Errorf("no search terms")
	}

	return dir, fs.Args(), nil
}

// parseSyncArgs parses sync flags; --full and --quiet are allowed.
func parseSyncArgs(args []string) (forceFull, quiet bool, err error) {
	fs := flag.NewFlagSet("sync", flag.ContinueOnError)
	fs.SetOutput(io.Discard) // syncUsage explains the contract on any error
	fs.BoolVar(&forceFull, "full", false, "force a full sync")
	fs.BoolVar(&quiet, "quiet", false, "suppress progress output on stderr")
	if err := fs.Parse(args); err != nil {
		return false, false, err
	}
	if fs.NArg() > 0 {
		return false, false, fmt.Errorf("unexpected argument: %s", fs.Arg(0))
	}

	return forceFull, quiet, nil
}

// progress renders sync progress on stderr; nil is valid and silent.
type progress struct {
	tty   bool
	shown bool
}

// newProgress returns the reporter for a run, or nil under --quiet.
func newProgress(quiet bool) *progress {
	if quiet {
		return nil
	}

	return &progress{tty: stderrIsTTY()}
}

// report records one progress tick: an updating line on a TTY, a heartbeat
// every 100 pages (and at the end) otherwise.
func (p *progress) report(done, total int) {
	if p == nil {
		return
	}
	p.shown = true
	switch {
	case p.tty:
		fmt.Fprintf(os.Stderr, "\rsync: %d/%d pages", done, total)
	case done%100 == 0 || done == total:
		fmt.Fprintf(os.Stderr, "sync: %d/%d pages\n", done, total)
	}
}

// close ends the updating TTY line, if one was started.
func (p *progress) close() {
	if p != nil && p.shown && p.tty {
		fmt.Fprintln(os.Stderr)
	}
}

// stderrIsTTY reports whether stderr is an interactive terminal: progress
// then renders as a single updating line instead of newline heartbeats.
func stderrIsTTY() bool {
	info, err := os.Stderr.Stat()

	return err == nil && info.Mode()&os.ModeCharDevice != 0
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
