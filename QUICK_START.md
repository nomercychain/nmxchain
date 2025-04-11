# NoMercyChain Quick Start Guide

This guide will help you quickly get started with the NoMercyChain project.

## Prerequisites

- Go 1.18 or higher
- Node.js 16 or higher
- npm 8 or higher
- Git

## Quick Setup

### 1. Clone the Repository

```bash
git clone https://github.com/nomercychain/nmxchain.git
cd nmxchain
```

### 2. Run the Setup Script

#### For Windows (PowerShell):

```powershell
# Run the PowerShell script
.\scripts\setup_dev_env.ps1
```

#### For Linux/macOS:

```bash
# Make the script executable
chmod +x scripts/setup_dev_env.sh
# Run the script
./scripts/setup_dev_env.sh
```

This script will:
- Check if Go and Node.js are installed
- Set up Go dependencies
- Set up frontend dependencies
- Create necessary scripts for building, testing, and running a local testnet

### 3. Build the Project

#### For Windows (PowerShell):

```powershell
.\scripts\build.ps1
```

#### For Linux/macOS:

```bash
./scripts/build.sh
```

### 4. Start a Local Testnet

#### For Windows (PowerShell):

```powershell
.\scripts\local_testnet\setup_local_testnet.ps1
```

#### For Linux/macOS:

```bash
./scripts/local_testnet/setup_local_testnet.sh
```

### 5. Start the Frontend

```bash
cd client
npm start
```

Then open your browser and navigate to `http://localhost:3000`

## Project Structure

- `x/dynacontracts/`: AI-powered smart contracts module
- `x/deai/`: Decentralized AI agents module
- `x/truthgpt/`: TruthGPT Oracle module
- `x/hyperchains/`: HyperChains module
- `client/`: Frontend dashboard

## Key Features

1. **DynaContracts**: AI-powered smart contracts that adapt based on external data and AI models
2. **DeAI Agents**: User-controlled AI agents that can perform actions on the blockchain
3. **TruthGPT Oracle**: AI-powered oracles for verifying information and detecting misinformation
4. **HyperChains**: AI-generated Layer 3 chains that can be created from natural language prompts

## Next Steps

1. Explore the codebase to understand the project structure
2. Review the `DEVELOPMENT_GUIDE.md` file for detailed instructions
3. Check the `NEXT_STEPS.md` file for a prioritized list of tasks
4. Start implementing the features outlined in the development guide

## Useful Commands

### Blockchain Commands

The commands are the same for both Windows and Unix-based systems:

```bash
# Check node status
nmxchaind status

# Query account balance
nmxchaind query bank balances <address>

# Send tokens
nmxchaind tx bank send <from-address> <to-address> <amount>unmx --chain-id nomercychain-local-1 --keyring-backend test

# Create an AI agent
nmxchaind tx deai create-agent <name> <description> <type> <config> --from <address> --chain-id nomercychain-local-1 --keyring-backend test

# Create a HyperChain from prompt
nmxchaind tx hyperchains create-chain-from-prompt <prompt> --from <address> --chain-id nomercychain-local-1 --keyring-backend test
```

### Frontend Commands

```bash
# Start development server
cd client
npm start

# Run tests
cd client
npm test

# Build for production
cd client
npm run build
```

## Documentation

For more detailed information, refer to the following documentation:

- `SETUP.md`: Setup instructions
- `DEVELOPMENT_GUIDE.md`: Comprehensive development guide
- `NEXT_STEPS.md`: Development roadmap and tasks
- `VS_CODE_SETUP.md`: Visual Studio Code setup guide for Windows

## Getting Help

If you encounter any issues or have questions, please:

1. Check the documentation
2. Review the codebase
3. Reach out to the development team

Happy coding!