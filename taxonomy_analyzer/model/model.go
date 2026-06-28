package model

import "time"

type FixAttempt struct {
	Attempt       int    `json:"attempt"`
	ErrorSeen     string `json:"error_seen"`
	Kind          string `json:"kind"`
	ValuePath     string `json:"value_path"`
	ValueInjected string `json:"value_injected"`
}

type FixedResult struct {
	Resolved   bool         `json:"resolved"`
	StopReason string       `json:"stop_reason"`
	FixChain   []FixAttempt `json:"fix_chain"`
	// FinalError is the error the chart was still failing with when the fixer
	// gave up — the real blocker, which can differ from the run's original
	// error_message once injections have changed the failure.
	FinalError string `json:"final_error,omitempty"`
}

type FixOutcome struct {
	Attempted int `json:"attempted"`
	Resolved  int `json:"resolved"`
	// NonReproduced counts runs the fixer marked resolved WITHOUT applying any
	// fix (empty fix chain): the original failure simply did not recur when the
	// chart was re-rendered (e.g. a floating dependency version or helm-version
	// drift between the original run and the fixer run). These are not fixes and
	// are kept out of Resolved so a non-fixable bucket can never show a fix.
	NonReproduced int            `json:"non_reproduced,omitempty"`
	Unresolved    int            `json:"unresolved"`
	ByStopReason  map[string]int `json:"by_stop_reason,omitempty"`
}

type RunResult struct {
	ChartPath    string       `json:"chart_path"`
	ValuesFiles  []string     `json:"values_files,omitempty"`
	HelmCommand  string       `json:"helm_command"`
	Success      bool         `json:"success"`
	ErrorMessage string       `json:"error_message,omitempty"`
	Fixed        *FixedResult `json:"fixed,omitempty"`
}

type ChartSummary struct {
	ChartPath       string      `json:"chart_path"`
	TotalRuns       int         `json:"total_runs"`
	Successes       int         `json:"successes"`
	Failures        int         `json:"failures"`
	DepBuildFailure bool        `json:"dep_build_failure,omitempty"`
	DepBuildError   string      `json:"dep_build_error,omitempty"`
	Runs            []RunResult `json:"runs"`
}

type RepoResult struct {
	RepoURL          string         `json:"repo_url"`
	RepoName         string         `json:"repo_name"`
	ClonedDir        string         `json:"cloned_dir"`
	TotalCharts      int            `json:"total_charts"`
	TotalRuns        int            `json:"total_runs"`
	TotalSuccess     int            `json:"total_successes"`
	TotalFailures    int            `json:"total_failures"`
	TotalDepFailures int            `json:"total_dep_failures"`
	Charts           []ChartSummary `json:"charts"`
	Kept             bool           `json:"kept"`
	DepFailed        bool           `json:"dep_failed"`
}

type ErrorOccurrence struct {
	RepoURL     string   `json:"repo_url"`
	RepoName    string   `json:"repo_name"`
	ChartPath   string   `json:"chart_path"`
	ValuesFiles []string `json:"values_files,omitempty"`
	HelmCommand string   `json:"helm_command,omitempty"`
	ErrorSource string   `json:"error_source"`
	// ErrorKind/ErrorSubKind classify the START error (the run's original
	// error_message) — what the fixer first set out to fix.
	ErrorKind    string `json:"error_kind"`
	ErrorSubKind string `json:"error_sub_kind"`
	ErrorMessage string `json:"error_message"`
	// FinalErrorKind/FinalErrorSubKind classify the FINAL error (the blocker the
	// fixer gave up on). After value injection the failure often shifts to a
	// different taxonomy than the start, so it is classified independently. Set
	// only for unresolved fixes that recorded a final_error.
	FinalErrorKind    string       `json:"final_error_kind,omitempty"`
	FinalErrorSubKind string       `json:"final_error_sub_kind,omitempty"`
	Fixed             *FixedResult `json:"fixed,omitempty"`
}

// TransitionStat counts unresolved fixes whose START taxonomy "From" shifted to
// FINAL taxonomy "To" once injections changed the failure (To may equal From
// when the kind is unchanged, or be unknown.unclassified for a new blocker).
type TransitionStat struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Count int    `json:"count"`
}

type TaxonomyBucket struct {
	Count int `json:"count"`
	// FinalCount is how many unresolved fixes ended on this taxonomy as their
	// FINAL blocker (regardless of what they started as). A bucket can have a
	// high FinalCount but low Count when it is mostly a destination, not a start.
	FinalCount int               `json:"final_count,omitempty"`
	Examples   []ErrorOccurrence `json:"examples"`
	FixOutcome *FixOutcome       `json:"fix_outcome,omitempty"`
	// NonFixable marks a subkind the fixer cannot resolve by value injection
	// (chart-author validation or structural blockers) — see classifier.IsNonFixable.
	NonFixable bool `json:"non_fixable,omitempty"`
}

type ReportTotals struct {
	Repos             int `json:"repos"`
	Runs              int `json:"runs"`
	TemplateFailures  int `json:"template_failures"`
	DependencyErrors  int `json:"dependency_failures"`
	ClassifiedErrors  int `json:"classified_errors"`
	UnclassifiedError int `json:"unclassified_errors"`
	// DuplicatesCollapsed counts occurrences dropped because the same error
	// message already appeared for the same project (value-file combinations and
	// repeated charts). Raw run counts above (Runs, TemplateFailures) are NOT
	// deduped; the classified/fix totals and taxonomy buckets ARE.
	DuplicatesCollapsed int `json:"duplicates_collapsed,omitempty"`
	// NonFixableErrors is how many classified occurrences fall in a non-fixable
	// subkind (classifier.IsNonFixable) — guardrails/structural blockers excluded
	// from the fixable surface.
	NonFixableErrors int `json:"non_fixable_errors,omitempty"`
	FixAttempted     int `json:"fix_attempted,omitempty"`
	FixResolved      int `json:"fix_resolved,omitempty"`
	// FixNonReproduced counts resolved-with-empty-chain runs (see
	// FixOutcome.NonReproduced): the fixer applied nothing and the failure just
	// did not recur on re-render. Excluded from FixResolved.
	FixNonReproduced int `json:"fix_non_reproduced,omitempty"`
	FixUnresolved    int `json:"fix_unresolved,omitempty"`
}

type TaxonomyReport struct {
	GeneratedAt   time.Time                 `json:"generated_at"`
	SourceCatalog string                    `json:"source_catalog"`
	FixedCatalog  string                    `json:"fixed_catalog,omitempty"`
	Totals        ReportTotals              `json:"totals"`
	ByKind        map[string]TaxonomyBucket `json:"by_kind"`
	BySubKind     map[string]TaxonomyBucket `json:"by_sub_kind"`
	ByRepo        map[string]map[string]int `json:"by_repo"`
	// Transitions records how the START taxonomy shifted to the FINAL taxonomy
	// across unresolved fixes (sorted by count desc).
	Transitions   []TransitionStat  `json:"transitions,omitempty"`
	Unclassified  []ErrorOccurrence `json:"unclassified"`
	AllClassified []ErrorOccurrence `json:"all_classified"`
}
