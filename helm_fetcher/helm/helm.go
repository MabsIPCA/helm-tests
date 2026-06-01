package helm

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

const (
	depBuildMaxRetries     = 4
	depBuildInitialBackoff = 15 * time.Second
	depBuildBackoffFactor  = 2.0
)

// FindCharts recursively searches for Chart.yaml files starting from baseDir and returns the directories containing them.
func FindCharts(baseDir string) []string {
	var charts []string
	_ = filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.Name() == "Chart.yaml" {
			charts = append(charts, filepath.Dir(path))
		}
		return nil
	})
	return charts
}

// FindValuesFiles looks for .yaml/.yml files in the given chartDir, excluding Chart*.yaml/yml and values.yaml/yml.
func FindValuesFiles(chartDir string) []string {
	var all []string
	for _, pat := range []string{"*.yaml", "*.yml"} {
		matches, _ := filepath.Glob(filepath.Join(chartDir, pat))
		for _, m := range matches {
			base := filepath.Base(m)
			// Skip the default values file (Helm loads it automatically)
			if base == "values.yaml" || base == "values.yml" {
				continue
			}
			// skip non *values* yaml/yml files (e.g. Chart.yaml, Chart-svc.yaml, etc.)
			if !strings.Contains(strings.ToLower(base), "values") {
				continue
			}
			// Skip any Chart*.yaml / Chart*.yml (chart descriptor files, e.g. Chart.yaml, Chart-svc.yaml)
			if strings.HasPrefix(strings.ToLower(base), "chart") {
				continue
			}
			// Skip anything that is not a regular file (e.g. symlinked dirs)
			info, err := os.Stat(m)
			if err != nil || !info.Mode().IsRegular() {
				continue
			}
			all = append(all, m)
		}
	}
	return all
}

// Combinations generates all non-empty Combinations of the input items.
func Combinations(items []string) [][]string {
	n := len(items)
	if n == 0 {
		return nil
	}
	total := 1 << n
	var out [][]string
	for mask := 1; mask < total; mask++ {
		var combo []string
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				combo = append(combo, items[i])
			}
		}
		out = append(out, combo)
	}
	return out
}

// RunHelmDependencyBuild executes "helm dependency build" for the given chart path and logs the result.
// It retries up to depBuildMaxRetries times with exponential backoff for transient failures.
// On permanent failure, the returned error contains the combined stdout+stderr output.
func RunHelmDependencyBuild(chartPath string) error {
	var lastErr error
	backoff := depBuildInitialBackoff
	maxAttempts := depBuildMaxRetries + 1

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		if attempt > 1 {
			log.Info().
				Str("chart", chartPath).
				Int("attempt", attempt).
				Int("max_attempts", maxAttempts).
				Dur("backoff", backoff).
				Msg("Retrying helm dependency build after backoff")
			time.Sleep(backoff)
			backoff = time.Duration(float64(backoff) * depBuildBackoffFactor)
		}

		cmd := exec.Command("helm", "dependency", "build", chartPath)
		out, err := cmd.CombinedOutput()
		if err == nil {
			if attempt > 1 {
				log.Info().Str("chart", chartPath).Int("attempt", attempt).Msg("helm dependency build succeeded on retry")
			} else {
				log.Info().Str("chart", chartPath).Msg("helm dependency build succeeded")
			}
			return nil
		}

		msg := strings.TrimSpace(string(out))
		if msg == "" {
			msg = err.Error()
		}
		lastErr = fmt.Errorf("%s", msg)
		log.Warn().
			Str("chart", chartPath).
			Str("output", msg).
			Int("attempt", attempt).
			Int("max_attempts", maxAttempts).
			Msg("helm dependency build failed")
	}

	log.Warn().Str("chart", chartPath).Int("attempts", maxAttempts).Msg("helm dependency build failed after all retries (continuing)")
	return lastErr
}

// RunHelmTemplate executes "helm template" for the given chart path and values files, returning the command string, output, and any error.
func RunHelmTemplate(chartPath string, valuesFiles []string) (cmdStr, output string, err error) {
	args := []string{"template", "test", chartPath}
	for _, vf := range valuesFiles {
		args = append(args, "-f", vf)
	}
	cmdStr = "helm " + strings.Join(args, " ")

	cmd := exec.Command("helm", args...)
	out, runErr := cmd.CombinedOutput()
	return cmdStr, string(out), runErr
}

// RunHelmTemplateWithSets runs "helm template test <chartPath> [-f vf...] [--set k=v...]"
// and returns the command string, combined output, and any exec error.
func RunHelmTemplateWithSets(chartPath string, valFiles, setFlags []string) (cmdStr, output string, err error) {
	args := []string{"template", "test", chartPath}
	for _, vf := range valFiles {
		args = append(args, "-f", vf)
	}
	for _, sf := range setFlags {
		args = append(args, "--set", sf)
	}
	cmdStr = "helm " + strings.Join(args, " ")

	cmd := exec.Command("helm", args...)
	out, runErr := cmd.CombinedOutput()
	return cmdStr, string(out), runErr
}
