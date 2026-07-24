# §4.4.3(?) / A5 — capability crosswalk: Kubescape vs Trivy vs kube-linter

**Status: measured 2026-07-23, native binaries only (no Docker), not yet applied to any thesis
`.tex` file.** Fills the item `04-helmgoat-misconfig.md` left open at its end: *"The capability
crosswalk table (`tab:ch4-crosswalk`) and the head-to-head run across Trivy/KubeLinter/Kubescape
(A5) — untouched."* Section number `§4.4.3` is a guess by numbering-continuity from `04`
(`§4.4.1`) and `05` (`§4.4.2`) — **not confirmed against the actual `.tex` outline**, unlike those
two files. Verify the real label/number before citing.

Raw scan output kept alongside this file in `scan_out/engine-crosswalk/` for re-verification.

## Tool provenance

All three installed natively in WSL2, no container involved (same discipline as `04`/`05`, which
found a Docker bind-mount silently altering KICS output — no such check was repeated here, but the
same install-native policy was followed as a precaution):

| Tool | Version | Install method |
|---|---|---|
| Kubescape | `v4.0.11` (build `8fb2eb1`, built 2026-07-22T14:16:46Z) | official `install.sh` from `kubescape/kubescape` |
| Trivy | `0.72.0` | official `contrib/install.sh` from `aquasecurity/trivy`, vulnerability DB v2, downloaded 2026-07-21 00:22 UTC, updated upstream 2026-07-20 19:22 UTC |
| kube-linter | `0.8.3` | binary downloaded directly from `stackrox/kube-linter` GitHub release `v0.8.3` (the project's own `install.sh` 404s — points at a dead URL) |

Chart source commit: `mad-deployment-service@151f207f2ada52e1db89ec0080526266f94af174` (branch
`main`), the **same commit** `04` and `05` used. Working tree was clean at that commit except for
build-artifact side effects documented below, which this file's own tool runs produced.

**CVE counts will drift** as Trivy's DB and Kubescape's backing vulnerability data update — treat
every CVE figure below as a snapshot dated 2026-07-23, not a fixed ground truth (same caveat `05`
gives for its 4,155-CVE figure).

## Two different questions, two different targets

This file covers two separate experiments:

- **§A — misconfig + CVE capability**, run against the whole `helm/` tree (all three chart
  variants: `madgoat/`, `madgoat-infra/`, `render-problems/*`) — the broadest available surface,
  chosen to exercise as many control categories as possible.
- **§B — Helm rendering robustness** (dependency resolution + render-failure handling), run
  against `render-problems/madgoat-render` specifically — the **same target A** chart `04` and
  `05` use, because it's the only chart in the repo with deliberate render failures
  (`templates/render-errors.yaml`, `r01`–`r03` all `true` by default) and a real local
  `file://../madgoat-infra-render` dependency to test "dep up" against.

## §A — misconfiguration scan, full `helm/` tree

| Tool | Command | Result |
|---|---|---|
| Kubescape | `kubescape scan helm/` | Compliance score **74/100**. Full control breakdown in `scan_out/engine-crosswalk/kubescape-helm-misconfig.json` |
| Trivy | `trivy fs helm/ --scanners misconfig,vuln` | **290** misconfigurations across 48 detected template files. Vuln scanner returned 0 (see §A CVE below — expected) |
| kube-linter | `kube-linter lint helm/` | **100** findings |

Kubescape's highest-severity findings (from the JSON, `resourceIDs` cross-referenced against
`resources[].object`):

| Control | Description | Resources affected |
|---|---|---:|
| `C-0012` | Application credentials in config files | 19 `ConfigMap`s (`env-*`, `realm-config`) |
| `C-0057` | Privileged container | 3 Deployments: `webapp`, `profile`, `mad4shell-unsafe` |
| `C-0267` | Workload with cluster-takeover RBAC | `-traefik` Deployment (all 3 chart copies) |
| `C-0016` | `allowPrivilegeEscalation` not set to `false` | 25 Deployments |

kube-linter's 100 findings, by check:

| Check | Count |
|---:|---|
| `no-read-only-root-fs` | 34 |
| `run-as-non-root` | 34 |
| `unset-cpu-requirements` | 13 |
| `unset-memory-requirements` | 13 |
| `privilege-escalation-container` | 3 |
| `privileged-container` | 3 |

All of kube-linter's findings are a subset of controls Kubescape also flags (same root causes,
narrower rule set — no secrets/RBAC/network/supply-chain categories in kube-linter).

## §A — image CVE scanning, single local command

**The core question: can each tool render local Helm charts *and* CVE-scan every referenced
container image in one invocation, without a live cluster?**

