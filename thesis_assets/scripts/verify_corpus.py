#!/usr/bin/env python3
"""Recompute the thesis corpus figures from the helm_fetcher catalogs and assert them.

Covers the Ch. 3.2 corpus figures (June per-source runs) and the Ch. 3.3 v2 repo-split
figures (catalog_sources_v2.json). Run from helm_fetcher/ (or pass --fetcher-dir). Exits
non-zero if any figure cited in the thesis no longer matches the data, so a stale number
cannot survive a re-run unnoticed.

    cd <helm-tests>/helm_fetcher && py <this>/verify_corpus.py
"""
from __future__ import annotations

import argparse
import json
import sys
from pathlib import Path

# The two authoritative per-source runs. results/github/ is the SUPERSEDED April snapshot
# (FINAL_RESULTS.md:55) and must never be used here — see 02-taxonomy.md.
GITHUB = "runs/20260528_215041_github/catalog_by_project.json"
ARTIFACTHUB = "runs/20260529_135859_artifacthub/catalog_by_project.json"
SUPERSEDED = "results/github/catalog_by_project.json"

# The current (v2) merged source catalog — basis for the Ch. 3.3 taxonomy numbers.
V2 = "catalog_sources_v2.json"

# The taxonomy analyzer's final report (relative to the fetcher dir). This is the product
# cited by Ch. 3.3 tab:ch5-fix-results; it is reproducible from V2 + catalog_fixed_cumulative_v2.
REPORT = "../taxonomy_analyzer/out/cumulative_v2/taxonomy_report.json"

# Figures as cited in the thesis, keyed to their location.
EXPECTED = {
    "entries (chap3-2:501)": 973,
    "distinct repos (correction 3)": 963,
    "charts (chap3-2:501)": 72_225,
    "runs (chap3-2:501)": 10_110,
    "successes (FINAL_RESULTS)": 7_664,
    "template failures (chap3-2:503)": 2_446,
    "dep failures (chap3-2:503)": 218,
    "kept repos (chap3-2:503)": 153,
    "repos with failures (correction 1)": 200,
    "failures in kept repos (correction 1)": 2_102,
    "seed overlap (correction 3)": 10,
    # D3 funnel: 92.5% of discovered charts sit in dep-failed archive monorepos.
    "dep-failed repos (D3)": 218,
    "charts in dep-failed repos (D3)": 66_837,
    "charts in renderable repos (D3)": 5_388,
    "repos with charts but 0 runs (D3)": 93,
}

# Ch. 3.3 v2 repo-level failure attribution (catalog_sources_v2.json). The dep-failure test is
# applied first, so a repo can log template failures from charts visited before the abort and
# still be recorded dep-failed. See 02-taxonomy.md "Repo-level failure attribution".
EXPECTED_V2 = {
    "v2 template failures (chap3-3)": 2_141,
    "v2 repos with template failures": 178,
    "v2 dep-failed repos w/ tmpl failures": 36,
    "v2 failures in those dep-failed repos": 177,
    "v2 kept repos w/ tmpl failures": 142,
    "v2 failures in those kept repos": 1_964,
}

# Ch. 3.3 taxonomy totals, read from the analyzer report (taxonomy_report.json "totals").
EXPECTED_TAX = {
    "tax repos": 973,
    "tax runs": 9_600,
    "tax template failures": 2_141,
    "tax dependency failures": 206,
    "tax duplicates collapsed": 930,
    "tax classified errors": 1_388,
    "tax unclassified errors": 29,
    "tax fix attempted": 1_211,
    "tax fix resolved": 723,
    "tax fix unresolved": 488,
}

# Ch. 3.3 tab:ch5-fix-results — template sub-kinds: (count, resolved). unresolved = count - resolved.
EXPECTED_SUBKIND = {
    "template.nil_pointer": (587, 540),
    "template.required_value": (263, 159),
    "template.kube_version_incompatible": (24, 24),
    "template.type_mismatch": (15, 0),
    "template.missing_template": (110, 0),
    "template.values_schema_validation": (92, 0),
    "template.malformed_yaml": (23, 0),
    "template.author_assertion": (21, 0),
    "template.runtime_eval": (21, 0),
    "template.custom_validation": (20, 0),
    "template.library_chart_not_installable": (11, 0),
    "template.parse_error": (9, 0),
    "template.unsupported_builtin": (8, 0),
    "template.chart_metadata": (4, 0),
    "template.invalid_value": (2, 0),
    "template.dependency_check_failed": (1, 0),
}

