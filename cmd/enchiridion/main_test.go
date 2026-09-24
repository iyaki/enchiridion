package main

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
	"testing"
	"time"

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
	for _, args := range [][]string{nil, {"nope"}} {
		if code := run(args); code != exitUsage {
			t.Fatalf("run(%v): got exit %d, want %d", args, code, exitUsage)
		}
	}
}

func TestRunHelpExitsZero(t *testing.T) {
	for _, args := range [][]string{{"help"}, {"--help"}, {"-h"}} {
		if code := run(args); code != exitOK {
			t.Fatalf("run(%v): got exit %d, want %d", args, code, exitOK)
		}
	}
}

func TestCommandHelpExitsZero(t *testing.T) {
	for _, tc := range []struct {
		name string
		code int
	}{
		{"sync --help", runSync([]string{"--help"})},
		{"sync -h", runSync([]string{"-h"})},
		{"pull --help", runPull([]string{"--help"})},
		{"pull -h", runPull([]string{"-h"})},
		{"doctor --help", runDoctor([]string{"--help"})},
		{"doctor -h", runDoctor([]string{"-h"})},
		{"search --help", runSearch([]string{"--help"})},
		{"search -h", runSearch([]string{"-h"})},
	} {
		if tc.code != exitOK {
			t.Fatalf("%s: got exit %d, want %d", tc.name, tc.code, exitOK)
		}
	}
}

func TestParseSyncArgs(t *testing.T) {
	cases := []struct {
		name      string
		args      []string
		forceFull bool
		quiet     bool
		wantErr   bool
	}{
		{"no flags", nil, false, false, false},
		{"--full", []string{"--full"}, true, false, false},
		{"--quiet", []string{"--quiet"}, false, true, false},
		{"--full --quiet", []string{"--full", "--quiet"}, true, true, false},
		{"--bogus", []string{"--bogus"}, false, false, true},
		{"stray argument", []string{"stray"}, false, false, true},
	}
	for _, tc := range cases {
		forceFull, quiet, err := parseSyncArgs(tc.args)
		if (err != nil) != tc.wantErr || forceFull != tc.forceFull || quiet != tc.quiet {
			t.Errorf("%s: got (%t, %t, %v), want (%t, %t, err=%t)",
				tc.name, forceFull, quiet, err, tc.forceFull, tc.quiet, tc.wantErr)
		}
	}
}

func TestRunSyncMissingConfigExitsOne(t *testing.T) {
	t.Setenv(envToken, "")
	t.Setenv(envSource, "")

	if code := runSync(nil); code != exitError {
		t.Fatalf("sync without config: got exit %d, want %d", code, exitError)
	}
}

func TestParsePullArgs(t *testing.T) {
	const home = "/tmp/ench-home"
	t.Setenv(envHome, home)

	out, err := parsePullArgs(nil)
	if err != nil || out != home {
		t.Fatalf("no flags: got (%q, %v), want (%q, nil)", out, err, home)
	}

	project, err := parsePullArgs([]string{"--project"})
	if err != nil || project != "data" {
		t.Fatalf("--project: got (%q, %v), want (\"data\", nil)", project, err)
	}

	custom, err := parsePullArgs([]string{"--out", "vendor/kb"})
	if err != nil || custom != "vendor/kb" {
		t.Fatalf("--out: got (%q, %v), want (\"vendor/kb\", nil)", custom, err)
	}

	overrides, err := parsePullArgs([]string{"--project", "--out", "vendor/kb"})
	if err != nil || overrides != "vendor/kb" {
		t.Fatalf("--project --out: got (%q, %v), want (\"vendor/kb\", nil)", overrides, err)
	}
}

func TestParsePullArgsErrors(t *testing.T) {
	if _, err := parsePullArgs([]string{"--bogus"}); err == nil {
		t.Fatal("--bogus: got nil error, want error")
	}
	if _, err := parsePullArgs([]string{"stray"}); err == nil {
		t.Fatal("stray argument: got nil error, want error")
	}
}

func TestRunPullMissingTokenExitsOne(t *testing.T) {
	t.Setenv(envGithubToken, "")

	if code := runPull(nil); code != exitError {
		t.Fatalf("pull without token: got exit %d, want %d", code, exitError)
	}
}

