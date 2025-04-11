package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/nomercychain/nmxchain/x/neuropos/types"
)

// NewQuerier creates a new querier for neuropos module
func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryParams:
			return queryParams(ctx, k, legacyQuerierCdc)
		case types.QueryValidator:
			return queryValidator(ctx, req, k, legacyQuerierCdc)
		case types.QueryValidators:
			return queryValidators(ctx, k, legacyQuerierCdc)
		case types.QueryDelegation:
			return queryDelegation(ctx, req, k, legacyQuerierCdc)
		case types.QueryDelegatorDelegations:
			return queryDelegatorDelegations(ctx, req, k, legacyQuerierCdc)
		case types.QueryValidatorDelegations:
			return queryValidatorDelegations(ctx, req, k, legacyQuerierCdc)
		case types.QueryUnbondingDelegation:
			return queryUnbondingDelegation(ctx, req, k, legacyQuerierCdc)
		case types.QueryDelegatorUnbondingDelegations:
			return queryDelegatorUnbondingDelegations(ctx, req, k, legacyQuerierCdc)
		case types.QueryValidatorUnbondingDelegations:
			return queryValidatorUnbondingDelegations(ctx, req, k, legacyQuerierCdc)
		case types.QueryRedelegation:
			return queryRedelegation(ctx, req, k, legacyQuerierCdc)
		case types.QueryDelegatorRedelegations:
			return queryDelegatorRedelegations(ctx, req, k, legacyQuerierCdc)
		case types.QueryValidatorRedelegations:
			return queryValidatorRedelegations(ctx, req, k, legacyQuerierCdc)
		case types.QueryNeuralNetwork:
			return queryNeuralNetwork(ctx, req, k, legacyQuerierCdc)
		case types.QueryNeuralNetworks:
			return queryNeuralNetworks(ctx, k, legacyQuerierCdc)
		case types.QueryNeuralNetworkWeights:
			return queryNeuralNetworkWeights(ctx, req, k, legacyQuerierCdc)
		case types.QueryTrainingData:
			return queryTrainingData(ctx, req, k, legacyQuerierCdc)
		case types.QueryNetworkTrainingData:
			return queryNetworkTrainingData(ctx, req, k, legacyQuerierCdc)
		case types.QueryNeuralPrediction:
			return queryNeuralPrediction(ctx, req, k, legacyQuerierCdc)
		case types.QueryNetworkNeuralPredictions:
			return queryNetworkNeuralPredictions(ctx, req, k, legacyQuerierCdc)
		case types.QueryValidatorPerformance:
			return queryValidatorPerformance(ctx, req, k, legacyQuerierCdc)
		case types.QueryValidatorPerformances:
			return queryValidatorPerformances(ctx, k, legacyQuerierCdc)
		case types.QueryValidatorReputation:
			return queryValidatorReputation(ctx, req, k, legacyQuerierCdc)
		case types.QueryValidatorReputations:
			return queryValidatorReputations(ctx, k, legacyQuerierCdc)
		case types.QueryValidatorSlashEvents:
			return queryValidatorSlashEvents(ctx, req, k, legacyQuerierCdc)
		case types.QueryNetworkState:
			return queryNetworkState(ctx, k, legacyQuerierCdc)
		case types.QueryAnomalyReports:
			return queryAnomalyReports(ctx, k, legacyQuerierCdc)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown %s query endpoint: %s", types.ModuleName, path[0])
		}
	}
}

