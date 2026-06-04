$ErrorActionPreference = "Stop"
$projectRoot = Split-Path -Parent $PSScriptRoot
$venvPython = Join-Path $projectRoot ".venv\Scripts\python.exe"
$distName = "Dynon2Savvy"

Write-Host "Ensuring PyInstaller is installed..."
& $venvPython -m pip install pyinstaller --quiet

Write-Host "Cleaning previous builds..."
if (Test-Path "build") { Remove-Item -Recurse -Force "build" }
if (Test-Path "dist") { Remove-Item -Recurse -Force "dist" }

Write-Host "Running PyInstaller using spec from packaging/..."
& $venvPython -m PyInstaller --clean packaging/Dynon2Savvy.spec

$distDir = Join-Path $projectRoot "dist\$distName"

# Post-process: ensure d2sCLI.exe and custom-mappings.json are at root (for the app's path logic)
Copy-Item -Force (Join-Path $projectRoot "d2sCLI.exe") (Join-Path $distDir "d2sCLI.exe")
Copy-Item -Force (Join-Path $projectRoot "custom-mappings.json") (Join-Path $distDir "custom-mappings.json")

# Remove from _internal if duplicated there
$internal = Join-Path $distDir "_internal"
if (Test-Path (Join-Path $internal "d2sCLI.exe")) { Remove-Item (Join-Path $internal "d2sCLI.exe") -Force }
if (Test-Path (Join-Path $internal "custom-mappings.json")) { Remove-Item (Join-Path $internal "custom-mappings.json") -Force }

# Explicitly ensure no d2sAppState.json is shipped (app creates its own)
$stateFile = Join-Path $distDir "d2sAppState.json"
if (Test-Path $stateFile) {
    Remove-Item $stateFile -Force
    Write-Host "Removed d2sAppState.json (app will create clean one on first run)"
}

Write-Host ""
Write-Host "Build complete. Package folder: $distDir"
Write-Host "To create release zip: Compress-Archive -Path $distDir -DestinationPath (Join-Path $projectRoot 'dist\Dynon2Savvy-v2.0.0-win-x64.zip')"
