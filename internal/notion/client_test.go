package notion

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/iyaki/enchiridion/internal/model"
)

type recorder struct {
	requests int
	bodies   []string
	sleeps   []time.Duration
	handler  func(w http.ResponseWriter, r *http.Request)
}

func newTestClient(t *testing.T, rec *recorder) *Client {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec.requests++
		if r.Body != nil {
			buf, _ := io.ReadAll(r.Body)
			rec.bodies = append(rec.bodies, string(buf))
		}
		rec.handler(w, r)
	}))
	t.Cleanup(srv.Close)

	c := NewClient("secret_test")
	c.baseURL = srv.URL
	c.sleep = func(d time.Duration) { rec.sleeps = append(rec.sleeps, d) }

	return c
}

type queryResponse struct {
	Results    []page `json:"results"`
	HasMore    bool   `json:"has_more"`
	NextCursor string `json:"next_cursor"`
}

type blocksResponse struct {
	Results    []apiBlock `json:"results"`
	HasMore    bool       `json:"has_more"`
	NextCursor string     `json:"next_cursor"`
}

func writeJSON(t *testing.T, w http.ResponseWriter, v any) {
	t.Helper()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		t.Errorf("write json response: %v", err)
	}
}

// boldRun builds a bold rich run with the full annotation struct the API uses.
func boldRun(s string) richRun {
	return richRun{
		PlainText: s,
		Annotations: struct {
			Code          bool `json:"code"`
			Bold          bool `json:"bold"`
			Italic        bool `json:"italic"`
			Strikethrough bool `json:"strikethrough"`
		}{Bold: true},
	}
}

func assertQueryHeaders(t *testing.T, r *http.Request) {
	t.Helper()
	if r.Method != http.MethodPost || !strings.HasSuffix(r.URL.Path, "/data_sources/ds123/query") {
		t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
	}
	if auth := r.Header.Get("Authorization"); auth != "Bearer secret_test" {
		t.Errorf("Authorization = %q", auth)
	}
	if v := r.Header.Get("Notion-Version"); v != "2025-09-03" {
		t.Errorf("Notion-Version = %q", v)
	}
}

func assertPageSize(t *testing.T, rec *recorder) {
	t.Helper()
	var body struct {
		PageSize int `json:"page_size"`
	}
	_ = json.Unmarshal([]byte(rec.bodies[rec.requests-1]), &body)
	if body.PageSize != 100 {
		t.Errorf("page_size = %d, want 100", body.PageSize)
	}
}

func TestQueryPagesPaginatesInOrder(t *testing.T) {
	rec := &recorder{}
	c := newTestClient(t, rec)
	rec.handler = func(w http.ResponseWriter, r *http.Request) {
		assertQueryHeaders(t, r)
		assertPageSize(t, rec)

		if rec.requests == 1 {
			pg := page{
				ID: "id-1", URL: "https://notion.so/id-1",
				LastEditedTime: "2026-09-20T10:00:00.000Z",
			}
			writeJSON(t, w, queryResponse{Results: []page{pg}, HasMore: true, NextCursor: "cursor-2"})

			return
		}
		var body struct {
			StartCursor string `json:"start_cursor"`
		}
		_ = json.Unmarshal([]byte(rec.bodies[rec.requests-1]), &body)
		if body.StartCursor != "cursor-2" {
			t.Errorf("start_cursor = %q, want cursor-2", body.StartCursor)
		}
		pg := page{ID: "id-2", URL: "https://notion.so/id-2", LastEditedTime: "2026-09-21T10:00:00.000Z"}
		writeJSON(t, w, queryResponse{Results: []page{pg}})
	}

	metas, err := c.QueryPages("ds123", QueryFilter{})
	if err != nil {
		t.Fatalf("QueryPages: %v", err)
	}
	if len(metas) != 2 {
		t.Fatalf("got %d pages, want 2", len(metas))
	}
	for i, want := range []string{"id-1", "id-2"} {
		if metas[i].ID != want {
			t.Errorf("page %d id = %q, want %q", i, metas[i].ID, want)
		}
	}
}

