package sync

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/iyaki/enchiridion/internal/model"
)

func TestSlugify(t *testing.T) {
	cases := []struct{ in, want string }{
		{"El desafío del lenguaje ubicuo", "el-desafio-del-lenguaje-ubicuo"},
		{"  Hello, World!  ", "hello-world"},
		{"", "untitled"},
		{"¿Qué es DDD?", "que-es-ddd"},
		{"---", "untitled"},
	}
	for _, tc := range cases {
		if got := Slugify(tc.in); got != tc.want {
			t.Errorf("Slugify(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

func TestPageFileName(t *testing.T) {
	meta := model.PageMeta{ID: "a1b2c3d4-e5f6-7890", Title: "Mi Página"}
	if got, want := PageFileName(meta), "mi-pagina--a1b2c3d4.md"; got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestFrontmatter(t *testing.T) {
	meta := model.PageMeta{
		ID:         "a1b2c3d4-e5f6",
		NotionURL:  "https://notion.so/a1b2c3d4",
		Title:      `He said "hi"`,
		LastEdited: "2026-09-20T10:00:00.000Z",
		SourceURL:  "https://src.example.com/post",
		Tags:       []string{"article", "ddd"},
	}

	want := "---\n" +
		"title: \"He said \\\"hi\\\"\"\n" +
		"notion_id: a1b2c3d4-e5f6\n" +
		"notion_url: https://notion.so/a1b2c3d4\n" +
		"last_edited: 2026-09-20T10:00:00.000Z\n" +
		"source_url: https://src.example.com/post\n" +
		"tags: [\"article\", \"ddd\"]\n" +
		"---"
	if got := Frontmatter(meta); got != want {
		t.Fatalf("got:\n%s\nwant:\n%s", got, want)
	}
}

func TestFrontmatterOmitsEmptyOptionalFields(t *testing.T) {
	meta := model.PageMeta{ID: "id1", Title: "t", LastEdited: "ts"}
	got := Frontmatter(meta)
	if strings.Contains(got, "source_url") || strings.Contains(got, "tags") {
		t.Fatalf("optional fields present without values:\n%s", got)
	}
}

func TestPageContent(t *testing.T) {
	meta := model.PageMeta{ID: "id1", Title: "t", LastEdited: "ts"}
	blocks := []model.Block{{Type: model.TypeParagraph, RichText: []model.RichText{{Plain: "body"}}}}

	want := Frontmatter(meta) + "\nbody\n"
	if got := PageContent(meta, blocks); got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestWritePageRenameSafe(t *testing.T) {
	dir := t.TempDir()
	meta := model.PageMeta{ID: "aaaa1111", Title: "Old Title", LastEdited: "ts1"}

	first, err := WritePage(dir, meta, nil)
	if err != nil {
		t.Fatalf("first write: %v", err)
	}

	meta.Title = "New Title"
	second, err := WritePage(dir, meta, nil)
	if err != nil {
		t.Fatalf("rewrite: %v", err)
	}

	if first != second {
		t.Fatalf("rename produced a second path: %q vs %q", first, second)
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatalf("read dir: %v", err)
	}
	if len(entries) != 1 {
		t.Fatalf("got %d files, want 1", len(entries))
	}

	data, err := os.ReadFile(second)
	if err != nil {
		t.Fatalf("read rewritten file: %v", err)
	}
	if !strings.Contains(string(data), "New Title") {
		t.Fatalf("rewritten file lacks new title:\n%s", data)
	}
}

func TestSweep(t *testing.T) {
	dir := t.TempDir()
	kept := "kept.md"
	for name, content := range map[string]string{
		kept:       "---\nnotion_id: keep-1\n---\n",
		"stale.md": "---\nnotion_id: gone-1\n---\n",
		"noid.md":  "no frontmatter here",
	} {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0o600); err != nil {
			t.Fatalf("setup %s: %v", name, err)
		}
	}

	removed, err := Sweep(dir, map[string]bool{"keep-1": true})
	if err != nil {
		t.Fatalf("Sweep: %v", err)
	}

	wantRemoved := []string{"stale.md", "noid.md"}
	if len(removed) != len(wantRemoved) {
		t.Fatalf("removed %v, want %v", removed, wantRemoved)
	}
	if _, err := os.Stat(filepath.Join(dir, kept)); err != nil {
		t.Fatalf("kept file was removed: %v", err)
	}
}
