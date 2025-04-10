# NoMercyChain (NMXChain)

NoMercyChain is a fully functioning, scalable, AI-powered Layer 1 blockchain built on Cosmos SDK. This blockchain integrates AI at the protocol level and supports dynamic smart contracts, predictive security, user-controlled AI agents, and auto-generated Layer 3 chains.

## Features

- **NeuroPoS Consensus**: AI-enhanced Proof of Stake consensus mechanism
- **DynaContracts**: AI-optimized smart contracts that can evolve with DAO governance
- **AI Validator Logic**: Integrated neural networks for anomaly detection and dynamic adjustment
- **$NMX Token**: Native token for gas fees, staking, governance, and rewards
- **DeAI Agent System**: User-created AI agents stored as NFTs
- **TruthGPT Oracle**: AI-powered oracle system for verifying external data
- **HyperChains**: AI-generated Layer 3 chains for specific use cases
- **Governance DAO**: Token-weighted proposals with AI agent delegate voting
- **ZK-AI Computation**: Off-chain AI computations verified via ZK-snarks

## Repository Structure

- `/cmd` - Main applications and entry points
- `/x` - Custom Cosmos SDK modules
- `/app` - Application initialization and wiring
- `/proto` - Protocol buffer definitions
- `/client` - Client-side applications (web, mobile)
- `/docs` - Documentation
- `/scripts` - Utility scripts
- `/test` - Test files

## Getting Started

### Prerequisites

- Windows 10 or later
- PowerShell 5.1 or later
- Go 1.18 or later
- Git

### Building from Source

1. Clone the repository:
   ```bash
   git clone https://github.com/nomercychain/nmxchain.git
   cd nmxchain
   ```

2. Build the chain:
   ```bash
   ./scripts/build.ps1
   ```

### Running a Local Testnet

#### Quick Start (New Method)

##### Linux/macOS
```bash
# Make the start script executable
chmod +x start.sh

# Start the blockchain and frontend
./start.sh
```

##### Windows
```bash
# Start the blockchain and frontend
start.bat
```

The application will be available at:
- Frontend: http://localhost:3000
- Blockchain API: http://localhost:1317
- Blockchain RPC: http://localhost:26657

#### Legacy Method

1. Initialize and start the testnet:
   ```bash
   ./scripts/local_testnet/setup_local_testnet.ps1
   ```

2. Monitor the testnet:
   ```bash
   ./scripts/monitor_testnet.ps1 -Continuous
   ```

3. Test the DynaContract module:
   ```bash
   ./scripts/test_dynacontract.ps1
   ```

4. Test module integration:
   ```bash
   ./scripts/test_integration.ps1
   ```

### Resetting the Testnet

If you need to reset the testnet:
```bash
./scripts/reset_testnet.ps1
```

## Documentation

- [Quick Start Guide](QUICK_START.md)
- [Development Guide](DEVELOPMENT_GUIDE.md)
- [Testnet Guide](TESTNET_GUIDE.md)
- [Setup Guide](SETUP.md)
- [VS Code Setup](VS_CODE_SETUP.md)

## Module Documentation

- [DynaContract Module](x/dynacontract/README.md)
- [DeAI Module](x/deai/README.md)
- [HyperChain Module](x/hyperchain/README.md)
- [NeuroPoS Module](x/neuropos/README.md)
- [TruthGPT Module](x/truthgpt/README.md)

## License

Copyright (c) 2023 NoMercyChain

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.#   n m x c h a i n 
 
 