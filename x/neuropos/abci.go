package neuropos

import (
	"time"

	"github.com/nomercychain/nmxchain/x/neuropos/keeper"
	"github.com/nomercychain/nmxchain/x/neuropos/types"
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeginBlocker is called at the beginning of every block
func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	// Update network state metrics
	k.UpdateNetworkState(ctx)

	// Process missed blocks for validators
	processMissedBlocks(ctx, req, k)

	// Update neural networks periodically
	updateNeuralNetworks(ctx, k)

	// Detect anomalies using AI models
	anomalyReports := k.DetectAnomalies(ctx)

	// Process anomaly reports
	for _, report := range anomalyReports {
		// Store the anomaly report
		k.SetAnomalyReport(ctx, report)

		// If the anomaly is severe enough, take action
		if report.Confidence.GT(sdk.NewDecWithPrec(9, 1)) { // 90% confidence
			// Example: slash the validator that caused the anomaly
			// This is a simplified example - real implementation would be more nuanced
			validator, found := k.StakingKeeper.GetValidator(ctx, sdk.ValAddress(report.ValidatorAddress))
			if found {
				consAddr, err := validator.GetConsAddr()
				if err == nil {
					// Slash a small amount for detected anomalies
					k.StakingKeeper.Slash(ctx, consAddr, ctx.BlockHeight(), validator.ConsensusPower(sdk.DefaultPowerReduction), sdk.NewDecWithPrec(1, 2)) // 1%
					
					// Update validator reputation
					k.UpdateValidatorReputation(ctx, validator.GetOperator().String(), sdk.NewDecWithPrec(-5, 2), "anomaly detected") // -0.05
				}
			}
		}
	}
}

// EndBlocker is called at the end of every block
func EndBlocker(ctx sdk.Context, req abci.RequestEndBlock, k keeper.Keeper) []abci.ValidatorUpdate {
	// Adjust block parameters based on network state
	k.AdjustBlockParameters(ctx)

	// Update validator performances
	updateValidatorPerformances(ctx, req, k)

	// Return validator updates
	// In a real implementation, this might include AI-based validator scoring
	return []abci.ValidatorUpdate{}
}

// processSigningInfo processes the signing info for validators
func processSigningInfo(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	// Process validator signing info
	for _, voteInfo := range req.LastCommitInfo.GetVotes() {
		// Skip if validator didn't sign
		if !voteInfo.SignedLastBlock {
			continue
		}

		// Get the validator
		val := sdk.ConsAddress(voteInfo.Validator.Address)
		validator, found := k.StakingKeeper.GetValidatorByConsAddr(ctx, val)
		if !found {
			continue
		}

		// Update validator performance
		k.UpdateValidatorPerformance(ctx, validator.GetOperator().String(), false, true, false)
	}
}

// processMissedBlocks processes missed blocks for validators
func processMissedBlocks(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	// Process missed blocks
	for _, evidence := range req.ByzantineValidators {
		val := sdk.ConsAddress(evidence.Validator.Address)
		validator, found := k.StakingKeeper.GetValidatorByConsAddr(ctx, val)
		if !found {
			continue
		}

		// Update validator performance
		k.UpdateValidatorPerformance(ctx, validator.GetOperator().String(), false, false, true)

		// If double signing, slash severely
		if evidence.Type == abci.EvidenceType_DUPLICATE_VOTE {
			// Get slash fraction for double signing
			slashFraction := k.SlashFractionDoubleSign(ctx)

			// Slash the validator
			k.StakingKeeper.Slash(ctx, val, evidence.Height, validator.ConsensusPower(sdk.DefaultPowerReduction), slashFraction)
			
			// Jail the validator
			k.StakingKeeper.Jail(ctx, val)

			// Add slash event
			k.AddValidatorSlashEvent(ctx, validator.GetOperator().String(), evidence.Height, "double signing", slashFraction, sdk.TokensFromConsensusPower(validator.ConsensusPower(sdk.DefaultPowerReduction), sdk.DefaultPowerReduction))
		}
	}

	// Process validators who missed signing the last block
	for _, voteInfo := range req.LastCommitInfo.GetVotes() {
		// Skip if validator signed
		if voteInfo.SignedLastBlock {
			continue
		}

		// Get the validator
		val := sdk.ConsAddress(voteInfo.Validator.Address)
		validator, found := k.StakingKeeper.GetValidatorByConsAddr(ctx, val)
		if !found {
			continue
		}

		// Update validator performance
		k.UpdateValidatorPerformance(ctx, validator.GetOperator().String(), false, false, true)

		// Get validator signing info
		signingInfo, found := k.GetValidatorSigningInfo(ctx, validator.GetOperator().String())
		if !found {
			// Initialize signing info if not found
			signingInfo = types.ValidatorSigningInfo{
				ValidatorAddress:    validator.GetOperator().String(),
				StartHeight:         ctx.BlockHeight(),
				IndexOffset:         0,
				JailedUntil:         time.Time{},
				Tombstoned:          false,
				MissedBlocksCounter: 0,
				SignedBlocksWindow:  k.SignedBlocksWindow(ctx),
				MinSignedPerWindow: k.MinSignedPerWindow(ctx),
			}
		}

		// Increment missed blocks counter
		signingInfo.MissedBlocksCounter++

		// Check if the validator should be jailed
		if signingInfo.MissedBlocksCounter > uint64(signingInfo.SignedBlocksWindow) {
			ratio := sdk.NewDec(int64(signingInfo.MissedBlocksCounter)).Quo(sdk.NewDec(signingInfo.SignedBlocksWindow))
			if ratio.GT(sdk.OneDec().Sub(signingInfo.MinSignedPerWindow)) {
				// Jail the validator
				k.StakingKeeper.Jail(ctx, val)

				// Set jailed until time
				signingInfo.JailedUntil = ctx.BlockTime().Add(k.DowntimeJailDuration(ctx))

				// Slash the validator
				slashFraction := k.SlashFractionDowntime(ctx)
				k.StakingKeeper.Slash(ctx, val, ctx.BlockHeight(), validator.ConsensusPower(sdk.DefaultPowerReduction), slashFraction)

				// Add slash event
				k.AddValidatorSlashEvent(ctx, validator.GetOperator().String(), ctx.BlockHeight(), "downtime", slashFraction, sdk.TokensFromConsensusPower(validator.ConsensusPower(sdk.DefaultPowerReduction), sdk.DefaultPowerReduction))
			}
		}

		// Save signing info
		k.SetValidatorSigningInfo(ctx, signingInfo)
	}
}