| Tool | Single-command local CVE scan? | How |
|---|---|---|
| **Kubescape** | **Yes** | `kubescape scan helm/ --scan-images` — one command, renders the charts, runs misconfig controls, then pulls and CVE-scans (embedded Grype) every image referenced in `values.yaml`. Both result sets land in one report. |
| **Trivy** | **No** | `trivy fs`/`trivy config` only resolves vulnerabilities from OS-package manifests or language lockfiles physically present in the scanned directory — it does not read `image:` fields out of rendered Helm output and pull them. Confirmed empirically: `trivy fs helm/ --scanners vuln,misconfig` reports 0 vulnerabilities despite `vuln` being enabled (0 language-specific files detected). CVE coverage requires either a separate `trivy image <ref>` per image, or `trivy kubernetes` against a **live cluster**. |
| **kube-linter** | **No** | No image-awareness at all — confirmed via `kube-linter lint --help`: no `image`/`vuln` subcommand, no CVE-related flag of any kind. Pure static config linter. |

`kubescape scan helm/ --scan-images` result: **6,245 vulnerabilities** across 13 images referenced
in the chart set (261 Critical, 1,555 High, 3,280 Medium, 1,149 Other [Low+Negligible+Unknown]).
Full JSON: `scan_out/engine-crosswalk/kubescape-helm-scan-images.json`. Note the top-level
aggregate JSON does **not** attribute individual CVEs back to a specific image — per-image
breakdown below was obtained separately via `kubescape scan image <ref>` run once per image
(13 runs, Grype-format JSON, `scan_out/engine-crosswalk/kubescape-per-image-grype/`).

Per-image severity counts:

| Image | Total | Critical | High | Medium | Low |
|---|---:|---:|---:|---:|---:|
| `mongo:6.0` | 831 | 83 | 276 | 369 | 69 |
| `ghcr.io/mad-goat-project/mc-minio:mc-minio` | 1063 | 34 | 232 | 558 | 233 |
| `ghcr.io/mad-goat-project/keycloak:main` | 437 | 1 | 107 | 253 | 75 |
| `docker.io/traefik:v3.5.3` | 252 | 26 | 107 | 101 | 15 |
| `ghcr.io/mad-goat-project/mad-goat-docs:main` | 410 | 24 | 185 | 180 | 19 |
| `ghcr.io/mad-goat-project/mad-profile-service:main` | 251 | 22 | 89 | 125 | 12 |
| `ghcr.io/mad-goat-project/mad-web-app:madgoat-tech` | 262 | 19 | 128 | 101 | 14 |
| `ghcr.io/mad-goat-project/mad-lessons-service:main` | 223 | 11 | 86 | 92 | 34 |
| `ghcr.io/mad-goat-project/mad-goat4shell-service:unsafe` | 997 | 10 | 63 | 638 | 267 |
| `ghcr.io/mad-goat-project/mad-goat4shell-service:safe` | 993 | 8 | 62 | 638 | 266 |
| `postgres:14.1-alpine` | 151 | 8 | 79 | 56 | 8 |
| `rabbitmq:3-management-alpine` | 162 | 8 | 54 | 83 | 16 |
| `ghcr.io/mad-goat-project/mad-scoreboard-service:main` | 213 | 7 | 87 | 86 | 33 |

Notable: `mad-goat4shell-service:unsafe` carries `log4j-core 2.8.2` with `GHSA-jfh8-c2jp-5v3q`
(Log4Shell, CVE-2021-44228) and `GHSA-7rjr-3q55-vv33` (CVE-2021-45046); the `:safe` tag lacks
those two but still carries Spring/Tomcat-family criticals. **This 13-image list and per-image
counts are a different measurement than `05-helmgoat-image-cve.md`'s "14 images, 4,155 CVEs"**
figure — that run scanned only `render-problems/madgoat-render` via the `feat/image-bom` KICS
branch + `trivy image`, on a Trivy DB snapshot one day older; this run scanned the whole `helm/`
tree via Kubescape's embedded Grype. The two are not directly comparable and neither supersedes
the other — do not merge these numbers without re-deriving both on the same target and DB
snapshot.

## §B — Helm rendering: render-failure handling

Target: `helm/render-problems/madgoat-render` at its committed defaults — `values.yaml` has
`renderErrors.r01_nilPointer`, `r02_typeMismatch`, `r03_requiredValue` all `true`, and
`Chart.yaml` declares `kubeVersion: ">= 1.99.0 < 2.0.0"` (satisfiable by no real or default-tooling
cluster capability).

