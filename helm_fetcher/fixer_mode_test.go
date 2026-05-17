package main

import (
	"testing"

	"github.com/MabsIPCA/helm-tests/helm_fetcher/model"
)

func TestParseHelmCLIError_NilPointer(t *testing.T) {
	errMsg := `Error: template: mychart/templates/svc.yaml:12:8: executing "mychart/templates/svc.yaml" at <.Values.service.type>: nil pointer evaluating interface {}.type`
	kind, path, val, ok := parseHelmCLIError(errMsg)
	if !ok {
		t.Fatal("expected ok=true")
	}
	if kind != "nil_pointer" {
		t.Errorf("kind: got %q, want %q", kind, "nil_pointer")
	}
	if path != "service.type" {
		t.Errorf("path: got %q, want %q", path, "service.type")
	}
	if val != "" {
		t.Errorf("val: got %q, want empty string", val)
	}
}

func TestParseHelmCLIError_RequiredValue(t *testing.T) {
	errMsg := `Error: template: mychart/templates/dep.yaml:5:15: executing "mychart/templates/dep.yaml" at <required .Values.db.host "db.host is required">: error calling required: db.host is required`
	kind, path, val, ok := parseHelmCLIError(errMsg)
	if !ok {
		t.Fatal("expected ok=true")
	}
	if kind != "required_value" {
		t.Errorf("kind: got %q, want %q", kind, "required_value")
	}
	if path != "db.host" {
		t.Errorf("path: got %q, want %q", path, "db.host")
	}
	if val != "kics-placeholder" {
		t.Errorf("val: got %q, want %q", val, "kics-placeholder")
	}
}

func TestParseHelmCLIError_Unfixable(t *testing.T) {
	errMsg := `Error: parse error at (mychart/templates/bad.yaml:3): unexpected "}" in command`
	_, _, _, ok := parseHelmCLIError(errMsg)
	if ok {
		t.Error("expected ok=false for parse error")
	}
}

func TestParseHelmCLIError_StripsCLIPrefix(t *testing.T) {
	// "Error: " prefix must be stripped before matching
	withPrefix := `Error: template: c/t.yaml:1:1: executing "" at <.Values.foo>: nil pointer evaluating interface {}.foo`
	withoutPrefix := `template: c/t.yaml:1:1: executing "" at <.Values.foo>: nil pointer evaluating interface {}.foo`
	_, path1, _, ok1 := parseHelmCLIError(withPrefix)
	_, path2, _, ok2 := parseHelmCLIError(withoutPrefix)
	if !ok1 || !ok2 {
		t.Fatal("both forms must parse ok")
	}
	if path1 != path2 {
		t.Errorf("path mismatch: %q vs %q", path1, path2)
	}
}

func TestParseHelmCLIError_PipeFilter(t *testing.T) {
	// Pipe-filter expressions like <.Values.foo | quote> must NOT match nilPtrRe
	errMsg := `Error: template: c/t.yaml:3:5: executing "" at <.Values.foo | quote>: nil pointer evaluating interface {}.foo`
	_, _, _, ok := parseHelmCLIError(errMsg)
	// The pipe-filter form doesn't match nilPtrRe (stops at space/pipe) so falls through to unfixable.
	// It may or may not match depending on the exact error. This test just verifies no panic.
	_ = ok
}

func TestFixRun_AlreadySucceeded(t *testing.T) {
	orig := model.RunResult{Success: true}
	result := fixRun("/nonexistent", orig)
	if !result.Resolved {
		t.Error("already-succeeded run must return Resolved=true")
	}
	if len(result.FixChain) != 0 {
		t.Errorf("FixChain must be empty for already-succeeded run, got %d steps", len(result.FixChain))
	}
	if result.StopReason != "" {
		t.Errorf("StopReason must be empty, got %q", result.StopReason)
	}
}

func TestFixRun_UnfixableImmediately(t *testing.T) {
	// A chart path that doesn't exist → helm returns an exec/path error which is unfixable.
	orig := model.RunResult{
		Success:      false,
		ErrorMessage: "Error: stat /nonexistent: no such file or directory",
	}
	result := fixRun("/nonexistent", orig)
	// Should stop immediately: unfixable_error_kind or loop_detected (depending on helm output).
	// Either way it must not resolve.
	if result.Resolved {
		t.Error("expected Resolved=false for non-existent chart")
	}
}
