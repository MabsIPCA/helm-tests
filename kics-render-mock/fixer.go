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
	nilPtrRe   = regexp.MustCompile(`at <(\.Values\.[^ |>]+)`)
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
		JSONValues:   orig.JSONValues,
		LiteralValues: orig.LiteralValues,
	}
}

func fixInvocation(chartPath string, inv invocation) FixedRenderEntry {
	entry := FixedRenderEntry{FixChain: []FixStep{}}

	patch := map[string]string{}
	seenErrs := map[string]bool{}

	for attempt := 1; attempt <= maxFixIterations; attempt++ {
		patchedOpts := applyPatch(inv.valOpts, patch)
		res := runOnce(chartPath, patchedOpts, false)

		if res.err == nil {
			entry.Resolved = true
			fixed := toStandardResult(res)
			entry.FixedResult = &fixed
			return entry
		}

		errStr := res.err.Error()
		if seenErrs[errStr] {
			entry.StopReason = "loop_detected"
			return entry
		}
		seenErrs[errStr] = true

		fix, ok := parseError(errStr)
		if !ok {
			entry.StopReason = "unfixable_error_kind"
			return entry
		}

		entry.FixChain = append(entry.FixChain, FixStep{
			Attempt:       attempt,
			ErrorSeen:     errStr,
			Kind:          fix.kind,
			ValuePath:     fix.path,
			ValueInjected: fix.value,
		})
		patch[fix.path] = fix.value
	}

	entry.StopReason = "max_iterations"
	return entry
}

func runFixDir(testDir, suite string, base RenderOutput) (FixedRenderOutput, error) {
	cfg, err := loadConfig(testDir)
	if err != nil {
		return FixedRenderOutput{}, err
	}
	invocations := buildInvocations(cfg, testDir)

	out := FixedRenderOutput{
		Suite:      suite,
		TestNumber: base.TestNumber,
		TestName:   base.TestName,
		ChartPath:  base.ChartPath,
		Renders:    make([]FixedRenderEntry, 0, len(base.Renders)),
	}

	for i, inv := range invocations {
		baseEntry := base.Renders[i]
		fixed := FixedRenderEntry{
			RenderEntry: baseEntry,
			FixChain:    []FixStep{},
		}

		if baseEntry.Standard.Error == nil {
			fixed.Resolved = true
			out.Renders = append(out.Renders, fixed)
			continue
		}

		loopResult := fixInvocation(testDir, inv)
		fixed.Resolved = loopResult.Resolved
		fixed.StopReason = loopResult.StopReason
		fixed.FixChain = loopResult.FixChain
		fixed.FixedResult = loopResult.FixedResult
		out.Renders = append(out.Renders, fixed)
	}

	return out, nil
}
