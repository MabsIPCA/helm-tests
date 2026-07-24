# §4.4.1 — HELM GOAT misconfiguration detection, baseline vs enhanced KICS

**Status: measured 2026-07-21, native builds only (no Docker).** Replaces the TEMP/empty cells in
`tab:ch4-baseline-enhanced` (chap4:158-170) — thesis-todos.md items **A1** and **A2**. Raw scan
output kept alongside this file for re-verification in `scan_out/`.

An earlier pass used `docker run -v <chart>:/path ...`. That was dropped: the Docker bind mount
(WSL2 ↔ Windows Docker Desktop) silently altered one LOW-severity query's match count (86 vs 92,
183 vs 189 total on the same target, same binary, same chart, only the mount differed). Every
number below is from binaries built and run natively inside WSL, no container involved.

## Which chart is "HELM GOAT" here

Two chart variants exist in `mad-deployment-service/helm/`: the plain `madgoat/` conversion, and
`render-problems/madgoat-render/`, which adds `templates/render-errors.yaml` toggling in the
render-failure taxonomy from chap3-1 (`nil_pointer`, `type_mismatch`, `required_value`,
`unsupported_builtin`, all **on** by default except `r04_unsupportedBuiltin`). This file uses
**only `render-problems/`** — never the plain `madgoat/` chart — because it's the variant that
actually has something for the enhanced renderer to fix; the plain chart wouldn't demonstrate
anything.

Chart source commit: `mad-deployment-service@151f207f2ada52e1db89ec0080526266f94af174` (branch
`main`). `Chart.lock` and `charts/madgoat-infra-render-0.1.0.tgz` under `madgoat-render/` are
untracked local build artifacts from `helm dependency build`, not yet committed — the template
source itself is committed at that hash.

`global.vulnerabilities.*` (the OWASP toggles) are all `true` in `values.yaml` by default — no
`--set` overrides used.

## Two different scan targets — both real, only one gives a valid pair