// updateValidatorPerformances updates the performance metrics for validators
func updateValidatorPerformances(ctx sdk.Context, req abci.RequestEndBlock, k keeper.Keeper) {
	// Get the proposer for this block
	proposerConsAddr := ctx.BlockHeader().ProposerAddress
	proposer, found := k.StakingKeeper.GetValidatorByConsAddr(ctx, sdk.ConsAddress(proposerConsAddr))
	if found {
		// Update proposer performance
		k.UpdateValidatorPerformance(ctx, proposer.GetOperator().String(), true, true, false)
	}

	// Update all validators' reputations periodically
	if ctx.BlockHeight()%100 == 0 { // Every 100 blocks
		validators := k.StakingKeeper.GetAllValidators(ctx)
		for _, validator := range validators {
			// Get validator performance
			performance, found := k.GetValidatorPerformance(ctx, validator.GetOperator().String())
			if !found {
				continue
			}

			// Calculate reputation change based on performance
			params := k.GetParams(ctx)
			baseChange := performance.PerformanceScore.Sub(sdk.NewDecWithPrec(5, 1)) // 0.5 is neutral
			reputationChange := baseChange.Mul(params.ReputationBonusRate)

			// Apply neural network influence
			nnInfluence := k.CalculateNeuralNetworkInfluence(ctx, validator.GetOperator().String())
			if !nnInfluence.IsZero() {
				reputationChange = reputationChange.Add(nnInfluence.Mul(params.NeuralNetworkInfluenceRate))
			}

			// Update reputation
			k.UpdateValidatorReputation(ctx, validator.GetOperator().String(), reputationChange, "periodic update")
		}
	}
}

// updateNeuralNetworks updates neural networks periodically
func updateNeuralNetworks(ctx sdk.Context, k keeper.Keeper) {
	// Check if it's time to update neural networks
	params := k.GetParams(ctx)
	lastUpdateKey := append([]byte("last_nn_update"), 0x00)
	store := ctx.KVStore(k.storeKey)
	lastUpdateBytes := store.Get(lastUpdateKey)

	var lastUpdate time.Time
	if lastUpdateBytes == nil {
		lastUpdate = time.Time{}
	} else {
		err := lastUpdate.UnmarshalBinary(lastUpdateBytes)
		if err != nil {
			lastUpdate = time.Time{}
		}
	}

	// If it's time to update
	if ctx.BlockTime().Sub(lastUpdate) >= params.NeuralNetworkUpdateInterval {
		// Get all neural networks
		networks := k.GetAllNeuralNetworks(ctx)

		// Update each network
		for _, network := range networks {
			// Skip networks that are already being updated or trained
			if network.Status == types.NeuralNetworkStatusUpdating || network.Status == types.NeuralNetworkStatusTraining {
				continue
			}

			// Get training data for this network
			trainingData := k.GetTrainingDataByNetwork(ctx, network.ID)
			if len(trainingData) == 0 {
				continue
			}

			// Set network status to updating
			network.Status = types.NeuralNetworkStatusUpdating
			k.SetNeuralNetwork(ctx, network)

			// In a real implementation, this would trigger the actual neural network update
			// For now, we'll simulate an update by improving accuracy slightly
			network.Accuracy = sdk.MinDec(sdk.OneDec(), network.Accuracy.Add(sdk.NewDecWithPrec(1, 2))) // +0.01
			network.Loss = sdk.MaxDec(sdk.ZeroDec(), network.Loss.Sub(sdk.NewDecWithPrec(1, 2)))       // -0.01
			network.Status = types.NeuralNetworkStatusActive
			network.LastUpdatedTime = ctx.BlockTime()
			k.SetNeuralNetwork(ctx, network)

			// Emit event
			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					types.EventTypeUpdateNeuralNetwork,
					sdk.NewAttribute(types.AttributeKeyNeuralNetworkID, network.ID),
					sdk.NewAttribute(types.AttributeKeyNeuralNetworkAccuracy, network.Accuracy.String()),
					sdk.NewAttribute(types.AttributeKeyNeuralNetworkLoss, network.Loss.String()),
				),
			)
		}

		// Update last update time
		nowBytes, _ := ctx.BlockTime().MarshalBinary()
		store.Set(lastUpdateKey, nowBytes)
	}
}