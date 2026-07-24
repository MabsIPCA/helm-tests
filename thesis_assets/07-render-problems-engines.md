# §4.4.x(?) / A5b — five engines on `madgoat-render/` with dependency resolution forced

**Status: measured 2026-07-23, native binaries only (no Docker), not yet applied to any thesis
`.tex` file.** Companion to `06-engine-crosswalk.md`. Where `06` §B pointed three engines at
`render-problems/madgoat-render` *after* neutralising both the render errors and the dependency,
this file keeps the chart **at its committed defaults** and points **five** engines at it —
deliberately forcing every engine down its dependency-resolution ("dep up") path at the same time
as its render-failure path. Section number is a guess by continuity; **verify the real label before
citing** (same caveat `06` carries).

The five engines are the three from `06` — **Kubescape, Trivy, kube-linter** — plus **KICS in its
two configurations, default and enhanced**. (An early draft also ran the full `hcs` image-CVE
pipeline and targeted the `render-problems/` parent directory; both were dropped — image-CVE
methodology lives in `05-helmgoat-image-cve.md`, and the single-chart target below is the one that
actually forces dep-up.)

Raw scan output kept alongside this file in `scan_out/madgoat-render-depup/` for re-verification.

## The target and why dep-up is forced

`helm/render-problems/madgoat-render/` at commit `mad-deployment-service@151f207` (branch `main`),
**the same commit `04`/`05`/`06` used**, working tree clean. At committed defaults the single chart
root presents **three faults at once**:

| Fault | Detail |
|---|---|
| Unvendored dependency | `Chart.yaml` declares `dependencies: [{name: madgoat-infra-render, repository: "file://../madgoat-infra-render"}]`, but there is **no `charts/` and no `Chart.lock`** — so any engine that wants the subchart must run the equivalent of `helm dependency update` itself. **This is what "forces dep up".** |
| Render errors | `values.yaml` sets `renderErrors.r01_nilPointer`/`r02_typeMismatch`/`r03_requiredValue` all `true`; the parent chart fails to render (`nil pointer evaluating interface {}.db`). |
| Invalid `kubeVersion` | `Chart.yaml` declares `kubeVersion: ">= 1.99.0 < 2.0.0"`, satisfiable by no default-tooling cluster. |

Because the subchart is reachable *only through* the render-broken parent, "resolve the dependency"
and "survive the render error" become two separate hurdles — and the interesting result is that
**only KICS-enhanced clears both**; every other engine clears at most one.

**Disk-mutation discipline:** the target was `git status`-clean before each run and restored clean
after (`rm` of any untracked `Chart.lock`/`charts/`, then `git status --porcelain
helm/render-problems` empty). **Kubescape and KICS-enhanced both mutate the tree** (each vendors the
missing dependency to `charts/` + `Chart.lock`); each was reset before the next engine ran and again
at the end.

## Tool provenance

All built/run natively in WSL2, no container (same install-native policy as `04`/`05`/`06`). The
Linux Go toolchain was reinstalled this session (`go1.23.5`, which auto-fetched each module's pinned
`toolchain` — `go1.26.2` for KICS master, `go1.25.0` for the `v2.1.20` tag); the KICS binaries are
freshly built from source, not downloaded.

| Tool | Version / commit | Build / install |
|---|---|---|
| KICS **default** | fork tag `v2.1.20` (matches upstream Checkmarx release; the `--experimental-helm-scan` flag does not exist in it) | `git archive v2.1.20` → `go build ./cmd/console` |
| KICS **enhanced** | fork `master` `3cdfdb2905` | `go build ./cmd/console`, run with `--experimental-helm-scan` |
| Kubescape | `v4.0.11` (build `8fb2eb1`, 2026-07-22) | `~/.kubescape/bin/kubescape` (same binary `06` used) |
| Trivy | `0.72.0` | `~/.local/bin/trivy` (same binary `06` used) |
| kube-linter | `0.8.3` | `~/.local/bin/kube-linter` (same binary `06` used) |

## Results — one target, five verdicts

