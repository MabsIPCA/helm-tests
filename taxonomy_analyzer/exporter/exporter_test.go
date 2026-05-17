package exporter_test

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/MabsIPCA/helm-tests/taxonomy_analyzer/exporter"
	"github.com/MabsIPCA/helm-tests/taxonomy_analyzer/model"
)

func buildReportWithFix() model.TaxonomyReport {
	occ := model.ErrorOccurrence{
		RepoURL:      "https://github.com/test/repo",
		RepoName:     "test/repo",
		ChartPath:    "chart",
		HelmCommand:  "helm template test chart",
		ErrorSource:  "template",
		ErrorKind:    "template",
		ErrorSubKind: "nil_pointer",
		ErrorMessage: "nil pointer evaluating interface {}.type",
		Fixed: &model.FixedResult{
			Resolved:   false,
			StopReason: "unfixable_error_kind",
		},
	}
	bucket := model.TaxonomyBucket{
		Count:    1,
		Examples: []model.ErrorOccurrence{occ},
		FixOutcome: &model.FixOutcome{
			Attempted:    1,
			Unresolved:   1,
			ByStopReason: map[string]int{"unfixable_error_kind": 1},
		},
	}
	return model.TaxonomyReport{
		GeneratedAt:   time.Now().UTC(),
		SourceCatalog: "catalog.json",
		FixedCatalog:  "catalog_fixed.json",
		Totals: model.ReportTotals{
			Repos:            1,
			Runs:             1,
			TemplateFailures: 1,
			ClassifiedErrors: 1,
			FixAttempted:     1,
			FixResolved:      0,
			FixUnresolved:    1,
		},
		ByKind:        map[string]model.TaxonomyBucket{"template": bucket},
		BySubKind:     map[string]model.TaxonomyBucket{"template.nil_pointer": bucket},
		ByRepo:        map[string]map[string]int{"test/repo": {"template.nil_pointer": 1}},
		AllClassified: []model.ErrorOccurrence{occ},
	}
}

func buildReportWithoutFix() model.TaxonomyReport {
	occ := model.ErrorOccurrence{
		RepoURL:      "https://github.com/test/repo",
		RepoName:     "test/repo",
		ChartPath:    "chart",
		HelmCommand:  "helm template test chart",
		ErrorSource:  "template",
		ErrorKind:    "template",
		ErrorSubKind: "nil_pointer",
		ErrorMessage: "nil pointer evaluating interface {}.type",
	}
	return model.TaxonomyReport{
		GeneratedAt:   time.Now().UTC(),
		SourceCatalog: "catalog.json",
		Totals: model.ReportTotals{
			Repos: 1, Runs: 1, TemplateFailures: 1, ClassifiedErrors: 1,
		},
		ByKind:        map[string]model.TaxonomyBucket{"template": {Count: 1, Examples: []model.ErrorOccurrence{occ}}},
		BySubKind:     map[string]model.TaxonomyBucket{"template.nil_pointer": {Count: 1, Examples: []model.ErrorOccurrence{occ}}},
		ByRepo:        map[string]map[string]int{"test/repo": {"template.nil_pointer": 1}},
		AllClassified: []model.ErrorOccurrence{occ},
	}
}

func TestWriteMarkdown_IncludesFixStats_WhenFixDataPresent(t *testing.T) {
	dir := t.TempDir()
	if err := exporter.WriteAll(dir, buildReportWithFix()); err != nil {
		t.Fatalf("WriteAll: %v", err)
	}

	data, err := os.ReadFile(filepath.Join(dir, "taxonomy_summary.md"))
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	md := string(data)

	for _, want := range []string{"Fix attempts", "Fix resolved", "Fix unresolved", "Fix Resolved", "Fix Unresolved"} {
		if !strings.Contains(md, want) {
			t.Errorf("markdown missing %q", want)
		}
	}
}

func TestWriteMarkdown_NoFixStats_WhenNoFixData(t *testing.T) {
	dir := t.TempDir()
	if err := exporter.WriteAll(dir, buildReportWithoutFix()); err != nil {
		t.Fatalf("WriteAll: %v", err)
	}

	data, err := os.ReadFile(filepath.Join(dir, "taxonomy_summary.md"))
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	md := string(data)

	if strings.Contains(md, "Fix attempts") {
		t.Error("markdown should not contain 'Fix attempts' when no fix data")
	}
}

func TestWriteCSV_IncludesFixColumns(t *testing.T) {
	dir := t.TempDir()
	if err := exporter.WriteAll(dir, buildReportWithFix()); err != nil {
		t.Fatalf("WriteAll: %v", err)
	}

	f, err := os.Open(filepath.Join(dir, "taxonomy_occurrences.csv"))
	if err != nil {
		t.Fatalf("open csv: %v", err)
	}
	defer f.Close()

	rows, err := csv.NewReader(f).ReadAll()
	if err != nil {
		t.Fatalf("read csv: %v", err)
	}
	if len(rows) < 2 {
		t.Fatalf("expected header + 1 data row, got %d rows", len(rows))
	}

	header := rows[0]
	if header[len(header)-2] != "fix_resolved" {
		t.Errorf("second-to-last column: got %q, want \"fix_resolved\"", header[len(header)-2])
	}
	if header[len(header)-1] != "fix_stop_reason" {
		t.Errorf("last column: got %q, want \"fix_stop_reason\"", header[len(header)-1])
	}

	row := rows[1]
	if row[len(row)-2] != "false" {
		t.Errorf("fix_resolved value: got %q, want \"false\"", row[len(row)-2])
	}
	if row[len(row)-1] != "unfixable_error_kind" {
		t.Errorf("fix_stop_reason value: got %q, want \"unfixable_error_kind\"", row[len(row)-1])
	}
}

func TestWriteCSV_EmptyFixColumns_WhenNoFixData(t *testing.T) {
	dir := t.TempDir()
	if err := exporter.WriteAll(dir, buildReportWithoutFix()); err != nil {
		t.Fatalf("WriteAll: %v", err)
	}

	f, err := os.Open(filepath.Join(dir, "taxonomy_occurrences.csv"))
	if err != nil {
		t.Fatalf("open csv: %v", err)
	}
	defer f.Close()

	rows, err := csv.NewReader(f).ReadAll()
	if err != nil {
		t.Fatalf("read csv: %v", err)
	}
	if len(rows) < 2 {
		t.Fatalf("expected header + 1 data row, got %d rows", len(rows))
	}

	row := rows[1]
	if row[len(row)-2] != "" {
		t.Errorf("fix_resolved (no fix data): got %q, want \"\"", row[len(row)-2])
	}
	if row[len(row)-1] != "" {
		t.Errorf("fix_stop_reason (no fix data): got %q, want \"\"", row[len(row)-1])
	}
}
