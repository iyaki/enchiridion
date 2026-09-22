// Package notion provides a minimal HTTP client for the Notion API surface
// enchiridion needs: querying a data source and reading page blocks. Every
// outbound request flows through here so the whole chain stays auditable
// (ADR-08).
package notion

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/iyaki/enchiridion/internal/model"
)

const (
	defaultBaseURL = "https://api.notion.com/v1"
	apiVersion     = "2025-09-03"
	pageSize       = 100
	maxAttempts    = 3
	requestTimeout = 30 * time.Second
	maxErrorBody   = 4096
)

// QueryFilter narrows a data source query. Empty fields mean "no filter".
type QueryFilter struct {
	// LastEditedAfter (RFC3339) keeps only pages edited on or after the
	// instant — the incremental sync filter.
	LastEditedAfter string
}

// Client talks to the Notion API over HTTP.
type Client struct {
	baseURL string
	token   string
	http    *http.Client
	sleep   func(time.Duration) // injectable for tests
}

// NewClient returns a client authenticated with token.
func NewClient(token string) *Client {
	return &Client{
		baseURL: defaultBaseURL,
		token:   token,
		http:    &http.Client{Timeout: requestTimeout},
		sleep:   time.Sleep,
	}
}

// QueryPages returns every page of the data source — or only those edited on
// or after filter.LastEditedAfter — following pagination.
func (c *Client) QueryPages(dataSourceID string, filter QueryFilter) ([]model.PageMeta, error) {
	var metas []model.PageMeta
	cursor := ""
	for {
		body := map[string]any{"page_size": pageSize}
		if cursor != "" {
			body["start_cursor"] = cursor
		}
		if filter.LastEditedAfter != "" {
			body["filter"] = map[string]any{
				"timestamp":        "last_edited_time",
				"last_edited_time": map[string]string{"on_or_after": filter.LastEditedAfter},
			}
		}

		var out struct {
			Results    []page `json:"results"`
			HasMore    bool   `json:"has_more"`
			NextCursor string `json:"next_cursor"`
		}
		if err := c.do("POST", "/data_sources/"+dataSourceID+"/query", body, &out); err != nil {
			return nil, err
		}
		for _, p := range out.Results {
			metas = append(metas, p.meta())
		}
		if !out.HasMore {
			return metas, nil
		}
		cursor = out.NextCursor
	}
}

// PageBlocks returns the distilled content blocks of a page, following
// pagination, resolving table rows, and recursively flattening the children of
// every block in document order (ADR-17) so the renderer stays flat.
func (c *Client) PageBlocks(pageID string) ([]model.Block, error) {
	raws, err := c.children(pageID)
	if err != nil {
		return nil, err
	}

	return c.distillAll(raws)
}

// Ping verifies that the token is accepted and the data source is reachable:
// the check doctor performs before a sync, without reading any content.
func (c *Client) Ping(dataSourceID string) error {
	return c.do("GET", "/data_sources/"+dataSourceID, nil, &struct {
		Object string `json:"object"`
	}{})
}

// distillAll flattens raw blocks (and, recursively, their children) into a
// single document-ordered slice (ADR-17). Pure containers (synced_block,
// column_list, column) contribute no block of their own — only their children.
func (c *Client) distillAll(raws []apiBlock) ([]model.Block, error) {
	var blocks []model.Block
	for _, raw := range raws {
		switch {
		case raw.Type == "synced_block":
			// Own children or the original's: fully handled here.
			var err error
			blocks, err = c.distillSyncedBlock(blocks, raw)
			if err != nil {
				return nil, err
			}

			continue
		case raw.Type == "column_list", raw.Type == "column":
			// Container: emit children only.
		default:
			blk, err := c.distill(raw)
			if err != nil {
				return nil, err
			}
			blocks = append(blocks, blk)
		}
		var err error
		blocks, err = c.appendChildren(blocks, raw)
		if err != nil {
			return nil, err
		}
	}

	return blocks, nil
}

// distillSyncedBlock splices synced content in document order: an instance
// (synced_from set) inherits the original block's children; an original
// contributes its own. When the source is unreachable the block surfaces as
// a visible comment — never silence, never a failed page (ADR-17).
func (c *Client) distillSyncedBlock(blocks []model.Block, raw apiBlock) ([]model.Block, error) {
	source := raw.ID
	if raw.SyncedBlock != nil && raw.SyncedBlock.SyncedFrom != nil && raw.SyncedBlock.SyncedFrom.BlockID != "" {
		source = raw.SyncedBlock.SyncedFrom.BlockID
	}
	kids, err := c.children(source)
	if err != nil {
		return append(blocks, model.Block{Type: raw.Type}), nil
	}
	flat, err := c.distillAll(kids)
	if err != nil {
		return nil, err
	}

	return append(blocks, flat...), nil
}

// appendChildren fetches raw's children, flattens them, and appends the
// result in document order.
func (c *Client) appendChildren(blocks []model.Block, raw apiBlock) ([]model.Block, error) {
	// table_row children are cells, resolved by distillTable.
	if !raw.HasChildren || raw.Type == model.TypeTable {
		return blocks, nil
	}

	kids, err := c.children(raw.ID)
	if err != nil {
		return nil, err
	}
	flat, err := c.distillAll(kids)
	if err != nil {
		return nil, err
	}

	return append(blocks, flat...), nil
}

