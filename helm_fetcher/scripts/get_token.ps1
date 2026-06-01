$tok = (gh auth token).Trim()
if (!$tok) {
    Write-Error 'gh auth token returned empty'
    exit 1
}
$lines = @(Get-Content .env -ErrorAction SilentlyContinue) -notmatch '^GITHUB_TOKEN='
$utf8NoBom = [System.Text.UTF8Encoding]::new($false)
[System.IO.File]::WriteAllLines((Join-Path (Get-Location).Path '.env'), ($lines + "GITHUB_TOKEN=$tok"), $utf8NoBom)
Write-Host 'Token saved to .env'
