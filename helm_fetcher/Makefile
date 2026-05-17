
# ---------------------------------------------------------------------------
# helm_fetcher – Makefile
# ---------------------------------------------------------------------------

# Number of top entries to process from the selected source
TOP       ?= 500
# Number of entries fetched per API request/page (artifacthub max 60, github max 100)
PAGE_SIZE ?= 60
# Discovery source: artifacthub or github
SOURCE    ?= artifacthub
# In github-search-json mode, number of repos to fetch
SEARCH_TOP ?= 1000
# In github-search-json mode, star ordering: desc or asc
SEARCH_ORDER ?= desc
# In full github source mode, input JSON produced by run-github-search-json
SEARCH_IN ?= artifacthub_search.json
# In github-search-json mode, output JSON file path
SEARCH_OUT ?= github_search.json
# In artifacthub-search-json mode, output JSON file path
ARTIFACT_SEARCH_OUT ?= artifacthub_search.json

# Plot generation settings
PLOT_DIR ?= plots
PLOT_TOP_PROBLEMS ?= 40
PLOT_TOP_STARS ?= 50
PLOT_TOP_POPULARITY ?= 200

# Output files
ALL_JSON  := catalog_all_results.json
FAIL_CSV  := catalog_failures.csv
SEARCH_JSON := github_search.json

# Cloned repos directory (override per run if needed)
CLONE_DIR ?= E:/helm_clones_artifacthub
# In fixer mode, the catalog JSON to read (produced by a previous 'run' invocation)
CATALOG_IN ?= catalog_kept.json

.PHONY: run run-github-search-json run-artifacthub-search-json run-fixer plot-problems plot-popularity plot-problems-popularity plot-all clean-cloned clean-output clean help

## run: fetch repos from selected source, run helm template on all combinations
run:
	go run . \
		-mode=full \
		-source=$(SOURCE) \
		-top=$(TOP) \
		-page-size=$(PAGE_SIZE) \
		-search-in=$(SEARCH_IN) \
		-clone-dir=$(CLONE_DIR)

## run-github-search-json: generate JSON only from GitHub search (no clone/template)
run-github-search-json:
	go run . \
		-mode=github-search-json \
		-page-size=$(PAGE_SIZE) \
		-search-top=$(SEARCH_TOP) \
		-search-order=$(SEARCH_ORDER) \
		-search-out=$(SEARCH_OUT)

## run-artifacthub-search-json: generate JSON only from Artifact Hub search (no clone/template)
run-artifacthub-search-json:
	go run . \
		-mode=artifacthub-search-json \
		-page-size=$(PAGE_SIZE) \
		-search-top=$(SEARCH_TOP) \
		-search-out=$(ARTIFACT_SEARCH_OUT)

## run-fixer: read catalog_kept.json, clone repos, run fix loop, write catalog_fixed.json
run-fixer:
	go run . \
		-mode=fixer \
		-catalog-in=$(CATALOG_IN) \
		-clone-dir=$(CLONE_DIR)

## plot-problems: generate SVG chart with problems per project
plot-problems:
	python scripts/plot_problems_per_project.py \
		--catalog-in=catalog_by_project.json \
		--output=$(PLOT_DIR)/problems_per_project.svg \
		--top=$(PLOT_TOP_PROBLEMS)

## plot-popularity: generate SVG chart with repo popularity (stars)
plot-popularity:
	python scripts/plot_repo_popularity.py \
		--search-in=$(SEARCH_IN) \
		--output=$(PLOT_DIR)/repo_popularity_stars.svg \
		--top=$(PLOT_TOP_STARS)

## plot-problems-popularity: generate SVG scatter chart with problems vs stars
plot-problems-popularity:
	python scripts/plot_problems_per_popularity.py \
		--catalog-in=catalog_by_project.json \
		--search-in=$(SEARCH_IN) \
		--output=$(PLOT_DIR)/problems_per_popularity.svg \
		--top=$(PLOT_TOP_POPULARITY)

## plot-all: generate both plots
plot-all: plot-problems plot-popularity plot-problems-popularity

## clean-cloned: remove all cloned repositories
clean-cloned:
	@echo "Removing $(CLONE_DIR)/ ..."
	@if exist "$(CLONE_DIR)" rmdir /s /q "$(CLONE_DIR)"
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
	@echo   Usage: make [target] [SOURCE=github] [TOP=400] [PAGE_SIZE=60] [CLONE_DIR=E:/helm_clones]
	@echo.
	@echo   Targets:
	@echo     run                    - run full catalog flow (load search JSON, clone, helm template)
	@echo     run-github-search-json - generate GitHub search JSON only
	@echo     run-artifacthub-search-json - generate Artifact Hub search JSON only
	@echo     run-fixer              - read catalog and apply fix loop (writes catalog_fixed.json)
	@echo     plot-problems          - generate plots/problems_per_project.svg
	@echo     plot-popularity        - generate plots/repo_popularity_stars.svg
	@echo     plot-problems-popularity - generate plots/problems_per_popularity.svg
	@echo     plot-all               - generate all plot files
	@echo     clean-cloned  - delete the configured clone directory
	@echo     clean-output  - delete catalog_all_results.json, catalog_failures.csv and github_search.json
	@echo     clean         - clean-cloned + clean-output
	@echo     help          - show this message
	@echo.
	@echo   Variables:
	@echo     SOURCE       (default: github) discovery source for run: artifacthub or github
	@echo     TOP          (default: 400) number of top entries to process in run mode
	@echo     PAGE_SIZE    (default: 60) entries per API request/page (artifacthub max 60, github max 100)
	@echo     SEARCH_IN    (default: artifacthub_search.json) input JSON for full mode (artifacthub or github)
	@echo     SEARCH_TOP   (default: 1000) repos fetched in run-github-search-json mode
	@echo     SEARCH_ORDER (default: desc) star order in run-github-search-json mode: desc or asc
	@echo     SEARCH_OUT   (default: github_search.json) output JSON path in run-github-search-json mode
	@echo     ARTIFACT_SEARCH_OUT (default: artifacthub_search.json) output JSON path in run-artifacthub-search-json mode
	@echo     CATALOG_IN   (default: catalog_kept.json) input catalog for run-fixer mode
	@echo     CLONE_DIR    (default: E:/helm_clones) clone destination for run/clean-cloned
	@echo     PLOT_DIR     (default: plots) directory where chart SVG files are written
	@echo     PLOT_TOP_PROBLEMS (default: 40) bars shown in problems-per-project chart
	@echo     PLOT_TOP_STARS    (default: 50) bars shown in repo-popularity chart
	@echo     PLOT_TOP_POPULARITY (default: 200) points shown in problems-vs-popularity chart
	@echo.

