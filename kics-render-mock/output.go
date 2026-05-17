package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type ValuesOptionsSummary struct {
	Values       []string `json:"values"`
	ValueFiles   []string `json:"valueFiles"`
	StringValues []string `json:"stringValues,omitempty"`
}

type SplitManifestEntry struct {
	Path            string `json:"path"`
	Content         string `json:"content"`
	OriginalContent string `json:"originalContent"`
	SplitID         string `json:"splitID,omitempty"`
}

type RunResult struct {
	Error           *string              `json:"error"`
	RawManifest     *string              `json:"rawManifest"`
	PartialManifest *string              `json:"partialManifest"`
	SplitManifests  []SplitManifestEntry `json:"splitManifests"`
}

type DebugRunResult struct {
	Error           *string              `json:"error"`
	DebugLogs       []string             `json:"debugLogs"`
	RawManifest     *string              `json:"rawManifest"`
	PartialManifest *string              `json:"partialManifest"`
	SplitManifests  []SplitManifestEntry `json:"splitManifests"`
}

type RenderEntry struct {
	Toggle        string               `json:"toggle,omitempty"`
	DataType      string               `json:"dataType,omitempty"`
	CaseName      string               `json:"caseName,omitempty"`
	ValuesOptions ValuesOptionsSummary `json:"valuesOptions"`
	Standard      RunResult            `json:"standard"`
	Debug         DebugRunResult       `json:"debug"`
}

type RenderOutput struct {
	Suite      string        `json:"suite"`
	TestNumber int           `json:"testNumber"`
	TestName   string        `json:"testName"`
	ChartPath  string        `json:"chartPath"`
	Renders    []RenderEntry `json:"renders"`
}

func writeOutput(testDir string, out RenderOutput) error {
	data, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(testDir, "kics-render-output.json"), data, 0644)
}

type FixStep struct {
	Attempt       int    `json:"attempt"`
	ErrorSeen     string `json:"errorSeen"`
	Kind          string `json:"kind"`
	ValuePath     string `json:"valuePath"`
	ValueInjected string `json:"valueInjected"`
}

type FixedRenderEntry struct {
	RenderEntry
	Resolved    bool       `json:"resolved"`
	StopReason  string     `json:"stopReason"`
	FixChain    []FixStep  `json:"fixChain"`
	FixedResult *RunResult `json:"fixedResult,omitempty"`
}

type FixedRenderOutput struct {
	Suite      string             `json:"suite"`
	TestNumber int                `json:"testNumber"`
	TestName   string             `json:"testName"`
	ChartPath  string             `json:"chartPath"`
	Renders    []FixedRenderEntry `json:"renders"`
}

func writeFixedOutput(testDir string, out FixedRenderOutput) error {
	data, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(testDir, "kics-fixed-output.json"), data, 0644)
}
