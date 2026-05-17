package helm_test

import (
	"strings"
	"testing"

	"github.com/MabsIPCA/helm-tests/helm_fetcher/helm"
)

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
