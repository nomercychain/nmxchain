# DynaContract Module

## Overview

The DynaContract module enables dynamic smart contracts on the NoMercyChain blockchain. These contracts can be created, updated, executed, and managed through a set of transactions and queries.

## Features

- **Dynamic Contract Creation**: Create smart contracts with custom code and initialization parameters
- **Contract Execution**: Execute contract functions with optional parameters and funds
- **Contract Updates**: Update contract code and metadata
- **Contract Templates**: Create and instantiate contracts from templates
- **AI Integration**: Connect contracts with AI agents from the DeAI module
- **Learning Capabilities**: Add learning data to adaptive contracts
- **Permission Management**: Grant and revoke permissions for contract operations

## Transactions

### Create Contract

Create a new dynamic contract:

```bash
nmxchaind tx dynacontract create-contract [name] [version] [code-file] [init-msg] \
  --from [key] --chain-id [chain-id]
```

### Update Contract

Update an existing contract:

```bash
nmxchaind tx dynacontract update-contract [id] [name] [version] [code-file] [update-msg] \
  --from [key] --chain-id [chain-id]
```

### Execute Contract

Execute a contract function:

```bash
nmxchaind tx dynacontract execute-contract [id] [execute-msg] [coins] \
  --from [key] --chain-id [chain-id]
```

### Delete Contract

Delete a contract:

```bash
nmxchaind tx dynacontract delete-contract [id] \
  --from [key] --chain-id [chain-id]
```

### Upgrade Contract

Upgrade a contract to a new version:

```bash
nmxchaind tx dynacontract upgrade-contract [id] [version] [code-file] [migrate-msg] \
  --from [key] --chain-id [chain-id]
```

## Queries

### List Contracts

List all contracts:

```bash
nmxchaind query dynacontract list-contracts
```

### Show Contract

Show details of a specific contract:

```bash
nmxchaind query dynacontract show-contract [id]
```

### Query Contract

Query a contract's state:

```bash
nmxchaind query dynacontract query-contract [id] [query-data]
```

### Get Contract History

Get the history of a contract:

```bash
nmxchaind query dynacontract get-contract-history [id]
```

### Get Contract State

Get the current state of a contract:

```bash
nmxchaind query dynacontract get-contract-state [id]
```

## Parameters

The DynaContract module has the following parameters:

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| MaxContractSize | uint64 | 1048576 | Maximum size of contract code in bytes (1MB) |
| MaxContractGas | uint64 | 10000000 | Maximum gas limit for contract execution (10M) |
| MaxLearningDataSize | uint64 | 5242880 | Maximum size of learning data in bytes (5MB) |
| MaxMetadataSize | uint64 | 102400 | Maximum size of metadata in bytes (100KB) |
| MinContractDeposit | sdk.Coin | 100 NMX | Minimum deposit required to create a contract |
| ExecutionFeeRate | sdk.Dec | 0.05 | Percentage of execution fees distributed to contract owners (5%) |

## Events

The DynaContract module emits the following events:

- `create_dyna_contract`: When a new contract is created
- `update_dyna_contract`: When a contract is updated
- `execute_dyna_contract`: When a contract is executed
- `add_learning_data`: When learning data is added to a contract
- `create_dyna_contract_template`: When a new template is created
- `instantiate_from_template`: When a contract is instantiated from a template
- `grant_contract_permission`: When permission is granted
- `revoke_contract_permission`: When permission is revoked