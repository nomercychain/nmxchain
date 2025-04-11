package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/nomercychain/nmxchain/x/deai/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// NewQuerier creates a new querier for deai module
func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryAIAgent:
			return queryAIAgent(ctx, path[1:], req, k, legacyQuerierCdc)
		case types.QueryAIAgents:
			return queryAIAgents(ctx, path[1:], req, k, legacyQuerierCdc)
		case types.QueryAIAgentState:
			return queryAIAgentState(ctx, path[1:], req, k, legacyQuerierCdc)
		case types.QueryAIAgentActions:
			return queryAIAgentActions(ctx, path[1:], req, k, legacyQuerierCdc)
		case types.QueryAIAgentAction:
			return queryAIAgentAction(ctx, path[1:], req, k, legacyQuerierCdc)
		case types.QueryAIAgentModels:
			return queryAIAgentModels(ctx, path[1:], req, k, legacyQuerierCdc)
		case types.QueryAIAgentModel:
			return queryAIAgentModel(ctx, path[1:], req, k, legacyQuerierCdc)
		case types.QueryAIAgentTrainingData:
			return queryAIAgentTrainingData(ctx, path[1:], req, k, legacyQuerierCdc)
		case types.QueryMarketplaceListings:
			return queryMarketplaceListings(ctx, path[1:], req, k, legacyQuerierCdc)
		case types.QueryMarketplaceListing:
			return queryMarketplaceListing(ctx, path[1:], req, k, legacyQuerierCdc)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown deai query endpoint")
		}
	}
}

func queryAIAgent(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryAIAgentRequest
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	agent, found := k.GetAIAgent(ctx, params.ID)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrAgentNotFound, params.ID)
	}

	res := types.QueryAIAgentResponse{
		Agent: agent,
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, res)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryAIAgents(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryAIAgentsRequest
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	var agents []types.AIAgent
	if params.Owner != "" {
		ownerAddr, err := sdk.AccAddressFromBech32(params.Owner)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, err.Error())
		}
		agents = k.GetAIAgentsByOwner(ctx, ownerAddr)
	} else {
		agents = k.GetAllAIAgents(ctx)
	}

	res := types.QueryAIAgentsResponse{
		Agents: agents,
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, res)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryAIAgentState(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryAIAgentStateRequest
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	state, found := k.GetAIAgentState(ctx, params.AgentID)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrAgentStateNotFound, params.AgentID)
	}

	res := types.QueryAIAgentStateResponse{
		State: state,
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, res)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryAIAgentActions(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryAIAgentActionsRequest
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	actions := k.GetAIAgentActionsByAgent(ctx, params.AgentID)

	res := types.QueryAIAgentActionsResponse{
		Actions: actions,
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, res)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryAIAgentAction(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryAIAgentActionRequest
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	action, found := k.GetAIAgentAction(ctx, params.ID)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrActionNotFound, params.ID)
	}

	res := types.QueryAIAgentActionResponse{
		Action: action,
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, res)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryAIAgentModels(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	models := k.GetAllAIAgentModels(ctx)

	res := types.QueryAIAgentModelsResponse{
		Models: models,
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, res)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryAIAgentModel(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryAIAgentModelRequest
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	model, found := k.GetAIAgentModel(ctx, params.ID)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrModelNotFound, params.ID)
	}

	res := types.QueryAIAgentModelResponse{
		Model: model,
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, res)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryAIAgentTrainingData(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryAIAgentTrainingDataRequest
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	trainingData := k.GetAIAgentTrainingDataByAgent(ctx, params.AgentID)

	res := types.QueryAIAgentTrainingDataResponse{
		TrainingData: trainingData,
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, res)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryMarketplaceListings(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryMarketplaceListingsRequest
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	listings := k.GetAllAIAgentMarketplaceListings(ctx)

	// Filter by status if specified
	if params.Status != "" {
		var filteredListings []types.AIAgentMarketplaceListing
		for _, listing := range listings {
			if listing.Status == params.Status {
				filteredListings = append(filteredListings, listing)
			}
		}
		listings = filteredListings
	}

	// Filter by listing type if specified
	if params.ListingType != "" {
		var filteredListings []types.AIAgentMarketplaceListing
		for _, listing := range listings {
			if listing.ListingType == params.ListingType {
				filteredListings = append(filteredListings, listing)
			}
		}
		listings = filteredListings
	}

	res := types.QueryMarketplaceListingsResponse{
		Listings: listings,
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, res)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryMarketplaceListing(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryMarketplaceListingRequest
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	listing, found := k.GetAIAgentMarketplaceListing(ctx, params.ID)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrListingNotFound, params.ID)
	}

	res := types.QueryMarketplaceListingResponse{
		Listing: listing,
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, res)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}