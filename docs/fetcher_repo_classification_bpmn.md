# Helm Fetcher — per-repo keep / drop classification (BPMN)

Process executed by `helm_fetcher` in `-mode=full`, once per seeded repository.
Classification logic: `main.go:287-302`; output split: `exporter/exporter.go:66-69`.
(Vector version: `fetcher_repo_classification_bpmn.svg`.)

```mermaid
flowchart TD
    start(["● Seeded repo"])
    clone("Clone repository<br/>git clone")
    gClone{"Clone OK &<br/>chart found?"}
    skip(["○ Not cataloged"])

    discover("Discover charts<br/>FindCharts +<br/>collapse vendored subcharts")
    depbuild("helm dependency build<br/>(per chart)")
    template("helm template<br/>default + every value-file combo<br/>→ records runs")
    agg("Aggregate per repo:<br/>DepFailures / Failures / Successes")

    gDep{"DepFailures<br/>&gt; 0 ?"}
    gFail{"Failures<br/>&gt; 0 ?"}
    gKeep{"-keep-clones ?"}

    dep("DepFailed = true<br/>Kept = false<br/><b>clone retained</b>")
    kept("Kept = true<br/><b>clone retained</b>")
    removed("Kept = false<br/>(clean, no failures)")
    delClone("Delete clone dir")

    depOut[/"Dependency-failure catalog"/]
    keptOut[/"Kept catalog"/]
    remOut[/"Removed catalog"/]
    allOut[/"Full catalog<br/>(all repos)"/]

    eDep(["◉ Dep-Failed"])
    eKept(["◉ Kept"])
    eRem(["◉ Removed"])

    start --> clone --> gClone
    gClone -- no --> skip
    gClone -- yes --> discover --> depbuild --> template --> agg --> gDep

    gDep -- yes --> dep --> depOut --> eDep
    gDep -- no --> gFail
    gFail -- yes --> kept --> keptOut --> eKept
    gFail -- no --> removed --> remOut
    removed --> gKeep
    gKeep -- no --> delClone --> eRem
    gKeep -- yes --> eRem

    dep -.-> allOut
    kept -.-> allOut
    removed -.-> allOut
```

## Reading it

- **Exclusive gateways (diamonds)** are evaluated top-down; the **dep-failure test runs first**, so a repo is bucketed `Dep-Failed` as soon as *any* chart fails `helm dependency build` — even if its other charts templated successfully (e.g. a monorepo).
- **Data objects** (`/…/`) are the output catalogs. Every repo — dep-failed, kept, or removed — is also written to the **full catalog** (dotted flows), which is the authoritative record of all repos.
- **Clone lifecycle:** only the *clean* branch may delete the clone (and only when `-keep-clones=false`). Dep-failed and template-failed clones are always retained.
