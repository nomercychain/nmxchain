package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/nomercychain/nmxchain/x/deai/types"
)

var _ types.QueryServer = Keeper{}

// AIAgent returns information about a specific AI agent
func (k Keeper) AIAgent(c context.Context, req *types.QueryAIAgentRequest) (*types.QueryAIAgentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	agent, found := k.GetAIAgent(ctx, req.Id)
	if !found {
		return nil, status.Error(codes.NotFound, "AI agent not found")
	}

	return &types.QueryAIAgentResponse{
		Agent: agent,
	}, nil
}

// AIAgents returns all AI agents
func (k Keeper) AIAgents(c context.Context, req *types.QueryAIAgentsRequest) (*types.QueryAIAgentsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	var agents []types.AIAgent
	var pageRes *query.PageResponse
	var err error

	if req.Owner != "" {
		owner, err := sdk.AccAddressFromBech32(req.Owner)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "invalid owner address")
		}
		agents = k.GetAIAgentsByOwner(ctx, owner)
	} else {
		agents = k.GetAllAIAgents(ctx)
	}

	// Apply pagination
	if req.Pagination != nil {
		start, end := query.Paginate(len(agents), req.Pagination.Offset, req.Pagination.Limit, 100)
		if start < 0 || end < 0 {
			agents = []types.AIAgent{}
		} else {
			agents = agents[start:end]
		}
		pageRes, err = query.NewPaginationResponse(uint64(len(agents)), req.Pagination.Offset, req.Pagination.Limit)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &types.QueryAIAgentsResponse{
		Agents:     agents,
		Pagination: pageRes,
	}, nil
}

// AIAgentState returns the state of a specific AI agent
func (k Keeper) AIAgentState(c context.Context, req *types.QueryAIAgentStateRequest) (*types.QueryAIAgentStateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	state, found := k.GetAIAgentState(ctx, req.AgentId)
	if !found {
		return nil, status.Error(codes.NotFound, "AI agent state not found")
	}

	return &types.QueryAIAgentStateResponse{
		State: state,
	}, nil
}

// AIAgentActions returns actions for a specific AI agent
func (k Keeper) AIAgentActions(c context.Context, req *types.QueryAIAgentActionsRequest) (*types.QueryAIAgentActionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	actions := k.GetAIAgentActionsByAgent(ctx, req.AgentId)
	var pageRes *query.PageResponse
	var err error

	// Apply pagination
	if req.Pagination != nil {
		start, end := query.Paginate(len(actions), req.Pagination.Offset, req.Pagination.Limit, 100)
		if start < 0 || end < 0 {
			actions = []types.AIAgentAction{}
		} else {
			actions = actions[start:end]
		}
		pageRes, err = query.NewPaginationResponse(uint64(len(actions)), req.Pagination.Offset, req.Pagination.Limit)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &types.QueryAIAgentActionsResponse{
		Actions:    actions,
		Pagination: pageRes,
	}, nil
}

// AIAgentAction returns a specific AI agent action
func (k Keeper) AIAgentAction(c context.Context, req *types.QueryAIAgentActionRequest) (*types.QueryAIAgentActionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	action, found := k.GetAIAgentAction(ctx, req.Id)
	if !found {
		return nil, status.Error(codes.NotFound, "AI agent action not found")
	}

	return &types.QueryAIAgentActionResponse{
		Action: action,
	}, nil
}

// AIAgentModels returns all AI agent models
func (k Keeper) AIAgentModels(c context.Context, req *types.QueryAIAgentModelsRequest) (*types.QueryAIAgentModelsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	models := k.GetAllAIAgentModels(ctx)
	var pageRes *query.PageResponse
	var err error

	// Apply pagination
	if req.Pagination != nil {
		start, end := query.Paginate(len(models), req.Pagination.Offset, req.Pagination.Limit, 100)
		if start < 0 || end < 0 {
			models = []types.AIAgentModel{}
		} else {
			models = models[start:end]
		}
		pageRes, err = query.NewPaginationResponse(uint64(len(models)), req.Pagination.Offset, req.Pagination.Limit)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &types.QueryAIAgentModelsResponse{
		Models:     models,
		Pagination: pageRes,
	}, nil
}

// AIAgentModel returns a specific AI agent model
func (k Keeper) AIAgentModel(c context.Context, req *types.QueryAIAgentModelRequest) (*types.QueryAIAgentModelResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	model, found := k.GetAIAgentModel(ctx, req.Id)
	if !found {
		return nil, status.Error(codes.NotFound, "AI agent model not found")
	}

	return &types.QueryAIAgentModelResponse{
		Model: model,
	}, nil
}

// AIAgentTrainingData returns training data for a specific AI agent
func (k Keeper) AIAgentTrainingData(c context.Context, req *types.QueryAIAgentTrainingDataRequest) (*types.QueryAIAgentTrainingDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	trainingData := k.GetAIAgentTrainingDataByAgent(ctx, req.AgentId)
	var pageRes *query.PageResponse
	var err error

	// Apply pagination
	if req.Pagination != nil {
		start, end := query.Paginate(len(trainingData), req.Pagination.Offset, req.Pagination.Limit, 100)
		if start < 0 || end < 0 {
			trainingData = []types.AIAgentTrainingData{}
		} else {
			trainingData = trainingData[start:end]
		}
		pageRes, err = query.NewPaginationResponse(uint64(len(trainingData)), req.Pagination.Offset, req.Pagination.Limit)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &types.QueryAIAgentTrainingDataResponse{
		TrainingData: trainingData,
		Pagination:   pageRes,
	}, nil
}

// MarketplaceListings returns all marketplace listings
func (k Keeper) MarketplaceListings(c context.Context, req *types.QueryMarketplaceListingsRequest) (*types.QueryMarketplaceListingsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	listings := k.GetAllAIAgentMarketplaceListings(ctx)
	var filteredListings []types.AIAgentMarketplaceListing

	// Filter by status if specified
	if req.Status != "" {
		for _, listing := range listings {
			if listing.Status == req.Status {
				filteredListings = append(filteredListings, listing)
			}
		}
		listings = filteredListings
		filteredListings = []types.AIAgentMarketplaceListing{}
	}

	// Filter by listing type if specified
	if req.ListingType != "" {
		for _, listing := range listings {
			if listing.ListingType == req.ListingType {
				filteredListings = append(filteredListings, listing)
			}
		}
		listings = filteredListings
	}

	var pageRes *query.PageResponse
	var err error

	// Apply pagination
	if req.Pagination != nil {
		start, end := query.Paginate(len(listings), req.Pagination.Offset, req.Pagination.Limit, 100)
		if start < 0 || end < 0 {
			listings = []types.AIAgentMarketplaceListing{}
		} else {
			listings = listings[start:end]
		}
		pageRes, err = query.NewPaginationResponse(uint64(len(listings)), req.Pagination.Offset, req.Pagination.Limit)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &types.QueryMarketplaceListingsResponse{
		Listings:   listings,
		Pagination: pageRes,
	}, nil
}

// MarketplaceListing returns a specific marketplace listing
func (k Keeper) MarketplaceListing(c context.Context, req *types.QueryMarketplaceListingRequest) (*types.QueryMarketplaceListingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	listing, found := k.GetAIAgentMarketplaceListing(ctx, req.Id)
	if !found {
		return nil, status.Error(codes.NotFound, "marketplace listing not found")
	}

	return &types.QueryMarketplaceListingResponse{
		Listing: listing,
	}, nil
}
