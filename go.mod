# NoMercyChain Minimal Build Script
# This PowerShell script builds a minimal version of the NoMercyChain binary

Write-Host "Building NoMercyChain (Minimal)..." -ForegroundColor Green

# Create build directory if it doesn't exist
New-Item -ItemType Directory -Force -Path build | Out-Null

# Create a temporary directory
$tempDir = "temp_build"
if (Test-Path $tempDir) {
    Remove-Item -Recurse -Force $tempDir
}
New-Item -ItemType Directory -Force -Path $tempDir | Out-Null

# Change to the temporary directory
Push-Location $tempDir

# Create a minimal go.mod file
@"
module nmxchain

go 1.20

require github.com/spf13/cobra v1.7.0

require (
github.com/inconshreveable/mousetrap v1.1.0 // indirect
github.com/spf13/pflag v1.0.5 // indirect
)
