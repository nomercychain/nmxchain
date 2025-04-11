# NoMercyChain Build Script for Windows
# This PowerShell script builds the NoMercyChain binary

Write-Host "Building NoMercyChain..." -ForegroundColor Green

# Create build directory if it doesn't exist
New-Item -ItemType Directory -Force -Path build | Out-Null

# Build the binary
go build -o build\nmxchaind.exe .\cmd\nmxchaind

Write-Host "Build completed. Binary is located at build\nmxchaind.exe" -ForegroundColor Green
