package sync

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/iyaki/enchiridion/internal/model"
	"github.com/iyaki/enchiridion/internal/notion"
)

// fakeAPI is a map-backed NotionAPI: deterministic, no HTTP.
type fakeAPI struct {
	pages       map[string]model.PageMeta
	blocks      map[string][]model.Block
	errPage     string // page whose blocks fetch fails
	queryFilter notion.QueryFilter
}

func (f *fakeAPI) QueryPages(_ string, filter notion.QueryFilter) ([]model.PageMeta, error) {
	f.queryFilter = filter
	marked, err := time.Parse(time.RFC3339, filter.LastEditedAfter)
	var metas []model.PageMeta
	for _, meta := range f.pages {
		edited, _ := time.Parse(time.RFC3339, meta.LastEdited)
		if err == nil && edited.Before(marked) {
			continue // the real API filters server-side
		}
		metas = append(metas, meta)
	}

	return metas, nil
}

func (f *fakeAPI) PageBlocks(pageID string) ([]model.Block, error) {
	if pageID == f.errPage {
		return nil, errors.New("boom")
	}

	return f.blocks[pageID], nil
}

// queryErrAPI fails the page query itself: the data source is unreachable.
type queryErrAPI struct{ err error }

func (f *queryErrAPI) QueryPages(string, notion.QueryFilter) ([]model.PageMeta, error) {
	return nil, f.err
}

func (f *queryErrAPI) PageBlocks(string) ([]model.Block, error) {
	return nil, errors.New("unreachable")
}

var (
	now       = time.Date(2026, 9, 20, 12, 0, 0, 0, time.UTC)
	editedOld = now.Add(-2 * time.Hour).Format(time.RFC3339)
	editedNew = now.Add(30 * time.Minute).Format(time.RFC3339)
)

func TestSelectMode(t *testing.T) {
	fresh := &State{LastFullAt: now.Add(-24 * time.Hour), Watermark: now.Add(-time.Hour)}
	stale := &State{LastFullAt: now.Add(-31 * 24 * time.Hour), Watermark: now.Add(-time.Hour)}

	cases := []struct {
		name       string
		state      *State
		mirrorLost bool
		force      bool
		want       string
	}{
		{"no state", nil, false, false, modeFull},
		{"empty mirror", fresh, true, false, modeFull},
		{"stale last full", stale, false, false, modeFull},
		{"fresh state", fresh, false, false, modeIncremental},
		{"forced", fresh, false, true, modeFull},
		{"forced without state", nil, false, true, modeFull},
	}
	for _, tc := range cases {
		if got := SelectMode(tc.state, tc.mirrorLost, now, tc.force); got != tc.want {
			t.Errorf("%s: got %q, want %q", tc.name, got, tc.want)
		}
	}
}

