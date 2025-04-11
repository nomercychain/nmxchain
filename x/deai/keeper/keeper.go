package keeper

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/nomercychain/nmxchain/x/deai/types"
)

// Keeper of the deai store
type Keeper struct {
	storeKey   sdk.StoreKey
	memKey     sdk.StoreKey
	cdc        codec.BinaryCodec
	paramstore paramtypes.Subspace

	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
	nftKeeper     types.NFTKeeper
}

// NewKeeper creates a new deai Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	nftKeeper types.NFTKeeper,
) Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		storeKey:      storeKey,
		memKey:        memKey,
		cdc:           cdc,
		paramstore:    ps,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
		nftKeeper:     nftKeeper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// SetAIAgent sets an AI agent
func (k Keeper) SetAIAgent(ctx sdk.Context, agent types.AIAgent) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.AIAgentKey, []byte(agent.ID)...)
	value := k.cdc.MustMarshal(&agent)
	store.Set(key, value)
}

// GetAIAgent returns an AI agent by ID
func (k Keeper) GetAIAgent(ctx sdk.Context, id string) (types.AIAgent, bool) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.AIAgentKey, []byte(id)...)
	value := store.Get(key)
	if value == nil {
		return types.AIAgent{}, false
	}

	var agent types.AIAgent
	k.cdc.MustUnmarshal(value, &agent)
	return agent, true
}

// DeleteAIAgent deletes an AI agent
func (k Keeper) DeleteAIAgent(ctx sdk.Context, id string) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.AIAgentKey, []byte(id)...)
	store.Delete(key)
}

// GetAllAIAgents returns all AI agents
func (k Keeper) GetAllAIAgents(ctx sdk.Context) []types.AIAgent {
	var agents []types.AIAgent
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.AIAgentKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var agent types.AIAgent
		k.cdc.MustUnmarshal(iterator.Value(), &agent)
		agents = append(agents, agent)
	}

	return agents
}

// GetAIAgentsByOwner returns all AI agents owned by an address
func (k Keeper) GetAIAgentsByOwner(ctx sdk.Context, owner sdk.AccAddress) []types.AIAgent {
	var agents []types.AIAgent
	allAgents := k.GetAllAIAgents(ctx)

	for _, agent := range allAgents {
		if agent.Owner.Equals(owner) {
			agents = append(agents, agent)
		}
	}

	return agents
}

// SetAIAgentModel sets an AI agent model
func (k Keeper) SetAIAgentModel(ctx sdk.Context, model types.AIAgentModel) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.AIAgentModelKey, []byte(model.ID)...)
	value := k.cdc.MustMarshal(&model)
	store.Set(key, value)
}

// GetAIAgentModel returns an AI agent model by ID
func (k Keeper) GetAIAgentModel(ctx sdk.Context, id string) (types.AIAgentModel, bool) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.AIAgentModelKey, []byte(id)...)
	value := store.Get(key)
	if value == nil {
		return types.AIAgentModel{}, false
	}

	var model types.AIAgentModel
	k.cdc.MustUnmarshal(value, &model)
	return model, true
}

// GetAllAIAgentModels returns all AI agent models
func (k Keeper) GetAllAIAgentModels(ctx sdk.Context) []types.AIAgentModel {
	var models []types.AIAgentModel
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.AIAgentModelKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var model types.AIAgentModel
		k.cdc.MustUnmarshal(iterator.Value(), &model)
		models = append(models, model)
	}

	return models
}

// SetAIAgentState sets an AI agent state
func (k Keeper) SetAIAgentState(ctx sdk.Context, state types.AIAgentState) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.AIAgentStateKey, []byte(state.AgentID)...)
	value := k.cdc.MustMarshal(&state)
	store.Set(key, value)
}

// GetAIAgentState returns an AI agent state by agent ID
func (k Keeper) GetAIAgentState(ctx sdk.Context, agentID string) (types.AIAgentState, bool) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.AIAgentStateKey, []byte(agentID)...)
	value := store.Get(key)
	if value == nil {
		return types.AIAgentState{}, false
	}

	var state types.AIAgentState
	k.cdc.MustUnmarshal(value, &state)
	return state, true
}

// SetAIAgentAction records an AI agent action
func (k Keeper) SetAIAgentAction(ctx sdk.Context, action types.AIAgentAction) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.AIAgentActionKey, []byte(action.ID)...)
	value := k.cdc.MustMarshal(&action)
	store.Set(key, value)
}

