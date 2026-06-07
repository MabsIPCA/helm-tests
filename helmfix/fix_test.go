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
		{"bare required values path", `Error: execution error at (jupyter/templates/jupyter.yaml:27:7): .Values.global.source_cidr required!`, "global.source_cidr"},
		{"bare required with entry word", `Error: execution error at (aqua-enforcer/templates/enforcer-token-secret.yaml:14:13): A valid .Values.enforcerToken entry required!`, "enforcerToken"},
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

// TestParseError_RequiredValueForSubField covers the "required value for
// <path> either as .value ..." pattern, where the placeholder must be injected
// at <path>.value (the chart reads the .value sub-field).
func TestParseError_RequiredValueForSubField(t *testing.T) {
	cases := []struct {
		name   string
		errStr string
		want   string
	}{
		{"secrets key", `Error: execution error at (mcp-server-21st-dev-magic/templates/secrets.yaml:10:9): required value for secrets.TWENTY_FIRST_API_KEY either as .value or .valueFrom.name and .valueFrom.key`, "secrets.TWENTY_FIRST_API_KEY.value"},
		{"key with digits", `Error: execution error at (mcp-server-aws-s3/templates/secrets.yaml:10:9): required value for secrets.AWS_S3_BUCKET either as .value or .valueFrom.name and .valueFrom.key`, "secrets.AWS_S3_BUCKET.value"},
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

// TestParseError_MustBeSetAndQuoted covers the newer "<path> must be set" and
// quoted-dotted-path required shapes.
func TestParseError_MustBeSetAndQuoted(t *testing.T) {
	cases := []struct {
		name   string
		errStr string
		want   string
	}{
		{"must be set", `Error: execution error at (envoy-gateway/templates/oidc-security-policy.yaml:21:15): global.domain must be set`, "global.domain"},
		{"must be set deep", `Error: execution error at (couchdb/templates/statefulset.yaml:25:28): A value for couchdbConfig.couchdb.uuid must be set`, "couchdbConfig.couchdb.uuid"},
		{"quoted is required", `Error: execution error at (cloud-connector/templates/secret.yaml:20:23): value 'sysdig.secureAPIToken' is required, but is not set`, "sysdig.secureAPIToken"},
		{"hyphenated subchart alias via --set hint", `Error: execution error at (sync-test/charts/op-geth/templates/x.yaml:1:4): op-geth.secrets.nodeKey.value is required: --set op-geth.secrets.nodeKey.value=foo`, "op-geth.secrets.nodeKey.value"},
		{"lowercase values prefix", `Error: execution error at (dp-cluster-proxies/templates/spark-master.yml:62:13): A valid .values.loadBalancer entry required!`, "loadBalancer"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			kind, path, _, ok := ParseError(tc.errStr)
			if !ok {
				t.Fatalf("expected parseable, got false")
			}
			if kind != KindRequiredValue {
				t.Errorf("kind: got %q, want %q", kind, KindRequiredValue)
			}
			if path != tc.want {
				t.Errorf("path: got %q, want %q", path, tc.want)
			}
		})
	}
}

// TestParseError_KubeVersion covers picking the highest compatible Kubernetes
// version from a chart's kubeVersion constraint.
func TestParseError_KubeVersion(t *testing.T) {
	cases := []struct {
		name   string
		errStr string
		want   string
	}{
		{"upper bound range", `Error: chart requires kubeVersion: >= 1.30.0 < 1.34.0 which is incompatible with Kubernetes v1.36.0`, "1.33.0"},
		{"prerelease upper bound", `Error: chart requires kubeVersion: <= 1.22.5-x which is incompatible with Kubernetes v1.36.0`, "1.22.0"},
		{"exact version", `Error: chart requires kubeVersion: 1.20.0 which is incompatible with Kubernetes v1.36.0`, "1.20.0"},
		{"lower bound above default", `Error: chart requires kubeVersion: >= 1.40.0 which is incompatible with Kubernetes v1.36.0`, "1.40.0"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			kind, path, value, ok := ParseError(tc.errStr)
			if !ok {
				t.Fatalf("expected parseable, got false")
			}
			if kind != KindKubeVersion {
				t.Errorf("kind: got %q, want %q", kind, KindKubeVersion)
			}
			if path != "" {
				t.Errorf("path: got %q, want empty (kube-version is not a --set)", path)
			}
			if value != tc.want {
				t.Errorf("version: got %q, want %q", value, tc.want)
			}
		})
	}
}

