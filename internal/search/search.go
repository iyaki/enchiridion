// Package search ranks mirror markdown files by term matches: the caller
// (agent or human) judges relevance, the engine only orders candidates.
package search

import (
	"cmp"
	"io/fs"
	"os"
	"slices"
	"strings"
)

// Field weights: a field counts once per term it matches.
const (
	weightTitle    = 3
	weightTags     = 2
	weightFilename = 1
	weightBody     = 1
)

// Frontmatter markers, matching what internal/sync emits.
const (
	fmOpen  = "---\n"
	fmClose = "---"
)

// Hit is one matching mirror file.
type Hit struct {
	Path  string // path relative to the search root, slash-joined
	Title string // frontmatter title, or filename base when absent
	Score int
}

// fields holds the lowered searchable surfaces of one file.
type fields struct {
	title    string
	tags     string
	filename string
	body     string
}

// Run returns the *.md files under root (walked recursively) that contain
// ALL terms (case-insensitive substring), ranked by Score desc then Path asc.
func Run(root string, terms []string) ([]Hit, error) {
	if len(terms) == 0 {
		return nil, nil
	}
	// ponytail: os.Root keeps gosec (G122/G304) happy and paths relative;
	// per-file skipping is unnecessary on a small local mirror.
	bounded, err := os.OpenRoot(root)
	if err != nil {
		return nil, err
	}
	defer bounded.Close() //nolint:errcheck // read-only handle
	fsys := bounded.FS()

	var hits []Hit
	walkErr := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !d.Type().IsRegular() || !strings.HasSuffix(d.Name(), ".md") {
			return nil
		}
		content, err := fs.ReadFile(fsys, path)
		if err != nil {
			return err
		}
		title, f := split(string(content), d.Name())
		score, ok := score(f, terms)
		if !ok {
			return nil
		}
		hits = append(hits, Hit{Path: path, Title: title, Score: score})

		return nil
	})
	if walkErr != nil {
		return nil, walkErr
	}
	slices.SortFunc(hits, func(a, b Hit) int {
		return cmp.Or(b.Score-a.Score, strings.Compare(a.Path, b.Path))
	})

	return hits, nil
}

// split separates frontmatter from body and lowers the searchable fields.
// The title is the frontmatter title, falling back to the filename base
// (without extension); malformed frontmatter is body.
func split(content, filename string) (title string, f fields) {
	base := strings.TrimSuffix(filename, ".md")
	f = fields{filename: strings.ToLower(base), body: strings.ToLower(content)}
	block, body, ok := frontmatter(content)
	if !ok {
		return base, f
	}
	title = titleField(block)
	f.title = strings.ToLower(title)
	f.tags = strings.ToLower(tagsLine(block))
	f.body = strings.ToLower(body)
	if title == "" {
		title = base
	}

	return title, f
}

// frontmatter splits a document into its frontmatter block and body. The
// block is recognized only when the first line is exactly "---" and a
// closing "---" line follows; anything else is all body.
func frontmatter(content string) (block, body string, ok bool) {
	rest, found := strings.CutPrefix(content, fmOpen)
	if !found {
		return "", content, false
	}
	end := strings.Index(rest, "\n"+fmClose)
	if end < 0 {
		return "", content, false
	}
	after := len(fmOpen) + end + 1 + len(fmClose) // prefix, block, "\n", closing

	return rest[:end], strings.TrimPrefix(content[after:], "\n"), true
}

// titleField returns the title value from a frontmatter block, with the
// surrounding quotes stripped; empty when absent.
func titleField(block string) string {
	for _, line := range strings.Split(block, "\n") {
		if value, found := strings.CutPrefix(line, `title: "`); found {
			return strings.TrimSuffix(value, `"`)
		}
	}

	return ""
}

// tagsLine returns the raw tags list line; empty when absent.
func tagsLine(block string) string {
	for _, line := range strings.Split(block, "\n") {
		if value, found := strings.CutPrefix(line, "tags: ["); found {
			return value
		}
	}

	return ""
}

// score returns the summed weights of the terms matching any field; ok is
// false when any term matches nothing (AND semantics).
func score(f fields, terms []string) (int, bool) {
	surfaces := []struct {
		text   string
		weight int
	}{{f.title, weightTitle}, {f.tags, weightTags}, {f.filename, weightFilename}, {f.body, weightBody}}

	total := 0
	for _, term := range terms {
		term = strings.ToLower(term)
		matched := 0
		for _, s := range surfaces {
			if strings.Contains(s.text, term) {
				matched += s.weight
			}
		}
		if matched == 0 {
			return 0, false
		}
		total += matched
	}

	return total, true
}
