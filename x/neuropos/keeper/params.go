package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nomercychain/nmxchain/x/neuropos/types"
)

// GetParams gets all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.Params{
		UnbondingTime:               k.UnbondingTime(ctx),
		MaxValidators:               k.MaxValidators(ctx),
		MinSelfDelegation:           k.MinSelfDelegation(ctx),
		HistoricalEntries:           k.HistoricalEntries(ctx),
		NeuralNetworkUpdateInterval: k.NeuralNetworkUpdateInterval(ctx),
		NeuralNetworkLearningRate:   k.NeuralNetworkLearningRate(ctx),
		NeuralNetworkArchitecture:   k.NeuralNetworkArchitecture(ctx),
		ReputationDecayRate:         k.ReputationDecayRate(ctx),
		PerformanceAssessmentWindow: k.PerformanceAssessmentWindow(ctx),
		MinValidatorReputation:      k.MinValidatorReputation(ctx),
		MaxMissedBlocks:             k.MaxMissedBlocks(ctx),
		SignedBlocksWindow:          k.SignedBlocksWindow(ctx),
		MinSignedPerWindow:          k.MinSignedPerWindow(ctx),
		DowntimeJailDuration:        k.DowntimeJailDuration(ctx),
		SlashFractionDoubleSign:     k.SlashFractionDoubleSign(ctx),
		SlashFractionDowntime:       k.SlashFractionDowntime(ctx),
		ReputationBonusRate:         k.ReputationBonusRate(ctx),
		ReputationPenaltyRate:       k.ReputationPenaltyRate(ctx),
		NeuralNetworkInfluenceRate:  k.NeuralNetworkInfluenceRate(ctx),
	}
}

// SetParams sets the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// UnbondingTime returns the unbonding time param
func (k Keeper) UnbondingTime(ctx sdk.Context) (res time.Duration) {
	k.paramstore.Get(ctx, types.KeyUnbondingTime, &res)
	return
}

// MaxValidators returns the maximum number of validators param
func (k Keeper) MaxValidators(ctx sdk.Context) (res uint32) {
	k.paramstore.Get(ctx, types.KeyMaxValidators, &res)
	return
}

// MinSelfDelegation returns the minimum self delegation param
func (k Keeper) MinSelfDelegation(ctx sdk.Context) (res int64) {
	k.paramstore.Get(ctx, types.KeyMinSelfDelegation, &res)
	return
}

// HistoricalEntries returns the historical entries param
func (k Keeper) HistoricalEntries(ctx sdk.Context) (res uint32) {
	k.paramstore.Get(ctx, types.KeyHistoricalEntries, &res)
	return
}

// NeuralNetworkUpdateInterval returns the neural network update interval param
func (k Keeper) NeuralNetworkUpdateInterval(ctx sdk.Context) (res time.Duration) {
	k.paramstore.Get(ctx, types.KeyNeuralNetworkUpdateInterval, &res)
	return
}

// NeuralNetworkLearningRate returns the neural network learning rate param
func (k Keeper) NeuralNetworkLearningRate(ctx sdk.Context) (res sdk.Dec) {
	k.paramstore.Get(ctx, types.KeyNeuralNetworkLearningRate, &res)
	return
}

// NeuralNetworkArchitecture returns the neural network architecture param
func (k Keeper) NeuralNetworkArchitecture(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyNeuralNetworkArchitecture, &res)
	return
}

// ReputationDecayRate returns the reputation decay rate param
func (k Keeper) ReputationDecayRate(ctx sdk.Context) (res sdk.Dec) {
	k.paramstore.Get(ctx, types.KeyReputationDecayRate, &res)
	return
}

// PerformanceAssessmentWindow returns the performance assessment window param
func (k Keeper) PerformanceAssessmentWindow(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyPerformanceAssessmentWindow, &res)
	return
}

// MinValidatorReputation returns the minimum validator reputation param
func (k Keeper) MinValidatorReputation(ctx sdk.Context) (res sdk.Dec) {
	k.paramstore.Get(ctx, types.KeyMinValidatorReputation, &res)
	return
}

// MaxMissedBlocks returns the maximum missed blocks param
func (k Keeper) MaxMissedBlocks(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyMaxMissedBlocks, &res)
	return
}

// SignedBlocksWindow returns the signed blocks window param
func (k Keeper) SignedBlocksWindow(ctx sdk.Context) (res int64) {
	k.paramstore.Get(ctx, types.KeySignedBlocksWindow, &res)
	return
}

// MinSignedPerWindow returns the minimum signed per window param
func (k Keeper) MinSignedPerWindow(ctx sdk.Context) (res sdk.Dec) {
	k.paramstore.Get(ctx, types.KeyMinSignedPerWindow, &res)
	return
}

// DowntimeJailDuration returns the downtime jail duration param
func (k Keeper) DowntimeJailDuration(ctx sdk.Context) (res time.Duration) {
	k.paramstore.Get(ctx, types.KeyDowntimeJailDuration, &res)
	return
}

// SlashFractionDoubleSign returns the slash fraction for double sign param
func (k Keeper) SlashFractionDoubleSign(ctx sdk.Context) (res sdk.Dec) {
	k.paramstore.Get(ctx, types.KeySlashFractionDoubleSign, &res)
	return
}

// SlashFractionDowntime returns the slash fraction for downtime param
func (k Keeper) SlashFractionDowntime(ctx sdk.Context) (res sdk.Dec) {
	k.paramstore.Get(ctx, types.KeySlashFractionDowntime, &res)
	return
}

// ReputationBonusRate returns the reputation bonus rate param
func (k Keeper) ReputationBonusRate(ctx sdk.Context) (res sdk.Dec) {
	k.paramstore.Get(ctx, types.KeyReputationBonusRate, &res)
	return
}

// ReputationPenaltyRate returns the reputation penalty rate param
func (k Keeper) ReputationPenaltyRate(ctx sdk.Context) (res sdk.Dec) {
	k.paramstore.Get(ctx, types.KeyReputationPenaltyRate, &res)
	return
}

// NeuralNetworkInfluenceRate returns the neural network influence rate param
func (k Keeper) NeuralNetworkInfluenceRate(ctx sdk.Context) (res sdk.Dec) {
	k.paramstore.Get(ctx, types.KeyNeuralNetworkInfluenceRate, &res)
	return
}