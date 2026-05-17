package model_test

import (
	"encoding/json"
	"testing"

	"github.com/MabsIPCA/helm-tests/helm_fetcher/model"
)

func TestFixStep_JSONRoundtrip(t *testing.T) {
	step := model.FixStep{
		Attempt:       1,
		ErrorSeen:     "Error: nil pointer evaluating interface {}.foo",
		Kind:          "nil_pointer",
		ValuePath:     "service.type",
		ValueInjected: "",
	}
	data, err := json.Marshal(step)
	if err != nil {
		t.Fatal(err)
	}
	var got model.FixStep
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}
	if got != step {
		t.Errorf("roundtrip: got %+v, want %+v", got, step)
	}
}

func TestFixedRunResult_JSONKeys(t *testing.T) {
	r := model.FixedRunResult{
		Resolved:     true,
		StopReason:   "",
		FixChain:     []model.FixStep{},
		FinalCommand: "helm template test /chart --set foo=bar",
	}
	data, err := json.Marshal(r)
	if err != nil {
		t.Fatal(err)
	}
	m := map[string]interface{}{}
	json.Unmarshal(data, &m)
	for _, key := range []string{"resolved", "stop_reason", "fix_chain", "final_command"} {
		if _, ok := m[key]; !ok {
			t.Errorf("missing JSON key %q", key)
		}
	}
}

func TestRunResultWithFix_EmbedPromotesFields(t *testing.T) {
	orig := model.RunResult{
		ChartPath:    "/some/chart",
		HelmCommand:  "helm template test /some/chart",
		Success:      false,
		ErrorMessage: "Error: nil pointer",
	}
	withFix := model.RunResultWithFix{RunResult: orig}
	data, err := json.Marshal(withFix)
	if err != nil {
		t.Fatal(err)
	}
	m := map[string]interface{}{}
	json.Unmarshal(data, &m)
	for _, key := range []string{"chart_path", "helm_command", "success", "error_message"} {
		if _, ok := m[key]; !ok {
			t.Errorf("embedded RunResult field %q not promoted to JSON", key)
		}
	}
	if _, ok := m["fixed"]; ok {
		t.Error("fixed field should be omitted when nil")
	}
}
