package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

// TestConfig holds the configuration for a test
type TestConfig struct {
	TestNumber       int                `json:"testNumber"`
	TestName         string             `json:"testName"`
	Description      string             `json:"description"`
	ChartName        string             `json:"chartName"`
	DirName          string             `json:"-"`
	TestCases        []TestCase         `json:"testCases"`
	ExpectedBehavior []ExpectedBehavior `json:"expectedBehavior,omitempty"`
}

// TestCase defines a single test case
type TestCase struct {
	Name           string                 `json:"name"`
	Description    string                 `json:"description"`
	ValuesFiles    []string               `json:"valuesFiles"`
	SetFlags       []string               `json:"setFlags,omitempty"`
	SetStringFlags []string               `json:"setStringFlags,omitempty"`
	Expected       map[string]interface{} `json:"expected"`
	ShouldFail     bool                   `json:"shouldFail,omitempty"`
}

// ExpectedBehavior documents expected override behavior
type ExpectedBehavior struct {
	Field       string `json:"field"`
	Behavior    string `json:"behavior"`
	Description string `json:"description"`
}

// TestResult holds the result of a single test
type TestResult struct {
	Pass         bool
	Error        string
	Before       map[string]interface{}
	After        map[string]interface{}
	Intermediate []IntermediateResult // Values after each override file/set flag
	Actual       map[string]interface{}
	Mismatches   []string
}

// IntermediateResult holds values after applying a specific override
type IntermediateResult struct {
	Label  string // e.g., "values-override-1.yaml" or "--set key=value"
	Values map[string]interface{}
}

// Test directories to scan for config.json
var testDirs = []string{
	"test-01-basic-data-types",
	"test-02-multiple-values-files",
	"test-03-values-with-defaults",
	"test-04-type-changes",
	"test-05-list-override",
	"test-06-deep-nested-override",
	"test-07-null-value-override",
	"test-08-subchart-override",
	"test-09-global-values-override",
	"test-10-set-flag-override",
}

func loadConfig(testDir string) (*TestConfig, error) {
	configPath := filepath.Join(testDir, "config.json")
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config: %v", err)
	}

	var config TestConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config: %v", err)
	}

	config.DirName = filepath.Base(testDir)
	return &config, nil
}

func runHelmTemplate(testDir string, chartName string, valuesFiles []string, setFlags []string, setStringFlags []string) (map[string]interface{}, error) {
	args := []string{"template", chartName, "."}

	for _, vf := range valuesFiles {
		args = append(args, "-f", vf)
	}

	for _, sf := range setFlags {
		args = append(args, "--set", sf)
	}

	for _, ssf := range setStringFlags {
		args = append(args, "--set-string", ssf)
	}

	cmd := exec.Command("helm", args...)
	cmd.Dir = testDir

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("helm template failed: %v\nOutput: %s", err, string(output))
	}

	// Parse the YAML output to extract values
	// We expect the template to output the values in a parseable format
	var result map[string]interface{}

	// Find the all-values section in the output
	outputStr := string(output)
	lines := strings.Split(outputStr, "\n")
	inAllValues := false
	var valuesLines []string

	for _, line := range lines {
		if strings.HasPrefix(line, "all-values:") {
			inAllValues = true
			continue
		}
		if inAllValues {
			if strings.HasPrefix(line, "---") || (len(line) > 0 && line[0] != ' ' && line[0] != '\t') {
				break
			}
			// Remove the 4-space indent from values-display template
			if len(line) >= 4 {
				valuesLines = append(valuesLines, line[4:])
			} else {
				valuesLines = append(valuesLines, line)
			}
		}
	}

	valuesYaml := strings.Join(valuesLines, "\n")
	if err := yaml.Unmarshal([]byte(valuesYaml), &result); err != nil {
		return nil, fmt.Errorf("failed to parse values output: %v", err)
	}

	return result, nil
}

func getNestedValue(data map[string]interface{}, path string) (interface{}, bool) {
	parts := strings.Split(path, ".")
	current := interface{}(data)

	for _, part := range parts {
		switch v := current.(type) {
		case map[string]interface{}:
			var ok bool
			current, ok = v[part]
			if !ok {
				return nil, false
			}
		default:
			return nil, false
		}
	}
	return current, true
}

