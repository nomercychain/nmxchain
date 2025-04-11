# NoMercyChain Development Environment Setup Script for Windows
# This PowerShell script sets up the development environment for the NoMercyChain project on Windows

Write-Host "Setting up NoMercyChain development environment..." -ForegroundColor Green

# Check if Go is installed
$goInstalled = $null
try {
    $goInstalled = Get-Command go -ErrorAction SilentlyContinue
} catch {
    $goInstalled = $null
}

if ($null -eq $goInstalled) {
    Write-Host "Go is not installed. Please install Go 1.18 or higher from https://golang.org/dl/" -ForegroundColor Yellow
    Write-Host "After installing Go, restart this script." -ForegroundColor Yellow
    exit 1
} else {
    $goVersion = (go version)
    Write-Host "Go is installed: $goVersion" -ForegroundColor Green
}

# Check if Node.js is installed
$nodeInstalled = $null
try {
    $nodeInstalled = Get-Command node -ErrorAction SilentlyContinue
} catch {
    $nodeInstalled = $null
}

if ($null -eq $nodeInstalled) {
    Write-Host "Node.js is not installed. Please install Node.js 16 or higher from https://nodejs.org/" -ForegroundColor Yellow
    Write-Host "After installing Node.js, restart this script." -ForegroundColor Yellow
    exit 1
} else {
    $nodeVersion = (node -v)
    Write-Host "Node.js is installed: $nodeVersion" -ForegroundColor Green
}

# Check if npm is installed
$npmInstalled = $null
try {
    $npmInstalled = Get-Command npm -ErrorAction SilentlyContinue
} catch {
    $npmInstalled = $null
}

if ($null -eq $npmInstalled) {
    Write-Host "npm is not installed. It should be installed with Node.js." -ForegroundColor Yellow
    Write-Host "Please reinstall Node.js from https://nodejs.org/" -ForegroundColor Yellow
    exit 1
} else {
    $npmVersion = (npm -v)
    Write-Host "npm is installed: $npmVersion" -ForegroundColor Green
}

# Set up Go dependencies
Write-Host "Setting up Go dependencies..." -ForegroundColor Green
go mod tidy

# Set up frontend dependencies
Write-Host "Setting up frontend dependencies..." -ForegroundColor Green
Set-Location -Path client
npm install
Set-Location -Path ..

# Create necessary directories
Write-Host "Creating necessary directories..." -ForegroundColor Green
New-Item -ItemType Directory -Force -Path scripts\local_testnet | Out-Null

# Set up local testnet script
$localTestnetScript = @'
# NoMercyChain Local Testnet Setup Script for Windows
# This PowerShell script sets up a local testnet for development and testing

Write-Host "Setting up NoMercyChain local testnet..." -ForegroundColor Green

# Initialize the chain
nmxchaind init local-testnet --chain-id nomercychain-local-1

# Create test accounts
nmxchaind keys add validator --keyring-backend test
nmxchaind keys add user1 --keyring-backend test
nmxchaind keys add user2 --keyring-backend test

# Add genesis accounts
$validatorAddress = (nmxchaind keys show validator -a --keyring-backend test)
$user1Address = (nmxchaind keys show user1 -a --keyring-backend test)
$user2Address = (nmxchaind keys show user2 -a --keyring-backend test)

nmxchaind add-genesis-account $validatorAddress 10000000000unmx
nmxchaind add-genesis-account $user1Address 1000000000unmx
nmxchaind add-genesis-account $user2Address 1000000000unmx

# Create validator transaction
nmxchaind gentx validator 5000000000unmx `
  --chain-id nomercychain-local-1 `
  --moniker="local-validator" `
  --commission-rate="0.10" `
  --commission-max-rate="0.20" `
  --commission-max-change-rate="0.01" `
  --min-self-delegation="1" `
  --keyring-backend test

# Collect gentxs
nmxchaind collect-gentxs

# Validate genesis
nmxchaind validate-genesis

# Start the chain
Write-Host "Starting the local testnet..." -ForegroundColor Green
nmxchaind start
'@

Set-Content -Path scripts\local_testnet\setup_local_testnet.ps1 -Value $localTestnetScript

# Create build script
$buildScript = @'
# NoMercyChain Build Script for Windows
# This PowerShell script builds the NoMercyChain binary

Write-Host "Building NoMercyChain..." -ForegroundColor Green

# Create build directory if it doesn't exist
New-Item -ItemType Directory -Force -Path build | Out-Null

# Build the binary
go build -o build\nmxchaind.exe .\cmd\nmxchaind

Write-Host "Build completed. Binary is located at build\nmxchaind.exe" -ForegroundColor Green
'@

Set-Content -Path scripts\build.ps1 -Value $buildScript

# Create test script
$testScript = @'
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
'@

Set-Content -Path scripts\run_tests.ps1 -Value $testScript

Write-Host "Development environment setup complete!" -ForegroundColor Green
Write-Host ""
Write-Host "Next steps:" -ForegroundColor Cyan
Write-Host "1. Review the DEVELOPMENT_GUIDE.md file for detailed instructions" -ForegroundColor Cyan
Write-Host "2. Run 'scripts\build.ps1' to build the NoMercyChain binary" -ForegroundColor Cyan
Write-Host "3. Run 'scripts\local_testnet\setup_local_testnet.ps1' to set up a local testnet" -ForegroundColor Cyan
Write-Host "4. Run 'scripts\run_tests.ps1' to run all tests" -ForegroundColor Cyan
Write-Host ""
Write-Host "Happy coding!" -ForegroundColor Cyan