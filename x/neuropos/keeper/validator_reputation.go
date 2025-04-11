package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nomercychain/nmxchain/x/neuropos/types"
)

// SetValidatorPerformance sets a validator's performance metrics
func (k Keeper) SetValidatorPerformance(ctx sdk.Context, performance types.ValidatorPerformance) {
	store := ctx.KVStore(k.storeKey)
	key := types.ValidatorPerformanceKey(performance.ValidatorAddress)
	value := k.cdc.MustMarshal(&performance)
	store.Set(key, value)
}

// GetValidatorPerformance returns a validator's performance metrics
func (k Keeper) GetValidatorPerformance(ctx sdk.Context, validatorAddr string) (types.ValidatorPerformance, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.ValidatorPerformanceKey(validatorAddr)
	value := store.Get(key)
	if value == nil {
		return types.ValidatorPerformance{}, false
	}

	var performance types.ValidatorPerformance
	k.cdc.MustUnmarshal(value, &performance)
	return performance, true
}

// GetAllValidatorPerformances returns all validator performance metrics
func (k Keeper) GetAllValidatorPerformances(ctx sdk.Context) []types.ValidatorPerformance {
	var performances []types.ValidatorPerformance
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ValidatorPerformanceKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var performance types.ValidatorPerformance
		k.cdc.MustUnmarshal(iterator.Value(), &performance)
		performances = append(performances, performance)
	}

	return performances
}

// SetValidatorReputation sets a validator's reputation
func (k Keeper) SetValidatorReputation(ctx sdk.Context, reputation types.ValidatorReputation) {
	store := ctx.KVStore(k.storeKey)
	key := types.ValidatorReputationKey(reputation.ValidatorAddress)
	value := k.cdc.MustMarshal(&reputation)
	store.Set(key, value)
}

// GetValidatorReputation returns a validator's reputation
func (k Keeper) GetValidatorReputation(ctx sdk.Context, validatorAddr string) (types.ValidatorReputation, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.ValidatorReputationKey(validatorAddr)
	value := store.Get(key)
	if value == nil {
		return types.ValidatorReputation{}, false
	}

	var reputation types.ValidatorReputation
	k.cdc.MustUnmarshal(value, &reputation)
	return reputation, true
}

// GetAllValidatorReputations returns all validator reputations
func (k Keeper) GetAllValidatorReputations(ctx sdk.Context) []types.ValidatorReputation {
	var reputations []types.ValidatorReputation
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ValidatorReputationKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var reputation types.ValidatorReputation
		k.cdc.MustUnmarshal(iterator.Value(), &reputation)
		reputations = append(reputations, reputation)
	}

	return reputations
}

// AddValidatorSlashEvent adds a slash event for a validator
func (k Keeper) AddValidatorSlashEvent(ctx sdk.Context, validatorAddr string, height int64, reason string, slashFactor sdk.Dec, tokens sdk.Int) {
	store := ctx.KVStore(k.storeKey)
	key := types.ValidatorSlashEventKey(validatorAddr, height)

	slashEvent := types.ValidatorSlashEvent{
		ValidatorAddress: validatorAddr,
		Height:          height,
		Timestamp:       ctx.BlockTime(),
		Reason:          reason,
		SlashFactor:     slashFactor,
		Tokens:          tokens,
	}

	value := k.cdc.MustMarshal(&slashEvent)
	store.Set(key, value)

	// Update validator reputation
	k.UpdateValidatorReputationAfterSlash(ctx, validatorAddr, slashFactor, reason)
}

// GetValidatorSlashEvents returns all slash events for a validator
func (k Keeper) GetValidatorSlashEvents(ctx sdk.Context, validatorAddr string) []types.ValidatorSlashEvent {
	var slashEvents []types.ValidatorSlashEvent
	store := ctx.KVStore(k.storeKey)
	prefix := append(types.ValidatorSlashEventKeyPrefix, []byte(validatorAddr+"/")...)
	iterator := sdk.KVStorePrefixIterator(store, prefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var slashEvent types.ValidatorSlashEvent
		k.cdc.MustUnmarshal(iterator.Value(), &slashEvent)
		slashEvents = append(slashEvents, slashEvent)
	}

	return slashEvents
}

// SetValidatorSigningInfo sets a validator's signing info
func (k Keeper) SetValidatorSigningInfo(ctx sdk.Context, signingInfo types.ValidatorSigningInfo) {
	store := ctx.KVStore(k.storeKey)
	key := types.ValidatorSigningInfoKey(signingInfo.ValidatorAddress)
	value := k.cdc.MustMarshal(&signingInfo)
	store.Set(key, value)
}

