package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfig_RenderProblems(t *testing.T) {
	dir := t.TempDir()
	raw := `{
		"testNumber": 2,
		"testName": "Nil Pointer",
		"chartName": "test-nil-pointer",
		"testFunctions": [
			{"name": "testNilMapAccess", "func": "nil-map-access"},
			{"name": "testNilNestedMap", "func": "nil-nested-map"}
		],
		"dataTypes": ["nilMap", "nilNested"]
	}`
	if err := os.WriteFile(filepath.Join(dir, "config.json"), []byte(raw), 0644); err != nil {
		t.Fatal(err)
	}

	cfg, err := loadConfig(dir)
	if err != nil {
		t.Fatalf("loadConfig: %v", err)
	}
	if cfg.TestNumber != 2 {
		t.Errorf("TestNumber: got %d", cfg.TestNumber)
	}
	if len(cfg.TestFunctions) != 2 {
		t.Errorf("TestFunctions len: got %d", len(cfg.TestFunctions))
	}
	if cfg.TestFunctions[0].Name != "testNilMapAccess" {
		t.Errorf("TestFunctions[0].Name: got %q", cfg.TestFunctions[0].Name)
	}
	if len(cfg.DataTypes) != 2 {
		t.Errorf("DataTypes len: got %d", len(cfg.DataTypes))
	}
}

func TestLoadConfig_OverrideTests(t *testing.T) {
	dir := t.TempDir()
	raw := `{
		"testNumber": 1,
		"testName": "Basic Data Types",
		"chartName": "test-basic",
		"testCases": [
			{"name": "default-values", "valuesFiles": [], "expected": {}},
			{"name": "string-override", "valuesFiles": ["values-string.yaml"], "expected": {}}
		]
	}`
	if err := os.WriteFile(filepath.Join(dir, "config.json"), []byte(raw), 0644); err != nil {
		t.Fatal(err)
	}

	cfg, err := loadConfig(dir)
	if err != nil {
		t.Fatalf("loadConfig: %v", err)
	}
	if len(cfg.TestCases) != 2 {
		t.Errorf("TestCases len: got %d", len(cfg.TestCases))
	}
	if cfg.TestCases[1].ValuesFiles[0] != "values-string.yaml" {
		t.Errorf("TestCases[1].ValuesFiles[0]: got %q", cfg.TestCases[1].ValuesFiles[0])
	}
}

func TestBuildInvocations_RenderProblems(t *testing.T) {
	cfg := &testConfig{
		TestFunctions: []testFunction{
			{Name: "testNilMapAccess", Func: "nil-map-access"},
		},
		DataTypes: []string{"nilMap", "nilNested"},
	}
	invs := buildInvocations(cfg, "/test/dir")
	if len(invs) != 2 {
		t.Fatalf("expected 2 invocations, got %d", len(invs))
	}
	if invs[0].toggle != "nil-map-access" {
		t.Errorf("toggle: got %q", invs[0].toggle)
	}
	if invs[0].dataType != "nilMap" {
		t.Errorf("dataType: got %q", invs[0].dataType)
	}
	if len(invs[0].valOpts.Values) != 2 {
		t.Fatalf("Values len: got %d", len(invs[0].valOpts.Values))
	}
	if invs[0].valOpts.Values[0] != "testNilMapAccess=true" {
		t.Errorf("Values[0]: got %q", invs[0].valOpts.Values[0])
	}
	if invs[0].valOpts.Values[1] != "testDataType=nilMap" {
		t.Errorf("Values[1]: got %q", invs[0].valOpts.Values[1])
	}
}

func TestBuildInvocations_OverrideTests(t *testing.T) {
	cfg := &testConfig{
		TestCases: []testCase{
			{Name: "string-override", ValuesFiles: []string{"values-string.yaml"}},
		},
	}
	invs := buildInvocations(cfg, "/test/dir")
	if len(invs) != 1 {
		t.Fatalf("expected 1 invocation, got %d", len(invs))
	}
	if invs[0].caseName != "string-override" {
		t.Errorf("caseName: got %q", invs[0].caseName)
	}
	if len(invs[0].valOpts.ValueFiles) != 1 {
		t.Fatalf("ValueFiles len: got %d", len(invs[0].valOpts.ValueFiles))
	}
	expected := filepath.Join("/test/dir", "values-string.yaml")
	if invs[0].valOpts.ValueFiles[0] != expected {
		t.Errorf("ValueFiles[0]: got %q, want %q", invs[0].valOpts.ValueFiles[0], expected)
	}
}

func TestBuildInvocations_DefaultDataType(t *testing.T) {
	cfg := &testConfig{
		TestFunctions: []testFunction{
			{Name: "testFoo", Func: "foo"},
		},
		DataTypes: []string{}, // empty → default
	}
	invs := buildInvocations(cfg, "/test/dir")
	if len(invs) != 1 {
		t.Fatalf("expected 1 invocation, got %d", len(invs))
	}
	// When dataType is "default", testDataType is NOT added to Values
	if len(invs[0].valOpts.Values) != 1 {
		t.Errorf("Values len for default dataType: got %d, want 1", len(invs[0].valOpts.Values))
	}
}
