package pull

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// buildTarball packs name->content files under a single top-level directory,
// mirroring the layout GitHub's repository tarballs use.
func buildTarball(t *testing.T, top string, files map[string]string) []byte {
	t.Helper()
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	tw := tar.NewWriter(gz)
	for name, content := range files {
		err := tw.WriteHeader(&tar.Header{
			Name: top + "/" + name,
			Mode: 0o644,
			Size: int64(len(content)),
		})
		if err != nil {
			t.Fatalf("write header %s: %v", name, err)
		}
		if _, err := tw.Write([]byte(content)); err != nil {
			t.Fatalf("write content %s: %v", name, err)
		}
	}
	if err := tw.Close(); err != nil {
		t.Fatalf("close tar: %v", err)
	}
	if err := gz.Close(); err != nil {
		t.Fatalf("close gzip: %v", err)
	}

	return buf.Bytes()
}

func newTestClient(t *testing.T, handler http.HandlerFunc) *Client {
	t.Helper()
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)

	c := NewClient("tok")
	c.baseURL = srv.URL

	return c
}

func TestPullExtractsOnlyMirrorDirs(t *testing.T) {
	var gotPath, gotAuth string
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		gotPath, gotAuth = r.URL.Path, r.Header.Get("Authorization")
		w.Header().Set("Content-Type", "application/x-gtar")
		_, _ = w.Write(buildTarball(t, "ench-main-abc", map[string]string{
			"data/knowledge/a.md":       "A",
			"data/knowledge/sub/b.md":   "B",
			"data/tools/t.md":           "T",
			"data/other.txt":            "junk",
			"README.md":                 "junk",
			"specs/architecture.md":     "junk",
			"data/knowledge/evil/xx.md": "x",
		}))
	})

	out := t.TempDir()
	stats, err := c.Pull("iyaki/enchiridion", out)
	if err != nil {
		t.Fatalf("Pull: %v", err)
	}

	if gotPath != "/repos/iyaki/enchiridion/tarball" {
		t.Errorf("request path = %q", gotPath)
	}
	if gotAuth != "Bearer tok" {
		t.Errorf("Authorization = %q", gotAuth)
	}
	assertFiles(t, out, map[string]string{
		"knowledge/a.md":     "A",
		"knowledge/sub/b.md": "B",
		"tools/t.md":         "T",
	})
	assertAbsent(t, out, "other.txt", "README.md", "architecture.md")
	if stats.Knowledge != 3 || stats.Tools != 1 {
		t.Errorf("stats = %+v, want knowledge=3 tools=1", stats)
	}
}

// assertFiles fails unless every relative path under root holds want.
func assertFiles(t *testing.T, root string, want map[string]string) {
	t.Helper()
	for rel, content := range want {
		got, err := os.ReadFile(filepath.Join(root, rel))
		if err != nil {
			t.Fatalf("read %s: %v", rel, err)
		}
		if string(got) != content {
			t.Errorf("%s = %q, want %q", rel, got, content)
		}
	}
}

// assertAbsent fails if any relative path exists under root.
func assertAbsent(t *testing.T, root string, rels ...string) {
	t.Helper()
	for _, rel := range rels {
		if _, err := os.Stat(filepath.Join(root, rel)); !os.IsNotExist(err) {
			t.Errorf("%s should not be extracted", rel)
		}
	}
}

func TestPullReplacesTargetDirsForDeletionPropagation(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write(buildTarball(t, "ench-main", map[string]string{
			"data/knowledge/fresh.md": "fresh",
		}))
	})

	out := t.TempDir()
	for _, stale := range []string{"knowledge/stale.md", "tools/old.md"} {
		path := filepath.Join(out, stale)
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, []byte("stale"), 0o644); err != nil {
			t.Fatal(err)
		}
	}

	if _, err := c.Pull("iyaki/enchiridion", out); err != nil {
		t.Fatalf("Pull: %v", err)
	}
	if _, err := os.Stat(filepath.Join(out, "knowledge/stale.md")); !os.IsNotExist(err) {
		t.Errorf("stale knowledge file survived the pull")
	}
	if _, err := os.Stat(filepath.Join(out, "tools/old.md")); !os.IsNotExist(err) {
		t.Errorf("stale tools file survived the pull")
	}
	if _, err := os.Stat(filepath.Join(out, "knowledge/fresh.md")); err != nil {
		t.Errorf("fresh file missing: %v", err)
	}
}

