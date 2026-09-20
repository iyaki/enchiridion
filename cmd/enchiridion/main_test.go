package main

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/iyaki/enchiridion/internal/sync"
)

// Child mode: smoke tests re-execute the test binary to cover main().
func TestMain(m *testing.M) {
	if os.Getenv("ENCHIRIDION_MAIN") != "" {
		main()

		return
	}
	os.Exit(m.Run())
}

func TestRunVersionExitsZero(t *testing.T) {
	if code := run([]string{"version"}); code != 0 {
		t.Fatalf("version: got exit %d, want 0", code)
	}
}

func TestRunUnknownCommandExitsTwo(t *testing.T) {
	for _, args := range [][]string{nil, {"nope"}, {"--help"}} {
		if code := run(args); code != exitUsage {
			t.Fatalf("run(%v): got exit %d, want %d", args, code, exitUsage)
		}
	}
}

func TestParseSyncArgs(t *testing.T) {
	forced, err := parseSyncArgs([]string{"--full"})
	if err != nil || !forced {
		t.Fatalf("--full: got (%t, %v), want (true, nil)", forced, err)
	}

	plain, err := parseSyncArgs(nil)
	if err != nil || plain {
		t.Fatalf("no flags: got (%t, %v), want (false, nil)", plain, err)
	}

	if _, err := parseSyncArgs([]string{"--bogus"}); err == nil {
		t.Fatal("--bogus: got nil error, want error")
	}
	if _, err := parseSyncArgs([]string{"stray"}); err == nil {
		t.Fatal("stray argument: got nil error, want error")
	}
}

func TestRunSyncMissingConfigExitsOne(t *testing.T) {
	t.Setenv(envToken, "")
	t.Setenv(envSource, "")

	if code := runSync(nil); code != exitError {
		t.Fatalf("sync without config: got exit %d, want %d", code, exitError)
	}
}

func TestRunSyncSuccess(t *testing.T) {
	t.Setenv(envToken, "secret")
	t.Setenv(envSource, "ds-1")
	home := t.TempDir()
	t.Setenv(envHome, home)

	var gotOpts sync.Options
	orig := runEngine
	runEngine = func(_ sync.NotionAPI, opts sync.Options) (sync.Stats, error) {
		gotOpts = opts

		return sync.Stats{Mode: "full", Kept: 1, Written: 1}, nil
	}
	defer func() { runEngine = orig }()

	if code := runSync([]string{"--full"}); code != exitOK {
		t.Fatalf("sync: got exit %d, want %d", code, exitOK)
	}
	if !gotOpts.ForceFull || gotOpts.DataSourceID != "ds-1" || gotOpts.Home != home {
		t.Fatalf("unexpected options: %+v", gotOpts)
	}
}

func TestRunSyncEngineErrorExitsOne(t *testing.T) {
	t.Setenv(envToken, "secret")
	t.Setenv(envSource, "ds-1")
	t.Setenv(envHome, t.TempDir())

	orig := runEngine
	runEngine = func(sync.NotionAPI, sync.Options) (sync.Stats, error) {
		return sync.Stats{}, errors.New("1 of 2 pages failed")
	}
	defer func() { runEngine = orig }()

	if code := runSync(nil); code != exitError {
		t.Fatalf("sync: got exit %d, want %d", code, exitError)
	}
}

func TestResolveHome(t *testing.T) {
	t.Setenv(envHome, "/custom/home")
	t.Setenv(envXDGData, "/xdg")
	if got, err := resolveHome(); err != nil || got != "/custom/home" {
		t.Fatalf("ENCHIRIDION_HOME: got (%q, %v)", got, err)
	}

	t.Setenv(envHome, "")
	if got, err := resolveHome(); err != nil || got != filepath.Join("/xdg", "enchiridion") {
		t.Fatalf("XDG_DATA_HOME: got (%q, %v)", got, err)
	}

	t.Setenv(envXDGData, "")
	t.Setenv("HOME", "/home/user")
	want := filepath.Join("/home/user", ".local", "share", "enchiridion")
	if got, err := resolveHome(); err != nil || got != want {
		t.Fatalf("default: got (%q, %v), want %q", got, err, want)
	}
}

func TestMainSmokeVersion(t *testing.T) {
	t.Setenv("ENCHIRIDION_MAIN", "1")
	out, err := exec.Command(os.Args[0], "version").CombinedOutput()
	if err != nil {
		t.Fatalf("version: unexpected error: %v (%s)", err, out)
	}
	if !strings.Contains(string(out), "enchiridion") {
		t.Fatalf("version: unexpected output %q", out)
	}
}

func TestMainSmokeUsageExitsTwo(t *testing.T) {
	t.Setenv("ENCHIRIDION_MAIN", "1")
	_, err := exec.Command(os.Args[0]).CombinedOutput()
	var exitErr *exec.ExitError
	if !errors.As(err, &exitErr) || exitErr.ExitCode() != exitUsage {
		t.Fatalf("no args: got %v, want exit code %d", err, exitUsage)
	}
}

func TestMainSmokeSyncWithoutConfig(t *testing.T) {
	t.Setenv("ENCHIRIDION_MAIN", "1")
	cmd := exec.Command(os.Args[0], "sync")
	cmd.Env = append(os.Environ(), envToken+"=", envSource+"=")

	out, err := cmd.CombinedOutput()
	var exitErr *exec.ExitError
	if !errors.As(err, &exitErr) || exitErr.ExitCode() != exitError {
		t.Fatalf("sync without config: got %v, want exit %d (output: %s)", err, exitError, out)
	}
	if !strings.Contains(string(out), envToken) || !strings.Contains(string(out), envSource) {
		t.Fatalf("error must name the missing variables, got: %s", out)
	}
}
