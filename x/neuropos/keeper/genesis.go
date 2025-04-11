package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nomercychain/nmxchain/x/neuropos/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	// Set all the validators
	for _, validator := range genState.Validators {
		valAddr, err := sdk.ValAddressFromBech32(validator.OperatorAddress)
		if err != nil {
			panic(err)
		}

		// Set validator in store
		// In a real implementation, this would interact with the staking module
		// For now, we'll just store the validator in our own store
		k.SetValidator(ctx, validator)

		// Initialize validator reputation
		k.InitializeValidatorReputation(ctx, validator.OperatorAddress)
	}

	// Set all the delegations
	for _, delegation := range genState.Delegations {
		// Set delegation in store
		// In a real implementation, this would interact with the staking module
		// For now, we'll just store the delegation in our own store
		k.SetDelegation(ctx, delegation)
	}

	// Set all the unbonding delegations
	for _, ubd := range genState.UnbondingDelegations {
		// Set unbonding delegation in store
		// In a real implementation, this would interact with the staking module
		// For now, we'll just store the unbonding delegation in our own store
		k.SetUnbondingDelegation(ctx, ubd)
	}

	// Set all the redelegations
	for _, red := range genState.Redelegations {
		// Set redelegation in store
		// In a real implementation, this would interact with the staking module
		// For now, we'll just store the redelegation in our own store
		k.SetRedelegation(ctx, red)
	}

	// Set all the neural networks
	for _, network := range genState.NeuralNetworks {
		k.SetNeuralNetwork(ctx, network)
	}

	// Set all the neural network weights
	for _, weights := range genState.NeuralNetworkWeights {
		k.SetNeuralNetworkWeights(ctx, weights)
	}

	// Set all the training data
	for _, data := range genState.TrainingData {
		k.SetTrainingData(ctx, data)
	}

	// Set all the neural predictions
	for _, prediction := range genState.NeuralPredictions {
		k.SetNeuralPrediction(ctx, prediction)
	}

	// Set all the validator performances
	for _, performance := range genState.ValidatorPerformances {
		k.SetValidatorPerformance(ctx, performance)
	}

	// Set all the validator reputations
	for _, reputation := range genState.ValidatorReputations {
		k.SetValidatorReputation(ctx, reputation)
	}

	// Set all the validator slash events
	for _, event := range genState.ValidatorSlashEvents {
		store := ctx.KVStore(k.storeKey)
		key := types.ValidatorSlashEventKey(event.ValidatorAddress, event.Height)
		value := k.cdc.MustMarshal(&event)
		store.Set(key, value)
	}

	// Set all the validator signing infos
	for _, info := range genState.ValidatorSigningInfos {
		k.SetValidatorSigningInfo(ctx, info)
	}

	// Set the params
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// Get all validators
	validators := k.GetAllValidators(ctx)
	genesis.Validators = validators

	// Get all delegations
	delegations := k.GetAllDelegations(ctx)
	genesis.Delegations = delegations

	// Get all unbonding delegations
	unbondingDelegations := k.GetAllUnbondingDelegations(ctx)
	genesis.UnbondingDelegations = unbondingDelegations

	// Get all redelegations
	redelegations := k.GetAllRedelegations(ctx)
	genesis.Redelegations = redelegations

	// Get all neural networks
	neuralNetworks := k.GetAllNeuralNetworks(ctx)
	genesis.NeuralNetworks = neuralNetworks

	// Get all neural network weights
	neuralNetworkWeights := k.GetAllNeuralNetworkWeights(ctx)
	genesis.NeuralNetworkWeights = neuralNetworkWeights

	// Get all training data
	trainingData := k.GetAllTrainingData(ctx)
	genesis.TrainingData = trainingData

	// Get all neural predictions
	neuralPredictions := k.GetAllNeuralPredictions(ctx)
	genesis.NeuralPredictions = neuralPredictions

	// Get all validator performances
	validatorPerformances := k.GetAllValidatorPerformances(ctx)
	genesis.ValidatorPerformances = validatorPerformances

	// Get all validator reputations
	validatorReputations := k.GetAllValidatorReputations(ctx)
	genesis.ValidatorReputations = validatorReputations

	// Get all validator slash events
	validatorSlashEvents := k.GetAllValidatorSlashEvents(ctx)
	genesis.ValidatorSlashEvents = validatorSlashEvents

	// Get all validator signing infos
	validatorSigningInfos := k.GetAllValidatorSigningInfos(ctx)
	genesis.ValidatorSigningInfos = validatorSigningInfos

	// Get params
	genesis.Params = k.GetParams(ctx)

	return genesis
}

// SetValidator sets a validator in the store
func (k Keeper) SetValidator(ctx sdk.Context, validator types.Validator) {
	store := ctx.KVStore(k.storeKey)
	key := types.ValidatorKey(validator.OperatorAddress)
	value := k.cdc.MustMarshal(&validator)
	store.Set(key, value)
}

// GetValidator gets a validator from the store
func (k Keeper) GetValidator(ctx sdk.Context, valAddr string) (types.Validator, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.ValidatorKey(valAddr)
	value := store.Get(key)
	if value == nil {
		return types.Validator{}, false
	}

	var validator types.Validator
	k.cdc.MustUnmarshal(value, &validator)
	return validator, true
}

// GetAllValidators gets all validators from the store
func (k Keeper) GetAllValidators(ctx sdk.Context) []types.Validator {
	var validators []types.Validator
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ValidatorKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var validator types.Validator
		k.cdc.MustUnmarshal(iterator.Value(), &validator)
		validators = append(validators, validator)
	}

	return validators
}

