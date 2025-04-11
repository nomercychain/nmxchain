# NoMercyChain Force Build Script
# This script tries multiple approaches to build the chain

$ErrorActionPreference = "Stop"

Write-Host "NoMercyChain Force Build" -ForegroundColor Cyan
Write-Host "======================" -ForegroundColor Cyan

# Step 1: Clean up any previous build artifacts
if (Test-Path "build") {
    Write-Host "Cleaning up previous build..." -ForegroundColor Yellow
    Remove-Item -Path "build" -Recurse -Force
}

# Step 2: Create build directory
New-Item -ItemType Directory -Force -Path "build" | Out-Null

# Step 3: Try to build with existing go.mod
Write-Host "`nAttempt 1: Building with existing go.mod..." -ForegroundColor Green
Write-Host "Running go mod tidy..." -ForegroundColor Yellow
go mod tidy
Write-Host "Building binary..." -ForegroundColor Yellow
$buildResult = go build -o build/nmxchaind.exe ./cmd/nmxchaind 2>&1
if ($LASTEXITCODE -eq 0) {
    Write-Host "Build successful!" -ForegroundColor Green
} else {
    Write-Host "Build failed with existing go.mod. Trying alternative approach..." -ForegroundColor Yellow
    Write-Host "Error was: $buildResult" -ForegroundColor Red
    
    # Step 4: Try with a simplified go.mod
    Write-Host "`nAttempt 2: Building with simplified go.mod..." -ForegroundColor Green
    
    # Backup original go.mod
    Copy-Item go.mod go.mod.bak
    Copy-Item go.sum go.sum.bak -ErrorAction SilentlyContinue
    
    # Create a simplified go.mod
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
	github.com/tendermint/tendermint => github.com/cometbft/cometbft v0.37.2
	github.com/tendermint/tm-db => github.com/cometbft/cometbft-db v0.7.0
)
"@
    
    Set-Content -Path "go.mod" -Value $goModContent
    Remove-Item go.sum -ErrorAction SilentlyContinue
    
    Write-Host "Running go mod tidy..." -ForegroundColor Yellow
    go mod tidy
    Write-Host "Building binary..." -ForegroundColor Yellow
    $buildResult = go build -o build/nmxchaind.exe ./cmd/nmxchaind 2>&1
    if ($LASTEXITCODE -eq 0) {
        Write-Host "Build successful!" -ForegroundColor Green
    } else {
        Write-Host "Build failed with simplified go.mod. Trying final approach..." -ForegroundColor Yellow
        Write-Host "Error was: $buildResult" -ForegroundColor Red
        
        # Step 5: Try with vendor directory
        Write-Host "`nAttempt 3: Building with vendor directory..." -ForegroundColor Green
        
        # Restore original go.mod
        Copy-Item go.mod.bak go.mod
        Copy-Item go.sum.bak go.sum -ErrorAction SilentlyContinue
        
        # Create vendor directory
        Write-Host "Creating vendor directory..." -ForegroundColor Yellow
        go mod vendor
        
        Write-Host "Building with vendor directory..." -ForegroundColor Yellow
        $buildResult = go build -mod=vendor -o build/nmxchaind.exe ./cmd/nmxchaind 2>&1
        if ($LASTEXITCODE -eq 0) {
            Write-Host "Build successful!" -ForegroundColor Green
        } else {
            Write-Host "All build attempts failed." -ForegroundColor Red
            Write-Host "Error was: $buildResult" -ForegroundColor Red
            
            # Restore original go.mod
            Copy-Item go.mod.bak go.mod
            Copy-Item go.sum.bak go.sum -ErrorAction SilentlyContinue
            
            # Clean up
            Remove-Item go.mod.bak
            Remove-Item go.sum.bak -ErrorAction SilentlyContinue
            
            exit 1
        }
    }
    
    # Clean up
    Remove-Item go.mod.bak
    Remove-Item go.sum.bak -ErrorAction SilentlyContinue
}

Write-Host "`nBuild completed successfully!" -ForegroundColor Green
Write-Host "Binary location: $PWD\build\nmxchaind.exe" -ForegroundColor Green

# Add build directory to PATH for this session
$env:PATH = "$PWD\build;$env:PATH"
Write-Host "Added build directory to PATH for this session." -ForegroundColor Green
Write-Host "You can now run 'nmxchaind' commands directly." -ForegroundColor Green

# Provide next steps
Write-Host "`nNext steps:" -ForegroundColor Cyan
Write-Host "1. Initialize the chain: .\init-and-start.ps1" -ForegroundColor White
Write-Host "2. Or manually initialize: .\build\nmxchaind.exe init testnode --chain-id nomercychain-testnet-1" -ForegroundColor White