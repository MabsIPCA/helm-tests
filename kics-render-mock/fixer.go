package main

import (
	"fmt"
	"os/exec"
	"strings"

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

// fixInvocation mirrors helm_fetcher's fixRun (fixer_mode.go) step for step, so
// that the only difference between the two fix rates is the render path itself
// -- exec "helm template" there, action.Install.Run() here. Keep the two in sync:
// a capability present in one and not the other makes the comparison measure the
// implementations rather than the SDK.
func fixInvocation(chartPath string, inv invocation) FixedRenderEntry {
	entry := FixedRenderEntry{FixChain: []FixStep{}}

	patch := map[string]string{}
	seenErrs := map[string]bool{}
	kubeVersion := ""   // --kube-version override, set when a kubeVersion error is seen
	depFetched := false // dependencies fetched at most once per invocation

	for attempt := 1; attempt <= helmfix.MaxFixIterations; attempt++ {
		patchedOpts := applyPatch(inv.valOpts, patch)
		res := runOnceKube(chartPath, patchedOpts, false, kubeVersion)

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

		// Chart or values file Helm cannot parse: never fixable by injection.
		if helmfix.IsUnparseableYAML(errStr) {
			entry.StopReason = "non_fixable_yaml"
			return entry
		}

		// Nil pointer on a Helm built-in is not a value, so --set cannot fix it.
		if helmfix.IsUnsupportedBuiltin(errStr) {
			entry.StopReason = "non_fixable_builtin"
			return entry
		}

		// Missing subcharts: fetch once (build, then update) and re-render, which
		// unmasks the chart's real, deeper error. A recurring dep error then falls
		// through to the unfixable path below.
		if !depFetched && helmfix.IsMissingDependency(errStr) {
			depFetched = true
			outcome := "fetched dependencies (build/update)"
			if fetchOut, fetchErr := runHelmDependencyFetch(chartPath); fetchErr != nil {
				outcome = "dependency fetch failed: " + firstLine(fetchOut)
			}
			entry.FixChain = append(entry.FixChain, FixStep{
				Attempt:       attempt,
				ErrorSeen:     errStr,
				Kind:          "dependency_fetch",
				ValueInjected: outcome,
			})
			continue
		}

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
		// kubeVersion is a render-wide override, not a --set at a value path.
		if kind == helmfix.KindKubeVersion {
			kubeVersion = value
		} else {
			patch[path] = value
		}
	}

	// Budget exhausted: the fix applied on the final iteration was never
	// verified, so render once more before giving up.
	if res := runOnceKube(chartPath, applyPatch(inv.valOpts, patch), false, kubeVersion); res.err == nil {
		entry.Resolved = true
		fixed := toStandardResult(res)
		entry.FixedResult = &fixed
		return entry
	}

	entry.StopReason = "max_iterations"
	return entry
}

// runHelmDependencyFetch mirrors helm_fetcher's helm.RunHelmDependencyFetch:
// try "helm dependency build" and fall back to "helm dependency update".
// Vendoring subcharts is a filesystem step, not a render, so it stays identical
// to the fetcher and leaves action.Install.Run() the only path under comparison.
func runHelmDependencyFetch(chartPath string) (string, error) {
	if out, err := runDepSubcommand(chartPath, "build"); err == nil {
		return out, nil
	}
	return runDepSubcommand(chartPath, "update")
}

func runDepSubcommand(chartPath, sub string) (string, error) {
	out, err := exec.Command("helm", "dependency", sub, chartPath).CombinedOutput()
	return string(out), err
}

func firstLine(s string) string {
	if i := strings.IndexAny(s, "\r\n"); i >= 0 {
		return s[:i]
	}
	return s
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
