package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

type TestFunction struct {
	Name string
	Func string
}

var testFunctions = []TestFunction{
	{Name: "testPrint", Func: "print"},
	{Name: "testPrintln", Func: "println"},
	{Name: "testPrintf", Func: "printf"},
	{Name: "testTrim", Func: "trim"},
	{Name: "testTrimAll", Func: "trimAll"},
	{Name: "testTrimPrefix", Func: "trimPrefix"},
	{Name: "testTrimSuffix", Func: "trimSuffix"},
	{Name: "testLower", Func: "lower"},
	{Name: "testUpper", Func: "upper"},
	{Name: "testTitle", Func: "title"},
	{Name: "testUntitle", Func: "untitle"},
	{Name: "testRepeat", Func: "repeat"},
	{Name: "testSubstr", Func: "substr"},
	{Name: "testNospace", Func: "nospace"},
	{Name: "testTrunc", Func: "trunc"},
	{Name: "testAbbrev", Func: "abbrev"},
	{Name: "testAbbrevboth", Func: "abbrevboth"},
	{Name: "testInitials", Func: "initials"},
	{Name: "testRandAlphaNum", Func: "randAlphaNum"},
	{Name: "testRandAlpha", Func: "randAlpha"},
	{Name: "testRandNumeric", Func: "randNumeric"},
	{Name: "testRandAscii", Func: "randAscii"},
	{Name: "testWrap", Func: "wrap"},
	{Name: "testWrapWith", Func: "wrapWith"},
	{Name: "testContains", Func: "contains"},
	{Name: "testHasPrefix", Func: "hasPrefix"},
	{Name: "testHasSuffix", Func: "hasSuffix"},
	{Name: "testQuote", Func: "quote"},
	{Name: "testSquote", Func: "squote"},
	{Name: "testCat", Func: "cat"},
	{Name: "testIndent", Func: "indent"},
	{Name: "testNindent", Func: "nindent"},
	{Name: "testReplace", Func: "replace"},
	{Name: "testPlural", Func: "plural"},
	{Name: "testSnakecase", Func: "snakecase"},
	{Name: "testCamelcase", Func: "camelcase"},
	{Name: "testKebabcase", Func: "kebabcase"},
	{Name: "testSwapcase", Func: "swapcase"},
	{Name: "testShuffle", Func: "shuffle"},
	{Name: "testAtoi", Func: "atoi"},
	{Name: "testInt64", Func: "int64"},
	{Name: "testFloat64", Func: "float64"},
}

var dataTypes = []string{
	"nilValue",
	"stringValue",
	"numberValue",
	"boolValue",
	"listValue",
	"dictValue",
}

type TestResult struct {
	Pass  bool
	Error string
}

func runTest(funcName, dataType string) TestResult {
	cmd := exec.Command("helm", "template", "test-string", ".",
		"--set", fmt.Sprintf("%s=true", funcName),
		"--set", fmt.Sprintf("testDataType=%s", dataType))

	output, err := cmd.CombinedOutput()
	if err != nil {
		return TestResult{Pass: false, Error: extractError(string(output))}
	}
	return TestResult{Pass: true}
}

func extractError(output string) string {
	patterns := []string{
		`wrong type[^;]*; expected [^;]+; got [^\s]+`,
		`invalid value[^;]*; expected [^\s]+`,
		`error calling \w+: [^:]+`,
	}
	for _, pattern := range patterns {
		re := regexp.MustCompile(`(?i)` + pattern)
		if match := re.FindString(output); match != "" {
			if len(match) > 50 {
				return match[:50] + "..."
			}
			return match
		}
	}
	return "template error"
}

func main() {
	// results[function][dataType]
	results := make(map[string]map[string]TestResult)
	for _, tf := range testFunctions {
		results[tf.Func] = make(map[string]TestResult)
		for _, dt := range dataTypes {
			results[tf.Func][dt] = runTest(tf.Name, dt)
		}
	}

	// Generate markdown
	var sb strings.Builder

	sb.WriteString("# Test 13: String Operation Errors\n\n")
	sb.WriteString(fmt.Sprintf("**Generated:** %s\n\n", time.Now().Format("2006-01-02 15:04:05")))

	// Stats
	total, passed := 0, 0
	for _, tf := range testFunctions {
		for _, dt := range dataTypes {
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
	for _, dt := range dataTypes {
		sb.WriteString(fmt.Sprintf(" %s |", dt))
	}
	sb.WriteString("\n|----------|")
	for range dataTypes {
		sb.WriteString("------|")
	}
	sb.WriteString("\n")

	for _, tf := range testFunctions {
		sb.WriteString(fmt.Sprintf("| `%s` |", tf.Func))
		for _, dt := range dataTypes {
			if results[tf.Func][dt].Pass {
				sb.WriteString(" ✅ |")
			} else {
				sb.WriteString(" ❌ |")
			}
		}
		sb.WriteString("\n")
	}

	// Failure details - one row per function, columns for each data type
	sb.WriteString("\n## Failure Details\n\n")
	sb.WriteString("| Function |")
	for _, dt := range dataTypes {
		sb.WriteString(fmt.Sprintf(" %s |", dt))
	}
	sb.WriteString("\n|----------|")
	for range dataTypes {
		sb.WriteString("------|")
	}
	sb.WriteString("\n")

	for _, tf := range testFunctions {
		// Check if this function has any failures
		hasFailure := false
		for _, dt := range dataTypes {
			if !results[tf.Func][dt].Pass {
				hasFailure = true
				break
			}
		}
		if !hasFailure {
			continue
		}

		sb.WriteString(fmt.Sprintf("| `%s` |", tf.Func))
		for _, dt := range dataTypes {
			if results[tf.Func][dt].Pass {
				sb.WriteString(" - |")
			} else {
				sb.WriteString(fmt.Sprintf(" %s |", results[tf.Func][dt].Error))
			}
		}
		sb.WriteString("\n")
	}

	os.WriteFile("./TEST-RESULTS.md", []byte(sb.String()), 0644)
}
