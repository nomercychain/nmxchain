# NoMercyChain Testnet Guide

This guide provides instructions for setting up and interacting with the NoMercyChain testnet.

## Prerequisites

- Windows 10 or later
- PowerShell 5.1 or later
- Go 1.18 or later
- Git

## Setup

### Clone the Repository

```bash
git clone https://github.com/nomercychain/nmxchain.git
cd nmxchain
```

### Build the Chain

```bash
./scripts/build.ps1
```

### Initialize the Testnet

```bash
./scripts/local_testnet/setup_local_testnet.ps1
```

This script will:
1. Initialize the chain
2. Create validator and user accounts
3. Configure the genesis file
4. Start the chain

## Testnet Accounts

The setup script creates the following accounts:

- **validator**: The validator node account
- **faucet**: An account with a large balance for distributing tokens
- **user1**: A test user account
- **user2**: Another test user account

To get the mnemonic for any of these accounts:

```bash
./scripts/get_test_account_mnemonic.ps1 [account-name]
```

## Sending Tokens

To send tokens between accounts:

```bash
./scripts/send_tokens.ps1 -FromAccount [sender] -ToAddress [recipient-address] -Amount [amount]unmx
```

Example:
```bash
./scripts/send_tokens.ps1 -FromAccount faucet -ToAddress nmx1... -Amount 1000000unmx
```

## Resetting the Testnet

If you need to reset the testnet:

```bash
./scripts/reset_testnet.ps1
```

## Changing Log Level

To change the log level:

```bash
./scripts/set_log_level.ps1 -LogLevel [level]
```

Valid log levels: debug, info, warn, error, dpanic, panic, fatal

## Interacting with Modules

### DynaContract Module

Create a contract:
```bash
nmxchaind tx dynacontract create-contract [name] [version] [code-file] [init-msg] \
  --from [key] --chain-id nomercychain-testnet-1 --keyring-backend test
```

Execute a contract:
```bash
nmxchaind tx dynacontract execute-contract [id] [execute-msg] [coins] \
  --from [key] --chain-id nomercychain-testnet-1 --keyring-backend test
```

Query a contract:
```bash
nmxchaind query dynacontract show-contract [id]
```

### DeAI Module

Create an AI agent:
```bash
nmxchaind tx deai create-agent [name] [description] [model-file] [init-params] \
  --from [key] --chain-id nomercychain-testnet-1 --keyring-backend test
```

Query an agent:
```bash
nmxchaind query deai show-agent [id]
```

### Hyperchain Module

Create a hyperchain:
```bash
nmxchaind tx hyperchain create-chain [name] [description] [config-file] \
  --from [key] --chain-id nomercychain-testnet-1 --keyring-backend test
```

Query a hyperchain:
```bash
nmxchaind query hyperchain show-chain [id]
```

## REST API

The REST API is available at http://localhost:1317

Example endpoints:
- List contracts: http://localhost:1317/nomercychain/nmxchain/dynacontract/contracts
- Get contract: http://localhost:1317/nomercychain/nmxchain/dynacontract/contracts/{id}

## Swagger Documentation

Swagger documentation is available at http://localhost:1317/swagger/

## Troubleshooting

### Common Issues

1. **Chain fails to start**
   - Check logs for errors
   - Ensure ports 26656, 26657, and 1317 are available
   - Try resetting the testnet

2. **Transaction errors**
   - Verify account has sufficient balance
   - Check that the chain-id is correct
   - Ensure the keyring backend is set to "test"

3. **Connection refused**
   - Verify the chain is running
   - Check that the RPC endpoint is correct (default: http://localhost:26657)

### Getting Help

If you encounter issues not covered in this guide, please:
1. Check the logs for error messages
2. Review the documentation
3. Open an issue on GitHub with detailed information about the problem