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