**Target A — single chart root** (`render-problems/madgoat-render/` only, subchart resolved via
the vendored `charts/*.tgz`, the normal way you'd point a scanner at "the HELM GOAT chart"):

| Metric | Baseline (`v2.1.20`) | Enhanced (fork, flag on) | Gain |
|---|---:|---:|---:|
| Misconfiguration findings (`total_counter`) | **6** | **322** | **+316** |
| Distinct rule categories | **1** | **9** | **+8** |
| Distinct queries fired | 3 | 30 | +27 |
| Files scanned / parsed | 8 / 2 | 13 / 13 | — |

This is the pair used in the table below — it's the only target where the flag visibly does
something, which is the entire point of this section.

**Target B — the `render-problems/` parent directory** (both `madgoat-render/` and
`madgoat-infra-render/` as sibling directories, i.e. what you get if you point KICS at the parent
folder instead of the chart root): baseline gives **189** findings — this is a real,
independently-reproduced number (matches the user's own run exactly), **not fabrication**. But
running the enhanced fork with `--experimental-helm-scan` against the *same* target B also gives
**189**, identical severity breakdown, byte-for-byte. The flag does not engage when the scan path
contains more than one `Chart.yaml` root — confirmed natively, not a mount artifact. So target B
has no usable "enhanced" counterpart and cannot populate this table; it's recorded here only
because it resolves a cross-check (see below).

## Update — the chap4:233 (now chap4:251) "189" cross-check is resolved, see `05-helmgoat-image-cve.md`

thesis-todos.md flagged that A2's misconfiguration count "must equal the 189 quoted at
chap4:233." A4 has now been run (full `hcs` pipeline, `05-helmgoat-image-cve.md`): its internal
KICS pass gives **322**, exactly matching this file's target-A enhanced number. The old 189 there
was simply the invented TEMP value, now replaced. It is a coincidence, and nothing more, that
target B below also happens to measure 189 — the two are unrelated scans of different scopes, not
the same figure counted twice.

**Target B — the `render-problems/` parent directory** (both `madgoat-render/` and
`madgoat-infra-render/` as sibling directories, i.e. what you get if you point KICS at the parent
folder instead of the chart root): baseline gives **189** findings — this is a real,
independently-reproduced number (matches the user's own run exactly), **not fabrication**. But
running the enhanced fork with `--experimental-helm-scan` against the *same* target B also gives
**189**, identical severity breakdown, byte-for-byte. The flag does not engage when the scan path
contains more than one `Chart.yaml` root — confirmed natively, not a mount artifact. So target B
has no usable "enhanced" counterpart and cannot populate this table; it's recorded here only as a
documented curiosity, not because anything in the thesis still depends on it.

## The two tool configurations

| | Baseline | Enhanced |
|---|---|---|
| Source | `MabsIPCA/kics` fork repo, tag `v2.1.20` (`git archive v2.1.20`, matches upstream Checkmarx release — no `--experimental-helm-scan` flag present) | same fork repo, commit `3cdfdb29059d719249c4659498c9999f15f9104f` (branch `master`) |
| Flag | *(none)* | `--experimental-helm-scan` |
| Rule set | 716 queries (`queries_total`) | 716 queries — same count, confirms "same rule set" |

Baseline note: there is no `--type helm` in `v2.1.20` — Helm isn't a scan type upstream
(`Error: unknown argument(s) for --type: helm`), so it runs type-unrestricted. On target A that
means only the platform-agnostic "Common" regex queries get a chance to fire, since the chart
can't be parsed as valid Kubernetes YAML without rendering.

## Commands to regenerate (native, no Docker)

```sh
# --- Baseline: build stock v2.1.20 from the fork repo's own tag ---
cd /path/to/kics                       # the MabsIPCA/kics fork checkout
git archive v2.1.20 | tar -x -C /tmp/kicsv2-1-20
cd /tmp/kicsv2-1-20 && rm -rf vendor && go build -o bin/kics ./cmd/console

# --- Enhanced: build the fork's current source (has --experimental-helm-scan) ---
rsync -a --exclude='.git' --exclude='vendor' /path/to/kics/ /tmp/kics-fork-build/
cd /tmp/kics-fork-build && rm -rf vendor && go build -o bin/kics ./cmd/console

CHART_A=/path/to/mad-deployment-service/helm/render-problems/madgoat-render
CHART_B=/path/to/mad-deployment-service/helm/render-problems
OUT=/path/to/helm-tests/thesis_assets/scan_out

# Target A (the table pair)
/tmp/kicsv2-1-20/bin/kics scan -p "$CHART_A" \
  -q /tmp/kicsv2-1-20/assets/queries -b /tmp/kicsv2-1-20/assets/libraries \
  -o "$OUT" --output-name helmgoat_baseline_chart_native --report-formats json

/tmp/kics-fork-build/bin/kics scan -p "$CHART_A" \
  -q /tmp/kics-fork-build/assets/queries -b /tmp/kics-fork-build/assets/libraries \
  --experimental-helm-scan \
  -o "$OUT" --output-name helmgoat_enhanced_chart_native --report-formats json

# Target B (the 189 cross-check, baseline only — enhanced is identical, not worth re-running)
/tmp/kicsv2-1-20/bin/kics scan -p "$CHART_B" \
  -q /tmp/kicsv2-1-20/assets/queries -b /tmp/kicsv2-1-20/assets/libraries \
  -o "$OUT" --output-name helmgoat_baseline_native --report-formats json
```

(`-q`/`-b` are required explicitly — the binary's default `./assets/queries` is relative to
*current working directory*, not the binary's own location, so running `bin/kics` from anywhere
else without `-q`/`-b` fails with `unable to find queries`.)

## Target A detail

Severity split:

| Severity | Baseline | Enhanced |
|---|---:|---:|
| CRITICAL | 0 | 0 |
| HIGH | 6 | 22 |
| MEDIUM | 0 | 134 |
| LOW | 0 | 152 |
| INFO | 0 | 14 |

Baseline's only 3 queries, all category **Secret Management**, all HIGH: `Passwords And Secrets -
Generic Password` (1), `Passwords And Secrets - Generic Secret` (2), `Passwords And Secrets -
Password in URL` (3) — text/regex "Common"-platform queries that don't require valid Kubernetes
YAML, which is why they're the only thing that survives the chart not rendering.

Enhanced's 9 categories: Access Control, Availability, Best Practices, Build Process, Insecure
Configurations, Insecure Defaults, Networking and Firewall, Resource Management, Secret
Management. File list confirms the full stack was reached, including the deliberately
render-broken `templates/render-errors.yaml` and every `charts/madgoat-infra-render/templates/
*.yaml` file — none of which baseline could parse (2/8 files parsed at baseline vs 13/13
enhanced).

## Table replacement for `tab:ch4-baseline-enhanced` (chap4:168-170) — applied

| Metric | Baseline KICS | Enhanced KICS |
|---|---:|---:|
| Misconfiguration findings | ~~17~~ → **6** | ~~?~~ → **322** |
| Distinct rule categories | ~~6~~ → **1** | ~~34~~ → **9** |
| Misconfiguration gain | n/a | ~~?~~ → **316** |
| Image CVE findings | 0 | still open — out of scope here, see A4 |

Applied to `chap4_experimentalEvaluation.tex:163-175` in the thesis repo (`../thesis`).

## Not covered by this file

- Image CVE findings (A4) — now covered separately in `05-helmgoat-image-cve.md`.
- The capability crosswalk table (`tab:ch4-crosswalk`) and the head-to-head run across Trivy/
  KubeLinter/Kubescape (A5) — untouched.
