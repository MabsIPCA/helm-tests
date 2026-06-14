package helm_test

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/MabsIPCA/helm-tests/helm_fetcher/helm"
)

// writeChart creates dir and an (empty) Chart.yaml inside it.
func writeChart(t *testing.T, dir string) {
	t.Helper()
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "Chart.yaml"), []byte("name: x\n"), 0o644); err != nil {
		t.Fatal(err)
	}
}

func TestIsVendoredSubchart(t *testing.T) {
	root := t.TempDir()

	parent := filepath.Join(root, "adp")
	sub := filepath.Join(parent, "charts", "adp-agent-config")
	nested := filepath.Join(sub, "charts", "leaf")
	// A "charts/" dir whose parent is NOT a chart (monorepo of standalone charts).
	standalone := filepath.Join(root, "repo", "charts", "free-standing")

	writeChart(t, parent)
	writeChart(t, sub)
	writeChart(t, nested)
	writeChart(t, standalone)

	cases := []struct {
		name string
		dir  string
		want bool
	}{
		{"parent umbrella chart", parent, false},
		{"vendored subchart", sub, true},
		{"nested vendored subchart", nested, true},
		{"standalone chart under non-chart charts dir", standalone, false},
	}
	for _, c := range cases {
		if got := helm.IsVendoredSubchart(c.dir); got != c.want {
			t.Errorf("%s: IsVendoredSubchart(%q) = %v, want %v", c.name, c.dir, got, c.want)
		}
	}
}

func TestFindCharts_SkipsRelativeComponents(t *testing.T) {
	root := t.TempDir()
	// Dev-monorepo layout (like tmforum-oda/oda-canvas): no root Chart.yaml,
	// umbrella + components are siblings under charts/, wired via file://.
	umbrella := filepath.Join(root, "charts", "canvas-oda")
	component := filepath.Join(root, "charts", "api-operator-istio")
	standalone := filepath.Join(root, "charts", "unrelated")

	writeChart(t, umbrella)
	writeChart(t, component)
	writeChart(t, standalone)
	// canvas-oda depends on api-operator-istio via a relative file:// path.
	dep := "dependencies:\n  - name: api-operator-istio\n    repository: 'file://../api-operator-istio'\n"
	if err := os.WriteFile(filepath.Join(umbrella, "Chart.yaml"), []byte("name: canvas-oda\n"+dep), 0o644); err != nil {
		t.Fatal(err)
	}

	got := helm.FindCharts(root)
	for _, g := range got {
		if g == component {
			t.Errorf("FindCharts returned the file:// component %q; it should be skipped", component)
		}
	}
	// umbrella and the unrelated standalone chart must still be discovered.
	for _, want := range []string{umbrella, standalone} {
		found := false
		for _, g := range got {
			if g == want {
				found = true
			}
		}
		if !found {
			t.Errorf("FindCharts omitted %q (got %v)", want, got)
		}
	}
}

