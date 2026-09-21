// Package model holds the distilled domain types shared by the Notion client,
// the renderer, and the sync engine. Pure data: no I/O, no behavior.
package model

// Supported block types (renderer contract: specs/architecture.md).
const (
	TypeParagraph    = "paragraph"
	TypeHeading1     = "heading_1"
	TypeHeading2     = "heading_2"
	TypeHeading3     = "heading_3"
	TypeBulletedItem = "bulleted_list_item"
	TypeNumberedItem = "numbered_list_item"
	TypeQuote        = "quote"
	TypeCallout      = "callout"
	TypeCode         = "code"
	TypeDivider      = "divider"
	TypeBookmark     = "bookmark"
	TypeEmbed        = "embed"
	TypeLinkPreview  = "link_preview"
	TypeImage        = "image"
	TypeToDo         = "to_do"
	TypePDF          = "pdf"
	TypeFile         = "file"
	TypeVideo        = "video"
	TypeToggle       = "toggle"
	TypeChildPage    = "child_page"
	TypeTable        = "table"
)

// RichText is a single styled text run.
type RichText struct {
	Plain  string
	Code   bool
	Bold   bool
	Italic bool
	Strike bool
	Href   string
}

// Cell is one table cell: the styled text runs it contains.
type Cell []RichText

// Block is a distilled Notion content block. The Notion client maps raw API
// blocks into this shape (including table rows, which it fetches recursively);
// the renderer consumes it and emits visible markers for anything it cannot
// render (specs/architecture.md — renderer contract).
type Block struct {
	Type      string
	RichText  []RichText // text-carrying blocks
	Language  string     // code
	URL       string     // bookmark, embed, link_preview, image, pdf, file, video (external)
	Internal  bool       // image/file hosted by Notion: its URL expires (ADR-05)
	Checked   bool       // to_do: checkbox state
	Title     string     // child_page
	HasHeader bool       // table: first row renders as header
	Rows      [][]Cell   // table: rows -> cells -> runs
}

// PageMeta is the per-page identity and classification the sync engine writes
// into the mirror frontmatter (specs/architecture.md — mirror format).
type PageMeta struct {
	ID         string
	NotionURL  string
	Title      string
	Tags       []string
	SourceURL  string // empty when the page has no URL property
	LastEdited string // RFC3339, as returned by the API
}
