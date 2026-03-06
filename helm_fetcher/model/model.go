package model

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
	ChartPath string      `json:"chart_path"`
	TotalRuns int         `json:"total_runs"`
	Successes int         `json:"successes"`
	Failures  int         `json:"failures"`
	Runs      []RunResult `json:"runs"`
}

// RepoResult holds all results for a single repository.
type RepoResult struct {
	RepoURL       string         `json:"repo_url"`
	RepoName      string         `json:"repo_name"`
	ClonedDir     string         `json:"cloned_dir"`
	TotalCharts   int            `json:"total_charts"`
	TotalRuns     int            `json:"total_runs"`
	TotalSuccess  int            `json:"total_successes"`
	TotalFailures int            `json:"total_failures"`
	Charts        []ChartSummary `json:"charts"`
	Kept          bool           `json:"kept"` // true if repo kept due to failures
}

// GitHubSearchResult models the JSON returned by the GitHub code-search API.
type GitHubSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []struct {
		Name       string `json:"name"`
		Path       string `json:"path"`
		Repository struct {
			FullName string `json:"full_name"`
			HTMLURL  string `json:"html_url"`
		} `json:"repository"`
	} `json:"items"`
}
