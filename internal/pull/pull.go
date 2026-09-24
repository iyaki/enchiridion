// Package pull downloads the published mirror from the enchiridion GitHub
// repository so a consumer project can vendor it into its own tree — no
// Notion credentials and no shared machine with an enchiridion installation
// (ADR-18). The source of the pull is the repository's default-branch
// tarball; only data/knowledge and data/tools are extracted.
package pull

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	// DefaultRepo is the distribution repository carrying the committed mirror.
	DefaultRepo = "iyaki/enchiridion"

	apiBase      = "https://api.github.com"
	maxErrorBody = 4 << 10
	topDirParts  = 2 // tarball name split: top-level dir + remainder
	dirPerm      = 0o755
	filePerm     = 0o644
)

// Stats reports how many mirror files were written per directory.
type Stats struct {
	Knowledge int
	Tools     int
}

// Client fetches the published mirror from the GitHub API.
type Client struct {
	baseURL string
	http    *http.Client
	token   string
}

// NewClient returns a client for the GitHub API. The token is optional: with
// it, requests are authenticated; without it, requests stay anonymous, which
// works for public repositories (ADR-19).
func NewClient(token string) *Client {
	return &Client{
		baseURL: apiBase,
		http:    &http.Client{},
		token:   token,
	}
}

// Pull replaces out/knowledge and out/tools with the mirror from repo's
// default branch. The download completes before anything on disk is touched,
// so a failed pull never damages an existing mirror.
func (c *Client) Pull(repo, out string) (Stats, error) {
	req, err := http.NewRequest(http.MethodGet, c.baseURL+"/repos/"+repo+"/tarball", nil)
	if err != nil {
		return Stats{}, err
	}
	c.authorize(req)

	resp, err := c.http.Do(req)
	if err != nil {
		return Stats{}, fmt.Errorf("github request failed: %w", err)
	}
	defer func() { _, _ = io.Copy(io.Discard, resp.Body); _ = resp.Body.Close() }()

	if err := c.checkStatus(resp, repo); err != nil {
		return Stats{}, err
	}

	blob, err := io.ReadAll(resp.Body)
	if err != nil {
		return Stats{}, fmt.Errorf("download mirror tarball: %w", err)
	}

	if err := os.RemoveAll(filepath.Join(out, "knowledge")); err != nil {
		return Stats{}, err
	}
	if err := os.RemoveAll(filepath.Join(out, "tools")); err != nil {
		return Stats{}, err
	}

	return extract(bytes.NewReader(blob), out)
}

// Check reports whether the token can see repo on the GitHub API: the
// reachability surface doctor verifies before suggesting a pull, without
// downloading anything.
func (c *Client) Check(repo string) error {
	req, err := http.NewRequest(http.MethodGet, c.baseURL+"/repos/"+repo, nil)
	if err != nil {
		return err
	}
	c.authorize(req)

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("github request failed: %w", err)
	}
	defer func() { _, _ = io.Copy(io.Discard, resp.Body); _ = resp.Body.Close() }()

	return c.checkStatus(resp, repo)
}

// authorize sets the standard GitHub API headers shared by every request.
// Without a token the Authorization header is omitted: GitHub rejects a bare
// "Bearer " as invalid credentials, while anonymous requests work for public
// repositories.
func (c *Client) authorize(req *http.Request) {
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
}

// checkStatus maps GitHub error statuses to actionable messages; nil means
// the response body is good to consume.
func (c *Client) checkStatus(resp *http.Response, repo string) error {
	switch {
	case resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden:
		return fmt.Errorf("github rejected the credentials (HTTP %d): "+
			"check GITHUB_TOKEN and that it can read the %s repository",
			resp.StatusCode, repo)
	case resp.StatusCode == http.StatusNotFound:
		return fmt.Errorf("repository or mirror not accessible (HTTP 404): "+
			"check GITHUB_TOKEN access and ENCHIRIDION_REPO (%s)", repo)
	case resp.StatusCode != http.StatusOK:
		msg, _ := io.ReadAll(io.LimitReader(resp.Body, maxErrorBody))

		return fmt.Errorf("github error (HTTP %d): %s", resp.StatusCode, strings.TrimSpace(string(msg)))
	}

	return nil
}

// extract unpacks the mirror files of a repository tarball into out.
func extract(r io.Reader, out string) (Stats, error) {
	gz, err := gzip.NewReader(r)
	if err != nil {
		return Stats{}, fmt.Errorf("open tarball: %w", err)
	}

	var stats Stats
	tr := tar.NewReader(gz)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return stats, fmt.Errorf("read tarball: %w", err)
		}
		if hdr.Typeflag != tar.TypeReg {
			continue
		}
		rel, ok := mirrorRelPath(hdr.Name)
		if !ok {
			continue
		}
		target := filepath.Join(out, rel)
		if !strings.HasPrefix(target, filepath.Clean(out)+string(os.PathSeparator)) {
			return stats, fmt.Errorf("tarball entry %q escapes the target directory", hdr.Name)
		}
		if err := writeFile(tr, target); err != nil {
			return stats, err
		}
		if strings.HasPrefix(rel, "knowledge/") {
			stats.Knowledge++
		} else {
			stats.Tools++
		}
	}

	return stats, nil
}

// mirrorRelPath strips the tarball's top-level directory and reports the
// path relative to out for mirror files (data/knowledge, data/tools);
// ok is false for everything else.
func mirrorRelPath(name string) (rel string, ok bool) {
	parts := strings.SplitN(filepath.ToSlash(strings.TrimPrefix(name, "/")), "/", topDirParts)
	if len(parts) < topDirParts {
		return "", false
	}
	rel = parts[1]
	switch {
	case strings.HasPrefix(rel, "data/knowledge/"), strings.HasPrefix(rel, "data/tools/"):
		return strings.TrimPrefix(rel, "data/"), true
	}

	return "", false
}

func writeFile(r io.Reader, target string) (err error) {
	if err := os.MkdirAll(filepath.Dir(target), dirPerm); err != nil {
		return err
	}
	// #nosec G304 -- target is confined to out by the zip-slip check above.
	f, err := os.OpenFile(target, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, filePerm)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := f.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	_, err = io.Copy(f, r)

	return err
}
