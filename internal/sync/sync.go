package sync

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"time"

	"github.com/iyaki/enchiridion/internal/model"
)

const (
	// FullInterval is the age after which the last full sync forces another
	// one: full mode re-renders every page under the current renderer and
	// resets the last-full date (both modes propagate deletions, ADR-21).
	FullInterval = 30 * 24 * time.Hour
	// OverlapMargin re-synchronizes this much before the watermark so clock
	// drift between Notion and the local machine never skips an edit.
	OverlapMargin = time.Minute

	lockFile        = ".sync.lock"
	modeFull        = "full"
	modeIncremental = "incremental"
)

// NotionAPI is the client surface the engine consumes.
type NotionAPI interface {
	QueryPages(dataSourceID string) ([]model.PageMeta, error)
	PageBlocks(pageID string) ([]model.Block, error)
}

// Options configures one sync run.
type Options struct {
	DataSourceID string
	Home         string    // resolved cache root
	ForceFull    bool      // --full
	Now          time.Time // single consistent instant for the run
	// Progress, when set, reports pages processed (done, total); done
	// counts failures too. Nil means silent.
	Progress func(done, total int)
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
	dirs, err := classDirs(opts.Home)
	if err != nil {
		return Stats{}, err
	}
	mode := SelectMode(state, mirrorIsEmpty(dirs[dirKnowledge]) && mirrorIsEmpty(dirs[dirTools]), opts.Now, opts.ForceFull)

	metas, err := query(api, opts)
	if err != nil {
		return Stats{}, err
	}

	toWrite := metas
	if mode == modeIncremental {
		toWrite = editedSince(metas, state.Watermark.Add(-OverlapMargin))
	}

	stats := syncAll(api, dirs, toWrite, opts.Progress)
	stats.Mode = mode
	keptByDir := keptSets(metas)
	for _, name := range []string{dirKnowledge, dirTools} {
		removed, sweepErr := Sweep(dirs[name], keptByDir[name])
		if sweepErr != nil {
			return stats, sweepErr
		}
		stats.Removed += len(removed)
	}

	if stats.Failed > 0 {
		return stats, fmt.Errorf("%d of %d pages failed", stats.Failed, len(metas))
	}

	if err := persistState(opts, state, mode); err != nil {
		return stats, err
	}

	return stats, nil
}

// classDirs resolves both class directories under home, creating them so a
// first full-mode sweep of an old-layout cache cannot fail on a missing dir.
func classDirs(home string) (map[string]string, error) {
	dirs := map[string]string{
		dirKnowledge: filepath.Join(home, dirKnowledge),
		dirTools:     filepath.Join(home, dirTools),
	}
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, dirPerm); err != nil {
			return nil, err
		}
	}

	return dirs, nil
}

// query fetches the full page manifest: one unfiltered metadata query per run
// (~⌈N/100⌉ requests, ADR-21). What to rewrite is selected client-side.
func query(api NotionAPI, opts Options) ([]model.PageMeta, error) {
	return api.QueryPages(opts.DataSourceID)
}

// editedSince keeps pages edited at or after cutoff; a page whose timestamp
// fails to parse is kept (conservative: rewrite it).
func editedSince(metas []model.PageMeta, cutoff time.Time) []model.PageMeta {
	var out []model.PageMeta
	for _, meta := range metas {
		edited, err := time.Parse(time.RFC3339, meta.LastEdited)
		if err != nil || !edited.Before(cutoff) {
			out = append(out, meta)
		}
	}

	return out
}

// keptSets maps, per class directory, the set of page IDs the manifest says
// belong there: the sweep baseline (ADR-21).
func keptSets(metas []model.PageMeta) map[string]map[string]bool {
	kept := map[string]map[string]bool{dirKnowledge: {}, dirTools: {}}
	for _, meta := range metas {
		kept[Class(meta)][meta.ID] = true
	}

	return kept
}

// syncAll writes every page in metas, logging and counting failures without
// stopping (specs/architecture.md — error behavior). progress (nullable) is
// called once before the loop and after every page; done counts processed
// pages, failures included.
func syncAll(api NotionAPI, dirs map[string]string, metas []model.PageMeta,
	progress func(done, total int),
) Stats {
	stats := Stats{Kept: len(metas)}
	if progress != nil {
		progress(0, len(metas))
	}
	for i, meta := range metas {
		if err := writeOne(api, dirs[Class(meta)], meta); err != nil {
			fmt.Fprintf(os.Stderr, "sync: page %s: %v\n", meta.ID, err)
			stats.Failed++
		} else {
			stats.Written++
		}
		if progress != nil {
			progress(i+1, len(metas))
		}
	}

	return stats
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

// writeOne fetches a page's blocks and writes it into the mirror.
func writeOne(api NotionAPI, dir string, meta model.PageMeta) error {
	blocks, err := api.PageBlocks(meta.ID)
	if err != nil {
		return err
	}

	_, err = WritePage(dir, meta, blocks)

	return err
}

// lock acquires exclusive access to the cache so concurrent runs (local +
// cron) cannot corrupt the state (specs/architecture.md — exclusivity).
func lock(home string) (func(), error) {
	root, err := os.OpenRoot(home)
	if err != nil {
		return nil, err
	}

	f, err := root.OpenFile(lockFile, os.O_CREATE|os.O_RDWR, filePerm)
	if err != nil {
		_ = root.Close()

		return nil, err
	}

	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX|syscall.LOCK_NB); err != nil {
		_ = f.Close()
		_ = root.Close()

		return nil, fmt.Errorf("sync already in progress (locked %s)", filepath.Join(home, lockFile))
	}

	return func() {
		_ = syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
		_ = f.Close()
		_ = root.Close()
	}, nil
}

// mirrorIsEmpty reports whether the mirror has no pages yet.
func mirrorIsEmpty(dir string) bool {
	matches, _ := filepath.Glob(filepath.Join(dir, "*.md"))

	return len(matches) == 0
}
