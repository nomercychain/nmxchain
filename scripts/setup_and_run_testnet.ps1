# NoMercyChain Setup and Run Testnet Script
# This PowerShell script sets up and runs the NoMercyChain testnet

param (
    [switch]$Reset = $false,
    [switch]$SkipBuild = $false
)

$ErrorActionPreference = "Stop"

Write-Host "NoMercyChain Testnet Setup and Run" -ForegroundColor Cyan
Write-Host "=================================" -ForegroundColor Cyan

# Step 1: Update dependencies
Write-Host "Step 1: Updating dependencies..." -ForegroundColor Green
& "$PSScriptRoot/update_dependencies.ps1"

if (!$SkipBuild) {
    # Step 2: Build the chain
    Write-Host "Step 2: Building the chain..." -ForegroundColor Green
    & "$PSScriptRoot/build.ps1"
}

# Check if the binary exists
if (!(Test-Path "$PSScriptRoot/../build/nmxchaind.exe")) {
    Write-Host "Error: Chain binary not found. Please build the chain first." -ForegroundColor Red
    exit 1
}

# Add the binary to the PATH
$env:PATH = "$PSScriptRoot/../build;$env:PATH"

# Step 3: Reset the chain if requested
if ($Reset) {
    Write-Host "Step 3: Resetting the chain..." -ForegroundColor Green
    & "$PSScriptRoot/reset_testnet.ps1"
}

# Step 4: Set up the local testnet
Write-Host "Step 4: Setting up the local testnet..." -ForegroundColor Green
& "$PSScriptRoot/local_testnet/setup_local_testnet.ps1"

Write-Host "`nNoMercyChain Testnet is now running!" -ForegroundColor Cyan
Write-Host "RPC Endpoint: http://localhost:26657" -ForegroundColor Green
Write-Host "REST API: http://localhost:1317" -ForegroundColor Green
Write-Host "Swagger UI: http://localhost:1317/swagger/" -ForegroundColor Green
Write-Host "`nPress Ctrl+C to stop the chain." -ForegroundColor Yellow