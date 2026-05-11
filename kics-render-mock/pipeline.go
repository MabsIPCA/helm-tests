package main

import (
	"fmt"
	"io"
	"log"
	"path/filepath"
	"strings"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/chartutil"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/cli/values"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/release"
)

const kicsHelmID = "# KICS_HELM_ID_"

var helmSettings = cli.New()

// renderResult is the raw output of one runOnce call.
type renderResult struct {
	rel  *release.Release
	err  error
	logs []string
}

// addID mirrors KICS's addID: injects a # KICS_HELM_ID_<line>: comment before
// each "apiVersion:" line in a template file. Mutates file.Data in place.
func addID(file *chart.File) *chart.File {
	split := strings.Split(string(file.Data), "\n")
	for i := 0; i < len(split); i++ {
		if strings.Contains(split[i], "apiVersion:") {
			split = append(split, "")
			copy(split[i+1:], split[i:])
			split[i] = fmt.Sprintf("%s%d:", kicsHelmID, i)
			i++
		}
	}
	file.Data = []byte(strings.Join(split, "\n"))
	return file
}

// setID mirrors KICS's setID: applies addID to all templates in the chart and
// its dependencies.
func setID(c *chart.Chart) *chart.Chart {
	for _, t := range c.Templates {
		addID(t)
	}
	for _, dep := range c.Dependencies() {
		setID(dep)
	}
	return c
}

// updateName mirrors KICS's updateName: prefixes all template file names with
// the chart name hierarchy and collects them into a flat slice.
func updateName(files []*chart.File, c *chart.Chart, name string) []*chart.File {
	if name != c.Name() {
		name = filepath.Join(name, c.Name())
	}
	for _, t := range c.Templates {
		t.Name = filepath.Join(name, t.Name)
	}
	files = append(files, c.Templates...)
	for _, dep := range c.Dependencies() {
		files = updateName(files, dep, filepath.Join(name, "charts"))
	}
	return files
}

// runOnce and splitManifestYAML will be added in Tasks 4 and 5.
// Add placeholder references to suppress unused import errors:
var _ = action.NewInstall
var _ = loader.Load
var _ chartutil.VersionSet
var _ getter.Providers
var _ = values.Options{}
var _ = log.SetOutput
var _ = io.Discard
