package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/nomercychain/nmxchain/x/neuropos/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

// Params returns the module parameters
func (k queryServer) Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.GetParams(ctx)

	return &types.QueryParamsResponse{Params: params}, nil
}

// Validator returns validator info by address
func (k queryServer) Validator(goCtx context.Context, req *types.QueryValidatorRequest) (*types.QueryValidatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.ValidatorAddr == "" {
		return nil, status.Error(codes.InvalidArgument, "validator address cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetValidator(ctx, req.ValidatorAddr)
	if !found {
		return nil, status.Errorf(codes.NotFound, "validator %s not found", req.ValidatorAddr)
	}

	return &types.QueryValidatorResponse{Validator: val}, nil
}

// Validators returns all validators
func (k queryServer) Validators(goCtx context.Context, req *types.QueryValidatorsRequest) (*types.QueryValidatorsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	validators := k.GetAllValidators(ctx)

	return &types.QueryValidatorsResponse{Validators: validators}, nil
}

// Delegation returns delegation info for a validator
func (k queryServer) Delegation(goCtx context.Context, req *types.QueryDelegationRequest) (*types.QueryDelegationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.DelegatorAddr == "" {
		return nil, status.Error(codes.InvalidArgument, "delegator address cannot be empty")
	}

	if req.ValidatorAddr == "" {
		return nil, status.Error(codes.InvalidArgument, "validator address cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	delegation, found := k.GetDelegation(ctx, req.DelegatorAddr, req.ValidatorAddr)
	if !found {
		return nil, status.Errorf(codes.NotFound, "delegation not found for delegator %s and validator %s", req.DelegatorAddr, req.ValidatorAddr)
	}

	return &types.QueryDelegationResponse{Delegation: delegation}, nil
}

// DelegatorDelegations returns all delegations for a delegator
func (k queryServer) DelegatorDelegations(goCtx context.Context, req *types.QueryDelegatorDelegationsRequest) (*types.QueryDelegatorDelegationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.DelegatorAddr == "" {
		return nil, status.Error(codes.InvalidArgument, "delegator address cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var delegations []types.Delegation
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.DelegationKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var delegation types.Delegation
		k.cdc.MustUnmarshal(iterator.Value(), &delegation)
		if delegation.DelegatorAddress == req.DelegatorAddr {
			delegations = append(delegations, delegation)
		}
	}

	return &types.QueryDelegatorDelegationsResponse{Delegations: delegations}, nil
}

// ValidatorDelegations returns all delegations to a validator
func (k queryServer) ValidatorDelegations(goCtx context.Context, req *types.QueryValidatorDelegationsRequest) (*types.QueryValidatorDelegationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.ValidatorAddr == "" {
		return nil, status.Error(codes.InvalidArgument, "validator address cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var delegations []types.Delegation
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.DelegationKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var delegation types.Delegation
		k.cdc.MustUnmarshal(iterator.Value(), &delegation)
		if delegation.ValidatorAddress == req.ValidatorAddr {
			delegations = append(delegations, delegation)
		}
	}

	return &types.QueryValidatorDelegationsResponse{Delegations: delegations}, nil
}

// UnbondingDelegation returns unbonding delegation info for a delegator and validator
func (k queryServer) UnbondingDelegation(goCtx context.Context, req *types.QueryUnbondingDelegationRequest) (*types.QueryUnbondingDelegationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.DelegatorAddr == "" {
		return nil, status.Error(codes.InvalidArgument, "delegator address cannot be empty")
	}

	if req.ValidatorAddr == "" {
		return nil, status.Error(codes.InvalidArgument, "validator address cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	ubd, found := k.GetUnbondingDelegation(ctx, req.DelegatorAddr, req.ValidatorAddr)
	if !found {
		return nil, status.Errorf(codes.NotFound, "unbonding delegation not found for delegator %s and validator %s", req.DelegatorAddr, req.ValidatorAddr)
	}

	return &types.QueryUnbondingDelegationResponse{Unbond: ubd}, nil
}

// DelegatorUnbondingDelegations returns all unbonding delegations for a delegator
func (k queryServer) DelegatorUnbondingDelegations(goCtx context.Context, req *types.QueryDelegatorUnbondingDelegationsRequest) (*types.QueryDelegatorUnbondingDelegationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.DelegatorAddr == "" {
		return nil, status.Error(codes.InvalidArgument, "delegator address cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var ubds []types.UnbondingDelegation
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.UnbondingDelegationKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var ubd types.UnbondingDelegation
		k.cdc.MustUnmarshal(iterator.Value(), &ubd)
		if ubd.DelegatorAddress == req.DelegatorAddr {
			ubds = append(ubds, ubd)
		}
	}

	return &types.QueryDelegatorUnbondingDelegationsResponse{Unbonds: ubds}, nil
}

// ValidatorUnbondingDelegations returns all unbonding delegations from a validator
func (k queryServer) ValidatorUnbondingDelegations(goCtx context.Context, req *types.QueryValidatorUnbondingDelegationsRequest) (*types.QueryValidatorUnbondingDelegationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.ValidatorAddr == "" {
		return nil, status.Error(codes.InvalidArgument, "validator address cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var ubds []types.UnbondingDelegation
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.UnbondingDelegationKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var ubd types.UnbondingDelegation
		k.cdc.MustUnmarshal(iterator.Value(), &ubd)
		if ubd.ValidatorAddress == req.ValidatorAddr {
			ubds = append(ubds, ubd)
		}
	}

	return &types.QueryValidatorUnbondingDelegationsResponse{Unbonds: ubds}, nil
}

// Redelegation returns redelegation info for a delegator, source validator, and destination validator
func (k queryServer) Redelegation(goCtx context.Context, req *types.QueryRedelegationRequest) (*types.QueryRedelegationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.DelegatorAddr == "" {
		return nil, status.Error(codes.InvalidArgument, "delegator address cannot be empty")
	}

	if req.SrcValidatorAddr == "" {
		return nil, status.Error(codes.InvalidArgument, "source validator address cannot be empty")
	}

	if req.DstValidatorAddr == "" {
		return nil, status.Error(codes.InvalidArgument, "destination validator address cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	red, found := k.GetRedelegation(ctx, req.DelegatorAddr, req.SrcValidatorAddr, req.DstValidatorAddr)
	if !found {
		return nil, status.Errorf(codes.NotFound, "redelegation not found for delegator %s, source validator %s, and destination validator %s", req.DelegatorAddr, req.SrcValidatorAddr, req.DstValidatorAddr)
	}

	return &types.QueryRedelegationResponse{Redelegation: red}, nil
}

// DelegatorRedelegations returns all redelegations for a delegator
func (k queryServer) DelegatorRedelegations(goCtx context.Context, req *types.QueryDelegatorRedelegationsRequest) (*types.QueryDelegatorRedelegationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.DelegatorAddr == "" {
		return nil, status.Error(codes.InvalidArgument, "delegator address cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var reds []types.Redelegation
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.RedelegationKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var red types.Redelegation
		k.cdc.MustUnmarshal(iterator.Value(), &red)
		if red.DelegatorAddress == req.DelegatorAddr {
			reds = append(reds, red)
		}
	}

	return &types.QueryDelegatorRedelegationsResponse{Redelegations: reds}, nil
}

// ValidatorRedelegations returns all redelegations from a source validator
func (k queryServer) ValidatorRedelegations(goCtx context.Context, req *types.QueryValidatorRedelegationsRequest) (*types.QueryValidatorRedelegationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.ValidatorAddr == "" {
		return nil, status.Error(codes.InvalidArgument, "validator address cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var reds []types.Redelegation
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.RedelegationKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var red types.Redelegation
		k.cdc.MustUnmarshal(iterator.Value(), &red)
		if red.ValidatorSrcAddress == req.ValidatorAddr || red.ValidatorDstAddress == req.ValidatorAddr {
			reds = append(reds, red)
		}
	}

	return &types.QueryValidatorRedelegationsResponse{Redelegations: reds}, nil
}

// NeuralNetwork returns neural network info by ID
func (k queryServer) NeuralNetwork(goCtx context.Context, req *types.QueryNeuralNetworkRequest) (*types.QueryNeuralNetworkResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.NetworkId == "" {
		return nil, status.Error(codes.InvalidArgument, "network ID cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	network, found := k.GetNeuralNetwork(ctx, req.NetworkId)
	if !found {
		return nil, status.Errorf(codes.NotFound, "neural network %s not found", req.NetworkId)
	}

	return &types.QueryNeuralNetworkResponse{Network: network}, nil
}

// NeuralNetworks returns all neural networks
func (k queryServer) NeuralNetworks(goCtx context.Context, req *types.QueryNeuralNetworksRequest) (*types.QueryNeuralNetworksResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	networks := k.GetAllNeuralNetworks(ctx)

	return &types.QueryNeuralNetworksResponse{Networks: networks}, nil
}

// NeuralNetworkWeights returns neural network weights by ID and version
func (k queryServer) NeuralNetworkWeights(goCtx context.Context, req *types.QueryNeuralNetworkWeightsRequest) (*types.QueryNeuralNetworkWeightsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.NetworkId == "" {
		return nil, status.Error(codes.InvalidArgument, "network ID cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var weights types.NeuralNetworkWeights
	var found bool

	if req.Version == 0 {
		// Get latest version
		weights, found = k.GetLatestNeuralNetworkWeights(ctx, req.NetworkId)
	} else {
		// Get specific version
		weights, found = k.GetNeuralNetworkWeights(ctx, req.NetworkId, req.Version)
	}

	if !found {
		return nil, status.Errorf(codes.NotFound, "neural network weights for network %s and version %d not found", req.NetworkId, req.Version)
	}

	return &types.QueryNeuralNetworkWeightsResponse{Weights: weights}, nil
}

// TrainingData returns training data by ID
func (k queryServer) TrainingData(goCtx context.Context, req *types.QueryTrainingDataRequest) (*types.QueryTrainingDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.DataId == "" {
		return nil, status.Error(codes.InvalidArgument, "data ID cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	data, found := k.GetTrainingData(ctx, req.DataId)
	if !found {
		return nil, status.Errorf(codes.NotFound, "training data %s not found", req.DataId)
	}

	return &types.QueryTrainingDataResponse{Data: data}, nil
}

// NetworkTrainingData returns all training data for a neural network
func (k queryServer) NetworkTrainingData(goCtx context.Context, req *types.QueryNetworkTrainingDataRequest) (*types.QueryNetworkTrainingDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.NetworkId == "" {
		return nil, status.Error(codes.InvalidArgument, "network ID cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	data := k.GetTrainingDataByNetwork(ctx, req.NetworkId)

	return &types.QueryNetworkTrainingDataResponse{Data: data}, nil
}

// NeuralPrediction returns neural prediction by ID
func (k queryServer) NeuralPrediction(goCtx context.Context, req *types.QueryNeuralPredictionRequest) (*types.QueryNeuralPredictionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.PredictionId == "" {
		return nil, status.Error(codes.InvalidArgument, "prediction ID cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	prediction, found := k.GetNeuralPrediction(ctx, req.PredictionId)
	if !found {
		return nil, status.Errorf(codes.NotFound, "neural prediction %s not found", req.PredictionId)
	}

	return &types.QueryNeuralPredictionResponse{Prediction: prediction}, nil
}

// NetworkNeuralPredictions returns all neural predictions for a neural network
func (k queryServer) NetworkNeuralPredictions(goCtx context.Context, req *types.QueryNetworkNeuralPredictionsRequest) (*types.QueryNetworkNeuralPredictionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.NetworkId == "" {
		return nil, status.Error(codes.InvalidArgument, "network ID cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	predictions := k.GetNeuralPredictionsByNetwork(ctx, req.NetworkId)

	return &types.QueryNetworkNeuralPredictionsResponse{Predictions: predictions}, nil
}

// ValidatorPerformance returns validator performance by address
func (k queryServer) ValidatorPerformance(goCtx context.Context, req *types.QueryValidatorPerformanceRequest) (*types.QueryValidatorPerformanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.ValidatorAddr == "" {
		return nil, status.Error(codes.InvalidArgument, "validator address cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	performance, found := k.GetValidatorPerformance(ctx, req.ValidatorAddr)
	if !found {
		return nil, status.Errorf(codes.NotFound, "validator performance for %s not found", req.ValidatorAddr)
	}

	return &types.QueryValidatorPerformanceResponse{Performance: performance}, nil
}

// ValidatorPerformances returns all validator performances
func (k queryServer) ValidatorPerformances(goCtx context.Context, req *types.QueryValidatorPerformancesRequest) (*types.QueryValidatorPerformancesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	performances := k.GetAllValidatorPerformances(ctx)

	return &types.QueryValidatorPerformancesResponse{Performances: performances}, nil
}

// ValidatorReputation returns validator reputation by address
func (k queryServer) ValidatorReputation(goCtx context.Context, req *types.QueryValidatorReputationRequest) (*types.QueryValidatorReputationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.ValidatorAddr == "" {
		return nil, status.Error(codes.InvalidArgument, "validator address cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	reputation, found := k.GetValidatorReputation(ctx, req.ValidatorAddr)
	if !found {
		return nil, status.Errorf(codes.NotFound, "validator reputation for %s not found", req.ValidatorAddr)
	}

	return &types.QueryValidatorReputationResponse{Reputation: reputation}, nil
}

// ValidatorReputations returns all validator reputations
func (k queryServer) ValidatorReputations(goCtx context.Context, req *types.QueryValidatorReputationsRequest) (*types.QueryValidatorReputationsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	reputations := k.GetAllValidatorReputations(ctx)

	return &types.QueryValidatorReputationsResponse{Reputations: reputations}, nil
}

// ValidatorSlashEvents returns all slash events for a validator
func (k queryServer) ValidatorSlashEvents(goCtx context.Context, req *types.QueryValidatorSlashEventsRequest) (*types.QueryValidatorSlashEventsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.ValidatorAddr == "" {
		return nil, status.Error(codes.InvalidArgument, "validator address cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	events := k.GetValidatorSlashEvents(ctx, req.ValidatorAddr)

	return &types.QueryValidatorSlashEventsResponse{SlashEvents: events}, nil
}

// NetworkState returns the current network state
func (k queryServer) NetworkState(goCtx context.Context, req *types.QueryNetworkStateRequest) (*types.QueryNetworkStateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	state, found := k.GetNetworkState(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "network state not found")
	}

	return &types.QueryNetworkStateResponse{State: state}, nil
}

// AnomalyReports returns all anomaly reports
func (k queryServer) AnomalyReports(goCtx context.Context, req *types.QueryAnomalyReportsRequest) (*types.QueryAnomalyReportsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	reports := k.GetAllAnomalyReports(ctx)

	return &types.QueryAnomalyReportsResponse{Reports: reports}, nil
}