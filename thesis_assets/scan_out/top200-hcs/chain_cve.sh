#!/usr/bin/env bash
# Waits for the image inventory (phase 1) to reach all 9861 charts, then runs
# the trivy CVE scan (phase 2) unattended.
set -uo pipefail
export PATH=$HOME/.local/bin:$HOME/go/bin:$PATH
export TRIVY_BIN=$(command -v trivy)
HCS=/mnt/c/Users/miabs/GolandProjects/helm-tests/thesis_assets/scan_out/top200-hcs
SC=/mnt/c/Users/miabs/GolandProjects/helm-tests/thesis_assets/scripts/hcs_corpus.py
echo "[chain] waiting for phase-1 inventory to complete ($(date))"
while [ "$(wc -l < "$HCS/images_full.jsonl" 2>/dev/null || echo 0)" -lt 9861 ]; do sleep 30; done
echo "[chain] inventory complete: $(wc -l < "$HCS/images_full.jsonl") charts. unique images:"
python3 "$SC" agg --images "$HCS/images_full.jsonl"
echo "[chain] launching phase-2 CVE scan ($(date))"
python3 "$SC" cve --images "$HCS/images_full.jsonl" --out "$HCS/cve_full.jsonl" \
   --workers 6 --timeout 300 --tmp /tmp/hcs-cve
echo "[chain] CVE scan complete ($(date))"
python3 "$SC" agg --images "$HCS/images_full.jsonl" --cve "$HCS/cve_full.jsonl"
