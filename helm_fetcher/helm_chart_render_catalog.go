package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

// ChartResult holds the result for a chart and values combination
type ChartResult struct {
	RepoURL      string                 `json:"repo_url"`
	ChartPath    string                 `json:"chart_path"`
	ValuesFile   string                 `json:"values_file"`
	ValuesCombo  map[string]interface{} `json:"values_combo"`
	Success      bool                   `json:"success"`
	ErrorMessage string                 `json:"error_message,omitempty"`
}

func searchGitHubCharts(ctx context.Context, query string) ([]string, error) {
	url := fmt.Sprintf("https://api.github.com/search/code?q=%s+filename:Chart.yaml", query)
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)

	githubToken := os.Getenv("GITHUB_TOKEN")
	if githubToken == "" {
		log.Warn().Msg("GITHUB_TOKEN environment variable not set")
	}
	if githubToken != "" {
		req.Header.Set("Authorization", "Bearer "+githubToken)
		// Add User-Agent header (required by GitHub API)
		req.Header.Set("User-Agent", "helm-fetcher")
	}

	fmt.Printf("Requesting: %s\nHeaders: %v\n", url, req.Header)

	client := &http.Client{Timeout: 30 * 1e9} // 30 seconds
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Warn().Err(err).Msg("Failed to close response body")
		}
	}(resp.Body)
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("GitHub API returned status %d: %s", resp.StatusCode, resp.Status)
	}
	var result struct {
		Items []struct {
			Repository struct {
				HTMLURL string `json:"html_url"`
			} `json:"repository"`
		} `json:"items"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	var repos []string
	for _, item := range result.Items {
		repos = append(repos, item.Repository.HTMLURL)
	}
	return repos, nil
}

func cloneRepo(repoURL string, destDir string) error {
	cmd := exec.Command("git", "clone", repoURL, destDir)
	return cmd.Run()
}

func findCharts(baseDir string) ([]string, error) {
	var charts []string
	err := filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.Name() == "Chart.yaml" {
			charts = append(charts, filepath.Dir(path))
		}
		return nil
	})
	if err != nil {
		log.Warn().Msg("Error walking directory: " + err.Error())
		return nil, err
	}
	return charts, nil
}

func parseValuesYaml(valuesPath string) (map[string]interface{}, error) {
	// Parse values.yaml and return map
	data, err := os.ReadFile(valuesPath)
	if err != nil {
		return nil, err
	}
	var values map[string]interface{}
	if err := json.Unmarshal(data, &values); err != nil {
		return nil, err
	}
	return values, nil
}

func generateValueCombinations(values map[string]interface{}) []map[string]interface{} {
	// Placeholder: generate combinations (for demo, just use original values)
	return []map[string]interface{}{values}
}

func runHelmTemplate(chartPath string, values map[string]interface{}) (string, error) {
	// Write values to temp file
	valuesFile := filepath.Join(chartPath, "temp-values.yaml")
	data, _ := json.Marshal(values)
	_ = os.WriteFile(valuesFile, data, 0644)
	cmd := exec.Command("helm", "template", "test", chartPath, "-f", valuesFile)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), err
	}
	err = os.Remove(valuesFile)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to remove temp values file")
	}

	return string(output), nil
}

func main() {
	_ = godotenv.Load()
	ctx := context.Background()
	query := "helm chart"
	repos, err := searchGitHubCharts(ctx, query)
	if err != nil {
		fmt.Println("GitHub search failed:", err)
		return
	}
	results := []ChartResult{}
	for _, repoURL := range repos {
		destDir := filepath.Join("./cloned", strings.ReplaceAll(repoURL, "/", "_"))
		if err := cloneRepo(repoURL, destDir); err != nil {
			fmt.Println("Clone failed:", repoURL)
			continue
		}
		charts, _ := findCharts(destDir)
		for _, chartPath := range charts {
			valuesPath := filepath.Join(chartPath, "values.yaml")
			values, err := parseValuesYaml(valuesPath)
			if err != nil {
				continue
			}
			combos := generateValueCombinations(values)
			for _, combo := range combos {
				output, err := runHelmTemplate(chartPath, combo)
				res := ChartResult{
					RepoURL:      repoURL,
					ChartPath:    chartPath,
					ValuesFile:   valuesPath,
					ValuesCombo:  combo,
					Success:      err == nil,
					ErrorMessage: "",
				}
				if err != nil {
					res.ErrorMessage = output
				}
				results = append(results, res)
			}
		}
	}
	// Save results
	out, _ := json.MarshalIndent(results, "", "  ")
	_ = os.WriteFile("catalog_results.json", out, 0644)
	fmt.Println("Cataloging complete. Results saved to catalog_results.json.")
}
