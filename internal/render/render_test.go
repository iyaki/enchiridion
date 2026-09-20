package render

import (
	"testing"

	"github.com/iyaki/enchiridion/internal/model"
)

func rt(s string) model.RichText { return model.RichText{Plain: s} }

func TestRendersContractBlocks(t *testing.T) {
	cases := map[string]struct {
		in   []model.Block
		want string
	}{
		"paragraph with inline styles and link": {
			in: []model.Block{{Type: model.TypeParagraph, RichText: []model.RichText{
				rt("plain "),
				{Plain: "cfg", Code: true},
				{Plain: " bold", Bold: true},
				{Plain: " it", Italic: true},
				{Plain: " gone", Strike: true},
				{Plain: "docs", Href: "https://example.com"},
			}}},
			want: "plain `cfg`** bold**_ it_~~ gone~~[docs](https://example.com)",
		},
		"heading levels": {
			in: []model.Block{
				{Type: model.TypeHeading1, RichText: []model.RichText{rt("One")}},
				{Type: model.TypeHeading2, RichText: []model.RichText{rt("Two")}},
				{Type: model.TypeHeading3, RichText: []model.RichText{rt("Three")}},
			},
			want: "# One\n\n## Two\n\n### Three",
		},
		"tight bulleted run": {
			in: []model.Block{
				{Type: model.TypeBulletedItem, RichText: []model.RichText{rt("a")}},
				{Type: model.TypeBulletedItem, RichText: []model.RichText{rt("b")}},
			},
			want: "- a\n- b",
		},
		"numbered run resets after non-numbered block": {
			in: []model.Block{
				{Type: model.TypeNumberedItem, RichText: []model.RichText{rt("first")}},
				{Type: model.TypeNumberedItem, RichText: []model.RichText{rt("second")}},
				{Type: model.TypeParagraph, RichText: []model.RichText{rt("break")}},
				{Type: model.TypeNumberedItem, RichText: []model.RichText{rt("restart")}},
			},
			want: "1. first\n2. second\n\nbreak\n\n1. restart",
		},
		"quote and callout": {
			in: []model.Block{
				{Type: model.TypeQuote, RichText: []model.RichText{rt("q")}},
				{Type: model.TypeCallout, RichText: []model.RichText{rt("c")}},
			},
			want: "> q\n\n> c",
		},
		"code fence with language": {
			in: []model.Block{{
				Type:     model.TypeCode,
				Language: "js",
				RichText: []model.RichText{rt("const x = 1")},
			}},
			want: "```js\nconst x = 1\n```",
		},
		"divider": {
			in:   []model.Block{{Type: model.TypeDivider}},
			want: "---",
		},
		"link-style blocks": {
			in: []model.Block{
				{Type: model.TypeBookmark, URL: "https://a.b"},
				{Type: model.TypeEmbed, URL: "https://c.d"},
				{Type: model.TypeLinkPreview, URL: "https://e.f"},
			},
			want: "[https://a.b](https://a.b)\n\n[https://c.d](https://c.d)\n\n[https://e.f](https://e.f)",
		},
		"link and image without source become visible markers": {
			in: []model.Block{
				{Type: model.TypeBookmark},
				{Type: model.TypeImage},
			},
			want: "<!-- link without URL -->\n\n<!-- image without source -->",
		},
		"external and internal images": {
			in: []model.Block{
				{Type: model.TypeImage, URL: "https://img.example/x.png"},
				{Type: model.TypeImage, URL: "https://notion.so/expire.png", Internal: true},
			},
			want: "![image](https://img.example/x.png)\n\n" +
				"<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->",
		},
		"toggle and child page": {
			in: []model.Block{
				{Type: model.TypeToggle, RichText: []model.RichText{rt("Details")}},
				{Type: model.TypeChildPage, Title: "Specs"},
			},
			want: "**Details**\n\n<!-- child page: Specs -->",
		},
		"table with header, pipes and newlines": {
			in: []model.Block{{
				Type:      model.TypeTable,
				HasHeader: true,
				Rows: [][]model.Cell{
					{{rt("Name")}, {rt("Value")}},
					{{rt("a|b")}, {{Plain: "x\ny", Bold: true}}},
				},
			}},
			want: "| Name | Value |\n| --- | --- |\n| a\\|b | **x<br>y** |",
		},
		"table without declared header still renders separator": {
			in: []model.Block{{
				Type: model.TypeTable,
				Rows: [][]model.Cell{{{rt("only")}}},
			}},
			want: "| only |\n| --- |",
		},
		"empty table becomes visible marker": {
			in:   []model.Block{{Type: model.TypeTable}},
			want: "<!-- unsupported block: empty table -->",
		},
		"unsupported block becomes visible marker": {
			in:   []model.Block{{Type: "morph"}},
			want: "<!-- unsupported block: morph -->",
		},
		"empty document": {in: nil, want: ""},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if got := Markdown(tc.in); got != tc.want {
				t.Fatalf("got:\n%q\nwant:\n%q", got, tc.want)
			}
		})
	}
}

func TestLooseSeparatorsBetweenDifferentBlocks(t *testing.T) {
	in := []model.Block{
		{Type: model.TypeParagraph, RichText: []model.RichText{rt("p")}},
		{Type: model.TypeBulletedItem, RichText: []model.RichText{rt("b")}},
		{Type: model.TypeHeading1, RichText: []model.RichText{rt("h")}},
	}

	want := "p\n\n- b\n\n# h"
	if got := Markdown(in); got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}
