package helmfix

import (
	"regexp"
	"strings"
)

const (
	KindNilPointer    = "nil_pointer"
	KindRequiredValue = "required_value"
	KindTypeMismatch  = "type_mismatch"
	KicsPlaceholder   = "kics-placeholder"
	MaxFixIterations  = 10
)

var (
	// atValuesRe captures the offending value path from a Go-template error's
	// "at <.Values.x.y ...>" clause (nil-pointer and wrong-type errors).
	atValuesRe = regexp.MustCompile(`at <(\.Values\.[^ |>]+)`)

	// valuesPathRe matches an explicit ".Values.x.y" path anywhere in a message.
	valuesPathRe = regexp.MustCompile(`\.Values\.([A-Za-z_]\w*(?:\.[A-Za-z_]\w*)*)`)

	// requiredPathRes recover a value path embedded in a custom required/fail
	// message, tried in priority order. All but the last require a dotted
	// (>=2 segment) path so we don't guess from a bare English word.
	requiredPathRes = []*regexp.Regexp{
		regexp.MustCompile(`(?i)missing required value\s+\.?([A-Za-z_][\w.]*)`),
		regexp.MustCompile(`(?i)please (?:define|specify|provide|set)\s+\.?([A-Za-z_]\w*(?:\.[A-Za-z_]\w*)+)`),
		regexp.MustCompile(`(?i)--set(?:-string|-file)?\s+\.?([A-Za-z_]\w*(?:\.[A-Za-z_]\w*)+)\s*=`),
		regexp.MustCompile(`\.?([A-Za-z_]\w*(?:\.[A-Za-z_]\w*)+)\s+is\s+(?:required|mandatory)`),
		// Dotted path immediately (or one word, e.g. "entry") before a bare
		// "required"/"mandatory", covering "x.y required!" / "x.y entry required".
		regexp.MustCompile(`\.?([A-Za-z_]\w*(?:\.[A-Za-z_]\w*)+)(?:\s+\w+)?\s+(?:required|mandatory)\b`),
		regexp.MustCompile(`\.([A-Za-z_]\w*)\s+is\s+(?:required|mandatory)`),
	}

	// requiredMarkers identify a custom "this value must be provided" failure
	// (the Helm `required` builtin or a hand-written `fail` guard).
	requiredMarkers = []string{
		"error calling required",
		"required",
		"mandatory",
		"please define",
		"please specify",
		"please provide",
		"missing required value",
		"must be set",
		"must specify",
		"must provide",
	}
)

// ParseError parses a Helm error string (CLI "Error: ..." or raw SDK error)
// and returns the fix kind, value path, and value to inject.
// ok is false for errors that cannot be fixed by value injection.
func ParseError(errStr string) (kind, path, value string, ok bool) {
	s := strings.TrimPrefix(errStr, "Error: ")

	// 1. nil pointer: inject an empty value so the chain can build the missing
	//    parent map and re-render.
	if strings.Contains(s, "nil pointer") {
		if m := atValuesRe.FindStringSubmatch(s); m != nil {
			return KindNilPointer, strings.TrimPrefix(m[1], ".Values."), "", true
		}
	}

	// 2. wrong type where a string was expected: inject a string placeholder at
	//    the offending .Values path so the scalar type-checks.
	if strings.Contains(s, "wrong type for value") && strings.Contains(s, "expected string") {
		if m := atValuesRe.FindStringSubmatch(s); m != nil {
			return KindTypeMismatch, strings.TrimPrefix(m[1], ".Values."), KicsPlaceholder, true
		}
	}

	// 3. required/fail guards: recover the value path from the message text and
	//    inject a placeholder.
	if isRequiredMessage(s) {
		if p := extractRequiredPath(s); p != "" {
			return KindRequiredValue, p, KicsPlaceholder, true
		}
	}

	return "", "", "", false
}

// isRequiredMessage reports whether s reads like a "value must be provided"
// failure raised by the `required` builtin or a `fail` guard.
func isRequiredMessage(s string) bool {
	low := strings.ToLower(s)
	for _, m := range requiredMarkers {
		if strings.Contains(low, m) {
			return true
		}
	}
	return false
}

// extractRequiredPath pulls a ".Values" value path out of a required/fail
// message, returning "" when no path can be recovered.
func extractRequiredPath(s string) string {
	// Explicit .Values.x.y references are unambiguous, so try them first on the
	// whole error string (covers `required .Values.x "msg"` SDK errors).
	if m := valuesPathRe.FindStringSubmatch(s); m != nil {
		return m[1]
	}
	// Otherwise inspect only the human message, stripping the SDK location
	// envelope so we don't mistake a "templates/foo.yaml" path for a value path.
	body := messageBody(s)
	for _, re := range requiredPathRes {
		if m := re.FindStringSubmatch(body); m != nil {
			return strings.TrimPrefix(m[1], ".")
		}
	}
	return ""
}

// messageBody strips the "execution error at (PATH:line:col): " envelope that
// the Helm SDK wraps around custom error messages, returning the message text.
func messageBody(s string) string {
	const marker = "execution error at ("
	if i := strings.Index(s, marker); i >= 0 {
		rest := s[i+len(marker):]
		if j := strings.Index(rest, "): "); j >= 0 {
			return rest[j+len("): "):]
		}
	}
	return s
}
