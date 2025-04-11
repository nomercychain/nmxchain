# NoMercyChain Reset Testnet Script
# This script resets the testnet to its initial state

# Configuration
$HOME_DIR = "$HOME/.nmxchain"

Write-Host "NoMercyChain Testnet Reset Script" -ForegroundColor Cyan
Write-Host "=================================" -ForegroundColor Cyan

# Check if the data directory exists
if (!(Test-Path $HOME_DIR)) {
    Write-Host "Data directory not found at $HOME_DIR. Nothing to reset." -ForegroundColor Yellow
    exit 0
}

# Confirm reset
Write-Host "WARNING: This will delete all blockchain data at $HOME_DIR" -ForegroundColor Red
$confirmation = Read-Host "Are you sure you want to proceed? (y/n)"

if ($confirmation -ne "y") {
    Write-Host "Reset cancelled." -ForegroundColor Yellow
    exit 0
}

# Stop the running node if any
$nmxProcess = Get-Process -Name "nmxchaind" -ErrorAction SilentlyContinue
if ($nmxProcess) {
    Write-Host "Stopping running node..." -ForegroundColor Yellow
    Stop-Process -Name "nmxchaind" -Force
    Start-Sleep -Seconds 2
}

# Reset the chain data
Write-Host "Resetting chain data..." -ForegroundColor Yellow
nmxchaind unsafe-reset-all

# Alternatively, delete the entire data directory for a clean start
Write-Host "Deleting data directory..." -ForegroundColor Yellow
Remove-Item -Path $HOME_DIR -Recurse -Force

Write-Host "Testnet reset complete." -ForegroundColor Green
Write-Host "Run scripts/local_testnet/setup_local_testnet.ps1 to start a fresh testnet." -ForegroundColor Yellow