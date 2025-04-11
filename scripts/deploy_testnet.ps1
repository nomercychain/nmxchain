# NoMercyChain Testnet Deployment Script
# This script prepares and deploys the testnet

param (
    [string]$ChainID = "nomercychain-testnet-1",
    [string]$ValidatorMoniker = "primary-validator",
    [string]$ValidatorStake = "100000000unmx",
    [string]$DeploymentDir = "$HOME/.nmxchain-testnet",
    [switch]$BuildOnly = $false
)

$ErrorActionPreference = "Stop"

Write-Host "NoMercyChain Testnet Deployment" -ForegroundColor Cyan
Write-Host "=============================" -ForegroundColor Cyan

# Step 1: Build the chain
Write-Host "Step 1: Building the chain..." -ForegroundColor Green
& "$PSScriptRoot/build.ps1"

if ($BuildOnly) {
    Write-Host "Build completed. Exiting as -BuildOnly was specified." -ForegroundColor Green
    exit 0
}

# Step 2: Create deployment directory
Write-Host "Step 2: Creating deployment directory..." -ForegroundColor Green
if (Test-Path $DeploymentDir) {
    Write-Host "Deployment directory already exists. Do you want to remove it? (y/n)" -ForegroundColor Yellow
    $response = Read-Host
    if ($response -eq "y") {
        Remove-Item -Recurse -Force $DeploymentDir
    } else {
        Write-Host "Deployment aborted." -ForegroundColor Red
        exit 1
    }
}
New-Item -ItemType Directory -Path $DeploymentDir | Out-Null

# Step 3: Initialize the chain
Write-Host "Step 3: Initializing the chain..." -ForegroundColor Green
$env:HOME_DIR = $DeploymentDir
nmxchaind init $ValidatorMoniker --chain-id $ChainID --home $DeploymentDir

# Step 4: Copy custom genesis file
Write-Host "Step 4: Configuring genesis..." -ForegroundColor Green
Copy-Item -Force "$PSScriptRoot/local_testnet/genesis.json" "$DeploymentDir/config/genesis.json"

# Step 5: Create validator key
Write-Host "Step 5: Creating validator key..." -ForegroundColor Green
nmxchaind keys add $ValidatorMoniker --keyring-backend test --home $DeploymentDir

# Step 6: Create additional accounts
Write-Host "Step 6: Creating additional accounts..." -ForegroundColor Green
nmxchaind keys add faucet --keyring-backend test --home $DeploymentDir
nmxchaind keys add user1 --keyring-backend test --home $DeploymentDir
nmxchaind keys add user2 --keyring-backend test --home $DeploymentDir

# Step 7: Get key addresses
$VALIDATOR_ADDRESS=$(nmxchaind keys show $ValidatorMoniker -a --keyring-backend test --home $DeploymentDir)
$FAUCET_ADDRESS=$(nmxchaind keys show faucet -a --keyring-backend test --home $DeploymentDir)
$USER1_ADDRESS=$(nmxchaind keys show user1 -a --keyring-backend test --home $DeploymentDir)
$USER2_ADDRESS=$(nmxchaind keys show user2 -a --keyring-backend test --home $DeploymentDir)

# Step 8: Add genesis accounts
Write-Host "Step 8: Adding genesis accounts..." -ForegroundColor Green
nmxchaind add-genesis-account $VALIDATOR_ADDRESS $ValidatorStake --home $DeploymentDir
nmxchaind add-genesis-account $FAUCET_ADDRESS "1000000000unmx" --home $DeploymentDir
nmxchaind add-genesis-account $USER1_ADDRESS "10000000unmx" --home $DeploymentDir
nmxchaind add-genesis-account $USER2_ADDRESS "10000000unmx" --home $DeploymentDir

# Step 9: Create validator transaction
Write-Host "Step 9: Creating validator transaction..." -ForegroundColor Green
nmxchaind gentx $ValidatorMoniker $ValidatorStake `
  --chain-id $ChainID `
  --moniker="$ValidatorMoniker" `
  --commission-rate="0.10" `
  --commission-max-rate="0.20" `
  --commission-max-change-rate="0.01" `
  --min-self-delegation="1" `
  --keyring-backend test `
  --home $DeploymentDir

# Step 10: Collect genesis transactions
Write-Host "Step 10: Collecting genesis transactions..." -ForegroundColor Green
nmxchaind collect-gentxs --home $DeploymentDir

# Step 11: Validate genesis
Write-Host "Step 11: Validating genesis..." -ForegroundColor Green
nmxchaind validate-genesis --home $DeploymentDir