// GetValidatorSigningInfo returns a validator's signing info
func (k Keeper) GetValidatorSigningInfo(ctx sdk.Context, validatorAddr string) (types.ValidatorSigningInfo, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.ValidatorSigningInfoKey(validatorAddr)
	value := store.Get(key)
	if value == nil {
		return types.ValidatorSigningInfo{}, false
	}

	var signingInfo types.ValidatorSigningInfo
	k.cdc.MustUnmarshal(value, &signingInfo)
	return signingInfo, true
}

// InitializeValidatorReputation initializes a validator's reputation
func (k Keeper) InitializeValidatorReputation(ctx sdk.Context, validatorAddr string) {
	// Check if reputation already exists
	_, found := k.GetValidatorReputation(ctx, validatorAddr)
	if found {
		return
	}

	// Initialize with default reputation (1.0)
	reputation := types.ValidatorReputation{
		ValidatorAddress: validatorAddr,
		Reputation:      sdk.OneDec(),
		LastUpdated:     ctx.BlockTime(),
		History:         []types.ReputationChange{},
	}

	k.SetValidatorReputation(ctx, reputation)

	// Initialize performance metrics
	performance := types.ValidatorPerformance{
		ValidatorAddress:   validatorAddr,
		BlocksProposed:     0,
		BlocksValidated:    0,
		MissedBlocks:       0,
		PredictionAccuracy: sdk.OneDec(),
		LastUpdated:        ctx.BlockTime(),
		PerformanceScore:   sdk.OneDec(),
		AssessmentWindow:   k.GetParams(ctx).PerformanceAssessmentWindow,
	}

	k.SetValidatorPerformance(ctx, performance)

	// Initialize signing info
	signingInfo := types.ValidatorSigningInfo{
		ValidatorAddress:    validatorAddr,
		StartHeight:         ctx.BlockHeight(),
		IndexOffset:         0,
		JailedUntil:         time.Time{},
		Tombstoned:          false,
		MissedBlocksCounter: 0,
		SignedBlocksWindow:  k.GetParams(ctx).SignedBlocksWindow,
		MinSignedPerWindow: k.GetParams(ctx).MinSignedPerWindow,
	}

	k.SetValidatorSigningInfo(ctx, signingInfo)
}

// UpdateValidatorReputation updates a validator's reputation based on performance
func (k Keeper) UpdateValidatorReputation(ctx sdk.Context, validatorAddr string, change sdk.Dec, reason string) error {
	// Get current reputation
	reputation, found := k.GetValidatorReputation(ctx, validatorAddr)
	if !found {
		// Initialize reputation if not found
		k.InitializeValidatorReputation(ctx, validatorAddr)
		reputation, _ = k.GetValidatorReputation(ctx, validatorAddr)
	}

	// Apply decay to current reputation
	params := k.GetParams(ctx)
	timeSinceLastUpdate := ctx.BlockTime().Sub(reputation.LastUpdated)
	decayFactor := sdk.OneDec().Sub(params.ReputationDecayRate.MulInt64(int64(timeSinceLastUpdate.Hours() / 24))) // Daily decay
	if decayFactor.LT(sdk.ZeroDec()) {
		decayFactor = sdk.ZeroDec()
	}

	decayedReputation := reputation.Reputation.Mul(decayFactor)

	// Apply the new change
	newReputation := decayedReputation.Add(change)

	// Ensure reputation is between 0 and 1
	if newReputation.LT(sdk.ZeroDec()) {
		newReputation = sdk.ZeroDec()
	} else if newReputation.GT(sdk.OneDec()) {
		newReputation = sdk.OneDec()
	}

	// Record the change in history
	reputationChange := types.ReputationChange{
		Timestamp: ctx.BlockTime(),
		Change:    change,
		Reason:    reason,
	}

	// Update the reputation
	reputation.Reputation = newReputation
	reputation.LastUpdated = ctx.BlockTime()
	reputation.History = append(reputation.History, reputationChange)

	// Limit history size to prevent unbounded growth
	if len(reputation.History) > 100 {
		reputation.History = reputation.History[len(reputation.History)-100:]
	}

	k.SetValidatorReputation(ctx, reputation)

	// Check if reputation is below minimum threshold
	if newReputation.LT(params.MinValidatorReputation) {
		// In a real implementation, this might trigger jailing or other penalties
		// For now, we'll just emit an event
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeUpdateValidatorReputation,
				sdk.NewAttribute(types.AttributeKeyValidator, validatorAddr),
				sdk.NewAttribute(types.AttributeKeyValidatorReputation, newReputation.String()),
				sdk.NewAttribute(types.AttributeKeyReputationChange, change.String()),
				sdk.NewAttribute("below_threshold", "true"),
			),
		)
	}

	return nil
}

