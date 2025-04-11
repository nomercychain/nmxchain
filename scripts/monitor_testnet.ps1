# NoMercyChain Testnet Monitoring Script
# This script monitors the testnet and displays key information

param (
    [int]$RefreshInterval = 10,  # Refresh interval in seconds
    [switch]$Continuous = $false # Run continuously if true
)

function Get-NodeStatus {
    try {
        $status = Invoke-RestMethod -Uri "http://localhost:26657/status" -Method Get -ErrorAction Stop
        return $status.result
    } catch {
        Write-Host "Error getting node status: $_" -ForegroundColor Red
        return $null
    }
}

function Get-LatestBlock {
    try {
        $block = Invoke-RestMethod -Uri "http://localhost:26657/block" -Method Get -ErrorAction Stop
        return $block.result.block
    } catch {
        Write-Host "Error getting latest block: $_" -ForegroundColor Red
        return $null
    }
}

function Get-ValidatorSet {
    try {
        $validators = Invoke-RestMethod -Uri "http://localhost:26657/validators" -Method Get -ErrorAction Stop
        return $validators.result.validators
    } catch {
        Write-Host "Error getting validators: $_" -ForegroundColor Red
        return $null
    }
}

function Get-PeerCount {
    try {
        $netInfo = Invoke-RestMethod -Uri "http://localhost:26657/net_info" -Method Get -ErrorAction Stop
        return $netInfo.result.n_peers
    } catch {
        Write-Host "Error getting peer count: $_" -ForegroundColor Red
        return 0
    }
}

function Get-AccountBalance {
    param (
        [string]$Address
    )
    
    try {
        $balance = nmxchaind query bank balances $Address --output json 2>$null | ConvertFrom-Json
        return $balance.balances
    } catch {
        Write-Host "Error getting balance for $Address: $_" -ForegroundColor Red
        return $null
    }
}

function Get-ModuleStats {
    # Get DynaContract stats
    try {
        $contracts = nmxchaind query dynacontract dyna-contracts --output json 2>$null | ConvertFrom-Json
        $contractCount = if ($contracts.contracts) { $contracts.contracts.Count } else { 0 }
    } catch {
        $contractCount = 0
    }
    
    # Get DeAI stats
    try {
        $agents = nmxchaind query deai ai-agents --output json 2>$null | ConvertFrom-Json
        $agentCount = if ($agents.agents) { $agents.agents.Count } else { 0 }
    } catch {
        $agentCount = 0
    }
    
    # Get Hyperchain stats
    try {
        $chains = nmxchaind query hyperchain chains --output json 2>$null | ConvertFrom-Json
        $chainCount = if ($chains.chains) { $chains.chains.Count } else { 0 }
    } catch {
        $chainCount = 0
    }
    
    return @{
        ContractCount = $contractCount
        AgentCount = $agentCount
        ChainCount = $chainCount
    }
}

function Display-TestnetInfo {
    Clear-Host
    Write-Host "NoMercyChain Testnet Monitor" -ForegroundColor Cyan
    Write-Host "===========================" -ForegroundColor Cyan
    
    # Get node status
    $status = Get-NodeStatus
    if ($status) {
        Write-Host "`nNode Status:" -ForegroundColor Green
        Write-Host "  Node ID:      $($status.node_info.id)" -ForegroundColor White
        Write-Host "  Network:      $($status.node_info.network)" -ForegroundColor White
        Write-Host "  Version:      $($status.node_info.version)" -ForegroundColor White
        Write-Host "  Syncing:      $($status.sync_info.catching_up)" -ForegroundColor White
        Write-Host "  Latest Block: $($status.sync_info.latest_block_height) ($($status.sync_info.latest_block_time))" -ForegroundColor White
    }
    
    # Get peer count
    $peerCount = Get-PeerCount
    Write-Host "`nNetwork:" -ForegroundColor Green
    Write-Host "  Connected Peers: $peerCount" -ForegroundColor White
    
    # Get latest block
    $block = Get-LatestBlock
    if ($block) {
        Write-Host "`nLatest Block:" -ForegroundColor Green
        Write-Host "  Height:    $($block.header.height)" -ForegroundColor White
        Write-Host "  Time:      $($block.header.time)" -ForegroundColor White
        Write-Host "  Proposer:  $($block.header.proposer_address)" -ForegroundColor White
        Write-Host "  Tx Count:  $($block.data.txs.Count)" -ForegroundColor White
    }
    
    # Get validator set
    $validators = Get-ValidatorSet
    if ($validators) {
        Write-Host "`nValidators:" -ForegroundColor Green
        Write-Host "  Count: $($validators.Count)" -ForegroundColor White
        
        foreach ($validator in $validators) {
            $votingPower = [int]$validator.voting_power
            $votingPowerPercentage = ($votingPower / $validators.ForEach({[int]$_.voting_power}) | Measure-Object -Sum).Sum * 100
            Write-Host "  - $($validator.address) (Power: $votingPower, $([math]::Round($votingPowerPercentage, 2))%)" -ForegroundColor White
        }
    }
    
    # Get account balances
    Write-Host "`nAccount Balances:" -ForegroundColor Green
    
    $accounts = @("validator", "faucet", "user1", "user2")
    foreach ($account in $accounts) {
        $address = nmxchaind keys show $account -a --keyring-backend test 2>$null
        if ($address) {
            $balance = Get-AccountBalance -Address $address
            if ($balance) {
                $balanceStr = ($balance | ForEach-Object { "$($_.amount) $($_.denom)" }) -join ", "
                Write-Host "  $account ($address): $balanceStr" -ForegroundColor White
            } else {
                Write-Host "  $account ($address): Error getting balance" -ForegroundColor Red
            }
        }
    }
    
    # Get module stats
    $moduleStats = Get-ModuleStats
    Write-Host "`nModule Statistics:" -ForegroundColor Green
    Write-Host "  DynaContract: $($moduleStats.ContractCount) contracts" -ForegroundColor White
    Write-Host "  DeAI:         $($moduleStats.AgentCount) AI agents" -ForegroundColor White
    Write-Host "  Hyperchain:   $($moduleStats.ChainCount) chains" -ForegroundColor White
    
    # Display refresh time
    Write-Host "`nLast Updated: $(Get-Date)" -ForegroundColor Gray
    if ($Continuous) {
        Write-Host "Refreshing every $RefreshInterval seconds. Press Ctrl+C to exit." -ForegroundColor Gray
    }
}

# Main monitoring loop
if ($Continuous) {
    try {
        while ($true) {
            Display-TestnetInfo
            Start-Sleep -Seconds $RefreshInterval
        }
    } catch {
        Write-Host "`nMonitoring stopped: $_" -ForegroundColor Red
    } finally {
        Write-Host "`nExiting monitor..." -ForegroundColor Yellow
    }
} else {
    Display-TestnetInfo
}