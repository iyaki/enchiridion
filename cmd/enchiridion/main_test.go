package main

import (
	"errors"
	"os"
	"os/exec"
	"strings"
	"testing"
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
