package sync

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"

	"github.com/iyaki/enchiridion/internal/model"
	"github.com/iyaki/enchiridion/internal/render"
)

const (
	idPrefixLen = 8
	untitled    = "untitled"
	dirPerm     = 0o755
	filePerm    = 0o600
	// maxSlugLen keeps file names well under the 255-byte filesystem limit
	// even for pathologically long titles.
	maxSlugLen = 100
)

// Accented Latin letters folded to ASCII. // ponytail: explicit table because
// NFD normalization needs golang.org/x/text, barred by ADR-08; covers the KB's
// languages, extend if titles ever leave Latin-1.
var accents = strings.NewReplacer(
	"à", "a", "á", "a", "â", "a", "ã", "a", "ä", "a", "å", "a",
	"è", "e", "é", "e", "ê", "e", "ë", "e",
	"ì", "i", "í", "i", "î", "i", "ï", "i",
	"ò", "o", "ó", "o", "ô", "o", "õ", "o", "ö", "o",
	"ù", "u", "ú", "u", "û", "u", "ü", "u",
	"ý", "y", "ÿ", "y", "ñ", "n", "ç", "c",
)

var nonSlug = regexp.MustCompile(`[^a-z0-9]+`)

// Slugify converts a title into the kebab-case slug used in mirror file names
// (specs/architecture.md — mirror format): lowercase, no accents, ASCII only.
func Slugify(title string) string {
	s := accents.Replace(strings.ToLower(title))
	s = strings.Map(func(r rune) rune {
		if unicode.Is(unicode.Mn, r) {
			return -1 // already-decomposed marks (NFD input)
		}

		return r
	}, s)
	s = strings.Trim(nonSlug.ReplaceAllString(s, "-"), "-")
	if s == "" {
		return untitled
	}

	return s
}

// PageFileName is the mirror file name for a page: {slug}--{id8}.md. The slug
// is capped so over-long titles cannot exceed filesystem limits.
func PageFileName(meta model.PageMeta) string {
	return fileName(meta, idPrefixLen)
}

// fileName builds the mirror file name with an explicit id prefix length.
func fileName(meta model.PageMeta, idLen int) string {
	id := meta.ID
	if len(id) > idLen {
		id = id[:idLen]
	}

	slug := Slugify(meta.Title)
	if len(slug) > maxSlugLen {
		slug = slug[:maxSlugLen]
	}

	return slug + "--" + id + ".md"
}

// quote escapes a value for its double-quoted frontmatter field.
func quote(s string) string {
	return strings.ReplaceAll(s, `"`, `\"`)
}

// Frontmatter renders the page identity block (specs/architecture.md — mirror
// format). Optional fields appear only when they carry a value.
func Frontmatter(meta model.PageMeta) string {
	var b strings.Builder
	b.WriteString("---\n")
	fmt.Fprintf(&b, "title: \"%s\"\n", quote(meta.Title))
	fmt.Fprintf(&b, "notion_id: %s\n", meta.ID)
	fmt.Fprintf(&b, "notion_url: %s\n", meta.NotionURL)
	fmt.Fprintf(&b, "last_edited: %s\n", meta.LastEdited)
	if meta.SourceURL != "" {
		fmt.Fprintf(&b, "source_url: %s\n", meta.SourceURL)
	}
	if len(meta.Tags) > 0 {
		quoted := make([]string, len(meta.Tags))
		for i, tag := range meta.Tags {
			quoted[i] = `"` + quote(tag) + `"`
		}
		fmt.Fprintf(&b, "tags: [%s]\n", strings.Join(quoted, ", "))
	}
	b.WriteString("---")

	return b.String()
}

// PageContent is the full mirror document for a page: frontmatter, body,
// trailing newline.
func PageContent(meta model.PageMeta, blocks []model.Block) string {
	return Frontmatter(meta) + "\n" + render.Markdown(blocks) + "\n"
}

// WritePage writes the page body into the mirror. If a file for the same
// notion_id already exists it is rewritten in place — even under a new title —
// so renames never accumulate duplicates (specs/architecture.md — identity
// consistency). Returns the written path.
func WritePage(dir string, meta model.PageMeta, blocks []model.Block) (string, error) {
	if err := os.MkdirAll(dir, dirPerm); err != nil {
		return "", err
	}

	root, err := os.OpenRoot(dir)
	if err != nil {
		return "", err
	}
	defer func() { _ = root.Close() }()

	// ponytail: linear scan per page; build an id->path index if the KB grows
	// past a few thousand entries.
	entries, err := os.ReadDir(dir)
	if err != nil {
		return "", err
	}
	for _, entry := range entries {
		name := entry.Name()
		if isMarkdown(name) && idOf(root, name) == meta.ID {
			path := filepath.Join(dir, name)

			return path, os.WriteFile(path, []byte(PageContent(meta, blocks)), filePerm)
		}
	}

	// Time-ordered page IDs can share their first characters across pages
	// created in the same batch, so the short prefix may not name this page;
	// extend it until the name is unambiguous.
	name := PageFileName(meta)
	for idLen := idPrefixLen + 1; idTaken(root, name, meta.ID); idLen++ {
		if idLen > len(meta.ID) {
			break
		}

		name = fileName(meta, idLen)
	}

	path := filepath.Join(dir, name)

	return path, os.WriteFile(path, []byte(PageContent(meta, blocks)), filePerm)
}

// isMarkdown reports whether a directory entry is a mirror page file.
func isMarkdown(name string) bool {
	return strings.HasSuffix(name, ".md")
}

// idTaken reports whether name exists in root and belongs to a different page.
func idTaken(root *os.Root, name, id string) bool {
	owner := idOf(root, name)

	return owner != "" && owner != id
}

// idOf returns the notion_id declared in a mirror file's frontmatter, or ""
// when absent.
func idOf(root *os.Root, name string) string {
	f, err := root.Open(name)
	if err != nil {
		return ""
	}
	defer func() { _ = f.Close() }()

	data, err := io.ReadAll(f)
	if err != nil {
		return ""
	}

	for _, line := range strings.Split(string(data), "\n") {
		if rest, ok := strings.CutPrefix(line, "notion_id: "); ok {
			return strings.TrimSpace(rest)
		}
	}

	return ""
}

// Sweep removes every mirror file whose page no longer exists in the source
// (or carries no notion_id at all) and returns the removed file names. Full
// mode only: it is the sole propagator of deletions (ADR-04).
func Sweep(dir string, keptIDs map[string]bool) ([]string, error) {
	root, err := os.OpenRoot(dir)
	if err != nil {
		return nil, err
	}
	defer func() { _ = root.Close() }()

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var removed []string
	for _, entry := range entries {
		name := entry.Name()
		if !isMarkdown(name) || keptIDs[idOf(root, name)] {
			continue
		}

		if err := root.Remove(name); err != nil {
			return removed, err
		}
		removed = append(removed, name)
	}

	return removed, nil
}
