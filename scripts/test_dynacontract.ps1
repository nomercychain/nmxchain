# NoMercyChain DynaContract Test Script
# This script tests the dynacontract module functionality

Write-Host "NoMercyChain DynaContract Test Script" -ForegroundColor Cyan
Write-Host "=====================================" -ForegroundColor Cyan

# Configuration
$CHAIN_ID = "nomercychain-testnet-1"
$KEYRING_BACKEND = "test"
$USER = "user1"
$CONTRACT_NAME = "TestContract"
$CONTRACT_VERSION = "1.0.0"
$CONTRACT_CODE_FILE = "$PSScriptRoot/test_data/sample_contract.js"
$INIT_MSG = '{"init": true, "owner": "user1"}'
$EXECUTE_MSG = '{"action": "increment", "value": 1}'

# Create the test data directory if it doesn't exist
$TEST_DATA_DIR = "$PSScriptRoot/test_data"
if (!(Test-Path $TEST_DATA_DIR)) {
    New-Item -ItemType Directory -Path $TEST_DATA_DIR | Out-Null
}

# Create a sample contract file if it doesn't exist
if (!(Test-Path $CONTRACT_CODE_FILE)) {
    $SAMPLE_CONTRACT = @'
// Sample Counter Contract
let state = { count: 0 };

function init(msg) {
    state.owner = msg.owner;
    state.count = 0;
    return { success: true, message: "Contract initialized" };
}

function execute(msg) {
    if (msg.action === "increment") {
        state.count += msg.value;
        return { success: true, count: state.count };
    } else if (msg.action === "decrement") {
        state.count -= msg.value;
        return { success: true, count: state.count };
    } else if (msg.action === "reset") {
        state.count = 0;
        return { success: true, count: state.count };
    } else {
        return { success: false, error: "Unknown action" };
    }
}

function query(msg) {
    if (msg.action === "getCount") {
        return { count: state.count };
    } else if (msg.action === "getOwner") {
        return { owner: state.owner };
    } else {
        return { error: "Unknown query" };
    }
}
'@
    Set-Content -Path $CONTRACT_CODE_FILE -Value $SAMPLE_CONTRACT
    Write-Host "Created sample contract file at $CONTRACT_CODE_FILE" -ForegroundColor Green
}

# Get user address
$USER_ADDRESS = nmxchaind keys show $USER -a --keyring-backend $KEYRING_BACKEND
if ([string]::IsNullOrEmpty($USER_ADDRESS)) {
    Write-Host "Error: Could not get address for user $USER" -ForegroundColor Red
    exit 1
}

Write-Host "Testing with user: $USER ($USER_ADDRESS)" -ForegroundColor Yellow

# Step 1: Create a contract
Write-Host "`nStep 1: Creating a contract..." -ForegroundColor Cyan
$CREATE_RESULT = nmxchaind tx dynacontract create-dyna-contract `
    --name $CONTRACT_NAME `
    --description "A test contract" `
    --contract-type 1 `
    --code $CONTRACT_CODE_FILE `
    --abi "{}" `
    --gas-limit 1000000 `
    --deposit "1000000unmx" `
    --from $USER `
    --chain-id $CHAIN_ID `
    --keyring-backend $KEYRING_BACKEND `
    --broadcast-mode block `
    --yes

if ($LASTEXITCODE -ne 0) {
    Write-Host "Error: Failed to create contract" -ForegroundColor Red
    exit 1
}

# Extract contract ID from the result
$CONTRACT_ID = ($CREATE_RESULT | Select-String -Pattern "contract_id: (.+)").Matches.Groups[1].Value
if ([string]::IsNullOrEmpty($CONTRACT_ID)) {
    Write-Host "Error: Could not extract contract ID from result" -ForegroundColor Red
    exit 1
}

Write-Host "Contract created with ID: $CONTRACT_ID" -ForegroundColor Green

# Step 2: Query the contract
Write-Host "`nStep 2: Querying the contract..." -ForegroundColor Cyan
$QUERY_RESULT = nmxchaind query dynacontract dyna-contract $CONTRACT_ID --output json
if ($LASTEXITCODE -ne 0) {
    Write-Host "Error: Failed to query contract" -ForegroundColor Red
    exit 1
}

Write-Host "Contract details:" -ForegroundColor Green
$QUERY_RESULT | ConvertFrom-Json | Format-List

# Step 3: Execute the contract
Write-Host "`nStep 3: Executing the contract..." -ForegroundColor Cyan
$EXECUTE_RESULT = nmxchaind tx dynacontract execute-dyna-contract `
    $CONTRACT_ID `
    $EXECUTE_MSG `
    --fee "1000unmx" `
    --from $USER `
    --chain-id $CHAIN_ID `
    --keyring-backend $KEYRING_BACKEND `
    --broadcast-mode block `
    --yes

if ($LASTEXITCODE -ne 0) {
    Write-Host "Error: Failed to execute contract" -ForegroundColor Red
    exit 1
}

# Extract execution ID from the result
$EXECUTION_ID = ($EXECUTE_RESULT | Select-String -Pattern "execution_id: (.+)").Matches.Groups[1].Value
if ([string]::IsNullOrEmpty($EXECUTION_ID)) {
    Write-Host "Error: Could not extract execution ID from result" -ForegroundColor Red
    exit 1
}

Write-Host "Contract executed with execution ID: $EXECUTION_ID" -ForegroundColor Green

# Step 4: Query contract executions
Write-Host "`nStep 4: Querying contract executions..." -ForegroundColor Cyan
$EXECUTIONS_RESULT = nmxchaind query dynacontract dyna-contract-executions $CONTRACT_ID --output json
if ($LASTEXITCODE -ne 0) {
    Write-Host "Error: Failed to query contract executions" -ForegroundColor Red
    exit 1
}

Write-Host "Contract executions:" -ForegroundColor Green
$EXECUTIONS_RESULT | ConvertFrom-Json | Format-List

# Step 5: Update the contract
Write-Host "`nStep 5: Updating the contract..." -ForegroundColor Cyan
$UPDATE_RESULT = nmxchaind tx dynacontract update-dyna-contract `
    --contract-id $CONTRACT_ID `
    --name "$CONTRACT_NAME Updated" `
    --description "An updated test contract" `
    --from $USER `
    --chain-id $CHAIN_ID `
    --keyring-backend $KEYRING_BACKEND `
    --broadcast-mode block `
    --yes

if ($LASTEXITCODE -ne 0) {
    Write-Host "Error: Failed to update contract" -ForegroundColor Red
    exit 1
}

Write-Host "Contract updated successfully" -ForegroundColor Green

# Step 6: Query the updated contract
Write-Host "`nStep 6: Querying the updated contract..." -ForegroundColor Cyan
$UPDATED_QUERY_RESULT = nmxchaind query dynacontract dyna-contract $CONTRACT_ID --output json
if ($LASTEXITCODE -ne 0) {
    Write-Host "Error: Failed to query updated contract" -ForegroundColor Red
    exit 1
}

Write-Host "Updated contract details:" -ForegroundColor Green
$UPDATED_QUERY_RESULT | ConvertFrom-Json | Format-List

Write-Host "`nDynaContract module test completed successfully!" -ForegroundColor Cyan