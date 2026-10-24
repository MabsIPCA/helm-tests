# Helm chart testing and repair research

Tools and datasets for a master's thesis on Helm values overrides, chart
rendering failures, and automated repair. Small test suites sit alongside a
pipeline that discovers real-world charts, records failures, attempts repairs,
and classifies the results.

## Find what you need

| Goal | Start here |
|---|---|
| Understand values overrides | [Override test suite](override_tests_suite/README.md) |
| Reproduce rendering failures | [Render problem fixtures](render_problems/) |
| Discover and scan chart repositories | [Helm Fetcher](helm_fetcher/README.md) |
| Inspect the shared repair implementation | [helmfix](helmfix/) |
| Exercise the KICS-style Helm SDK pipeline | [kics-render-mock](kics-render-mock/) |
| Classify failures and inspect repair chains | [Taxonomy analyzer](taxonomy_analyzer/README.md) |
| Understand the research corpus and limitations | [Final catalog results](helm_fetcher/FINAL_RESULTS.md) |
| Trace thesis figures to evidence | [Thesis assets](thesis_assets/README.md) |

## Setup

- Go **1.25.3 or newer** for the toolchain declared by most modules.
- Helm **3** and Git on `PATH` for chart execution and repository cloning.
- Python 3 for plotting and thesis analysis scripts.
- PowerShell and GNU Make for the provided Windows automation targets.

Each tool or suite has its own `go.mod`; run Go commands from that directory.
Keep the repository layout intact: `helm_fetcher` and `kics-render-mock`
reference the sibling `helmfix` module through local `replace` directives.
Dependencies can be resolved by Go on first use. Local `vendor/` directories
are ignored; regenerate them with `go mod vendor` when dependencies change.

## Run the test suites

From the repository root, run the override suite:

```powershell
cd override_tests_suite
go run run-tests.go
```

For one override scenario, use `go run run-tests.go -test 01`.

Run the rendering scenarios from their own directory, starting at the repository root:

```powershell
cd render_problems
go run run-tests.go
```

The runners generate `TEST-RESULTS.md` reports within the test directories,
recording commands, values, and rendering outcomes. Render fixtures deliberately
include failing charts; consult the generated report for each scenario.

## Research workflow

1. **Discover and render** with [Helm Fetcher](helm_fetcher/README.md#discover-and-scan).
   Capture a GitHub or Artifact Hub seed, clone repositories, and write catalogs
   of chart rendering and dependency failures.
2. **Attempt repairs** with the fetcher's `fixer` mode and the shared `helmfix`
   library. Keep each fixed catalog paired with its source catalog.
3. **Classify and inspect** with [taxonomy_analyzer](taxonomy_analyzer/README.md).
   It produces JSON, CSV, Markdown, and a self-contained HTML viewer.
4. **Report and plot** with the [fetcher helpers](helm_fetcher/scripts/README.md)
   and [thesis evidence records](thesis_assets/README.md).

To regenerate the taxonomy from the authoritative checked-in v2 pair, starting
at the repository root:

```powershell
cd taxonomy_analyzer
go run . -input ../helm_fetcher/catalog_sources_v2.json -fixed ../helm_fetcher/catalog_fixed_cumulative_v2.json -out out/cumulative_v2
```

Open `out/cumulative_v2/viewer.html` from that directory. Read
[FINAL_RESULTS.md](helm_fetcher/FINAL_RESULTS.md) before interpreting the numbers:
it documents corpus provenance, superseded snapshots, and the dependency-failure
early stop that affects coverage. A fresh scan does not recreate the historical
corpus automatically.

## Repository layout

| Directory | Responsibility |
|---|---|
| `override_tests_suite/` | Values precedence, types, lists, nested values, subcharts, and globals |
| `render_problems/` | Rendering failure fixtures and their runner |
| `helmfix/` | Shared repair library and tests |
| `helm_fetcher/` | Discovery, cloning, rendering, repair, catalogs, and plots |
| `kics-render-mock/` | Helm SDK rendering pipeline and repair integration experiments |
| `taxonomy_analyzer/` | Failure classification, aggregation, exports, and report viewer |
| `docs/` | Workflow diagrams and implementation notes |
| `thesis_assets/` | Research evidence, figure provenance, and analysis scripts |
| `thesis/` | Optional local thesis checkout, maintained as a separate Git repository |

## Development and data hygiene

Run `go test ./...` inside the module you change, such as `helmfix/`,
`helm_fetcher/`, or `taxonomy_analyzer/`. Use the suite runners above for the
chart fixtures.

The fetcher keeps research catalogs, seed snapshots, reports, and plots in Git.
Its executables, logs, automatic backups, and local assistant settings are ignored;
credentials belong in the ignored `.env` file or environment variables. Historical
`runs/`, `results/`, and `backup/` datasets remain at their existing paths so
research references continue to resolve.

See [LICENSE](LICENSE) for licensing.
