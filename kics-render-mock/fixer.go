package main

import (
	"fmt"
	"regexp"
	"strings"

	"helm.sh/helm/v3/pkg/cli/values"
)

const (
	kindNilPointer    = "nil_pointer"
	kindRequiredValue = "required_value"
	maxFixIterations  = 10
)

var (
	nilPtrRe   = regexp.MustCompile(`at <(\.Values\.[^>]+)>`)
	requiredRe = regexp.MustCompile(`required (\.Values\.[^ "]+)`)
)

type parsedFix struct {
	kind  string
	path  string
	value string
}

func parseError(errStr string) (parsedFix, bool) {
	if strings.Contains(errStr, "nil pointer") {
		if m := nilPtrRe.FindStringSubmatch(errStr); m != nil {
			path := strings.TrimPrefix(m[1], ".Values.")
			return parsedFix{kind: kindNilPointer, path: path, value: ""}, true
		}
	}
	if strings.Contains(errStr, "error calling required") {
		if m := requiredRe.FindStringSubmatch(errStr); m != nil {
			path := strings.TrimPrefix(m[1], ".Values.")
			return parsedFix{kind: kindRequiredValue, path: path, value: "kics-placeholder"}, true
		}
	}
	return parsedFix{}, false
}

func applyPatch(orig *values.Options, patch map[string]string) *values.Options {
	extra := make([]string, 0, len(patch))
	for k, v := range patch {
		extra = append(extra, fmt.Sprintf("%s=%s", k, v))
	}
	return &values.Options{
		ValueFiles:   orig.ValueFiles,
		Values:       append(append([]string{}, orig.Values...), extra...),
		StringValues: orig.StringValues,
		FileValues:   orig.FileValues,
	}
}
