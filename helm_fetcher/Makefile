
# ---------------------------------------------------------------------------
# helm_fetcher – Makefile
# ---------------------------------------------------------------------------

# Number of top entries to process from the selected source
TOP       ?= 400
# Number of entries fetched per API request/page (artifacthub max 60, github max 100)
PAGE_SIZE ?= 60
# Discovery source: artifacthub or github
SOURCE    ?= artifacthub

# Output files
ALL_JSON  := catalog_all_results.json
FAIL_CSV  := catalog_failures.csv

# Cloned repos directory
CLONE_DIR := cloned

.PHONY: run clean-cloned clean-output clean help

## run: fetch repos from selected source, run helm template on all combinations
run:
	go run main.go \
		-source=$(SOURCE) \
		-top=$(TOP) \
		-page-size=$(PAGE_SIZE)

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
	@echo   Usage: make [target] [SOURCE=artifacthub] [TOP=400] [PAGE_SIZE=60]
	@echo.
	@echo   Targets:
	@echo     run           - run the catalog script
	@echo     clean-cloned  - delete the cloned/ directory
	@echo     clean-output  - delete catalog_all_results.json and catalog_failures.csv
	@echo     clean         - clean-cloned + clean-output
	@echo     help          - show this message
	@echo.
	@echo   Variables:
	@echo     SOURCE    (default: artifacthub) discovery source: artifacthub or github
	@echo     TOP       (default: 400) number of top entries to process
	@echo     PAGE_SIZE (default: 60) entries per API request/page (artifacthub max 60)
	@echo.

