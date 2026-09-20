package notion

import (
	"strings"

	"github.com/iyaki/enchiridion/internal/model"
)

// Raw API shapes decoded from JSON, mapped into model types by the client.

type richRun struct {
	PlainText   string `json:"plain_text"`
	Annotations struct {
		Code          bool `json:"code"`
		Bold          bool `json:"bold"`
		Italic        bool `json:"italic"`
		Strikethrough bool `json:"strikethrough"`
	} `json:"annotations"`
	Href string `json:"href"`
}

type urlPayload struct {
	URL string `json:"url"`
}

type textPayload struct {
	RichText []richRun `json:"rich_text"`
	Language string    `json:"language"`
}

type imagePayload struct {
	External *urlPayload `json:"external"`
	File     *urlPayload `json:"file"`
}

type childPagePayload struct {
	Title string `json:"title"`
}

type tablePayload struct {
	HasColumnHeader bool `json:"has_column_header"`
}

type tableRowPayload struct {
	Cells [][]richRun `json:"cells"`
}

// textPayload is embedded (flattened on decode) because most block types carry
// rich_text at the top level of their payload.
type apiBlock struct {
	textPayload

	Type        string            `json:"type"`
	ID          string            `json:"id"`
	HasChildren bool              `json:"has_children"`
	Bookmark    *urlPayload       `json:"bookmark"`
	Embed       *urlPayload       `json:"embed"`
	LinkPreview *urlPayload       `json:"link_preview"`
	Image       *imagePayload     `json:"image"`
	ChildPage   *childPagePayload `json:"child_page"`
	Table       *tablePayload     `json:"table"`
	TableRow    *tableRowPayload  `json:"table_row"`
}

func (b apiBlock) text() []richRun { return b.RichText }

func richText(runs []richRun) []model.RichText {
	out := make([]model.RichText, 0, len(runs))
	for _, r := range runs {
		out = append(out, model.RichText{
			Plain:  r.PlainText,
			Code:   r.Annotations.Code,
			Bold:   r.Annotations.Bold,
			Italic: r.Annotations.Italic,
			Strike: r.Annotations.Strikethrough,
			Href:   r.Href,
		})
	}

	return out
}

func (b apiBlock) url(blockType string) string {
	switch blockType {
	case model.TypeBookmark:
		if b.Bookmark != nil {
			return b.Bookmark.URL
		}
	case model.TypeEmbed:
		if b.Embed != nil {
			return b.Embed.URL
		}
	case model.TypeLinkPreview:
		if b.LinkPreview != nil {
			return b.LinkPreview.URL
		}
	}

	return ""
}

type run struct {
	PlainText string `json:"plain_text"`
}

type option struct {
	Name string `json:"name"`
}

type property struct {
	Type        string   `json:"type"`
	Title       []run    `json:"title"`
	Select      *option  `json:"select"`
	MultiSelect []option `json:"multi_select"`
	Status      *option  `json:"status"`
	URL         string   `json:"url"`
}

type page struct {
	ID             string              `json:"id"`
	URL            string              `json:"url"`
	LastEditedTime string              `json:"last_edited_time"`
	Properties     map[string]property `json:"properties"`
}

func (p page) meta() model.PageMeta {
	meta := model.PageMeta{
		ID:         p.ID,
		NotionURL:  p.URL,
		LastEdited: p.LastEditedTime,
	}

	for _, prop := range p.Properties {
		applyProperty(&meta, prop)
	}

	return meta
}

func applyProperty(meta *model.PageMeta, prop property) {
	switch prop.Type {
	case "title":
		setTitle(meta, prop.Title)
	case "select", "multi_select", "status":
		applyTags(meta, prop)
	case "url":
		if prop.URL != "" {
			meta.SourceURL = prop.URL
		}
	}
}

func setTitle(meta *model.PageMeta, runs []run) {
	var title strings.Builder
	for _, r := range runs {
		title.WriteString(r.PlainText)
	}
	if title.Len() > 0 {
		meta.Title = title.String()
	}
}

func applyTags(meta *model.PageMeta, prop property) {
	switch prop.Type {
	case "select":
		if prop.Select != nil {
			meta.Tags = append(meta.Tags, prop.Select.Name)
		}
	case "multi_select":
		for _, o := range prop.MultiSelect {
			meta.Tags = append(meta.Tags, o.Name)
		}
	case "status":
		if prop.Status != nil {
			meta.Tags = append(meta.Tags, prop.Status.Name)
		}
	}
}
