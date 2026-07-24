#!/usr/bin/env python3
"""Run five static scanners over one Helm source, count each run's findings, and
report which OWASP Kubernetes Top 10 IDs each scanner detected.

The five runs are:

    kics-default    stock KICS v2.1.20 (no --experimental-helm-scan)
    kics-enhanced   fork KICS + --experimental-helm-scan (renders + fixer + dep-up)
    kubescape       kubescape scan <src>
    trivy           trivy config <src>
    kube-linter     kube-linter lint <src>

For every run the tool prints two values per scanner:

    findings   the raw count of results the scanner produced
    owasp      the set of OWASP Kubernetes Top 10 IDs (K01..K10) whose mapped
               rule fired, derived from the scanner rule ids in vulnerabilities.yaml

The rule id -> OWASP id mapping is the same single source of truth that
`gen-tables.py` uses for `tab:scanner-coverage`; a copy of that file lives next
to this script (`vulnerabilities.yaml`). Only the OWASP-mapped rules contribute
to the `owasp` column; `findings` counts everything the scanner reported.

Binaries and query assets are resolved from the environment (see BINARIES below),
falling back to $PATH and the build locations used while producing the thesis
`scan_out/` artifacts. Run with `--help` for options.

This is a WSL/Linux tool: it shells out to native scanner binaries. It does not
mutate the source tree — any `Chart.lock`/`charts/` a scanner vendors as a side
effect is removed afterward if it was not present before the run.
"""
import argparse
import json
import os
import re
import shutil
import subprocess
import sys
import tempfile

SCRIPT_DIR = os.path.dirname(os.path.abspath(__file__))
DATA = os.path.join(SCRIPT_DIR, "vulnerabilities.yaml")


# ── binary / asset resolution ────────────────────────────────────────────────
def _first(*cands):
    """Return the first candidate that exists or resolves on $PATH, else the last."""
    for c in cands:
        if not c:
            continue
        if os.path.sep in c:
            if os.path.exists(c):
                return c
        elif shutil.which(c):
            return shutil.which(c)
    return cands[-1]


def resolve_binaries():
    home = os.path.expanduser("~")
    return {
        "kics_default": os.environ.get("KICS_DEFAULT_BIN")
        or _first("/tmp/kicsv2/bin/kics", "kics"),
        "kics_default_assets": os.environ.get("KICS_DEFAULT_ASSETS", "/tmp/kicsv2/assets"),
        "kics_enhanced": os.environ.get("KICS_ENHANCED_BIN")
        or _first("/tmp/kics-fork/bin/kics", "kics"),
        "kics_enhanced_assets": os.environ.get("KICS_ENHANCED_ASSETS", "/tmp/kics-fork/assets"),
        "kubescape": os.environ.get("KUBESCAPE_BIN")
        or _first(os.path.join(home, ".kubescape/bin/kubescape"), "kubescape"),
        "trivy": os.environ.get("TRIVY_BIN") or _first("trivy", os.path.join(home, ".local/bin/trivy")),
        "kube_linter": os.environ.get("KUBELINTER_BIN")
        or _first("kube-linter", os.path.join(home, ".local/bin/kube-linter")),
    }


# ── vulnerabilities.yaml → rule maps ─────────────────────────────────────────
def parse_scanners(s):
    """Parse `{ kics: "...", trivy: KSV012, ... }` (borrowed from gen-tables.py)."""
    result = {}
    s = s.strip().lstrip("{").rstrip("}")
    for m in re.finditer(r'(\w+):\s*(?:"([^"]*?)"|([^,}]+?))\s*(?:,|$)', s):
        key = m.group(1)
        val = (m.group(2) if m.group(2) is not None else m.group(3)).strip()
        result[key] = val
    return result


