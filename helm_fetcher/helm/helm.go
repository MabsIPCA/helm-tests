package helm

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
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
//
// Two classes of non-standalone charts are skipped because rendering them on
// their own yields false failures (missing parent globals / shared named
// templates / coalesced values):
//   - vendored subcharts under a parent's "charts/" directory (IsVendoredSubchart), and
//   - relative components referenced via a "file://" dependency by a sibling
//     chart in the same tree (RelativeComponentDirs) — the umbrella that owns
//     them is still discovered and renders them inline.
func FindCharts(baseDir string) []string {
	components := RelativeComponentDirs(baseDir)
	var charts []string
	_ = filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.IsDir() || info.Name() != "Chart.yaml" {
			return nil
		}
		chartDir := filepath.Dir(path)
		if IsVendoredSubchart(chartDir) || components[filepath.Clean(chartDir)] {
			return nil
		}
		charts = append(charts, chartDir)
		return nil
	})
	return charts
}

// fileDepRe captures the target of a "file://" Helm dependency repository.
var fileDepRe = regexp.MustCompile(`file://(\S+)`)

// RelativeComponentDirs scans every Chart.yaml under baseDir for dependencies
// declared with a "file://" repository and returns the set of chart directories
// those point at (cleaned, absolute). Such a chart is a component of the
// referencing umbrella and must be rendered through it, not standalone.
func RelativeComponentDirs(baseDir string) map[string]bool {
	out := map[string]bool{}
	_ = filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || info.Name() != "Chart.yaml" {
			return nil
		}
		data, readErr := os.ReadFile(path)
		if readErr != nil {
			return nil
		}
		ownerDir := filepath.Dir(path)
		for _, m := range fileDepRe.FindAllStringSubmatch(string(data), -1) {
			rel := strings.Trim(m[1], `"'`)
			target := rel
			if !filepath.IsAbs(target) {
				target = filepath.Join(ownerDir, rel)
			}
			out[filepath.Clean(target)] = true
		}
		return nil
	})
	return out
}

// IsVendoredSubchart reports whether chartDir is a subchart vendored under a
// parent chart's "charts/" directory, i.e. the Helm dependency layout
// <parent>/charts/<subchart> where <parent> has its own Chart.yaml. Such a
// chart must be rendered through its parent, not on its own.
//
// A "charts/" directory whose parent is NOT itself a chart (e.g. a monorepo
// that simply stores independent charts under charts/) is not treated as a
// vendoring boundary, so those charts are still discovered.
func IsVendoredSubchart(chartDir string) bool {
	dir := filepath.Clean(chartDir)
	for {
		parent := filepath.Dir(dir)
		if parent == dir {
			return false // reached filesystem root
		}
		if filepath.Base(dir) == "charts" && fileExists(filepath.Join(parent, "Chart.yaml")) {
			return true
		}
		dir = parent
	}
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.Mode().IsRegular()
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

// RunHelmTemplateFix runs "helm template test <chartPath> [-f vf...] [--set k=v...]
// [--kube-version v]" and returns the command string, combined output, and any
// exec error. kubeVersion is omitted when empty. Used by the fix loop so it can
// override both values (--set) and the Kubernetes version in one render.
func RunHelmTemplateFix(chartPath string, valFiles, setFlags []string, kubeVersion string) (cmdStr, output string, err error) {
	args := []string{"template", "test", chartPath}
	for _, vf := range valFiles {
		args = append(args, "-f", vf)
	}
	for _, sf := range setFlags {
		args = append(args, "--set", sf)
	}
	if kubeVersion != "" {
		args = append(args, "--kube-version", kubeVersion)
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
