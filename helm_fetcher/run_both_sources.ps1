<#
.SYNOPSIS
    Runs helm_fetcher for both GitHub and ArtifactHub sources sequentially.

.DESCRIPTION
    Each source gets its own clone directory and results subdirectory.
    Output streams live to the terminal AND is captured to per-source log files.

    Results layout after the run:
        results/
          github/
            run.log              -- full go output
            catalog_*.json/md/csv
          artifacthub/
            run.log
            catalog_*.json/md/csv
          run_<timestamp>.log    -- high-level progress summary

.USAGE
    cd helm_fetcher
    .\run_both_sources.ps1

    Override defaults:
    .\run_both_sources.ps1 -Top 300 -GithubCloneDir D:/gh_clones -ArtifactCloneDir D:/ah_clones
#>

param(
    [int]   $Top              = 500,
    [string]$GithubCloneDir   = "D:/helm_clones_github",
    [string]$ArtifactCloneDir = "D:/helm_clones_artifacthub"
)

$ErrorActionPreference = "Continue"

$ScriptDir = Split-Path -Parent $MyInvocation.MyCommand.Definition
Set-Location $ScriptDir

# Ensure git, helm, and curl are reachable — add common Windows install paths if missing.
$extraPaths = @(
    "C:\Program Files\Git\cmd",
    "C:\Program Files\Git\bin",
    "C:\Program Files (x86)\Git\cmd",
    "C:\Program Files\Helm",
    "C:\ProgramData\chocolatey\bin"
)
foreach ($p in $extraPaths) {
    if ((Test-Path $p) -and ($env:PATH -notlike "*$p*")) {
        $env:PATH = "$p;$env:PATH"
    }
}

# Verify required tools are now available
foreach ($tool in @("git", "helm", "go")) {
    if (-not (Get-Command $tool -ErrorAction SilentlyContinue)) {
        Write-Host "[ERROR] '$tool' not found in PATH. Install it and retry." -ForegroundColor Red
        exit 1
    }
}

$ResultsRoot = Join-Path $ScriptDir "results"
$GithubDir   = Join-Path $ResultsRoot "github"
$ArtifactDir = Join-Path $ResultsRoot "artifacthub"

New-Item -ItemType Directory -Force -Path $GithubDir   | Out-Null
New-Item -ItemType Directory -Force -Path $ArtifactDir | Out-Null

$RunTs   = Get-Date -Format "yyyyMMdd_HHmmss"
$LogFile = Join-Path $ResultsRoot "run_$RunTs.log"

$CatalogFiles = @(
    "catalog_by_project.json", "catalog_kept.json", "catalog_removed.json",
    "catalog_dep_failures.json", "catalog_results.md", "catalog_kept.md",
    "catalog_removed.md", "catalog_dep_failures.md",
    "catalog_failures.csv", "catalog_kept_failures.csv", "catalog_dep_failures.csv"
)

function Write-Log([string]$Msg) {
    $line = "[$(Get-Date -Format 'yyyy-MM-dd HH:mm:ss')] $Msg"
    Write-Host $line -ForegroundColor Cyan
    Add-Content -Path $LogFile -Value $line -Encoding utf8
}

function Move-CatalogFiles([string]$DestDir) {
    foreach ($f in $CatalogFiles) {
        $src = Join-Path $ScriptDir $f
        if (Test-Path $src) {
            Move-Item -Force $src (Join-Path $DestDir $f)
        }
    }
}

function Invoke-Source {
    param(
        [string]$Label,
        [string]$Source,
        [string]$SearchIn,
        [string]$CloneDir,
        [string]$ResultsDir
    )

    Write-Log ""
    Write-Log ">>> [$Label]  source=$Source  search_in=$SearchIn  top=$Top  clone_dir=$CloneDir"
    $start   = Get-Date
    $runLog  = Join-Path $ResultsDir "run.log"

    # Run go directly so output streams live to the terminal.
    # 2>&1 merges stderr (where zerolog writes) into the pipeline so Tee-Object captures both.
    & go run . `
        -mode=full `
        "-source=$Source" `
        "-top=$Top" `
        "-search-in=$SearchIn" `
        "-clone-dir=$CloneDir" 2>&1 | Tee-Object -FilePath $runLog

    $exitCode  = $LASTEXITCODE
    $elapsed   = (Get-Date) - $start
    $elapsedStr = $elapsed.ToString("hh\:mm\:ss")

    if ($exitCode -eq 0) {
        Write-Log "<<< [$Label] DONE  exit=0  elapsed=$elapsedStr"
    } else {
        Write-Log "<<< [$Label] FAILED  exit=$exitCode  elapsed=$elapsedStr  (see $runLog)"
    }

    Move-CatalogFiles -DestDir $ResultsDir
    Write-Log "    Catalog files saved to: $ResultsDir"
}

# ---------------------------------------------------------------------------
$totalStart = Get-Date
Write-Log "================================================================"
Write-Log "helm_fetcher overnight run   top=$Top   started=$RunTs"
Write-Log "================================================================"

Invoke-Source `
    -Label      "1/2  GitHub" `
    -Source     "github" `
    -SearchIn   "github_search.json" `
    -CloneDir   $GithubCloneDir `
    -ResultsDir $GithubDir

Invoke-Source `
    -Label      "2/2  ArtifactHub" `
    -Source     "artifacthub" `
    -SearchIn   "artifacthub_search.json" `
    -CloneDir   $ArtifactCloneDir `
    -ResultsDir $ArtifactDir

$totalElapsed = ((Get-Date) - $totalStart).ToString("hh\:mm\:ss")
Write-Log ""
Write-Log "================================================================"
Write-Log "All done.  Total elapsed: $totalElapsed"
Write-Log "  GitHub results:      $GithubDir"
Write-Log "  ArtifactHub results: $ArtifactDir"
Write-Log "  This log:            $LogFile"
Write-Log "================================================================"
