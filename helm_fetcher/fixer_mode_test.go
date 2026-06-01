package main

import (
	"testing"

	"github.com/MabsIPCA/helm-tests/helm_fetcher/model"
	"github.com/MabsIPCA/helm-tests/helmfix"
)

func TestParseError_NilPointer(t *testing.T) {
	errMsg := `Error: template: mychart/templates/svc.yaml:12:8: executing "mychart/templates/svc.yaml" at <.Values.service.type>: nil pointer evaluating interface {}.type`
	kind, path, val, ok := helmfix.ParseError(errMsg)
	if !ok {
		t.Fatal("expected ok=true")
	}
	if kind != helmfix.KindNilPointer {
		t.Errorf("kind: got %q, want %q", kind, helmfix.KindNilPointer)
	}
	if path != "service.type" {
		t.Errorf("path: got %q, want %q", path, "service.type")
	}
	if val != "" {
		t.Errorf("val: got %q, want empty string", val)
	}
}

func TestParseError_RequiredValue(t *testing.T) {
	errMsg := `Error: template: mychart/templates/dep.yaml:5:15: executing "mychart/templates/dep.yaml" at <required .Values.db.host "db.host is required">: error calling required: db.host is required`
	kind, path, val, ok := helmfix.ParseError(errMsg)
	if !ok {
		t.Fatal("expected ok=true")
	}
	if kind != helmfix.KindRequiredValue {
		t.Errorf("kind: got %q, want %q", kind, helmfix.KindRequiredValue)
	}
	if path != "db.host" {
		t.Errorf("path: got %q, want %q", path, "db.host")
	}
	if val != helmfix.KicsPlaceholder {
		t.Errorf("val: got %q, want %q", val, helmfix.KicsPlaceholder)
	}
}

func TestParseError_Unfixable(t *testing.T) {
	errMsg := `Error: parse error at (mychart/templates/bad.yaml:3): unexpected "}" in command`
	_, _, _, ok := helmfix.ParseError(errMsg)
	if ok {
		t.Error("expected ok=false for parse error")
	}
}

func TestParseError_StripsCLIPrefix(t *testing.T) {
	withPrefix := `Error: template: c/t.yaml:1:1: executing "" at <.Values.foo>: nil pointer evaluating interface {}.foo`
	withoutPrefix := `template: c/t.yaml:1:1: executing "" at <.Values.foo>: nil pointer evaluating interface {}.foo`
	_, path1, _, ok1 := helmfix.ParseError(withPrefix)
	_, path2, _, ok2 := helmfix.ParseError(withoutPrefix)
	if !ok1 || !ok2 {
		t.Fatal("both forms must parse ok")
	}
	if path1 != path2 {
		t.Errorf("path mismatch: %q vs %q", path1, path2)
	}
}

func TestParseError_PipeFilter(t *testing.T) {
	errMsg := `Error: template: c/t.yaml:3:5: executing "" at <.Values.foo | quote>: nil pointer evaluating interface {}.foo`
	kind, path, _, ok := helmfix.ParseError(errMsg)
	if !ok {
		t.Fatal("expected parseable, got false")
	}
	if kind != helmfix.KindNilPointer {
		t.Errorf("kind: got %q", kind)
	}
	if path != "foo" {
		t.Errorf("path: got %q, want %q", path, "foo")
	}
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
	orig := model.RunResult{
		Success:      false,
		ErrorMessage: "Error: stat /nonexistent: no such file or directory",
	}
	result := fixRun("/nonexistent", orig)
	if result.Resolved {
		t.Error("expected Resolved=false for non-existent chart")
	}
}
