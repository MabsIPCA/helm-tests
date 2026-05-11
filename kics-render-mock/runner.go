package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"helm.sh/helm/v3/pkg/cli/values"
)

// testConfig is a unified shape that covers both suite config.json formats.
// Presence of TestFunctions → render_problems suite.
// Presence of TestCases    → override_tests_suite.
type testConfig struct {
	TestNumber    int            `json:"testNumber"`
	TestName      string         `json:"testName"`
	ChartName     string         `json:"chartName"`
	TestFunctions []testFunction `json:"testFunctions,omitempty"`
	DataTypes     []string       `json:"dataTypes,omitempty"`
	TestCases     []testCase     `json:"testCases,omitempty"`
}

type testFunction struct {
	Name string `json:"name"` // --set flag name, e.g. "testNilMapAccess"
	Func string `json:"func"` // display key,    e.g. "nil-map-access"
}

type testCase struct {
	Name           string   `json:"name"`
	ValuesFiles    []string `json:"valuesFiles"`
	SetFlags       []string `json:"setFlags,omitempty"`
	SetStringFlags []string `json:"setStringFlags,omitempty"`
}

// invocation is one render call to make.
type invocation struct {
	toggle   string // render_problems: display key (Func field)
	dataType string // render_problems: data type name
	caseName string // override_tests_suite: test case name
	valOpts  *values.Options
}

func loadConfig(testDir string) (*testConfig, error) {
	data, err := os.ReadFile(filepath.Join(testDir, "config.json"))
	if err != nil {
		return nil, fmt.Errorf("read config.json: %w", err)
	}
	var cfg testConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parse config.json: %w", err)
	}
	return &cfg, nil
}

func buildInvocations(cfg *testConfig, testDir string) []invocation {
	if len(cfg.TestFunctions) > 0 {
		return buildRenderProblemsInvocations(cfg)
	}
	return buildOverrideTestsInvocations(cfg, testDir)
}

func buildRenderProblemsInvocations(cfg *testConfig) []invocation {
	dataTypes := cfg.DataTypes
	if len(dataTypes) == 0 {
		dataTypes = []string{"default"}
	}
	var invs []invocation
	for _, tf := range cfg.TestFunctions {
		for _, dt := range dataTypes {
			vals := []string{fmt.Sprintf("%s=true", tf.Name)}
			if dt != "default" {
				vals = append(vals, fmt.Sprintf("testDataType=%s", dt))
			}
			invs = append(invs, invocation{
				toggle:   tf.Func,
				dataType: dt,
				valOpts:  &values.Options{Values: vals},
			})
		}
	}
	return invs
}

func buildOverrideTestsInvocations(cfg *testConfig, testDir string) []invocation {
	var invs []invocation
	for _, tc := range cfg.TestCases {
		absFiles := make([]string, len(tc.ValuesFiles))
		for i, vf := range tc.ValuesFiles {
			absFiles[i] = filepath.Join(testDir, vf)
		}
		invs = append(invs, invocation{
			caseName: tc.Name,
			valOpts: &values.Options{
				ValueFiles:   absFiles,
				Values:       tc.SetFlags,
				StringValues: tc.SetStringFlags,
			},
		})
	}
	return invs
}
