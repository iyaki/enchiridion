package search

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// writeMirror creates .md files under a temp root and returns the root.
func writeMirror(t *testing.T, files map[string]string) string {
	t.Helper()
	root := t.TempDir()
	for rel, content := range files {
		path := filepath.Join(root, filepath.FromSlash(rel))
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatalf("mkdir: %v", err)
		}
		if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
			t.Fatalf("write %s: %v", rel, err)
		}
	}

	return root
}

// fm renders the same frontmatter shape the sync engine emits
// (internal/sync/pagefile.go Frontmatter).
func fm(title string, tags ...string) string {
	s := "---\ntitle: \"" + title + "\"\nnotion_id: id\nnotion_url: https://notion.so/id\nlast_edited: ts\n"
	if len(tags) > 0 {
		quoted := make([]string, len(tags))
		for i, tag := range tags {
			quoted[i] = `"` + tag + `"`
		}
		s += "tags: [" + strings.Join(quoted, ", ") + "]\n"
	}

	return s + "---\n"
}

func TestRunRanksTitleAboveBody(t *testing.T) {
	root := writeMirror(t, map[string]string{
		"knowledge/a.md": fm("Accessibility auditing") + "unrelated body\n",
		"knowledge/b.md": fm("Something else") + "talks about accessibility here\n",
	})

	hits, err := Run(root, []string{"accessibility"})
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	if len(hits) != 2 {
		t.Fatalf("got %d hits, want 2", len(hits))
	}
	if hits[0].Path != "knowledge/a.md" {
		t.Fatalf("first hit %q, want knowledge/a.md (title match outranks body)", hits[0].Path)
	}
	if hits[0].Title != "Accessibility auditing" {
		t.Fatalf("title %q, want the frontmatter title", hits[0].Title)
	}
}

func TestRunANDSemanticsCaseInsensitive(t *testing.T) {
	root := writeMirror(t, map[string]string{
		"knowledge/both.md": fm("Go Testing") + "this one covers RACE conditions\n",
		"knowledge/one.md":  fm("Go Testing alone") + "nothing else here\n",
	})

	hits, err := Run(root, []string{"go", "RACE"})
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	if len(hits) != 1 || hits[0].Path != "knowledge/both.md" {
		t.Fatalf("got %v, want only knowledge/both.md (every term must match)", hits)
	}
}

func TestRunTagsWeight(t *testing.T) {
	root := writeMirror(t, map[string]string{
		"knowledge/tagged.md": fm("Page One", "newsletter") + "unrelated body\n",
		"knowledge/body.md":   fm("Page Two") + "mentions newsletter once\n",
	})

	hits, err := Run(root, []string{"newsletter"})
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	if len(hits) != 2 {
		t.Fatalf("got %d hits, want 2", len(hits))
	}
	if hits[0].Path != "knowledge/tagged.md" {
		t.Fatalf("first hit %q, want knowledge/tagged.md (tags match outranks body)", hits[0].Path)
	}
}

func TestRunNoFrontmatterFallback(t *testing.T) {
	root := writeMirror(t, map[string]string{
		"notes/plain.md": "just some prose about ferrets\n",
	})

	hits, err := Run(root, []string{"ferrets"})
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	if len(hits) != 1 {
		t.Fatalf("got %d hits, want 1", len(hits))
	}
	if hits[0].Title != "plain" {
		t.Fatalf("title %q, want the filename base without extension", hits[0].Title)
	}
}

func TestRunDeterministicOrder(t *testing.T) {
	root := writeMirror(t, map[string]string{
		"knowledge/z.md": fm("Zed") + "dungeon\n",
		"knowledge/a.md": fm("Ayy") + "dungeon\n",
	})

	hits, err := Run(root, []string{"dungeon"})
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	if len(hits) != 2 || hits[0].Path != "knowledge/a.md" || hits[1].Path != "knowledge/z.md" {
		t.Fatalf("got %v, want paths sorted ascending at equal score", hits)
	}
}

func TestRunNoMatchesNoError(t *testing.T) {
	root := writeMirror(t, map[string]string{
		"knowledge/a.md": fm("A") + "body\n",
	})

	hits, err := Run(root, []string{"zzzznoterm"})
	if err != nil || hits != nil {
		t.Fatalf("got (%v, %v), want (nil, nil)", hits, err)
	}

	if hits, err := Run(root, nil); err != nil || hits != nil {
		t.Fatalf("no terms: got (%v, %v), want (nil, nil)", hits, err)
	}
}
