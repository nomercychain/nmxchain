# NoMercyChain Testnet Setup Guide

This guide provides step-by-step instructions for setting up and running the NoMercyChain testnet.

## Prerequisites

- Windows 10 or later
- PowerShell 5.1 or later
- Go 1.20 or later
- Git

## Setup Steps

### 1. Clone the Repository

```bash
git clone https://github.com/nomercychain/nmxchain.git
cd nmxchain
```

### 2. Update Dependencies

Run the dependency update script to ensure all required Go modules are installed:

```bash
./scripts/update_dependencies.ps1
```

### 3. Build the Chain

Build the NoMercyChain binary:

```bash
./scripts/build.ps1
```

This will create the `nmxchaind.exe` binary in the `build` directory.

### 4. Set Up and Run the Testnet

Use the all-in-one setup script to set up and run the testnet:

```bash
./scripts/setup_and_run_testnet.ps1
```

This script will:
- Update dependencies
- Build the chain (if not already built)
- Set up the local testnet
- Start the chain

### 5. Reset the Testnet (Optional)

If you need to reset the testnet to its initial state:

```bash
./scripts/reset_testnet.ps1
```

Then set up and run the testnet again:

```bash
./scripts/setup_and_run_testnet.ps1
```

## Testnet Access Points

Once the testnet is running, you can access it at:

- **RPC Endpoint**: http://localhost:26657
- **REST API**: http://localhost:1317
- **Swagger UI**: http://localhost:1317/swagger/

## Working with the Testnet

### Creating a New Account

```bash
nmxchaind keys add <name> --keyring-backend test
```

### Checking Account Balance

```bash
nmxchaind query bank balances <address>
```

### Sending Tokens

```bash
nmxchaind tx bank send <from_address> <to_address> <amount>unmx --chain-id nomercychain-testnet-1 --keyring-backend test
```

### Querying Transactions

```bash
nmxchaind query tx <txhash>
```

## Troubleshooting

### Chain Fails to Start

If the chain fails to start, try resetting the testnet:

```bash
./scripts/reset_testnet.ps1
./scripts/setup_and_run_testnet.ps1
```

### Build Errors

If you encounter build errors, ensure your Go environment is properly set up and try updating the dependencies:

```bash
./scripts/update_dependencies.ps1
```

### Port Conflicts

If there are port conflicts (e.g., ports 26656, 26657, or 1317 are already in use), stop any processes using those ports before starting the testnet.

## Next Steps

After setting up the testnet, you can:

1. Explore the REST API using the Swagger UI
2. Create and execute DynaContracts
3. Create and interact with AI Agents
4. Set up and join HyperChains
5. Use the TruthGPT Oracle for data verification