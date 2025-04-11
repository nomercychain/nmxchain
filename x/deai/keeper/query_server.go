package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nomercychain/nmxchain/x/deai/types"
)

type queryServer struct {
	Keeper
}

// NewQueryServerImpl returns an implementation of the QueryServer interface
// for the provided Keeper.
func NewQueryServerImpl(keeper Keeper) types.QueryServer {
	return &queryServer{Keeper: keeper}
}

var _ types.QueryServer = queryServer{}

// AIAgent returns information about a specific AI agent
func (k queryServer) AIAgent(goCtx context.Context, req *types.QueryAIAgentRequest) (*types.QueryAIAgentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	agent, found := k.GetAIAgent(ctx, req.ID)
	if !found {
		return nil, status.Error(codes.NotFound, "AI agent not found")
	}

	return &types.QueryAIAgentResponse{
		Agent: agent,
	}, nil
}

// AIAgents returns all AI agents
func (k queryServer) AIAgents(goCtx context.Context, req *types.QueryAIAgentsRequest) (*types.QueryAIAgentsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	
	var agents []types.AIAgent
	
	// If owner is specified, filter by owner
	if req.Owner != "" {
		ownerAddr, err := sdk.AccAddressFromBech32(req.Owner)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "invalid owner address")
		}
		agents = k.GetAIAgentsByOwner(ctx, ownerAddr)
	} else {
		agents = k.GetAllAIAgents(ctx)
	}

	return &types.QueryAIAgentsResponse{
		Agents: agents,
	}, nil
}

// AIAgentState returns the state of a specific AI agent
func (k queryServer) AIAgentState(goCtx context.Context, req *types.QueryAIAgentStateRequest) (*types.QueryAIAgentStateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	state, found := k.GetAIAgentState(ctx, req.AgentID)
	if !found {
		return nil, status.Error(codes.NotFound, "AI agent state not found")
	}

	return &types.QueryAIAgentStateResponse{
		State: state,
	}, nil
}

// AIAgentActions returns actions for a specific AI agent
func (k queryServer) AIAgentActions(goCtx context.Context, req *types.QueryAIAgentActionsRequest) (*types.QueryAIAgentActionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	actions := k.GetAIAgentActionsByAgent(ctx, req.AgentID)

	return &types.QueryAIAgentActionsResponse{
		Actions: actions,
	}, nil
}

// AIAgentAction returns a specific AI agent action
func (k queryServer) AIAgentAction(goCtx context.Context, req *types.QueryAIAgentActionRequest) (*types.QueryAIAgentActionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	action, found := k.GetAIAgentAction(ctx, req.ID)
	if !found {
		return nil, status.Error(codes.NotFound, "AI agent action not found")
	}

	return &types.QueryAIAgentActionResponse{
		Action: action,
	}, nil
}

// AIAgentModels returns all AI agent models
func (k queryServer) AIAgentModels(goCtx context.Context, req *types.QueryAIAgentModelsRequest) (*types.QueryAIAgentModelsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	models := k.GetAllAIAgentModels(ctx)

	return &types.QueryAIAgentModelsResponse{
		Models: models,
	}, nil
}

// AIAgentModel returns a specific AI agent model
func (k queryServer) AIAgentModel(goCtx context.Context, req *types.QueryAIAgentModelRequest) (*types.QueryAIAgentModelResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	model, found := k.GetAIAgentModel(ctx, req.ID)
	if !found {
		return nil, status.Error(codes.NotFound, "AI agent model not found")
	}

	return &types.QueryAIAgentModelResponse{
		Model: model,
	}, nil
}

// AIAgentTrainingData returns training data for a specific AI agent
func (k queryServer) AIAgentTrainingData(goCtx context.Context, req *types.QueryAIAgentTrainingDataRequest) (*types.QueryAIAgentTrainingDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	trainingData := k.GetAIAgentTrainingDataByAgent(ctx, req.AgentID)

	return &types.QueryAIAgentTrainingDataResponse{
		TrainingData: trainingData,
	}, nil
}

// MarketplaceListings returns all marketplace listings
func (k queryServer) MarketplaceListings(goCtx context.Context, req *types.QueryMarketplaceListingsRequest) (*types.QueryMarketplaceListingsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	listings := k.GetAllAIAgentMarketplaceListings(ctx)
	
	// Filter by status if specified
	if req.Status != "" {
		var filteredListings []types.AIAgentMarketplaceListing
		for _, listing := range listings {
			if listing.Status == req.Status {
				filteredListings = append(filteredListings, listing)
			}
		}
		listings = filteredListings
	}
	
	// Filter by listing type if specified
	if req.ListingType != "" {
		var filteredListings []types.AIAgentMarketplaceListing
		for _, listing := range listings {
			if listing.ListingType == req.ListingType {
				filteredListings = append(filteredListings, listing)
			}
		}
		listings = filteredListings
	}

	return &types.QueryMarketplaceListingsResponse{
		Listings: listings,
	}, nil
}

// MarketplaceListing returns a specific marketplace listing
func (k queryServer) MarketplaceListing(goCtx context.Context, req *types.QueryMarketplaceListingRequest) (*types.QueryMarketplaceListingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	listing, found := k.GetAIAgentMarketplaceListing(ctx, req.ID)
	if !found {
		return nil, status.Error(codes.NotFound, "marketplace listing not found")
	}

	return &types.QueryMarketplaceListingResponse{
		Listing: listing,
	}, nil
}