// UpdateValidatorReputationAfterSlash updates a validator's reputation after a slash event
func (k Keeper) UpdateValidatorReputationAfterSlash(ctx sdk.Context, validatorAddr string, slashFactor sdk.Dec, reason string) {
	// Calculate reputation penalty based on slash factor
	params := k.GetParams(ctx)
	penaltyRate := params.ReputationPenaltyRate
	penalty := slashFactor.Mul(penaltyRate).Neg() // Negative change to reduce reputation

	// Update reputation
	k.UpdateValidatorReputation(ctx, validatorAddr, penalty, "slash: "+reason)
}

// UpdateValidatorPerformance updates a validator's performance metrics
func (k Keeper) UpdateValidatorPerformance(ctx sdk.Context, validatorAddr string, proposedBlock bool, validatedBlock bool, missedBlock bool) {
	// Get current performance
	performance, found := k.GetValidatorPerformance(ctx, validatorAddr)
	if !found {
		// Initialize performance if not found
		k.InitializeValidatorReputation(ctx, validatorAddr)
		performance, _ = k.GetValidatorPerformance(ctx, validatorAddr)
	}

	// Update metrics
	if proposedBlock {
		performance.BlocksProposed++
	}

	if validatedBlock {
		performance.BlocksValidated++
	}

	if missedBlock {
		performance.MissedBlocks++
	}

	// Calculate performance score
	totalBlocks := performance.BlocksProposed + performance.BlocksValidated
	if totalBlocks > 0 {
		missRate := sdk.NewDecFromInt(sdk.NewIntFromUint64(performance.MissedBlocks)).Quo(sdk.NewDecFromInt(sdk.NewIntFromUint64(totalBlocks)))
		performance.PerformanceScore = sdk.OneDec().Sub(missRate)
	}

	// Update last updated time
	performance.LastUpdated = ctx.BlockTime()

	// Save performance
	k.SetValidatorPerformance(ctx, performance)

	// Update reputation based on performance
	params := k.GetParams(ctx)
	if performance.AssessmentWindow > 0 && totalBlocks >= performance.AssessmentWindow {
		// Calculate reputation change based on performance score
		baseChange := performance.PerformanceScore.Sub(sdk.NewDecWithPrec(5, 1)) // 0.5 is neutral
		reputationChange := baseChange.Mul(params.ReputationBonusRate)

		// Apply neural network influence if available
		nnInfluence := k.CalculateNeuralNetworkInfluence(ctx, validatorAddr)
		if !nnInfluence.IsZero() {
			reputationChange = reputationChange.Add(nnInfluence.Mul(params.NeuralNetworkInfluenceRate))
		}

		// Update reputation
		k.UpdateValidatorReputation(ctx, validatorAddr, reputationChange, "performance assessment")

		// Reset assessment window
		performance.BlocksProposed = 0
		performance.BlocksValidated = 0
		performance.MissedBlocks = 0
		performance.LastUpdated = ctx.BlockTime()
		k.SetValidatorPerformance(ctx, performance)
	}

	// Check for excessive missed blocks
	if performance.MissedBlocks > params.MaxMissedBlocks {
		// In a real implementation, this might trigger jailing
		// For now, we'll just emit an event
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeValidatorPerformance,
				sdk.NewAttribute(types.AttributeKeyValidator, validatorAddr),
				sdk.NewAttribute(types.AttributeKeyMissedBlocks, sdk.NewIntFromUint64(performance.MissedBlocks).String()),
				sdk.NewAttribute("excessive_missed_blocks", "true"),
			),
		)
	}
}

// CalculateNeuralNetworkInfluence calculates the influence of neural network contributions on reputation
func (k Keeper) CalculateNeuralNetworkInfluence(ctx sdk.Context, validatorAddr string) sdk.Dec {
	// This is a placeholder for the actual neural network influence calculation
	// In a real implementation, this would analyze the validator's neural network contributions
	// For now, we'll return a small random value

	// Get all predictions by this validator
	var validatorPredictions []types.NeuralPrediction
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.NeuralNetworkPredictionKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var prediction types.NeuralPrediction
		k.cdc.MustUnmarshal(iterator.Value(), &prediction)

		// Check if this validator is in the validator set
		for _, valAddr := range prediction.ValidatorSet {
			if valAddr == validatorAddr {
				validatorPredictions = append(validatorPredictions, prediction)
				break
			}
		}
	}

	// If no predictions, return zero influence
	if len(validatorPredictions) == 0 {
		return sdk.ZeroDec()
	}

	// Calculate average confidence of predictions
	totalConfidence := sdk.ZeroDec()
	for _, pred := range validatorPredictions {
		totalConfidence = totalConfidence.Add(pred.Confidence)
	}

	avgConfidence := totalConfidence.Quo(sdk.NewDecFromInt(sdk.NewIntFromUint64(uint64(len(validatorPredictions)))))

	// Scale to a small influence value (-0.1 to 0.1)
	influence := avgConfidence.Sub(sdk.NewDecWithPrec(5, 1)).Mul(sdk.NewDecWithPrec(2, 1))

	return influence
}