func compareValues(expected, actual interface{}) bool {
	// Handle nil comparisons
	if expected == nil && actual == nil {
		return true
	}
	if expected == nil || actual == nil {
		return false
	}

	// Type-specific comparison
	switch e := expected.(type) {
	case map[string]interface{}:
		a, ok := actual.(map[string]interface{})
		if !ok {
			return false
		}
		for k, v := range e {
			av, ok := a[k]
			if !ok || !compareValues(v, av) {
				return false
			}
		}
		return true
	case []interface{}:
		a, ok := actual.([]interface{})
		if !ok {
			return false
		}
		if len(e) != len(a) {
			return false
		}
		for i, v := range e {
			if !compareValues(v, a[i]) {
				return false
			}
		}
		return true
	case float64:
		// JSON numbers are always float64
		switch a := actual.(type) {
		case float64:
			return e == a
		case int:
			return e == float64(a)
		case int64:
			return e == float64(a)
		}
		return false
	case int:
		switch a := actual.(type) {
		case int:
			return e == a
		case int64:
			return int64(e) == a
		case float64:
			return float64(e) == a
		}
		return false
	case string:
		a, ok := actual.(string)
		return ok && e == a
	case bool:
		a, ok := actual.(bool)
		return ok && e == a
	default:
		return fmt.Sprintf("%v", expected) == fmt.Sprintf("%v", actual)
	}
}

func runTest(testDir string, config *TestConfig, tc TestCase) TestResult {
	// Get "before" values - render without override files to get default values
	var before map[string]interface{}
	if len(tc.ValuesFiles) > 0 || len(tc.SetFlags) > 0 || len(tc.SetStringFlags) > 0 {
		beforeValues, err := runHelmTemplate(testDir, config.ChartName, []string{}, []string{}, []string{})
		if err == nil {
			before = beforeValues
		}
	}

	// Capture intermediate values when overrides are used
	var intermediate []IntermediateResult
	totalOverrides := len(tc.ValuesFiles) + len(tc.SetFlags) + len(tc.SetStringFlags)
	if totalOverrides >= 1 {
		// Build up values files incrementally
		for i := 1; i <= len(tc.ValuesFiles); i++ {
			vals, err := runHelmTemplate(testDir, config.ChartName, tc.ValuesFiles[:i], []string{}, []string{})
			if err == nil {
				intermediate = append(intermediate, IntermediateResult{
					Label:  tc.ValuesFiles[i-1],
					Values: vals,
				})
			}
		}
		// Add set flags incrementally after all values files
		for i := 1; i <= len(tc.SetFlags); i++ {
			vals, err := runHelmTemplate(testDir, config.ChartName, tc.ValuesFiles, tc.SetFlags[:i], []string{})
			if err == nil {
				intermediate = append(intermediate, IntermediateResult{
					Label:  "--set " + tc.SetFlags[i-1],
					Values: vals,
				})
			}
		}
		// Add set-string flags incrementally
		for i := 1; i <= len(tc.SetStringFlags); i++ {
			vals, err := runHelmTemplate(testDir, config.ChartName, tc.ValuesFiles, tc.SetFlags, tc.SetStringFlags[:i])
			if err == nil {
				intermediate = append(intermediate, IntermediateResult{
					Label:  "--set-string " + tc.SetStringFlags[i-1],
					Values: vals,
				})
			}
		}
	}

	// Get "after" values - render with all override files
	actual, err := runHelmTemplate(testDir, config.ChartName, tc.ValuesFiles, tc.SetFlags, tc.SetStringFlags)
	if err != nil {
		if tc.ShouldFail {
			return TestResult{Pass: true, Error: "Expected failure occurred", Before: before, Intermediate: intermediate}
		}
		return TestResult{Pass: false, Error: err.Error(), Before: before, Intermediate: intermediate}
	}

	if tc.ShouldFail {
		return TestResult{Pass: false, Error: "Expected failure but succeeded", Before: before, After: actual, Intermediate: intermediate}
	}

	// Compare expected values
	var mismatches []string
	for path, expectedValue := range tc.Expected {
		actualValue, found := getNestedValue(actual, path)
		if !found {
			// If expected is null and key is not found, this is a pass
			// because Helm removes null keys from the output
			if expectedValue == nil {
				continue
			}
			mismatches = append(mismatches, fmt.Sprintf("%s: not found (expected: %v)", path, expectedValue))
			continue
		}
		if !compareValues(expectedValue, actualValue) {
			mismatches = append(mismatches, fmt.Sprintf("%s: expected %v (%T), got %v (%T)",
				path, expectedValue, expectedValue, actualValue, actualValue))
		}
	}

	if len(mismatches) > 0 {
		return TestResult{Pass: false, Mismatches: mismatches, Before: before, After: actual, Intermediate: intermediate, Actual: actual}
	}

	return TestResult{Pass: true, Before: before, After: actual, Intermediate: intermediate, Actual: actual}
}

