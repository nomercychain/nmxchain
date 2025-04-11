# NoMercyChain Testnet Implementation Plan

This document outlines the steps required to implement a fully functional testnet for NoMercyChain and connect the frontend to real blockchain data.

## 1. Complete Blockchain Module Implementation

### 1.1 Integrate Custom Modules into App

Update `app/app.go` to include our custom modules:

```go
// Import custom modules
import (
    // Existing imports...
    "github.com/nomercychain/x/deai"
    deaikeeper "github.com/nomercychain/x/deai/keeper"
    deaitypes "github.com/nomercychain/x/deai/types"
    
    "github.com/nomercychain/x/dynacontracts"
    dynacontractskeeper "github.com/nomercychain/x/dynacontracts/keeper"
    dynacontractstypes "github.com/nomercychain/x/dynacontracts/types"
    
    "github.com/nomercychain/x/hyperchains"
    hyperchainskeeper "github.com/nomercychain/x/hyperchains/keeper"
    hyperchaintypes "github.com/nomercychain/x/hyperchains/types"
    
    "github.com/nomercychain/x/neuropos"
    neuroposkeeper "github.com/nomercychain/x/neuropos/keeper"
    neuropostypes "github.com/nomercychain/x/neuropos/types"
    
    "github.com/nomercychain/x/truthgpt"
    truthgptkeeper "github.com/nomercychain/x/truthgpt/keeper"
    truthgpttypes "github.com/nomercychain/x/truthgpt/types"
)

// Add custom modules to ModuleBasics
ModuleBasics = module.NewBasicManager(
    // Existing modules...
    deai.AppModuleBasic{},
    dynacontracts.AppModuleBasic{},
    hyperchains.AppModuleBasic{},
    neuropos.AppModuleBasic{},
    truthgpt.AppModuleBasic{},
)

// Add custom keepers to NMXApp struct
type NMXApp struct {
    // Existing keepers...
    DeAIKeeper         deaikeeper.Keeper
    DynaContractsKeeper dynacontractskeeper.Keeper
    HyperChainsKeeper   hyperchainskeeper.Keeper
    NeuroPOSKeeper      neuroposkeeper.Keeper
    TruthGPTKeeper      truthgptkeeper.Keeper
}
```

### 1.2 Complete Module Implementations

For each module, ensure the following components are implemented:

- **Types**: Message types, events, and state objects
- **Keeper**: State management and business logic
- **Handler**: Message processing
- **Querier**: Query handling
- **Genesis**: Genesis state management
- **Module**: Module registration and initialization

## 2. Blockchain API Implementation

### 2.1 REST API Endpoints

Implement REST API endpoints for each module:

```go
// In each module's client/rest/rest.go file
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
    registerQueryRoutes(clientCtx, r)
    registerTxRoutes(clientCtx, r)
}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
    // Register query endpoints
}

func registerTxRoutes(clientCtx client.Context, r *mux.Router) {
    // Register transaction endpoints
}
```

### 2.2 gRPC Endpoints

Implement gRPC services for each module:

```protobuf
// In each module's proto/query.proto file
service Query {
  rpc ModuleState(QueryModuleStateRequest) returns (QueryModuleStateResponse);
  // Additional RPC methods
}
```

## 3. Testnet Deployment

### 3.1 Update Genesis Configuration

Modify the genesis configuration to include our custom modules:

```json
{
  "app_state": {
    "deai": {
      // DeAI module genesis state
    },
    "dynacontracts": {
      // DynaContracts module genesis state
    },
    "hyperchains": {
      // HyperChains module genesis state
    },
    "neuropos": {
      // NeuroPOS module genesis state
    },
    "truthgpt": {
      // TruthGPT module genesis state
    }
  }
}
```

### 3.2 Testnet Deployment Script

Create a script to deploy the testnet:

```bash
#!/bin/bash

# Initialize chain
nmxchaind init testnet --chain-id nomercychain-testnet-1

# Create validator keys
nmxchaind keys add validator1 --keyring-backend test
nmxchaind keys add validator2 --keyring-backend test

# Add genesis accounts
nmxchaind add-genesis-account $(nmxchaind keys show validator1 -a --keyring-backend test) 10000000000unmx
nmxchaind add-genesis-account $(nmxchaind keys show validator2 -a --keyring-backend test) 10000000000unmx

# Create validator transactions
nmxchaind gentx validator1 5000000000unmx --chain-id nomercychain-testnet-1 --keyring-backend test
nmxchaind gentx validator2 5000000000unmx --chain-id nomercychain-testnet-1 --keyring-backend test

# Collect gentxs
nmxchaind collect-gentxs

# Validate genesis
nmxchaind validate-genesis

# Start the chain
nmxchaind start
```

### 3.3 Persistent Peers Configuration

Configure persistent peers for testnet nodes:

```toml
# In config.toml
persistent_peers = "validator1_node_id@validator1_ip:26656,validator2_node_id@validator2_ip:26656"
```

## 4. Frontend Integration

### 4.1 API Client Implementation

