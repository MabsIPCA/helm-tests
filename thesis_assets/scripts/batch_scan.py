#!/usr/bin/env python3
"""Batch-run the five engines over a corpus of Helm chart sources and build an
engine-comparison table.

Discovery is delegated to the fetcher's own mechanism: feed this a chart-list
file produced by `helm_fetcher/cmd/listcharts` (which uses `helm.FindCharts`, so
the sources are exactly what a fetcher run would enumerate — vendored subcharts
and file:// components already excluded).

    # 1. discover (same mechanism as the fetcher)
    go run ./cmd/listcharts -catalog catalog_cumulative.json > charts_all.txt

    # 2. run the batch (resumable; safe to Ctrl-C and restart)
    python3 batch_scan.py run --charts charts_all.txt --out results.jsonl

    # 3. aggregate any time (partial or complete)
    python3 batch_scan.py agg --out results.jsonl --latex compare.tex

Every chart yields one JSONL record with each engine's finding count, ok flag,
wall time, and detected OWASP Kubernetes Top 10 ids. The run is resumable: on
restart, charts already present in the JSONL are skipped. Raw per-scanner JSON is
not retained at corpus scale (it would be terabytes) — the JSONL is the durable
record of every result.

The run is non-destructive: any Chart.lock/charts/ an engine vendors (KICS-
enhanced, Kubescape) is removed after the chart if it was not present before, so
the frozen corpus on disk is left as found.
"""
import argparse
import json
import os
import statistics
import sys
import threading
import time
from concurrent.futures import ThreadPoolExecutor, as_completed

sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))
import scan_coverage as sc   # noqa: E402

ENGINES = ["kics-default", "kics-enhanced", "kubescape", "trivy", "kube-linter"]
SHORT = {"kics-default": "kics-def", "kics-enhanced": "kics-enh", "kubescape": "kubescape",
         "trivy": "trivy", "kube-linter": "kube-lint"}


def scan_one(src, bins, maps, timeout, tmp_root, engines=ENGINES):
    """Run the selected engines over one chart. Returns the JSONL record dict."""
    import shutil
    import tempfile
    outdir = tempfile.mkdtemp(prefix="bs-", dir=tmp_root)
    roots, pre = sc.snapshot_untracked(src)
    rec = {"chart": src, "ts": int(time.time()), "engines": {}}
    runners = {
        "kics-default": lambda: sc.run_kics(
            bins["kics_default"], bins["kics_default_assets"], src, outdir, False, timeout),
        "kics-enhanced": lambda: sc.run_kics(
            bins["kics_enhanced"], bins["kics_enhanced_assets"], src, outdir, True, timeout),
        "kubescape": lambda: sc.run_kubescape(bins["kubescape"], src, outdir, timeout),
        "trivy": lambda: sc.run_trivy(bins["trivy"], src, outdir, timeout),
        "kube-linter": lambda: sc.run_kube_linter(bins["kube_linter"], src, outdir, timeout),
    }
    mapkey = {"kics-default": "kics", "kics-enhanced": "kics", "kubescape": "kubescape",
              "trivy": "trivy", "kube-linter": "kube_linter"}
    try:
        for name in engines:
            t0 = time.time()
            try:
                count, fired, ok = runners[name]()
            except Exception as e:                       # never let one engine kill the chart
                count, fired, ok = 0, set(), False
                rec.setdefault("errors", {})[name] = str(e)[:200]
            owasp, vulns = sc.owasp_hits(fired, maps[mapkey[name]])
            rec["engines"][name] = {
                "findings": count, "ok": ok, "secs": round(time.time() - t0, 1),
                "owasp": sorted(owasp), "vulns": sorted(vulns),
            }
            sc.restore_untracked(roots, pre)             # undo dep-up between engines
    finally:
        sc.restore_untracked(roots, pre)
        shutil.rmtree(outdir, ignore_errors=True)
    return rec


def load_done(path):
    done = set()
    if os.path.exists(path):
        with open(path, encoding="utf-8") as f:
            for line in f:
                line = line.strip()
                if not line:
                    continue
                try:
                    done.add(json.loads(line)["chart"])
                except (json.JSONDecodeError, KeyError):
                    pass
    return done


