#!/usr/bin/env bash
# Durable launcher for the corpus engine-comparison batch. Resumable: re-running
# this after any interruption continues where results.jsonl left off.
#
#   bash run.sh            # run/continue the batch (background-friendly)
#   bash run.sh agg        # aggregate current results into the comparison table
set -euo pipefail

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
SCRIPTS="$HERE/../../scripts"
BINS="$HOME/scan-bins"

# Persistent engine binaries (survive reboots, unlike /tmp builds).
export KICS_DEFAULT_BIN="$BINS/kicsv2/kics"
export KICS_DEFAULT_ASSETS="$BINS/kicsv2/assets"
export KICS_ENHANCED_BIN="$BINS/kics-fork/kics"
export KICS_ENHANCED_ASSETS="$BINS/kics-fork/assets"
export KUBESCAPE_BIN="$HOME/.kubescape/bin/kubescape"
export TRIVY_BIN="$(command -v trivy || echo "$HOME/.local/bin/trivy")"
export KUBELINTER_BIN="$(command -v kube-linter || echo "$HOME/.local/bin/kube-linter")"

CHARTS="$HERE/charts_all.txt"
OUT="$HERE/results.jsonl"

if [ "${1:-run}" = "agg" ]; then
  exec python3 "$SCRIPTS/batch_scan.py" agg --out "$OUT" \
      --latex "$HERE/engine_comparison.tex" --json "$HERE/engine_comparison.json"
fi

exec python3 "$SCRIPTS/batch_scan.py" run \
  --charts "$CHARTS" --out "$OUT" \
  --workers "${WORKERS:-4}" --timeout "${TIMEOUT:-180}" \
  ${ENGINES:+--engines "$ENGINES"} \
  --tmp /tmp/batch-scan
