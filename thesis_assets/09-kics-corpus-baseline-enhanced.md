# §4.4.1(?) / A7 — Baseline vs Enhanced KICS at corpus scale (top-200 repos)

**Status: measured 2026-07-24, native-Windows builds (no WSL, no Docker), not yet applied to any
thesis `.tex` file.** Scales `04-helmgoat-misconfig.md`'s single-chart baseline-vs-enhanced KICS
result (HELM GOAT: 6 → 322) up to a real corpus: every chart discovered inside the **top-100 GitHub
+ top-100 ArtifactHub** repositories. Section number is a guess by proximity to `04` (§4.4.1) — this
is the corpus-scale companion to that single-chart measurement; **verify the real label/number
before citing.** "Enhanced KICS (used in HCS)" is the same fork build `hcs` wraps (see the caveat in
§Tool provenance).

Raw per-chart results and the aggregate tables are kept in `scan_out/top200-kics/`.

## Headline

| Metric | Baseline KICS | Enhanced KICS (used in HCS) |
|---|---:|---:|
| **Charts analysed** | 9,861 | 9,861 |
| **Findings raised** | 204,086 | 508,248 |
| **Net finding gain** | n/a | **+304,162** |

Both builds were run against the **identical** 9,861 charts (9,857 / 9,838 completed without
error — a 19-chart difference, immaterial). "Findings raised" is each build's KICS `total_counter`,
so the two are directly comparable. **Enhanced raises +304,162 more findings — +149 %, ≈2.5× the
baseline.**

## The scope: what "top 200" means here

The source set is every chart directory `helm.FindCharts` discovers inside the top-100 GitHub and
top-100 ArtifactHub repositories (the same discovery a `helm_fetcher` run uses — vendored subcharts
and `file://` umbrella components already excluded). The two top-100 lists are the first 100 entries
of the star-ranked seed files (`github_search.json`, `artifacthub_search.json`).

| | |
|---|---:|
| Top-200 repos (100 GH + 100 AH, deduped) | 196 |
| …with a clone on disk | 194 |
| …that actually contain ≥1 chart | 154 |
| **Charts discovered / analysed** | **9,861** (GitHub 9,306 · ArtifactHub 555) |

**A blunt caveat: the chart count is dominated by versioned near-duplicates in archive monorepos.**
Just the two rancher catalogs supply ~49 % of the run:

| Repo | Charts |
|---|---:|
| `rancher/partner-charts` | 2,695 |
| `rancher/charts` | 2,092 |
| `trueforge-org/truecharts` | 798 |
| `beclab/apps` | 412 |
| `k0rdent/catalog` | 411 |
| `kubernetes/charts` | 356 |

So "9,861 charts" is **raw chart directories, not distinct charts** — many are
`.../rancher-monitoring/103.0.0`, `.../104.0.0`, … copies of one chart. This run deliberately did
**not** deduplicate (the user chose the raw sweep); a one-version-per-chart variant would cut the set
to roughly 3–4k distinct charts. The gain figures below are per chart-directory.

## Tool provenance

