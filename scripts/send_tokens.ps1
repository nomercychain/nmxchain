# NoMercyChain Send Tokens Script
# This script sends tokens from one account to another

param (
    [Parameter(Mandatory=$true)]
    [string]$FromAccount,
    
    [Parameter(Mandatory=$true)]
    [string]$ToAddress,
    
    [Parameter(Mandatory=$true)]
    [string]$Amount
)

# Configuration
$CHAIN_ID = "nomercychain-testnet-1"

Write-Host "NoMercyChain Send Tokens Script" -ForegroundColor Cyan
Write-Host "===============================" -ForegroundColor Cyan

# Validate parameters
if ([string]::IsNullOrEmpty($FromAccount)) {
    Write-Host "Error: From account name is required" -ForegroundColor Red
    exit 1
}

if ([string]::IsNullOrEmpty($ToAddress)) {
    Write-Host "Error: To address is required" -ForegroundColor Red
    exit 1
}

if ([string]::IsNullOrEmpty($Amount)) {
    Write-Host "Error: Amount is required" -ForegroundColor Red
    exit 1
}

# Set keyring backend
$keyringBackend = "test"

# Get the from address
try {
    $fromAddress = nmxchaind keys show $FromAccount -a --keyring-backend $keyringBackend
    if ([string]::IsNullOrEmpty($fromAddress)) {
        throw "Could not get address for account $FromAccount"
    }
} catch {
    Write-Host "Error: Could not find account $FromAccount. Make sure it exists in the keyring." -ForegroundColor Red
    exit 1
}

# Send tokens
Write-Host "Sending $Amount from $FromAccount ($fromAddress) to $ToAddress..." -ForegroundColor Yellow

try {
    $result = nmxchaind tx bank send $fromAddress $ToAddress $Amount `
        --chain-id $CHAIN_ID `
        --keyring-backend $keyringBackend `
        --broadcast-mode block `
        --yes
    
    if ($LASTEXITCODE -ne 0) {
        throw "Transaction failed"
    }
    
    Write-Host "Transaction successful!" -ForegroundColor Green
    
    # Extract and display the transaction hash
    $txHash = ($result | Select-String -Pattern "txhash: (.+)").Matches.Groups[1].Value
    if (-not [string]::IsNullOrEmpty($txHash)) {
        Write-Host "Transaction hash: $txHash" -ForegroundColor Cyan
    }
    
    # Check the recipient's balance
    Write-Host "Checking recipient's balance..." -ForegroundColor Yellow
    nmxchaind query bank balances $ToAddress --chain-id $CHAIN_ID
    
} catch {
    Write-Host "Error: Transaction failed. $($_.Exception.Message)" -ForegroundColor Red
    exit 1
}

Write-Host "Tokens sent successfully." -ForegroundColor Green