package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/MabsIPCA/helm-tests/taxonomy_analyzer/analyzer"
	"github.com/MabsIPCA/helm-tests/taxonomy_analyzer/exporter"
	"github.com/MabsIPCA/helm-tests/taxonomy_analyzer/loader"
	"github.com/MabsIPCA/helm-tests/taxonomy_analyzer/model"
)

func main() {
	defaultCatalog := filepath.Join("..", "helm_fetcher", "backup", "run_002_274", "catalog_by_project.json")

	catalogPath := flag.String("input", defaultCatalog, "Path to catalog_by_project.json")
	outputDir := flag.String("out", "out", "Output directory for taxonomy reports")
	flag.Parse()

	absCatalog, err := filepath.Abs(*catalogPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to resolve input path: %v\n", err)
		os.Exit(1)
	}

	collector := analyzer.New(absCatalog)
	err = loader.StreamCatalog(absCatalog, func(repo model.RepoResult) error {
		collector.ConsumeRepo(repo)
		return nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to scan catalog: %v\n", err)
		os.Exit(1)
	}

	report := collector.Report()
	if err := exporter.WriteAll(*outputDir, report); err != nil {
		fmt.Fprintf(os.Stderr, "failed to export report: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Taxonomy analysis complete. Output directory: %s\n", *outputDir)
	fmt.Printf("Classified errors: %d | Unclassified errors: %d\n", report.Totals.ClassifiedErrors, report.Totals.UnclassifiedError)
}