func TestQueryPagesMapsProperties(t *testing.T) {
	rec := &recorder{}
	c := newTestClient(t, rec)
	rec.handler = func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(t, w, queryResponse{Results: []page{{
			ID:             "id-1",
			URL:            "https://notion.so/id-1",
			LastEditedTime: "2026-09-20T10:00:00.000Z",
			Properties: map[string]property{
				"Name": {Type: "title", Title: []run{{PlainText: "Design Patterns"}}},
				"URL":  {Type: "url", URL: "https://src.example/a"},
				"Category": {Type: "multi_select", MultiSelect: []option{
					{Name: "arquitectura"}, {Name: "articulo"},
				}},
			},
		}}})
	}

	metas, err := c.QueryPages("ds123", QueryFilter{})
	if err != nil {
		t.Fatalf("QueryPages: %v", err)
	}
	if metas[0].Title != "Design Patterns" || metas[0].SourceURL != "https://src.example/a" {
		t.Errorf("meta = %+v", metas[0])
	}
	if got := strings.Join(metas[0].Tags, ","); got != "arquitectura,articulo" {
		t.Errorf("tags = %q", got)
	}
}

func TestQueryPagesIncrementalFilter(t *testing.T) {
	rec := &recorder{}
	c := newTestClient(t, rec)
	rec.handler = func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(t, w, queryResponse{})
	}

	_, err := c.QueryPages("ds123", QueryFilter{LastEditedAfter: "2026-01-02T00:00:00Z"})
	if err != nil {
		t.Fatalf("QueryPages: %v", err)
	}

	var body struct {
		Filter struct {
			Timestamp      string `json:"timestamp"`
			LastEditedTime struct {
				OnOrAfter string `json:"on_or_after"`
			} `json:"last_edited_time"`
		} `json:"filter"`
	}
	if err := json.Unmarshal([]byte(rec.bodies[0]), &body); err != nil {
		t.Fatalf("decode request body: %v", err)
	}
	if body.Filter.Timestamp != "last_edited_time" ||
		body.Filter.LastEditedTime.OnOrAfter != "2026-01-02T00:00:00Z" {
		t.Errorf("filter = %+v", body.Filter)
	}
}

func TestNilAndEmptyPropertiesMapSafely(t *testing.T) {
	rec := &recorder{}
	c := newTestClient(t, rec)
	rec.handler = func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(t, w, queryResponse{Results: []page{{
			ID:             "id-9",
			URL:            "https://notion.so/id-9",
			LastEditedTime: "2026-09-20T00:00:00Z",
			Properties: map[string]property{
				"Title": {Type: "title"},
				"Cat":   {Type: "select"},
				"Est":   {Type: "status"},
				"Link":  {Type: "url"},
			},
		}}})
	}

	metas, err := c.QueryPages("ds123", QueryFilter{})
	if err != nil {
		t.Fatalf("QueryPages: %v", err)
	}
	if metas[0].Title != "" || metas[0].SourceURL != "" || len(metas[0].Tags) != 0 {
		t.Errorf("meta = %+v, want empty title/tags/source", metas[0])
	}
}

func TestPageBlocksTextMapping(t *testing.T) {
	rec := &recorder{}
	c := newTestClient(t, rec)
	rec.handler = func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(t, w, blocksResponse{Results: []apiBlock{
			{Type: model.TypeParagraph, Paragraph: &textPayload{RichText: []richRun{
				{PlainText: "plain "}, boldRun("bold"),
			}}},
			{Type: model.TypeCode, Code: &textPayload{
				RichText: []richRun{{PlainText: "const x = 1"}}, Language: "js",
			}},
			{Type: model.TypeChildPage, ChildPage: &childPagePayload{Title: "Specs"}},
			{Type: "morph"},
		}})
	}

	blocks, err := c.PageBlocks("page1")
	if err != nil {
		t.Fatalf("PageBlocks: %v", err)
	}

	wantTypes := []string{model.TypeParagraph, model.TypeCode, model.TypeChildPage, "morph"}
	if len(blocks) != len(wantTypes) {
		t.Fatalf("got %d blocks, want %d", len(blocks), len(wantTypes))
	}
	for i, want := range wantTypes {
		if blocks[i].Type != want {
			t.Errorf("block %d type = %q, want %q", i, blocks[i].Type, want)
		}
	}
	if got := blocks[0].RichText[1]; !got.Bold || got.Plain != "bold" {
		t.Errorf("paragraph runs = %+v", blocks[0].RichText)
	}
	if blocks[1].Language != "js" {
		t.Errorf("code language = %q", blocks[1].Language)
	}
	if blocks[2].Title != "Specs" {
		t.Errorf("child page title = %q", blocks[2].Title)
	}
}