// GetAIAgentAction returns an AI agent action by ID
func (k Keeper) GetAIAgentAction(ctx sdk.Context, id string) (types.AIAgentAction, bool) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.AIAgentActionKey, []byte(id)...)
	value := store.Get(key)
	if value == nil {
		return types.AIAgentAction{}, false
	}

	var action types.AIAgentAction
	k.cdc.MustUnmarshal(value, &action)
	return action, true
}

// GetAIAgentActionsByAgent returns all actions for an AI agent
func (k Keeper) GetAIAgentActionsByAgent(ctx sdk.Context, agentID string) []types.AIAgentAction {
	var actions []types.AIAgentAction
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.AIAgentActionKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var action types.AIAgentAction
		k.cdc.MustUnmarshal(iterator.Value(), &action)
		if action.AgentID == agentID {
			actions = append(actions, action)
		}
	}

	return actions
}

// SetAIAgentTrainingData sets AI agent training data
func (k Keeper) SetAIAgentTrainingData(ctx sdk.Context, data types.AIAgentTrainingData) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.AIAgentTrainingKey, []byte(data.ID)...)
	value := k.cdc.MustMarshal(&data)
	store.Set(key, value)
}

// GetAIAgentTrainingData returns AI agent training data by ID
func (k Keeper) GetAIAgentTrainingData(ctx sdk.Context, id string) (types.AIAgentTrainingData, bool) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.AIAgentTrainingKey, []byte(id)...)
	value := store.Get(key)
	if value == nil {
		return types.AIAgentTrainingData{}, false
	}

	var data types.AIAgentTrainingData
	k.cdc.MustUnmarshal(value, &data)
	return data, true
}

// GetAIAgentTrainingDataByAgent returns all training data for an AI agent
func (k Keeper) GetAIAgentTrainingDataByAgent(ctx sdk.Context, agentID string) []types.AIAgentTrainingData {
	var dataList []types.AIAgentTrainingData
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.AIAgentTrainingKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var data types.AIAgentTrainingData
		k.cdc.MustUnmarshal(iterator.Value(), &data)
		if data.AgentID == agentID {
			dataList = append(dataList, data)
		}
	}

	return dataList
}

// SetAIAgentMarketplaceListing sets an AI agent marketplace listing
func (k Keeper) SetAIAgentMarketplaceListing(ctx sdk.Context, listing types.AIAgentMarketplaceListing) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.AIAgentMarketplaceKey, []byte(listing.ID)...)
	value := k.cdc.MustMarshal(&listing)
	store.Set(key, value)
}

// GetAIAgentMarketplaceListing returns an AI agent marketplace listing by ID
func (k Keeper) GetAIAgentMarketplaceListing(ctx sdk.Context, id string) (types.AIAgentMarketplaceListing, bool) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.AIAgentMarketplaceKey, []byte(id)...)
	value := store.Get(key)
	if value == nil {
		return types.AIAgentMarketplaceListing{}, false
	}

	var listing types.AIAgentMarketplaceListing
	k.cdc.MustUnmarshal(value, &listing)
	return listing, true
}

// GetAllAIAgentMarketplaceListings returns all AI agent marketplace listings
func (k Keeper) GetAllAIAgentMarketplaceListings(ctx sdk.Context) []types.AIAgentMarketplaceListing {
	var listings []types.AIAgentMarketplaceListing
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.AIAgentMarketplaceKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var listing types.AIAgentMarketplaceListing
		k.cdc.MustUnmarshal(iterator.Value(), &listing)
		listings = append(listings, listing)
	}

	return listings
}

// CreateAIAgent creates a new AI agent
func (k Keeper) CreateAIAgent(ctx sdk.Context, creator sdk.AccAddress, name string, description string, agentType types.AIAgentType, modelID string, permissions json.RawMessage, metadata json.RawMessage) (string, error) {
	// Generate a unique ID for the agent
	id := fmt.Sprintf("%s-%d", creator.String(), ctx.BlockHeight())

	// Create the agent
	agent := types.AIAgent{
		ID:          id,
		Name:        name,
		Description: description,
		Owner:       creator,
		Creator:     creator,
		AgentType:   agentType,
		Status:      types.AIAgentStatusInactive, // Start as inactive until trained
		ModelID:     modelID,
		CreatedAt:   ctx.BlockTime(),
		UpdatedAt:   ctx.BlockTime(),
		Permissions: permissions,
		Metadata:    metadata,
	}

	// Store the agent
	k.SetAIAgent(ctx, agent)

	// Create initial state
	state := types.AIAgentState{
		AgentID:     id,
		StateData:   json.RawMessage(`{}`),
		StorageType: "chain",
		UpdatedAt:   ctx.BlockTime(),
	}
	k.SetAIAgentState(ctx, state)

	return id, nil
}

