#!/usr/bin/env python3
"""Corpus-scale HCS pipeline in two phases, mirroring `hcs`:

  phase 1  images : run `kics --experimental-helm-scan --image-bom` on every chart,
                    collect the CycloneDX image inventory (one JSONL record/chart).
  phase 2  cve    : dedupe the images, run `trivy image <ref> --scanners vuln` once
                    per unique image, and sum the CVEs by severity.
  agg             : report unique-image and CVE totals.

Phase 1 is a render pass (cost ~ the enhanced-KICS run). Phase 2's cost is a
`trivy image` pull+scan per *unique* image — the expensive, network/disk-heavy
part — so `images` supports `--limit` to size it on a sample first.

Binaries via env: KICS_IMGBOM_BIN, KICS_IMGBOM_ASSETS, TRIVY_BIN.
"""
import argparse
import json
import os
import subprocess
import sys
import tempfile
import threading
import time
from concurrent.futures import ThreadPoolExecutor, as_completed


def _run(cmd, timeout=None, **kw):
    try:
        return subprocess.run(cmd, capture_output=True, text=True, timeout=timeout, **kw)
    except (subprocess.TimeoutExpired, OSError):
        return None


def bins():
    return {
        "kics": os.environ.get("KICS_IMGBOM_BIN", "kics"),
        "assets": os.environ.get("KICS_IMGBOM_ASSETS", "assets"),
        "trivy": os.environ.get("TRIVY_BIN", "trivy"),
    }


def load_done(path, key="chart"):
    done = set()
    if os.path.exists(path):
        for line in open(path, encoding="utf-8"):
            line = line.strip()
            if line:
                try:
                    done.add(json.loads(line)[key])
                except (json.JSONDecodeError, KeyError):
                    pass
    return done


# ── phase 1: image inventory ─────────────────────────────────────────────────
def kics_images(b, chart, timeout, tmp_root):
    import shutil
    out = tempfile.mkdtemp(prefix="hcs-", dir=tmp_root)
    # protect the frozen corpus from the fixer's dep-up
    pre_lock = os.path.exists(os.path.join(chart, "Chart.lock"))
    pre_charts = os.path.exists(os.path.join(chart, "charts"))
    try:
        _run([b["kics"], "scan", "-p", chart, "--experimental-helm-scan", "--image-bom",
              "-q", os.path.join(b["assets"], "queries"),
              "-b", os.path.join(b["assets"], "libraries"),
              "-o", out, "--output-name", "r", "--report-formats", "json"], timeout=timeout)
        bom = os.path.join(out, "kics-image-bom.json")
        imgs, ok = [], False
        if os.path.exists(bom):
            ok = True
            try:
                d = json.load(open(bom))
                for c in d.get("components", []) or []:
                    name = c.get("name", "")
                    ver = c.get("version", "")
                    imgs.append(f"{name}:{ver}" if ver else name)
            except (json.JSONDecodeError, ValueError):
                ok = False
        return {"chart": chart, "ok": ok, "images": sorted(set(imgs))}
    finally:
        if not pre_lock and os.path.exists(os.path.join(chart, "Chart.lock")):
            os.remove(os.path.join(chart, "Chart.lock"))
        if not pre_charts and os.path.isdir(os.path.join(chart, "charts")):
            shutil.rmtree(os.path.join(chart, "charts"), ignore_errors=True)
        shutil.rmtree(out, ignore_errors=True)


def cmd_images(a):
    b = bins()
    charts = [l.strip() for l in open(a.charts, encoding="utf-8") if l.strip()]
    done = load_done(a.out)
    todo = [c for c in charts if c not in done]
    if a.limit:
        todo = todo[:a.limit]
    tmp = a.tmp or tempfile.gettempdir()
    os.makedirs(tmp, exist_ok=True)
    print(f"images: {len(charts)} charts, {len(done)} done, {len(todo)} to do "
          f"({a.workers} workers)", file=sys.stderr)
    lock = threading.Lock()
    out = open(a.out, "a", encoding="utf-8")
    n = 0
    t0 = time.time()

    def write(rec):
        nonlocal n
        with lock:
            out.write(json.dumps(rec) + "\n"); out.flush(); os.fsync(out.fileno())
            n += 1
            if n % 50 == 0 or n == len(todo):
                r = n / max(time.time() - t0, 1)
                print(f"  {n}/{len(todo)} ({r*3600:.0f}/h)", file=sys.stderr)
    try:
        with ThreadPoolExecutor(max_workers=a.workers) as ex:
            futs = {ex.submit(kics_images, b, c, a.timeout, tmp): c for c in todo}
            for f in as_completed(futs):
                try:
                    write(f.result())
                except Exception as e:
                    write({"chart": futs[f], "ok": False, "images": [], "err": str(e)[:150]})
    finally:
        out.close()
    print(f"images: wrote {n} records", file=sys.stderr)


def unique_images(images_jsonl):
    imgs = set()
    for line in open(images_jsonl, encoding="utf-8"):
        line = line.strip()
        if line:
            try:
                imgs.update(json.loads(line).get("images", []))
            except json.JSONDecodeError:
                pass
    # drop obviously unscannable refs (templated leftovers, empty tags)
    return sorted(i for i in imgs if i and "{{" not in i and not i.startswith(":"))


# ── phase 2: CVE scan per unique image ───────────────────────────────────────
SEV = ["CRITICAL", "HIGH", "MEDIUM", "LOW", "UNKNOWN"]


