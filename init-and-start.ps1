# NoMercyChain Initialize and Start Script

$ErrorActionPreference = "Stop"

Write-Host "NoMercyChain Initialize and Start" -ForegroundColor Cyan
Write-Host "===============================" -ForegroundColor Cyan

# Step 1: Build the chain
Write-Host "`nStep 1: Building the chain..." -ForegroundColor Green
.\build-local.ps1
if ($LASTEXITCODE -ne 0) {
    Write-Host "Build failed. Exiting." -ForegroundColor Red
    exit 1
}

# Step 2: Initialize the chain if not already initialized
if (!(Test-Path "$HOME\.nmxchain\config\genesis.json")) {
    Write-Host "`nStep 2: Initializing the chain..." -ForegroundColor Green
    
    # Initialize the chain
    Write-Host "Initializing chain..." -ForegroundColor Yellow
    .\build\nmxchaind.exe init testnode --chain-id nomercychain-testnet-1
    
    # Create validator account
    Write-Host "Creating validator account..." -ForegroundColor Yellow
    .\build\nmxchaind.exe keys add validator --keyring-backend test
    
    # Add genesis account
    Write-Host "Adding genesis account..." -ForegroundColor Yellow
    $validatorAddress = (.\build\nmxchaind.exe keys show validator -a --keyring-backend test)
    .\build\nmxchaind.exe add-genesis-account $validatorAddress 10000000000unmx
    
    # Create genesis transaction
    Write-Host "Creating genesis transaction..." -ForegroundColor Yellow
    .\build\nmxchaind.exe gentx validator 1000000000unmx --chain-id nomercychain-testnet-1 --keyring-backend test
    
    # Collect genesis transactions
    Write-Host "Collecting genesis transactions..." -ForegroundColor Yellow
    .\build\nmxchaind.exe collect-gentxs
    
    # Validate genesis
    Write-Host "Validating genesis..." -ForegroundColor Yellow
    .\build\nmxchaind.exe validate-genesis
} else {
    Write-Host "`nStep 2: Chain already initialized. Skipping initialization." -ForegroundColor Green
}

# Step 3: Start the chain
Write-Host "`nStep 3: Starting the blockchain..." -ForegroundColor Green
Write-Host "The blockchain will start in this window. Press Ctrl+C to stop." -ForegroundColor Yellow
Write-Host "Open a new terminal to interact with the chain or start the frontend." -ForegroundColor Yellow
Write-Host "`nBlockchain RPC: http://localhost:26657" -ForegroundColor Green
Write-Host "Blockchain API: http://localhost:1317" -ForegroundColor Green

# Start the chain in the current window
.\build\nmxchaind.exe start --rpc.laddr tcp://0.0.0.0:26657 --api.enable true --api.address tcp://0.0.0.0:1317