| Engine | Command | Findings | Exit | Renders parent? | Fixes render error? | Resolves `file://` dep? | Disk mutation |
|---|---|---:|:--:|:--:|:--:|:--:|:--:|
| KICS default | `kics scan -p madgoat-render/` | **6** | 50 | ❌ raw-file scan | ❌ | ❌ | No |
| KICS enhanced | `… --experimental-helm-scan` | **322** | 50 | ✅ | ✅ (6→322) | ✅ vendors subchart | **Yes** |
| Kubescape | `kubescape scan madgoat-render/` | **1 resource**, score **94.07** | 0 | ❌ render aborts | ❌ | ✅ writes `.tgz` (but wasted) | **Yes** |
| Trivy | `trivy config madgoat-render/` | **0** | 0 | ❌ refused | ❌ | ❌ hard refuse | No |
| kube-linter | `kube-linter lint madgoat-render/` | **0** | 0 | ❌ render aborts | ❌ | ❌ never reached | No |

**Headline:** dep-up and render-repair are two independent hurdles, and **only KICS-enhanced clears
both**. Its fixer repairs the render errors *and* vendors the missing dependency, so it renders the
full chart — parent **and** subchart — for **322** findings. Kubescape *resolves the dependency* (it
really does write `madgoat-infra-render-0.1.0.tgz` into `charts/`) but the unfixed render error
collapses the whole render, so the resolution is wasted (1 resource, score 94). Trivy refuses on the
missing dependency before anything else; kube-linter's parent render aborts. So KICS-enhanced is the
only engine that surfaces the chart's real misconfigurations at all — and now the only one that
delivers full coverage of a render-broken, dep-unvendored chart from a single command.

## KICS: enhanced clears both hurdles — fixes the render *and* dep-ups (6 → 322)

On a **single** chart root the `--experimental-helm-scan` flag engages (unlike the multi-root case
`04` documented, where it is a no-op). Its fixer does **two** things in sequence — repairs the
render errors by value injection, and vendors the missing dependency — so the full chart renders,
parent **and** subchart:

| | KICS default `v2.1.20` | KICS enhanced `master` + flag |
|---|---:|---:|
| `total_counter` | 6 | **322** |
| Severity (C/H/M/L/Info) | 0 / 6 / 0 / 0 / 0 | 0 / 22 / 134 / 152 / 14 |
| Distinct categories | 1 (Secret Management) | 9 |
| Files parsed / scanned | 2 / 8 | 13 / 13 |
| Findings on parent's own templates | 6 (`values.yaml` secrets) | 139 (`deployment.yaml` ×123, `service.yaml` ×8, `configmap.yaml` ×5, `render-errors.yaml` ×3) |
| Findings on `madgoat-infra-render` subchart | 0 | **183** |
| Disk mutation | none | **`Chart.lock` + `charts/madgoat-infra-render-0.1.0.tgz`** |

Both fixer steps are visible in the `--log-level DEBUG` log of a run on a verified-clean tree (`ls`
shows only `Chart.yaml`/`templates/`/`values.yaml` before the run):

```
DBG helm fixer: render resolved after value injection attempts=6 chart=madgoat-render
DBG helm fixer: vendored missing chart dependencies, re-rendering chart=madgoat-render dependencies=1
```

- The **value-injection** step (6 attempts) repairs `r01`–`r03` and tolerates the invalid
  `kubeVersion`, taking the parent from the raw-scan baseline (6) to its fully-rendered templates
  (139).
- The **dependency-vendoring** step performs the equivalent of `helm dependency build`, pulling
  `file://../madgoat-infra-render` into `charts/` and re-rendering, which adds the subchart's 183
  findings. This is a real disk mutation — a `Chart.lock` (with digest) and the `.tgz` are written
  into the chart directory, the same side effect Kubescape has (restored here after the run).

**139 + 183 = 322**, matching `04` Target A's enhanced number exactly. The reconciliation with `04`
is now cleaner than the previous draft of this file claimed: `04`'s 322 was measured on an
**already-vendored** tree, and it was an open question whether the fork dep-ups on its own. It now
does — with `charts/` removed to *force* dep-up, enhanced KICS still reaches 322 by vendoring the
subchart itself. So the 322 no longer depends on a human running `helm dependency build` first.

