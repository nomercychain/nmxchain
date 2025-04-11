# NoMercyChain Set Log Level Script
# This script sets the log level for the NoMercyChain node

param (
    [Parameter(Mandatory=$true)]
    [ValidateSet("debug", "info", "warn", "error", "dpanic", "panic", "fatal")]
    [string]$LogLevel
)

Write-Host "NoMercyChain Set Log Level Script" -ForegroundColor Cyan
Write-Host "===============================" -ForegroundColor Cyan

# Configuration
$CONFIG_FILE = "$HOME/.nmxchain/config/config.toml"

# Check if the config file exists
if (!(Test-Path $CONFIG_FILE)) {
    Write-Host "Error: Config file not found at $CONFIG_FILE. Please run setup_local_testnet.ps1 first." -ForegroundColor Red
    exit 1
}

# Set the log level
Write-Host "Setting log level to $LogLevel..." -ForegroundColor Yellow

try {
    # Read the config file
    $configContent = Get-Content $CONFIG_FILE -Raw
    
    # Update the log level
    $configContent = $configContent -replace 'log_level = ".*"', "log_level = `"$LogLevel`""
    
    # Write the updated config back to the file
    Set-Content -Path $CONFIG_FILE -Value $configContent
    
    Write-Host "Log level set to $LogLevel successfully." -ForegroundColor Green
    Write-Host "You need to restart the node for the changes to take effect." -ForegroundColor Yellow
    
} catch {
    Write-Host "Error: Failed to set log level. $($_.Exception.Message)" -ForegroundColor Red
    exit 1
}