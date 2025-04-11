package types

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Validators:             []Validator{},
		Delegations:            []Delegation{},
		UnbondingDelegations:   []UnbondingDelegation{},
		Redelegations:          []Redelegation{},
		NeuralNetworks:         []NeuralNetwork{},
		NeuralNetworkWeights:   []NeuralNetworkWeights{},
		TrainingData:           []TrainingData{},
		NeuralPredictions:      []NeuralPrediction{},
		ValidatorPerformances:  []ValidatorPerformance{},
		ValidatorReputations:   []ValidatorReputation{},
		ValidatorSlashEvents:   []ValidatorSlashEvent{},
		ValidatorSigningInfos:  []ValidatorSigningInfo{},
		Params:                 DefaultParams(),
	}
}

// Validate performs basic genesis state validation
func (gs GenesisState) Validate() error {
	// Validate validators
	validatorAddresses := make(map[string]bool)
	for _, validator := range gs.Validators {
		if validatorAddresses[validator.OperatorAddress] {
			return fmt.Errorf("duplicate validator address: %s", validator.OperatorAddress)
		}
		validatorAddresses[validator.OperatorAddress] = true

		if validator.DelegatorShares.IsNegative() {
			return fmt.Errorf("validator delegator shares cannot be negative: %s", validator.DelegatorShares)
		}

		if validator.Tokens.IsNegative() {
			return fmt.Errorf("validator tokens cannot be negative: %s", validator.Tokens)
		}

		if validator.MinSelfDelegation.IsNegative() {
			return fmt.Errorf("validator min self delegation cannot be negative: %s", validator.MinSelfDelegation)
		}

		if validator.Commission.CommissionRates.Rate.IsNegative() || validator.Commission.CommissionRates.Rate.GT(sdk.OneDec()) {
			return fmt.Errorf("validator commission rate must be between 0 and 1: %s", validator.Commission.CommissionRates.Rate)
		}

		if validator.Commission.CommissionRates.MaxRate.IsNegative() || validator.Commission.CommissionRates.MaxRate.GT(sdk.OneDec()) {
			return fmt.Errorf("validator max commission rate must be between 0 and 1: %s", validator.Commission.CommissionRates.MaxRate)
		}

		if validator.Commission.CommissionRates.MaxChangeRate.IsNegative() || validator.Commission.CommissionRates.MaxChangeRate.GT(sdk.OneDec()) {
			return fmt.Errorf("validator max commission change rate must be between 0 and 1: %s", validator.Commission.CommissionRates.MaxChangeRate)
		}

		if validator.Reputation.IsNegative() || validator.Reputation.GT(sdk.OneDec()) {
			return fmt.Errorf("validator reputation must be between 0 and 1: %s", validator.Reputation)
		}

		if validator.PerformanceScore.IsNegative() || validator.PerformanceScore.GT(sdk.OneDec()) {
			return fmt.Errorf("validator performance score must be between 0 and 1: %s", validator.PerformanceScore)
		}

		if validator.NeuralNetworkContribution.IsNegative() || validator.NeuralNetworkContribution.GT(sdk.OneDec()) {
			return fmt.Errorf("validator neural network contribution must be between 0 and 1: %s", validator.NeuralNetworkContribution)
		}
	}

	// Validate delegations
	delegationKeys := make(map[string]bool)
	for _, delegation := range gs.Delegations {
		key := fmt.Sprintf("%s/%s", delegation.DelegatorAddress, delegation.ValidatorAddress)
		if delegationKeys[key] {
			return fmt.Errorf("duplicate delegation key: %s", key)
		}
		delegationKeys[key] = true

		if delegation.Shares.IsNegative() {
			return fmt.Errorf("delegation shares cannot be negative: %s", delegation.Shares)
		}

		// Check if the validator exists
		if !validatorAddresses[delegation.ValidatorAddress] {
			return fmt.Errorf("delegation references non-existent validator: %s", delegation.ValidatorAddress)
		}
	}

	// Validate unbonding delegations
	unbondingDelegationKeys := make(map[string]bool)
	for _, ubd := range gs.UnbondingDelegations {
		key := fmt.Sprintf("%s/%s", ubd.DelegatorAddress, ubd.ValidatorAddress)
		if unbondingDelegationKeys[key] {
			return fmt.Errorf("duplicate unbonding delegation key: %s", key)
		}
		unbondingDelegationKeys[key] = true

		// Check if the validator exists
		if !validatorAddresses[ubd.ValidatorAddress] {
			return fmt.Errorf("unbonding delegation references non-existent validator: %s", ubd.ValidatorAddress)
		}

		// Validate entries
		for _, entry := range ubd.Entries {
			if entry.InitialBalance.IsNegative() {
				return fmt.Errorf("unbonding delegation initial balance cannot be negative: %s", entry.InitialBalance)
			}

			if entry.Balance.IsNegative() {
				return fmt.Errorf("unbonding delegation balance cannot be negative: %s", entry.Balance)
			}
		}
	}

	// Validate redelegations
	redelegationKeys := make(map[string]bool)
	for _, red := range gs.Redelegations {
		key := fmt.Sprintf("%s/%s/%s", red.DelegatorAddress, red.ValidatorSrcAddress, red.ValidatorDstAddress)
		if redelegationKeys[key] {
			return fmt.Errorf("duplicate redelegation key: %s", key)
		}
		redelegationKeys[key] = true

		// Check if the validators exist
		if !validatorAddresses[red.ValidatorSrcAddress] {
			return fmt.Errorf("redelegation references non-existent source validator: %s", red.ValidatorSrcAddress)
		}

		if !validatorAddresses[red.ValidatorDstAddress] {
			return fmt.Errorf("redelegation references non-existent destination validator: %s", red.ValidatorDstAddress)
		}

		// Validate entries
		for _, entry := range red.Entries {
			if entry.InitialBalance.IsNegative() {
				return fmt.Errorf("redelegation initial balance cannot be negative: %s", entry.InitialBalance)
			}

			if entry.SharesDst.IsNegative() {
				return fmt.Errorf("redelegation shares dst cannot be negative: %s", entry.SharesDst)
			}
		}
	}

	// Validate neural networks
	neuralNetworkIDs := make(map[string]bool)
	for _, nn := range gs.NeuralNetworks {
		if neuralNetworkIDs[nn.ID] {
			return fmt.Errorf("duplicate neural network ID: %s", nn.ID)
		}
		neuralNetworkIDs[nn.ID] = true

		if nn.Accuracy.IsNegative() || nn.Accuracy.GT(sdk.OneDec()) {
			return fmt.Errorf("neural network accuracy must be between 0 and 1: %s", nn.Accuracy)
		}

		if nn.Loss.IsNegative() {
			return fmt.Errorf("neural network loss cannot be negative: %s", nn.Loss)
		}
	}

	// Validate neural network weights
	weightKeys := make(map[string]bool)
	for _, weight := range gs.NeuralNetworkWeights {
		key := fmt.Sprintf("%s/%d", weight.NetworkID, weight.Version)
		if weightKeys[key] {
			return fmt.Errorf("duplicate neural network weight key: %s", key)
		}
		weightKeys[key] = true

		// Check if the neural network exists
		if !neuralNetworkIDs[weight.NetworkID] {
			return fmt.Errorf("neural network weight references non-existent neural network: %s", weight.NetworkID)
		}

		// Validate weights
		if len(weight.Weights) == 0 {
			return fmt.Errorf("neural network weights cannot be empty")
		}
	}

	// Validate training data
	trainingDataIDs := make(map[string]bool)
	for _, td := range gs.TrainingData {
		if trainingDataIDs[td.ID] {
			return fmt.Errorf("duplicate training data ID: %s", td.ID)
		}
		trainingDataIDs[td.ID] = true

		// Check if the neural network exists
		if !neuralNetworkIDs[td.NetworkID] {
			return fmt.Errorf("training data references non-existent neural network: %s", td.NetworkID)
		}

		// Validate features and labels
		if len(td.Features) == 0 {
			return fmt.Errorf("training data features cannot be empty")
		}

		if len(td.Labels) == 0 {
			return fmt.Errorf("training data labels cannot be empty")
		}
	}

	// Validate neural predictions
	predictionIDs := make(map[string]bool)
	for _, pred := range gs.NeuralPredictions {
		if predictionIDs[pred.ID] {
			return fmt.Errorf("duplicate neural prediction ID: %s", pred.ID)
		}
		predictionIDs[pred.ID] = true

		// Check if the neural network exists
		if !neuralNetworkIDs[pred.NetworkID] {
			return fmt.Errorf("neural prediction references non-existent neural network: %s", pred.NetworkID)
		}

		// Validate input and output
		if len(pred.Input) == 0 {
			return fmt.Errorf("neural prediction input cannot be empty")
		}

		if len(pred.Output) == 0 {
			return fmt.Errorf("neural prediction output cannot be empty")
		}

		if pred.Confidence.IsNegative() || pred.Confidence.GT(sdk.OneDec()) {
			return fmt.Errorf("neural prediction confidence must be between 0 and 1: %s", pred.Confidence)
		}

		// Check if the validators exist
		for _, valAddr := range pred.ValidatorSet {
			if !validatorAddresses[valAddr] {
				return fmt.Errorf("neural prediction references non-existent validator: %s", valAddr)
			}
		}
	}

	// Validate validator performances
	performanceKeys := make(map[string]bool)
	for _, perf := range gs.ValidatorPerformances {
		if performanceKeys[perf.ValidatorAddress] {
			return fmt.Errorf("duplicate validator performance key: %s", perf.ValidatorAddress)
		}
		performanceKeys[perf.ValidatorAddress] = true

		// Check if the validator exists
		if !validatorAddresses[perf.ValidatorAddress] {
			return fmt.Errorf("validator performance references non-existent validator: %s", perf.ValidatorAddress)
		}

		if perf.PredictionAccuracy.IsNegative() || perf.PredictionAccuracy.GT(sdk.OneDec()) {
			return fmt.Errorf("validator prediction accuracy must be between 0 and 1: %s", perf.PredictionAccuracy)
		}

		if perf.PerformanceScore.IsNegative() || perf.PerformanceScore.GT(sdk.OneDec()) {
			return fmt.Errorf("validator performance score must be between 0 and 1: %s", perf.PerformanceScore)
		}
	}

	// Validate validator reputations
	reputationKeys := make(map[string]bool)
	for _, rep := range gs.ValidatorReputations {
		if reputationKeys[rep.ValidatorAddress] {
			return fmt.Errorf("duplicate validator reputation key: %s", rep.ValidatorAddress)
		}
		reputationKeys[rep.ValidatorAddress] = true

		// Check if the validator exists
		if !validatorAddresses[rep.ValidatorAddress] {
			return fmt.Errorf("validator reputation references non-existent validator: %s", rep.ValidatorAddress)
		}

		if rep.Reputation.IsNegative() || rep.Reputation.GT(sdk.OneDec()) {
			return fmt.Errorf("validator reputation must be between 0 and 1: %s", rep.Reputation)
		}

		// Validate history
		for _, change := range rep.History {
			if change.Change.IsZero() {
				return fmt.Errorf("reputation change cannot be zero")
			}
		}
	}

	// Validate validator slash events
	slashEventKeys := make(map[string]bool)
	for _, event := range gs.ValidatorSlashEvents {
		key := fmt.Sprintf("%s/%d", event.ValidatorAddress, event.Height)
		if slashEventKeys[key] {
			return fmt.Errorf("duplicate validator slash event key: %s", key)
		}
		slashEventKeys[key] = true

		// Check if the validator exists
		if !validatorAddresses[event.ValidatorAddress] {
			return fmt.Errorf("validator slash event references non-existent validator: %s", event.ValidatorAddress)
		}

		if event.SlashFactor.IsNegative() || event.SlashFactor.GT(sdk.OneDec()) {
			return fmt.Errorf("validator slash factor must be between 0 and 1: %s", event.SlashFactor)
		}

		if event.Tokens.IsNegative() {
			return fmt.Errorf("validator slash tokens cannot be negative: %s", event.Tokens)
		}
	}

	// Validate validator signing infos
	signingInfoKeys := make(map[string]bool)
	for _, info := range gs.ValidatorSigningInfos {
		if signingInfoKeys[info.ValidatorAddress] {
			return fmt.Errorf("duplicate validator signing info key: %s", info.ValidatorAddress)
		}
		signingInfoKeys[info.ValidatorAddress] = true

		// Check if the validator exists
		if !validatorAddresses[info.ValidatorAddress] {
			return fmt.Errorf("validator signing info references non-existent validator: %s", info.ValidatorAddress)
		}

		if info.MinSignedPerWindow.IsNegative() || info.MinSignedPerWindow.GT(sdk.OneDec()) {
			return fmt.Errorf("validator min signed per window must be between 0 and 1: %s", info.MinSignedPerWindow)
		}
	}

	// Validate params
	if err := gs.Params.Validate(); err != nil {
		return fmt.Errorf("invalid params: %w", err)
	}

	return nil
}

// GetGenesisStateFromAppState returns the genesis state from the app state
func GetGenesisStateFromAppState(appState map[string]json.RawMessage) *GenesisState {
	var genesisState GenesisState

	if appState[ModuleName] != nil {
		ModuleCdc.MustUnmarshalJSON(appState[ModuleName], &genesisState)
	}

	return &genesisState
}