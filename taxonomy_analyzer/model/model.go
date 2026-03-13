package model

import "time"

// RunResult holds a single helm template execution result.
type RunResult struct {
	ChartPath    string   `json:"chart_path"`
	ValuesFiles  []string `json:"values_files,omitempty"`
	HelmCommand  string   `json:"helm_command"`
	Success      bool     `json:"success"`
	ErrorMessage string   `json:"error_message,omitempty"`
}

// ChartSummary groups all runs for a single chart.
type ChartSummary struct {
	ChartPath       string      `json:"chart_path"`
	TotalRuns       int         `json:"total_runs"`
	Successes       int         `json:"successes"`
	Failures        int         `json:"failures"`
	DepBuildFailure bool        `json:"dep_build_failure,omitempty"`
	DepBuildError   string      `json:"dep_build_error,omitempty"`
	Runs            []RunResult `json:"runs"`
}

// RepoResult holds all results for a single repository.
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

// ErrorOccurrence is a single failed run/dependency build enriched with taxonomy fields.
type ErrorOccurrence struct {
	RepoURL      string   `json:"repo_url"`
	RepoName     string   `json:"repo_name"`
	ChartPath    string   `json:"chart_path"`
	ValuesFiles  []string `json:"values_files,omitempty"`
	HelmCommand  string   `json:"helm_command,omitempty"`
	ErrorSource  string   `json:"error_source"` // template | dependency
	ErrorKind    string   `json:"error_kind"`
	ErrorSubKind string   `json:"error_sub_kind"`
	ErrorMessage string   `json:"error_message"`
}

// TaxonomyBucket groups occurrences by classification.
type TaxonomyBucket struct {
	Count    int               `json:"count"`
	Examples []ErrorOccurrence `json:"examples"`
}

// ReportTotals contains high-level scan totals.
type ReportTotals struct {
	Repos             int `json:"repos"`
	Runs              int `json:"runs"`
	TemplateFailures  int `json:"template_failures"`
	DependencyErrors  int `json:"dependency_failures"`
	ClassifiedErrors  int `json:"classified_errors"`
	UnclassifiedError int `json:"unclassified_errors"`
}

// TaxonomyReport is the final analyzer output.
type TaxonomyReport struct {
	GeneratedAt   time.Time                 `json:"generated_at"`
	SourceCatalog string                    `json:"source_catalog"`
	Totals        ReportTotals              `json:"totals"`
	ByKind        map[string]TaxonomyBucket `json:"by_kind"`
	BySubKind     map[string]TaxonomyBucket `json:"by_sub_kind"`
	ByRepo        map[string]map[string]int `json:"by_repo"`
	Unclassified  []ErrorOccurrence         `json:"unclassified"`
	AllClassified []ErrorOccurrence         `json:"all_classified"`
}
