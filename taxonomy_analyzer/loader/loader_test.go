package loader_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/MabsIPCA/helm-tests/taxonomy_analyzer/loader"
)

const fixedSample = `[
  {
    "repo_url": "https://github.com/test/repo-a",
    "repo_name": "test/repo-a",
    "cloned_dir": "/tmp/repo-a",
    "total_charts": 1, "total_runs": 2, "total_successes": 0, "total_failures": 2, "total_dep_failures": 0,
    "charts": [
      {
        "chart_path": "/tmp/repo-a/chart",
        "total_runs": 2, "successes": 0, "failures": 2,
        "runs": [
          {
            "chart_path": "cloned/repo-a/chart",
            "helm_command": "helm template test cloned/repo-a/chart",
            "success": false,
            "error_message": "nil pointer evaluating interface {}.type",
            "fixed": {
              "resolved": true,
              "stop_reason": "",
              "fix_chain": [{"attempt": 1, "error_seen": "nil pointer evaluating interface {}.type", "kind": "nil_pointer", "value_path": "service.type", "value_injected": "ClusterIP"}]
            }
          },
          {
            "chart_path": "cloned/repo-a/chart",
            "helm_command": "helm template test cloned/repo-a/chart -f values.yaml",
            "success": false,
            "error_message": "nil pointer evaluating interface {}.repository",
            "fixed": {"resolved": false, "stop_reason": "unfixable_error_kind", "fix_chain": []}
          }
        ]
      }
    ],
    "kept": true, "dep_failed": false
  }
]`

func TestLoadFixedIndex_IndexesByHelmCommand(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "catalog_fixed.json")
	if err := os.WriteFile(path, []byte(fixedSample), 0o644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}

	idx, err := loader.LoadFixedIndex(path)
	if err != nil {
		t.Fatalf("LoadFixedIndex: %v", err)
	}

	if len(idx) != 2 {
		t.Errorf("index length: got %d, want 2", len(idx))
	}

	cmd1 := "helm template test cloned/repo-a/chart"
	if idx[cmd1] == nil {
		t.Fatalf("missing key %q", cmd1)
	}
	if !idx[cmd1].Resolved {
		t.Errorf("%q: Resolved got false, want true", cmd1)
	}
	if len(idx[cmd1].FixChain) != 1 {
		t.Errorf("%q: FixChain len got %d, want 1", cmd1, len(idx[cmd1].FixChain))
	}
	if idx[cmd1].FixChain[0].Kind != "nil_pointer" {
		t.Errorf("%q: FixChain[0].Kind got %q, want %q", cmd1, idx[cmd1].FixChain[0].Kind, "nil_pointer")
	}

	cmd2 := "helm template test cloned/repo-a/chart -f values.yaml"
	if idx[cmd2] == nil {
		t.Fatalf("missing key %q", cmd2)
	}
	if idx[cmd2].Resolved {
		t.Errorf("%q: Resolved got true, want false", cmd2)
	}
	if idx[cmd2].StopReason != "unfixable_error_kind" {
		t.Errorf("%q: StopReason got %q, want %q", cmd2, idx[cmd2].StopReason, "unfixable_error_kind")
	}
}

func TestLoadFixedIndex_MissingFile_ReturnsError(t *testing.T) {
	_, err := loader.LoadFixedIndex("/nonexistent/path.json")
	if err == nil {
		t.Error("expected error for missing file, got nil")
	}
}