Both KICS builds are **native Windows `.exe`**, built with the Windows Go toolchain and run directly
against the `D:\` clones — a deliberate choice: running the linux binaries under WSL reads every
chart file through the 9p filesystem, which throttled the scan to ~500–900 charts/h; native builds
reading `D:\` directly ran at ~2,500–3,000 charts/h, finishing 9,861 charts in ~4 h instead of days.

| Build | Version / commit | Binary |
|---|---|---|
| Baseline | fork tag `v2.1.20` (upstream Checkmarx release; no `--experimental-helm-scan`) | `C:\Users\miabs\scan-bins-win\kicsv2\kics.exe` |
| Enhanced | fork `master` `3cdfdb2905` (+ current working-tree edits), run with `--experimental-helm-scan` | `C:\Users\miabs\scan-bins-win\kics-fork\kics.exe` |

Discovery: `helm_fetcher/cmd/listcharts` (made OS-aware — walks native `D:\` on Windows, still emits
`/mnt/d` paths for WSL consumers), run natively in **69 s** for all 194 clones. Orchestration:
`thesis_assets/scripts/batch_scan.py run --engines kics-default,kics-enhanced`, 12 workers, 180 s
per-engine timeout, one JSONL record per chart. Non-destructive: any `Chart.lock`/`charts/` a build
vendored during a scan was removed afterward if absent before.

**Native ≠ WSL counts — do not pool.** Native-Windows KICS gives slightly different totals than the
WSL/linux build (enhanced `madgoat-render`: 316 native vs 322 WSL — CRLF / path-separator nuances in
rendering). This run is internally consistent (baseline and enhanced both native), so the
baseline-vs-enhanced *comparison* is valid, but these absolute numbers must not be merged with the
WSL figures in `04`/`07` or the earlier partial WSL corpus run.

## Full results

| | Baseline | Enhanced |
|---|---:|---:|
| Charts analysed | 9,861 | 9,861 |
| Ran ok (no timeout/error) | 9,857 | 9,838 |
| Charts with ≥1 finding | 5,838 | **7,337** |
| Total findings | 204,086 | **508,248** |
| Mean findings / chart | 20.7 | 51.5 |
| Median findings / chart | 2 | 15 |
| Distinct OWASP K-Top-10 ids | 6 | 6 |
| Avg wall-time / chart | 9.5 s | 8.6 s |

**+1,499 charts gained** (5,838 → 7,337 with findings): charts baseline saw as effectively empty —
it cannot render Helm, so on most charts it only matches raw text/secret-regex queries — that
enhanced surfaced findings on once its fixer rendered and dep-up'd them. The median jump (2 → 15) is
the clearest signal: baseline's median chart yields almost nothing, enhanced's yields a full
rendered manifest's worth.

OWASP Kubernetes Top 10 detection — **charts on which each build flagged the id** (mapping via
`scripts/vulnerabilities.yaml`, synced to `tab:scanner-coverage`):

| OWASP | Baseline | Enhanced |
|---|---:|---:|
| K01 Insecure workload config | 3,773 | 5,591 |
| K02 Supply-chain / RBAC | 1,737 | 2,325 |
| K03 Secrets in config | 1,630 | 2,245 |
| K05 Network segmentation | 15 | 28 |
| K06 Broken auth / exposure | 1 | 3 |
| K09 Misconfigured cluster components | 618 | 890 |

Same breadth (6 ids — KICS has no query for K04/K07/K08/K10 in the chart-scope mapping), but enhanced
detects each id on markedly more charts, led by K01 (+1,818 charts) and K03 (+615).

## Reconciliation with `04`

`04-helmgoat-misconfig.md` measured the same two builds on the single render-broken HELM GOAT chart:
**6 → 322** (a >50× jump), because that chart is a worst case where baseline renders nothing. At
corpus scale the ratio is far smaller — **+149 % (≈2.5×)** — and that is the honest, generalisable
number: across thousands of real charts, many are plain enough YAML that baseline's raw scan already
catches a substantial share, so the enhanced delta is a solid multiple rather than two orders of
magnitude. The corpus figure *contextualises* the HELM GOAT headline: the flag's value is real and
large, but the 50× is a single-chart worst case, not the corpus norm.

(An earlier 4,508-chart WSL partial gave +48 %; it was skewed to bitnami's simpler charts and is
superseded by this fuller, native run — do not cite the +48 %.)

## Commands to regenerate

```powershell
# 0. native KICS builds (Windows Go), assets copied alongside each exe
$src="C:\path\to\kics"; $dst="C:\Users\miabs\scan-bins-win"
cd $src; go build -o "$dst\kics-fork\kics.exe" ./cmd/console   # enhanced (master)
git archive v2.1.20 -o v2.tar; mkdir v2; tar -xf v2.tar -C v2
cd v2; go build -o "$dst\kicsv2\kics.exe" ./cmd/console        # baseline v2.1.20
copy assets\  ->  $dst\kics-fork\assets, $dst\kicsv2\assets

# 1. discover charts in the top-200 repos (native, ~69s)
cd helm_fetcher
python -c "extract first 100 repo_url from github_search.json + artifacthub_search.json -> top200_repos.txt"
go build -o listcharts.exe ./cmd/listcharts
.\listcharts.exe -catalog catalog_cumulative.json -repos top200_repos.txt > charts_top200_native.txt
#   (strip BOM/CRLF for WSL use; convert /mnt/d -> D:\ for the native run)

# 2. run both KICS builds over the list (native, ~4h, resumable)
powershell -File thesis_assets\scan_out\top200-kics\run_win.ps1

# 3. aggregate
powershell -File thesis_assets\scan_out\top200-kics\run_win.ps1 -Agg
```

Artifacts: `scan_out/top200-kics/` — `charts_top200_win.txt` (the 9,861 sources), `results_win.jsonl`
(per-chart, both builds), `engine_comparison_win.{tex,json}`, `run_win.ps1` (resumable launcher),
`top200_repos.txt` (the 196-repo filter).

## Not covered by this file

- **Deduplication** — the run counts every versioned chart directory; ~49 % is rancher near-
  duplicates. A distinct-chart variant is not measured here.
- **The other three engines** (Kubescape, Trivy, kube-linter) — this run is KICS-only by request; the
  five-engine comparison on a 4,508-chart WSL partial is in `scan_out/corpus-batch/` (different
  platform, not poolable with these numbers).
- **Full 58,232-chart corpus** — discovered (`scan_out/corpus-batch/charts_all.txt`) but not scanned
  to completion; the full sweep was abandoned as not reachable in useful time.
- **CVE / image scanning** — out of scope (see `05-helmgoat-image-cve.md`).
