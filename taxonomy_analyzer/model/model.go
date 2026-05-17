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
}

type FixOutcome struct {
	Attempted    int            `json:"attempted"`
	Resolved     int            `json:"resolved"`
	Unresolved   int            `json:"unresolved"`
	ByStopReason map[string]int `json:"by_stop_reason,omitempty"`
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
	RepoURL      string       `json:"repo_url"`
	RepoName     string       `json:"repo_name"`
	ChartPath    string       `json:"chart_path"`
	ValuesFiles  []string     `json:"values_files,omitempty"`
	HelmCommand  string       `json:"helm_command,omitempty"`
	ErrorSource  string       `json:"error_source"`
	ErrorKind    string       `json:"error_kind"`
	ErrorSubKind string       `json:"error_sub_kind"`
	ErrorMessage string       `json:"error_message"`
	Fixed        *FixedResult `json:"fixed,omitempty"`
}

type TaxonomyBucket struct {
	Count      int               `json:"count"`
	Examples   []ErrorOccurrence `json:"examples"`
	FixOutcome *FixOutcome       `json:"fix_outcome,omitempty"`
}

type ReportTotals struct {
	Repos             int `json:"repos"`
	Runs              int `json:"runs"`
	TemplateFailures  int `json:"template_failures"`
	DependencyErrors  int `json:"dependency_failures"`
	ClassifiedErrors  int `json:"classified_errors"`
	UnclassifiedError int `json:"unclassified_errors"`
	FixAttempted      int `json:"fix_attempted,omitempty"`
	FixResolved       int `json:"fix_resolved,omitempty"`
	FixUnresolved     int `json:"fix_unresolved,omitempty"`
}

type TaxonomyReport struct {
	GeneratedAt   time.Time                 `json:"generated_at"`
	SourceCatalog string                    `json:"source_catalog"`
	FixedCatalog  string                    `json:"fixed_catalog,omitempty"`
	Totals        ReportTotals              `json:"totals"`
	ByKind        map[string]TaxonomyBucket `json:"by_kind"`
	BySubKind     map[string]TaxonomyBucket `json:"by_sub_kind"`
	ByRepo        map[string]map[string]int `json:"by_repo"`
	Unclassified  []ErrorOccurrence         `json:"unclassified"`
	AllClassified []ErrorOccurrence         `json:"all_classified"`
}
