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

func TestQueryPagesPaginatesAndMapsProperties(t *testing.T) {
	rec := &recorder{}
	c := newTestClient(t, rec)
	rec.handler = func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || !strings.HasSuffix(r.URL.Path, "/data_sources/ds123/query") {
			t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
		if auth := r.Header.Get("Authorization"); auth != "Bearer secret_test" {
			t.Errorf("Authorization = %q", auth)
		}
		if v := r.Header.Get("Notion-Version"); v != "2025-09-03" {
			t.Errorf("Notion-Version = %q", v)
		}

		var body struct {
			PageSize    int    `json:"page_size"`
			StartCursor string `json:"start_cursor"`
		}
		_ = json.Unmarshal([]byte(rec.bodies[rec.requests-1]), &body)
		if body.PageSize != 100 {
			t.Errorf("page_size = %d, want 100", body.PageSize)
		}

		switch rec.requests {
		case 1:
			writeJSON(t, w, queryResponse{
				Results: []page{{
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
				}},
				HasMore:    true,
				NextCursor: "cursor-2",
			})
		case 2:
			if body.StartCursor != "cursor-2" {
				t.Errorf("start_cursor = %q, want cursor-2", body.StartCursor)
			}
			writeJSON(t, w, queryResponse{
				Results: []page{{
					ID:             "id-2",
					URL:            "https://notion.so/id-2",
					LastEditedTime: "2026-09-21T10:00:00.000Z",
					Properties: map[string]property{
						"Name":   {Type: "title", Title: []run{{PlainText: "Second"}}},
						"Temas":  {Type: "multi_select", MultiSelect: []option{{Name: "ddd"}}},
						"Estado": {Type: "status", Status: &option{Name: "published"}},
					},
				}},
			})
		}
	}

	metas, err := c.QueryPages("ds123", QueryFilter{})
	if err != nil {
		t.Fatalf("QueryPages: %v", err)
	}
	if len(metas) != 2 {
		t.Fatalf("got %d pages, want 2", len(metas))
	}
	if metas[0].Title != "Design Patterns" || metas[0].SourceURL != "https://src.example/a" {
		t.Errorf("page 1 meta = %+v", metas[0])
	}
	if got := strings.Join(metas[0].Tags, ","); got != "arquitectura,articulo" {
		t.Errorf("page 1 tags = %q", got)
	}
	if metas[1].Title != "Second" {
		t.Errorf("page 2 title = %q", metas[1].Title)
	}
	// Cross-property order depends on map iteration; compare as a set.
	if got := strings.Join(metas[1].Tags, ","); !containsSameItems(got, "ddd,published") {
		t.Errorf("page 2 tags = %q", got)
	}
}

// containsSameItems checks comma-separated lists hold the same items in any order.
func containsSameItems(got, want string) bool {
	toSet := func(s string) map[string]bool {
		set := map[string]bool{}
		for _, item := range strings.Split(s, ",") {
			set[item] = true
		}

		return set
	}
	a, b := toSet(got), toSet(want)
	if len(a) != len(b) {
		return false
	}
	for k := range a {
		if !b[k] {
			return false
		}
	}

	return true
}

func TestQueryPagesIncrementalFilter(t *testing.T) {
	rec := &recorder{}
	c := newTestClient(t, rec)
	rec.handler = func(w http.ResponseWriter, r *http.Request) {
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

func TestPageBlocksMapsBlockTypes(t *testing.T) {
	rec := &recorder{}
	c := newTestClient(t, rec)
	rec.handler = func(w http.ResponseWriter, r *http.Request) {
		switch rec.requests {
		case 1:
			writeJSON(t, w, blocksResponse{Results: []apiBlock{
				{Type: model.TypeParagraph, textPayload: textPayload{RichText: []richRun{
					{PlainText: "plain "}, boldRun("bold"),
				}}},
				{Type: model.TypeCode, textPayload: textPayload{
					RichText: []richRun{{PlainText: "const x = 1"}}, Language: "js",
				}},
				{Type: model.TypeBookmark, Bookmark: &urlPayload{URL: "https://a.b"}},
				{Type: model.TypeImage, Image: &imagePayload{External: &urlPayload{URL: "https://img.example/x.png"}}},
				{Type: model.TypeImage, Image: &imagePayload{File: &urlPayload{URL: "https://notion.so/expiring.png"}}},
				{Type: model.TypeChildPage, ChildPage: &childPagePayload{Title: "Specs"}},
				{Type: "morph"},
				{
					Type: model.TypeTable, ID: "tbl-1", HasChildren: true,
					Table: &tablePayload{HasColumnHeader: true},
				},
			}})
		default:
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
	}

	blocks, err := c.PageBlocks("page1")
	if err != nil {
		t.Fatalf("PageBlocks: %v", err)
	}

	wantTypes := []string{
		model.TypeParagraph, model.TypeCode, model.TypeBookmark, model.TypeImage,
		model.TypeImage, model.TypeChildPage, "morph", model.TypeTable,
	}
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
	if blocks[2].URL != "https://a.b" {
		t.Errorf("bookmark url = %q", blocks[2].URL)
	}
	if blocks[3].Internal {
		t.Errorf("external image marked internal")
	}
	if !blocks[4].Internal || blocks[4].URL != "https://notion.so/expiring.png" {
		t.Errorf("internal image = %+v", blocks[4])
	}
	if blocks[5].Title != "Specs" {
		t.Errorf("child page title = %q", blocks[5].Title)
	}

	tbl := blocks[7]
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

func TestRetries429HonoringRetryAfter(t *testing.T) {
	rec := &recorder{}
	c := newTestClient(t, rec)
	rec.handler = func(w http.ResponseWriter, r *http.Request) {
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

func TestNilAndEmptyPropertiesMapSafely(t *testing.T) {
	rec := &recorder{}
	c := newTestClient(t, rec)
	rec.handler = func(w http.ResponseWriter, r *http.Request) {
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

func TestBlocksWithoutPayloadsMapSafely(t *testing.T) {
	rec := &recorder{}
	c := newTestClient(t, rec)
	rec.handler = func(w http.ResponseWriter, r *http.Request) {
		writeJSON(t, w, blocksResponse{Results: []apiBlock{
			{Type: model.TypeBookmark},
			{Type: model.TypeImage},
		}})
	}

	blocks, err := c.PageBlocks("page1")
	if err != nil {
		t.Fatalf("PageBlocks: %v", err)
	}
	if blocks[0].URL != "" {
		t.Errorf("bookmark without payload: url = %q", blocks[0].URL)
	}
	if blocks[1].URL != "" || blocks[1].Internal {
		t.Errorf("image without payload: %+v", blocks[1])
	}
}

func TestRateLimitWithoutRetryAfterUsesBackoff(t *testing.T) {
	rec := &recorder{}
	c := newTestClient(t, rec)
	rec.handler = func(w http.ResponseWriter, r *http.Request) {
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
	rec.handler = func(w http.ResponseWriter, r *http.Request) {
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
	rec.handler = func(w http.ResponseWriter, r *http.Request) {
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
	rec.handler = func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"message":"Invalid data source"}`))
	}

	_, err := c.QueryPages("nope", QueryFilter{})
	if err == nil || !strings.Contains(err.Error(), "Invalid data source") {
		t.Fatalf("err = %v, want surfaced API message", err)
	}
}
