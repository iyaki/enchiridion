package sync

import "github.com/iyaki/enchiridion/internal/model"

const (
	dirKnowledge = "knowledge"
	dirTools     = "tools"
)

// Category values as they arrive in Tags (confirmed against the live mirror).
var (
	knowledgeCategories = map[string]bool{"Article": true, "Note": true}
	toolCategories      = map[string]bool{
		"Tool": true, "Service": true, "Website": true,
		"Framework/Library": true, "Game": true,
	}
)

// Class returns the mirror directory name for a page (ADR-15). Knowledge
// wins: a page carrying any knowledge category is knowledge even when it
// also carries tool categories; unknown or absent categories default to
// knowledge.
func Class(meta model.PageMeta) string {
	for _, tag := range meta.Tags {
		if knowledgeCategories[tag] {
			return dirKnowledge
		}
	}
	for _, tag := range meta.Tags {
		if toolCategories[tag] {
			return dirTools
		}
	}

	return dirKnowledge
}
