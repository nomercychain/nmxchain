# NoMercyChain Update Dependencies Script
# This PowerShell script updates the Go module dependencies

Write-Host "Updating NoMercyChain dependencies..." -ForegroundColor Green

# Navigate to the root directory
Set-Location $PSScriptRoot\..

# Initialize the Go modules
Write-Host "Initializing Go modules..." -ForegroundColor Cyan
go mod tidy

# Get required dependencies
Write-Host "Getting required dependencies..." -ForegroundColor Cyan
go get github.com/cosmos/cosmos-sdk@v0.47.5
go get github.com/cosmos/ibc-go/v7@v7.3.1
go get github.com/spf13/cast@v1.5.1
go get github.com/spf13/cobra@v1.7.0
go get github.com/stretchr/testify@v1.8.4
go get github.com/tendermint/tendermint@v0.37.0
go get github.com/tendermint/tm-db@v0.6.7

# Tidy up the modules
Write-Host "Tidying up modules..." -ForegroundColor Cyan
go mod tidy

Write-Host "Dependencies updated successfully!" -ForegroundColor Green