// TestIsUnparseableYAML covers the malformed-YAML / unmarshal class that must
// be recognized as inherently non-fixable.
func TestIsUnparseableYAML(t *testing.T) {
	nonFixable := []string{
		`Error: failed to parse values_ovms.yaml: cannot unmarshal yaml document: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type map[string]interface {}`,
		`Error: cannot load Chart.yaml: error converting YAML to JSON: yaml: line 4: found character that cannot start any token`,
		`Error: cannot load values.yaml: error reading yaml document: invalid Yaml document separator`,
	}
	for _, e := range nonFixable {
		if !IsUnparseableYAML(e) {
			t.Errorf("expected unparseable-YAML for %q", e)
		}
		// And it must NOT be mistaken for a fixable error.
		if _, _, _, ok := ParseError(e); ok {
			t.Errorf("malformed YAML should be unfixable, got parseable: %q", e)
		}
	}

	fixable := `Error: execution error at (a/templates/x.yaml:1:2): agent.identifier is required!`
	if IsUnparseableYAML(fixable) {
		t.Errorf("a normal required error must not be flagged unparseable-YAML")
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

func TestIsUnsupportedBuiltin(t *testing.T) {
	cases := []struct {
		name    string
		errStr  string
		builtin bool
	}{
		{
			name:    "release_time_removed_in_helm3",
			errStr:  `Error: tensorflow-serving/templates/deployment.yaml:16:34: executing "tensorflow-serving/templates/deployment.yaml" at <.Release.Time.Seconds>: nil pointer evaluating interface {}.Seconds`,
			builtin: true,
		},
		{
			name:    "release_root_context_dollar",
			errStr:  `template: x/templates/a.yaml:1:1: executing "x" at <$.Release.Time>: nil pointer evaluating interface {}.Time`,
			builtin: true,
		},
		{
			name:    "capabilities_builtin",
			errStr:  `template: x/templates/a.yaml:1:1: executing "x" at <.Capabilities.KubeVersion.Minor>: nil pointer evaluating interface {}.Minor`,
			builtin: true,
		},
		{
			name:    "values_nil_pointer_is_fixable_not_builtin",
			errStr:  `template: x/templates/a.yaml:1:1: executing "x" at <.Values.ingress.hosts>: nil pointer evaluating interface {}.hosts`,
			builtin: false,
		},
		{
			name:    "release_name_without_nil_pointer_is_not_builtin_error",
			errStr:  `something mentioning .Release.Name but no nil pointer`,
			builtin: false,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsUnsupportedBuiltin(tc.errStr); got != tc.builtin {
				t.Errorf("IsUnsupportedBuiltin = %v, want %v", got, tc.builtin)
			}
		})
	}

	// A built-in nil pointer must NOT be reported as a fixable kind by ParseError.
	relErr := `Error: x/templates/a.yaml:16:34: executing "x" at <.Release.Time.Seconds>: nil pointer evaluating interface {}.Seconds`
	if _, _, _, ok := ParseError(relErr); ok {
		t.Error("ParseError reported a built-in nil pointer as fixable; want ok=false")
	}
}

func TestParseError_NilPointerRootContextDollar(t *testing.T) {
	// Templates that reference the root context use "$.Values..." — the "$" must
	// not defeat path extraction. Real case: argocd-httproute.yaml.
	errStr := `Error: argocd/templates/argocd-httproute.yaml:1:18: executing "argocd/templates/argocd-httproute.yaml" at <$.Values.global.security>: nil pointer evaluating interface {}.security`
	kind, path, value, ok := ParseError(errStr)
	if !ok {
		t.Fatal("expected parseable, got false")
	}
	if kind != KindNilPointer {
		t.Errorf("kind: got %q, want %q", kind, KindNilPointer)
	}
	if path != "global.security" {
		t.Errorf("path: got %q, want %q", path, "global.security")
	}
	if value != "" {
		t.Errorf("value: got %q, want empty string", value)
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