def parse_vulns(path):
    entries, cur = [], {}
    with open(path, encoding="utf-8") as f:
        for raw in f:
            line = raw.rstrip()
            if line.startswith("- id:"):
                if cur:
                    entries.append(cur)
                cur = {"id": line.split(":", 1)[1].strip()}
            elif line.startswith("  ") and ":" in line and not line.strip().startswith("#"):
                key, _, val = line.strip().partition(":")
                val = val.strip().strip('"')
                cur[key] = parse_scanners(val) if key == "scanners" else val
    if cur:
        entries.append(cur)
    return entries


def norm_trivy(tok):
    """KSV012 / AVD-KSV-0109 / KSV-0111 -> canonical 'KSV<int>' (KCV likewise)."""
    out = set()
    for m in re.finditer(r"(KSV|KCV)-?0*(\d+)", tok, re.I):
        out.add(f"{m.group(1).upper()}{int(m.group(2))}")
    return out


def build_maps(entries):
    """rule-token -> owasp2025 id, per scanner. Also keep the fine k0X id."""
    kics, trivy, kube, ks = {}, {}, {}, {}
    for e in entries:
        owasp = e.get("owasp2025", "")
        vid = e.get("id", "")
        sc = e.get("scanners", {})
        if not owasp:
            continue
        payload = (owasp, vid)
        for tok in re.split(r"[,/\s]+", sc.get("kics", "") or ""):
            tok = tok.strip()
            if tok:
                kics[tok] = payload        # full uuid or 8-char id
                kics[tok[:8]] = payload    # 8-char prefix (table cells are 8-char)
        for tok in re.split(r"[/\s]+", sc.get("trivy", "") or ""):
            for canon in norm_trivy(tok):
                trivy[canon] = payload
        if sc.get("kubelinter"):
            kube[sc["kubelinter"].strip()] = payload
        for tok in re.split(r"[,/\s]+", sc.get("kubescape", "") or ""):
            tok = tok.strip()
            if tok:
                ks[tok] = payload
    return {"kics": kics, "trivy": trivy, "kube_linter": kube, "kubescape": ks}


# ── scanner runners: return (findings_count, set_of_fired_rule_tokens) ─────────
def _run(cmd, timeout=None, **kw):
    """Run cmd; return the CompletedProcess, or None on timeout/spawn failure."""
    try:
        return subprocess.run(cmd, capture_output=True, text=True, timeout=timeout, **kw)
    except (subprocess.TimeoutExpired, OSError):
        return None


# Each runner returns (findings:int, fired:set[str], ok:bool). `ok` means the
# tool completed and produced parseable output (even if it found nothing);
# ok=False means it timed out, crashed, or wrote no report.
def run_kics(binary, assets, src, outdir, enhanced, timeout=None):
    name = "kics-enhanced" if enhanced else "kics-default"
    out = os.path.join(outdir, name)
    os.makedirs(out, exist_ok=True)
    cmd = [binary, "scan", "-p", src,
           "-q", os.path.join(assets, "queries"),
           "-b", os.path.join(assets, "libraries"),
           "-o", out, "--output-name", "r", "--report-formats", "json"]
    if enhanced:
        cmd.append("--experimental-helm-scan")
    _run(cmd, timeout=timeout)
    jf = os.path.join(out, "r.json")
    if not os.path.exists(jf):
        return 0, set(), False
    try:
        d = json.load(open(jf))
    except (json.JSONDecodeError, ValueError):
        return 0, set(), False
    fired = set()
    for q in d.get("queries", []):
        qid = q.get("query_id", "")
        fired.add(qid)
        fired.add(qid[:8])
    return d.get("total_counter", 0), fired, True


def run_trivy(binary, src, outdir, timeout=None):
    jf = os.path.join(outdir, "trivy.json")
    _run([binary, "config", src, "--format", "json", "--output", jf], timeout=timeout)
    if not os.path.exists(jf):
        return 0, set(), False
    try:
        d = json.load(open(jf))
    except (json.JSONDecodeError, ValueError):
        return 0, set(), False
    count, fired = 0, set()
    for r in d.get("Results", []):
        for m in r.get("Misconfigurations", []):
            count += 1
            for tok in (m.get("ID"), m.get("AVDID")):
                fired |= norm_trivy(tok or "")
    return count, fired, True


