// Package sync mirrors the Notion knowledge base into a local markdown cache:
// watermark state, page file management, and mode-selected sync runs
// (specs/architecture.md).
package sync

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

// stateFile travels committed alongside the mirror so a fresh clone
// increments correctly (specs/architecture.md — sync state).
const stateFile = ".sync-state.json"

// State is the sync watermark, committed alongside the mirror.
type State struct {
	LastFullAt time.Time `json:"last_full_at"`
	Watermark  time.Time `json:"watermark"`
}

// LoadState reads the watermark from home. A missing or corrupt file yields
// (nil, nil): mode selection degrades to a full sync (auto-backfill) instead
// of blocking the run.
func LoadState(home string) (*State, error) {
	data, err := os.ReadFile(filepath.Join(home, stateFile))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}

		return nil, err
	}

	var s State
	if err := json.Unmarshal(data, &s); err != nil {
		return nil, nil
	}

	return &s, nil
}

// SaveState writes the watermark into home with restricted permissions.
func SaveState(home string, s State) error {
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(home, stateFile), append(data, '\n'), 0o600)
}
