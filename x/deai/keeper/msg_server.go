package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/nomercychain/nmxchain/x/deai/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// CreateAIAgent creates a new AI agent
func (k msgServer) CreateAIAgent(goCtx context.Context, msg *types.MsgCreateAIAgent) (*types.MsgCreateAIAgentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Generate a unique ID for the agent
	id := fmt.Sprintf("%s-%d", msg.Creator.String(), ctx.BlockHeight())

	// Create the agent
	agent := types.AIAgent{
		ID:          id,
		Name:        msg.Name,
		Description: msg.Description,
		Owner:       msg.Creator,
		Creator:     msg.Creator,
		AgentType:   msg.AgentType,
		Status:      types.AIAgentStatusInactive, // Start as inactive until trained
		ModelID:     msg.ModelID,
		CreatedAt:   ctx.BlockTime(),
		UpdatedAt:   ctx.BlockTime(),
		Permissions: msg.Permissions,
		Metadata:    msg.Metadata,
	}

	// Store the agent
	k.SetAIAgent(ctx, agent)

	// Create initial state
	state := types.AIAgentState{
		AgentID:     id,
		StateData:   []byte("{}"),
		StorageType: "chain",
		UpdatedAt:   ctx.BlockTime(),
	}
	k.SetAIAgentState(ctx, state)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"ai_agent_created",
			sdk.NewAttribute("agent_id", id),
			sdk.NewAttribute("creator", msg.Creator.String()),
			sdk.NewAttribute("name", msg.Name),
			sdk.NewAttribute("model_id", msg.ModelID),
		),
	)

	return &types.MsgCreateAIAgentResponse{
		ID: id,
	}, nil
}

// UpdateAIAgent updates an existing AI agent
func (k msgServer) UpdateAIAgent(goCtx context.Context, msg *types.MsgUpdateAIAgent) (*types.MsgUpdateAIAgentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the agent
	agent, found := k.GetAIAgent(ctx, msg.AgentID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("agent not found: %s", msg.AgentID))
	}

	// Check if the caller is the owner
	if !agent.Owner.Equals(msg.Owner) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "only the owner can update the agent")
	}

	// Update the agent
	agent.Name = msg.Name
	agent.Description = msg.Description
	agent.Permissions = msg.Permissions
	agent.Metadata = msg.Metadata
	agent.UpdatedAt = ctx.BlockTime()

	// Store the updated agent
	k.SetAIAgent(ctx, agent)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"ai_agent_updated",
			sdk.NewAttribute("agent_id", msg.AgentID),
			sdk.NewAttribute("owner", msg.Owner.String()),
			sdk.NewAttribute("name", msg.Name),
		),
	)

	return &types.MsgUpdateAIAgentResponse{}, nil
}

// TrainAIAgent trains an AI agent with new data
func (k msgServer) TrainAIAgent(goCtx context.Context, msg *types.MsgTrainAIAgent) (*types.MsgTrainAIAgentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the agent
	agent, found := k.GetAIAgent(ctx, msg.AgentID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("agent not found: %s", msg.AgentID))
	}

	// Check if the caller is the owner
	if !agent.Owner.Equals(msg.Owner) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "only the owner can train the agent")
	}

	// Set agent status to training
	agent.Status = types.AIAgentStatusTraining
	agent.UpdatedAt = ctx.BlockTime()
	k.SetAIAgent(ctx, agent)

	// Create training data record
	trainingData := types.AIAgentTrainingData{
		ID:          fmt.Sprintf("%s-%d", msg.AgentID, ctx.BlockHeight()),
		AgentID:     msg.AgentID,
		DataType:    msg.DataType,
		Data:        msg.Data,
		Source:      msg.Source,
		Timestamp:   ctx.BlockTime(),
		StorageType: "chain",
	}
	k.SetAIAgentTrainingData(ctx, trainingData)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"ai_agent_training_started",
			sdk.NewAttribute("agent_id", msg.AgentID),
			sdk.NewAttribute("owner", msg.Owner.String()),
			sdk.NewAttribute("data_type", msg.DataType),
			sdk.NewAttribute("training_data_id", trainingData.ID),
		),
	)

	return &types.MsgTrainAIAgentResponse{
		TrainingDataID: trainingData.ID,
	}, nil
}

