# Native-Windows launcher for the two-KICS top-200 batch (no WSL, no 9p).
# Resumable: re-running continues from results_win.jsonl.
#   powershell -File run_win.ps1          # run / continue
#   powershell -File run_win.ps1 -Agg     # aggregate the comparison table
param([switch]$Agg)

$ErrorActionPreference = "Stop"
$here    = Split-Path -Parent $MyInvocation.MyCommand.Path
$scripts = Join-Path $here "..\..\scripts"
$bins    = "C:\Users\miabs\scan-bins-win"

$env:KICS_DEFAULT_BIN    = "$bins\kicsv2\kics.exe"
$env:KICS_DEFAULT_ASSETS = "$bins\kicsv2\assets"
$env:KICS_ENHANCED_BIN   = "$bins\kics-fork\kics.exe"
$env:KICS_ENHANCED_ASSETS= "$bins\kics-fork\assets"

$py = if (Get-Command py -ErrorAction SilentlyContinue) { "py" } else { "python" }
$charts = Join-Path $here "charts_top200_win.txt"
$out    = Join-Path $here "results_win.jsonl"
$tmp    = Join-Path $env:TEMP "batch-top200-win"

if ($Agg) {
  & $py "$scripts\batch_scan.py" agg --out $out `
      --latex (Join-Path $here "engine_comparison_win.tex") `
      --json  (Join-Path $here "engine_comparison_win.json")
  return
}

& $py "$scripts\batch_scan.py" run --charts $charts --out $out `
   --engines kics-default,kics-enhanced --workers 12 --timeout 180 --tmp $tmp