// TrainAIAgent trains an AI agent with new data
func (k Keeper) TrainAIAgent(ctx sdk.Context, agentID string, owner sdk.AccAddress, dataType string, data json.RawMessage, source string) error {
	// Get the agent
	agent, found := k.GetAIAgent(ctx, agentID)
	if !found {
		return fmt.Errorf("agent not found: %s", agentID)
	}

	// Check if the caller is the owner
	if !agent.Owner.Equals(owner) {
		return fmt.Errorf("only the owner can train the agent")
	}

	// Set agent status to training
	agent.Status = types.AIAgentStatusTraining
	agent.UpdatedAt = ctx.BlockTime()
	k.SetAIAgent(ctx, agent)

	// Create training data record
	trainingData := types.AIAgentTrainingData{
		ID:          fmt.Sprintf("%s-%d", agentID, ctx.BlockHeight()),
		AgentID:     agentID,
		DataType:    dataType,
		Data:        data,
		Source:      source,
		Timestamp:   ctx.BlockTime(),
		StorageType: "chain",
	}
	k.SetAIAgentTrainingData(ctx, trainingData)

	// In a real implementation, we would initiate the training process here
	// For now, we'll just update the agent status after a delay

	// This would be handled by an external service or in EndBlocker
	// For this example, we'll just set it back to active
	agent.Status = types.AIAgentStatusActive
	k.SetAIAgent(ctx, agent)

	return nil
}

// ExecuteAIAgentAction executes an action using an AI agent
func (k Keeper) ExecuteAIAgentAction(ctx sdk.Context, agentID string, caller sdk.AccAddress, actionType string, actionData json.RawMessage) (json.RawMessage, error) {
	// Get the agent
	agent, found := k.GetAIAgent(ctx, agentID)
	if !found {
		return nil, fmt.Errorf("agent not found: %s", agentID)
	}

	// Check if the agent is active
	if agent.Status != types.AIAgentStatusActive {
		return nil, fmt.Errorf("agent is not active")
	}

	// Check permissions
	// In a real implementation, we would check if the caller has permission to use this agent
	// For simplicity, we'll just check if the caller is the owner
	if !agent.Owner.Equals(caller) {
		return nil, fmt.Errorf("only the owner can use this agent")
	}

	// Get the agent state
	state, found := k.GetAIAgentState(ctx, agentID)
	if !found {
		return nil, fmt.Errorf("agent state not found")
	}

	// In a real implementation, we would:
	// 1. Load the AI model
	// 2. Process the action using the model
	// 3. Update the agent state based on the result
	// 4. Return the result

	// For now, we'll just record the action and return a dummy result
	action := types.AIAgentAction{
		ID:         fmt.Sprintf("%s-%s-%d", agentID, actionType, ctx.BlockHeight()),
		AgentID:    agentID,
		ActionType: actionType,
		Timestamp:  ctx.BlockTime(),
		Data:       actionData,
		Result:     json.RawMessage(`{"status": "success"}`),
		Status:     "completed",
		GasUsed:    1000, // Dummy value
	}

	k.SetAIAgentAction(ctx, action)

	// Update the agent state
	state.UpdatedAt = ctx.BlockTime()
	k.SetAIAgentState(ctx, state)

	return action.Result, nil
}

