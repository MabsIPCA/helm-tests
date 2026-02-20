package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"
)

// TestConfig holds the configuration for a test
type TestConfig struct {
	TestNumber    int            `json:"testNumber"`
	TestName      string         `json:"testName"`
	ChartName     string         `json:"chartName"`
	DirName       string         `json:"-"`
	TestFunctions []TestFunction `json:"testFunctions"`
	DataTypes     []string       `json:"dataTypes"`
	ErrorPatterns []string       `json:"errorPatterns"`
	SubDirs       []string       `json:"subDirs,omitempty"`
}

// TestFunction defines a function to test
type TestFunction struct {
	Name   string `json:"name"`
	Func   string `json:"func"`
	SubDir string `json:"subDir,omitempty"`
}

// TestResult holds the result of a single test
type TestResult struct {
	Pass  bool
	Error string
}

// Test directories to scan for config.json
var testDirs = []string{
	"test-01-logic-default-functions",
	"test-02-nil-pointer",
	"test-03-template-syntax-error",
	"test-04-missing-template",
	"test-05-division-by-zero",
	"test-06-invalid-regex",
	"test-07-file-read-error",
	"test-08-invalid-types",
	"test-09-list-operation-error",
	"test-10-string-operation-error",
	"test-11-map-dict-error",
	"test-12-date-time-error",
	"test-13-crypto-error",
	"test-14-helper-template",
	"test-15-tpl-function",
	"test-16-files-advanced",
	"test-17-math-operations",
	"test-18-uuid-functions",
	"test-19-url-functions",
	"test-20-encoding-functions",
	"test-21-reflection-functions",
	"test-22-semver-functions",
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

	// If subDirs are specified, load each sub-dir config and merge
	if len(config.SubDirs) > 0 {
		dataTypeSet := make(map[string]bool)
		for _, dt := range config.DataTypes {
			dataTypeSet[dt] = true
		}
		errorPatternSet := make(map[string]bool)
		for _, ep := range config.ErrorPatterns {
			errorPatternSet[ep] = true
		}

		for _, subDir := range config.SubDirs {
			subConfigPath := filepath.Join(testDir, subDir, "config.json")
			subData, err := os.ReadFile(subConfigPath)
			if err != nil {
				return nil, fmt.Errorf("failed to read sub-config %s: %v", subDir, err)
			}

			var subConfig TestConfig
			if err := json.Unmarshal(subData, &subConfig); err != nil {
				return nil, fmt.Errorf("failed to parse sub-config %s: %v", subDir, err)
			}

			// Tag each function with its subDir and merge
			for i := range subConfig.TestFunctions {
				subConfig.TestFunctions[i].SubDir = subDir
			}
			config.TestFunctions = append(config.TestFunctions, subConfig.TestFunctions...)

			// Merge dataTypes
			for _, dt := range subConfig.DataTypes {
				if !dataTypeSet[dt] {
					config.DataTypes = append(config.DataTypes, dt)
					dataTypeSet[dt] = true
				}
			}

			// Merge errorPatterns
			for _, ep := range subConfig.ErrorPatterns {
				if !errorPatternSet[ep] {
					config.ErrorPatterns = append(config.ErrorPatterns, ep)
					errorPatternSet[ep] = true
				}
			}
		}
	}

	// Default to a single "default" data type if none are specified
	if len(config.DataTypes) == 0 {
		config.DataTypes = []string{"default"}
	}

	return &config, nil
}

func runTest(testDir, chartName, funcName, dataType, subDir string) TestResult {
	runDir := testDir
	if subDir != "" {
		runDir = filepath.Join(testDir, subDir)
		chartName = subDir
	}
	args := []string{"template", chartName, ".", "--set", fmt.Sprintf("%s=true", funcName)}
	if dataType != "default" {
		args = append(args, "--set", fmt.Sprintf("testDataType=%s", dataType))
	}
	cmd := exec.Command("helm", args...)
	cmd.Dir = runDir

	output, err := cmd.CombinedOutput()
	if err != nil {
		return TestResult{Pass: false, Error: string(output)}
	}
	return TestResult{Pass: true}
}

func extractError(output string, patterns []string) string {
	for _, pattern := range patterns {
		re := regexp.MustCompile(`(?i)` + pattern)
		if match := re.FindString(output); match != "" {
			match = strings.TrimSpace(match)
			if len(match) > 60 {
				return match[:60] + "..."
			}
			return match
		}
	}
	if strings.Contains(output, "Error:") {
		lines := strings.Split(output, "\n")
		for _, line := range lines {
			if strings.Contains(line, "Error:") {
				line = strings.TrimSpace(line)
				if len(line) > 60 {
					return line[:60] + "..."
				}
				return line
			}
		}
	}
	return "template error"
}

