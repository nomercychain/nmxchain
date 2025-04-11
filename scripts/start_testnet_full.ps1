# NoMercyChain Full Testnet Startup Script
# This script starts both the backend chain and frontend application

param (
    [switch]$ResetChain = $false,
    [switch]$BackendOnly = $false,
    [switch]$FrontendOnly = $false
)

$ErrorActionPreference = "Stop"

Write-Host "NoMercyChain Full Testnet Startup" -ForegroundColor Cyan
Write-Host "===============================" -ForegroundColor Cyan

# Check if the chain binary exists
if (!(Test-Path "$PSScriptRoot/../build/nmxchaind.exe")) {
    Write-Host "Chain binary not found. Building chain..." -ForegroundColor Yellow
    & "$PSScriptRoot/build.ps1"
}

# Start backend if not frontend only
if (!$FrontendOnly) {
    # Reset chain if requested
    if ($ResetChain) {
        Write-Host "Resetting chain data..." -ForegroundColor Yellow
        & "$PSScriptRoot/reset_testnet.ps1"
    }

    # Check if chain is already running
    $chainRunning = $false
    try {
        $response = Invoke-RestMethod -Uri "http://localhost:26657/status" -Method Get -ErrorAction SilentlyContinue
        if ($response.result.node_info.network -eq "nomercychain-testnet-1") {
            $chainRunning = $true
            Write-Host "Chain is already running." -ForegroundColor Green
        }
    } catch {
        # Chain is not running, which is fine
    }

    # Start the chain if it's not already running
    if (!$chainRunning) {
        Write-Host "Starting the chain..." -ForegroundColor Green
        
        # Check if the chain has been initialized
        if (!(Test-Path "$HOME/.nmxchain/config/genesis.json")) {
            Write-Host "Chain not initialized. Setting up local testnet..." -ForegroundColor Yellow
            Start-Process -FilePath "powershell.exe" -ArgumentList "-File `"$PSScriptRoot/local_testnet/setup_local_testnet.ps1`"" -NoNewWindow
            
            # Wait for the chain to start
            Write-Host "Waiting for chain to start..." -ForegroundColor Yellow
            $chainStarted = $false
            $attempts = 0
            while (!$chainStarted -and $attempts -lt 30) {
                try {
                    $response = Invoke-RestMethod -Uri "http://localhost:26657/status" -Method Get -ErrorAction SilentlyContinue
                    if ($response.result.node_info.network -eq "nomercychain-testnet-1") {
                        $chainStarted = $true
                        Write-Host "Chain started successfully." -ForegroundColor Green
                    }
                } catch {
                    # Chain not started yet
                }
                $attempts++
                Start-Sleep -Seconds 2
            }
            
            if (!$chainStarted) {
                Write-Host "Failed to start chain. Please check logs." -ForegroundColor Red
                exit 1
            }
        } else {
            # Chain is initialized but not running, start it
            Start-Process -FilePath "powershell.exe" -ArgumentList "-Command `"& nmxchaind start`"" -NoNewWindow
            
            # Wait for the chain to start
            Write-Host "Waiting for chain to start..." -ForegroundColor Yellow
            $chainStarted = $false
            $attempts = 0
            while (!$chainStarted -and $attempts -lt 15) {
                try {
                    $response = Invoke-RestMethod -Uri "http://localhost:26657/status" -Method Get -ErrorAction SilentlyContinue
                    if ($response.result.node_info.network -eq "nomercychain-testnet-1") {
                        $chainStarted = $true
                        Write-Host "Chain started successfully." -ForegroundColor Green
                    }
                } catch {
                    # Chain not started yet
                }
                $attempts++
                Start-Sleep -Seconds 2
            }
            
            if (!$chainStarted) {
                Write-Host "Failed to start chain. Please check logs." -ForegroundColor Red
                exit 1
            }
        }
    }
}

# Start frontend if not backend only
if (!$BackendOnly) {
    Write-Host "Starting frontend application..." -ForegroundColor Green
    
    # Change to client directory
    Push-Location "$PSScriptRoot/../client"
    
    # Start the frontend
    & "$PSScriptRoot/../client/start_testnet.ps1"
    
    # Return to original directory
    Pop-Location
}

if (!$BackendOnly -and !$FrontendOnly) {
    Write-Host "`nNoMercyChain Testnet is now running!" -ForegroundColor Cyan
    Write-Host "Backend: http://localhost:26657 (RPC) and http://localhost:1317 (REST API)" -ForegroundColor Green
    Write-Host "Frontend: http://localhost:3000" -ForegroundColor Green
    Write-Host "`nPress Ctrl+C to stop the frontend. The backend will continue running in the background." -ForegroundColor Yellow
    Write-Host "To stop the backend, run: Get-Process -Name nmxchaind | Stop-Process" -ForegroundColor Yellow
} elseif ($BackendOnly) {
    Write-Host "`nNoMercyChain Backend is now running!" -ForegroundColor Cyan
    Write-Host "Backend: http://localhost:26657 (RPC) and http://localhost:1317 (REST API)" -ForegroundColor Green
    Write-Host "`nTo stop the backend, run: Get-Process -Name nmxchaind | Stop-Process" -ForegroundColor Yellow
} elseif ($FrontendOnly) {
    Write-Host "`nNoMercyChain Frontend is now running!" -ForegroundColor Cyan
    Write-Host "Frontend: http://localhost:3000" -ForegroundColor Green
    Write-Host "`nPress Ctrl+C to stop the frontend." -ForegroundColor Yellow
}