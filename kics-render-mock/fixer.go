package main

import (
	"fmt"

	"helm.sh/helm/v3/pkg/cli/values"

	"github.com/MabsIPCA/helm-tests/helmfix"
)

func applyPatch(orig *values.Options, patch map[string]string) *values.Options {
	extra := make([]string, 0, len(patch))
	for k, v := range patch {
		extra = append(extra, fmt.Sprintf("%s=%s", k, v))
	}
	return &values.Options{
		ValueFiles:    orig.ValueFiles,
		Values:        append(append([]string{}, orig.Values...), extra...),
		StringValues:  orig.StringValues,
		FileValues:    orig.FileValues,
		JSONValues:    orig.JSONValues,
		LiteralValues: orig.LiteralValues,
	}
}

func fixInvocation(chartPath string, inv invocation) FixedRenderEntry {
	entry := FixedRenderEntry{FixChain: []FixStep{}}

	patch := map[string]string{}
	seenErrs := map[string]bool{}

	for attempt := 1; attempt <= helmfix.MaxFixIterations; attempt++ {
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

		kind, path, value, ok := helmfix.ParseError(errStr)
		if !ok {
			entry.StopReason = "unfixable_error_kind"
			return entry
		}

		entry.FixChain = append(entry.FixChain, FixStep{
			Attempt:       attempt,
			ErrorSeen:     errStr,
			Kind:          kind,
			ValuePath:     path,
			ValueInjected: value,
		})
		patch[path] = value
	}

	entry.StopReason = "max_iterations"
	return entry
}

// fixFromInvocations executes the fix loop for pre-built invocations matched
// against their corresponding base RenderOutput entries. Used by runFixDir
// (test-suite mode) and runCatalogMode (catalog mode).
func fixFromInvocations(chartPath string, base RenderOutput, invocations []invocation) FixedRenderOutput {
	out := FixedRenderOutput{
		Suite:      base.Suite,
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

		loopResult := fixInvocation(chartPath, inv)
		fixed.Resolved = loopResult.Resolved
		fixed.StopReason = loopResult.StopReason
		fixed.FixChain = loopResult.FixChain
		fixed.FixedResult = loopResult.FixedResult
		out.Renders = append(out.Renders, fixed)
	}

	return out
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
