package main

import (
	"strings"
	"testing"

	"helm.sh/helm/v3/pkg/chart"
)

func TestAddID_InjectsMarkerBeforeAPIVersion(t *testing.T) {
	f := &chart.File{
		Name: "templates/test.yaml",
		Data: []byte("apiVersion: v1\nkind: ConfigMap\nmetadata:\n  name: test"),
	}
	result := addID(f)
	lines := strings.Split(string(result.Data), "\n")

	found := false
	for i, l := range lines {
		if strings.Contains(l, "# KICS_HELM_ID_0:") {
			found = true
			if i+1 >= len(lines) || !strings.Contains(lines[i+1], "apiVersion:") {
				t.Errorf("KICS_HELM_ID_0: must immediately precede apiVersion:, got next line: %q", lines[i+1])
			}
			break
		}
	}
	if !found {
		t.Errorf("KICS_HELM_ID_0: marker not found in:\n%s", string(result.Data))
	}
}

func TestAddID_MultipleAPIVersions(t *testing.T) {
	f := &chart.File{
		Name: "templates/multi.yaml",
		Data: []byte("apiVersion: v1\nkind: ConfigMap\n---\napiVersion: apps/v1\nkind: Deployment"),
	}
	result := addID(f)
	content := string(result.Data)
	if !strings.Contains(content, "# KICS_HELM_ID_0:") {
		t.Error("missing KICS_HELM_ID_0:")
	}
	if !strings.Contains(content, kicsHelmID) {
		t.Error("missing additional KICS_HELM_ID marker")
	}
}