// ListAIAgentForSale lists an AI agent for sale on the marketplace
func (k Keeper) ListAIAgentForSale(ctx sdk.Context, agentID string, seller sdk.AccAddress, price sdk.Coins, rentalPrice sdk.Coins, rentalDuration uint64, listingType string, expiresAt time.Time) (string, error) {
	// Get the agent
	agent, found := k.GetAIAgent(ctx, agentID)
	if !found {
		return "", fmt.Errorf("agent not found: %s", agentID)
	}

	// Check if the caller is the owner
	if !agent.Owner.Equals(seller) {
		return "", fmt.Errorf("only the owner can list the agent for sale")
	}

	// Check if the agent is already listed
	if agent.Status == types.AIAgentStatusForSale || agent.Status == types.AIAgentStatusForRent {
		return "", fmt.Errorf("agent is already listed on the marketplace")
	}

	// Generate a unique ID for the listing
	listingID := fmt.Sprintf("%s-%d", agentID, ctx.BlockHeight())

	// Create the listing
	listing := types.AIAgentMarketplaceListing{
		ID:              listingID,
		AgentID:         agentID,
		Seller:          seller,
		Price:           price,
		RentalPrice:     rentalPrice,
		RentalDuration:  rentalDuration,
		ListingType:     listingType,
		Status:          "active",
		CreatedAt:       ctx.BlockTime(),
		ExpiresAt:       expiresAt,
	}

	// Update the agent status
	if listingType == "sale" {
		agent.Status = types.AIAgentStatusForSale
	} else if listingType == "rent" {
		agent.Status = types.AIAgentStatusForRent
	} else {
		agent.Status = types.AIAgentStatusForSale
	}
	agent.UpdatedAt = ctx.BlockTime()

	// Store the listing and updated agent
	k.SetAIAgentMarketplaceListing(ctx, listing)
	k.SetAIAgent(ctx, agent)

	return listingID, nil
}

// BuyAIAgent buys an AI agent from the marketplace
func (k Keeper) BuyAIAgent(ctx sdk.Context, listingID string, buyer sdk.AccAddress) error {
	// Get the listing
	listing, found := k.GetAIAgentMarketplaceListing(ctx, listingID)
	if !found {
		return fmt.Errorf("listing not found: %s", listingID)
	}

	// Check if the listing is active
	if listing.Status != "active" {
		return fmt.Errorf("listing is not active")
	}

	// Check if the listing is for sale
	if listing.ListingType != "sale" && listing.ListingType != "both" {
		return fmt.Errorf("agent is not for sale")
	}

	// Get the agent
	agent, found := k.GetAIAgent(ctx, listing.AgentID)
	if !found {
		return fmt.Errorf("agent not found: %s", listing.AgentID)
	}

	// Transfer payment from buyer to seller
	err := k.bankKeeper.SendCoinsFromAccountToAccount(ctx, buyer, listing.Seller, listing.Price)
	if err != nil {
		return err
	}

	// Update the agent owner
	agent.Owner = buyer
	agent.Status = types.AIAgentStatusActive
	agent.UpdatedAt = ctx.BlockTime()

	// Update the listing status
	listing.Status = "completed"

	// Store the updated agent and listing
	k.SetAIAgent(ctx, agent)
	k.SetAIAgentMarketplaceListing(ctx, listing)

	return nil
}

// RentAIAgent rents an AI agent from the marketplace
func (k Keeper) RentAIAgent(ctx sdk.Context, listingID string, renter sdk.AccAddress) error {
	// Get the listing
	listing, found := k.GetAIAgentMarketplaceListing(ctx, listingID)
	if !found {
		return fmt.Errorf("listing not found: %s", listingID)
	}

	// Check if the listing is active
	if listing.Status != "active" {
		return fmt.Errorf("listing is not active")
	}

	// Check if the listing is for rent
	if listing.ListingType != "rent" && listing.ListingType != "both" {
		return fmt.Errorf("agent is not for rent")
	}

	// Get the agent
	agent, found := k.GetAIAgent(ctx, listing.AgentID)
	if !found {
		return fmt.Errorf("agent not found: %s", listing.AgentID)
	}

	// Transfer rental payment from renter to seller
	err := k.bankKeeper.SendCoinsFromAccountToAccount(ctx, renter, listing.Seller, listing.RentalPrice)
	if err != nil {
		return err
	}

	// In a real implementation, we would:
	// 1. Create a rental agreement
	// 2. Set up a temporary permission for the renter
	// 3. Schedule the return of the agent after the rental period

	// For now, we'll just update the listing status
	listing.Status = "rented"

	// Store the updated listing
	k.SetAIAgentMarketplaceListing(ctx, listing)

	return nil
}

// ProcessAIAgentActions processes pending AI agent actions
func (k Keeper) ProcessAIAgentActions(ctx sdk.Context) {
	// In a real implementation, this would:
	// 1. Get all active AI agents
	// 2. For each agent, check if it has any scheduled or automated actions
	// 3. Execute those actions based on the agent's AI model and current state
	// 4. Update the agent state and record the actions

	// For now, this is just a placeholder
}