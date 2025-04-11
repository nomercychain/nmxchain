# NoMercyChain Simple Build Script
# This PowerShell script builds the NoMercyChain binary with minimal dependencies

Write-Host "Building NoMercyChain (Simplified)..." -ForegroundColor Green

# Create build directory if it doesn't exist
New-Item -ItemType Directory -Force -Path build | Out-Null

# Set environment variables to skip problematic imports
$env:GO111MODULE = "on"

# Update go.mod
Write-Host "Updating go.mod..." -ForegroundColor Cyan
go mod tidy

# Build the binary with minimal features
go build -tags "netgo ledger" -ldflags "-s -w -X github.com/cosmos/cosmos-sdk/version.Name=nmxchain -X github.com/cosmos/cosmos-sdk/version.AppName=nmxchaind" -o build\nmxchaind.exe .\cmd\nmxchaind

if (Test-Path "build\nmxchaind.exe") {
    Write-Host "Build completed. Binary is located at build\nmxchaind.exe" -ForegroundColor Green
} else {
    Write-Host "Build failed. Binary not created." -ForegroundColor Red
}
