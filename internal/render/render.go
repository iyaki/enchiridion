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
		numbered = writeBlock(&b, blk, numbered)
		prev = blk.Type
	}

	return b.String()
}

func isList(blockType string) bool {
	return blockType == model.TypeBulletedItem || blockType == model.TypeNumberedItem
}

func writeBlock(b *strings.Builder, blk model.Block, numbered int) int {
	switch blk.Type {
	case model.TypeParagraph:
		b.WriteString(inline(blk.RichText))
	case model.TypeHeading1:
		b.WriteString("# " + inline(blk.RichText))
	case model.TypeHeading2:
		b.WriteString("## " + inline(blk.RichText))
	case model.TypeHeading3:
		b.WriteString("### " + inline(blk.RichText))
	case model.TypeBulletedItem:
		b.WriteString("- " + inline(blk.RichText))
	case model.TypeNumberedItem:
		numbered++
		b.WriteString(strconv.Itoa(numbered) + ". " + inline(blk.RichText))
	case model.TypeQuote, model.TypeCallout:
		b.WriteString("> " + inline(blk.RichText))
	case model.TypeCode:
		b.WriteString("```" + blk.Language + "\n" + plain(blk.RichText) + "\n```")
	case model.TypeDivider:
		b.WriteString("---")
	case model.TypeBookmark, model.TypeEmbed, model.TypeLinkPreview:
		if blk.URL == "" {
			b.WriteString("<!-- link without URL -->")

			break
		}
		b.WriteString("[" + blk.URL + "](" + blk.URL + ")")
	case model.TypeImage:
		switch {
		case blk.Internal:
			b.WriteString("<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->")
		case blk.URL == "":
			b.WriteString("<!-- image without source -->")
		default:
			b.WriteString("![image](" + blk.URL + ")")
		}
	case model.TypeToggle:
		b.WriteString("**" + inline(blk.RichText) + "**")
	case model.TypeChildPage:
		b.WriteString("<!-- child page: " + blk.Title + " -->")
	case model.TypeTable:
		b.WriteString(table(blk))
	default:
		b.WriteString("<!-- unsupported block: " + blk.Type + " -->")
	}

	return numbered
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
