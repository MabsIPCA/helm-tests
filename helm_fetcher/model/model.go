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
	ChartPath       string      `json:"chart_path"`
	TotalRuns       int         `json:"total_runs"`
	Successes       int         `json:"successes"`
	Failures        int         `json:"failures"`
	DepBuildFailure bool        `json:"dep_build_failure,omitempty"` // true when helm dep build failed
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
	Kept             bool           `json:"kept"`       // true if repo kept due to template failures
	DepFailed        bool           `json:"dep_failed"` // true if any chart had dep build failures
}

// DepFailureEntry records a single helm dependency build failure for separate cataloging.
type DepFailureEntry struct {
	RepoURL   string `json:"repo_url"`
	RepoName  string `json:"repo_name"`
	ChartPath string `json:"chart_path"`
	Error     string `json:"error"`
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

// ArtifactHubSearchResponse models the package search payload from Artifact Hub.
type ArtifactHubSearchResponse struct {
	Packages []ArtifactHubPackage `json:"packages"`
}

// ArtifactHubPackage contains the fields needed to resolve a source repository.
type ArtifactHubPackage struct {
	Name       string             `json:"name"`
	Version    string             `json:"version"`
	Repository ArtifactRepository `json:"repository"`
	Links      []ArtifactLink     `json:"links"`
}

type ArtifactRepository struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type ArtifactLink struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