// SetDelegation sets a delegation in the store
func (k Keeper) SetDelegation(ctx sdk.Context, delegation types.Delegation) {
	store := ctx.KVStore(k.storeKey)
	key := types.DelegationKey(delegation.DelegatorAddress, delegation.ValidatorAddress)
	value := k.cdc.MustMarshal(&delegation)
	store.Set(key, value)
}

// GetDelegation gets a delegation from the store
func (k Keeper) GetDelegation(ctx sdk.Context, delAddr string, valAddr string) (types.Delegation, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.DelegationKey(delAddr, valAddr)
	value := store.Get(key)
	if value == nil {
		return types.Delegation{}, false
	}

	var delegation types.Delegation
	k.cdc.MustUnmarshal(value, &delegation)
	return delegation, true
}

// GetAllDelegations gets all delegations from the store
func (k Keeper) GetAllDelegations(ctx sdk.Context) []types.Delegation {
	var delegations []types.Delegation
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.DelegationKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var delegation types.Delegation
		k.cdc.MustUnmarshal(iterator.Value(), &delegation)
		delegations = append(delegations, delegation)
	}

	return delegations
}

// SetUnbondingDelegation sets an unbonding delegation in the store
func (k Keeper) SetUnbondingDelegation(ctx sdk.Context, ubd types.UnbondingDelegation) {
	store := ctx.KVStore(k.storeKey)
	key := types.UnbondingDelegationKey(ubd.DelegatorAddress, ubd.ValidatorAddress)
	value := k.cdc.MustMarshal(&ubd)
	store.Set(key, value)
}

// GetUnbondingDelegation gets an unbonding delegation from the store
func (k Keeper) GetUnbondingDelegation(ctx sdk.Context, delAddr string, valAddr string) (types.UnbondingDelegation, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.UnbondingDelegationKey(delAddr, valAddr)
	value := store.Get(key)
	if value == nil {
		return types.UnbondingDelegation{}, false
	}

	var ubd types.UnbondingDelegation
	k.cdc.MustUnmarshal(value, &ubd)
	return ubd, true
}

// GetAllUnbondingDelegations gets all unbonding delegations from the store
func (k Keeper) GetAllUnbondingDelegations(ctx sdk.Context) []types.UnbondingDelegation {
	var ubds []types.UnbondingDelegation
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.UnbondingDelegationKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var ubd types.UnbondingDelegation
		k.cdc.MustUnmarshal(iterator.Value(), &ubd)
		ubds = append(ubds, ubd)
	}

	return ubds
}

// SetRedelegation sets a redelegation in the store
func (k Keeper) SetRedelegation(ctx sdk.Context, red types.Redelegation) {
	store := ctx.KVStore(k.storeKey)
	key := types.RedelegationKey(red.DelegatorAddress, red.ValidatorSrcAddress, red.ValidatorDstAddress)
	value := k.cdc.MustMarshal(&red)
	store.Set(key, value)
}

// GetRedelegation gets a redelegation from the store
func (k Keeper) GetRedelegation(ctx sdk.Context, delAddr string, valSrcAddr string, valDstAddr string) (types.Redelegation, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.RedelegationKey(delAddr, valSrcAddr, valDstAddr)
	value := store.Get(key)
	if value == nil {
		return types.Redelegation{}, false
	}

	var red types.Redelegation
	k.cdc.MustUnmarshal(value, &red)
	return red, true
}

// GetAllRedelegations gets all redelegations from the store
func (k Keeper) GetAllRedelegations(ctx sdk.Context) []types.Redelegation {
	var reds []types.Redelegation
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.RedelegationKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var red types.Redelegation
		k.cdc.MustUnmarshal(iterator.Value(), &red)
		reds = append(reds, red)
	}

	return reds
}

// GetAllNeuralNetworkWeights gets all neural network weights from the store
func (k Keeper) GetAllNeuralNetworkWeights(ctx sdk.Context) []types.NeuralNetworkWeights {
	var weights []types.NeuralNetworkWeights
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.NeuralNetworkWeightKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var weight types.NeuralNetworkWeights
		k.cdc.MustUnmarshal(iterator.Value(), &weight)
		weights = append(weights, weight)
	}

	return weights
}

// GetAllTrainingData gets all training data from the store
func (k Keeper) GetAllTrainingData(ctx sdk.Context) []types.TrainingData {
	var dataList []types.TrainingData
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.NeuralNetworkTrainingDataKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var data types.TrainingData
		k.cdc.MustUnmarshal(iterator.Value(), &data)
		dataList = append(dataList, data)
	}

	return dataList
}

// GetAllNeuralPredictions gets all neural predictions from the store
func (k Keeper) GetAllNeuralPredictions(ctx sdk.Context) []types.NeuralPrediction {
	var predictions []types.NeuralPrediction
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.NeuralNetworkPredictionKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var prediction types.NeuralPrediction
		k.cdc.MustUnmarshal(iterator.Value(), &prediction)
		predictions = append(predictions, prediction)
	}

	return predictions
}

// GetAllValidatorSlashEvents gets all validator slash events from the store
func (k Keeper) GetAllValidatorSlashEvents(ctx sdk.Context) []types.ValidatorSlashEvent {
	var events []types.ValidatorSlashEvent
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ValidatorSlashEventKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var event types.ValidatorSlashEvent
		k.cdc.MustUnmarshal(iterator.Value(), &event)
		events = append(events, event)
	}

	return events
}

// GetAllValidatorSigningInfos gets all validator signing infos from the store
func (k Keeper) GetAllValidatorSigningInfos(ctx sdk.Context) []types.ValidatorSigningInfo {
	var infos []types.ValidatorSigningInfo
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ValidatorSigningInfoKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var info types.ValidatorSigningInfo
		k.cdc.MustUnmarshal(iterator.Value(), &info)
		infos = append(infos, info)
	}

	return infos
}