# Dependency taxonomy figures cited in 02-taxonomy.md ("4 classified -> 177").
EXPECTED_DEP_SUBKIND = {
    "dependency.missing_repository": 84,
    "dependency.missing_subchart": 41,
    "dependency.chart_validation": 14,
    "dependency.lock_file_out_of_sync": 9,
}


def load(path: Path) -> list[dict]:
    with path.open(encoding="utf-8") as fh:
        data = json.load(fh)
    return data if isinstance(data, list) else (data.get("repos") or next(iter(data.values())))


def total(repos: list[dict], field: str) -> int:
    return sum(r.get(field, 0) for r in repos)


def main() -> int:
    ap = argparse.ArgumentParser()
    ap.add_argument("--fetcher-dir", type=Path, default=Path.cwd())
    args = ap.parse_args()
    base = args.fetcher_dir

    missing = [p for p in (GITHUB, ARTIFACTHUB) if not (base / p).is_file()]
    if missing:
        print(f"error: catalog(s) not found under {base}: {', '.join(missing)}", file=sys.stderr)
        print("hint: run from the helm_fetcher directory, or pass --fetcher-dir", file=sys.stderr)
        return 2

    gh, ah = load(base / GITHUB), load(base / ARTIFACTHUB)
    corpus = gh + ah

    names_gh = {r["repo_name"] for r in gh}
    names_ah = {r["repo_name"] for r in ah}
    failing = [r for r in corpus if r.get("total_failures", 0) > 0]
    kept = [r for r in corpus if r.get("kept")]
    dep_failed = [r for r in corpus if r.get("dep_failed")]
    renderable = [r for r in corpus if not r.get("dep_failed")]
    stranded = [r for r in corpus
                if r.get("total_charts", 0) > 0 and r.get("total_runs", 0) == 0]

    actual = {
        "entries (chap3-2:501)": len(corpus),
        "distinct repos (correction 3)": len(names_gh | names_ah),
        "charts (chap3-2:501)": total(corpus, "total_charts"),
        "runs (chap3-2:501)": total(corpus, "total_runs"),
        "successes (FINAL_RESULTS)": total(corpus, "total_successes"),
        "template failures (chap3-2:503)": total(corpus, "total_failures"),
        "dep failures (chap3-2:503)": total(corpus, "total_dep_failures"),
        "kept repos (chap3-2:503)": len(kept),
        "repos with failures (correction 1)": len(failing),
        "failures in kept repos (correction 1)": total(kept, "total_failures"),
        "seed overlap (correction 3)": len(names_gh & names_ah),
        "dep-failed repos (D3)": len(dep_failed),
        "charts in dep-failed repos (D3)": total(dep_failed, "total_charts"),
        "charts in renderable repos (D3)": total(renderable, "total_charts"),
        "repos with charts but 0 runs (D3)": len(stranded),
    }

    print(f"corpus = {GITHUB}\n       + {ARTIFACTHUB}\n")
    width = max(len(k) for k in EXPECTED)
    failures = 0
    for key, want in EXPECTED.items():
        got = actual[key]
        ok = got == want
        failures += not ok
        print(f"  {'ok  ' if ok else 'FAIL'}  {key:<{width}}  expected {want:>7,}  got {got:>7,}")

    rate = total(corpus, "total_successes") / total(corpus, "total_runs") * 100
    rate_ok = round(rate, 1) == 75.8
    failures += not rate_ok
    print(f"  {'ok  ' if rate_ok else 'FAIL'}  {'success rate (chap3-2:501)':<{width}}  expected    75.8%  got   {rate:.1f}%")

    # Ch. 3.3 v2 repo-level failure attribution.
    v2_path = base / V2
    if v2_path.is_file():
        v2 = load(v2_path)
        v2_tf = [r for r in v2 if r.get("total_failures", 0) > 0]
        v2_dep = [r for r in v2_tf if r.get("dep_failed")]
        v2_kept = [r for r in v2_tf if not r.get("dep_failed")]
        actual_v2 = {
            "v2 template failures (chap3-3)": total(v2, "total_failures"),
            "v2 repos with template failures": len(v2_tf),
            "v2 dep-failed repos w/ tmpl failures": len(v2_dep),
            "v2 failures in those dep-failed repos": total(v2_dep, "total_failures"),
            "v2 kept repos w/ tmpl failures": len(v2_kept),
            "v2 failures in those kept repos": total(v2_kept, "total_failures"),
        }
        print(f"\ncorpus (Ch.3.3 v2) = {V2}\n")
        width_v2 = max(len(k) for k in EXPECTED_V2)
        for key, want in EXPECTED_V2.items():
            got = actual_v2[key]
            ok = got == want
            failures += not ok
            print(f"  {'ok  ' if ok else 'FAIL'}  {key:<{width_v2}}  expected {want:>7,}  got {got:>7,}")
    else:
        print(f"\n  note  {V2} not found — Ch.3.3 v2 figures not checked")

    # Ch. 3.3 taxonomy analyzer report (out/cumulative_v2/taxonomy_report.json).
    report_path = base / REPORT
    if report_path.is_file():
        with report_path.open(encoding="utf-8") as fh:
            rep = json.load(fh)
        t, bsk = rep["totals"], rep["by_sub_kind"]
        actual_tax = {
            "tax repos": t["repos"],
            "tax runs": t["runs"],
            "tax template failures": t["template_failures"],
            "tax dependency failures": t["dependency_failures"],
            "tax duplicates collapsed": t["duplicates_collapsed"],
            "tax classified errors": t["classified_errors"],
            "tax unclassified errors": t["unclassified_errors"],
            "tax fix attempted": t["fix_attempted"],
            "tax fix resolved": t["fix_resolved"],
            "tax fix unresolved": t["fix_unresolved"],
        }
        print(f"\ntaxonomy (Ch.3.3) = {REPORT}\n")
        width_t = max(len(k) for k in EXPECTED_TAX)
        for key, want in EXPECTED_TAX.items():
            got = actual_tax[key]
            ok = got == want
            failures += not ok
            print(f"  {'ok  ' if ok else 'FAIL'}  {key:<{width_t}}  expected {want:>7,}  got {got:>7,}")

        # tab:ch5-fix-results — per-template-sub-kind count + resolved.
        print()
        for sk, (want_c, want_r) in EXPECTED_SUBKIND.items():
            e = bsk.get(sk, {})
            got_c = e.get("count", 0)
            got_r = e.get("fix_outcome", {}).get("resolved", 0)
            ok = got_c == want_c and got_r == want_r
            failures += not ok
            print(f"  {'ok  ' if ok else 'FAIL'}  {sk:<40}  count {want_c:>4}/{got_c:<4}  resolved {want_r:>3}/{got_r}")

        # dependency taxonomy counts cited in prose.
        print()
        for sk, want in EXPECTED_DEP_SUBKIND.items():
            got = bsk.get(sk, {}).get("count", 0)
            ok = got == want
            failures += not ok
            print(f"  {'ok  ' if ok else 'FAIL'}  {sk:<40}  count {want:>4}/{got}")

        # derived rates cited in prose.
        rec = round(t["fix_resolved"] / t["fix_attempted"] * 100, 1)
        inj_applicable = sum(EXPECTED_SUBKIND[k][0] for k in (
            "template.nil_pointer", "template.required_value",
            "template.kube_version_incompatible", "template.type_mismatch"))
        inj = round(t["fix_resolved"] / inj_applicable * 100, 1)
        for label, got, want in (("overall recovery", rec, 59.7),
                                 ("injection-applicable (723/889)", inj, 81.3)):
            ok = got == want
            failures += not ok
            print(f"  {'ok  ' if ok else 'FAIL'}  {label:<40}  expected {want:>6}%  got {got:>6}%")
    else:
        print(f"\n  note  {REPORT} not found — Ch.3.3 taxonomy figures not checked")

    # The superseded catalog must not be mistaken for the corpus. Report the overlap that
    # makes the two non-interchangeable (see 02-taxonomy.md).
    stale_path = base / SUPERSEDED
    if stale_path.is_file():
        stale = load(stale_path)
        shared = len({r["repo_name"] for r in stale} & names_gh)
        print(f"\n  note  superseded {SUPERSEDED}: {len(stale)} repos, "
              f"{total(stale, 'total_runs'):,} runs, {total(stale, 'total_failures'):,} failures")
        print(f"  note  it shares only {shared} of the {len(names_gh)} authoritative GitHub repos")

    print()
    if failures:
        print(f"{failures} figure(s) NO LONGER MATCH the catalogs — update thesis_assets/01-corpus.md / 02-taxonomy.md")
        return 1
    print("all Ch. 3.2 + Ch. 3.3 (v2) figures reproduce from the catalogs")
    return 0


if __name__ == "__main__":
    sys.exit(main())
