package main

import (
	"testing"

	"helm.sh/helm/v3/pkg/cli/values"
)

func TestParseError_NilPointer(t *testing.T) {
	errStr := `template: mychart/templates/deploy.yaml:12:8: executing "mychart/templates/deploy.yaml" at <.Values.ingress.hosts>: nil pointer evaluating interface {}.hosts`
	fix, ok := parseError(errStr)
	if !ok {
		t.Fatal("expected parseable, got false")
	}
	if fix.kind != kindNilPointer {
		t.Errorf("kind: got %q, want %q", fix.kind, kindNilPointer)
	}
	if fix.path != "ingress.hosts" {
		t.Errorf("path: got %q, want %q", fix.path, "ingress.hosts")
	}
	if fix.value != "" {
		t.Errorf("value: got %q, want empty string", fix.value)
	}
}

func TestParseError_RequiredValue(t *testing.T) {
	errStr := `template: mychart/templates/deploy.yaml:5:15: executing "mychart/templates/deploy.yaml" at <required .Values.db.host "db.host is required">: error calling required: db.host is required`
	fix, ok := parseError(errStr)
	if !ok {
		t.Fatal("expected parseable, got false")
	}
	if fix.kind != kindRequiredValue {
		t.Errorf("kind: got %q, want %q", fix.kind, kindRequiredValue)
	}
	if fix.path != "db.host" {
		t.Errorf("path: got %q, want %q", fix.path, "db.host")
	}
	if fix.value != "kics-placeholder" {
		t.Errorf("value: got %q, want %q", fix.value, "kics-placeholder")
	}
}

func TestParseError_Unfixable(t *testing.T) {
	errStr := `template: mychart/templates/deploy.yaml:1:1: executing "..." at <.Values.foo>: error calling fail: chart cannot be installed directly`
	_, ok := parseError(errStr)
	if ok {
		t.Fatal("expected unfixable, got parseable")
	}
}

func TestParseError_NilPointerWithoutAtClause(t *testing.T) {
	errStr := `template: mychart/templates/deploy.yaml:12: nil pointer evaluating something that has no at clause`
	_, ok := parseError(errStr)
	if ok {
		t.Fatal("expected unfixable without at <.Values.x> clause, got parseable")
	}
}

func TestParseError_NilPointerWithPipeFilter(t *testing.T) {
	errStr := `template: mychart/templates/deploy.yaml:5:8: executing "mychart/templates/deploy.yaml" at <.Values.foo | quote>: nil pointer evaluating interface {}.foo`
	fix, ok := parseError(errStr)
	if !ok {
		t.Fatal("expected parseable, got false")
	}
	if fix.path != "foo" {
		t.Errorf("path: got %q, want %q", fix.path, "foo")
	}
}

func TestApplyPatch_AppendsToOriginalValues(t *testing.T) {
	orig := &values.Options{
		Values:     []string{"foo=bar"},
		ValueFiles: []string{"/some/values.yaml"},
	}
	patch := map[string]string{"db.host": "kics-placeholder"}
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
	if patched.Values[1] != "db.host=kics-placeholder" {
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
