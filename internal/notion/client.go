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
// pagination and resolving table rows recursively so the renderer stays flat.
func (c *Client) PageBlocks(pageID string) ([]model.Block, error) {
	raws, err := c.children(pageID)
	if err != nil {
		return nil, err
	}

	blocks := make([]model.Block, 0, len(raws))
	for _, raw := range raws {
		blk, err := c.distill(raw)
		if err != nil {
			return nil, err
		}
		blocks = append(blocks, blk)
	}

	return blocks, nil
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
	switch {
	case isTextBlock(raw.Type):
		blk.RichText = richText(raw.text())
	case raw.Type == model.TypeCode:
		blk.RichText = richText(raw.text())
		blk.Language = raw.Language
	case isLinkBlock(raw.Type):
		blk.URL = raw.url(raw.Type)
	case raw.Type == model.TypeImage:
		distillImage(&blk, raw.Image)
	case raw.Type == model.TypeChildPage:
		blk.Title = raw.ChildPage.Title
	case raw.Type == model.TypeTable:
		if err := c.distillTable(&blk, raw); err != nil {
			return model.Block{}, err
		}
	}

	return blk, nil
}

func isTextBlock(blockType string) bool {
	switch blockType {
	case model.TypeParagraph, model.TypeHeading1, model.TypeHeading2,
		model.TypeHeading3, model.TypeBulletedItem, model.TypeNumberedItem,
		model.TypeQuote, model.TypeCallout, model.TypeToggle:
		return true
	}

	return false
}

func isLinkBlock(blockType string) bool {
	switch blockType {
	case model.TypeBookmark, model.TypeEmbed, model.TypeLinkPreview:
		return true
	}

	return false
}

func distillImage(blk *model.Block, img *imagePayload) {
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

func (c *Client) distillTable(blk *model.Block, raw apiBlock) error {
	blk.HasHeader = raw.Table.HasColumnHeader
	if !raw.HasChildren {
		return nil
	}

	rows, err := c.tableRows(raw.ID)
	if err != nil {
		return err
	}
	blk.Rows = rows

	return nil
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