func queryParams(ctx sdk.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	params := k.GetParams(ctx)

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryValidator(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryValidatorParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	validator, found := k.GetValidator(ctx, params.ValidatorAddr)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrNoValidatorFound, params.ValidatorAddr)
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, validator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryValidators(ctx sdk.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	validators := k.GetAllValidators(ctx)

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, validators)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryDelegation(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryDelegationParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	delegation, found := k.GetDelegation(ctx, params.DelegatorAddr, params.ValidatorAddr)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrNoDelegation, "delegation not found")
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, delegation)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryDelegatorDelegations(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryDelegatorParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	var delegations []types.Delegation
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.DelegationKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var delegation types.Delegation
		k.cdc.MustUnmarshal(iterator.Value(), &delegation)
		if delegation.DelegatorAddress == params.DelegatorAddr {
			delegations = append(delegations, delegation)
		}
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, delegations)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryValidatorDelegations(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryValidatorParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	var delegations []types.Delegation
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.DelegationKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var delegation types.Delegation
		k.cdc.MustUnmarshal(iterator.Value(), &delegation)
		if delegation.ValidatorAddress == params.ValidatorAddr {
			delegations = append(delegations, delegation)
		}
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, delegations)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryUnbondingDelegation(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryDelegationParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	ubd, found := k.GetUnbondingDelegation(ctx, params.DelegatorAddr, params.ValidatorAddr)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrNoUnbondingDelegation, "unbonding delegation not found")
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, ubd)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryDelegatorUnbondingDelegations(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryDelegatorParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	var ubds []types.UnbondingDelegation
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.UnbondingDelegationKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var ubd types.UnbondingDelegation
		k.cdc.MustUnmarshal(iterator.Value(), &ubd)
		if ubd.DelegatorAddress == params.DelegatorAddr {
			ubds = append(ubds, ubd)
		}
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, ubds)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryValidatorUnbondingDelegations(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryValidatorParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	var ubds []types.UnbondingDelegation
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.UnbondingDelegationKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var ubd types.UnbondingDelegation
		k.cdc.MustUnmarshal(iterator.Value(), &ubd)
		if ubd.ValidatorAddress == params.ValidatorAddr {
			ubds = append(ubds, ubd)
		}
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, ubds)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryRedelegation(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryRedelegationParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	red, found := k.GetRedelegation(ctx, params.DelegatorAddr, params.SrcValidatorAddr, params.DstValidatorAddr)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrNoRedelegation, "redelegation not found")
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, red)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryDelegatorRedelegations(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryDelegatorParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	var reds []types.Redelegation
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.RedelegationKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var red types.Redelegation
		k.cdc.MustUnmarshal(iterator.Value(), &red)
		if red.DelegatorAddress == params.DelegatorAddr {
			reds = append(reds, red)
		}
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, reds)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryValidatorRedelegations(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryValidatorParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	var reds []types.Redelegation
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.RedelegationKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var red types.Redelegation
		k.cdc.MustUnmarshal(iterator.Value(), &red)
		if red.ValidatorSrcAddress == params.ValidatorAddr || red.ValidatorDstAddress == params.ValidatorAddr {
			reds = append(reds, red)
		}
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, reds)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryNeuralNetwork(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryNeuralNetworkParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	network, found := k.GetNeuralNetwork(ctx, params.NetworkID)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrNoNeuralNetworkFound, params.NetworkID)
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, network)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryNeuralNetworks(ctx sdk.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	networks := k.GetAllNeuralNetworks(ctx)

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, networks)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryNeuralNetworkWeights(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryNeuralNetworkWeightsParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	var weights types.NeuralNetworkWeights
	var found bool

	if params.Version == 0 {
		// Get latest version
		weights, found = k.GetLatestNeuralNetworkWeights(ctx, params.NetworkID)
	} else {
		// Get specific version
		weights, found = k.GetNeuralNetworkWeights(ctx, params.NetworkID, params.Version)
	}

	if !found {
		return nil, sdkerrors.Wrapf(types.ErrNoNeuralNetworkWeightsFound, "network: %s, version: %d", params.NetworkID, params.Version)
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, weights)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryTrainingData(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryTrainingDataParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	data, found := k.GetTrainingData(ctx, params.DataID)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrNoTrainingDataFound, params.DataID)
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, data)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryNetworkTrainingData(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryNeuralNetworkParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	data := k.GetTrainingDataByNetwork(ctx, params.NetworkID)

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, data)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryNeuralPrediction(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryNeuralPredictionParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	prediction, found := k.GetNeuralPrediction(ctx, params.PredictionID)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrNoNeuralPredictionFound, params.PredictionID)
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, prediction)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryNetworkNeuralPredictions(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryNeuralNetworkParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	predictions := k.GetNeuralPredictionsByNetwork(ctx, params.NetworkID)

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, predictions)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryValidatorPerformance(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryValidatorParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	performance, found := k.GetValidatorPerformance(ctx, params.ValidatorAddr)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrNoValidatorFound, params.ValidatorAddr)
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, performance)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryValidatorPerformances(ctx sdk.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	performances := k.GetAllValidatorPerformances(ctx)

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, performances)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryValidatorReputation(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryValidatorParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	reputation, found := k.GetValidatorReputation(ctx, params.ValidatorAddr)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrNoValidatorFound, params.ValidatorAddr)
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, reputation)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryValidatorReputations(ctx sdk.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	reputations := k.GetAllValidatorReputations(ctx)

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, reputations)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryValidatorSlashEvents(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryValidatorParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	events := k.GetValidatorSlashEvents(ctx, params.ValidatorAddr)

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, events)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryNetworkState(ctx sdk.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	state, found := k.GetNetworkState(ctx)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrNoNetworkState, "network state not found")
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, state)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryAnomalyReports(ctx sdk.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	reports := k.GetAllAnomalyReports(ctx)

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, reports)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}