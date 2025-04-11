package deai

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/nomercychain/nmxchain/x/deai/keeper"
	"github.com/nomercychain/nmxchain/x/deai/types"
)

// NewHandler creates a new handler for the deai module
func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgCreateAIAgent:
			res, err := msgServer.CreateAIAgent(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgUpdateAIAgent:
			res, err := msgServer.UpdateAIAgent(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgTrainAIAgent:
			res, err := msgServer.TrainAIAgent(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgExecuteAIAgent:
			res, err := msgServer.ExecuteAIAgent(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgListAIAgentForSale:
			res, err := msgServer.ListAIAgentForSale(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgBuyAIAgent:
			res, err := msgServer.BuyAIAgent(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgRentAIAgent:
			res, err := msgServer.RentAIAgent(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgCancelMarketListing:
			res, err := msgServer.CancelMarketListing(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// BeginBlocker executes all ABCI BeginBlock logic respective to the deai module.
func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	// Process any pending AI agent training tasks
	processPendingTrainingTasks(ctx, k)
}

// EndBlocker executes all ABCI EndBlock logic respective to the deai module.
func EndBlocker(ctx sdk.Context, req abci.RequestEndBlock, k keeper.Keeper) []abci.ValidatorUpdate {
	// Process expired marketplace listings
	processExpiredListings(ctx, k)
	
	// Process AI agent executions
	processAIAgentExecutions(ctx, k)
	
	return []abci.ValidatorUpdate{}
}

// processPendingTrainingTasks processes any pending AI agent training tasks
func processPendingTrainingTasks(ctx sdk.Context, k keeper.Keeper) {
	// Get all agents in training status
	agents := k.GetAllAIAgents(ctx)
	for _, agent := range agents {
		if agent.Status == types.AIAgentStatusTraining {
			// In a real implementation, we would check if the training is complete
			// For now, we'll just set the agent status to active after a delay
			if ctx.BlockHeight()%10 == 0 { // Simulate training completion every 10 blocks
				agent.Status = types.AIAgentStatusActive
				agent.UpdatedAt = ctx.BlockTime()
				k.SetAIAgent(ctx, agent)
				
				// Emit an event
				ctx.EventManager().EmitEvent(
					sdk.NewEvent(
						"agent_training_completed",
						sdk.NewAttribute("agent_id", agent.ID),
						sdk.NewAttribute("timestamp", ctx.BlockTime().String()),
					),
				)
			}
		}
	}
}

// processExpiredListings processes expired marketplace listings
func processExpiredListings(ctx sdk.Context, k keeper.Keeper) {
	// Get all marketplace listings
	listings := k.GetAllAIAgentMarketplaceListings(ctx)
	for _, listing := range listings {
		if listing.Status == "active" && !listing.ExpiresAt.IsZero() && listing.ExpiresAt.Before(ctx.BlockTime()) {
			// Cancel the listing
			listing.Status = "expired"
			k.SetAIAgentMarketplaceListing(ctx, listing)
			
			// Update the agent status
			agent, found := k.GetAIAgent(ctx, listing.AgentID)
			if found {
				agent.Status = types.AIAgentStatusActive
				agent.UpdatedAt = ctx.BlockTime()
				k.SetAIAgent(ctx, agent)
				
				// Emit an event
				ctx.EventManager().EmitEvent(
					sdk.NewEvent(
						"marketplace_listing_expired",
						sdk.NewAttribute("listing_id", listing.ID),
						sdk.NewAttribute("agent_id", listing.AgentID),
						sdk.NewAttribute("timestamp", ctx.BlockTime().String()),
					),
				)
			}
		}
	}
}

// processAIAgentExecutions processes AI agent executions
func processAIAgentExecutions(ctx sdk.Context, k keeper.Keeper) {
	// In a real implementation, we would process any pending AI agent executions
	// For now, we'll just simulate the completion of executions
	
	// This would typically involve:
	// 1. Fetching pending executions from the store
	// 2. Processing them using the AI model
	// 3. Updating the execution status and result
	// 4. Distributing fees to the agent owner
}