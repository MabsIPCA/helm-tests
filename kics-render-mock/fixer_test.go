package main

import (
	"testing"

	"helm.sh/helm/v3/pkg/cli/values"

	"github.com/MabsIPCA/helm-tests/helmfix"
)

func TestParseError_NilPointer(t *testing.T) {
	errStr := `template: mychart/templates/deploy.yaml:12:8: executing "mychart/templates/deploy.yaml" at <.Values.ingress.hosts>: nil pointer evaluating interface {}.hosts`
	kind, path, value, ok := helmfix.ParseError(errStr)
	if !ok {
		t.Fatal("expected parseable, got false")
	}
	if kind != helmfix.KindNilPointer {
		t.Errorf("kind: got %q, want %q", kind, helmfix.KindNilPointer)
	}
	if path != "ingress.hosts" {
		t.Errorf("path: got %q, want %q", path, "ingress.hosts")
	}
	if value != "" {
		t.Errorf("value: got %q, want empty string", value)
	}
}

func TestParseError_RequiredValue(t *testing.T) {
	errStr := `template: mychart/templates/deploy.yaml:5:15: executing "mychart/templates/deploy.yaml" at <required .Values.db.host "db.host is required">: error calling required: db.host is required`
	kind, path, value, ok := helmfix.ParseError(errStr)
	if !ok {
		t.Fatal("expected parseable, got false")
	}
	if kind != helmfix.KindRequiredValue {
		t.Errorf("kind: got %q, want %q", kind, helmfix.KindRequiredValue)
	}
	if path != "db.host" {
		t.Errorf("path: got %q, want %q", path, "db.host")
	}
	if value != helmfix.KicsPlaceholder {
		t.Errorf("value: got %q, want %q", value, helmfix.KicsPlaceholder)
	}
}

func TestParseError_Unfixable(t *testing.T) {
	errStr := `template: mychart/templates/deploy.yaml:1:1: executing "..." at <.Values.foo>: error calling fail: chart cannot be installed directly`
	_, _, _, ok := helmfix.ParseError(errStr)
	if ok {
		t.Fatal("expected unfixable, got parseable")
	}
}

func TestParseError_NilPointerWithoutAtClause(t *testing.T) {
	errStr := `template: mychart/templates/deploy.yaml:12: nil pointer evaluating something that has no at clause`
	_, _, _, ok := helmfix.ParseError(errStr)
	if ok {
		t.Fatal("expected unfixable without at <.Values.x> clause, got parseable")
	}
}

func TestParseError_NilPointerWithPipeFilter(t *testing.T) {
	errStr := `template: mychart/templates/deploy.yaml:5:8: executing "mychart/templates/deploy.yaml" at <.Values.foo | quote>: nil pointer evaluating interface {}.foo`
	_, path, _, ok := helmfix.ParseError(errStr)
	if !ok {
		t.Fatal("expected parseable, got false")
	}
	if path != "foo" {
		t.Errorf("path: got %q, want %q", path, "foo")
	}
}

func TestApplyPatch_AppendsToOriginalValues(t *testing.T) {
	orig := &values.Options{
		Values:     []string{"foo=bar"},
		ValueFiles: []string{"/some/values.yaml"},
	}
	patch := map[string]string{"db.host": helmfix.KicsPlaceholder}
	patched := applyPatch(orig, patch)

	if len(patched.ValueFiles) != 1 || patched.ValueFiles[0] != "/some/values.yaml" {
		t.Error("ValueFiles not preserved")
	}
	if len(patched.Values) != 2 {
		t.Fatalf("Values len: got %d, want 2", len(patched.Values))
	}
	if patched.Values[0] != "foo=bar" {
		t.Errorf("original value not at index 0: got %q", patched.Values[0])
	}
	if patched.Values[1] != "db.host="+helmfix.KicsPlaceholder {
		t.Errorf("patch not appended correctly: got %q", patched.Values[1])
	}
}

func TestApplyPatch_DoesNotMutateOriginal(t *testing.T) {
	orig := &values.Options{Values: []string{"foo=bar"}}
	_ = applyPatch(orig, map[string]string{"extra": "val"})
	if len(orig.Values) != 1 {
		t.Errorf("original Values was mutated: len=%d", len(orig.Values))
	}
}

func TestApplyPatch_EmptyPatch(t *testing.T) {
	orig := &values.Options{Values: []string{"foo=bar"}}
	patched := applyPatch(orig, map[string]string{})
	if len(patched.Values) != 1 {
		t.Errorf("Values len: got %d, want 1", len(patched.Values))
	}
}
