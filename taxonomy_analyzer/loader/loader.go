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
