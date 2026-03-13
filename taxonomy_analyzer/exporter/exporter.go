package exporter

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/MabsIPCA/helm-tests/taxonomy_analyzer/model"
)

// WriteAll emits JSON, markdown and CSV outputs in outputDir.
func WriteAll(outputDir string, report model.TaxonomyReport) error {
	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		return fmt.Errorf("create output dir: %w", err)
	}

	if err := writeJSON(filepath.Join(outputDir, "taxonomy_report.json"), report); err != nil {
		return err
	}
	if err := writeMarkdown(filepath.Join(outputDir, "taxonomy_summary.md"), report); err != nil {
		return err
	}
	if err := writeCSV(filepath.Join(outputDir, "taxonomy_occurrences.csv"), report.AllClassified); err != nil {
		return err
	}
	if err := writeErrorsByBucketMarkdown(filepath.Join(outputDir, "taxonomy_errors_by_bucket.md"), report); err != nil {
		return err
	}

	return nil
}

func writeJSON(path string, report model.TaxonomyReport) error {
	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal report json: %w", err)
	}
	if err := os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("write report json: %w", err)
	}
	return nil
}

func writeMarkdown(path string, report model.TaxonomyReport) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create markdown report: %w", err)
	}
	defer f.Close()

	fmt.Fprintln(f, "# Helm Error Taxonomy Report")
	fmt.Fprintln(f)
	fmt.Fprintf(f, "Generated at: `%s`\n\n", report.GeneratedAt.Format("2006-01-02 15:04:05 MST"))
	fmt.Fprintf(f, "Source catalog: `%s`\n\n", report.SourceCatalog)

	fmt.Fprintln(f, "## Totals")
	fmt.Fprintln(f)
	fmt.Fprintln(f, "| Metric | Value |")
	fmt.Fprintln(f, "|---|---:|")
	fmt.Fprintf(f, "| Repositories | %d |\n", report.Totals.Repos)
	fmt.Fprintf(f, "| Helm runs | %d |\n", report.Totals.Runs)
	fmt.Fprintf(f, "| Template failures | %d |\n", report.Totals.TemplateFailures)
	fmt.Fprintf(f, "| Dependency failures | %d |\n", report.Totals.DependencyErrors)
	fmt.Fprintf(f, "| Classified errors | %d |\n", report.Totals.ClassifiedErrors)
	fmt.Fprintf(f, "| Unclassified errors | %d |\n", report.Totals.UnclassifiedError)
	fmt.Fprintln(f)

	fmt.Fprintln(f, "## Taxonomy by Kind")
	fmt.Fprintln(f)
	fmt.Fprintln(f, "| Kind | Count |")
	fmt.Fprintln(f, "|---|---:|")
	for _, key := range sortedKeysByCount(report.ByKind) {
		fmt.Fprintf(f, "| `%s` | %d |\n", key, report.ByKind[key].Count)
	}
	fmt.Fprintln(f)

	fmt.Fprintln(f, "## Taxonomy by SubKind")
	fmt.Fprintln(f)
	fmt.Fprintln(f, "| SubKind | Count |")
	fmt.Fprintln(f, "|---|---:|")
	for _, key := range sortedKeysByCount(report.BySubKind) {
		fmt.Fprintf(f, "| `%s` | %d |\n", key, report.BySubKind[key].Count)
	}
	fmt.Fprintln(f)

	if len(report.Unclassified) > 0 {
		fmt.Fprintln(f, "## Unclassified Samples")
		fmt.Fprintln(f)
		limit := 10
		if len(report.Unclassified) < limit {
			limit = len(report.Unclassified)
		}
		for i := 0; i < limit; i++ {
			u := report.Unclassified[i]
			msg := strings.TrimSpace(u.ErrorMessage)
			msg = strings.ReplaceAll(msg, "\n", " ")
			if len(msg) > 180 {
				msg = msg[:180] + "..."
			}
			fmt.Fprintf(f, "- `%s` `%s` `%s`: %s\n", u.RepoName, u.ChartPath, u.ErrorSource, msg)
		}
	}

	return nil
}

func writeCSV(path string, occurrences []model.ErrorOccurrence) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create csv report: %w", err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	header := []string{"repo_url", "repo_name", "chart_path", "values_files", "helm_command", "error_source", "error_kind", "error_sub_kind", "error_message"}
	if err := w.Write(header); err != nil {
		return fmt.Errorf("write csv header: %w", err)
	}

	for _, occ := range occurrences {
		if err := w.Write([]string{
			occ.RepoURL,
			occ.RepoName,
			occ.ChartPath,
			strings.Join(occ.ValuesFiles, " | "),
			occ.HelmCommand,
			occ.ErrorSource,
			occ.ErrorKind,
			occ.ErrorSubKind,
			occ.ErrorMessage,
		}); err != nil {
			return fmt.Errorf("write csv row: %w", err)
		}
	}

	return nil
}

func sortedKeysByCount(buckets map[string]model.TaxonomyBucket) []string {
	keys := make([]string, 0, len(buckets))
	for key := range buckets {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		if buckets[keys[i]].Count == buckets[keys[j]].Count {
			return keys[i] < keys[j]
		}
		return buckets[keys[i]].Count > buckets[keys[j]].Count
	})
	return keys
}

func writeErrorsByBucketMarkdown(path string, report model.TaxonomyReport) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create bucket markdown report: %w", err)
	}
	defer f.Close()

	buckets := make(map[string][]model.ErrorOccurrence)
	for _, occ := range report.AllClassified {
		bucket := fmt.Sprintf("%s.%s", occ.ErrorKind, occ.ErrorSubKind)
		buckets[bucket] = append(buckets[bucket], occ)
	}

	fmt.Fprintln(f, "# Complete Errors By Taxonomy Bucket")
	fmt.Fprintln(f)
	fmt.Fprintf(f, "Generated at: `%s`\n\n", report.GeneratedAt.Format("2006-01-02 15:04:05 MST"))
	fmt.Fprintf(f, "Source catalog: `%s`\n\n", report.SourceCatalog)

	for _, key := range sortedOccurrenceBucketKeys(buckets) {
		entries := buckets[key]
		fmt.Fprintf(f, "## `%s` (%d)\n\n", key, len(entries))
		for i, occ := range entries {
			fmt.Fprintf(f, "### %d. `%s`\n\n", i+1, occ.RepoName)
			fmt.Fprintf(f, "- Chart: `%s`\n", occ.ChartPath)
			fmt.Fprintf(f, "- Source: `%s`\n", occ.ErrorSource)
			if len(occ.ValuesFiles) > 0 {
				fmt.Fprintf(f, "- Values files: `%s`\n", strings.Join(occ.ValuesFiles, "`, `"))
			}
			if occ.HelmCommand != "" {
				fmt.Fprintf(f, "- Command: `%s`\n", occ.HelmCommand)
			}
			fmt.Fprintln(f)
			fmt.Fprintln(f, "```text")
			fmt.Fprintln(f, strings.TrimSpace(occ.ErrorMessage))
			fmt.Fprintln(f, "```")
			fmt.Fprintln(f)
		}
	}

	return nil
}

func sortedOccurrenceBucketKeys(buckets map[string][]model.ErrorOccurrence) []string {
	keys := make([]string, 0, len(buckets))
	for k := range buckets {
		keys = append(keys, k)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		if len(buckets[keys[i]]) == len(buckets[keys[j]]) {
			return keys[i] < keys[j]
		}
		return len(buckets[keys[i]]) > len(buckets[keys[j]])
	})
	return keys
}