// ExecuteAIAgent executes an action using an AI agent
func (k msgServer) ExecuteAIAgent(goCtx context.Context, msg *types.MsgExecuteAIAgent) (*types.MsgExecuteAIAgentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the agent
	agent, found := k.GetAIAgent(ctx, msg.AgentID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("agent not found: %s", msg.AgentID))
	}

	// Check if the agent is active
	if agent.Status != types.AIAgentStatusActive {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "agent is not active")
	}

	// Check permissions
	// In a real implementation, we would check if the caller has permission to use this agent
	// For simplicity, we'll just check if the caller is the owner or if the agent is public
	if !agent.Owner.Equals(msg.Sender) {
		// Check if the agent has public permissions
		// This would be implemented based on the permissions field
		isPublic := false // Placeholder for permission check
		if !isPublic {
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "no permission to use this agent")
		}
	}

	// Collect the fee
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, msg.Sender, types.ModuleName, sdk.NewCoins(msg.Fee))
	if err != nil {
		return nil, sdkerrors.Wrap(err, "failed to collect fee")
	}

	// Get the agent state
	state, found := k.GetAIAgentState(ctx, msg.AgentID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "agent state not found")
	}

	// In a real implementation, we would:
	// 1. Load the AI model
	// 2. Process the action using the model
	// 3. Update the agent state based on the result
	// 4. Return the result

	// For now, we'll just record the action and return a dummy result
	actionID := fmt.Sprintf("%s-%s-%d", msg.AgentID, msg.ActionType, ctx.BlockHeight())
	action := types.AIAgentAction{
		ID:         actionID,
		AgentID:    msg.AgentID,
		ActionType: msg.ActionType,
		Timestamp:  ctx.BlockTime(),
		Data:       msg.Data,
		Result:     []byte(`{"status": "success", "message": "Action executed successfully"}`),
		Status:     "completed",
		GasUsed:    1000, // Dummy value
	}

	k.SetAIAgentAction(ctx, action)

	// Update the agent state
	state.UpdatedAt = ctx.BlockTime()
	k.SetAIAgentState(ctx, state)

	// Distribute the fee to the agent owner
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, agent.Owner, sdk.NewCoins(msg.Fee))
	if err != nil {
		return nil, sdkerrors.Wrap(err, "failed to distribute fee")
	}

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"ai_agent_executed",
			sdk.NewAttribute("agent_id", msg.AgentID),
			sdk.NewAttribute("sender", msg.Sender.String()),
			sdk.NewAttribute("action_type", msg.ActionType),
			sdk.NewAttribute("action_id", actionID),
			sdk.NewAttribute("fee", msg.Fee.String()),
		),
	)

	return &types.MsgExecuteAIAgentResponse{
		ActionID: actionID,
		Result:   action.Result,
	}, nil
}

// ListAIAgentForSale lists an AI agent for sale on the marketplace
func (k msgServer) ListAIAgentForSale(goCtx context.Context, msg *types.MsgListAIAgentForSale) (*types.MsgListAIAgentForSaleResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the agent
	agent, found := k.GetAIAgent(ctx, msg.AgentID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("agent not found: %s", msg.AgentID))
	}

	// Check if the caller is the owner
	if !agent.Owner.Equals(msg.Seller) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "only the owner can list the agent for sale")
	}

	// Check if the agent is already listed
	if agent.Status == types.AIAgentStatusForSale || agent.Status == types.AIAgentStatusForRent {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "agent is already listed on the marketplace")
	}

	// Generate a unique ID for the listing
	listingID := fmt.Sprintf("%s-%d", msg.AgentID, ctx.BlockHeight())

	// Calculate expiration time
	expiresAt := ctx.BlockTime().Add(time.Hour * 24 * time.Duration(msg.ExpirationDays))

	// Create the listing
	listing := types.AIAgentMarketplaceListing{
		ID:              listingID,
		AgentID:         msg.AgentID,
		Seller:          msg.Seller,
		Price:           msg.Price,
		RentalPrice:     msg.RentalPrice,
		RentalDuration:  msg.RentalDuration,
		ListingType:     msg.ListingType,
		Status:          "active",
		CreatedAt:       ctx.BlockTime(),
		ExpiresAt:       expiresAt,
	}

	// Update the agent status
	if msg.ListingType == "sale" {
		agent.Status = types.AIAgentStatusForSale
	} else if msg.ListingType == "rent" {
		agent.Status = types.AIAgentStatusForRent
	} else {
		agent.Status = types.AIAgentStatusForSale
	}
	agent.UpdatedAt = ctx.BlockTime()

	// Store the listing and updated agent
	k.SetAIAgentMarketplaceListing(ctx, listing)
	k.SetAIAgent(ctx, agent)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"ai_agent_listed",
			sdk.NewAttribute("listing_id", listingID),
			sdk.NewAttribute("agent_id", msg.AgentID),
			sdk.NewAttribute("seller", msg.Seller.String()),
			sdk.NewAttribute("listing_type", msg.ListingType),
			sdk.NewAttribute("expires_at", expiresAt.String()),
		),
	)

	return &types.MsgListAIAgentForSaleResponse{
		ListingID: listingID,
	}, nil
}

