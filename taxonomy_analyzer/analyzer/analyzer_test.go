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
	if subBucket.FixOutcome == nil {
		t.Fatal("expected non-nil FixOutcome")
	}
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
	if kindBucket.FixOutcome == nil {
		t.Fatal("expected non-nil FixOutcome for kind bucket")
	}
	if kindBucket.FixOutcome.Attempted != 2 {
		t.Errorf("template FixOutcome.Attempted: got %d, want 2", kindBucket.FixOutcome.Attempted)
	}
}

func TestConsumeRepo_FinalErrorTaxonomyTransition(t *testing.T) {
	// Run starts as a required_value error; after injecting placeholders the
	// fixer gives up on a missing-template blocker — a different taxonomy. The
	// start classification, final classification, and the transition must all
	// be recorded.
	fixedIndex := map[string]*model.FixedResult{
		"helm template test cloned/repo-a/chart": {
			Resolved:   false,
			StopReason: "max_iterations",
			// Injection happened (non-empty chain), so the differing FinalError is
			// a genuine post-injection shift, not a fixer-env artifact.
			FixChain:   []model.FixAttempt{{Attempt: 1, Kind: "required_value", ValuePath: "secrets.KEY.value", ValueInjected: "kics-placeholder"}},
			FinalError: `no template "docker.registry" associated with template "gotpl"`,
		},
		"helm template test cloned/repo-a/chart -f values.yaml": {
			Resolved: true,
		},
	}

	repo := repoWithTwoNilPointerRuns()
	repo.Charts[0].Runs[0].ErrorMessage = "execution error: required value for secrets.KEY either as .value or .valueFrom"

	a := analyzer.New("catalog.json", "catalog_fixed.json", fixedIndex)
	a.ConsumeRepo(repo)
	report := a.Report()

	var occ *model.ErrorOccurrence
	for i := range report.AllClassified {
		if report.AllClassified[i].HelmCommand == "helm template test cloned/repo-a/chart" {
			occ = &report.AllClassified[i]
			break
		}
	}
	if occ == nil {
		t.Fatal("occurrence for unresolved run not found")
	}
	if occ.ErrorSubKind != "required_value" {
		t.Errorf("start sub_kind: got %q, want required_value", occ.ErrorSubKind)
	}
	if occ.FinalErrorSubKind != "missing_template" {
		t.Errorf("final sub_kind: got %q, want missing_template", occ.FinalErrorSubKind)
	}

	if len(report.Transitions) != 1 {
		t.Fatalf("transitions: got %d, want 1 (%+v)", len(report.Transitions), report.Transitions)
	}
	got := report.Transitions[0]
	if got.From != "template.required_value" || got.To != "template.missing_template" || got.Count != 1 {
		t.Errorf("transition: got %+v, want {template.required_value -> template.missing_template 1}", got)
	}
}

func TestConsumeRepo_UnfixableEmptyChain_FinalEqualsOriginal(t *testing.T) {
	// values_schema_validation is never attempted: the fixer stops immediately
	// with an EMPTY chain. The FinalError it captured is a fixer-environment
	// artifact (deps not built -> a dependency error). Because nothing was
	// injected, the final must equal the ORIGINAL (values_schema_validation),
	// and no transition is recorded.
	fixedIndex := map[string]*model.FixedResult{
		"helm template test cloned/repo-a/chart": {
			Resolved:   false,
			StopReason: "unfixable_error_kind",
			FixChain:   nil, // nothing injected
			FinalError: "an error occurred while checking for chart dependencies",
		},
	}

	repo := repoWithTwoNilPointerRuns()
	repo.Charts[0].Runs = repo.Charts[0].Runs[:1] // single run
	repo.Charts[0].Runs[0].ErrorMessage = "values don't meet the specifications of the schema(s) in the following chart(s):\nfoo"

	a := analyzer.New("catalog.json", "catalog_fixed.json", fixedIndex)
	a.ConsumeRepo(repo)
	report := a.Report()

	occ := report.AllClassified[0]
	if occ.ErrorSubKind != "values_schema_validation" {
		t.Fatalf("start sub_kind: got %q, want values_schema_validation", occ.ErrorSubKind)
	}
	if occ.FinalErrorSubKind != "values_schema_validation" {
		t.Errorf("final sub_kind: got %q, want values_schema_validation (final must equal original)", occ.FinalErrorSubKind)
	}
	if got := report.BySubKind["template.values_schema_validation"].FinalCount; got != 1 {
		t.Errorf("values_schema_validation FinalCount: got %d, want 1", got)
	}
	if got := report.BySubKind["template.dependency_check_failed"].FinalCount; got != 0 {
		t.Errorf("dependency_check_failed FinalCount: got %d, want 0 (artifact must not leak)", got)
	}
	if len(report.Transitions) != 0 {
		t.Errorf("transitions: got %d, want 0 (no injection)", len(report.Transitions))
	}
}