func TestParseSearchArgs(t *testing.T) {
	dir, terms, err := parseSearchArgs([]string{"--dir", "data", "a11y", "audit"})
	if err != nil || dir != "data" || !slices.Equal(terms, []string{"a11y", "audit"}) {
		t.Fatalf("--dir + terms: got (%q, %v, %v)", dir, terms, err)
	}

	dir, terms, err = parseSearchArgs([]string{"go", "testing"})
	if err != nil || dir != "" || !slices.Equal(terms, []string{"go", "testing"}) {
		t.Fatalf("bare terms: got (%q, %v, %v)", dir, terms, err)
	}

	if _, _, err := parseSearchArgs([]string{"--bogus"}); err == nil {
		t.Fatal("--bogus: got nil error, want error")
	}
	if _, _, err := parseSearchArgs(nil); err == nil {
		t.Fatal("no terms: got nil error, want error")
	}
}

func TestRunSearchFindsAndMisses(t *testing.T) {
	root := t.TempDir()
	dir := filepath.Join(root, "knowledge")
	if err := os.MkdirAll(dir, dirPerm); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	page := "---\ntitle: \"Accessibility Auditing\"\nnotion_id: id\nnotion_url: u\nlast_edited: t\n---\nbody\n"
	if err := os.WriteFile(filepath.Join(dir, "a11y.md"), []byte(page), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}

	var code int
	out := captureStdout(t, func() {
		code = runSearch([]string{"--dir", root, "accessibility"})
	})
	if code != exitOK {
		t.Fatalf("matching term: got exit %d, want %d", code, exitOK)
	}
	if !strings.Contains(out, "knowledge/a11y.md") || !strings.Contains(out, "Accessibility Auditing") {
		t.Fatalf("matching term: stdout %q, want path and title", out)
	}

	out = captureStdout(t, func() {
		code = runSearch([]string{"--dir", root, "zzzznoterm"})
	})
	if code != exitError {
		t.Fatalf("no match: got exit %d, want %d", code, exitError)
	}
	if out != "" {
		t.Fatalf("no match: stdout %q, want empty", out)
	}
}

