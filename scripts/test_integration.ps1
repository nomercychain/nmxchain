# NoMercyChain Module Integration Test Script
# This script tests the integration between different modules

Write-Host "NoMercyChain Module Integration Test Script" -ForegroundColor Cyan
Write-Host "==========================================" -ForegroundColor Cyan

# Configuration
$CHAIN_ID = "nomercychain-testnet-1"
$KEYRING_BACKEND = "test"
$USER = "user1"
$TEST_DATA_DIR = "$PSScriptRoot/test_data"

# Create the test data directory if it doesn't exist
if (!(Test-Path $TEST_DATA_DIR)) {
    New-Item -ItemType Directory -Path $TEST_DATA_DIR | Out-Null
}

# Get user address
$USER_ADDRESS = nmxchaind keys show $USER -a --keyring-backend $KEYRING_BACKEND
if ([string]::IsNullOrEmpty($USER_ADDRESS)) {
    Write-Host "Error: Could not get address for user $USER" -ForegroundColor Red
    exit 1
}

Write-Host "Testing with user: $USER ($USER_ADDRESS)" -ForegroundColor Yellow

# Step 1: Create an AI agent in the DeAI module
Write-Host "`nStep 1: Creating an AI agent..." -ForegroundColor Cyan

# Create a sample AI model file
$AI_MODEL_FILE = "$TEST_DATA_DIR/sample_ai_model.json"
$SAMPLE_AI_MODEL = @'
{
  "name": "TestAIModel",
  "version": "1.0.0",
  "description": "A test AI model for integration testing",
  "parameters": {
    "input_size": 10,
    "hidden_size": 20,
    "output_size": 5
  },
  "weights": "base64encodedweights..."
}
'@
Set-Content -Path $AI_MODEL_FILE -Value $SAMPLE_AI_MODEL

$CREATE_AGENT_RESULT = nmxchaind tx deai create-ai-agent `
    --name "TestAgent" `
    --description "A test AI agent for integration testing" `
    --model-type "CLASSIFICATION" `
    --model-file $AI_MODEL_FILE `
    --deposit "5000000unmx" `
    --from $USER `
    --chain-id $CHAIN_ID `
    --keyring-backend $KEYRING_BACKEND `
    --broadcast-mode block `
    --yes

if ($LASTEXITCODE -ne 0) {
    Write-Host "Error: Failed to create AI agent" -ForegroundColor Red
    exit 1
}

# Extract agent ID from the result
$AGENT_ID = ($CREATE_AGENT_RESULT | Select-String -Pattern "agent_id: (.+)").Matches.Groups[1].Value
if ([string]::IsNullOrEmpty($AGENT_ID)) {
    Write-Host "Error: Could not extract agent ID from result" -ForegroundColor Red
    exit 1
}

Write-Host "AI agent created with ID: $AGENT_ID" -ForegroundColor Green

# Step 2: Create a dynamic contract that uses the AI agent
Write-Host "`nStep 2: Creating a dynamic contract with AI integration..." -ForegroundColor Cyan

# Create a sample contract file
$CONTRACT_CODE_FILE = "$TEST_DATA_DIR/ai_integrated_contract.js"
$AI_INTEGRATED_CONTRACT = @"
// AI-Integrated Contract
let state = { predictions: [] };

function init(msg) {
    state.owner = msg.owner;
    state.agentId = msg.agentId;
    return { success: true, message: "Contract initialized with AI agent: " + state.agentId };
}

function execute(msg) {
    if (msg.action === "predict") {
        // In a real implementation, this would call the AI agent
        // For this test, we'll simulate a prediction
        const prediction = {
            input: msg.input,
            result: "Simulated prediction for " + msg.input,
            timestamp: Date.now()
        };
        state.predictions.push(prediction);
        return { success: true, prediction: prediction };
    } else {
        return { success: false, error: "Unknown action" };
    }
}