| Tool | Command | Behavior | Exit code |
|---|---|---|---|
| Kubescape | `kubescape scan render-problems/madgoat-render` | Logs the exact Go-template error (`nil pointer evaluating interface {}.db`), then **silently degrades**: only **1 of ~47** expected resources survives into the final scan. Reports *"All controls passed. No issues found"* at **94/100**. | **0** |
| Trivy | `trivy config render-problems/madgoat-render` | Hard refuses the whole chart: `Skipping chart ... err="parse chart: chart requires kubeVersion: >= 1.99.0 < 2.0.0 which is incompatible with Kubernetes v1.36.0"`. 0 config files scanned. | 0 |
| kube-linter | `kube-linter lint render-problems/madgoat-render` | Hard refuses the whole chart: `failed to render: template: ...nil pointer evaluating interface {}.db`, then `Warning: no valid objects found`. | 0 (default; `--fail-if-no-objects-found` exists but wasn't used) |

**Kubescape's `--fail-coverage-below` does not catch this.** That flag's own help text describes
discounting the score for "silently failed GVR pull"-type scenarios, which reads as directly
relevant to a near-total render failure — tested at `--fail-coverage-below 80` against the broken
chart and it still exited 0 with the same "All controls passed" report. Whatever internal coverage
metric that flag tracks did not register the 47→1 resource collapse.

### Workarounds

| Tool | Can you route around the render failure from the CLI? |
|---|---|
| Kubescape | `--set renderErrors.r01_nilPointer=false --set renderErrors.r02_typeMismatch=false --set renderErrors.r03_requiredValue=false` → jumps from 1 to **47 resources**, subchart included. The invalid `kubeVersion` constraint was **never enforced at all**, with or without these flags. |
| Trivy | Same three `--helm-set` overrides **plus** `--helm-kube-version 1.99.0` (Trivy *does* enforce the constraint, unlike Kubescape) → **18 template files** rendered, including the doubly-nested vendored `traefik` subchart. |
| kube-linter | **No such flags exist.** `kube-linter lint --help` has no `--set`, `--values`, or `--helm-kube-version` of any kind. `--ignore-paths` was tested as a possible workaround (excluding `templates/render-errors.yaml`) and confirmed **not** to help — it filters findings post-render, not files pre-render; the chart still fails to render and still reports "no valid objects found". There is no way to scan this chart with kube-linter without editing the chart source itself. |

Raw JSON: `scan_out/engine-crosswalk/{kubescape,trivy}-render-madgoat-render-{broken,fixed}.json`.
kube-linter produced no JSON artifact for the broken case (0 objects found, nothing to write) —
its stdout is quoted verbatim above.

## §B — Helm rendering: dependency resolution ("dep up")

Same chart, with `Chart.lock` and `charts/` **removed**, isolating the dependency question from
the render-error question above by testing on a scratch copy with `r01`–`r03` forced `false`
(subchart's own `renderErrors.*` are `false` by default already, so only the parent needed
patching). Only `Chart.yaml`'s `dependencies: [{name: madgoat-infra-render, repository:
"file://../madgoat-infra-render"}]` was left to resolve.

| Tool | Behavior when `Chart.lock`/`charts/` are missing | Disk mutation |
|---|---|---|
| Kubescape | **Auto-resolves.** Silently performs the equivalent of `helm dependency update`: packages the sibling `../madgoat-infra-render` chart into a `.tgz`, resolves it, and renders the full chart — including the doubly-nested `traefik` `IngressRoute`/`Middleware` custom resources. | **Yes** — writes `Chart.lock` and `charts/madgoat-infra-render-0.1.0.tgz` back into the chart directory on disk, unprompted, as a side effect of a scan command. Reproduced twice in this session. |
| Trivy | **Hard refuses.** `parse chart: found in Chart.yaml, but missing in charts/ directory: madgoat-infra-render` → chart skipped, 0 misconfig files scanned. No dependency-build flag exists anywhere in `trivy config --help`. | No |
| kube-linter | **Silently ignores the missing dependency.** Renders only the parent chart's own templates (22 legitimate findings on `lesson`/`docs`/`profile`/`scoreboard`/`webapp`/`mad4shell-safe`/`mad4shell-unsafe`) with **zero mention** of `madgoat-infra-render` anywhere in output — no warning, no error. Exit code 1, but only because of those 22 real findings, not because anything was flagged as missing. The entire subchart (traefik, both DB templates, RBAC, NetworkPolicy) is silently absent from results with no signal that anything is incomplete. | No |

### Observed side effects from Kubescape's auto dep-up, still present in the working tree

Kubescape performed the same silent dependency-build the very first time it scanned the *whole*
`helm/` tree in this session (before this experiment was designed), not only in the isolated test
above. Two artifacts from that earlier run are still sitting in the repo as untracked files:

- `helm/madgoat/Chart.lock` + `helm/madgoat/charts/madgoat-infra-0.1.0.tgz` — the latter is
  covered by an existing `.gitignore` rule (`helm/madgoat/charts/`), so it can't leak into a
  commit; `Chart.lock` is **not** ignored and is currently untracked.
- `helm-manager-<hash>-charts.txt` and `helm-manager-<hash>-index.yaml` at the repo root — a
  cached Helm repo index, written while Kubescape attempted (and partially failed, per its own
  warning log) to register `https://traefik.github.io/charts` as a repo for `madgoat-infra`'s
  live-repo `Chart.lock` entry.