# Step 12: Configure client
Write-Host "Step 12: Configuring client..." -ForegroundColor Green
nmxchaind config chain-id $ChainID --home $DeploymentDir
nmxchaind config keyring-backend test --home $DeploymentDir
nmxchaind config broadcast-mode block --home $DeploymentDir

# Step 13: Update config for testnet
Write-Host "Step 13: Updating config for testnet..." -ForegroundColor Green
$CONFIG_FILE="$DeploymentDir/config/config.toml"
$APP_FILE="$DeploymentDir/config/app.toml"

# Enable API server
(Get-Content $APP_FILE) -replace "enable = false", "enable = true" | Set-Content $APP_FILE

# Enable Swagger
(Get-Content $APP_FILE) -replace "swagger = false", "swagger = true" | Set-Content $APP_FILE

# Enable CORS
(Get-Content $CONFIG_FILE) -replace "cors_allowed_origins = \[\]", "cors_allowed_origins = [\"*\"]" | Set-Content $CONFIG_FILE

# Step 14: Build the frontend
Write-Host "Step 14: Building the frontend..." -ForegroundColor Green
Push-Location "$PSScriptRoot/../client"
npm install
npm run build
Pop-Location

# Step 15: Create deployment package
Write-Host "Step 15: Creating deployment package..." -ForegroundColor Green
$DEPLOY_PACKAGE="$PSScriptRoot/../nomercychain-testnet-deploy.zip"
$FRONTEND_BUILD="$PSScriptRoot/../client/build"

# Create a temporary directory for deployment files
$TEMP_DEPLOY_DIR="$PSScriptRoot/../temp-deploy"
if (Test-Path $TEMP_DEPLOY_DIR) {
    Remove-Item -Recurse -Force $TEMP_DEPLOY_DIR
}
New-Item -ItemType Directory -Path $TEMP_DEPLOY_DIR | Out-Null

# Copy necessary files
Copy-Item -Path "$PSScriptRoot/../build/nmxchaind.exe" -Destination "$TEMP_DEPLOY_DIR/"
Copy-Item -Path "$DeploymentDir" -Destination "$TEMP_DEPLOY_DIR/" -Recurse
Copy-Item -Path $FRONTEND_BUILD -Destination "$TEMP_DEPLOY_DIR/frontend" -Recurse

# Create deployment scripts
@"
# Start NoMercyChain Testnet Node
Write-Host "Starting NoMercyChain Testnet Node..." -ForegroundColor Green
Start-Process -FilePath "nmxchaind.exe" -ArgumentList "start --home .nmxchain-testnet" -NoNewWindow
"@ | Set-Content -Path "$TEMP_DEPLOY_DIR/start_node.ps1"

@"
# Start NoMercyChain Frontend
Write-Host "Starting NoMercyChain Frontend..." -ForegroundColor Green
npx serve -s frontend -l 3000
"@ | Set-Content -Path "$TEMP_DEPLOY_DIR/start_frontend.ps1"

# Create the deployment package
if (Test-Path $DEPLOY_PACKAGE) {
    Remove-Item -Force $DEPLOY_PACKAGE
}
Compress-Archive -Path "$TEMP_DEPLOY_DIR/*" -DestinationPath $DEPLOY_PACKAGE

# Clean up
Remove-Item -Recurse -Force $TEMP_DEPLOY_DIR

Write-Host "`nTestnet deployment package created: $DEPLOY_PACKAGE" -ForegroundColor Green
Write-Host "Validator Address: $VALIDATOR_ADDRESS" -ForegroundColor Green
Write-Host "Faucet Address: $FAUCET_ADDRESS" -ForegroundColor Green
Write-Host "User1 Address: $USER1_ADDRESS" -ForegroundColor Green
Write-Host "User2 Address: $USER2_ADDRESS" -ForegroundColor Green

Write-Host "`nTo start the testnet locally:" -ForegroundColor Cyan
Write-Host "1. Extract the deployment package" -ForegroundColor White
Write-Host "2. Run start_node.ps1 to start the blockchain node" -ForegroundColor White
Write-Host "3. Run start_frontend.ps1 to start the frontend application" -ForegroundColor White
Write-Host "`nThe testnet will be accessible at:" -ForegroundColor Cyan
Write-Host "- RPC: http://localhost:26657" -ForegroundColor White
Write-Host "- REST API: http://localhost:1317" -ForegroundColor White
Write-Host "- Frontend: http://localhost:3000" -ForegroundColor White