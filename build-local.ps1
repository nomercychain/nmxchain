# NoMercyChain Local Build Script

Write-Host "Building NoMercyChain (Local Development)..." -ForegroundColor Green

# Step 1: Clean up any previous build artifacts
if (Test-Path "build") {
    Write-Host "Cleaning up previous build..." -ForegroundColor Yellow
    Remove-Item -Path "build" -Recurse -Force
}

# Step 2: Create build directory
New-Item -ItemType Directory -Force -Path "build" | Out-Null

# Step 3: Update go.mod to use local modules
Write-Host "Configuring for local development..." -ForegroundColor Yellow

# Create a temporary go.mod file
$goModContent = @"
module github.com/nomercychain/nmxchain

go 1.20

require (
	github.com/cosmos/cosmos-sdk v0.47.5
	github.com/cometbft/cometbft v0.37.2
	github.com/cometbft/cometbft-db v0.7.0
	github.com/spf13/cast v1.5.1
	github.com/spf13/cobra v1.7.0
	github.com/stretchr/testify v1.8.4
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/cosmos/cosmos-sdk v0.47.5
	github.com/cometbft/cometbft => github.com/cometbft/cometbft v0.37.2
	github.com/cometbft/cometbft-db => github.com/cometbft/cometbft-db v0.7.0
	github.com/tendermint/tendermint => github.com/cometbft/cometbft v0.37.2
	github.com/tendermint/tm-db => github.com/cometbft/cometbft-db v0.7.0
)
"@

Set-Content -Path "go.mod" -Value $goModContent

# Step 4: Run go mod tidy and then build
Write-Host "Running go mod tidy..." -ForegroundColor Yellow
go mod tidy
if ($LASTEXITCODE -ne 0) {
    Write-Host "Warning: go mod tidy returned non-zero exit code, but continuing with build..." -ForegroundColor Yellow
}

Write-Host "Building binary directly..." -ForegroundColor Yellow
go build -o build/nmxchaind.exe ./cmd/nmxchaind
if ($LASTEXITCODE -ne 0) {
    Write-Host "Build failed. Please check the error messages above." -ForegroundColor Red
    exit 1
}

Write-Host "Build completed successfully!" -ForegroundColor Green
Write-Host "Binary location: $PWD\build\nmxchaind.exe" -ForegroundColor Green

# Step 5: Add build directory to PATH for this session
$env:PATH = "$PWD\build;$env:PATH"
Write-Host "Added build directory to PATH for this session." -ForegroundColor Green
Write-Host "You can now run 'nmxchaind' commands directly." -ForegroundColor Green

# Step 6: Provide next steps
Write-Host "`nNext steps:" -ForegroundColor Cyan
Write-Host "1. Initialize the chain: .\build\nmxchaind.exe init testnode --chain-id nomercychain-testnet-1" -ForegroundColor White
Write-Host "2. Create validator account: .\build\nmxchaind.exe keys add validator --keyring-backend test" -ForegroundColor White
Write-Host "3. Add genesis account: .\build\nmxchaind.exe add-genesis-account <address> 10000000000unmx" -ForegroundColor White
Write-Host "4. Create genesis transaction: .\build\nmxchaind.exe gentx validator 1000000000unmx --chain-id nomercychain-testnet-1 --keyring-backend test" -ForegroundColor White
Write-Host "5. Collect genesis transactions: .\build\nmxchaind.exe collect-gentxs" -ForegroundColor White
Write-Host "6. Start the chain: .\build\nmxchaind.exe start --rpc.laddr tcp://0.0.0.0:26657 --api.enable true --api.address tcp://0.0.0.0:1317" -ForegroundColor White