func TestPageBlocksTableRowsResolvedRecursively(t *testing.T) {
	rec := &recorder{}
	c := newTestClient(t, rec)
	rec.handler = func(w http.ResponseWriter, _ *http.Request) {
		if rec.requests == 1 {
			table := apiBlock{
				Type: model.TypeTable, ID: "tbl-1", HasChildren: true,
				Table: &tablePayload{HasColumnHeader: true},
			}
			writeJSON(t, w, blocksResponse{Results: []apiBlock{table}})

			return
		}
		writeJSON(t, w, blocksResponse{Results: []apiBlock{
			{Type: "table_row", TableRow: &tableRowPayload{Cells: [][]richRun{
				{{PlainText: "Name"}},
			}}},
			{Type: "table_row", TableRow: &tableRowPayload{Cells: [][]richRun{
				{{PlainText: "a|b"}},
				{boldRun("x\ny")},
			}}},
		}})
	}

	blocks, err := c.PageBlocks("page1")
	if err != nil {
		t.Fatalf("PageBlocks: %v", err)
	}

	tbl := blocks[0]
	if !tbl.HasHeader || len(tbl.Rows) != 2 {
		t.Fatalf("table = %+v", tbl)
	}
	if tbl.Rows[0][0][0].Plain != "Name" {
		t.Errorf("header cell = %+v", tbl.Rows[0][0])
	}
	if tbl.Rows[1][0][0].Plain != "a|b" {
		t.Errorf("cell 0 = %+v", tbl.Rows[1][0])
	}
	if !tbl.Rows[1][1][0].Bold {
		t.Errorf("cell 1 = %+v", tbl.Rows[1][1])
	}
}

func TestPageBlocksLinksAndImages(t *testing.T) {
	rec := &recorder{}
	c := newTestClient(t, rec)
	rec.handler = func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(t, w, blocksResponse{Results: []apiBlock{
			{Type: model.TypeBookmark, Bookmark: &urlPayload{URL: "https://a.b"}},
			{Type: model.TypeBookmark},
			{Type: model.TypeImage, Image: &imagePayload{External: &urlPayload{URL: "https://img.example/x.png"}}},
			{Type: model.TypeImage, Image: &imagePayload{File: &urlPayload{URL: "https://notion.so/expiring.png"}}},
			{Type: model.TypeImage},
		}})
	}

	blocks, err := c.PageBlocks("page1")
	if err != nil {
		t.Fatalf("PageBlocks: %v", err)
	}
	if blocks[0].URL != "https://a.b" {
		t.Errorf("bookmark url = %q", blocks[0].URL)
	}
	if blocks[1].URL != "" {
		t.Errorf("payload-less bookmark url = %q, want empty", blocks[1].URL)
	}
	if blocks[2].URL != "https://img.example/x.png" || blocks[2].Internal {
		t.Errorf("external image = %+v", blocks[2])
	}
	if !blocks[3].Internal || blocks[3].URL != "https://notion.so/expiring.png" {
		t.Errorf("internal image = %+v", blocks[3])
	}
	if blocks[4].URL != "" || blocks[4].Internal {
		t.Errorf("source-less image = %+v", blocks[4])
	}
}

func TestRetries429HonoringRetryAfter(t *testing.T) {
	rec := &recorder{}
	c := newTestClient(t, rec)
	rec.handler = func(w http.ResponseWriter, _ *http.Request) {
		if rec.requests == 1 {
			w.Header().Set("Retry-After", "7")
			w.WriteHeader(http.StatusTooManyRequests)

			return
		}
		writeJSON(t, w, queryResponse{})
	}

	if _, err := c.QueryPages("ds123", QueryFilter{}); err != nil {
		t.Fatalf("QueryPages: %v", err)
	}
	if rec.requests != 2 {
		t.Errorf("requests = %d, want 2", rec.requests)
	}
	if len(rec.sleeps) != 1 || rec.sleeps[0] != 7*time.Second {
		t.Errorf("sleeps = %v, want [7s]", rec.sleeps)
	}
}