function query(msg) {
    if (msg.action === "getPredictions") {
        return { predictions: state.predictions };
    } else if (msg.action === "getAgentInfo") {
        return { agentId: state.agentId, owner: state.owner };
    } else {
        return { error: "Unknown query" };
    }
}
"@
Set-Content -Path $CONTRACT_CODE_FILE -Value $AI_INTEGRATED_CONTRACT

$INIT_MSG = @"
{
  "init": true,
  "owner": "$USER",
  "agentId": "$AGENT_ID"
}
"@

$CREATE_CONTRACT_RESULT = nmxchaind tx dynacontract create-dyna-contract `
    --name "AIIntegratedContract" `
    --description "A contract that integrates with an AI agent" `
    --contract-type 2 `
    --code $CONTRACT_CODE_FILE `
    --abi "{}" `
    --agent-id $AGENT_ID `
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
$CONTRACT_ID = ($CREATE_CONTRACT_RESULT | Select-String -Pattern "contract_id: (.+)").Matches.Groups[1].Value
if ([string]::IsNullOrEmpty($CONTRACT_ID)) {
    Write-Host "Error: Could not extract contract ID from result" -ForegroundColor Red
    exit 1
}

Write-Host "AI-integrated contract created with ID: $CONTRACT_ID" -ForegroundColor Green

# Step 3: Execute the contract with a prediction request
Write-Host "`nStep 3: Executing the contract with a prediction request..." -ForegroundColor Cyan

$EXECUTE_MSG = '{"action": "predict", "input": "test data"}'

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

Write-Host "Contract executed successfully" -ForegroundColor Green

# Step 4: Query contracts by AI agent
Write-Host "`nStep 4: Querying contracts by AI agent..." -ForegroundColor Cyan

$QUERY_BY_AGENT_RESULT = nmxchaind query dynacontract dyna-contracts-by-agent $AGENT_ID --output json
if ($LASTEXITCODE -ne 0) {
    Write-Host "Error: Failed to query contracts by agent" -ForegroundColor Red
    exit 1
}

Write-Host "Contracts using agent $AGENT_ID:" -ForegroundColor Green
$QUERY_BY_AGENT_RESULT | ConvertFrom-Json | Format-List

# Step 5: Add learning data to the contract
Write-Host "`nStep 5: Adding learning data to the contract..." -ForegroundColor Cyan

$LEARNING_DATA = '{"input": "sample input", "output": "sample output", "feedback": "positive"}'

$ADD_LEARNING_DATA_RESULT = nmxchaind tx dynacontract add-learning-data `
    --contract-id $CONTRACT_ID `
    --data-type "TRAINING" `
    --data $LEARNING_DATA `
    --source "integration_test" `
    --from $USER `
    --chain-id $CHAIN_ID `
    --keyring-backend $KEYRING_BACKEND `
    --broadcast-mode block `
    --yes

if ($LASTEXITCODE -ne 0) {
    Write-Host "Error: Failed to add learning data" -ForegroundColor Red
    exit 1
}

# Extract data ID from the result
$DATA_ID = ($ADD_LEARNING_DATA_RESULT | Select-String -Pattern "data_id: (.+)").Matches.Groups[1].Value
if ([string]::IsNullOrEmpty($DATA_ID)) {
    Write-Host "Error: Could not extract data ID from result" -ForegroundColor Red
    exit 1
}

Write-Host "Learning data added with ID: $DATA_ID" -ForegroundColor Green

# Step 6: Query learning data for the contract
Write-Host "`nStep 6: Querying learning data for the contract..." -ForegroundColor Cyan

$QUERY_LEARNING_DATA_RESULT = nmxchaind query dynacontract dyna-contract-learning-data $CONTRACT_ID --output json
if ($LASTEXITCODE -ne 0) {
    Write-Host "Error: Failed to query learning data" -ForegroundColor Red
    exit 1
}

Write-Host "Learning data for contract $CONTRACT_ID:" -ForegroundColor Green
$QUERY_LEARNING_DATA_RESULT | ConvertFrom-Json | Format-List

Write-Host "`nModule integration test completed successfully!" -ForegroundColor Cyan