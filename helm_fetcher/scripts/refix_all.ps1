<#
.SYNOPSIS
  Re-run the helm fixer over both sources (github + artifacthub) and rebuild the
  merged fixed catalog in one shot.

.DESCRIPTION
  github and artifacthub must run separately because each uses its own clone
  directory and the fixer always writes catalog_fixed.json. This script runs
  each in turn, snapshots the per-source result, then merges them into the
  cumulative fixed catalog via scripts/merge_fixed (no sed/jq needed).

  IMPORTANT: if you changed helmfix/fix.go, refresh the vendored copies first
  (cp helmfix/fix.go <module>/vendor/.../helmfix/fix.go) or the -mod=vendor
  build silently uses the OLD parser.

  Run from the helm_fetcher directory (the Makefile target does this for you).
#>
param(
  [string]$GithubCatalog   = "results/github/catalog_by_project.json",
  [string]$GithubCloneDir  = "D:/helm_clones_github",
  [string]$ArtifactCatalog = "runs/20260529_135859_artifacthub/catalog_by_project.json",
  [string]$ArtifactCloneDir = "D:/helm_clones_artifacthub",
  [string]$GithubFixed     = "catalog_fixed_github_final.json",
  [string]$ArtifactFixed   = "catalog_fixed_ah_final.json",
  [string]$Cumulative      = "catalog_fixed_cumulative.json",
  # Taxonomy regen (runs unless -SkipAnalyze is passed):
  [string]$SourcesCatalog  = "catalog_sources_merged.json",
  [string]$AnalyzeOut      = "out/cumulative_final",
  [switch]$SkipAnalyze
)

$ErrorActionPreference = "Stop"

function Invoke-Fixer($catalog, $cloneDir, $dest) {
  Write-Host "=== fixer: $catalog (clones: $cloneDir) ===" -ForegroundColor Cyan
  go run . -mode fixer -catalog-in $catalog -clone-dir $cloneDir
  if ($LASTEXITCODE -ne 0) { throw "fixer failed for $catalog (exit $LASTEXITCODE)" }
  Copy-Item -Force catalog_fixed.json $dest
  Write-Host "  -> $dest" -ForegroundColor Green
}

Invoke-Fixer $GithubCatalog   $GithubCloneDir   $GithubFixed
Invoke-Fixer $ArtifactCatalog $ArtifactCloneDir $ArtifactFixed

Write-Host "=== merge -> $Cumulative ===" -ForegroundColor Cyan
go run ./scripts/merge_fixed $Cumulative $GithubFixed $ArtifactFixed
if ($LASTEXITCODE -ne 0) { throw "merge failed (exit $LASTEXITCODE)" }

if ($SkipAnalyze) {
  Write-Host ""
  Write-Host "Skipping taxonomy regen. To regenerate the report + viewer later:" -ForegroundColor Yellow
  Write-Host "  cd ../taxonomy_analyzer; make analyze-cumulative" -ForegroundColor Yellow
  return
}

# Resolve to absolute paths so the regen works regardless of the cwd switch.
$srcAbs = (Resolve-Path $SourcesCatalog).Path
$cumAbs = (Resolve-Path $Cumulative).Path

Write-Host "=== taxonomy regen -> ../taxonomy_analyzer/$AnalyzeOut ===" -ForegroundColor Cyan
Push-Location ../taxonomy_analyzer
try {
  go run . -input $srcAbs --fixed $cumAbs -out $AnalyzeOut
  if ($LASTEXITCODE -ne 0) { throw "taxonomy regen failed (exit $LASTEXITCODE)" }
} finally {
  Pop-Location
}

Write-Host ""
Write-Host "Done. Report + viewer regenerated at taxonomy_analyzer/$AnalyzeOut" -ForegroundColor Green
