# NoMercyChain Build Script

Write-Host "Building NoMercyChain..." -ForegroundColor Green

# Step 1: Clean up any previous build artifacts
if (Test-Path "build") {
    Write-Host "Cleaning up previous build..." -ForegroundColor Yellow
    Remove-Item -Path "build" -Recurse -Force
}

# Step 2: Create build directory
New-Item -ItemType Directory -Force -Path "build" | Out-Null

# Step 3: Update Go dependencies
Write-Host "Updating Go dependencies..." -ForegroundColor Yellow
go mod tidy
if ($LASTEXITCODE -ne 0) {
    Write-Host "Failed to update dependencies. Trying to fix..." -ForegroundColor Red
    
    # Try to fix common dependency issues
    Write-Host "Downloading direct dependencies..." -ForegroundColor Yellow
    go get github.com/cosmos/cosmos-sdk@v0.47.5
    go get github.com/cometbft/cometbft@v0.37.2
    go get github.com/cometbft/cometbft-db@v0.7.0
    go get github.com/cosmos/ibc-go/v7@v7.3.1
    
    # Run tidy again
    Write-Host "Running go mod tidy again..." -ForegroundColor Yellow
    go mod tidy
    if ($LASTEXITCODE -ne 0) {
        Write-Host "Failed to update dependencies. Please check your go.mod file." -ForegroundColor Red
        exit 1
    }
}

# Step 4: Build the binary
Write-Host "Building binary..." -ForegroundColor Yellow
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
Write-Host "1. Initialize the chain: .\scripts\local_testnet\setup_local_testnet.ps1" -ForegroundColor White
Write-Host "2. Or start everything: .\scripts\start_testnet_full.ps1" -ForegroundColor White
Write-Host "3. Or use the simplified start script: .\start.bat" -ForegroundColor White