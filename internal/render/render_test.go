package render

import (
	"testing"

	"github.com/iyaki/enchiridion/internal/model"
)

func rt(s string) model.RichText { return model.RichText{Plain: s} }

func TestInlineRendering(t *testing.T) {
	blocks := []model.Block{{Type: model.TypeParagraph, RichText: []model.RichText{
		rt("plain "),
		{Plain: "cfg", Code: true},
		{Plain: " bold", Bold: true},
		{Plain: " it", Italic: true},
		{Plain: " gone", Strike: true},
		{Plain: "docs", Href: "https://example.com"},
	}}}

	want := "plain `cfg`** bold**_ it_~~ gone~~[docs](https://example.com)"
	if got := Markdown(blocks); got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestHeadingLevels(t *testing.T) {
	blocks := []model.Block{
		{Type: model.TypeHeading1, RichText: []model.RichText{rt("One")}},
		{Type: model.TypeHeading2, RichText: []model.RichText{rt("Two")}},
		{Type: model.TypeHeading3, RichText: []model.RichText{rt("Three")}},
	}

	want := "# One\n\n## Two\n\n### Three"
	if got := Markdown(blocks); got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestListBehavior(t *testing.T) {
	blocks := []model.Block{
		{Type: model.TypeBulletedItem, RichText: []model.RichText{rt("a")}},
		{Type: model.TypeBulletedItem, RichText: []model.RichText{rt("b")}},
		{Type: model.TypeNumberedItem, RichText: []model.RichText{rt("first")}},
		{Type: model.TypeNumberedItem, RichText: []model.RichText{rt("second")}},
		{Type: model.TypeParagraph, RichText: []model.RichText{rt("break")}},
		{Type: model.TypeNumberedItem, RichText: []model.RichText{rt("restart")}},
	}

	want := "- a\n- b\n\n1. first\n2. second\n\nbreak\n\n1. restart"
	if got := Markdown(blocks); got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestQuoteCodeDividerToggleChildPage(t *testing.T) {
	blocks := []model.Block{
		{Type: model.TypeQuote, RichText: []model.RichText{rt("q")}},
		{Type: model.TypeCallout, RichText: []model.RichText{rt("c")}},
		{Type: model.TypeCode, Language: "js", RichText: []model.RichText{rt("const x = 1")}},
		{Type: model.TypeDivider},
		{Type: model.TypeToggle, RichText: []model.RichText{rt("Details")}},
		{Type: model.TypeChildPage, Title: "Specs"},
	}

	want := "> q\n\n> c\n\n```js\nconst x = 1\n```\n\n---\n\n**Details**\n\n<!-- child page: Specs -->"
	if got := Markdown(blocks); got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestLinkAndImageRendering(t *testing.T) {
	blocks := []model.Block{
		{Type: model.TypeBookmark, URL: "https://a.b"},
		{Type: model.TypeEmbed, URL: "https://c.d"},
		{Type: model.TypeLinkPreview, URL: "https://e.f"},
		{Type: model.TypeImage, URL: "https://img.example/x.png"},
		{Type: model.TypeImage, URL: "https://notion.so/expire.png", Internal: true},
	}

	want := "[https://a.b](https://a.b)\n\n[https://c.d](https://c.d)\n\n[https://e.f](https://e.f)\n\n" +
		"![image](https://img.example/x.png)\n\n" +
		"<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->"
	if got := Markdown(blocks); got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestSourcelessBlocksRenderVisibleMarkers(t *testing.T) {
	blocks := []model.Block{
		{Type: model.TypeBookmark},
		{Type: model.TypeImage},
	}

	want := "<!-- link without URL -->\n\n<!-- image without source -->"
	if got := Markdown(blocks); got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestTableRendering(t *testing.T) {
	cases := map[string]struct {
		in   model.Block
		want string
	}{
		"header with pipes and newlines": {
			in: model.Block{
				Type: model.TypeTable, HasHeader: true,
				Rows: [][]model.Cell{
					{{rt("Name")}, {rt("Value")}},
					{{rt("a|b")}, {{Plain: "x\ny", Bold: true}}},
				},
			},
			want: "| Name | Value |\n| --- | --- |\n| a\\|b | **x<br>y** |",
		},
		"separator always rendered": {
			in:   model.Block{Type: model.TypeTable, Rows: [][]model.Cell{{{rt("only")}}}},
			want: "| only |\n| --- |",
		},
		"empty table becomes visible marker": {
			in:   model.Block{Type: model.TypeTable},
			want: "<!-- unsupported block: empty table -->",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if got := Markdown([]model.Block{tc.in}); got != tc.want {
				t.Fatalf("got %q, want %q", got, tc.want)
			}
		})
	}
}

func TestUnsupportedBlockBecomesVisibleMarker(t *testing.T) {
	if got := Markdown([]model.Block{{Type: "morph"}}); got != "<!-- unsupported block: morph -->" {
		t.Fatalf("got %q", got)
	}
	if got := Markdown(nil); got != "" {
		t.Fatalf("empty document: got %q, want empty", got)
	}
}

func TestLooseSeparatorsBetweenDifferentBlocks(t *testing.T) {
	blocks := []model.Block{
		{Type: model.TypeParagraph, RichText: []model.RichText{rt("p")}},
		{Type: model.TypeBulletedItem, RichText: []model.RichText{rt("b")}},
		{Type: model.TypeHeading1, RichText: []model.RichText{rt("h")}},
	}

	want := "p\n\n- b\n\n# h"
	if got := Markdown(blocks); got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}