func TestFindCharts_SkipsVendoredSubcharts(t *testing.T) {
	root := t.TempDir()
	parent := filepath.Join(root, "adp")
	sub := filepath.Join(parent, "charts", "adp-agent-config")
	standalone := filepath.Join(root, "standalone")

	writeChart(t, parent)
	writeChart(t, sub)
	writeChart(t, standalone)

	got := helm.FindCharts(root)
	sort.Strings(got)
	want := []string{parent, standalone}
	sort.Strings(want)

	if len(got) != len(want) {
		t.Fatalf("FindCharts returned %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("FindCharts[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}

func TestRunHelmTemplateWithSets_CmdStr(t *testing.T) {
	// cmdStr is always built before exec; helm binary need not be present for this test.
	cmdStr, _, _ := helm.RunHelmTemplateWithSets(
		"/chart",
		[]string{"/values.yaml"},
		[]string{"foo=bar", "baz=qux"},
	)
	if !strings.HasPrefix(cmdStr, "helm template test /chart") {
		t.Errorf("unexpected cmdStr prefix: %q", cmdStr)
	}
	if !strings.Contains(cmdStr, "-f /values.yaml") {
		t.Errorf("cmdStr missing -f flag: %q", cmdStr)
	}
	if !strings.Contains(cmdStr, "--set foo=bar") {
		t.Errorf("cmdStr missing --set foo=bar: %q", cmdStr)
	}
	if !strings.Contains(cmdStr, "--set baz=qux") {
		t.Errorf("cmdStr missing --set baz=qux: %q", cmdStr)
	}
}

func TestRunHelmTemplateWithSets_NoExtras(t *testing.T) {
	cmdStr, _, _ := helm.RunHelmTemplateWithSets("/chart", nil, nil)
	want := "helm template test /chart"
	if cmdStr != want {
		t.Errorf("got %q, want %q", cmdStr, want)
	}
}

func TestIsGitCryptEncrypted(t *testing.T) {
	root := t.TempDir()
	magic := append([]byte{0x00, 'G', 'I', 'T', 'C', 'R', 'Y', 'P', 'T', 0x00}, []byte("\x01\x02\x03nonce+ciphertext")...)
	write := func(name string, content []byte) string {
		p := filepath.Join(root, name)
		if err := os.WriteFile(p, content, 0o644); err != nil {
			t.Fatal(err)
		}
		return p
	}

	enc := write("secret-values.yaml", magic)
	plain := write("values-prod.yaml", []byte("replicas: 3\nimage: nginx\n"))
	short := write("tiny.yaml", []byte("a: 1")) // shorter than the magic header
	looksClose := write("decoy.yaml", []byte("GITCRYPT but not at the start"))

	cases := []struct {
		path string
		want bool
	}{
		{enc, true},
		{plain, false},
		{short, false},
		{looksClose, false},
		{filepath.Join(root, "does-not-exist.yaml"), false},
	}
	for _, c := range cases {
		if got := helm.IsGitCryptEncrypted(c.path); got != c.want {
			t.Errorf("IsGitCryptEncrypted(%q) = %v, want %v", c.path, got, c.want)
		}
	}
}

func TestIsParseableValuesFile(t *testing.T) {
	root := t.TempDir()
	write := func(name string, content []byte) string {
		p := filepath.Join(root, name)
		if err := os.WriteFile(p, content, 0o644); err != nil {
			t.Fatal(err)
		}
		return p
	}
	magic := append([]byte{0x00, 'G', 'I', 'T', 'C', 'R', 'Y', 'P', 'T', 0x00}, []byte("\x01\x02ciphertext")...)

	cases := []struct {
		name    string
		content []byte
		want    bool
	}{
		{"valid map", []byte("replicas: 3\nimage:\n  tag: v1\n"), true},
		{"empty file", []byte(""), true},
		{"whitespace only", []byte("\n  \n"), true},
		{"null doc", []byte("null\n"), true},
		{"git-crypt encrypted", magic, false},
		{"control characters", []byte("foo: \x07\x1b\x00bar\n"), false},
		{"broken yaml", []byte("foo: [unterminated\n"), false},
		{"scalar not a map", []byte("just-a-string\n"), false},
		{"sequence not a map", []byte("- a\n- b\n"), false},
	}
	for _, c := range cases {
		p := write(c.name+".yaml", c.content)
		if got := helm.IsParseableValuesFile(p); got != c.want {
			t.Errorf("%s: IsParseableValuesFile = %v, want %v", c.name, got, c.want)
		}
	}
	if helm.IsParseableValuesFile(filepath.Join(root, "missing.yaml")) {
		t.Error("missing file should not be parseable")
	}
}

func TestFindValuesFiles_PrefiltersUnparseable(t *testing.T) {
	chart := t.TempDir()
	if err := os.WriteFile(filepath.Join(chart, "Chart.yaml"), []byte("name: x\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	magic := append([]byte{0x00, 'G', 'I', 'T', 'C', 'R', 'Y', 'P', 'T', 0x00}, []byte("ciphertext")...)
	mustWrite := func(name string, b []byte) {
		if err := os.WriteFile(filepath.Join(chart, name), b, 0o644); err != nil {
			t.Fatal(err)
		}
	}
	mustWrite("secret-values.yaml", magic)                          // encrypted -> dropped
	mustWrite("broken-values.yaml", []byte("a: [oops\n"))           // malformed -> dropped
	mustWrite("values-prod.yaml", []byte("replicas: 3\n"))          // valid -> kept
	mustWrite("values-staging.yaml", []byte("image:\n  tag: v2\n")) // valid -> kept

	got := helm.FindValuesFiles(chart)
	gotBase := make([]string, 0, len(got))
	for _, g := range got {
		gotBase = append(gotBase, filepath.Base(g))
	}
	sort.Strings(gotBase)
	want := []string{"values-prod.yaml", "values-staging.yaml"}
	if strings.Join(gotBase, ",") != strings.Join(want, ",") {
		t.Errorf("FindValuesFiles = %v, want %v (unparseable files prefiltered)", gotBase, want)
	}
}
