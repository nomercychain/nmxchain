package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nomercychain/nmxchain/x/deai/types"
)

// InitGenesis initializes the deai module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) []abci.ValidatorUpdate {
	// Set all the agents
	for _, agent := range genState.Agents {
		k.SetAIAgent(ctx, agent)
	}

	// Set all the models
	for _, model := range genState.Models {
		k.SetAIAgentModel(ctx, model)
	}

	// Set all the states
	for _, state := range genState.States {
		k.SetAIAgentState(ctx, state)
	}

	// Set all the actions
	for _, action := range genState.Actions {
		k.SetAIAgentAction(ctx, action)
	}

	// Set all the training data
	for _, data := range genState.TrainingData {
		k.SetAIAgentTrainingData(ctx, data)
	}

	// Set all the marketplace listings
	for _, listing := range genState.MarketplaceListings {
		k.SetAIAgentMarketplaceListing(ctx, listing)
	}

	// Set module parameters
	k.SetParams(ctx, genState.Params)

	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the deai module's exported genesis.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return &types.GenesisState{
		Agents:              k.GetAllAIAgents(ctx),
		Models:              k.GetAllAIAgentModels(ctx),
		States:              k.GetAllAIAgentStates(ctx),
		Actions:             k.GetAllAIAgentActions(ctx),
		TrainingData:        k.GetAllAIAgentTrainingData(ctx),
		MarketplaceListings: k.GetAllAIAgentMarketplaceListings(ctx),
		Params:              k.GetParams(ctx),
	}
}

// GetAllAIAgentStates returns all AI agent states
func (k Keeper) GetAllAIAgentStates(ctx sdk.Context) []types.AIAgentState {
	var states []types.AIAgentState
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.AIAgentStateKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var state types.AIAgentState
		k.cdc.MustUnmarshal(iterator.Value(), &state)
		states = append(states, state)
	}

	return states
}

// GetAllAIAgentActions returns all AI agent actions
func (k Keeper) GetAllAIAgentActions(ctx sdk.Context) []types.AIAgentAction {
	var actions []types.AIAgentAction
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.AIAgentActionKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var action types.AIAgentAction
		k.cdc.MustUnmarshal(iterator.Value(), &action)
		actions = append(actions, action)
	}

	return actions
}

// GetAllAIAgentTrainingData returns all AI agent training data
func (k Keeper) GetAllAIAgentTrainingData(ctx sdk.Context) []types.AIAgentTrainingData {
	var dataList []types.AIAgentTrainingData
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.AIAgentTrainingKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var data types.AIAgentTrainingData
		k.cdc.MustUnmarshal(iterator.Value(), &data)
		dataList = append(dataList, data)
	}

	return dataList
}