> **History note.** An intermediate build measured during this session gave **139** (render fixed,
> subchart absent — dep-up not yet implemented) and a later one **6** (a regression where the new
> vendoring step ran but bypassed value injection, so the render failed and fell back to raw scan).
> Both were work-in-progress states; the **322** above is the current build with both steps
> chaining correctly. Only the 322 should be cited.

## Kubescape: resolves the dependency, then wastes it on the unfixed render error

`kubescape scan madgoat-render/` returns compliance score **94.07 / 100**, exit **0**, having
scanned exactly **1 resource** (a single `Secret`); control status **111 passed, 0 failed, 7
skipped**. It **did** resolve the dependency — after the run, `charts/madgoat-infra-render-0.1.0.tgz`
and `Chart.lock` are sitting in the chart directory on disk (the same unprompted `helm dependency
update` side effect `06` documents, reproduced here) — but it made **no difference to coverage**:

- The parent chart still fails to render (`nil pointer …` logged at `warn`), and because the
  subchart is reachable only *through* that parent render, the subchart's resources never
  materialise either. The entire chart collapses to the one `Secret` template that happens to
  render standalone.
- Score **94.07 / 100** and exit **0** give no hint that ~46 of ~47 resources were lost. This is
  `06` §B's "47 → 1 resource collapse, reports 94/100, exit 0" signature — reproduced with the
  dependency present but the render still broken.

Contrast with the parent-directory scope (measured in an earlier draft, artifacts not kept):
scanning `render-problems/` as a folder let Kubescape reach `madgoat-infra-render` as an
*independent sibling* chart (29 resources, score 74.77). Targeting `madgoat-render` alone — the
realistic "scan this chart" invocation — removes that escape hatch: the subchart is only a
dependency, and the dependency is dead weight behind a render error.

## Trivy: refuses on the missing dependency

`trivy config madgoat-render/` detects **0** config files, reports **0** misconfigurations, exit
**0**. It hard-skips the chart before render or `kubeVersion` are ever considered:

```
WARN  [helm scanner] Skipping chart  file_path="." err="parse chart: found in Chart.yaml, but missing in charts/ directory: madgoat-infra-render"
INFO  Detected config files  num=0
```

`trivy config` has no dependency-build flag, so on an unvendored chart it scans nothing. Safest
failure mode (no false confidence), zero coverage, and visible only in the `WARN` line — the exit
code is still 0.

## kube-linter: render aborts, zero objects, no JSON

`kube-linter lint madgoat-render/` prints `Warning: no valid objects found.` and exits **0**,
writing **no JSON at all** (`kubelinter.json` is empty — recorded via `kubelinter.stderr`). Unlike
`06` §B (which forced `r01`–`r03` false to get 22 findings), the committed defaults leave the
render errors on, so the parent render fails outright and kube-linter never gets any objects to
lint — the missing dependency is never even reached. `--fail-if-no-objects-found` exists but was
not used; with it, this would be the one engine that fails loudly.

## Cross-engine crosswalk (draft table content, `madgoat-render` at committed defaults)

| Capability on the render-broken, dep-unvendored single chart | KICS default | KICS enhanced | Kubescape | Trivy | kube-linter |
|---|:--:|:--:|:--:|:--:|:--:|
| Surfaces any real chart misconfigurations | ⚠️ 6 (secrets only) | ✅ **322** | ❌ (1 Secret) | ❌ 0 | ❌ 0 |
| Repairs the render error to reach the templates | ❌ | ✅ | ❌ | ❌ | ❌ |
| Resolves the `file://` dependency ("dep up") | ❌ | ✅ (mutates disk) | ✅ (mutates disk) | ❌ (hard refuse) | ❌ (never reached) |
| Delivers full coverage (parent + subchart) | ❌ | ✅ | ❌ (dep-up wasted) | ❌ | ❌ |
| Signals that coverage was lost / chart was dropped | n/a | n/a | ❌ (score 94, exit 0) | ✅ (WARN, exit 0) | ✅ (`no valid objects`, exit 0) |
| Disk mutation as a side effect of a scan | No | **Yes** | **Yes** | No | No |

