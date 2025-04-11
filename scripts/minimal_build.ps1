# NoMercyChain Minimal Build Script
# This PowerShell script builds a minimal version of the NoMercyChain binary

Write-Host "Building NoMercyChain (Minimal)..." -ForegroundColor Green

# Create build directory if it doesn't exist
New-Item -ItemType Directory -Force -Path build | Out-Null

# Create a temporary minimal go.mod file
$tempDir = New-Item -ItemType Directory -Force -Path "temp_build"
Set-Location $tempDir

# Initialize a new module
go mod init nmxchain
go get github.com/spf13/cobra@v1.7.0

# Copy the minimal main.go
Copy-Item -Path "..\cmd\nmxchaind\minimal_main.go" -Destination "main.go"

# Build the minimal binary
go build -o "..\build\nmxchaind.exe" main.go

# Clean up
Set-Location ..
Remove-Item -Recurse -Force $tempDir

if (Test-Path "build\nmxchaind.exe") {
    Write-Host "Build completed. Binary is located at build\nmxchaind.exe" -ForegroundColor Green
} else {
    Write-Host "Build failed. Binary not created." -ForegroundColor Red
}
