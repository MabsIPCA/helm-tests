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

// runOnce runs the KICS-faithful Helm pipeline once for the given chart path
// and values options. When debugMode is true, cfg.Log is wired to collect log
// lines and helmSettings.Debug is set to true for the duration of the call.
//
// Even when err != nil, rel may be non-nil (partial render before failure).
func runOnce(chartPath string, valOpts *values.Options, debugMode bool) renderResult {
	var logs []string

	cfg := new(action.Configuration)
	if debugMode {
		cfg.Log = func(format string, v ...interface{}) {
			logs = append(logs, fmt.Sprintf(format, v...))
		}
		helmSettings.Debug = true
		defer func() { helmSettings.Debug = false }()
	}

	client := action.NewInstall(cfg)
	client.DryRun = true
	client.ReleaseName = "kics-helm"
	client.Replace = true
	client.ClientOnly = true
	client.APIVersions = chartutil.VersionSet([]string{})
	client.IncludeCRDs = false
	client.Namespace = "kics-namespace"

	log.SetOutput(io.Discard)
	defer log.SetOutput(nil)

	p := getter.All(helmSettings)
	vals, err := valOpts.MergeValues(p)
	if err != nil {
		return renderResult{err: err, logs: logs}
	}

	chartReq, err := loader.Load(chartPath)
	if err != nil {
		return renderResult{err: err, logs: logs}
	}

	chartReq = setID(chartReq)

	rel, runErr := client.Run(chartReq, vals)
	return renderResult{rel: rel, err: runErr, logs: logs}
}

// splitManifestYAML mirrors KICS's splitManifestYAML: splits rel.Manifest by
// "---" separators, maps each section back to its source template file, and
// returns structured entries. originalContent is the post-setID template source.
func splitManifestYAML(rel *release.Release) []SplitManifestEntry {
	// Build path→originalContent map using the same updateName logic as KICS.
	sources := updateName(nil, rel.Chart, rel.Chart.Name())
	origData := make(map[string][]byte, len(sources))
	for _, f := range sources {
		key := filepath.ToSlash(f.Name) // manifest uses forward slashes
		origData[key] = []byte(strings.ReplaceAll(string(f.Data), "\r", ""))
	}

	var entries []SplitManifestEntry
	for _, section := range strings.Split(rel.Manifest, "---") {
		section = strings.ReplaceAll(section, "\r", "")
		if strings.TrimSpace(section) == "" {
			continue
		}

		lines := strings.Split(section, "\n")

		// Extract # Source: path
		path := ""
		for _, l := range lines {
			if strings.HasPrefix(l, "# Source: ") {
				path = strings.TrimSpace(strings.TrimPrefix(l, "# Source: "))
				break
			}
		}
		if path == "" {
			continue
		}

		orig, ok := origData[path]
		if !ok {
			continue
		}

		// Extract KICS split ID if present
		splitID := ""
		for _, l := range lines {
			if strings.Contains(l, kicsHelmID) {
				splitID = strings.TrimSpace(l)
				break
			}
		}

		entries = append(entries, SplitManifestEntry{
			Path:            path,
			Content:         section,
			OriginalContent: string(orig),
			SplitID:         splitID,
		})
	}
	return entries
}
