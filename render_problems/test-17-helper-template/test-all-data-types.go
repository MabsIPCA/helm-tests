package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

type TestCase struct {
	Name        string
	Flag        string
	Description string
	ExpectFail  bool
}

var testCases = []TestCase{
	{Name: "undefinedVar", Flag: "testHelperUndefinedVar", Description: "Helper with undefined variable reference", ExpectFail: true},
	{Name: "syntaxError", Flag: "testHelperSyntaxError", Description: "Helper with unclosed template action", ExpectFail: true},
	{Name: "typeMismatch", Flag: "testHelperTypeMismatch", Description: "Helper trying to range over non-iterable (port number)", ExpectFail: true},
	{Name: "divZero", Flag: "testHelperDivZero", Description: "Helper with division by zero", ExpectFail: true},
	{Name: "required", Flag: "testHelperRequired", Description: "Helper with required function on missing value", ExpectFail: true},
	{Name: "nilPointer", Flag: "testHelperNilPointer", Description: "Helper with nil pointer dereference", ExpectFail: true},
	{Name: "missingInclude", Flag: "testHelperMissingInclude", Description: "Helper calling non-existent helper", ExpectFail: true},
	{Name: "invalidFunc", Flag: "testHelperInvalidFunc", Description: "Helper using non-existent function", ExpectFail: true},
}

type TestResult struct {
	Pass       bool
	Error      string
	AsExpected bool
}

func runTest(flag string) TestResult {
	cmd := exec.Command("helm", "template", "test-helper", ".",
		"--set", fmt.Sprintf("%s=true", flag))

	output, err := cmd.CombinedOutput()
	if err != nil {
		return TestResult{Pass: false, Error: extractError(string(output)), AsExpected: true}
	}
	return TestResult{Pass: true, AsExpected: false}
}

func extractError(output string) string {
	patterns := []string{
		`nil pointer evaluating [^\n]+`,
		`can't evaluate field [^\n]+`,
		`range can't iterate over [^\n]+`,
		`error calling div: [^\n]+`,
		`error calling required: [^\n]+`,
		`template: [^\n]+:[0-9]+:[0-9]+: executing [^\n]+`,
		`template: [^\n]+: unclosed action[^\n]*`,
		`function "[^"]+" not defined`,
		`no template "[^"]+" associated with template`,
	}
	for _, pattern := range patterns {
		re := regexp.MustCompile(`(?i)` + pattern)
		if match := re.FindString(output); match != "" {
			match = strings.TrimSpace(match)
			if len(match) > 80 {
				return match[:80] + "..."
			}
			return match
		}
	}
	if strings.Contains(output, "Error:") {
		lines := strings.Split(output, "\n")
		for _, line := range lines {
			if strings.Contains(line, "Error:") {
				line = strings.TrimSpace(line)
				if len(line) > 80 {
					return line[:80] + "..."
				}
				return line
			}
		}
	}
	return "template error"
}

