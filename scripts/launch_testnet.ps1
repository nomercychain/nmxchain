# NoMercyChain Testnet Launch Script
# This script sets up and launches a testnet for NoMercyChain

Write-Host "NoMercyChain Testnet Launch Script" -ForegroundColor Cyan
Write-Host "====================================" -ForegroundColor Cyan

# Configuration
$CHAIN_ID = "nomercychain-testnet-1"
$VALIDATOR_MONIKER = "primary-validator"
$VALIDATOR_STAKE = "5000000000unmx" # 5,000 NMX
$FAUCET_AMOUNT = "1000000000unmx"   # 1,000 NMX per account

# Create testnet directory
$TESTNET_DIR = "$HOME/.nomercychain/testnet"
if (!(Test-Path $TESTNET_DIR)) {
    Write-Host "Creating testnet directory at $TESTNET_DIR" -ForegroundColor Green
    New-Item -ItemType Directory -Path $TESTNET_DIR -Force | Out-Null
}

# Step 1: Build the binary
Write-Host "Building NoMercyChain binary..." -ForegroundColor Green
Set-Location $PSScriptRoot/..
go build -o build/nmxchaind ./cmd/nmxchaind
if ($LASTEXITCODE -ne 0) {
    Write-Host "Error building binary" -ForegroundColor Red
    exit 1
}
$BINARY = "$PSScriptRoot/../build/nmxchaind"

# Step 2: Initialize the chain
Write-Host "Initializing the chain..." -ForegroundColor Green
& $BINARY init $VALIDATOR_MONIKER --chain-id $CHAIN_ID --home $TESTNET_DIR
if ($LASTEXITCODE -ne 0) {
    Write-Host "Error initializing chain" -ForegroundColor Red
    exit 1
}

# Step 3: Create validator key
Write-Host "Creating validator key..." -ForegroundColor Green
& $BINARY keys add validator --keyring-backend test --home $TESTNET_DIR
if ($LASTEXITCODE -ne 0) {
    Write-Host "Error creating validator key" -ForegroundColor Red
    exit 1
}

# Get validator address
$VALIDATOR_ADDRESS = & $BINARY keys show validator -a --keyring-backend test --home $TESTNET_DIR
Write-Host "Validator address: $VALIDATOR_ADDRESS" -ForegroundColor Yellow

# Step 4: Create test accounts
Write-Host "Creating test accounts..." -ForegroundColor Green
$TEST_ACCOUNTS = @("user1", "user2", "user3", "faucet")
foreach ($ACCOUNT in $TEST_ACCOUNTS) {
    & $BINARY keys add $ACCOUNT --keyring-backend test --home $TESTNET_DIR
    if ($LASTEXITCODE -ne 0) {
        Write-Host "Error creating $ACCOUNT key" -ForegroundColor Red
        exit 1
    }
}

# Get account addresses
$ACCOUNT_ADDRESSES = @{}
foreach ($ACCOUNT in $TEST_ACCOUNTS) {
    $ACCOUNT_ADDRESSES[$ACCOUNT] = & $BINARY keys show $ACCOUNT -a --keyring-backend test --home $TESTNET_DIR
    Write-Host "$ACCOUNT address: $($ACCOUNT_ADDRESSES[$ACCOUNT])" -ForegroundColor Yellow
}

# Step 5: Add genesis accounts
Write-Host "Adding genesis accounts..." -ForegroundColor Green
& $BINARY add-genesis-account $VALIDATOR_ADDRESS "10000000000unmx" --home $TESTNET_DIR
if ($LASTEXITCODE -ne 0) {
    Write-Host "Error adding validator genesis account" -ForegroundColor Red
    exit 1
}

foreach ($ACCOUNT in $TEST_ACCOUNTS) {
    & $BINARY add-genesis-account $ACCOUNT_ADDRESSES[$ACCOUNT] $FAUCET_AMOUNT --home $TESTNET_DIR
    if ($LASTEXITCODE -ne 0) {
        Write-Host "Error adding $ACCOUNT genesis account" -ForegroundColor Red
        exit 1
    }
}

# Step 6: Create validator gentx
Write-Host "Creating validator gentx..." -ForegroundColor Green
& $BINARY gentx validator $VALIDATOR_STAKE `
  --chain-id $CHAIN_ID `
  --moniker $VALIDATOR_MONIKER `
  --commission-rate "0.10" `
  --commission-max-rate "0.20" `
  --commission-max-change-rate "0.01" `
  --min-self-delegation "1" `
  --keyring-backend test `
  --home $TESTNET_DIR
if ($LASTEXITCODE -ne 0) {
    Write-Host "Error creating validator gentx" -ForegroundColor Red
    exit 1
}

# Step 7: Collect gentxs
Write-Host "Collecting gentxs..." -ForegroundColor Green
& $BINARY collect-gentxs --home $TESTNET_DIR
if ($LASTEXITCODE -ne 0) {
    Write-Host "Error collecting gentxs" -ForegroundColor Red
    exit 1
}

# Step 8: Validate genesis
Write-Host "Validating genesis..." -ForegroundColor Green
& $BINARY validate-genesis --home $TESTNET_DIR
if ($LASTEXITCODE -ne 0) {
    Write-Host "Error validating genesis" -ForegroundColor Red
    exit 1
}

# Step 9: Update config files
Write-Host "Updating config files..." -ForegroundColor Green

# Update config.toml
$CONFIG_FILE = "$TESTNET_DIR/config/config.toml"
$CONFIG_CONTENT = Get-Content $CONFIG_FILE -Raw
$CONFIG_CONTENT = $CONFIG_CONTENT -replace "timeout_commit = ""5s""", "timeout_commit = ""1s"""
$CONFIG_CONTENT = $CONFIG_CONTENT -replace "cors_allowed_origins = \[\]", "cors_allowed_origins = [""*""]"
Set-Content -Path $CONFIG_FILE -Value $CONFIG_CONTENT

# Update app.toml
$APP_FILE = "$TESTNET_DIR/config/app.toml"
$APP_CONTENT = Get-Content $APP_FILE -Raw
$APP_CONTENT = $APP_CONTENT -replace "enable = false", "enable = true"
$APP_CONTENT = $APP_CONTENT -replace "swagger = false", "swagger = true"
Set-Content -Path $APP_FILE -Value $APP_CONTENT

# Step 10: Start the chain
Write-Host "Starting the NoMercyChain testnet..." -ForegroundColor Green
Write-Host "Chain ID: $CHAIN_ID" -ForegroundColor Yellow
Write-Host "Validator Address: $VALIDATOR_ADDRESS" -ForegroundColor Yellow
Write-Host "Faucet Address: $($ACCOUNT_ADDRESSES["faucet"])" -ForegroundColor Yellow
Write-Host "RPC Endpoint: http://localhost:26657" -ForegroundColor Yellow
Write-Host "REST API: http://localhost:1317" -ForegroundColor Yellow
Write-Host "====================================" -ForegroundColor Cyan

# Start the chain
& $BINARY start --home $TESTNET_DIR

# Note: The script will stop here as the chain is running in the foreground
# To run in the background, you would need to use a different approach like Start-Process