def run_kube_linter(binary, src, outdir, timeout=None):
    p = _run([binary, "lint", src, "--format", "json"], timeout=timeout)
    if p is None:
        return 0, set(), False
    try:
        d = json.loads(p.stdout)
    except (json.JSONDecodeError, ValueError):
        return 0, set(), True    # "no valid objects found" -> ran, empty stdout
    reports = d.get("Reports") or []
    fired = {r.get("Check") for r in reports if r.get("Check")}
    return len(reports), fired, True


def run_kubescape(binary, src, outdir, timeout=None):
    jf = os.path.join(outdir, "kubescape.json")
    env = dict(os.environ, KUBESCAPE_SKIP_UPDATE_CHECK="true")
    _run([binary, "scan", src, "--format", "json", "--output", jf],
         timeout=timeout, env=env)
    if not os.path.exists(jf):
        return 0, set(), False
    try:
        d = json.load(open(jf))
    except (json.JSONDecodeError, ValueError):
        return 0, set(), False
    controls = d.get("summaryDetails", {}).get("controls", {})
    count, fired = 0, set()
    for cid, c in controls.items():
        failed = (c.get("ResourceCounters") or {}).get("failedResources", 0)
        if failed > 0:
            count += failed
            fired.add(c.get("controlID") or cid)
    return count, fired, True


# ── source-tree protection ───────────────────────────────────────────────────
def snapshot_untracked(src):
    """Record chart-root Chart.lock/charts that don't exist yet, so we can undo
    a scanner's dep-up side effect afterward."""
    roots = []
    if os.path.exists(os.path.join(src, "Chart.yaml")):
        roots.append(src)
    for name in sorted(os.listdir(src)) if os.path.isdir(src) else []:
        d = os.path.join(src, name)
        if os.path.isdir(d) and os.path.exists(os.path.join(d, "Chart.yaml")):
            roots.append(d)
    pre = set()
    for r in roots:
        for art in ("Chart.lock", "charts"):
            if os.path.exists(os.path.join(r, art)):
                pre.add(os.path.join(r, art))
    return roots, pre


def restore_untracked(roots, pre):
    for r in roots:
        for art in ("Chart.lock", "charts"):
            p = os.path.join(r, art)
            if p not in pre and os.path.exists(p):
                shutil.rmtree(p) if os.path.isdir(p) else os.remove(p)


# ── OWASP set from fired tokens ──────────────────────────────────────────────
ALL_OWASP = [f"K{i:02d}" for i in range(1, 11)]


def owasp_hits(fired, rule_map):
    owasp, vulns = set(), set()
    for tok in fired:
        if tok in rule_map:
            o, v = rule_map[tok]
            owasp.add(o)
            vulns.add(v)
    return owasp, vulns


# ── reporting ────────────────────────────────────────────────────────────────
def fmt_owasp(s):
    return ", ".join(sorted(s)) if s else "—"


def print_table(rows):
    w = max(len(r["scanner"]) for r in rows)
    print(f"\n{'Scanner'.ljust(w)}  {'Findings':>8}  OWASP K-Top-10 detected")
    print("-" * (w + 12 + 32))
    for r in rows:
        print(f"{r['scanner'].ljust(w)}  {r['findings']:>8}  {fmt_owasp(r['owasp'])}")
    print()


def emit_latex(rows, src):
    lines = [
        "% Generated by thesis_assets/scripts/scan_coverage.py — do not edit manually",
        f"% source: {src}",
        r"\begin{table}[htbp]",
        r"\centering\small",
        r"\caption{Scanner findings and OWASP Kubernetes Top 10 coverage}\label{tab:scanner-owasp-coverage}",
        r"\begin{tabular}{@{}lrl@{}}",
        r"\toprule",
        r"\textbf{Scanner} & \textbf{Findings} & \textbf{OWASP K-Top-10 detected} \\",
        r"\midrule",
    ]
    for r in rows:
        owasp = ", ".join(f"\\texttt{{{o}}}" for o in sorted(r["owasp"])) or "---"
        lines.append(f"{r['scanner']} & {r['findings']} & {owasp} \\\\")
    lines += [r"\bottomrule", r"\end{tabular}", r"\end{table}", ""]
    return "\n".join(lines)


