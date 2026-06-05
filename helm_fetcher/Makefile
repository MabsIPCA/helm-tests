
# ---------------------------------------------------------------------------
# helm_fetcher – Makefile
# ---------------------------------------------------------------------------

# Number of top entries to process from the selected source
TOP       ?= 500
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
CLONE_DIR ?= D:/helm_clones_artifacthub
# In fixer mode, the catalog JSON to read (produced by a previous 'run' invocation)
CATALOG_IN ?= catalog_kept.json
# Keep cloned repos even when they have no template failures (true = skip deletion, allows reuse on reruns)
KEEP_CLONES ?= true
# Per-run output subdirectory (auto-generated as runs/<timestamp>_<source> when empty)
RUN_DIR     ?=
# Cumulative catalog used by plot targets (rebuilt automatically after each run)
CATALOG_CUMULATIVE ?= catalog_cumulative.json

# refix-all: per-source input catalogs, clone dirs, and output fixed catalogs
GITHUB_CATALOG     ?= results/github/catalog_by_project.json
GITHUB_CLONE_DIR   ?= D:/helm_clones_github
ARTIFACT_CATALOG   ?= runs/20260529_135859_artifacthub/catalog_by_project.json
ARTIFACT_CLONE_DIR ?= D:/helm_clones_artifacthub
FIXED_GITHUB       ?= catalog_fixed_github_final.json
FIXED_ARTIFACT     ?= catalog_fixed_ah_final.json
FIXED_CUMULATIVE   ?= catalog_fixed_cumulative.json
# Taxonomy regen at the end of refix-all (input sources catalog + output dir)
SOURCES_CATALOG    ?= catalog_sources_merged.json
ANALYZE_OUT        ?= out/cumulative_final

## GH_USER: GitHub username used by setup-token
GH_USER ?=

.PHONY: run run-both run-github-search-json run-artifacthub-search-json get-github get-artifact run-fixer refix-all merge setup-token gen-token plot-problems plot-popularity plot-problems-popularity plot-all clean-cloned clean-output clean help

## run: fetch repos from selected source, run helm template on all combinations
run:
	go run . \
		-mode=full \
		-source=$(SOURCE) \
		-top=$(TOP) \
		-page-size=$(PAGE_SIZE) \
		-search-in=$(SEARCH_IN) \
		-clone-dir=$(CLONE_DIR) \
		-keep-clones=$(KEEP_CLONES) \
		-run-dir=$(RUN_DIR)

## run-both: run full catalog flow for both github and artifacthub sources (overnight)
run-both:
	powershell -ExecutionPolicy Bypass -File run_both_sources.ps1 -Top $(TOP) -GithubCloneDir D:/helm_clones_github -ArtifactCloneDir D:/helm_clones_artifacthub

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

## get-github: fetch top-500 GitHub repos and write github_search.json
get-github:
	go run . \
		-mode=github-search-json \
		-page-size=$(PAGE_SIZE) \
		-search-top=500 \
		-search-order=$(SEARCH_ORDER) \
		-search-out=$(SEARCH_OUT)

## get-artifact: fetch top-500 Artifact Hub repos and write artifacthub_search.json
get-artifact:
	go run . \
		-mode=artifacthub-search-json \
		-page-size=$(PAGE_SIZE) \
		-search-top=500 \
		-search-out=$(ARTIFACT_SEARCH_OUT)

## get-token: authenticate via GitHub CLI and save the token to .env
get-token:
	@powershell -ExecutionPolicy Bypass -File scripts/get_token.ps1

## setup-token: store a new GitHub PAT in .env (API) and Windows Credential Manager (git clone)
##   Usage: make setup-token GH_USER=<your-github-username> GITHUB_TOKEN=<ghp_...>
setup-token:
	@if "$(GITHUB_TOKEN)"=="" (echo ERROR: pass GITHUB_TOKEN=ghp_... & exit 1)
	@if "$(GH_USER)"=="" (echo ERROR: pass GH_USER=your-github-username & exit 1)
	@powershell -Command "$$lines = (Get-Content .env -ErrorAction SilentlyContinue) -notmatch '^GITHUB_TOKEN='; Set-Content .env -Value ($$lines + 'GITHUB_TOKEN=$(GITHUB_TOKEN)') -Encoding UTF8"
	@powershell -Command "$$c = \"protocol=https`nhost=github.com`nusername=$(GH_USER)`npassword=$(GITHUB_TOKEN)`n\"; $$c | git credential approve"
	@echo Token stored in .env and Windows Credential Manager.

## merge: rebuild catalog_cumulative.json from all existing runs/ subfolders (no new cloning)
merge:
	go run . -mode=merge