**Only KICS-enhanced both repairs the render *and* resolves the dependency**, and it is the only
engine that achieves full parent+subchart coverage. Kubescape clears only the dependency hurdle (and
wastes it behind the unfixed render); the other three clear neither. The cost is that KICS-enhanced,
like Kubescape, mutates the source tree (vendored `charts/` + `Chart.lock`) as a side effect of a
scan.

## OWASP Kubernetes Top 10 coverage per engine

`thesis_assets/scripts/scan_coverage.py` runs all five engines over one source, counts each run's
findings, and maps the rule ids that fired to the **OWASP Kubernetes Top 10** IDs (`K01`–`K10`). The
rule-id → OWASP mapping is the same one the thesis renders in `tab:scanner-coverage`
(`chap3-1_helmGoat.tex`); a copy lives at `thesis_assets/scripts/vulnerabilities.yaml` with its
`scanners:` cells synced to that table. Only OWASP-mapped rules contribute to the `OWASP` column; the
`Findings` column counts everything the scanner reported (so KICS-default's 6 secret-regex hits count
as findings but map to no `K`-ID). The tool is non-destructive — it removes any `Chart.lock`/`charts/`
an engine vendors during a run if it was not present beforehand.

Run against `madgoat-render` at committed defaults:

| Engine | Findings | OWASP K-Top-10 detected |
|---|---:|---|
| KICS default | 6 | — |
| KICS enhanced | 322 | **K01, K02, K05, K09** |
| Kubescape | 0 | — |
| Trivy | 0 | — |
| kube-linter | 0 | — |

Union detected by any engine: **K01, K02, K05, K09**. Chart-scope IDs not detected by any engine:
K03, K04, K06, K07, K08, K10.

This is the OWASP-coverage restatement of the headline: on the render-broken, dep-unvendored chart,
**KICS-enhanced is the only engine that detects any OWASP category at all**, and it spans four
(Insecure Workload Configurations K01, Supply-Chain/RBAC K02, Network Segmentation K05, Misconfigured
Cluster Components K09). The four zero-rows are the same collapse documented above — Kubescape's
render aborts to one non-mapped `Secret`, Trivy hard-refuses the missing dependency, kube-linter
finds no valid objects.

**Contrast — the vendored, render-clean chart.** The same tool over `helm/madgoat` (deps vendored,
no render errors — where every engine renders) spreads coverage across all five:

| Engine | Findings | OWASP K-Top-10 detected |
|---|---:|---|
| KICS default | 6 | — |
| KICS enhanced | 333 | K01, K02, K05, K09 |
| Kubescape | 239 | K01, K02, K03, K05 |
| Trivy | 136 | K01, K02, K03 |
| kube-linter | 48 | K01 |

Union there is **K01, K02, K03, K05, K09**. The single difference between the two targets — a broken
render plus an unvendored dependency — is what strips four of the five engines down to zero, leaving
only the one built to handle it. (kube-linter's narrow K01-only spread is expected: its K02/K03/K05
checks are disabled by default, the `†` rows of `tab:scanner-coverage`; and KICS detects no K03
because it has no ConfigMap-content query, per `chap3-1_helmGoat.tex`.)

Regenerate with:

```sh
python3 thesis_assets/scripts/scan_coverage.py \
  /path/to/mad-deployment-service/helm/render-problems/madgoat-render
# --latex FILE / --json FILE to emit tab:scanner-owasp-coverage or a machine-readable roll-up;
# --only kics-enhanced,trivy to run a subset. Binary/asset paths are overridable via
# KICS_ENHANCED_BIN / KUBESCAPE_BIN / TRIVY_BIN / KUBELINTER_BIN (see the script docstring).
```

## Commands to regenerate