func generateMarkdown(config *TestConfig, results map[string]map[string]TestResult) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("# Test %d: %s\n\n", config.TestNumber, config.TestName))
	sb.WriteString(fmt.Sprintf("**Generated:** %s\n\n", time.Now().Format("2006-01-02 15:04:05")))

	// Stats
	total, passed := 0, 0
	for _, tf := range config.TestFunctions {
		for _, dt := range config.DataTypes {
			total++
			if results[tf.Func][dt].Pass {
				passed++
			}
		}
	}
	sb.WriteString(fmt.Sprintf("**Total:** %d | **Passed:** %d (%.1f%%) | **Failed:** %d (%.1f%%)\n\n",
		total, passed, float64(passed)/float64(total)*100, total-passed, float64(total-passed)/float64(total)*100))

	// Matrix table
	sb.WriteString("## Test Matrix\n\n")
	sb.WriteString("| Function |")
	for _, dt := range config.DataTypes {
		sb.WriteString(fmt.Sprintf(" %s |", dt))
	}
	sb.WriteString("\n|----------|")
	for range config.DataTypes {
		sb.WriteString("------|")
	}
	sb.WriteString("\n")

	for _, tf := range config.TestFunctions {
		sb.WriteString(fmt.Sprintf("| `%s` |", tf.Func))
		for _, dt := range config.DataTypes {
			if results[tf.Func][dt].Pass {
				sb.WriteString(" ✅ |")
			} else {
				sb.WriteString(" ❌ |")
			}
		}
		sb.WriteString("\n")
	}

	// Failure details
	sb.WriteString("\n## Failure Details\n\n")
	sb.WriteString("| Function |")
	for _, dt := range config.DataTypes {
		sb.WriteString(fmt.Sprintf(" %s |", dt))
	}
	sb.WriteString("\n|----------|")
	for range config.DataTypes {
		sb.WriteString("------|")
	}
	sb.WriteString("\n")

	for _, tf := range config.TestFunctions {
		hasFailure := false
		for _, dt := range config.DataTypes {
			if !results[tf.Func][dt].Pass {
				hasFailure = true
				break
			}
		}
		if !hasFailure {
			continue
		}

		sb.WriteString(fmt.Sprintf("| `%s` |", tf.Func))
		for _, dt := range config.DataTypes {
			if results[tf.Func][dt].Pass {
				sb.WriteString(" - |")
			} else {
				errMsg := extractError(results[tf.Func][dt].Error, config.ErrorPatterns)
				sb.WriteString(fmt.Sprintf(" %s |", errMsg))
			}
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func runTestSuite(config *TestConfig, baseDir string) (int, int, error) {
	testDir := filepath.Join(baseDir, config.DirName)

	// Run all tests
	results := make(map[string]map[string]TestResult)
	var mu sync.Mutex
	var wg sync.WaitGroup
	for _, tf := range config.TestFunctions {
		tf := tf // capture loop variable
		wg.Add(1)
		go func() {
			defer wg.Done()
			funcResults := make(map[string]TestResult)
			for _, dt := range config.DataTypes {
				funcResults[dt] = runTest(testDir, config.ChartName, tf.Name, dt, tf.SubDir)
			}
			mu.Lock()
			results[tf.Func] = funcResults
			mu.Unlock()
		}()
	}
	wg.Wait()

	// Generate markdown
	markdown := generateMarkdown(config, results)

	// Write to file
	outputPath := filepath.Join(testDir, "TEST-RESULTS.md")
	err := os.WriteFile(outputPath, []byte(markdown), 0644)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to write results: %v", err)
	}

	// Calculate stats
	total, passed := 0, 0
	for _, tf := range config.TestFunctions {
		for _, dt := range config.DataTypes {
			total++
			if results[tf.Func][dt].Pass {
				passed++
			}
		}
	}

	return passed, total, nil
}

func main() {
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("Helm Template Error Tests")
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

	for _, dirName := range testDirs {
		testDir := filepath.Join(baseDir, dirName)

		// Load config from JSON
		config, err := loadConfig(testDir)
		if err != nil {
			fmt.Printf("\n[%s] ❌ ERROR: %v\n", dirName, err)
			allPassed = false
			continue
		}

		fmt.Printf("\n[Test %d] %s\n", config.TestNumber, config.TestName)
		fmt.Printf("  Directory: %s\n", config.DirName)

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
		}
		fmt.Printf("  %s Passed: %d/%d (%.1f%%)\n", status, passed, total, pct)
		fmt.Printf("  Results written to %s/TEST-RESULTS.md\n", config.DirName)
	}

	// Summary
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("Summary")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("Total Tests: %d\n", totalTests)
	fmt.Printf("Total Passed: %d (%.1f%%)\n", totalPassed, float64(totalPassed)/float64(totalTests)*100)
	fmt.Printf("Total Failed: %d (%.1f%%)\n", totalTests-totalPassed, float64(totalTests-totalPassed)/float64(totalTests)*100)
	fmt.Printf("Completed: %s\n", time.Now().Format("2006-01-02 15:04:05"))

	if !allPassed {
		os.Exit(1)
	}
}