def cmd_run(args):
    bins = sc.resolve_binaries()
    maps = sc.build_maps(sc.parse_vulns(sc.DATA))
    engines = [e.strip() for e in args.engines.split(",")] if args.engines else ENGINES
    bad = [e for e in engines if e not in ENGINES]
    if bad:
        sys.exit(f"batch: unknown engine(s): {bad}; valid: {ENGINES}")
    with open(args.charts, encoding="utf-8") as f:
        charts = [ln.strip() for ln in f if ln.strip()]
    done = load_done(args.out)
    todo = [c for c in charts if c not in done]
    if args.limit:
        todo = todo[:args.limit]
    tmp_root = args.tmp or "/tmp/batch-scan"
    os.makedirs(tmp_root, exist_ok=True)

    print(f"batch: {len(charts)} charts, {len(done)} already done, {len(todo)} to scan "
          f"({args.workers} workers, {args.timeout}s/engine, engines={engines})",
          file=sys.stderr)

    lock = threading.Lock()
    out = open(args.out, "a", encoding="utf-8")
    n = 0
    t_start = time.time()

    def write(rec):
        nonlocal n
        with lock:
            out.write(json.dumps(rec) + "\n")
            out.flush()
            os.fsync(out.fileno())
            n += 1
            if n % 25 == 0 or n == len(todo):
                rate = n / max(time.time() - t_start, 1)
                eta = (len(todo) - n) / rate / 3600 if rate else 0
                print(f"  {n}/{len(todo)}  ({rate*3600:.0f}/h, ETA {eta:.1f}h)  {rec['chart']}",
                      file=sys.stderr)

    try:
        if args.workers <= 1:
            for c in todo:
                write(scan_one(c, bins, maps, args.timeout, tmp_root, engines))
        else:
            with ThreadPoolExecutor(max_workers=args.workers) as ex:
                futs = {ex.submit(scan_one, c, bins, maps, args.timeout, tmp_root, engines): c
                        for c in todo}
                for fut in as_completed(futs):
                    try:
                        write(fut.result())
                    except Exception as e:               # pragma: no cover
                        write({"chart": futs[fut], "ts": int(time.time()),
                               "engines": {}, "fatal": str(e)[:200]})
    finally:
        out.close()
    print(f"batch: wrote {n} records to {args.out}", file=sys.stderr)


# ── aggregation ──────────────────────────────────────────────────────────────
def read_records(path):
    recs = []
    with open(path, encoding="utf-8") as f:
        for line in f:
            line = line.strip()
            if line:
                try:
                    recs.append(json.loads(line))
                except json.JSONDecodeError:
                    pass
    return recs


def aggregate(recs):
    n = len(recs)
    per = {e: {"ok": 0, "with_findings": 0, "total": 0, "secs": 0.0,
               "finding_list": [], "owasp": {k: 0 for k in sc.ALL_OWASP}}
           for e in ENGINES}
    for r in recs:
        for e in ENGINES:
            d = r.get("engines", {}).get(e)
            if not d:
                continue
            if d.get("ok"):
                per[e]["ok"] += 1
            f = d.get("findings", 0)
            per[e]["total"] += f
            per[e]["secs"] += d.get("secs", 0.0)
            per[e]["finding_list"].append(f)
            if f > 0:
                per[e]["with_findings"] += 1
            for k in d.get("owasp", []):
                if k in per[e]["owasp"]:
                    per[e]["owasp"][k] += 1
    return n, per


def fmt_summary(n, per):
    lines = [f"\nCharts scanned: {n}\n",
             f"{'Engine':<14}{'ok%':>6}{'w/find':>8}{'findings':>10}{'mean':>7}"
             f"{'median':>8}{'OWASP ids':>12}  {'avg s':>7}",
             "-" * 80]
    for e in ENGINES:
        p = per[e]
        fl = p["finding_list"] or [0]
        okpct = 100 * p["ok"] / n if n else 0
        ids = sum(1 for k, v in p["owasp"].items() if v > 0)
        lines.append(f"{e:<14}{okpct:>5.0f}%{p['with_findings']:>8}{p['total']:>10}"
                     f"{statistics.mean(fl):>7.1f}{statistics.median(fl):>8.0f}"
                     f"{ids:>12}  {p['secs']/max(len(fl),1):>7.1f}")
    lines.append("")
    # OWASP id x engine detection matrix (charts where each engine flagged the id)
    lines.append(f"{'OWASP':<8}" + "".join(f"{SHORT[e]:>11}" for e in ENGINES))
    lines.append("-" * (8 + 11 * len(ENGINES)))
    for k in sc.ALL_OWASP:
        row = "".join(f"{per[e]['owasp'][k]:>11}" for e in ENGINES)
        lines.append(f"{k:<8}{row}")
    return "\n".join(lines) + "\n"


