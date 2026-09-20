package sync

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"time"

	"github.com/iyaki/enchiridion/internal/model"
	"github.com/iyaki/enchiridion/internal/notion"
)

const (
	// FullInterval is the age after which the last full sync forces another
	// one: the only mode that propagates deletions (ADR-04).
	FullInterval = 30 * 24 * time.Hour
	// OverlapMargin re-synchronizes this much before the watermark so clock
	// drift between Notion and the local machine never skips an edit.
	OverlapMargin = time.Minute

	knowledgeDir    = "knowledge"
	lockFile        = ".sync.lock"
	modeFull        = "full"
	modeIncremental = "incremental"
)

// NotionAPI is the client surface the engine consumes.
type NotionAPI interface {
	QueryPages(dataSourceID string, filter notion.QueryFilter) ([]model.PageMeta, error)
	PageBlocks(pageID string) ([]model.Block, error)
}

// Options configures one sync run.
type Options struct {
	DataSourceID string
	Home         string    // resolved cache root
	ForceFull    bool      // --full
	Now          time.Time // single consistent instant for the run
}

// Stats summarizes one run.
type Stats struct {
	Mode    string // "full" | "incremental"
	Kept    int
	Written int
	Removed int
	Failed  int
}

// SelectMode picks the sync mode for a run (specs/architecture.md — sync
// modes, ADR-07). Pure.
func SelectMode(state *State, mirrorEmpty bool, now time.Time, forceFull bool) string {
	switch {
	case forceFull:
		return modeFull
	case state == nil || mirrorEmpty:
		return modeFull // auto-backfill: missing data resolves itself
	case now.Sub(state.LastFullAt) > FullInterval:
		return modeFull
	default:
		return modeIncremental
	}
}

// Run synchronizes the mirror under home and advances the watermark. The
// watermark is written only when every page succeeded, so a partial run
// retries every edit since the old mark next time (idempotent writes make
// that safe). Per-page failures are logged and do not stop the run.
func Run(api NotionAPI, opts Options) (Stats, error) {
	release, err := lock(opts.Home)
	if err != nil {
		return Stats{}, err
	}
	defer release()

	state, _ := LoadState(opts.Home)
	knowledge := filepath.Join(opts.Home, knowledgeDir)
	mode := SelectMode(state, mirrorIsEmpty(knowledge), opts.Now, opts.ForceFull)

	metas, err := query(api, opts, state, mode)
	if err != nil {
		return Stats{}, err
	}

	stats, keptIDs := syncAll(api, knowledge, mode, metas)
	if mode == modeFull {
		removed, sweepErr := Sweep(knowledge, keptIDs)
		if sweepErr != nil {
			return stats, sweepErr
		}
		stats.Removed = len(removed)
	}

	if stats.Failed > 0 {
		return stats, fmt.Errorf("%d of %d pages failed", stats.Failed, len(metas))
	}

	if err := persistState(opts, state, mode); err != nil {
		return stats, err
	}

	return stats, nil
}

// query fetches page metadata for the mode: everything on full, only pages
// edited since the watermark (minus the overlap margin) on incremental.
func query(api NotionAPI, opts Options, state *State, mode string) ([]model.PageMeta, error) {
	filter := notion.QueryFilter{}
	if mode == modeIncremental {
		filter.LastEditedAfter = state.Watermark.Add(-OverlapMargin).Format(time.RFC3339)
	}

	return api.QueryPages(opts.DataSourceID, filter)
}

// syncAll writes every page, logging and counting failures without stopping
// (specs/architecture.md — error behavior). Returns the stats and the set of
// page IDs that must remain in the mirror.
func syncAll(api NotionAPI, knowledge string, mode string, metas []model.PageMeta) (Stats, map[string]bool) {
	stats := Stats{Mode: mode, Kept: len(metas)}
	keptIDs := make(map[string]bool, len(metas))
	for _, meta := range metas {
		keptIDs[meta.ID] = true
		if err := writeOne(api, knowledge, meta); err != nil {
			fmt.Fprintf(os.Stderr, "sync: page %s: %v\n", meta.ID, err)
			stats.Failed++

			continue
		}
		stats.Written++
	}

	return stats, keptIDs
}

// persistState advances the watermark to the run instant; incremental runs
// keep the previous full-sync date.
func persistState(opts Options, state *State, mode string) error {
	newState := State{LastFullAt: opts.Now, Watermark: opts.Now}
	if mode == modeIncremental && state != nil {
		newState.LastFullAt = state.LastFullAt
	}

	return SaveState(opts.Home, newState)
}

// writeOne fetches a page's blocks and writes it into the mirror, printing
// the mirrored file name.
func writeOne(api NotionAPI, knowledge string, meta model.PageMeta) error {
	blocks, err := api.PageBlocks(meta.ID)
	if err != nil {
		return err
	}

	path, err := WritePage(knowledge, meta, blocks)
	if err != nil {
		return err
	}
	fmt.Printf("synced: %s\n", filepath.Base(path))

	return nil
}

// lock acquires exclusive access to the cache so concurrent runs (local +
// cron) cannot corrupt the state (specs/architecture.md — exclusivity).
func lock(home string) (func(), error) {
	f, err := os.OpenFile(filepath.Join(home, lockFile), os.O_CREATE|os.O_RDWR, filePerm)
	if err != nil {
		return nil, err
	}

	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX|syscall.LOCK_NB); err != nil {
		_ = f.Close()

		return nil, fmt.Errorf("sync already in progress (locked %s)", filepath.Join(home, lockFile))
	}

	return func() {
		_ = syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
		_ = f.Close()
	}, nil
}

// mirrorIsEmpty reports whether the mirror has no pages yet.
func mirrorIsEmpty(dir string) bool {
	matches, _ := filepath.Glob(filepath.Join(dir, "*.md"))

	return len(matches) == 0
}
