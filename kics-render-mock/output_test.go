package main

import (
	"encoding/json"
	"testing"
)

func TestRenderOutputMarshal(t *testing.T) {
	errStr := "template error"
	manifest := "---\napiVersion: v1"
	out := RenderOutput{
		Suite:      "render_problems",
		TestNumber: 2,
		TestName:   "Nil Pointer",
		ChartPath:  "/some/path",
		Renders: []RenderEntry{
			{
				Toggle:   "nil-map-access",
				DataType: "nilMap",
				ValuesOptions: ValuesOptionsSummary{
					Values:     []string{"testNilMapAccess=true", "testDataType=nilMap"},
					ValueFiles: []string{},
				},
				Standard: RunResult{
					Error:           &errStr,
					RawManifest:     nil,
					PartialManifest: nil,
					SplitManifests:  []SplitManifestEntry{},
				},
				Debug: DebugRunResult{
					Error:           &errStr,
					DebugLogs:       []string{"[debug] chart path: /some/path"},
					RawManifest:     nil,
					PartialManifest: &manifest,
					SplitManifests:  []SplitManifestEntry{},
				},
			},
		},
	}

	data, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var roundTrip RenderOutput
	if err := json.Unmarshal(data, &roundTrip); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if roundTrip.Suite != "render_problems" {
		t.Errorf("Suite: got %q", roundTrip.Suite)
	}
	if len(roundTrip.Renders) != 1 {
		t.Fatalf("Renders len: got %d", len(roundTrip.Renders))
	}
	r := roundTrip.Renders[0]
	if r.Standard.Error == nil || *r.Standard.Error != errStr {
		t.Errorf("Standard.Error: got %v", r.Standard.Error)
	}
	if r.Debug.PartialManifest == nil || *r.Debug.PartialManifest != manifest {
		t.Errorf("Debug.PartialManifest: got %v", r.Debug.PartialManifest)
	}
	if len(r.Debug.DebugLogs) != 1 {
		t.Errorf("Debug.DebugLogs len: got %d", len(r.Debug.DebugLogs))
	}
}