Create a JavaScript client for interacting with the blockchain:

```javascript
// client/src/api/blockchain.js
import axios from 'axios';

const API_BASE_URL = process.env.REACT_APP_API_URL || 'http://localhost:1317';

export const blockchainApi = {
  // Account endpoints
  getAccount: async (address) => {
    const response = await axios.get(`${API_BASE_URL}/cosmos/auth/v1beta1/accounts/${address}`);
    return response.data;
  },
  
  getBalance: async (address) => {
    const response = await axios.get(`${API_BASE_URL}/cosmos/bank/v1beta1/balances/${address}`);
    return response.data;
  },
  
  // Staking endpoints
  getValidators: async () => {
    const response = await axios.get(`${API_BASE_URL}/cosmos/staking/v1beta1/validators`);
    return response.data;
  },
  
  getDelegations: async (address) => {
    const response = await axios.get(`${API_BASE_URL}/cosmos/staking/v1beta1/delegations/${address}`);
    return response.data;
  },
  
  // Governance endpoints
  getProposals: async () => {
    const response = await axios.get(`${API_BASE_URL}/cosmos/gov/v1beta1/proposals`);
    return response.data;
  },
  
  // Custom module endpoints
  getDynaContracts: async () => {
    const response = await axios.get(`${API_BASE_URL}/nomercychain/dynacontracts/v1/contracts`);
    return response.data;
  },
  
  getHyperChains: async () => {
    const response = await axios.get(`${API_BASE_URL}/nomercychain/hyperchains/v1/chains`);
    return response.data;
  },
  
  getOracleQueries: async (address) => {
    const response = await axios.get(`${API_BASE_URL}/nomercychain/truthgpt/v1/queries/${address}`);
    return response.data;
  },
};
```

### 4.2 Transaction Signing

Implement transaction signing using CosmJS:

```javascript
// client/src/api/transactions.js
import { SigningStargateClient, GasPrice } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";

// Custom message types
import { MsgCreateDynaContract } from "./proto/dynacontracts/tx";
import { MsgCreateHyperChain } from "./proto/hyperchains/tx";
import { MsgSubmitOracleQuery } from "./proto/truthgpt/tx";

// Create registry with custom message types
const registry = new Registry();
registry.register("/nomercychain.dynacontracts.v1.MsgCreateDynaContract", MsgCreateDynaContract);
registry.register("/nomercychain.hyperchains.v1.MsgCreateHyperChain", MsgCreateHyperChain);
registry.register("/nomercychain.truthgpt.v1.MsgSubmitOracleQuery", MsgSubmitOracleQuery);

export const connectWallet = async (keplr) => {
  if (!keplr) {
    throw new Error("Keplr wallet not found");
  }
  
  await keplr.enable("nomercychain-testnet-1");
  const offlineSigner = keplr.getOfflineSigner("nomercychain-testnet-1");
  const accounts = await offlineSigner.getAccounts();
  
  const client = await SigningStargateClient.connectWithSigner(
    "http://localhost:26657",
    offlineSigner,
    { registry, gasPrice: GasPrice.fromString("0.025unmx") }
  );
  
  return { client, accounts };
};

export const sendTransaction = async (client, sender, msg, memo = "") => {
  const fee = {
    amount: [{ denom: "unmx", amount: "5000" }],
    gas: "200000",
  };
  
  const result = await client.signAndBroadcast(
    sender.address,
    [msg],
    fee,
    memo
  );
  
  return result;
};
```

### 4.3 Update React Components

Update React components to use real data:

```javascript
// Example for Staking.js
import React, { useState, useEffect, useContext } from 'react';
import { blockchainApi } from '../api/blockchain';
import { WalletContext } from '../context/WalletContext';

const Staking = () => {
  const { account } = useContext(WalletContext);
  const [validators, setValidators] = useState([]);
  const [delegations, setDelegations] = useState([]);
  const [loading, setLoading] = useState(true);
  
  useEffect(() => {
    const fetchData = async () => {
      try {
        const validatorsResponse = await blockchainApi.getValidators();
        setValidators(validatorsResponse.validators || []);
        
        if (account) {
          const delegationsResponse = await blockchainApi.getDelegations(account);
          setDelegations(delegationsResponse.delegation_responses || []);
        }
      } catch (error) {
        console.error("Error fetching staking data:", error);
      } finally {
        setLoading(false);
      }
    };
    
    fetchData();
  }, [account]);
  
  // Component rendering logic
};
```

## 5. Wallet Integration

### 5.1 Keplr Wallet Integration

Update the WalletContext to integrate with Keplr:

