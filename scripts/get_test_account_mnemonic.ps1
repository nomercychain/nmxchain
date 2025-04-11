# NoMercyChain Get Test Account Mnemonic Script
# This script retrieves the mnemonic for a test account

param (
    [Parameter(Mandatory=$true)]
    [string]$AccountName
)

Write-Host "NoMercyChain Get Test Account Mnemonic Script" -ForegroundColor Cyan
Write-Host "==========================================" -ForegroundColor Cyan

# Validate parameters
if ([string]::IsNullOrEmpty($AccountName)) {
    Write-Host "Error: Account name is required" -ForegroundColor Red
    exit 1
}

# Set keyring backend
$keyringBackend = "test"

# Check if the account exists
try {
    $accountInfo = nmxchaind keys show $AccountName --keyring-backend $keyringBackend
    if ([string]::IsNullOrEmpty($accountInfo)) {
        throw "Could not get info for account $AccountName"
    }
} catch {
    Write-Host "Error: Could not find account $AccountName. Make sure it exists in the keyring." -ForegroundColor Red
    exit 1
}

# Get the mnemonic
Write-Host "Retrieving mnemonic for account $AccountName..." -ForegroundColor Yellow
Write-Host "You will be prompted to enter the keyring password (default is empty, just press Enter)" -ForegroundColor Yellow

try {
    # Export the key with the mnemonic
    nmxchaind keys export $AccountName --unsafe --keyring-backend $keyringBackend
    
    Write-Host "IMPORTANT: Keep this mnemonic phrase safe and secure!" -ForegroundColor Red
    Write-Host "Anyone with access to this mnemonic can access your account and funds." -ForegroundColor Red
    
} catch {
    Write-Host "Error: Failed to retrieve mnemonic. $($_.Exception.Message)" -ForegroundColor Red
    exit 1
}