def emit_latex(n, per):
    L = [r"% Generated by thesis_assets/scripts/batch_scan.py agg — do not edit manually",
         f"% charts scanned: {n}",
         r"\begin{table}[htbp]", r"\centering\small",
         r"\caption{Engine comparison across the fetched Helm corpus"
         f" ({n} charts)}}\\label{{tab:engine-corpus-comparison}}",
         r"\begin{tabular}{@{}lrrrrr@{}}", r"\toprule",
         r"\textbf{Engine} & \textbf{ok\%} & \textbf{charts w/ finding} & "
         r"\textbf{total findings} & \textbf{mean/chart} & \textbf{OWASP ids} \\",
         r"\midrule"]
    for e in ENGINES:
        p = per[e]
        fl = p["finding_list"] or [0]
        okpct = 100 * p["ok"] / n if n else 0
        ids = sum(1 for v in p["owasp"].values() if v > 0)
        L.append(f"{e} & {okpct:.0f} & {p['with_findings']} & {p['total']} & "
                 f"{statistics.mean(fl):.1f} & {ids} \\\\")
    L += [r"\bottomrule", r"\end{tabular}", r"\end{table}", ""]
    # second table: OWASP id x engine detection matrix (chart counts)
    L += [r"\begin{table}[htbp]", r"\centering\small",
          r"\caption{OWASP Kubernetes Top 10 detection by engine (chart counts)}"
          r"\label{tab:engine-corpus-owasp}",
          r"\begin{tabular}{@{}l" + "r" * len(ENGINES) + r"@{}}", r"\toprule",
          r"\textbf{OWASP} & " + " & ".join(f"\\textbf{{{e}}}" for e in ENGINES) + r" \\",
          r"\midrule"]
    for k in sc.ALL_OWASP:
        L.append(f"{k} & " + " & ".join(str(per[e]["owasp"][k]) for e in ENGINES) + r" \\")
    L += [r"\bottomrule", r"\end{tabular}", r"\end{table}", ""]
    return "\n".join(L)


def cmd_agg(args):
    recs = read_records(args.out)
    n, per = aggregate(recs)
    print(fmt_summary(n, per))
    if args.latex:
        with open(args.latex, "w", encoding="utf-8", newline="\n") as f:
            f.write(emit_latex(n, per))
        print(f"Wrote LaTeX: {args.latex}", file=sys.stderr)
    if args.json:
        with open(args.json, "w", encoding="utf-8") as f:
            json.dump({"charts": n, "engines": {e: {k: v for k, v in per[e].items()
                       if k != "finding_list"} for e in ENGINES}}, f, indent=2)
        print(f"Wrote JSON: {args.json}", file=sys.stderr)


def main():
    ap = argparse.ArgumentParser(description=__doc__,
                                 formatter_class=argparse.RawDescriptionHelpFormatter)
    sub = ap.add_subparsers(dest="cmd", required=True)

    r = sub.add_parser("run", help="run the batch (resumable)")
    r.add_argument("--charts", required=True, help="chart-list file (one path per line)")
    r.add_argument("--out", required=True, help="JSONL results file (appended, resumable)")
    r.add_argument("--workers", type=int, default=1, help="parallel charts (default 1)")
    r.add_argument("--timeout", type=int, default=180, help="per-engine timeout seconds")
    r.add_argument("--engines", help="comma-separated subset of engines to run "
                   "(default all 5); e.g. kics-default,kics-enhanced")
    r.add_argument("--limit", type=int, default=0, help="scan at most N new charts this run")
    r.add_argument("--tmp", help="scratch dir for scanner output (default /tmp/batch-scan)")
    r.set_defaults(fn=cmd_run)

    a = sub.add_parser("agg", help="aggregate results into a comparison table")
    a.add_argument("--out", required=True, help="JSONL results file to read")
    a.add_argument("--latex", help="write LaTeX comparison tables to FILE")
    a.add_argument("--json", help="write aggregate JSON to FILE")
    a.set_defaults(fn=cmd_agg)

    args = ap.parse_args()
    args.fn(args)


if __name__ == "__main__":
    main()