```javascript
// client/src/context/WalletContext.js
import React, { createContext, useState, useEffect } from 'react';
import { connectWallet } from '../api/transactions';

export const WalletContext = createContext();

export const WalletProvider = ({ children }) => {
  const [account, setAccount] = useState(null);
  const [balance, setBalance] = useState(null);
  const [client, setClient] = useState(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  
  const connectKeplr = async () => {
    setLoading(true);
    setError(null);
    
    try {
      const keplr = window.keplr;
      if (!keplr) {
        throw new Error("Keplr wallet not found");
      }
      
      const { client, accounts } = await connectWallet(keplr);
      const account = accounts[0];
      
      setAccount(account.address);
      setClient(client);
      
      // Get balance
      const balanceResponse = await blockchainApi.getBalance(account.address);
      const nmxBalance = balanceResponse.balances.find(b => b.denom === "unmx");
      setBalance(nmxBalance ? (parseInt(nmxBalance.amount) / 1000000).toString() : "0");
      
    } catch (error) {
      console.error("Error connecting wallet:", error);
      setError(error.message);
    } finally {
      setLoading(false);
    }
  };
  
  const disconnectWallet = () => {
    setAccount(null);
    setBalance(null);
    setClient(null);
  };
  
  // Check if Keplr is already connected
  useEffect(() => {
    const checkKeplr = async () => {
      if (window.keplr && window.keplr.getOfflineSigner) {
        try {
          await connectKeplr();
        } catch (error) {
          console.error("Error auto-connecting to Keplr:", error);
        }
      }
    };
    
    checkKeplr();
  }, []);
  
  return (
    <WalletContext.Provider
      value={{
        account,
        balance,
        client,
        loading,
        error,
        connectWallet: connectKeplr,
        disconnectWallet,
      }}
    >
      {children}
    </WalletContext.Provider>
  );
};
```

### 5.2 Transaction Handling

Create a utility for handling transactions:

```javascript
// client/src/utils/transactionUtils.js
import { sendTransaction } from '../api/transactions';
import { MsgDelegate } from "@cosmjs/stargate";

export const stakingActions = {
  delegate: async (client, delegatorAddress, validatorAddress, amount) => {
    const msg = {
      typeUrl: "/cosmos.staking.v1beta1.MsgDelegate",
      value: {
        delegatorAddress,
        validatorAddress,
        amount: {
          denom: "unmx",
          amount: (parseFloat(amount) * 1000000).toString(),
        },
      },
    };
    
    return sendTransaction(client, { address: delegatorAddress }, msg);
  },
  
  undelegate: async (client, delegatorAddress, validatorAddress, amount) => {
    const msg = {
      typeUrl: "/cosmos.staking.v1beta1.MsgUndelegate",
      value: {
        delegatorAddress,
        validatorAddress,
        amount: {
          denom: "unmx",
          amount: (parseFloat(amount) * 1000000).toString(),
        },
      },
    };
    
    return sendTransaction(client, { address: delegatorAddress }, msg);
  },
};

export const governanceActions = {
  submitProposal: async (client, proposerAddress, title, description, deposit) => {
    const msg = {
      typeUrl: "/cosmos.gov.v1beta1.MsgSubmitProposal",
      value: {
        content: {
          typeUrl: "/cosmos.gov.v1beta1.TextProposal",
          value: {
            title,
            description,
          },
        },
        proposer: proposerAddress,
        initialDeposit: [
          {
            denom: "unmx",
            amount: (parseFloat(deposit) * 1000000).toString(),
          },
        ],
      },
    };
    
    return sendTransaction(client, { address: proposerAddress }, msg);
  },
  
  vote: async (client, voterAddress, proposalId, option) => {
    const msg = {
      typeUrl: "/cosmos.gov.v1beta1.MsgVote",
      value: {
        proposalId,
        voter: voterAddress,
        option,
      },
    };
    
    return sendTransaction(client, { address: voterAddress }, msg);
  },
};

// Add similar action objects for other modules
```

## 6. Testing and Deployment

### 6.1 Local Testing

1. Start the local testnet:
```bash
cd scripts/local_testnet
./setup_local_testnet.ps1
```

2. Start the frontend:
```bash
cd client
npm start
```

3. Test all functionality:
   - Wallet connection
   - Token transfers
   - Staking operations
   - Governance participation
   - Smart contract deployment
   - HyperChain creation
   - Oracle queries

### 6.2 Testnet Deployment

1. Set up testnet nodes on cloud servers (AWS, GCP, or Azure)
2. Configure firewall rules to allow required ports
3. Deploy the blockchain nodes
4. Deploy the frontend to a hosting service (Netlify, Vercel, or AWS S3)
5. Configure the frontend to connect to the testnet API endpoints

### 6.3 Monitoring and Maintenance

1. Set up monitoring using Prometheus and Grafana
2. Configure alerts for node downtime or performance issues
3. Implement regular backups of validator keys and data
4. Create a process for updates and upgrades

## 7. Documentation

### 7.1 User Documentation

1. Create a user guide for:
   - Wallet setup
   - Token management
   - Staking and delegation
   - Governance participation
   - Smart contract interaction
   - HyperChain creation
   - Oracle usage

### 7.2 Developer Documentation

1. Create API documentation
2. Document module interfaces and message types
3. Provide examples for common operations
4. Create tutorials for building on NoMercyChain

## 8. Community Building

1. Set up a Discord server for community support
2. Create a testnet faucet for distributing test tokens
3. Organize community testing events
4. Collect feedback for improvements