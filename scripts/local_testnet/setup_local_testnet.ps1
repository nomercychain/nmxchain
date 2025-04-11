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
