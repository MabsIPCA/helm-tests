package main

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"

	"github.com/MabsIPCA/helm-tests/helm_fetcher/exporter"
	"github.com/MabsIPCA/helm-tests/helm_fetcher/model"
)

// runRenderMDMode re-renders the markdown + CSV sidecars for an existing
// catalog_by_project.json WITHOUT re-running helm. This is needed after a mode
// that rewrites the JSON in place (e.g. refilter-values) but leaves the
// catalog_results.md / *.csv siblings stale. It reuses the exporter writers, so
// the output is byte-identical to what a full run would have produced. The
// authoritative catalog_by_project.json itself is left untouched.
func runRenderMDMode(catalogPath string) {
	if catalogPath == "" {
		log.Fatal().Msg("render-md mode requires -catalog-in (path to catalog_by_project.json, or a run dir)")
	}
	// Accept either a run directory or a direct path to the JSON file.
	if info, err := os.Stat(catalogPath); err == nil && info.IsDir() {
		catalogPath = filepath.Join(catalogPath, "catalog_by_project.json")
	}
	dir := filepath.Dir(catalogPath)

	repos := loadRepoResults(catalogPath)
	if repos == nil {
		log.Fatal().Str("file", catalogPath).Msg("Could not load catalog_by_project.json")
	}
	deps := loadDepFailureEntries(filepath.Join(dir, "catalog_dep_failures.json"))

	p := func(name string) string { return filepath.Join(dir, name) }
	_ = exporter.WriteMarkdownTable(p("catalog_results.md"), repos)
	_ = exporter.WriteMarkdownTable(p("catalog_kept.md"), filterKept(repos))
	_ = exporter.WriteMarkdownTable(p("catalog_removed.md"), filterRemoved(repos))
	_ = exporter.WriteDepFailuresMarkdown(p("catalog_dep_failures.md"), deps)
	_ = exporter.WriteFailuresCSV(p("catalog_failures.csv"), repos)
	_ = exporter.WriteFailuresCSV(p("catalog_kept_failures.csv"), filterKept(repos))
	_ = exporter.WriteDepFailuresCSV(p("catalog_dep_failures.csv"), deps)

	log.Info().
		Str("dir", dir).
		Int("repos", len(repos)).
		Int("dep_failures", len(deps)).
		Msg("Re-rendered markdown + CSV sidecars from catalog_by_project.json")
}

func filterKept(repos []model.RepoResult) []model.RepoResult {
	var out []model.RepoResult
	for _, r := range repos {
		if r.Kept {
			out = append(out, r)
		}
	}
	return out
}

func filterRemoved(repos []model.RepoResult) []model.RepoResult {
	var out []model.RepoResult
	for _, r := range repos {
		if !r.Kept && !r.DepFailed {
			out = append(out, r)
		}
	}
	return out
}

func loadRepoResults(path string) []model.RepoResult {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	var repos []model.RepoResult
	if err := json.Unmarshal(data, &repos); err != nil {
		log.Warn().Err(err).Str("file", path).Msg("Could not parse catalog_by_project.json")
		return nil
	}
	return repos
}

func loadDepFailureEntries(path string) []model.DepFailureEntry {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	var entries []model.DepFailureEntry
	if err := json.Unmarshal(data, &entries); err != nil {
		log.Warn().Err(err).Str("file", path).Msg("Could not parse catalog_dep_failures.json")
		return nil
	}
	return entries
}