func TestPullKeepsOldMirrorWhenDownloadFails(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	out := t.TempDir()
	keep := filepath.Join(out, "knowledge", "keep.md")
	if err := os.MkdirAll(filepath.Dir(keep), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(keep, []byte("keep"), 0o644); err != nil {
		t.Fatal(err)
	}

	if _, err := c.Pull("iyaki/enchiridion", out); err == nil {
		t.Fatal("Pull over a server error should fail")
	}
	content, err := os.ReadFile(keep)
	if err != nil || string(content) != "keep" {
		t.Errorf("existing mirror damaged by failed pull: %v", err)
	}
}

func TestPullRejectsEscapingTarEntries(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write(buildTarball(t, "ench-main", map[string]string{
			"data/knowledge/../../../evil.md": "evil",
			"data/knowledge/ok.md":            "ok",
		}))
	})

	root := t.TempDir()
	out := filepath.Join(root, "out")
	if _, err := c.Pull("iyaki/enchiridion", out); err == nil {
		t.Fatal("Pull over an escaping entry should fail")
	}
	if _, err := os.Stat(filepath.Join(root, "evil.md")); !os.IsNotExist(err) {
		t.Errorf("entry escaped the target directory")
	}
}

func TestPullSurfacesCredentialErrors(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(`{"message":"Bad credentials"}`))
	})

	_, err := c.Pull("iyaki/enchiridion", t.TempDir())
	if err == nil || !strings.Contains(err.Error(), "GITHUB_TOKEN") {
		t.Fatalf("err = %v, want a GITHUB_TOKEN-oriented message", err)
	}
}

func TestPullSurfacesNotFoundAsAccessProblem(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	_, err := c.Pull("iyaki/enchiridion", t.TempDir())
	if err == nil || !strings.Contains(err.Error(), "access") {
		t.Fatalf("err = %v, want an access-oriented message", err)
	}
}

func TestPullRejectsCorruptTarball(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte("this is not gzip"))
	})

	if _, err := c.Pull("iyaki/enchiridion", t.TempDir()); err == nil || !strings.Contains(err.Error(), "open tarball") {
		t.Fatalf("err = %v, want an open-tarball error", err)
	}
}

func TestPullRejectsCorruptTarEntries(t *testing.T) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	_, _ = gz.Write([]byte("not a tar stream"))
	_ = gz.Close()

	c := newTestClient(t, func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write(buf.Bytes())
	})

	if _, err := c.Pull("iyaki/enchiridion", t.TempDir()); err == nil || !strings.Contains(err.Error(), "read tarball") {
		t.Fatalf("err = %v, want a read-tarball error", err)
	}
}

func TestPullSurfacesTransportErrors(t *testing.T) {
	c := NewClient("tok")
	c.baseURL = "http://127.0.0.1:1" // nothing listens there

	_, err := c.Pull("iyaki/enchiridion", t.TempDir())
	if err == nil || !strings.Contains(err.Error(), "github request failed") {
		t.Fatalf("err = %v, want a transport error", err)
	}
}

func TestMirrorRelPath(t *testing.T) {
	cases := map[string]struct {
		in  string
		rel string
		ok  bool
	}{
		"knowledge file":            {"ench-main/data/knowledge/a.md", "knowledge/a.md", true},
		"tools file":                {"ench-main/data/tools/t.md", "tools/t.md", true},
		"non-mirror file":           {"ench-main/README.md", "", false},
		"other data file":           {"ench-main/data/other.txt", "", false},
		"lookalike prefix":          {"ench-main/data/knowledge-evil/x.md", "", false},
		"top-level entry only":      {"ench-main", "", false},
		"absolute-looking name":     {"/ench-main/data/tools/t.md", "tools/t.md", true},
		"nested under subdirectory": {"ench-main/data/knowledge/sub/b.md", "knowledge/sub/b.md", true},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			rel, ok := mirrorRelPath(tc.in)
			if ok != tc.ok || rel != tc.rel {
				t.Fatalf("mirrorRelPath(%q) = (%q, %t), want (%q, %t)", tc.in, rel, ok, tc.rel, tc.ok)
			}
		})
	}
}

func TestPullFailsWhenTargetIsNotADirectory(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write(buildTarball(t, "ench-main", map[string]string{
			"data/knowledge/a.md": "A",
		}))
	})

	blocker := filepath.Join(t.TempDir(), "blocker")
	if err := os.WriteFile(blocker, []byte("a file, not a directory"), 0o644); err != nil {
		t.Fatal(err)
	}

	if _, err := c.Pull("iyaki/enchiridion", blocker); err == nil {
		t.Fatal("Pull over a file-shaped target should fail")
	}
}
