package helmfix

import "testing"

func TestParseError_NilPointer(t *testing.T) {
	errStr := `template: mychart/templates/deploy.yaml:12:8: executing "mychart/templates/deploy.yaml" at <.Values.ingress.hosts>: nil pointer evaluating interface {}.hosts`
	kind, path, value, ok := ParseError(errStr)
	if !ok {
		t.Fatal("expected parseable, got false")
	}
	if kind != KindNilPointer {
		t.Errorf("kind: got %q, want %q", kind, KindNilPointer)
	}
	if path != "ingress.hosts" {
		t.Errorf("path: got %q, want %q", path, "ingress.hosts")
	}
	if value != "" {
		t.Errorf("value: got %q, want empty string", value)
	}
}

func TestParseError_NilPointerWithCLIPrefix(t *testing.T) {
	errStr := `Error: template: mychart/templates/svc.yaml:12:8: executing "mychart/templates/svc.yaml" at <.Values.service.type>: nil pointer evaluating interface {}.type`
	kind, path, _, ok := ParseError(errStr)
	if !ok {
		t.Fatal("expected parseable, got false")
	}
	if kind != KindNilPointer {
		t.Errorf("kind: got %q, want %q", kind, KindNilPointer)
	}
	if path != "service.type" {
		t.Errorf("path: got %q, want %q", path, "service.type")
	}
}

func TestParseError_RequiredValue(t *testing.T) {
	errStr := `template: mychart/templates/deploy.yaml:5:15: executing "mychart/templates/deploy.yaml" at <required .Values.db.host "db.host is required">: error calling required: db.host is required`
	kind, path, value, ok := ParseError(errStr)
	if !ok {
		t.Fatal("expected parseable, got false")
	}
	if kind != KindRequiredValue {
		t.Errorf("kind: got %q, want %q", kind, KindRequiredValue)
	}
	if path != "db.host" {
		t.Errorf("path: got %q, want %q", path, "db.host")
	}
	if value != KicsPlaceholder {
		t.Errorf("value: got %q, want %q", value, KicsPlaceholder)
	}
}

func TestParseError_Unfixable(t *testing.T) {
	errStr := `template: mychart/templates/deploy.yaml:1:1: executing "..." at <.Values.foo>: error calling fail: chart cannot be installed directly`
	_, _, _, ok := ParseError(errStr)
	if ok {
		t.Fatal("expected unfixable, got parseable")
	}
}

func TestParseError_NilPointerWithoutAtClause(t *testing.T) {
	errStr := `template: mychart/templates/deploy.yaml:12: nil pointer evaluating something that has no at clause`
	_, _, _, ok := ParseError(errStr)
	if ok {
		t.Fatal("expected unfixable without at <.Values.x> clause, got parseable")
	}
}

func TestParseError_NilPointerWithPipeFilter(t *testing.T) {
	errStr := `template: mychart/templates/deploy.yaml:5:8: executing "mychart/templates/deploy.yaml" at <.Values.foo | quote>: nil pointer evaluating interface {}.foo`
	kind, path, _, ok := ParseError(errStr)
	if !ok {
		t.Fatal("expected parseable, got false")
	}
	if kind != KindNilPointer {
		t.Errorf("kind: got %q", kind)
	}
	if path != "foo" {
		t.Errorf("path: got %q, want %q", path, "foo")
	}
}