func TestRateLimitWithoutRetryAfterUsesBackoff(t *testing.T) {
	rec := &recorder{}
	c := newTestClient(t, rec)
	rec.handler = func(w http.ResponseWriter, _ *http.Request) {
		if rec.requests == 1 {
			w.WriteHeader(http.StatusTooManyRequests)

			return
		}
		writeJSON(t, w, queryResponse{})
	}

	if _, err := c.QueryPages("ds123", QueryFilter{}); err != nil {
		t.Fatalf("QueryPages: %v", err)
	}
	if len(rec.sleeps) != 1 || rec.sleeps[0] != 1*time.Second {
		t.Errorf("sleeps = %v, want [1s] exponential fallback", rec.sleeps)
	}
}

func TestGivesUpAfterThreeAttemptsOnServerError(t *testing.T) {
	rec := &recorder{}
	c := newTestClient(t, rec)
	rec.handler = func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}

	_, err := c.QueryPages("ds123", QueryFilter{})
	if err == nil || !strings.Contains(err.Error(), "gave up after 3 attempts") {
		t.Fatalf("err = %v, want give-up after 3 attempts", err)
	}
	if rec.requests != 3 {
		t.Errorf("requests = %d, want 3", rec.requests)
	}
	if len(rec.sleeps) != 3 {
		t.Errorf("sleeps = %d entries, want 3", len(rec.sleeps))
	}
}

func TestUnauthorizedFailsFastWithoutRetries(t *testing.T) {
	rec := &recorder{}
	c := newTestClient(t, rec)
	rec.handler = func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}

	_, err := c.QueryPages("ds123", QueryFilter{})
	if err == nil || !strings.Contains(err.Error(), "credentials") {
		t.Fatalf("err = %v, want credential error", err)
	}
	if rec.requests != 1 {
		t.Errorf("requests = %d, want 1 (no retries on auth failure)", rec.requests)
	}
	if len(rec.sleeps) != 0 {
		t.Errorf("sleeps = %v, want none", rec.sleeps)
	}
}