// generateHelmCommand builds the helm template command string for display
func generateHelmCommand(chartName string, valuesFiles []string, setFlags []string, setStringFlags []string) string {
	var parts []string
	parts = append(parts, "helm", "template", chartName, ".")

	for _, vf := range valuesFiles {
		parts = append(parts, "-f", vf)
	}

	for _, sf := range setFlags {
		parts = append(parts, "--set", sf)
	}

	for _, ssf := range setStringFlags {
		parts = append(parts, "--set-string", ssf)
	}

	return strings.Join(parts, " ")
}

// formatValueForYaml formats a value as YAML-like string (single line for tables)
func formatValueForYaml(value interface{}) string {
	if value == nil {
		return "null"
	}
	switch v := value.(type) {
	case string:
		if v == "" {
			return `""`
		}
		return fmt.Sprintf(`"%s"`, v)
	case []interface{}:
		if len(v) == 0 {
			return "[]"
		}
		var items []string
		for _, item := range v {
			items = append(items, formatValueForYaml(item))
		}
		return "[" + strings.Join(items, ", ") + "]"
	case map[string]interface{}:
		if len(v) == 0 {
			return "{}"
		}
		// Sort keys for consistent output
		var keys []string
		for k := range v {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		var pairs []string
		for _, k := range keys {
			pairs = append(pairs, fmt.Sprintf("%s: %s", k, formatValueForYaml(v[k])))
		}
		return "{" + strings.Join(pairs, ", ") + "}"
	default:
		return fmt.Sprintf("%v", v)
	}
}

// isComplexValue checks if a value is a map or list (needs multi-line display)
func isComplexValue(value interface{}) bool {
	switch v := value.(type) {
	case []interface{}:
		return len(v) > 0
	case map[string]interface{}:
		return len(v) > 0
	default:
		return false
	}
}

// formatTableCell formats a value for display in a markdown table cell
func formatTableCell(value interface{}, found bool) string {
	if !found {
		return "_(not set)_"
	}
	return "`" + formatValueForYaml(value) + "`"
}

// formatValueBlock formats a value as a YAML code block
func formatValueBlock(value interface{}, found bool) string {
	if !found {
		return "```yaml\n# not set\n```"
	}
	formatted := formatValueForYaml(value)
	formatted = strings.TrimPrefix(formatted, "\n")
	return "```yaml\n" + formatted + "\n```"
}

// hasComplexValues checks if any expected field has complex values (maps/lists)
func hasComplexValues(before, after map[string]interface{}, intermediate []IntermediateResult, expected map[string]interface{}) bool {
	for path := range expected {
		if val, found := getNestedValue(before, path); found && isComplexValue(val) {
			return true
		}
		if val, found := getNestedValue(after, path); found && isComplexValue(val) {
			return true
		}
		for _, inter := range intermediate {
			if val, found := getNestedValue(inter.Values, path); found && isComplexValue(val) {
				return true
			}
		}
	}
	return false
}

// generateComparisonBlocks creates comparison using code blocks for each field
func generateComparisonBlocks(before, after map[string]interface{}, intermediate []IntermediateResult, expected map[string]interface{}) string {
	var sb strings.Builder

	// Sort paths for consistent output
	var paths []string
	for path := range expected {
		paths = append(paths, path)
	}
	sort.Strings(paths)

	for _, path := range paths {
		sb.WriteString(fmt.Sprintf("#### `%s`\n\n", path))

		beforeVal, beforeFound := getNestedValue(before, path)
		afterVal, afterFound := getNestedValue(after, path)

		if len(intermediate) > 1 {
			// Show progression through each override
			sb.WriteString("**Original:**\n")
			sb.WriteString(formatValueBlock(beforeVal, beforeFound) + "\n\n")

			for _, inter := range intermediate {
				val, found := getNestedValue(inter.Values, path)
				sb.WriteString(fmt.Sprintf("**After `%s`:**\n", inter.Label))
				sb.WriteString(formatValueBlock(val, found) + "\n\n")
			}

			sb.WriteString("**Final:**\n")
			sb.WriteString(formatValueBlock(afterVal, afterFound) + "\n\n")
		} else {
			// Simple before/after - no table, just two code blocks
			sb.WriteString("**Original:**\n\n")
			sb.WriteString(formatValueBlock(beforeVal, beforeFound) + "\n\n")
			sb.WriteString("**Final:**\n\n")
			sb.WriteString(formatValueBlock(afterVal, afterFound) + "\n\n")
		}
	}

	return sb.String()
}

// generateComparisonTable creates a side-by-side comparison table with intermediate values
func generateComparisonTable(before, after map[string]interface{}, intermediate []IntermediateResult, expected map[string]interface{}) string {
	var sb strings.Builder

	// Build header - always show intermediate columns if there are any
	if len(intermediate) >= 1 {
		// Multi-column table with intermediate values and final column
		sb.WriteString("| Field | Default |")
		for _, inter := range intermediate {
			label := "`" + inter.Label + "`"
			sb.WriteString(" " + label + " |")
		}
		sb.WriteString(" **Final** |\n")

		// Header separator
		sb.WriteString("|-------|---------|")
		for range intermediate {
			sb.WriteString("----------|")
		}
		sb.WriteString("----------|")
		sb.WriteString("\n")
	} else {
		// Simple two-column table (no overrides - shouldn't happen but fallback)
		sb.WriteString("| Field | Default | Final |\n")
		sb.WriteString("|-------|---------|-------|\n")
	}

	// Sort paths for consistent output
	var paths []string
	for path := range expected {
		paths = append(paths, path)
	}
	sort.Strings(paths)

	for _, path := range paths {
		beforeVal, beforeFound := getNestedValue(before, path)
		afterVal, afterFound := getNestedValue(after, path)

		beforeStr := formatTableCell(beforeVal, beforeFound)
		afterStr := formatTableCell(afterVal, afterFound)

		if len(intermediate) >= 1 {
			// Multi-column row with intermediate values and final
			sb.WriteString(fmt.Sprintf("| `%s` | %s |", path, beforeStr))
			for _, inter := range intermediate {
				val, found := getNestedValue(inter.Values, path)
				valStr := formatTableCell(val, found)
				sb.WriteString(" " + valStr + " |")
			}
			sb.WriteString(" " + afterStr + " |\n")
		} else {
			// Simple two-column row (fallback)
			sb.WriteString(fmt.Sprintf("| `%s` | %s | %s |\n", path, beforeStr, afterStr))
		}
	}

	return sb.String()
}

func generateMarkdown(config *TestConfig, results map[string]TestResult) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("# Test %d: %s\n\n", config.TestNumber, config.TestName))
	sb.WriteString(fmt.Sprintf("**Description:** %s\n\n", config.Description))
	sb.WriteString(fmt.Sprintf("**Generated:** %s\n\n", time.Now().Format("2006-01-02 15:04:05")))

	// Summary table
	sb.WriteString("## Summary\n\n")
	sb.WriteString("| Test Case | Description |\n")
	sb.WriteString("|-----------|-------------|\n")

	for _, tc := range config.TestCases {
		sb.WriteString(fmt.Sprintf("| [%s](#%s) | %s |\n", tc.Name, strings.ReplaceAll(tc.Name, " ", "-"), tc.Description))
	}

	// Test Case Details
	sb.WriteString("\n## Test Case Details\n\n")
	for _, tc := range config.TestCases {
		result := results[tc.Name]
		sb.WriteString(fmt.Sprintf("### %s\n\n", tc.Name))
		sb.WriteString(fmt.Sprintf("%s\n\n", tc.Description))

		// Show the helm command being executed
		helmCmd := generateHelmCommand(config.ChartName, tc.ValuesFiles, tc.SetFlags, tc.SetStringFlags)
		sb.WriteString("**Command:**\n\n")
		sb.WriteString(fmt.Sprintf("```bash\n%s\n```\n\n", helmCmd))

		// Show comparison table if there are override files or set flags
		if len(tc.ValuesFiles) > 0 || len(tc.SetFlags) > 0 || len(tc.SetStringFlags) > 0 {
			sb.WriteString("**Values Comparison:**\n\n")
			sb.WriteString(generateComparisonTable(result.Before, result.After, result.Intermediate, tc.Expected))
			sb.WriteString("\n")
		} else {
			// For default-values test, just show the values
			sb.WriteString("**Default Values:**\n\n")
			sb.WriteString("| Field | Value |\n")
			sb.WriteString("|-------|-------|\n")

			var paths []string
			for path := range tc.Expected {
				paths = append(paths, path)
			}
			sort.Strings(paths)

			for _, path := range paths {
				val, found := getNestedValue(result.After, path)
				valStr := "_(not set)_"
				if found {
					valStr = "`" + formatValueForYaml(val) + "`"
				}
				sb.WriteString(fmt.Sprintf("| `%s` | %s |\n", path, valStr))
			}
			sb.WriteString("\n")
		}

		// Show mismatches if any
		if !result.Pass && len(result.Mismatches) > 0 {
			sb.WriteString("**⚠️ Mismatches:**\n\n")
			sb.WriteString("```\n")
			for _, m := range result.Mismatches {
				sb.WriteString(m + "\n")
			}
			sb.WriteString("```\n\n")
		}

		if result.Error != "" {
			sb.WriteString(fmt.Sprintf("**⚠️ Error:**\n\n```\n%s\n```\n\n", result.Error))
		}

		sb.WriteString("---\n\n")
	}

	// Expected behaviors
	if len(config.ExpectedBehavior) > 0 {
		sb.WriteString("## Expected Override Behaviors\n\n")
		sb.WriteString("| Field | Behavior | Description |\n")
		sb.WriteString("|-------|----------|-------------|\n")
		for _, eb := range config.ExpectedBehavior {
			sb.WriteString(fmt.Sprintf("| `%s` | %s | %s |\n", eb.Field, eb.Behavior, eb.Description))
		}
	}

	return sb.String()
}

