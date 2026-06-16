# Helm Fetcher — Final Catalog Results

_Generated 2026-06-16. Top-500 seeds per source, value-prefiltered basis._

## Overview

Each source was seeded from a top-500 list and processed in `-mode=full`
(clone → discover charts → `helm dependency build` → `helm template` on the
default values plus every value-file combination). Results below are on the
**value-prefiltered basis**: overlay value files that are git-crypt encrypted or
otherwise unparseable as YAML are excluded from the combinations, so they no
longer produce false template failures.

| Source | Seed file (`top: 500`) | Run directory | Repos cataloged |
|---|---|---|---:|
| GitHub | `github_search.json` | `runs/20260528_215041_github/` | 487 / 500 |
| ArtifactHub | `artifacthub_search.json` | `runs/20260529_135859_artifacthub/` | 486 / 500 |

> Repos seeded but not cataloged yielded no chart or failed to clone.

## Results by source

| Metric | GitHub | ArtifactHub | Combined |
|---|---:|---:|---:|
| Repos | 487 | 486 | 973 |
| Charts discovered | 69,034 | 3,191 | 72,225 |
| Helm template runs | 8,012 | 2,098 | 10,110 |
| ✅ Successes | 5,814 | 1,850 | 7,664 |
| ❌ Template failures | 2,198 | 248 | 2,446 |
| 🔧 Dep-build failures (runs) | 194 | 24 | 218 |
| **Success rate** | **72.6%** | **88.2%** | **75.8%** |

### Repo status breakdown

| Status | GitHub | ArtifactHub | Combined |
|---|---:|---:|---:|
| Kept (has template failures) | 86 | 67 | 153 |
| Removed (clean, no failures) | 207 | 395 | 602 |
| Dep-failed (dependency build failed) | 194 | 24 | 218 |
| **Total** | **487** | **486** | **973** |

## Notes

- **Value prefilter** (`helm.FindValuesFiles` → `IsParseableValuesFile`): an
  overlay value file is only fed to `helm template` if it parses as a YAML map.
  Encrypted (git-crypt) and malformed files are dropped, matching what helm
  itself would refuse on load. This removes false failures from the catalog.
  - GitHub: 6 / 487 repos affected, 15 false failures removed (rate unchanged at
    72.6% — GitHub had few such files).
  - ArtifactHub: prefilter applied; basis reflected in the 88.2% above.
- **Authoritative data** is each run's `catalog_by_project.json`; the
  `catalog_results.md` siblings were regenerated from it via `-mode=render-md`
  so they match exactly. Pre-prefilter snapshots are preserved at
  `catalog_by_project.json.prefilter.bak`.
- The stale April top-500 GitHub data in `results/github/` and
  `backup/GHTop500/` is **superseded** by `runs/20260528_215041_github/` (only
  ~91/322 of its repos overlap the current seed).

## Source files

```
runs/20260528_215041_github/catalog_by_project.json     # GitHub, authoritative
runs/20260528_215041_github/catalog_results.md          # GitHub, rendered
runs/20260529_135859_artifacthub/catalog_by_project.json # ArtifactHub, authoritative
runs/20260529_135859_artifacthub/catalog_results.md      # ArtifactHub, rendered
github_search.json / artifacthub_search.json             # top-500 seeds
```