func TestClientErrorSurfacesBody(t *testing.T) {
	rec := &recorder{}
	c := newTestClient(t, rec)
	rec.handler = func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"message":"Invalid data source"}`))
	}

	_, err := c.QueryPages("nope", QueryFilter{})
	if err == nil || !strings.Contains(err.Error(), "Invalid data source") {
		t.Fatalf("err = %v, want surfaced API message", err)
	}
}

// flattenChildrenHandler serves a page whose blocks carry every child-bearing
// shape the flattening must cover (ADR-17), dispatched by block ID.
func flattenChildrenHandler(t *testing.T, w http.ResponseWriter, r *http.Request) {
	t.Helper()
	path := r.URL.Path
	switch {
	case strings.Contains(path, "tg-1"):
		writeJSON(t, w, blocksResponse{Results: []apiBlock{
			{Type: model.TypeParagraph, Paragraph: &textPayload{RichText: []richRun{{PlainText: "inside toggle"}}}},
		}})
	case strings.Contains(path, "td-1"):
		writeJSON(t, w, blocksResponse{Results: []apiBlock{
			{Type: model.TypeBulletedItem, BulletedItem: &textPayload{RichText: []richRun{{PlainText: "nested"}}}},
		}})
	case strings.Contains(path, "sb-1"):
		writeJSON(t, w, blocksResponse{Results: []apiBlock{
			{Type: model.TypeHeading2, Heading2: &textPayload{RichText: []richRun{{PlainText: "synced content"}}}},
		}})
	case strings.Contains(path, "cl-1"):
		writeJSON(t, w, blocksResponse{Results: []apiBlock{
			{Type: "column", ID: "col-1", HasChildren: true},
		}})
	case strings.Contains(path, "col-1"):
		writeJSON(t, w, blocksResponse{Results: []apiBlock{
			{Type: model.TypeBulletedItem, BulletedItem: &textPayload{RichText: []richRun{{PlainText: "in column"}}}},
		}})
	case strings.Contains(path, "tbl-1"):
		writeJSON(t, w, blocksResponse{Results: []apiBlock{
			{Type: "table_row", TableRow: &tableRowPayload{Cells: [][]richRun{{{PlainText: "cell"}}}}},
		}})
	default:
		writeJSON(t, w, blocksResponse{Results: []apiBlock{
			{Type: model.TypeToggle, ID: "tg-1", HasChildren: true,
				Toggle: &textPayload{RichText: []richRun{{PlainText: "how to"}}}},
			{Type: model.TypeToDo, ID: "td-1", HasChildren: true,
				ToDo: &toDoPayload{
					textPayload: textPayload{RichText: []richRun{{PlainText: "ship"}}},
					Checked:     true,
				}},
			{Type: "synced_block", ID: "sb-1", HasChildren: true},
			{Type: "column_list", ID: "cl-1", HasChildren: true},
			{Type: model.TypeTable, ID: "tbl-1", HasChildren: true,
				Table: &tablePayload{HasColumnHeader: false}},
		}})
	}
}

func TestPageBlocksFlattensAllChildren(t *testing.T) {
	rec := &recorder{}
	c := newTestClient(t, rec)
	rec.handler = func(w http.ResponseWriter, r *http.Request) {
		flattenChildrenHandler(t, w, r)
	}

	blocks, err := c.PageBlocks("page1")
	if err != nil {
		t.Fatalf("PageBlocks: %v", err)
	}

	want := []struct {
		typ, text string
		checked   bool
	}{
		{model.TypeToggle, "how to", false},
		{model.TypeParagraph, "inside toggle", false},
		{model.TypeToDo, "ship", true},
		{model.TypeBulletedItem, "nested", false},
		{model.TypeHeading2, "synced content", false},
		{model.TypeBulletedItem, "in column", false},
		{model.TypeTable, "", false},
	}
	if len(blocks) != len(want) {
		t.Fatalf("got %d blocks, want %d", len(blocks), len(want))
	}
	for i, w := range want {
		if blocks[i].Type != w.typ || plainText(blocks[i]) != w.text || blocks[i].Checked != w.checked {
			t.Errorf("block %d = %s %q checked=%v, want %s %q checked=%v",
				i, blocks[i].Type, plainText(blocks[i]), blocks[i].Checked, w.typ, w.text, w.checked)
		}
	}
	assertNoContainers(t, blocks)
}

// assertNoContainers fails if pure container blocks leaked into the output.
func assertNoContainers(t *testing.T, blocks []model.Block) {
	t.Helper()
	for _, blk := range blocks {
		if blk.Type == "synced_block" || blk.Type == "column_list" || blk.Type == "column" {
			t.Errorf("container block %q leaked into output", blk.Type)
		}
	}
}

func TestPageBlocksMediaURLs(t *testing.T) {
	rec := &recorder{}
	c := newTestClient(t, rec)
	rec.handler = func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(t, w, blocksResponse{Results: []apiBlock{
			{Type: model.TypePDF, PDF: &imagePayload{External: &urlPayload{URL: "https://example.com/doc.pdf"}}},
			{Type: model.TypeFile, File: &imagePayload{File: &urlPayload{URL: "https://notion.so/expire"}}},
			{Type: model.TypeVideo, Video: &imagePayload{External: &urlPayload{URL: "https://example.com/clip.mp4"}}},
		}})
	}

	blocks, err := c.PageBlocks("page1")
	if err != nil {
		t.Fatalf("PageBlocks: %v", err)
	}
	if len(blocks) != 3 {
		t.Fatalf("got %d blocks, want 3", len(blocks))
	}
	if blocks[0].URL != "https://example.com/doc.pdf" || blocks[0].Internal {
		t.Errorf("pdf = %+v", blocks[0])
	}
	if blocks[1].URL != "https://notion.so/expire" || !blocks[1].Internal {
		t.Errorf("file = %+v", blocks[1])
	}
	if blocks[2].URL != "https://example.com/clip.mp4" || blocks[2].Internal {
		t.Errorf("video = %+v", blocks[2])
	}
}

func plainText(blk model.Block) string {
	out := ""
	for _, r := range blk.RichText {
		out += r.Plain
	}

	return out
}
