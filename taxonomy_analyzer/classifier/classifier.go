package classifier

import "strings"

// Rule maps a set of message snippets to taxonomy labels.
type Rule struct {
	Kind     string
	SubKind  string
	Patterns []string
}

// Result is the classifier output.
type Result struct {
	Kind         string
	SubKind      string
	Classified   bool
	Normalized   string
	ErrorSource  string
	OriginalText string
}

// nonFixableYAMLPatterns identify a chart/values file that Helm cannot even
// parse (malformed YAML or a JSON unmarshal mismatch). These are inherently
// non-fixable by value injection and are bucketed as their own category.
var nonFixableYAMLPatterns = []string{
	"cannot unmarshal",
	"error unmarshaling json",
	"error converting yaml to json",
	"error reading yaml document",
}

var templateRules = []Rule{
	// malformed_yaml is checked first: the chart can't be parsed, so it is a
	// distinct non-fixable class (not a render/value problem).
	{Kind: "template", SubKind: "malformed_yaml", Patterns: nonFixableYAMLPatterns},
	// unsupported_builtin: nil pointer on a Helm built-in object (.Release.Time,
	// .Capabilities, ...) — not a value, so --set cannot fix it. Checked before
	// the generic nil_pointer rule so it isn't mistaken for a fixable case. The
	// classic case is .Release.Time, removed in Helm 3 (Helm-2-only charts).
	{Kind: "template", SubKind: "unsupported_builtin", Patterns: []string{
		"at <.release.", "at <$.release.",
		"at <.capabilities.", "at <$.capabilities.",
	}},
	{Kind: "template", SubKind: "nil_pointer", Patterns: []string{"nil pointer evaluating"}},
	{Kind: "template", SubKind: "parse_error", Patterns: []string{"parse error at", "unexpected \"=\" in operand", "unexpected end"}},
	{Kind: "template", SubKind: "missing_template", Patterns: []string{"no template", "error calling include"}},
	{Kind: "template", SubKind: "type_mismatch", Patterns: []string{"wrong type for value", "expected string", "expected int", "got interface"}},
	{Kind: "template", SubKind: "invalid_value", Patterns: []string{"invalid value; expected", "error calling index", "index of untyped nil", "proxy token must be a 32 byte random string", "proxy.secrettoken must be a 32 byte random string"}},
	{Kind: "template", SubKind: "dependency_check_failed", Patterns: []string{"an error occurred while checking for chart dependencies", "missing in charts/ directory"}},
	{Kind: "template", SubKind: "kube_version_incompatible", Patterns: []string{"chart requires kubeversion", "incompatible with kubernetes"}},
	{Kind: "template", SubKind: "values_merge_error", Patterns: []string{"values contain an error that may be a result of merging"}},
	{Kind: "template", SubKind: "values_schema_validation", Patterns: []string{"values don't meet the specifications of the schema"}},
	{Kind: "template", SubKind: "required_value", Patterns: []string{" is required", " is mandatory", "mandatory for", "please define ", "please specify ", "please provide ", "you did not specify", "must be set", "must be provided", "required value for", "either as .value", "a connection string to", "a host is mandatory", "must specify", "must provide", "need provide", "no value found for", "missing required value", "no root password and no existing secret was provided", "must be a long random string", "value or secretref must be provided", "expected non-empty", "expected the key", "expected ingress [main] to be enabled", "required!"}},
	{Kind: "template", SubKind: "library_chart_not_installable", Patterns: []string{"library charts are not installable"}},
	{Kind: "template", SubKind: "yaml_render", Patterns: []string{"yaml parse error", "did not find expected key"}},
	// custom_validation: a chart's hand-written guard (mutual exclusivity,
	// one-of, at-least-one, environment/version preconditions). Placed late so
	// the specific rules above win; catches bespoke "fail" assertions that would
	// otherwise be unclassified.
	{Kind: "template", SubKind: "custom_validation", Patterns: []string{
		" xor ", "either set", "either use", "either by", "either configure", "configure it either",
		"at least one", "one of ", "to be set together", "must be set together",
		"cannot be installed without", "cannot be used with", "must be installed",
		"you must ", "not specified", "please set ",
		"must be greater than", "must be less than", "must be defined", "duplicate port",
		// environment / cluster preconditions and discovery guards
		"crds installed", "crds first", "resource definition in the target cluster",
		"deployment found using label", "is missing", "value is missing", "undefined image",
	}},
	// chart_metadata: the chart package itself is invalid (bad apiVersion,
	// version, or dependency processing) — surfaces template-side too.
	{Kind: "template", SubKind: "chart_metadata", Patterns: []string{"invalid chart apiversion", "unsupported chart version", "chart dependencies processing failed"}},
	{Kind: "template", SubKind: "runtime_eval", Patterns: []string{"executing", "error calling"}},
}

var dependencyRules = []Rule{
	// A Chart.yaml / values.yaml that fails to load/parse is non-fixable, same
	// class as the template-side malformed_yaml. Checked first so the broad
	// "directory "/" not found" subchart patterns below don't shadow it.
	{Kind: "dependency", SubKind: "malformed_yaml", Patterns: nonFixableYAMLPatterns},
	{Kind: "dependency", SubKind: "missing_subchart", Patterns: []string{"an error occurred while checking for chart dependencies", "missing in charts/ directory", "directory ", " not found"}},
	{Kind: "dependency", SubKind: "missing_repository", Patterns: []string{"no repository definition for", "please add the missing repos", "please add them via 'helm repo add'"}},
	{Kind: "dependency", SubKind: "lock_file_out_of_sync", Patterns: []string{"lock file", "out of sync with the dependencies file"}},
	{Kind: "dependency", SubKind: "rate_limit", Patterns: []string{"toomanyrequests", "increase-rate-limit", "unauthenticated pull rate limit"}},
	{Kind: "dependency", SubKind: "network_dns", Patterns: []string{"no such host", "dial tcp", "lookup"}},
	{Kind: "dependency", SubKind: "cache_index_missing", Patterns: []string{"no cached repository", "index.yaml"}},
	{Kind: "dependency", SubKind: "repo_update", Patterns: []string{"getting updates for unmanaged helm repositories"}},
	{Kind: "dependency", SubKind: "unpack_error", Patterns: []string{"error unpacking subchart"}},
	{Kind: "dependency", SubKind: "version_resolution", Patterns: []string{"can't get a valid version for", "make sure a matching chart version exists"}},
	{Kind: "dependency", SubKind: "chart_validation", Patterns: []string{"validation: chart.metadata.name is required", "validation: chart.metadata.version", "is invalid", "invalid chart apiversion", "unsupported chart version", "improper constraint", "invalid version/constraint format", "larger than the maximum file size"}},
}

// Classify maps an error message to taxonomy labels.
func Classify(message string, source string) Result {
	normalized := strings.ToLower(strings.TrimSpace(message))
	result := Result{
		Kind:         "unknown",
		SubKind:      "unclassified",
		Classified:   false,
		Normalized:   normalized,
		ErrorSource:  source,
		OriginalText: message,
	}

	rules := templateRules
	if source == "dependency" {
		rules = dependencyRules
	}

	for _, rule := range rules {
		for _, pattern := range rule.Patterns {
			if strings.Contains(normalized, pattern) {
				result.Kind = rule.Kind
				result.SubKind = rule.SubKind
				result.Classified = true
				return result
			}
		}
	}

	return result
}
