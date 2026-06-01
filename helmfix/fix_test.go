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

// requiredCases covers value paths recovered from custom required/fail
// messages, drawn from real catalog failures.
func TestParseError_RequiredMessagePaths(t *testing.T) {
	cases := []struct {
		name   string
		errStr string
		want   string
	}{
		{"dotted is required", `Error: execution error at (legit-kubernetes-agent/templates/secret.yaml:7:17): agent.identifier is required!`, "agent.identifier"},
		{"deep dotted is required", `Error: execution error at (demeter-fabric/templates/sts.yaml:23:28): config.kafka.consumerCacheName is required`, "config.kafka.consumerCacheName"},
		{"leading dot dotted", `Error: execution error at (keycloak-backup/templates/cronjob.yaml:42:57): .backup.targetKeycloak is required!`, "backup.targetKeycloak"},
		{"leading dot single", `Error: execution error at (keycloak/templates/keycloak.yaml:13:23): .db is required!`, "db"},
		{"please define", `Error: execution error at (monitoring-umbrella/charts/loki/templates/validate.yaml:46:4): Please define loki.storage.bucketName.chunks`, "loki.storage.bucketName.chunks"},
		{"missing required value", `Error: execution error at (postgrest/templates/secret.yaml:8:8): Missing required value postgrest.dbUri`, "postgrest.dbUri"},
		{"set hint", `Error: execution error at (fume/templates/configmap.yaml:1:4): CANONICAL_BASE_URL is required. Set via --set configMap.CANONICAL_BASE_URL="https://fume.your-company.com"`, "configMap.CANONICAL_BASE_URL"},
		{"explicit values path", `Error: execution error at (powerdns-admin/templates/deployment.yaml:47:63): .Values.powerdnsAdmin.config.secretKey or .Values.powerdnsAdmin.config.secretKeySecret is required!`, "powerdnsAdmin.config.secretKey"},
		{"required when clause", `Error: execution error at (wireguard/templates/secret.yaml:2:11): wireguard.privateKey is required when wireguard.existingSecret is not set`, "wireguard.privateKey"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			kind, path, value, ok := ParseError(tc.errStr)
			if !ok {
				t.Fatalf("expected parseable, got false")
			}
			if kind != KindRequiredValue {
				t.Errorf("kind: got %q, want %q", kind, KindRequiredValue)
			}
			if path != tc.want {
				t.Errorf("path: got %q, want %q", path, tc.want)
			}
			if value != KicsPlaceholder {
				t.Errorf("value: got %q, want %q", value, KicsPlaceholder)
			}
		})
	}
}

// Required-style messages with no recoverable value path must stay unfixable.
func TestParseError_RequiredMessageNoPath(t *testing.T) {
	cases := []string{
		`Error: execution error at (amphitrite/templates/secret.yaml:11:27): A connection string to SQLServer is required`,
		`Error: execution error at (element-web/templates/deployment.yaml:18:28): Must specify a default homeserver`,
		`Error: execution error at (tenzu-back/templates/job.yaml:32:14): A host is mandatory for redis `,
	}
	for _, errStr := range cases {
		if _, _, _, ok := ParseError(errStr); ok {
			t.Errorf("expected unfixable (no path), got parseable for %q", errStr)
		}
	}
}

func TestParseError_WrongTypeExpectedString(t *testing.T) {
	errStr := `Error: template: acm-aws-cluster/templates/ssh-private-key-secret.yaml:7:19: executing "acm-aws-cluster/templates/ssh-private-key-secret.yaml" at <.Values.sshKey>: wrong type for value; expected string; got interface {}`
	kind, path, value, ok := ParseError(errStr)
	if !ok {
		t.Fatal("expected parseable, got false")
	}
	if kind != KindTypeMismatch {
		t.Errorf("kind: got %q, want %q", kind, KindTypeMismatch)
	}
	if path != "sshKey" {
		t.Errorf("path: got %q, want %q", path, "sshKey")
	}
	if value != KicsPlaceholder {
		t.Errorf("value: got %q, want %q", value, KicsPlaceholder)
	}
}

// b64enc-style type errors expose no value path and must remain unfixable.
func TestParseError_InvalidValueNoPath(t *testing.T) {
	errStr := `Error: template: lion/templates/django-secret.yaml:7:66: executing "lion/templates/django-secret.yaml" at <b64enc>: invalid value; expected string`
	if _, _, _, ok := ParseError(errStr); ok {
		t.Fatal("expected unfixable, got parseable")
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
