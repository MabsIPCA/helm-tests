
# ---------------------------------------------------------------------------
# helm_fetcher – Makefile
# ---------------------------------------------------------------------------

# Number of top entries to process from the selected source
TOP       ?= 400
# Number of entries fetched per API request/page (artifacthub max 60, github max 100)
PAGE_SIZE ?= 60
# Discovery source: artifacthub or github
SOURCE    ?= github
# In github-search-json mode, number of repos to fetch
SEARCH_TOP ?= 1000
# In github-search-json mode, star ordering: desc or asc
SEARCH_ORDER ?= desc
# In full github source mode, input JSON produced by run-github-search-json
SEARCH_IN ?= github_search.json
# In github-search-json mode, output JSON file path
SEARCH_OUT ?= github_search.json

# Output files
ALL_JSON  := catalog_all_results.json
FAIL_CSV  := catalog_failures.csv
SEARCH_JSON := github_search.json

# Cloned repos directory
CLONE_DIR := cloned

.PHONY: run run-github-search-json clean-cloned clean-output clean help

## run: fetch repos from selected source, run helm template on all combinations
run:
	go run . \
		-mode=full \
		-source=$(SOURCE) \
		-top=$(TOP) \
		-page-size=$(PAGE_SIZE) \
		-search-in=$(SEARCH_IN)

## run-github-search-json: generate JSON only from GitHub search (no clone/template)
run-github-search-json:
	go run . \
		-mode=github-search-json \
		-page-size=$(PAGE_SIZE) \
		-search-top=$(SEARCH_TOP) \
		-search-order=$(SEARCH_ORDER) \
		-search-out=$(SEARCH_OUT)

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
	@if exist $(SEARCH_JSON) del /q $(SEARCH_JSON)
	@echo "Done."

## clean: remove cloned repos AND output files
clean: clean-cloned clean-output

## help: show available targets
help:
	@echo.
	@echo   Usage: make [target] [SOURCE=github] [TOP=400] [PAGE_SIZE=60]
	@echo.
	@echo   Targets:
	@echo     run                    - run full catalog flow (search, clone, helm template)
	@echo     run-github-search-json - generate GitHub search JSON only
	@echo     clean-cloned  - delete the cloned/ directory
	@echo     clean-output  - delete catalog_all_results.json, catalog_failures.csv and github_search.json
	@echo     clean         - clean-cloned + clean-output
	@echo     help          - show this message
	@echo.
	@echo   Variables:
	@echo     SOURCE       (default: github) discovery source for run: artifacthub or github
	@echo     TOP          (default: 400) number of top entries to process in run mode
	@echo     PAGE_SIZE    (default: 60) entries per API request/page (artifacthub max 60, github max 100)
	@echo     SEARCH_IN    (default: github_search.json) input JSON for SOURCE=github in run mode
	@echo     SEARCH_TOP   (default: 1000) repos fetched in run-github-search-json mode
	@echo     SEARCH_ORDER (default: desc) star order in run-github-search-json mode: desc or asc
	@echo     SEARCH_OUT   (default: github_search.json) output JSON path in run-github-search-json mode
	@echo.