func runTestSuite(config *TestConfig, baseDir string) (int, int, error) {
	testDir := filepath.Join(baseDir, config.DirName)

	results := make(map[string]TestResult)
	for _, tc := range config.TestCases {
		results[tc.Name] = runTest(testDir, config, tc)
	}

	// Generate markdown
	markdown := generateMarkdown(config, results)

	// Write to file
	outputPath := filepath.Join(testDir, "TEST-RESULTS.md")
	err := os.WriteFile(outputPath, []byte(markdown), 0644)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to write results: %v", err)
	}

	// Calculate stats
	passed := 0
	for _, tc := range config.TestCases {
		if results[tc.Name].Pass {
			passed++
		}
	}

	return passed, len(config.TestCases), nil
}

func main() {
	specificTest := flag.String("test", "", "Run specific test (e.g., 01)")
	flag.Parse()

	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("Helm Values Override Tests")
	fmt.Printf("Started: %s\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(strings.Repeat("=", 60))

	baseDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		os.Exit(1)
	}

	allPassed := true
	totalTests := 0
	totalPassed := 0

	dirsToRun := testDirs
	if *specificTest != "" {
		dirsToRun = []string{}
		for _, d := range testDirs {
			if strings.Contains(d, fmt.Sprintf("test-%s-", *specificTest)) {
				dirsToRun = append(dirsToRun, d)
			}
		}
		if len(dirsToRun) == 0 {
			fmt.Printf("No test found matching: %s\n", *specificTest)
			os.Exit(1)
		}
	}

	for _, dirName := range dirsToRun {
		testDir := filepath.Join(baseDir, dirName)

		// Check if directory exists
		if _, err := os.Stat(testDir); os.IsNotExist(err) {
			fmt.Printf("\n[%s] ⚠️  SKIPPED: Directory not found\n", dirName)
			continue
		}

		// Load config from JSON
		config, err := loadConfig(testDir)
		if err != nil {
			fmt.Printf("\n[%s] ❌ ERROR: %v\n", dirName, err)
			allPassed = false
			continue
		}

		fmt.Printf("\n[Test %d] %s\n", config.TestNumber, config.TestName)
		fmt.Printf("  Directory: %s\n", config.DirName)
		fmt.Printf("  Test Cases: %d\n", len(config.TestCases))

		passed, total, err := runTestSuite(config, baseDir)
		if err != nil {
			fmt.Printf("  ❌ ERROR: %v\n", err)
			allPassed = false
			continue
		}

		totalTests += total
		totalPassed += passed

		pct := float64(passed) / float64(total) * 100
		status := "✅"
		if passed != total {
			status = "⚠️"
			allPassed = false
		}
		fmt.Printf("  %s Passed: %d/%d (%.1f%%)\n", status, passed, total, pct)
		fmt.Printf("  Results written to %s/TEST-RESULTS.md\n", config.DirName)
	}

	// Summary
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("Summary")
	fmt.Println(strings.Repeat("=", 60))
	if totalTests > 0 {
		fmt.Printf("Total Tests: %d\n", totalTests)
		fmt.Printf("Total Passed: %d (%.1f%%)\n", totalPassed, float64(totalPassed)/float64(totalTests)*100)
		fmt.Printf("Total Failed: %d (%.1f%%)\n", totalTests-totalPassed, float64(totalTests-totalPassed)/float64(totalTests)*100)
	} else {
		fmt.Println("No tests were executed.")
	}
	fmt.Printf("Completed: %s\n", time.Now().Format("2006-01-02 15:04:05"))

	if !allPassed {
		os.Exit(1)
	}
}
