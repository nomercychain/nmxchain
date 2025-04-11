# NoMercyChain Frontend Testnet Startup Script

Write-Host "Starting NoMercyChain Frontend with Testnet Configuration" -ForegroundColor Cyan
Write-Host "=======================================================" -ForegroundColor Cyan

# Copy testnet environment file
Write-Host "Setting up testnet environment..." -ForegroundColor Green
Copy-Item .env.testnet .env.local -Force

# Install dependencies if needed
if (!(Test-Path node_modules)) {
    Write-Host "Installing dependencies..." -ForegroundColor Green
    npm install
}

# Start the frontend
Write-Host "Starting frontend application..." -ForegroundColor Green
Write-Host "The application will be available at http://localhost:3000" -ForegroundColor Yellow
Write-Host "=======================================================" -ForegroundColor Cyan

npm start