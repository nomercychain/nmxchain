# NoMercyChain Test Script for Windows
# This PowerShell script runs all tests for the NoMercyChain project

Write-Host "Running NoMercyChain tests..." -ForegroundColor Green

# Run Go tests
Write-Host "Running Go tests..." -ForegroundColor Green
go test -v ./...

# Run frontend tests
Write-Host "Running frontend tests..." -ForegroundColor Green
Set-Location -Path client
npm test
Set-Location -Path ..

Write-Host "All tests completed." -ForegroundColor Green