// BuyAIAgent buys an AI agent from the marketplace
func (k msgServer) BuyAIAgent(goCtx context.Context, msg *types.MsgBuyAIAgent) (*types.MsgBuyAIAgentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the listing
	listing, found := k.GetAIAgentMarketplaceListing(ctx, msg.ListingID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("listing not found: %s", msg.ListingID))
	}

	// Check if the listing is active
	if listing.Status != "active" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "listing is not active")
	}

	// Check if the listing is for sale
	if listing.ListingType != "sale" && listing.ListingType != "both" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "agent is not for sale")
	}

	// Get the agent
	agent, found := k.GetAIAgent(ctx, listing.AgentID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("agent not found: %s", listing.AgentID))
	}

	// Transfer payment from buyer to seller
	err := k.bankKeeper.SendCoinsFromAccountToAccount(ctx, msg.Buyer, listing.Seller, listing.Price)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "failed to transfer payment")
	}

	// Update the agent owner
	agent.Owner = msg.Buyer
	agent.Status = types.AIAgentStatusActive
	agent.UpdatedAt = ctx.BlockTime()

	// Update the listing status
	listing.Status = "completed"

	// Store the updated agent and listing
	k.SetAIAgent(ctx, agent)
	k.SetAIAgentMarketplaceListing(ctx, listing)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"ai_agent_sold",
			sdk.NewAttribute("listing_id", msg.ListingID),
			sdk.NewAttribute("agent_id", listing.AgentID),
			sdk.NewAttribute("seller", listing.Seller.String()),
			sdk.NewAttribute("buyer", msg.Buyer.String()),
			sdk.NewAttribute("price", listing.Price.String()),
		),
	)

	return &types.MsgBuyAIAgentResponse{
		AgentID: listing.AgentID,
	}, nil
}

// RentAIAgent rents an AI agent from the marketplace
func (k msgServer) RentAIAgent(goCtx context.Context, msg *types.MsgRentAIAgent) (*types.MsgRentAIAgentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the listing
	listing, found := k.GetAIAgentMarketplaceListing(ctx, msg.ListingID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("listing not found: %s", msg.ListingID))
	}

	// Check if the listing is active
	if listing.Status != "active" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "listing is not active")
	}

	// Check if the listing is for rent
	if listing.ListingType != "rent" && listing.ListingType != "both" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "agent is not for rent")
	}

	// Check if the requested duration is valid
	if msg.Duration > listing.RentalDuration {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "requested duration exceeds maximum rental duration")
	}

	// Get the agent
	agent, found := k.GetAIAgent(ctx, listing.AgentID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("agent not found: %s", listing.AgentID))
	}

	// Calculate the rental price based on the requested duration
	rentalPrice := listing.RentalPrice
	if msg.Duration < listing.RentalDuration {
		// Adjust price proportionally
		for i := range rentalPrice {
			amount := rentalPrice[i].Amount.MulRaw(int64(msg.Duration)).QuoRaw(int64(listing.RentalDuration))
			rentalPrice[i] = sdk.NewCoin(rentalPrice[i].Denom, amount)
		}
	}

	// Transfer payment from renter to owner
	err := k.bankKeeper.SendCoinsFromAccountToAccount(ctx, msg.Renter, listing.Seller, rentalPrice)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "failed to transfer payment")
	}

	// Create a rental record
	// In a real implementation, we would create a record of the rental
	// and grant temporary access to the agent

	// For now, we'll just emit an event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"ai_agent_rented",
			sdk.NewAttribute("listing_id", msg.ListingID),
			sdk.NewAttribute("agent_id", listing.AgentID),
			sdk.NewAttribute("owner", listing.Seller.String()),
			sdk.NewAttribute("renter", msg.Renter.String()),
			sdk.NewAttribute("duration", fmt.Sprintf("%d", msg.Duration)),
			sdk.NewAttribute("price", rentalPrice.String()),
		),
	)

	return &types.MsgRentAIAgentResponse{
		AgentID:  listing.AgentID,
		Duration: msg.Duration,
	}, nil
}

// CancelMarketListing cancels a marketplace listing
func (k msgServer) CancelMarketListing(goCtx context.Context, msg *types.MsgCancelMarketListing) (*types.MsgCancelMarketListingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the listing
	listing, found := k.GetAIAgentMarketplaceListing(ctx, msg.ListingID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("listing not found: %s", msg.ListingID))
	}

	// Check if the listing is active
	if listing.Status != "active" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "listing is not active")
	}

	// Check if the caller is the seller
	if !listing.Seller.Equals(msg.Owner) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "only the seller can cancel the listing")
	}

	// Get the agent
	agent, found := k.GetAIAgent(ctx, listing.AgentID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("agent not found: %s", listing.AgentID))
	}

	// Update the listing status
	listing.Status = "cancelled"

	// Update the agent status
	agent.Status = types.AIAgentStatusActive
	agent.UpdatedAt = ctx.BlockTime()

	// Store the updated listing and agent
	k.SetAIAgentMarketplaceListing(ctx, listing)
	k.SetAIAgent(ctx, agent)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"marketplace_listing_cancelled",
			sdk.NewAttribute("listing_id", msg.ListingID),
			sdk.NewAttribute("agent_id", listing.AgentID),
			sdk.NewAttribute("seller", listing.Seller.String()),
		),
	)

	return &types.MsgCancelMarketListingResponse{}, nil
}