```sh
TARGET=/path/to/mad-deployment-service/helm/render-problems/madgoat-render
OUT=/path/to/helm-tests/thesis_assets/scan_out/madgoat-render-depup
MDS=/path/to/mad-deployment-service
# ensure dep-up is forced: no charts/, no Chart.lock
rm -rf "$TARGET/charts" "$TARGET/Chart.lock"

# --- KICS default (stock v2.1.20 from the fork tag; no flag) ---
cd /path/to/kics && git archive v2.1.20 | tar -x -C /tmp/kicsv2
cd /tmp/kicsv2 && rm -rf vendor && go build -o bin/kics ./cmd/console
/tmp/kicsv2/bin/kics scan -p "$TARGET" \
  -q /tmp/kicsv2/assets/queries -b /tmp/kicsv2/assets/libraries \
  -o "$OUT" --output-name kics-default --report-formats json

# --- KICS enhanced (fork master, --experimental-helm-scan) ---
rsync -a --exclude=.git --exclude=vendor /path/to/kics/ /tmp/kics-fork/
cd /tmp/kics-fork && go build -o bin/kics ./cmd/console
/tmp/kics-fork/bin/kics scan -p "$TARGET" \
  -q /tmp/kics-fork/assets/queries -b /tmp/kics-fork/assets/libraries \
  --experimental-helm-scan \
  -o "$OUT" --output-name kics-enhanced --report-formats json

# --- Kubescape (WARNING: writes Chart.lock + charts/*.tgz into the chart dir; restore after) ---
kubescape scan "$TARGET" --format json --output "$OUT/kubescape.json"
rm -rf "$TARGET/charts" "$TARGET/Chart.lock"   # restore

# --- Trivy ---
trivy config "$TARGET" --format json --output "$OUT/trivy-config.json"

# --- kube-linter (writes no JSON when 0 objects; keep stderr) ---
kube-linter lint "$TARGET" --format json > "$OUT/kubelinter.json" 2> "$OUT/kubelinter.stderr"
```

(`-q`/`-b` are required explicitly for KICS — the binary's default `./assets/queries` is relative
to the *current working directory*, not the binary's own location. Same gotcha `04` documents.)

## Reconciliation with the other files

- **`04-helmgoat-misconfig.md`** — Target A measured KICS default = 6, enhanced = 322 on this same
  chart, but on an **already-vendored** tree, leaving open whether the fork dep-ups on its own. `07`
  confirms the **6** exactly and closes that question: with `charts/` removed to *force* dep-up,
  enhanced KICS still reaches **322** by vendoring the subchart itself (139 parent + 183 subchart).
  So enhanced KICS repairs renders **and** dep-ups — the 322 does not depend on a prior manual
  `helm dependency build`.
- **`06-engine-crosswalk.md`** — `06` §B ran Kubescape/Trivy/kube-linter on this chart but with the
  render errors and dependency **neutralised** to isolate one behaviour at a time. `07` runs them
  (plus both KICS modes) at **committed defaults**, so all three faults fire together. The
  Kubescape 94/100 collapse, Trivy hard-refuse, and kube-linter silent behaviours are **consistent**
  with `06` — this is the "all faults on at once" version of the same crosswalk.

## Not covered by this file

- Image CVEs — out of scope (see `05-helmgoat-image-cve.md`); the `hcs` pipeline was dropped.
- The `render-problems/` **parent directory** as a target (both charts as siblings) — measured in
  an early draft, artifacts not kept; it does *not* force dep-up (Kubescape reaches the subchart as
  a sibling instead) and the KICS flag becomes a no-op there per `04`, so it was set aside in favour
  of the single-chart target above.
- Whether the other engines can be *made* to clear both hurdles (e.g. pre-`helm dependency build` +
  Trivy `--helm-kube-version`, or fixing the render by hand before Kubescape) — not benchmarked;
  KICS-enhanced already clears both unaided, and this file otherwise measures **default behaviour on
  the committed tree only**.
- CI-gate exit-code semantics (all four non-KICS engines exit 0 despite near-zero coverage;
  kube-linter's `--fail-if-no-objects-found`) — described, not benchmarked end-to-end.
- A statistically powered comparison — single-target case study, one run per engine.
