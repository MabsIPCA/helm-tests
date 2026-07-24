# §4.4.2 / A4 — HELM GOAT image CVE detection, full `hcs` run

**Status: measured 2026-07-21, native builds only (no Docker).** Fills `tab:ch4-image-cve`
(chap4:243-244), the "Image CVE findings" row of `tab:ch4-baseline-enhanced` (chap4:178), and the
prose at chap4:251 — thesis-todos.md item **A4**. Raw output: `scan_out/hcs/hcs.sarif`,
`scan_out/hcs/hcs-summary.md`.

## What ran

The full `hcs` pipeline (see `hcs/README.md`) against `render-problems/madgoat-render` (same
target A as `04-helmgoat-misconfig.md`, same chart commit `151f207`):

1. `kics scan --image-bom --experimental-helm-scan` — misconfigurations + image inventory (a
   CycloneDX BoM per discovered image), from `MabsIPCA/kics` branch **`feat/image-bom`** (commit
   `ac802547839e0dfc259d1fa667880c443df06c69`) — **not** the `master` branch used for `04`'s
   enhanced numbers. This is deliberate: `hcs`'s own `Dockerfile` pins `feat/image-bom`, so this
   is the actual binary `hcs` runs in production, not a substitute.
2. `trivy image` once per discovered image (14 runs), SARIF output.
3. `hcs`'s internal merge step combines all 15 SARIF runs into one `hcs.sarif`.

## Commands to regenerate

```sh
# 1. kics with --image-bom (branch feat/image-bom, not master)
cd /path/to/kics
git fetch fork feat/image-bom   # fork = https://github.com/MabsIPCA/kics.git
git archive fork/feat/image-bom | tar -x -C /tmp/kics-image-bom
cd /tmp/kics-image-bom && rm -rf vendor && go build -o bin/kics ./cmd/console

# 2. trivy (native build — GOEXPERIMENT=jsonv2 required, trivy imports encoding/json/v2)
GOEXPERIMENT=jsonv2 go install github.com/aquasecurity/trivy/cmd/trivy@latest

# 3. hcs
rsync -a --exclude='.git' /path/to/hcs/ /tmp/hcs-build/
cd /tmp/hcs-build && go build -o bin/hcs ./cmd/hcs

# 4. run
CHART=/path/to/mad-deployment-service/helm/render-problems/madgoat-render
OUT=/path/to/helm-tests/thesis_assets/scan_out/hcs
/tmp/hcs-build/bin/hcs scan "$CHART" \
  --output "$OUT/hcs.sarif" --summary "$OUT/hcs-summary.md" \
  --kics-bin /tmp/kics-image-bom/bin/kics \
  --kics-query-path /tmp/kics-image-bom/assets/queries \
  --trivy-bin ~/go/bin/trivy
```

`hcs` commit: `66ad605757e7f39287efe5015da7c29cca4d0d7f`. Trivy vulnerability DB: version 2,
updated 2026-07-20 19:22 UTC — **CVE counts will drift** as the DB updates; this is a snapshot,
not a fixed ground truth like the misconfiguration counts.

## Results

| Metric | Baseline KICS | Enhanced (`hcs`) |
|---|---:|---:|
| Container images inventoried | 0 | **14** |
| Image CVEs detected | 0 | **4,155** |
| Chart misconfiguration findings (KICS pass inside `hcs`) | — | **322** |
| Unified SARIF findings (322 + 4,155) | — | **4,477** |
| Scan runs (1 KICS + 14 Trivy) | — | **15** |

CVE severity breakdown, summed across the 14 images (from `hcs-summary.md`'s per-image table):

| Severity | Count |
|---|---:|
| Critical | 57 |
| High | 867 |
| Medium | 2,214 |
| Low | 1,017 |
| **Total** | **4,155** |

## The 322 cross-check — now resolved

`04-helmgoat-misconfig.md` flagged that thesis-todos.md wanted A2's misconfiguration count to
"equal the 189 quoted at chap4:233," and left that open pending this run. It's now resolved: the
misconfiguration count from `hcs`'s own internal KICS pass (`feat/image-bom` branch,
`--image-bom --experimental-helm-scan`) is **322**, exactly matching `04`'s independently-measured
target-A enhanced number (`master` branch, `--experimental-helm-scan` only). The old **189** was
simply the invented TEMP value at chap4:251 — not a different-but-valid measurement (unlike the
189 in the *other* location discussed in `04`, which came from scanning `render-problems/` as a
whole directory and is real, just from an unrelated scan target). Both TEMP occurrences of 189 are
now replaced with sourced numbers; they happen to be the same digits by coincidence, nothing more.

## Not verified in this pass

- **Log4Shell `:safe`/`:unsafe` narrative** (chap4, prose after this table, not `\myworries`-flagged
  so left untouched): both `mad-goat4shell-service:safe` (731 CVEs) and `:unsafe` (735 CVEs) carry
  vulnerabilities in this run — the claim is about a *specific* CVE (Log4Shell, CVE-2021-44228)
  being present in one and absent in the other, not about total CVE count. Not checked here; grep
  `hcs.sarif` for `CVE-2021-44228` per image before trusting that sentence.
- Distinct container images inventoried at corpus scale (`tab:ch4-corpus-cve`, chap4:343) — separate
  item (A6), untouched.