func TestRunSearchNoTermsExitsTwo(t *testing.T) {
	if code := runSearch(nil); code != exitUsage {
		t.Fatalf("no terms: got exit %d, want %d", code, exitUsage)
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

// captureStderr swaps os.Stderr for a temp file while fn runs and returns
// what was written. Not parallel-safe; no test in this file uses t.Parallel.
func captureStderr(t *testing.T, fn func()) string {
	t.Helper()
	f, err := os.CreateTemp(t.TempDir(), "stderr")
	if err != nil {
		t.Fatalf("temp file: %v", err)
	}

	orig := os.Stderr
	os.Stderr = f
	defer func() { os.Stderr = orig }()
	fn()
	if err := f.Close(); err != nil {
		t.Fatalf("close: %v", err)
	}

	data, err := os.ReadFile(f.Name())
	if err != nil {
		t.Fatalf("read: %v", err)
	}

	return string(data)
}

func TestRunSyncProgressOnStderr(t *testing.T) {
	t.Setenv(envToken, "secret")
	t.Setenv(envSource, "ds-1")
	t.Setenv(envHome, t.TempDir())

	orig := runEngine
	runEngine = func(_ sync.NotionAPI, opts sync.Options) (sync.Stats, error) {
		if opts.Progress == nil {
			t.Error("progress callback not wired")
		} else {
			opts.Progress(200, 300) // 200 % 100 == 0: must print on non-TTY
		}

		return sync.Stats{Mode: "full", Kept: 2}, nil
	}
	defer func() { runEngine = orig }()

	stderr := captureStderr(t, func() {
		if code := runSync(nil); code != exitOK {
			t.Errorf("sync: got exit %d, want %d", code, exitOK)
		}
	})
	if !strings.Contains(stderr, "sync: 200/300 pages") {
		t.Fatalf("stderr lacks progress line: %q", stderr)
	}

	stderr = captureStderr(t, func() {
		if code := runSync([]string{"--quiet"}); code != exitOK {
			t.Errorf("sync --quiet: got exit %d, want %d", code, exitOK)
		}
	})
	if strings.Contains(stderr, "sync:") {
		t.Fatalf("--quiet must suppress progress, got: %q", stderr)
	}
}

// captureStdout swaps os.Stdout for a temp file while fn runs and returns
// what was written. Not parallel-safe; no test in this file uses t.Parallel.
func captureStdout(t *testing.T, fn func()) string {
	t.Helper()
	f, err := os.CreateTemp(t.TempDir(), "stdout")
	if err != nil {
		t.Fatalf("temp file: %v", err)
	}

	orig := os.Stdout
	os.Stdout = f
	defer func() { os.Stdout = orig }()
	fn()
	if err := f.Close(); err != nil {
		t.Fatalf("close: %v", err)
	}

	data, err := os.ReadFile(f.Name())
	if err != nil {
		t.Fatalf("read: %v", err)
	}

	return string(data)
}

// fakePinger and fakeChecker stand in for doctor's clients.
type fakePinger struct{ err error }

func (f fakePinger) Ping(string) error { return f.err }

type fakeChecker struct{ err error }

func (f fakeChecker) Check(string) error { return f.err }

// restoreDoctor injects doctor fakes for the duration of the test.
func restoreDoctor(t *testing.T, p pinger, c repoChecker) {
	t.Helper()
	origP, origC := newPinger, newRepoChecker
	newPinger = func(string) pinger { return p }
	newRepoChecker = func(string) repoChecker { return c }
	t.Cleanup(func() { newPinger, newRepoChecker = origP, origC })
}

func TestRunDoctorAllGreenExitsZero(t *testing.T) {
	t.Setenv(envToken, "secret")
	t.Setenv(envSource, "ds-1")
	t.Setenv(envGithubToken, "ght")
	home := t.TempDir()
	t.Setenv(envHome, home)
	if err := sync.SaveState(home, sync.State{LastFullAt: time.Now(), Watermark: time.Now()}); err != nil {
		t.Fatalf("seed state: %v", err)
	}
	restoreDoctor(t, fakePinger{}, fakeChecker{})

	stdout := captureStdout(t, func() {
		if code := runDoctor(nil); code != exitOK {
			t.Errorf("doctor: got exit %d, want %d", code, exitOK)
		}
	})
	for _, want := range []string{"ok", "cache home:", "notion: credentials", "last full sync", "pull: GITHUB_TOKEN"} {
		if !strings.Contains(stdout, want) {
			t.Fatalf("doctor output lacks %q:\n%s", want, stdout)
		}
	}
}

func TestRunDoctorMissingConfigExitsOne(t *testing.T) {
	t.Setenv(envToken, "")
	t.Setenv(envSource, "")
	t.Setenv(envGithubToken, "")
	t.Setenv(envHome, t.TempDir())

	if code := run([]string{"doctor"}); code != exitError {
		t.Fatalf("doctor without config: got exit %d, want %d", code, exitError)
	}
}

func TestRunDoctorWarnOnlyStillExitsZero(t *testing.T) {
	t.Setenv(envToken, "secret")
	t.Setenv(envSource, "ds-1")
	t.Setenv(envGithubToken, "")
	t.Setenv(envHome, t.TempDir()) // no state file: watermark warn
	restoreDoctor(t, fakePinger{}, fakeChecker{})

	stdout := captureStdout(t, func() {
		if code := runDoctor(nil); code != exitOK {
			t.Errorf("doctor: got exit %d, want %d", code, exitOK)
		}
	})
	if !strings.Contains(stdout, "no watermark yet") || !strings.Contains(stdout, "warn") {
		t.Fatalf("doctor output lacks warnings:\n%s", stdout)
	}
}

func TestRunDoctorNotionFailureExitsOne(t *testing.T) {
	t.Setenv(envToken, "secret")
	t.Setenv(envSource, "ds-1")
	t.Setenv(envGithubToken, "")
	t.Setenv(envHome, t.TempDir())
	restoreDoctor(t, fakePinger{err: errors.New("notion rejected the credentials")}, fakeChecker{})

	if code := runDoctor(nil); code != exitError {
		t.Fatalf("doctor with notion failure: got exit %d, want %d", code, exitError)
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

func TestMainSmokeHelpPrintsToStdout(t *testing.T) {
	t.Setenv("ENCHIRIDION_MAIN", "1")
	var out, errOut bytes.Buffer
	cmd := exec.Command(os.Args[0], "--help")
	cmd.Stdout, cmd.Stderr = &out, &errOut

	err := cmd.Run()
	var exitErr *exec.ExitError
	if err != nil && (!errors.As(err, &exitErr) || exitErr.ExitCode() != exitOK) {
		t.Fatalf("--help: got %v, want exit code %d (stderr: %s)", err, exitOK, errOut.String())
	}
	if !strings.Contains(out.String(), "usage:") {
		t.Fatalf("--help must print usage to stdout, got: %q", out.String())
	}
	if errOut.Len() != 0 {
		t.Fatalf("--help must not write to stderr, got: %q", errOut.String())
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
