# NoMercyChain Project Status Checker
# This script checks the current state of the project and provides guidance

$ErrorActionPreference = "Stop"

Write-Host "NoMercyChain Project Status" -ForegroundColor Cyan
Write-Host "========================" -ForegroundColor Cyan

# Check if build directory exists
$buildExists = Test-Path "build/nmxchaind.exe"
if ($buildExists) {
    Write-Host "✅ Build exists: build/nmxchaind.exe" -ForegroundColor Green
} else {
    Write-Host "❌ Build not found. You need to build the project first." -ForegroundColor Red
    Write-Host "   Run: .\build-force.ps1" -ForegroundColor Yellow
}

# Check if chain is initialized
$chainInitialized = Test-Path "$HOME\.nmxchain\config\genesis.json"
if ($chainInitialized) {
    Write-Host "✅ Chain is initialized" -ForegroundColor Green
} else {
    Write-Host "❌ Chain is not initialized" -ForegroundColor Red
    if ($buildExists) {
        Write-Host "   Run: .\build\nmxchaind.exe init testnode --chain-id nomercychain-testnet-1" -ForegroundColor Yellow
    } else {
        Write-Host "   Build the project first, then initialize the chain" -ForegroundColor Yellow
    }
}

# Check if chain is running
$chainRunning = $false
try {
    $response = Invoke-RestMethod -Uri "http://localhost:26657/status" -Method Get -ErrorAction SilentlyContinue
    $chainRunning = $true
    Write-Host "✅ Chain is running" -ForegroundColor Green
    Write-Host "   Chain ID: $($response.result.node_info.network)" -ForegroundColor Green
    Write-Host "   Latest block: $($response.result.sync_info.latest_block_height)" -ForegroundColor Green
} catch {
    Write-Host "❌ Chain is not running" -ForegroundColor Red
    if ($chainInitialized -and $buildExists) {
        Write-Host "   Run: .\build\nmxchaind.exe start" -ForegroundColor Yellow
    } elseif ($buildExists) {
        Write-Host "   Initialize the chain first, then start it" -ForegroundColor Yellow
    } else {
        Write-Host "   Build the project first, then initialize and start the chain" -ForegroundColor Yellow
    }
}

# Check if client directory exists
$clientExists = Test-Path "client"
if ($clientExists) {
    Write-Host "✅ Client directory exists" -ForegroundColor Green
    
    # Check if client has node_modules
    $nodeModulesExist = Test-Path "client/node_modules"
    if ($nodeModulesExist) {
        Write-Host "✅ Client dependencies are installed" -ForegroundColor Green
    } else {
        Write-Host "❌ Client dependencies are not installed" -ForegroundColor Red
        Write-Host "   Run: cd client && npm install" -ForegroundColor Yellow
    }
    
    # Check if .env.testnet exists
    $envTestnetExists = Test-Path "client/.env.testnet"
    if ($envTestnetExists) {
        Write-Host "✅ Client testnet environment file exists" -ForegroundColor Green
    } else {
        Write-Host "❌ Client testnet environment file not found" -ForegroundColor Red
        Write-Host "   Create client/.env.testnet with appropriate configuration" -ForegroundColor Yellow
    }
} else {
    Write-Host "❌ Client directory not found" -ForegroundColor Red
    Write-Host "   Frontend development might not be set up yet" -ForegroundColor Yellow
}

# Provide guidance based on current state
Write-Host "`nRecommended next steps:" -ForegroundColor Cyan

if (!$buildExists) {
    Write-Host "1. Build the project: .\build-force.ps1" -ForegroundColor White
    Write-Host "2. Initialize the chain: .\build\nmxchaind.exe init testnode --chain-id nomercychain-testnet-1" -ForegroundColor White
    Write-Host "3. Start the chain: .\build\nmxchaind.exe start" -ForegroundColor White
} elseif (!$chainInitialized) {
    Write-Host "1. Initialize the chain: .\build\nmxchaind.exe init testnode --chain-id nomercychain-testnet-1" -ForegroundColor White
    Write-Host "2. Create validator account: .\build\nmxchaind.exe keys add validator --keyring-backend test" -ForegroundColor White
    Write-Host "3. Add genesis account: .\build\nmxchaind.exe add-genesis-account <address> 10000000000unmx" -ForegroundColor White
    Write-Host "4. Create genesis transaction: .\build\nmxchaind.exe gentx validator 1000000000unmx --chain-id nomercychain-testnet-1 --keyring-backend test" -ForegroundColor White
    Write-Host "5. Collect genesis transactions: .\build\nmxchaind.exe collect-gentxs" -ForegroundColor White
    Write-Host "6. Start the chain: .\build\nmxchaind.exe start" -ForegroundColor White
} elseif (!$chainRunning) {
    Write-Host "1. Start the chain: .\build\nmxchaind.exe start" -ForegroundColor White
    if ($clientExists) {
        Write-Host "2. Start the frontend: .\start-frontend.ps1" -ForegroundColor White
    }
} else {
    if ($clientExists) {
        Write-Host "1. Start the frontend: .\start-frontend.ps1" -ForegroundColor White
    }
    Write-Host "Chain is running. You can interact with it using the following commands:" -ForegroundColor White
    Write-Host "- Check account balance: .\build\nmxchaind.exe query bank balances <address>" -ForegroundColor White
    Write-Host "- Send tokens: .\build\nmxchaind.exe tx bank send validator <recipient-address> 1000000unmx --chain-id nomercychain-testnet-1 --keyring-backend test" -ForegroundColor White
    Write-Host "- List validators: .\build\nmxchaind.exe query staking validators" -ForegroundColor White
}

Write-Host "`nFor a complete setup in one step, you can try: .\init-and-start.ps1" -ForegroundColor White