def trivy_cve(b, ref, timeout, tmp_root):
    out = os.path.join(tmp_root, "t.json")
    p = _run([b["trivy"], "image", ref, "--format", "json", "--scanners", "vuln",
              "--quiet", "--output", out], timeout=timeout)
    counts = {s: 0 for s in SEV}
    if not os.path.exists(out):
        return {"image": ref, "ok": False, "counts": counts, "total": 0}
    try:
        d = json.load(open(out))
    except (json.JSONDecodeError, ValueError):
        return {"image": ref, "ok": False, "counts": counts, "total": 0}
    for res in d.get("Results", []) or []:
        for v in res.get("Vulnerabilities", []) or []:
            s = (v.get("Severity") or "UNKNOWN").upper()
            counts[s] = counts.get(s, 0) + 1
    try:
        os.remove(out)
    except OSError:
        pass
    return {"image": ref, "ok": True, "counts": counts, "total": sum(counts.values())}


def cmd_cve(a):
    b = bins()
    imgs = unique_images(a.images)
    if a.limit:
        imgs = imgs[:a.limit]
    done = load_done(a.out, key="image")
    todo = [i for i in imgs if i not in done]
    tmp = a.tmp or tempfile.mkdtemp(prefix="hcs-cve-")
    os.makedirs(tmp, exist_ok=True)
    print(f"cve: {len(imgs)} unique images, {len(done)} done, {len(todo)} to scan "
          f"({a.workers} workers)", file=sys.stderr)
    lock = threading.Lock()
    out = open(a.out, "a", encoding="utf-8")
    n = 0
    t0 = time.time()

    def write(rec):
        nonlocal n
        with lock:
            out.write(json.dumps(rec) + "\n"); out.flush(); os.fsync(out.fileno())
            n += 1
            if n % 10 == 0 or n == len(todo):
                r = n / max(time.time() - t0, 1)
                eta = (len(todo) - n) / r / 3600 if r else 0
                print(f"  {n}/{len(todo)} ({r*3600:.0f}/h, ETA {eta:.1f}h)", file=sys.stderr)
    try:
        with ThreadPoolExecutor(max_workers=a.workers) as ex:
            wtmp = {}
            futs = {}
            for i, img in enumerate(todo):
                d = os.path.join(tmp, str(i % a.workers))
                os.makedirs(d, exist_ok=True)
                futs[ex.submit(trivy_cve, b, img, a.timeout, d)] = img
            for f in as_completed(futs):
                try:
                    write(f.result())
                except Exception as e:
                    write({"image": futs[f], "ok": False, "counts": {}, "total": 0,
                           "err": str(e)[:150]})
    finally:
        out.close()
    print(f"cve: wrote {n} records", file=sys.stderr)


def cmd_agg(a):
    n_charts = ok_charts = 0
    if a.images and os.path.exists(a.images):
        for line in open(a.images, encoding="utf-8"):
            if line.strip():
                n_charts += 1
                try:
                    ok_charts += 1 if json.loads(line).get("ok") else 0
                except json.JSONDecodeError:
                    pass
        uniq = unique_images(a.images)
        print(f"charts inventoried : {n_charts} ({ok_charts} ok)")
        print(f"unique images      : {len(uniq)}")
    if a.cve and os.path.exists(a.cve):
        tot = {s: 0 for s in SEV}
        scanned = ok = grand = 0
        for line in open(a.cve, encoding="utf-8"):
            line = line.strip()
            if not line:
                continue
            r = json.loads(line)
            scanned += 1
            if r.get("ok"):
                ok += 1
            for s, c in (r.get("counts") or {}).items():
                tot[s] = tot.get(s, 0) + c
            grand += r.get("total", 0)
        print(f"images CVE-scanned : {scanned} ({ok} ok)")
        print(f"CVEs by severity   : " + "  ".join(f"{s[:4]}={tot[s]}" for s in SEV))
        print(f"TOTAL CVEs         : {grand}")


def main():
    ap = argparse.ArgumentParser(description=__doc__,
                                 formatter_class=argparse.RawDescriptionHelpFormatter)
    sub = ap.add_subparsers(dest="cmd", required=True)

    im = sub.add_parser("images", help="phase 1: inventory images per chart")
    im.add_argument("--charts", required=True)
    im.add_argument("--out", required=True)
    im.add_argument("--workers", type=int, default=12)
    im.add_argument("--timeout", type=int, default=180)
    im.add_argument("--limit", type=int, default=0)
    im.add_argument("--tmp")
    im.set_defaults(fn=cmd_images)

    cv = sub.add_parser("cve", help="phase 2: trivy CVE scan per unique image")
    cv.add_argument("--images", required=True, help="phase-1 JSONL")
    cv.add_argument("--out", required=True, help="CVE JSONL")
    cv.add_argument("--workers", type=int, default=4)
    cv.add_argument("--timeout", type=int, default=300)
    cv.add_argument("--limit", type=int, default=0)
    cv.add_argument("--tmp")
    cv.set_defaults(fn=cmd_cve)

    ag = sub.add_parser("agg", help="report totals")
    ag.add_argument("--images")
    ag.add_argument("--cve")
    ag.set_defaults(fn=cmd_agg)

    args = ap.parse_args()
    args.fn(args)


if __name__ == "__main__":
    main()