These were left in place as of this writing for the record; not yet cleaned up or committed.

## Capability crosswalk (draft `tab:ch4-crosswalk` content)

| Capability | Kubescape v4.0.11 | Trivy v0.72.0 | kube-linter v0.8.3 |
|---|---|---|---|
| Misconfiguration scan of local Helm charts | ✅ | ✅ (`config` / `fs --scanners misconfig`) | ✅ (`lint`) |
| Image CVE scan, single local command, no cluster | ✅ (`--scan-images`) | ❌ — needs separate `trivy image` per image, or a live cluster (`trivy kubernetes`) | ❌ — no image-scanning capability exists |
| Resolves nested/vendored subcharts once rendering succeeds | ✅ | ✅ | ✅ |
| Auto-builds missing `Chart.lock`/`charts/` ("dep up") | ✅ — but mutates the source tree unprompted | ❌ — hard refuses, 0 coverage | ❌ — silently proceeds with 0 coverage of the missing subchart, no warning |
| Enforces `Chart.yaml` `kubeVersion` constraint | ❌ — ignored entirely | ✅ — hard skip if incompatible | N/A — never reached, fails earlier on template render |
| Behavior on template render failure | Silent near-total degradation, still reports a passing score, exit 0 | Hard skip, clear log reason, exit 0 | Hard skip, clear log reason, exit 0 (opt-in hard-fail via `--fail-if-no-objects-found`) |
| CLI override for `values.yaml` to work around a render failure | `--set` / `--values` | `--helm-set` / `--helm-values` | none |
| CLI override for `kubeVersion` | n/a (unenforced) | `--helm-kube-version` | none |
| Disk mutation as a side effect of a read-only scan | Yes | No | No |

## Commands to regenerate

```sh
# --- §A: misconfig, full helm/ tree ---
kubescape scan helm/ --format json --output kubescape-helm-misconfig.json
trivy fs helm/ --scanners misconfig,vuln --format json --output trivy-fs-helm.json
kube-linter lint helm/ --format json --output kubelinter-helm.json

# --- §A: CVE, single local command ---
kubescape scan helm/ --scan-images --format json --output kubescape-helm-scan-images.json
# per-image breakdown (repeat per image found in the chart set's values.yaml):
kubescape scan image "<image:tag>" --format json --output "<image>.json"

# --- §B: render-failure handling (chart at its committed defaults) ---
kubescape scan helm/render-problems/madgoat-render
trivy config helm/render-problems/madgoat-render
kube-linter lint helm/render-problems/madgoat-render

# --- §B: render-failure workaround ---
kubescape scan helm/render-problems/madgoat-render \
  --set renderErrors.r01_nilPointer=false \
  --set renderErrors.r02_typeMismatch=false \
  --set renderErrors.r03_requiredValue=false
trivy config helm/render-problems/madgoat-render \
  --helm-kube-version 1.99.0 \
  --helm-set renderErrors.r01_nilPointer=false \
  --helm-set renderErrors.r02_typeMismatch=false \
  --helm-set renderErrors.r03_requiredValue=false

# --- §B: dep-up (run against a copy with Chart.lock/charts/ removed and r01-r03 pre-patched
#     false in values.yaml, to isolate dependency resolution from the render-error case above) ---
rm -rf madgoat-render/Chart.lock madgoat-render/charts
kubescape scan madgoat-render/
trivy config madgoat-render/
kube-linter lint madgoat-render/
```

## Not covered by this file

- Whether Kubescape's auto dep-up behavior is configurable/disableable (no such flag was found in
  `kubescape scan --help`, but this wasn't exhaustively searched against the source).
- Reconciling the 13-image/6,245-CVE figure here against `05`'s 14-image/4,155-CVE figure — see
  the explicit non-comparability note in §A above.
- kube-linter's `--fail-if-no-objects-found` and Trivy's exit-code behavior under a CI gate were
  described but not benchmarked end-to-end in a CI harness.
- Corpus-scale replication (all three tools run once, on one chart, in one session) — this is a
  single-chart case study, not a statistically powered comparison.