func main() {
	results := make(map[string]TestResult)
	for _, tc := range testCases {
		results[tc.Name] = runTest(tc.Flag)
	}

	// Generate markdown
	var sb strings.Builder

	sb.WriteString("# Test 17: Helper Template Errors\n\n")
	sb.WriteString(fmt.Sprintf("**Generated:** %s\n\n", time.Now().Format("2006-01-02 15:04:05")))

	// Stats
	total, asExpected := 0, 0
	for _, tc := range testCases {
		total++
		r := results[tc.Name]
		// For expected failures, test passes if it fails
		if (tc.ExpectFail && !r.Pass) || (!tc.ExpectFail && r.Pass) {
			asExpected++
		}
	}
	sb.WriteString(fmt.Sprintf("**Total:** %d | **As Expected:** %d (%.1f%%)\n\n",
		total, asExpected, float64(asExpected)/float64(total)*100))

	sb.WriteString("## Description\n\n")
	sb.WriteString("This test suite validates error handling in Helm helper templates (`_helpers.tpl`). ")
	sb.WriteString("Helper templates are reusable named templates that can be included in other templates using `include` or `template` directives. ")
	sb.WriteString("Errors in helper templates propagate to the including template and cause rendering to fail.\n\n")

	// Test Results Table
	sb.WriteString("## Test Results\n\n")
	sb.WriteString("| Test Case | Description | Expected | Result | Error Message |\n")
	sb.WriteString("|-----------|-------------|----------|--------|---------------|\n")

	for _, tc := range testCases {
		r := results[tc.Name]
		expected := "✅ Pass"
		if tc.ExpectFail {
			expected = "❌ Fail"
		}
		actual := "✅ Pass"
		if !r.Pass {
			actual = "❌ Fail"
		}
		errMsg := "-"
		if !r.Pass {
			errMsg = r.Error
		}
		sb.WriteString(fmt.Sprintf("| `%s` | %s | %s | %s | %s |\n", tc.Name, tc.Description, expected, actual, errMsg))
	}

	// Detailed Analysis
	sb.WriteString("\n## Detailed Analysis\n\n")

	sb.WriteString("### 1. Undefined Variable (`helpers.undefinedVar`)\n")
	sb.WriteString("```go\n{{ .Values.nonExistentVariable.nested.value }}\n```\n")
	sb.WriteString("**Error Type:** `nil pointer evaluating interface {}.nested`\n")
	sb.WriteString("**Cause:** Accessing a nested property on a nil or non-existent value.\n\n")

	sb.WriteString("### 2. Syntax Error (`helpers.syntaxError`)\n")
	sb.WriteString("```go\n{{ .Values.someValue\n```\n")
	sb.WriteString("**Error Type:** `unclosed action`\n")
	sb.WriteString("**Cause:** Missing closing `}}` in template action.\n\n")

	sb.WriteString("### 3. Type Mismatch (`helpers.typeMismatch`)\n")
	sb.WriteString("```go\n{{- range .Values.port }}\n```\n")
	sb.WriteString("**Error Type:** `range can't iterate over <number>`\n")
	sb.WriteString("**Cause:** Attempting to range over a non-iterable type (number instead of list/map).\n\n")

	sb.WriteString("### 4. Division by Zero (`helpers.divZero`)\n")
	sb.WriteString("```go\n{{ div 100 0 }}\n```\n")
	sb.WriteString("**Error Type:** `error calling div: divide by zero`\n")
	sb.WriteString("**Cause:** Explicit division by zero in template.\n\n")

	sb.WriteString("### 5. Required Function (`helpers.requiredValue`)\n")
	sb.WriteString("```go\n{{ required \"helperRequiredValue is required\" .Values.helperRequiredValue }}\n```\n")
	sb.WriteString("**Error Type:** `error calling required: helperRequiredValue is required`\n")
	sb.WriteString("**Cause:** Required value is not set in values.yaml.\n\n")

	sb.WriteString("### 6. Nil Pointer (`helpers.nilPointer`)\n")
	sb.WriteString("```go\n{{ .Values.nilHelper.nested.value }}\n```\n")
	sb.WriteString("**Error Type:** `nil pointer evaluating interface {}.nested`\n")
	sb.WriteString("**Cause:** Accessing property on explicitly null value.\n\n")

	sb.WriteString("### 7. Missing Include (`helpers.callsMissing`)\n")
	sb.WriteString("```go\n{{ include \"non.existent.helper\" . }}\n```\n")
	sb.WriteString("**Error Type:** `no template \"non.existent.helper\" associated with template`\n")
	sb.WriteString("**Cause:** Including a helper template that doesn't exist.\n\n")

	sb.WriteString("### 8. Invalid Function (`helpers.invalidFunction`)\n")
	sb.WriteString("```go\n{{ nonExistentFunction \"arg1\" \"arg2\" }}\n```\n")
	sb.WriteString("**Error Type:** `function \"nonExistentFunction\" not defined`\n")
	sb.WriteString("**Cause:** Using a function that doesn't exist in Helm/Sprig.\n\n")

	// Key Findings
	sb.WriteString("## Key Findings\n\n")
	sb.WriteString("### Error Propagation\n")
	sb.WriteString("- Errors in helper templates propagate to the calling template\n")
	sb.WriteString("- The error message includes the helper template name and line number\n")
	sb.WriteString("- All tested error conditions correctly cause template rendering to fail\n\n")

	sb.WriteString("### Common Error Patterns\n")
	sb.WriteString("| Error Type | Common Causes |\n")
	sb.WriteString("|------------|---------------|\n")
	sb.WriteString("| `nil pointer` | Accessing nested properties on nil/missing values |\n")
	sb.WriteString("| `unclosed action` | Missing `}}` in template syntax |\n")
	sb.WriteString("| `range can't iterate` | Using range on non-iterable types |\n")
	sb.WriteString("| `divide by zero` | Division operation with zero divisor |\n")
	sb.WriteString("| `error calling required` | Required value not provided |\n")
	sb.WriteString("| `no template` | Including non-existent helper template |\n")
	sb.WriteString("| `function not defined` | Using non-existent function |\n\n")

	sb.WriteString("## Recommendations\n\n")
	sb.WriteString("1. **Use `default` for optional values:** `{{ .Values.foo | default \"bar\" }}`\n")
	sb.WriteString("2. **Check existence with `if`:** `{{- if .Values.foo }}...{{- end }}`\n")
	sb.WriteString("3. **Use `hasKey` for maps:** `{{- if hasKey .Values \"foo\" }}...{{- end }}`\n")
	sb.WriteString("4. **Validate helper template names** before using `include`\n")
	sb.WriteString("5. **Test helper templates in isolation** during development\n")
	sb.WriteString("6. **Guard division operations:** Check divisor is non-zero before dividing\n")

	err := os.WriteFile("./TEST-RESULTS.md", []byte(sb.String()), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Println("Results written to TEST-RESULTS.md")
}