// children fetches every child block of a block/page, following pagination.
func (c *Client) children(blockID string) ([]apiBlock, error) {
	var raws []apiBlock
	cursor := ""
	for {
		path := "/blocks/" + blockID + "/children?page_size=" + strconv.Itoa(pageSize)
		if cursor != "" {
			path += "&start_cursor=" + cursor
		}
		var out struct {
			Results    []apiBlock `json:"results"`
			HasMore    bool       `json:"has_more"`
			NextCursor string     `json:"next_cursor"`
		}
		if err := c.do("GET", path, nil, &out); err != nil {
			return nil, err
		}
		raws = append(raws, out.Results...)
		if !out.HasMore {
			return raws, nil
		}
		cursor = out.NextCursor
	}
}

// distill maps a raw API block into a model.Block, delegating to focused
// helpers so each stays simple.
func (c *Client) distill(raw apiBlock) (model.Block, error) {
	blk := model.Block{Type: raw.Type}
	if tp := raw.text(); tp != nil {
		blk.RichText = richText(tp.RichText)
		blk.Language = tp.Language
	}
	switch {
	case raw.Type == model.TypeToDo:
		if raw.ToDo != nil {
			blk.RichText = richText(raw.ToDo.RichText)
			blk.Checked = raw.ToDo.Checked
		}
	case isLinkBlock(raw.Type):
		blk.URL = raw.url(raw.Type)
	case raw.media() != nil:
		distillMedia(&blk, raw.media())
	case raw.Type == model.TypeChildPage:
		blk.Title = raw.ChildPage.Title
	case raw.Type == model.TypeTable:
		return c.distillTable(raw)
	}

	return blk, nil
}

func isLinkBlock(blockType string) bool {
	switch blockType {
	case model.TypeBookmark, model.TypeEmbed, model.TypeLinkPreview:
		return true
	}

	return false
}

// distillMedia fills URL/Internal for block types whose payload is the
// external|file url shape: image, pdf, file, video.
func distillMedia(blk *model.Block, img *imagePayload) {
	if img == nil {
		return
	}
	switch {
	case img.External != nil:
		blk.URL = img.External.URL
	case img.File != nil:
		blk.URL = img.File.URL
		blk.Internal = true
	}
}

// distillTable distills a table block, resolving its rows into cells.
func (c *Client) distillTable(raw apiBlock) (model.Block, error) {
	blk := model.Block{Type: raw.Type}
	if raw.Table != nil {
		blk.HasHeader = raw.Table.HasColumnHeader
	}
	if !raw.HasChildren {
		return blk, nil
	}

	rows, err := c.tableRows(raw.ID)
	if err != nil {
		return model.Block{}, err
	}
	blk.Rows = rows

	return blk, nil
}

func (c *Client) tableRows(tableID string) ([][]model.Cell, error) {
	raws, err := c.children(tableID)
	if err != nil {
		return nil, err
	}

	var rows [][]model.Cell
	for _, raw := range raws {
		if raw.Type != "table_row" {
			continue
		}
		cells := make([]model.Cell, 0, len(raw.TableRow.Cells))
		for _, runs := range raw.TableRow.Cells {
			cells = append(cells, model.Cell(richText(runs)))
		}
		rows = append(rows, cells)
	}

	return rows, nil
}

// do performs one API call, retrying transient failures (429 honoring
// Retry-After, 5xx, transport errors) up to maxAttempts and failing fast on
// credential errors.
func (c *Client) do(method, path string, body any, out any) error {
	var lastErr error
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		req, err := c.request(method, path, body)
		if err != nil {
			return err
		}

		resp, err := c.http.Do(req)
		if err != nil {
			lastErr = fmt.Errorf("notion request failed: %w", err)
			c.pause(attempt)

			continue
		}

		switch {
		case resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden:
			drain(resp)

			return fmt.Errorf("notion rejected the credentials (HTTP %d): "+
				"check NOTION_TOKEN and that the integration has access to this data source",
				resp.StatusCode)
		case resp.StatusCode == http.StatusTooManyRequests:
			lastErr = fmt.Errorf("notion rate limit exceeded (HTTP 429)")
			c.pauseFor(attempt, retryAfter(resp))
			drain(resp)
		case resp.StatusCode >= http.StatusInternalServerError:
			lastErr = fmt.Errorf("notion server error (HTTP %d)", resp.StatusCode)
			c.pause(attempt)
			drain(resp)
		case resp.StatusCode >= http.StatusBadRequest:
			msg, _ := io.ReadAll(io.LimitReader(resp.Body, maxErrorBody))
			drain(resp)

			return fmt.Errorf("notion API error (HTTP %d): %s", resp.StatusCode, strings.TrimSpace(string(msg)))
		default:
			err = json.NewDecoder(resp.Body).Decode(out)
			drain(resp)

			return err
		}
	}

	return fmt.Errorf("%w (gave up after %d attempts)", lastErr, maxAttempts)
}

func (c *Client) request(method, path string, body any) (*http.Request, error) {
	var reader io.Reader
	if body != nil {
		buf, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reader = bytes.NewReader(buf)
	}

	req, err := http.NewRequest(method, c.baseURL+path, reader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Notion-Version", apiVersion)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

func (c *Client) pause(attempt int) {
	c.sleep(time.Duration(attempt) * time.Second)
}

func (c *Client) pauseFor(attempt int, d time.Duration) {
	if d <= 0 {
		c.pause(attempt)

		return
	}
	c.sleep(d)
}

func retryAfter(resp *http.Response) time.Duration {
	secs, err := strconv.Atoi(resp.Header.Get("Retry-After"))
	if err != nil || secs < 0 {
		return 0
	}

	return time.Duration(secs) * time.Second
}

func drain(resp *http.Response) {
	_, _ = io.Copy(io.Discard, resp.Body)
	_ = resp.Body.Close()
}