# ── main ─────────────────────────────────────────────────────────────────────
def main():
    ap = argparse.ArgumentParser(description=__doc__,
                                 formatter_class=argparse.RawDescriptionHelpFormatter)
    ap.add_argument("source", help="path to the Helm chart / directory to scan")
    ap.add_argument("--only", help="comma-separated subset of: "
                    "kics-default,kics-enhanced,kubescape,trivy,kube-linter")
    ap.add_argument("--latex", metavar="FILE", help="also write a LaTeX table to FILE")
    ap.add_argument("--json", metavar="FILE", help="also write the raw result as JSON")
    ap.add_argument("--keep-scan-out", metavar="DIR",
                    help="keep raw scanner JSON under DIR instead of a temp dir")
    args = ap.parse_args()

    src = os.path.abspath(args.source)
    if not os.path.exists(src):
        sys.exit(f"scan_coverage: source not found: {src}")

    bins = resolve_binaries()
    maps = build_maps(parse_vulns(DATA))
    selected = set((args.only or
                    "kics-default,kics-enhanced,kubescape,trivy,kube-linter").split(","))

    outdir = args.keep_scan_out or tempfile.mkdtemp(prefix="scan-cov-")
    os.makedirs(outdir, exist_ok=True)
    roots, pre = snapshot_untracked(src)

    runners = [
        ("kics-default", "kics", lambda: run_kics(
            bins["kics_default"], bins["kics_default_assets"], src, outdir, False)),
        ("kics-enhanced", "kics", lambda: run_kics(
            bins["kics_enhanced"], bins["kics_enhanced_assets"], src, outdir, True)),
        ("kubescape", "kubescape", lambda: run_kubescape(bins["kubescape"], src, outdir)),
        ("trivy", "trivy", lambda: run_trivy(bins["trivy"], src, outdir)),
        ("kube-linter", "kube_linter", lambda: run_kube_linter(bins["kube_linter"], src, outdir)),
    ]

    rows = []
    try:
        for name, mapkey, fn in runners:
            if name not in selected:
                continue
            print(f"  running {name} …", file=sys.stderr)
            count, fired, ok = fn()
            owasp, vulns = owasp_hits(fired, maps[mapkey])
            rows.append({"scanner": name, "findings": count, "ok": ok,
                         "owasp": owasp, "vulns": vulns})
            restore_untracked(roots, pre)   # undo any dep-up between runs
    finally:
        restore_untracked(roots, pre)
        if not args.keep_scan_out:
            shutil.rmtree(outdir, ignore_errors=True)

    print_table(rows)
    # coverage roll-up across all scanners
    union = set().union(*(r["owasp"] for r in rows)) if rows else set()
    missed = [k for k in ALL_OWASP if k not in union]
    print(f"Union of OWASP IDs detected by any scanner: {fmt_owasp(union)}")
    print(f"Chart-scope OWASP IDs not detected by any scanner: "
          f"{', '.join(missed) if missed else '—'}\n")

    if args.latex:
        with open(args.latex, "w", encoding="utf-8", newline="\n") as f:
            f.write(emit_latex(rows, os.path.relpath(src)))
        print(f"Wrote LaTeX table: {args.latex}")
    if args.json:
        payload = [{"scanner": r["scanner"], "findings": r["findings"],
                    "owasp": sorted(r["owasp"]), "vulns": sorted(r["vulns"])}
                   for r in rows]
        with open(args.json, "w", encoding="utf-8") as f:
            json.dump({"source": src, "results": payload}, f, indent=2)
        print(f"Wrote JSON: {args.json}")


if __name__ == "__main__":
    main()
