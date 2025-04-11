# NoMercyChain Testnet Guide

This guide provides instructions for setting up and running the NoMercyChain testnet, as well as connecting the frontend application to it.

## Prerequisites

- Go 1.18 or later
- Node.js 16 or later
- npm 8 or later
- PowerShell (for Windows) or Bash (for Linux/macOS)
- Keplr wallet browser extension

## Directory Structure

- `/cmd/nmxchaind`: Main blockchain application
- `/x`: Custom blockchain modules
- `/client`: Frontend application
- `/scripts`: Utility scripts

## 1. Launch the Testnet

### 1.1 Build and Start the Blockchain

Run the testnet launch script:

```powershell
# From the repository root
cd scripts
./launch_testnet.ps1
```

This script will:
1. Build the NoMercyChain binary
2. Initialize the blockchain
3. Create validator and test accounts
4. Configure the genesis file
5. Start the blockchain node

The blockchain node will run in the foreground. You'll need to keep this terminal window open.

### 1.2 Important Endpoints

- **RPC Endpoint**: http://localhost:26657
- **REST API**: http://localhost:1317
- **WebSocket**: ws://localhost:26657/websocket

## 2. Launch the Frontend

Open a new terminal window and run:

```powershell
# From the repository root
cd client
./start_testnet.ps1
```

This script will:
1. Set up the testnet environment configuration
2. Install dependencies if needed
3. Start the frontend application

The frontend will be available at http://localhost:3000

## 3. Connect Keplr Wallet

### 3.1 Add the Testnet to Keplr

When you first visit the frontend, you'll be prompted to add the NoMercyChain testnet to your Keplr wallet. Click "Approve" to add the network.

If you're not prompted automatically, you can add the network manually:

1. Open the frontend application
2. Click "Connect Wallet"
3. The application will request Keplr to add the NoMercyChain testnet

### 3.2 Import Test Accounts

The testnet launch script creates several test accounts. You can import these into Keplr:

1. Open Keplr extension
2. Click "Import account"
3. Enter the mnemonic phrase for one of the test accounts

To get the mnemonic phrase for a test account:

```powershell
# From the repository root
cd scripts
./get_test_account_mnemonic.ps1 user1
```

## 4. Get Test Tokens

The testnet is initialized with test accounts that have tokens. You can use the faucet account to send tokens to other accounts:

```powershell
# From the repository root
cd scripts
./send_tokens.ps1 faucet <recipient_address> 100unmx
```

## 5. Using the Application

### 5.1 Wallet

- View your account balance
- Send tokens to other addresses
- View transaction history

### 5.2 Staking

- View active validators
- Stake tokens to validators
- Claim staking rewards
- Unstake tokens

### 5.3 Governance

- View active proposals
- Create new proposals
- Vote on proposals
- Track proposal status

### 5.4 Smart Contracts

- Deploy DynaContracts
- Interact with contracts
- View contract details

### 5.5 HyperChains

- Create new HyperChains
- Configure chain parameters
- Join existing HyperChains

### 5.6 Oracle

- Submit oracle queries
- View query results
- Integrate oracle data into contracts

## 6. Development and Testing

### 6.1 Reset the Testnet

To reset the testnet to its initial state:

```powershell
# From the repository root
cd scripts
./reset_testnet.ps1
```

### 6.2 Blockchain Explorer

You can access a simple blockchain explorer at:

http://localhost:1317/swagger/

### 6.3 Logs and Debugging

Blockchain logs are printed to the console where the node is running. For more detailed logs, you can adjust the log level in the configuration:

```powershell
# From the repository root
cd scripts
./set_log_level.ps1 debug
```

## 7. Troubleshooting

### 7.1 Common Issues

#### Keplr Connection Issues

If you have trouble connecting Keplr:
1. Make sure the Keplr extension is installed and unlocked
2. Try refreshing the page
3. Check that the chain ID in `.env.testnet` matches the one in the launch script

#### Transaction Failures

If transactions fail:
1. Check that you have sufficient tokens for gas fees
2. Verify that the node is running and accessible
3. Check the blockchain logs for error messages

#### API Connection Issues

If the frontend can't connect to the blockchain:
1. Verify that the blockchain node is running
2. Check that the API endpoints in `.env.testnet` are correct
3. Ensure CORS is properly configured in the node's config.toml

### 7.2 Getting Help

If you encounter issues not covered in this guide:
1. Check the error logs
2. Review the documentation in the `/docs` directory
3. Open an issue on the GitHub repository

## 8. Next Steps

After successfully setting up and testing the NoMercyChain testnet, you might want to:

1. Explore the custom modules in the `/x` directory
2. Develop your own DynaContracts
3. Create a HyperChain for your specific use case
4. Contribute to the project by submitting pull requests

Happy building on NoMercyChain!