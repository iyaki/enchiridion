package main

import (
	"errors"
	"os"
	"os/exec"
	"strings"
	"testing"
)

// Child mode: TestMainSmoke re-executes the test binary to cover main().
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

func TestRunUnimplementedExitsTwo(t *testing.T) {
	if code := run([]string{"sync"}); code != 2 {
		t.Fatalf("sync: got exit %d, want 2", code)
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

func TestMainSmokeUnimplemented(t *testing.T) {
	t.Setenv("ENCHIRIDION_MAIN", "1")
	_, err := exec.Command(os.Args[0]).CombinedOutput()
	var exitErr *exec.ExitError
	if !errors.As(err, &exitErr) || exitErr.ExitCode() != 2 {
		t.Fatalf("no args: got %v, want exit code 2", err)
	}
}