## run-fixer: read catalog_kept.json, clone repos, run fix loop, write catalog_fixed.json
run-fixer:
	go run . \
		-mode=fixer \
		-catalog-in=$(CATALOG_IN) \
		-clone-dir=$(CLONE_DIR)

## refix-all: re-run the fixer over BOTH sources, rebuild catalog_fixed_cumulative.json,
##   and regenerate the taxonomy report + viewer in one shot.
##   Chains: github fixer -> artifacthub fixer -> merge -> taxonomy analyze-cumulative.
##   Pass SKIP_ANALYZE=1 to stop after the merge. If you edited helmfix/fix.go, refresh
##   the vendored copies first (see scripts/refix_all.ps1 header).
refix-all:
	powershell -ExecutionPolicy Bypass -File scripts/refix_all.ps1 \
		-GithubCatalog $(GITHUB_CATALOG) -GithubCloneDir $(GITHUB_CLONE_DIR) \
		-ArtifactCatalog $(ARTIFACT_CATALOG) -ArtifactCloneDir $(ARTIFACT_CLONE_DIR) \
		-GithubFixed $(FIXED_GITHUB) -ArtifactFixed $(FIXED_ARTIFACT) -Cumulative $(FIXED_CUMULATIVE) \
		-SourcesCatalog $(SOURCES_CATALOG) -AnalyzeOut $(ANALYZE_OUT) $(if $(SKIP_ANALYZE),-SkipAnalyze,)

## plot-problems: generate SVG chart with problems per project
plot-problems:
	python scripts/plot_problems_per_project.py \
		--catalog-in=$(CATALOG_CUMULATIVE) \
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
		--catalog-in=$(CATALOG_CUMULATIVE) \
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
	@if exist catalog_fixed.json del /q catalog_fixed.json
	@if exist catalog_fixer_report.md del /q catalog_fixer_report.md
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
	@echo     run-both               - run full catalog flow for BOTH github and artifacthub (overnight)
	@echo     get-github             - fetch top-500 GitHub repos and write github_search.json
	@echo     get-artifact           - fetch top-500 Artifact Hub repos and write artifacthub_search.json
	@echo     run-github-search-json - generate GitHub search JSON only (uses SEARCH_TOP)
	@echo     run-artifacthub-search-json - generate Artifact Hub search JSON only (uses SEARCH_TOP)
	@echo     gen-token              - open GitHub token page, paste token, save to .env
	@echo     setup-token            - store GitHub PAT in .env and Windows Credential Manager
	@echo     merge                  - rebuild catalog_cumulative.json from all runs/ subfolders
	@echo     run-fixer              - read catalog and apply fix loop (writes catalog_fixed.json)
	@echo     refix-all              - re-run fixer for BOTH sources, rebuild fixed catalog, regen report (SKIP_ANALYZE=1 to stop after merge)
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
	@echo     TOP          (default: 500) number of top entries to process in run mode
	@echo     PAGE_SIZE    (default: 60) entries per API request/page (artifacthub max 60, github max 100)
	@echo     SEARCH_IN    (default: github_search.json) input JSON for full mode (artifacthub or github)
	@echo     SEARCH_TOP   (default: 1000) repos fetched in run-github-search-json mode
	@echo     SEARCH_ORDER (default: desc) star order in run-github-search-json mode: desc or asc
	@echo     SEARCH_OUT   (default: github_search.json) output JSON path in run-github-search-json mode
	@echo     ARTIFACT_SEARCH_OUT (default: artifacthub_search.json) output JSON path in run-artifacthub-search-json mode
	@echo     CATALOG_IN   (default: catalog_kept.json) input catalog for run-fixer mode
	@echo     CLONE_DIR    (default: D:/helm_clones_artifacthub) clone destination for run/clean-cloned
	@echo     KEEP_CLONES        (default: true) keep cloned repos with no failures so reruns skip re-cloning
	@echo     RUN_DIR            (default: auto) per-run output dir, auto-generated as runs/TIMESTAMP_SOURCE
	@echo     CATALOG_CUMULATIVE (default: catalog_cumulative.json) catalog used by plot targets
	@echo     GITHUB_CATALOG / ARTIFACT_CATALOG   (refix-all) per-source input catalogs
	@echo     GITHUB_CLONE_DIR / ARTIFACT_CLONE_DIR (refix-all) per-source clone dirs
	@echo     FIXED_CUMULATIVE   (default: catalog_fixed_cumulative.json) refix-all merged output
	@echo     PLOT_DIR     (default: plots) directory where chart SVG files are written
	@echo     PLOT_TOP_PROBLEMS (default: 40) bars shown in problems-per-project chart
	@echo     PLOT_TOP_STARS    (default: 50) bars shown in repo-popularity chart
	@echo     PLOT_TOP_POPULARITY (default: 200) points shown in problems-vs-popularity chart
	@echo.

