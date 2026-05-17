package loader

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/MabsIPCA/helm-tests/taxonomy_analyzer/model"
)

// StreamCatalog decodes a large catalog array and invokes fn for each repository.
func StreamCatalog(catalogPath string, fn func(repo model.RepoResult) error) error {
	f, err := os.Open(catalogPath)
	if err != nil {
		return fmt.Errorf("open catalog: %w", err)
	}
	defer f.Close()

	decoder := json.NewDecoder(f)

	tok, err := decoder.Token()
	if err != nil {
		return fmt.Errorf("read opening token: %w", err)
	}
	delim, ok := tok.(json.Delim)
	if !ok || delim != '[' {
		return fmt.Errorf("invalid catalog format: expected JSON array")
	}

	for decoder.More() {
		var repo model.RepoResult
		if err := decoder.Decode(&repo); err != nil {
			return fmt.Errorf("decode repository entry: %w", err)
		}
		if err := fn(repo); err != nil {
			return err
		}
	}

	_, err = decoder.Token()
	if err != nil {
		return fmt.Errorf("read closing token: %w", err)
	}

	return nil
}

// LoadFixedIndex streams catalog_fixed.json and returns a map from
// helm_command -> *FixedResult for correlation with primary catalog runs.
func LoadFixedIndex(path string) (map[string]*model.FixedResult, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open fixed catalog: %w", err)
	}
	defer f.Close()

	decoder := json.NewDecoder(f)

	tok, err := decoder.Token()
	if err != nil {
		return nil, fmt.Errorf("read opening token: %w", err)
	}
	delim, ok := tok.(json.Delim)
	if !ok || delim != '[' {
		return nil, fmt.Errorf("invalid fixed catalog format: expected JSON array")
	}

	index := make(map[string]*model.FixedResult)
	for decoder.More() {
		var repo model.RepoResult
		if err := decoder.Decode(&repo); err != nil {
			return nil, fmt.Errorf("decode fixed repository entry: %w", err)
		}
		for _, chart := range repo.Charts {
			for _, run := range chart.Runs {
				if run.Fixed != nil && run.HelmCommand != "" {
					if _, exists := index[run.HelmCommand]; exists {
						fmt.Fprintf(os.Stderr, "warning: duplicate helm_command in fixed catalog, overwriting: %s\n", run.HelmCommand)
					}
					index[run.HelmCommand] = run.Fixed
				}
			}
		}
	}

	if _, err = decoder.Token(); err != nil {
		return nil, fmt.Errorf("read closing token: %w", err)
	}

	return index, nil
}
