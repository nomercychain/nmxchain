package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nomercychain/nmxchain/x/hyperchain/types"
)

// InitGenesis initializes the hyperchain module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	// Set all the hyperchains
	for _, hyperchain := range genState.Hyperchains {
		k.SetHyperchain(ctx, hyperchain)
	}

	// Set all the validators
	for _, validator := range genState.Validators {
		k.SetHyperchainValidator(ctx, validator)
	}

	// Set all the blocks
	for _, block := range genState.Blocks {
		k.SetHyperchainBlock(ctx, block)
	}

	// Set all the transactions
	for _, transaction := range genState.Transactions {
		k.SetHyperchainTransaction(ctx, transaction)
	}

	// Set all the bridges
	for _, bridge := range genState.Bridges {
		k.SetHyperchainBridge(ctx, bridge)
	}

	// Set all the bridge transactions
	for _, transaction := range genState.BridgeTransactions {
		k.SetHyperchainBridgeTransaction(ctx, transaction)
	}

	// Set all the permissions
	for _, permission := range genState.Permissions {
		k.SetHyperchainPermission(ctx, permission)
	}

	// Set the params
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the hyperchain module's exported genesis.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// Get all hyperchains
	genesis.Hyperchains = k.GetAllHyperchains(ctx)

	// Get all validators
	for _, hyperchain := range genesis.Hyperchains {
		validators := k.GetHyperchainValidatorsByChain(ctx, hyperchain.Id)
		genesis.Validators = append(genesis.Validators, validators...)
	}

	// Get all blocks
	for _, hyperchain := range genesis.Hyperchains {
		blocks := k.GetHyperchainBlocksByChain(ctx, hyperchain.Id)
		genesis.Blocks = append(genesis.Blocks, blocks...)
	}

	// Get all transactions
	for _, hyperchain := range genesis.Hyperchains {
		transactions := k.GetHyperchainTransactionsByChain(ctx, hyperchain.Id)
		genesis.Transactions = append(genesis.Transactions, transactions...)
	}

	// Get all bridges
	genesis.Bridges = k.GetAllHyperchainBridges(ctx)

	// Get all bridge transactions
	for _, bridge := range genesis.Bridges {
		transactions := k.GetHyperchainBridgeTransactionsByBridge(ctx, bridge.Id)
		genesis.BridgeTransactions = append(genesis.BridgeTransactions, transactions...)
	}

	// Get all permissions
	for _, hyperchain := range genesis.Hyperchains {
		permissions := k.GetHyperchainPermissionsByChain(ctx, hyperchain.Id)
		genesis.Permissions = append(genesis.Permissions, permissions...)
	}

	// Get params
	genesis.Params = k.GetParams(ctx)

	return genesis
}