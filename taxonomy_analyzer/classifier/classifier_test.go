package classifier_test

import (
	"testing"

	"github.com/MabsIPCA/helm-tests/taxonomy_analyzer/classifier"
)

func TestClassify_EncryptedValuesBeatsMalformedYAML(t *testing.T) {
	// A git-crypt encrypted value file trips "control characters are not allowed"
	// nested inside the generic "error converting yaml to json". It must classify
	// as encrypted_values (checked first), NOT malformed_yaml.
	msg := `Error: failed to parse /tmp/x/ncats-dev-values.yaml: cannot unmarshal yaml document: ` +
		`error converting YAML to JSON: yaml: control characters are not allowed`

	for _, src := range []string{"template", "dependency"} {
		res := classifier.Classify(msg, src)
		if !res.Classified {
			t.Fatalf("[%s] expected classified", src)
		}
		if res.SubKind != "encrypted_values" {
			t.Errorf("[%s] sub_kind: got %q, want encrypted_values", src, res.SubKind)
		}
	}
}

func TestClassify_AuthorAssertionSignature(t *testing.T) {
	// Helm's `required`/`fail` signature is "execution error at (...): <msg>".
	// Bespoke author messages that match no specific vocabulary must still be
	// caught as author_assertion, not left unclassified.
	cases := []string{
		"Error: execution error at (mychart/templates/tls.yaml:3:5): TLS is not configured. Set one of:\n  tls.existingSecret.name=<secret-name>\n  tls.generateJob.enabled=true",
		"Error: execution error at (app/templates/_helpers.tpl:12:1): You have to deploy monitoring.coreos.com/v1 first",
		"Error: execution error at (vm/templates/license.yaml:1:1): Pass valid license at .Values.license or .Values.global.license",
	}
	for _, msg := range cases {
		res := classifier.Classify(msg, "template")
		if res.SubKind != "author_assertion" {
			t.Errorf("sub_kind: got %q, want author_assertion for %q", res.SubKind, msg)
		}
	}
}

func TestClassify_SpecificRulesBeatAuthorAssertion(t *testing.T) {
	// An exec-error whose message DOES match a specific vocabulary keeps the more
	// descriptive bucket (required_value / custom_validation), not the catch-all.
	req := classifier.Classify("Error: execution error at (c/templates/x.yaml:2:3): database password is required", "template")
	if req.SubKind != "required_value" {
		t.Errorf("required: got %q, want required_value", req.SubKind)
	}
	cv := classifier.Classify("Error: execution error at (c/templates/x.yaml:2:3): set at least one of foo or bar", "template")
	if cv.SubKind != "custom_validation" {
		t.Errorf("custom: got %q, want custom_validation", cv.SubKind)
	}
}

func TestClassify_NilPointerNotAuthorAssertion(t *testing.T) {
	// Helm-internal evaluation errors use "executing ... at <expr>:" — they must
	// NOT be swept into author_assertion.
	msg := `Error: template: c/templates/x.yaml:5:10: executing "c/templates/x.yaml" at <.Values.foo.bar>: nil pointer evaluating interface {}.bar`
	res := classifier.Classify(msg, "template")
	if res.SubKind != "nil_pointer" {
		t.Errorf("sub_kind: got %q, want nil_pointer", res.SubKind)
	}
}

func TestClassify_GenuineMalformedYAMLStillMalformed(t *testing.T) {
	// A real YAML mistake (no control characters) must stay malformed_yaml.
	msg := "Error: failed to parse values.yaml: error converting YAML to JSON: yaml: line 4: did not find expected node content"
	res := classifier.Classify(msg, "template")
	if res.SubKind != "malformed_yaml" {
		t.Errorf("sub_kind: got %q, want malformed_yaml", res.SubKind)
	}
}
