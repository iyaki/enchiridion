package sync

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestStateRoundTrip(t *testing.T) {
	home := t.TempDir()
	want := State{
		LastFullAt: time.Date(2026, 9, 1, 10, 0, 0, 0, time.UTC),
		Watermark:  time.Date(2026, 9, 20, 8, 30, 0, 0, time.UTC),
	}

	if err := SaveState(home, want); err != nil {
		t.Fatalf("SaveState: %v", err)
	}

	got, err := LoadState(home)
	if err != nil {
		t.Fatalf("LoadState: %v", err)
	}
	if got == nil || !got.LastFullAt.Equal(want.LastFullAt) || !got.Watermark.Equal(want.Watermark) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}

func TestLoadStateMissingFile(t *testing.T) {
	got, err := LoadState(t.TempDir())
	if err != nil {
		t.Fatalf("LoadState on missing file: %v", err)
	}
	if got != nil {
		t.Fatalf("got %+v, want nil", got)
	}
}

func TestLoadStateMissingHome(t *testing.T) {
	got, err := LoadState(filepath.Join(t.TempDir(), "nope"))
	if err != nil {
		t.Fatalf("LoadState on missing home: %v", err)
	}
	if got != nil {
		t.Fatalf("got %+v, want nil", got)
	}
}

func TestLoadStateCorruptJSON(t *testing.T) {
	home := t.TempDir()
	if err := os.WriteFile(filepath.Join(home, stateFile), []byte("{not json"), 0o600); err != nil {
		t.Fatalf("write corrupt state: %v", err)
	}

	got, err := LoadState(home)
	if err != nil {
		t.Fatalf("LoadState on corrupt file: %v", err)
	}
	if got != nil {
		t.Fatalf("got %+v, want nil", got)
	}
}
