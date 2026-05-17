package analyzer_test

import (
	"testing"

	"github.com/MabsIPCA/helm-tests/taxonomy_analyzer/analyzer"
	"github.com/MabsIPCA/helm-tests/taxonomy_analyzer/model"
)

func repoWithTwoNilPointerRuns() model.RepoResult {
	return model.RepoResult{
		RepoURL:       "https://github.com/test/repo-a",
		RepoName:      "test/repo-a",
		TotalRuns:     2,
		TotalFailures: 2,
		Charts: []model.ChartSummary{
			{
				ChartPath: "/tmp/repo-a/chart",
				TotalRuns: 2,
				Failures:  2,
				Runs: []model.RunResult{
					{
						ChartPath:    "cloned/repo-a/chart",
						HelmCommand:  "helm template test cloned/repo-a/chart",
						Success:      false,
						ErrorMessage: "nil pointer evaluating interface {}.type",
					},
					{
						ChartPath:    "cloned/repo-a/chart",
						HelmCommand:  "helm template test cloned/repo-a/chart -f values.yaml",
						Success:      false,
						ErrorMessage: "nil pointer evaluating interface {}.repository",
					},
				},
			},
		},
	}
}

func TestConsumeRepo_WithFixedIndex_AccumulatesFixStats(t *testing.T) {
	fixedIndex := map[string]*model.FixedResult{
		"helm template test cloned/repo-a/chart": {
			Resolved:   true,
			StopReason: "",
			FixChain:   []model.FixAttempt{{Attempt: 1, Kind: "nil_pointer", ValuePath: "service.type", ValueInjected: "ClusterIP"}},
		},
		"helm template test cloned/repo-a/chart -f values.yaml": {
			Resolved:   false,
			StopReason: "unfixable_error_kind",
			FixChain:   []model.FixAttempt{},
		},
	}

	a := analyzer.New("catalog.json", "catalog_fixed.json", fixedIndex)
	a.ConsumeRepo(repoWithTwoNilPointerRuns())
	report := a.Report()

	if report.FixedCatalog != "catalog_fixed.json" {
		t.Errorf("FixedCatalog: got %q, want %q", report.FixedCatalog, "catalog_fixed.json")
	}
	if report.Totals.FixAttempted != 2 {
		t.Errorf("FixAttempted: got %d, want 2", report.Totals.FixAttempted)
	}
	if report.Totals.FixResolved != 1 {
		t.Errorf("FixResolved: got %d, want 1", report.Totals.FixResolved)
	}
	if report.Totals.FixUnresolved != 1 {
		t.Errorf("FixUnresolved: got %d, want 1", report.Totals.FixUnresolved)
	}

	subBucket := report.BySubKind["template.nil_pointer"]
	if subBucket.FixOutcome.Attempted != 2 {
		t.Errorf("nil_pointer FixOutcome.Attempted: got %d, want 2", subBucket.FixOutcome.Attempted)
	}
	if subBucket.FixOutcome.Resolved != 1 {
		t.Errorf("nil_pointer FixOutcome.Resolved: got %d, want 1", subBucket.FixOutcome.Resolved)
	}
	if subBucket.FixOutcome.Unresolved != 1 {
		t.Errorf("nil_pointer FixOutcome.Unresolved: got %d, want 1", subBucket.FixOutcome.Unresolved)
	}
	if subBucket.FixOutcome.ByStopReason["unfixable_error_kind"] != 1 {
		t.Errorf("nil_pointer ByStopReason[unfixable_error_kind]: got %d, want 1", subBucket.FixOutcome.ByStopReason["unfixable_error_kind"])
	}

	kindBucket := report.ByKind["template"]
	if kindBucket.FixOutcome.Attempted != 2 {
		t.Errorf("template FixOutcome.Attempted: got %d, want 2", kindBucket.FixOutcome.Attempted)
	}
}

func TestConsumeRepo_WithoutFixedIndex_NoFixStats(t *testing.T) {
	a := analyzer.New("catalog.json", "", nil)
	a.ConsumeRepo(repoWithTwoNilPointerRuns())
	report := a.Report()

	if report.Totals.FixAttempted != 0 {
		t.Errorf("FixAttempted: got %d, want 0", report.Totals.FixAttempted)
	}
	if report.BySubKind["template.nil_pointer"].FixOutcome.Attempted != 0 {
		t.Errorf("bucket FixOutcome.Attempted: got %d, want 0", report.BySubKind["template.nil_pointer"].FixOutcome.Attempted)
	}
}

func TestConsumeRepo_RunNotInFixedIndex_NoFixStats(t *testing.T) {
	fixedIndex := map[string]*model.FixedResult{} // empty — no matching entry

	a := analyzer.New("catalog.json", "catalog_fixed.json", fixedIndex)
	a.ConsumeRepo(repoWithTwoNilPointerRuns())
	report := a.Report()

	if report.Totals.FixAttempted != 0 {
		t.Errorf("FixAttempted: got %d, want 0", report.Totals.FixAttempted)
	}
}
