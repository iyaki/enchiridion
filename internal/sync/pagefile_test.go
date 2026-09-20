package sync

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/iyaki/enchiridion/internal/model"
)

// idOfPath reads the notion_id of a mirror file by path (test helper over
// the root-based idOf).
func idOfPath(t *testing.T, path string) string {
	t.Helper()
	root, err := os.OpenRoot(filepath.Dir(path))
	if err != nil {
		t.Fatalf("open root of %q: %v", path, err)
	}
	defer func() { _ = root.Close() }()

	return idOf(root, filepath.Base(path))
}

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

func TestPageFileNameCapsSlugLength(t *testing.T) {
	longTitle := strings.Repeat("word ", 60)
	meta := model.PageMeta{ID: "a1b2c3d4-e5f6-7890", Title: longTitle}
	got := PageFileName(meta)
	if len(got) > 255 {
		t.Fatalf("file name exceeds filesystem limit: %d bytes", len(got))
	}
	if !strings.HasSuffix(got, "--a1b2c3d4.md") {
		t.Fatalf("truncated name lost the id suffix: %q", got)
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

func TestWritePageIDPrefixCollision(t *testing.T) {
	dir := t.TempDir()
	// Time-ordered page IDs can share their first characters across pages
	// created in the same batch; a shared slug narrows the name further.
	pages := []model.PageMeta{
		{ID: "2b754f1c-aaaa-1111", Title: "Same Slug", LastEdited: "ts"},
		{ID: "2b754f1c-bbbb-2222", Title: "Same Slug", LastEdited: "ts"},
		{ID: "2b754f1c-cccc-3333", Title: "Same Slug", LastEdited: "ts"},
	}

	paths := make(map[string]string) // notion_id -> written path
	for _, meta := range pages {
		path, err := WritePage(dir, meta, nil)
		if err != nil {
			t.Fatalf("write %s: %v", meta.ID, err)
		}
		paths[meta.ID] = path
	}

	if len(paths) != len(pages) {
		t.Fatalf("got %d paths, want %d", len(paths), len(pages))
	}
	for id, path := range paths {
		if got := idOfPath(t, path); got != id {
			t.Fatalf("path %q belongs to %q, want %q", path, got, id)
		}
	}

	// Rewrites must land on the same paths (stable across syncs) and no file
	// may be duplicated.
	rewriteStable(t, dir, pages, paths)
}

func rewriteStable(t *testing.T, dir string, pages []model.PageMeta, paths map[string]string) {
	t.Helper()
	for _, meta := range pages {
		path, err := WritePage(dir, meta, nil)
		if err != nil {
			t.Fatalf("rewrite %s: %v", meta.ID, err)
		}
		if path != paths[meta.ID] {
			t.Fatalf("rewrite of %s moved the file: %q -> %q", meta.ID, paths[meta.ID], path)
		}
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatalf("read dir: %v", err)
	}
	if len(entries) != len(pages) {
		t.Fatalf("got %d files, want %d", len(entries), len(pages))
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