func wantStats(t *testing.T, got, want Stats) {
	t.Helper()
	if got != want {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}

func wantState(t *testing.T, home string, lastFullAt, watermark time.Time) {
	t.Helper()
	state, err := LoadState(home)
	if err != nil {
		t.Fatalf("LoadState: %v", err)
	}
	if state == nil || !state.LastFullAt.Equal(lastFullAt) || !state.Watermark.Equal(watermark) {
		t.Fatalf("got %+v, want last_full=%s watermark=%s", state, lastFullAt, watermark)
	}
}

func TestFullRun(t *testing.T) {
	home := t.TempDir()
	api := &fakeAPI{
		pages: map[string]model.PageMeta{
			"p1": {ID: "p1", Title: "Alpha", LastEdited: editedOld},
			"p2": {ID: "p2", Title: "Beta", LastEdited: editedNew},
		},
	}

	stats, err := Run(api, Options{DataSourceID: "ds", Home: home, Now: now})
	if err != nil {
		t.Fatalf("Run: %v", err)
	}

	wantStats(t, stats, Stats{Mode: modeFull, Kept: 2, Written: 2})
	if n := countMD(t, home); n != 2 {
		t.Fatalf("got %d mirror files, want 2", n)
	}
	wantState(t, home, now, now)
}

func TestIncrementalRun(t *testing.T) {
	home := t.TempDir()
	api := &fakeAPI{
		pages: map[string]model.PageMeta{
			"p1": {ID: "p1", Title: "Alpha", LastEdited: editedOld},
			"p2": {ID: "p2", Title: "Beta", LastEdited: editedOld},
		},
	}
	if _, err := Run(api, Options{Home: home, Now: now}); err != nil {
		t.Fatalf("seed full run: %v", err)
	}

	later := now.Add(time.Hour)
	api.pages["p2"] = model.PageMeta{ID: "p2", Title: "Beta 2", LastEdited: editedNew}
	stats, err := Run(api, Options{Home: home, Now: later})
	if err != nil {
		t.Fatalf("incremental run: %v", err)
	}

	wantStats(t, stats, Stats{Mode: modeIncremental, Kept: 1, Written: 1})
	wantFilter := now.Add(-OverlapMargin).Format(time.RFC3339)
	if api.queryFilter.LastEditedAfter != wantFilter {
		t.Fatalf("filter = %q, want %q", api.queryFilter.LastEditedAfter, wantFilter)
	}

	wantState(t, home, now, later)
	if n := countMD(t, home); n != 2 {
		t.Fatalf("got %d mirror files, want 2", n)
	}
}

func TestIncrementalRenameKeepsSingleFile(t *testing.T) {
	home := t.TempDir()
	api := &fakeAPI{pages: map[string]model.PageMeta{
		"p1": {ID: "p1", Title: "Old Title", LastEdited: editedOld},
	}}
	if _, err := Run(api, Options{Home: home, Now: now}); err != nil {
		t.Fatalf("seed full run: %v", err)
	}

	api.pages["p1"] = model.PageMeta{ID: "p1", Title: "New Title", LastEdited: editedNew}
	if _, err := Run(api, Options{Home: home, Now: now.Add(time.Hour)}); err != nil {
		t.Fatalf("incremental run: %v", err)
	}

	if n := countMD(t, home); n != 1 {
		t.Fatalf("got %d mirror files, want 1", n)
	}
	data, err := os.ReadFile(filepath.Join(home, "knowledge", "old-title--p1.md"))
	if err != nil {
		t.Fatalf("renamed file kept old slug but moved: %v", err)
	}
	if !strings.Contains(string(data), "New Title") {
		t.Fatalf("file lacks new title:\n%s", data)
	}
}

func TestFullRunSweepsDeletedPages(t *testing.T) {
	home := t.TempDir()
	api := &fakeAPI{pages: map[string]model.PageMeta{
		"p1": {ID: "p1", Title: "Alpha", LastEdited: editedOld},
		"p2": {ID: "p2", Title: "Beta", LastEdited: editedNew},
	}}
	if _, err := Run(api, Options{Home: home, Now: now}); err != nil {
		t.Fatalf("seed full run: %v", err)
	}

	delete(api.pages, "p2")
	stats, err := Run(api, Options{Home: home, Now: now.Add(time.Hour), ForceFull: true})
	if err != nil {
		t.Fatalf("sweeping run: %v", err)
	}

	if stats.Removed != 1 {
		t.Fatalf("removed %d files, want 1", stats.Removed)
	}
	if n := countMD(t, home); n != 1 {
		t.Fatalf("got %d mirror files, want 1", n)
	}
}

func TestPartialFailureKeepsWatermark(t *testing.T) {
	home := t.TempDir()
	old := State{LastFullAt: now.Add(-48 * time.Hour), Watermark: now.Add(-24 * time.Hour)}
	if err := SaveState(home, old); err != nil {
		t.Fatalf("seed state: %v", err)
	}

	api := &fakeAPI{
		pages: map[string]model.PageMeta{
			"p1": {ID: "p1", Title: "Alpha", LastEdited: editedOld},
			"p2": {ID: "p2", Title: "Beta", LastEdited: editedNew},
		},
		errPage: "p2",
	}

	stats, err := Run(api, Options{Home: home, Now: now, ForceFull: true})
	if err == nil {
		t.Fatal("Run succeeded despite page failure")
	}
	if !strings.Contains(err.Error(), "1 of 2 pages failed") {
		t.Fatalf("error = %q, want page-failure summary", err)
	}
	if stats.Failed != 1 || stats.Written != 1 {
		t.Fatalf("unexpected stats: %+v", stats)
	}

	state, err := LoadState(home)
	if err != nil {
		t.Fatalf("LoadState: %v", err)
	}
	if state == nil || !state.LastFullAt.Equal(old.LastFullAt) || !state.Watermark.Equal(old.Watermark) {
		t.Fatalf("watermark advanced despite failure: %+v", state)
	}
}

func TestQueryFailureFailsRunAndKeepsState(t *testing.T) {
	home := t.TempDir()
	if err := SaveState(home, State{LastFullAt: now, Watermark: now}); err != nil {
		t.Fatalf("seed state: %v", err)
	}

	api := &queryErrAPI{err: errors.New("notion rejected the credentials")}
	_, err := Run(api, Options{Home: home, Now: now})
	if err == nil {
		t.Fatal("Run succeeded despite query failure")
	}

	wantState(t, home, now, now) // the watermark must not advance
}

func TestRunLockedHome(t *testing.T) {
	home := t.TempDir()
	f, err := os.OpenFile(filepath.Join(home, lockFile), os.O_CREATE|os.O_RDWR, 0o600)
	if err != nil {
		t.Fatalf("open lock: %v", err)
	}
	defer func() { _ = f.Close() }()
	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX|syscall.LOCK_NB); err != nil {
		t.Fatalf("hold lock: %v", err)
	}

	_, err = Run(&fakeAPI{}, Options{Home: home, Now: now})
	if err == nil {
		t.Fatal("Run succeeded on locked home")
	}
	if !strings.Contains(err.Error(), "sync already in progress") {
		t.Fatalf("error = %q, want lock message", err)
	}
}

func countMD(t *testing.T, home string) int {
	t.Helper()
	matches, err := filepath.Glob(filepath.Join(home, "knowledge", "*.md"))
	if err != nil {
		t.Fatalf("glob mirror: %v", err)
	}

	return len(matches)
}
