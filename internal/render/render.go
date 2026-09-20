// Package render translates distilled knowledge-base blocks into markdown
// exactly per the renderer contract in specs/architecture.md. Pure: no I/O,
// no clock, no env.
package render

import (
	"strconv"
	"strings"

	"github.com/iyaki/enchiridion/internal/model"
)

// Markdown renders blocks into a single markdown document.
func Markdown(blocks []model.Block) string {
	var b strings.Builder
	numbered := 0
	prev := ""
	for i, blk := range blocks {
		if i > 0 {
			if blk.Type == prev && isList(blk.Type) {
				b.WriteString("\n")
			} else {
				b.WriteString("\n\n")
			}
		}
		if blk.Type != model.TypeNumberedItem {
			numbered = 0
		}
		writeBlock(&b, blk, &numbered)
		prev = blk.Type
	}

	return b.String()
}

func isList(blockType string) bool {
	return blockType == model.TypeBulletedItem || blockType == model.TypeNumberedItem
}

// writeBlock dispatches to focused writers; returns false for unknown types so
// the caller can emit a visible marker instead of silent output.
func writeBlock(b *strings.Builder, blk model.Block, numbered *int) bool {
	if writeTextBlock(b, blk, numbered) {
		return true
	}
	if writePayloadBlock(b, blk) {
		return true
	}
	b.WriteString("<!-- unsupported block: " + blk.Type + " -->")

	return true
}

func writeTextBlock(b *strings.Builder, blk model.Block, numbered *int) bool {
	switch blk.Type {
	case model.TypeParagraph:
		b.WriteString(inline(blk.RichText))
	case model.TypeHeading1, model.TypeHeading2, model.TypeHeading3:
		writeHeading(b, blk)
	case model.TypeBulletedItem:
		b.WriteString("- " + inline(blk.RichText))
	case model.TypeNumberedItem:
		*numbered++
		b.WriteString(strconv.Itoa(*numbered) + ". " + inline(blk.RichText))
	case model.TypeQuote, model.TypeCallout:
		b.WriteString("> " + inline(blk.RichText))
	case model.TypeDivider:
		b.WriteString("---")
	case model.TypeToggle:
		b.WriteString("**" + inline(blk.RichText) + "**")
	default:
		return false
	}

	return true
}

func writeHeading(b *strings.Builder, blk model.Block) {
	switch blk.Type {
	case model.TypeHeading1:
		b.WriteString("# " + inline(blk.RichText))
	case model.TypeHeading2:
		b.WriteString("## " + inline(blk.RichText))
	case model.TypeHeading3:
		b.WriteString("### " + inline(blk.RichText))
	}
}

func writePayloadBlock(b *strings.Builder, blk model.Block) bool {
	switch blk.Type {
	case model.TypeBookmark, model.TypeEmbed, model.TypeLinkPreview:
		writeLink(b, blk)
	case model.TypeImage:
		writeImage(b, blk)
	case model.TypeCode:
		b.WriteString("```" + blk.Language + "\n" + plain(blk.RichText) + "\n```")
	case model.TypeChildPage:
		b.WriteString("<!-- child page: " + blk.Title + " -->")
	case model.TypeTable:
		b.WriteString(table(blk))
	default:
		return false
	}

	return true
}

func writeLink(b *strings.Builder, blk model.Block) {
	if blk.URL == "" {
		b.WriteString("<!-- link without URL -->")

		return
	}
	b.WriteString("[" + blk.URL + "](" + blk.URL + ")")
}

func writeImage(b *strings.Builder, blk model.Block) {
	switch {
	case blk.Internal:
		b.WriteString("<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->")
	case blk.URL == "":
		b.WriteString("<!-- image without source -->")
	default:
		b.WriteString("![image](" + blk.URL + ")")
	}
}

// inline renders styled runs; order matters: code, bold, italic,
// strikethrough, then link (specs/architecture.md).
func inline(runs []model.RichText) string {
	var b strings.Builder
	for _, r := range runs {
		if r.Plain == "" {
			continue
		}
		t := r.Plain
		if r.Code {
			t = "`" + t + "`"
		}
		if r.Bold {
			t = "**" + t + "**"
		}
		if r.Italic {
			t = "_" + t + "_"
		}
		if r.Strike {
			t = "~~" + t + "~~"
		}
		if r.Href != "" {
			t = "[" + t + "](" + r.Href + ")"
		}
		b.WriteString(t)
	}

	return b.String()
}

func plain(runs []model.RichText) string {
	var b strings.Builder
	for _, r := range runs {
		b.WriteString(r.Plain)
	}

	return b.String()
}

// table renders a markdown table. Markdown requires a separator row after the
// first row regardless of header semantics; HasHeader only marks the first row
// as a header in the source.
func table(blk model.Block) string {
	if len(blk.Rows) == 0 {
		return "<!-- unsupported block: empty table -->"
	}

	var b strings.Builder
	for i, row := range blk.Rows {
		b.WriteString("|")
		for _, cell := range row {
			b.WriteString(" " + escapeCell(cell) + " |")
		}
		b.WriteString("\n")
		if i == 0 {
			b.WriteString("|")
			for range row {
				b.WriteString(" --- |")
			}
			b.WriteString("\n")
		}
	}

	return strings.TrimSuffix(b.String(), "\n")
}

func escapeCell(cell model.Cell) string {
	t := strings.ReplaceAll(inline(cell), "|", `\|`)

	return strings.ReplaceAll(t, "\n", "<br>")
}
