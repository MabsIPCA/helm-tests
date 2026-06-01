package helmfix

import (
	"regexp"
	"strings"
)

const (
	KindNilPointer    = "nil_pointer"
	KindRequiredValue = "required_value"
	KicsPlaceholder   = "kics-placeholder"
	MaxFixIterations  = 10
)

var (
	nilPtrRe   = regexp.MustCompile(`at <(\.Values\.[^ |>]+)`)
	requiredRe = regexp.MustCompile(`required (\.Values\.[^ "]+)`)
)

// ParseError parses a Helm error string (CLI "Error: ..." or raw SDK error)
// and returns the fix kind, value path, and value to inject.
// ok is false for errors that cannot be fixed by value injection.
func ParseError(errStr string) (kind, path, value string, ok bool) {
	s := strings.TrimPrefix(errStr, "Error: ")
	if strings.Contains(s, "nil pointer") {
		if m := nilPtrRe.FindStringSubmatch(s); m != nil {
			return KindNilPointer, strings.TrimPrefix(m[1], ".Values."), "", true
		}
	}
	if strings.Contains(s, "error calling required") {
		if m := requiredRe.FindStringSubmatch(s); m != nil {
			return KindRequiredValue, strings.TrimPrefix(m[1], ".Values."), KicsPlaceholder, true
		}
	}
	return "", "", "", false
}
