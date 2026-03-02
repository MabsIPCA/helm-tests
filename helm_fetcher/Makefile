
# ---------------------------------------------------------------------------
# helm_fetcher – Makefile
# ---------------------------------------------------------------------------

# Number of results per page (max 100 for GitHub code search)
PER_PAGE  ?= 30
# Number of pages to fetch  (≈ PER_PAGE × MAX_PAGES unique repos)
MAX_PAGES ?= 3

# Output files
ALL_JSON  := catalog_all_results.json
FAIL_CSV  := catalog_failures.csv

# Cloned repos directory
CLONE_DIR := cloned

.PHONY: run clean-cloned clean-output clean help

## run: fetch repos from GitHub, run helm template on all combinations
run:
	go run helm_chart_render_catalog.go \
		-per-page=$(PER_PAGE) \
		-max-pages=$(MAX_PAGES)

## clean-cloned: remove all cloned repositories
clean-cloned:
	@echo "Removing $(CLONE_DIR)/ ..."
	@if exist $(CLONE_DIR) rmdir /s /q $(CLONE_DIR)
	@echo "Done."

## clean-output: remove generated JSON and CSV output files
clean-output:
	@echo "Removing output files ..."
	@if exist $(ALL_JSON) del /q $(ALL_JSON)
	@if exist $(FAIL_CSV)  del /q $(FAIL_CSV)
	@echo "Done."

## clean: remove cloned repos AND output files
clean: clean-cloned clean-output

## help: show available targets
help:
	@echo.
	@echo   Usage: make [target] [PER_PAGE=30] [MAX_PAGES=3]
	@echo.
	@echo   Targets:
	@echo     run           - run the catalog script
	@echo     clean-cloned  - delete the cloned/ directory
	@echo     clean-output  - delete catalog_all_results.json and catalog_failures.csv
	@echo     clean         - clean-cloned + clean-output
	@echo     help          - show this message
	@echo.
	@echo   Variables:
	@echo     PER_PAGE  (default: 30)  results per GitHub search page
	@echo     MAX_PAGES (default:  3)  number of pages to fetch
	@echo.