func TestConsumeRepo_FinalErrorIsCumulative_NotAttempted(t *testing.T) {
	// A run that was never attempted (not in the fixed index) is still broken,
	// so its original error stands as the FINAL error: final_count is bumped and
	// the final taxonomy equals the start. No transition is recorded (nothing was
	// injected to shift the taxonomy).
	fixedIndex := map[string]*model.FixedResult{
		// only the first run is attempted+resolved; the second is not attempted
		"helm template test cloned/repo-a/chart": {Resolved: true},
	}

	a := analyzer.New("catalog.json", "catalog_fixed.json", fixedIndex)
	a.ConsumeRepo(repoWithTwoNilPointerRuns())
	report := a.Report()

	var attempted, notAttempted *model.ErrorOccurrence
	for i := range report.AllClassified {
		switch report.AllClassified[i].HelmCommand {
		case "helm template test cloned/repo-a/chart":
			attempted = &report.AllClassified[i]
		case "helm template test cloned/repo-a/chart -f values.yaml":
			notAttempted = &report.AllClassified[i]
		}
	}
	if attempted == nil || notAttempted == nil {
		t.Fatal("expected both occurrences present")
	}
	// Resolved run: no final error.
	if attempted.FinalErrorKind != "" {
		t.Errorf("resolved run final kind: got %q, want empty", attempted.FinalErrorKind)
	}
	// Not-attempted run: final == start.
	if notAttempted.FinalErrorSubKind != "nil_pointer" {
		t.Errorf("not-attempted final sub_kind: got %q, want nil_pointer", notAttempted.FinalErrorSubKind)
	}
	// final_count counts the not-attempted run only (resolved one excluded).
	if got := report.BySubKind["template.nil_pointer"].FinalCount; got != 1 {
		t.Errorf("template.nil_pointer FinalCount: got %d, want 1", got)
	}
	// No transitions: the only non-resolved run was never attempted.
	if len(report.Transitions) != 0 {
		t.Errorf("transitions: got %d, want 0 (%+v)", len(report.Transitions), report.Transitions)
	}
}

func TestConsumeRepo_DedupesSameErrorWithinProject(t *testing.T) {
	// The fetcher renders many value-file combinations of one chart; identical
	// errors must collapse to a single occurrence within the project. Here three
	// runs share the same malformed-yaml error and one run has a distinct error.
	sameErr := "error converting YAML to JSON: yaml: control characters are not allowed"
	repo := model.RepoResult{
		RepoURL:  "https://github.com/test/mono",
		RepoName: "test/mono",
		Charts: []model.ChartSummary{{
			ChartPath: "/tmp/mono/chart",
			Runs: []model.RunResult{
				{ChartPath: "cloned/mono/chart", HelmCommand: "helm template m chart -f a.yaml", ErrorMessage: sameErr},
				{ChartPath: "cloned/mono/chart", HelmCommand: "helm template m chart -f b.yaml", ErrorMessage: sameErr},
				{ChartPath: "cloned/mono/chart", HelmCommand: "helm template m chart -f a.yaml,b.yaml", ErrorMessage: sameErr},
				{ChartPath: "cloned/mono/chart", HelmCommand: "helm template m chart -f c.yaml", ErrorMessage: "nil pointer evaluating interface {}.type"},
			},
		}},
	}

	a := analyzer.New("catalog.json", "", nil)
	a.ConsumeRepo(repo)
	report := a.Report()

	if got := report.BySubKind["template.malformed_yaml"].Count; got != 1 {
		t.Errorf("malformed_yaml Count: got %d, want 1 (combos collapsed)", got)
	}
	if got := report.Totals.DuplicatesCollapsed; got != 2 {
		t.Errorf("DuplicatesCollapsed: got %d, want 2", got)
	}
	// The distinct error is preserved.
	if got := report.BySubKind["template.nil_pointer"].Count; got != 1 {
		t.Errorf("nil_pointer Count: got %d, want 1 (distinct error preserved)", got)
	}
	if got := len(report.AllClassified); got != 2 {
		t.Errorf("AllClassified: got %d, want 2", got)
	}
}

func TestConsumeRepo_SameErrorDifferentProjects_NotDeduped(t *testing.T) {
	// An identical message in two different projects must count once per project.
	sameErr := "error converting YAML to JSON: yaml: control characters are not allowed"
	mk := func(name string) model.RepoResult {
		return model.RepoResult{
			RepoURL:  "https://github.com/test/" + name,
			RepoName: "test/" + name,
			Charts: []model.ChartSummary{{
				ChartPath: "/tmp/" + name + "/chart",
				Runs:      []model.RunResult{{ChartPath: "cloned/" + name + "/chart", HelmCommand: "helm template " + name, ErrorMessage: sameErr}},
			}},
		}
	}

	a := analyzer.New("catalog.json", "", nil)
	a.ConsumeRepo(mk("alpha"))
	a.ConsumeRepo(mk("beta"))
	report := a.Report()

	if got := report.BySubKind["template.malformed_yaml"].Count; got != 2 {
		t.Errorf("malformed_yaml Count: got %d, want 2 (distinct projects)", got)
	}
	if got := report.Totals.DuplicatesCollapsed; got != 0 {
		t.Errorf("DuplicatesCollapsed: got %d, want 0", got)
	}
}

func TestConsumeRepo_WithoutFixedIndex_NoFixStats(t *testing.T) {
	a := analyzer.New("catalog.json", "", nil)
	a.ConsumeRepo(repoWithTwoNilPointerRuns())
	report := a.Report()

	if report.Totals.FixAttempted != 0 {
		t.Errorf("FixAttempted: got %d, want 0", report.Totals.FixAttempted)
	}
	if report.BySubKind["template.nil_pointer"].FixOutcome != nil {
		t.Errorf("expected nil FixOutcome when no fix index")
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
