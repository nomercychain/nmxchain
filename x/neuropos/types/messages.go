package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateValidator          = "create_validator"
	TypeMsgEditValidator            = "edit_validator"
	TypeMsgDelegate                 = "delegate"
	TypeMsgUndelegate               = "undelegate"
	TypeMsgBeginRedelegate          = "begin_redelegate"
	TypeMsgUnjail                   = "unjail"
	TypeMsgUpdateNeuralNetwork      = "update_neural_network"
	TypeMsgTrainNeuralNetwork       = "train_neural_network"
	TypeMsgSubmitNeuralPrediction   = "submit_neural_prediction"
	TypeMsgUpdateValidatorReputation = "update_validator_reputation"
)

var _ sdk.Msg = &MsgCreateValidator{}

// NewMsgCreateValidator creates a new MsgCreateValidator instance
func NewMsgCreateValidator(
	valAddr sdk.ValAddress,
	pubKey string,
	selfDelegation sdk.Coin,
	description Description,
	commissionRates CommissionRates,
	minSelfDelegation sdk.Int,
) *MsgCreateValidator {
	return &MsgCreateValidator{
		ValidatorAddress:  valAddr.String(),
		Pubkey:           pubKey,
		SelfDelegation:   selfDelegation,
		Description:      description,
		CommissionRates:  commissionRates,
		MinSelfDelegation: minSelfDelegation,
	}
}

// Route implements Msg
func (msg MsgCreateValidator) Route() string { return RouterKey }

// Type implements Msg
func (msg MsgCreateValidator) Type() string { return TypeMsgCreateValidator }

// ValidateBasic implements Msg
func (msg MsgCreateValidator) ValidateBasic() error {
	// Validate validator address
	_, err := sdk.ValAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validator address: %s", err)
	}

	// Validate pubkey
	if msg.Pubkey == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidPubKey, "empty pubkey")
	}

	// Validate self delegation
	if !msg.SelfDelegation.IsValid() || !msg.SelfDelegation.Amount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "self delegation must be a valid positive amount")
	}

	// Validate description
	if msg.Description.Moniker == "" {
		return sdkerrors.Wrap(ErrInvalidMoniker, "moniker cannot be empty")
	}

	// Validate commission rates
	if msg.CommissionRates.Rate.IsNegative() || msg.CommissionRates.Rate.GT(sdk.OneDec()) {
		return sdkerrors.Wrap(ErrInvalidCommissionRate, "commission rate must be between 0 and 1")
	}

	if msg.CommissionRates.MaxRate.IsNegative() || msg.CommissionRates.MaxRate.GT(sdk.OneDec()) {
		return sdkerrors.Wrap(ErrInvalidCommissionRate, "max commission rate must be between 0 and 1")
	}

	if msg.CommissionRates.MaxChangeRate.IsNegative() || msg.CommissionRates.MaxChangeRate.GT(sdk.OneDec()) {
		return sdkerrors.Wrap(ErrInvalidCommissionRate, "max commission change rate must be between 0 and 1")
	}

	if msg.CommissionRates.Rate.GT(msg.CommissionRates.MaxRate) {
		return sdkerrors.Wrap(ErrInvalidCommissionRate, "commission rate cannot be greater than max rate")
	}

	// Validate min self delegation
	if !msg.MinSelfDelegation.IsPositive() {
		return sdkerrors.Wrap(ErrInvalidInput, "minimum self delegation must be positive")
	}

	if msg.SelfDelegation.Amount.LT(msg.MinSelfDelegation) {
		return sdkerrors.Wrap(ErrSelfDelegationBelowMinimum, "self delegation below minimum")
	}

	return nil
}

// GetSignBytes implements Msg
func (msg MsgCreateValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg
func (msg MsgCreateValidator) GetSigners() []sdk.AccAddress {
	valAddr, err := sdk.ValAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sdk.AccAddress(valAddr)}
}

var _ sdk.Msg = &MsgEditValidator{}

// NewMsgEditValidator creates a new MsgEditValidator instance
func NewMsgEditValidator(
	valAddr sdk.ValAddress,
	description Description,
	commissionRate *sdk.Dec,
	minSelfDelegation *sdk.Int,
) *MsgEditValidator {
	return &MsgEditValidator{
		ValidatorAddress:  valAddr.String(),
		Description:      description,
		CommissionRate:   commissionRate,
		MinSelfDelegation: minSelfDelegation,
	}
}

// Route implements Msg
func (msg MsgEditValidator) Route() string { return RouterKey }

// Type implements Msg
func (msg MsgEditValidator) Type() string { return TypeMsgEditValidator }

// ValidateBasic implements Msg
func (msg MsgEditValidator) ValidateBasic() error {
	// Validate validator address
	_, err := sdk.ValAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validator address: %s", err)
	}

	// Validate description
	if msg.Description.Moniker == "" {
		return sdkerrors.Wrap(ErrInvalidMoniker, "moniker cannot be empty")
	}

	// Validate commission rate
	if msg.CommissionRate != nil {
		if msg.CommissionRate.IsNegative() || msg.CommissionRate.GT(sdk.OneDec()) {
			return sdkerrors.Wrap(ErrInvalidCommissionRate, "commission rate must be between 0 and 1")
		}
	}

	// Validate min self delegation
	if msg.MinSelfDelegation != nil {
		if !msg.MinSelfDelegation.IsPositive() {
			return sdkerrors.Wrap(ErrInvalidInput, "minimum self delegation must be positive")
		}
	}

	return nil
}

// GetSignBytes implements Msg
func (msg MsgEditValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg
func (msg MsgEditValidator) GetSigners() []sdk.AccAddress {
	valAddr, err := sdk.ValAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sdk.AccAddress(valAddr)}
}

var _ sdk.Msg = &MsgDelegate{}

// NewMsgDelegate creates a new MsgDelegate instance
func NewMsgDelegate(
	delAddr sdk.AccAddress,
	valAddr sdk.ValAddress,
	amount sdk.Coin,
) *MsgDelegate {
	return &MsgDelegate{
		DelegatorAddress: delAddr.String(),
		ValidatorAddress: valAddr.String(),
		Amount:          amount,
	}
}

// Route implements Msg
func (msg MsgDelegate) Route() string { return RouterKey }

// Type implements Msg
func (msg MsgDelegate) Type() string { return TypeMsgDelegate }

// ValidateBasic implements Msg
func (msg MsgDelegate) ValidateBasic() error {
	// Validate delegator address
	_, err := sdk.AccAddressFromBech32(msg.DelegatorAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid delegator address: %s", err)
	}

	// Validate validator address
	_, err = sdk.ValAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validator address: %s", err)
	}

	// Validate amount
	if !msg.Amount.IsValid() || !msg.Amount.Amount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "delegation amount must be a valid positive amount")
	}

	return nil
}

// GetSignBytes implements Msg
func (msg MsgDelegate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg
func (msg MsgDelegate) GetSigners() []sdk.AccAddress {
	delAddr, err := sdk.AccAddressFromBech32(msg.DelegatorAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{delAddr}
}

var _ sdk.Msg = &MsgUndelegate{}

// NewMsgUndelegate creates a new MsgUndelegate instance
func NewMsgUndelegate(
	delAddr sdk.AccAddress,
	valAddr sdk.ValAddress,
	amount sdk.Coin,
) *MsgUndelegate {
	return &MsgUndelegate{
		DelegatorAddress: delAddr.String(),
		ValidatorAddress: valAddr.String(),
		Amount:          amount,
	}
}

// Route implements Msg
func (msg MsgUndelegate) Route() string { return RouterKey }

// Type implements Msg
func (msg MsgUndelegate) Type() string { return TypeMsgUndelegate }

// ValidateBasic implements Msg
func (msg MsgUndelegate) ValidateBasic() error {
	// Validate delegator address
	_, err := sdk.AccAddressFromBech32(msg.DelegatorAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid delegator address: %s", err)
	}

	// Validate validator address
	_, err = sdk.ValAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validator address: %s", err)
	}

	// Validate amount
	if !msg.Amount.IsValid() || !msg.Amount.Amount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "unbonding amount must be a valid positive amount")
	}

	return nil
}

// GetSignBytes implements Msg
func (msg MsgUndelegate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg
func (msg MsgUndelegate) GetSigners() []sdk.AccAddress {
	delAddr, err := sdk.AccAddressFromBech32(msg.DelegatorAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{delAddr}
}

var _ sdk.Msg = &MsgBeginRedelegate{}

// NewMsgBeginRedelegate creates a new MsgBeginRedelegate instance
func NewMsgBeginRedelegate(
	delAddr sdk.AccAddress,
	valSrcAddr, valDstAddr sdk.ValAddress,
	amount sdk.Coin,
) *MsgBeginRedelegate {
	return &MsgBeginRedelegate{
		DelegatorAddress:    delAddr.String(),
		ValidatorSrcAddress: valSrcAddr.String(),
		ValidatorDstAddress: valDstAddr.String(),
		Amount:             amount,
	}
}

// Route implements Msg
func (msg MsgBeginRedelegate) Route() string { return RouterKey }

// Type implements Msg
func (msg MsgBeginRedelegate) Type() string { return TypeMsgBeginRedelegate }

// ValidateBasic implements Msg
func (msg MsgBeginRedelegate) ValidateBasic() error {
	// Validate delegator address
	_, err := sdk.AccAddressFromBech32(msg.DelegatorAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid delegator address: %s", err)
	}

	// Validate source validator address
	_, err = sdk.ValAddressFromBech32(msg.ValidatorSrcAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid source validator address: %s", err)
	}

	// Validate destination validator address
	_, err = sdk.ValAddressFromBech32(msg.ValidatorDstAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid destination validator address: %s", err)
	}

	// Validate amount
	if !msg.Amount.IsValid() || !msg.Amount.Amount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "redelegation amount must be a valid positive amount")
	}

	// Check that source and destination validators are different
	if msg.ValidatorSrcAddress == msg.ValidatorDstAddress {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "redelegation source and destination validators cannot be the same")
	}

	return nil
}

// GetSignBytes implements Msg
func (msg MsgBeginRedelegate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg
func (msg MsgBeginRedelegate) GetSigners() []sdk.AccAddress {
	delAddr, err := sdk.AccAddressFromBech32(msg.DelegatorAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{delAddr}
}

var _ sdk.Msg = &MsgUnjail{}

// NewMsgUnjail creates a new MsgUnjail instance
func NewMsgUnjail(valAddr sdk.ValAddress) *MsgUnjail {
	return &MsgUnjail{
		ValidatorAddress: valAddr.String(),
	}
}

// Route implements Msg
func (msg MsgUnjail) Route() string { return RouterKey }

// Type implements Msg
func (msg MsgUnjail) Type() string { return TypeMsgUnjail }

// ValidateBasic implements Msg
func (msg MsgUnjail) ValidateBasic() error {
	// Validate validator address
	_, err := sdk.ValAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validator address: %s", err)
	}

	return nil
}

// GetSignBytes implements Msg
func (msg MsgUnjail) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg
func (msg MsgUnjail) GetSigners() []sdk.AccAddress {
	valAddr, err := sdk.ValAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sdk.AccAddress(valAddr)}
}

var _ sdk.Msg = &MsgUpdateNeuralNetwork{}

// NewMsgUpdateNeuralNetwork creates a new MsgUpdateNeuralNetwork instance
func NewMsgUpdateNeuralNetwork(
	valAddr sdk.ValAddress,
	networkID string,
	architecture string,
	layers []Layer,
	weights json.RawMessage,
	metadata []byte,
) *MsgUpdateNeuralNetwork {
	return &MsgUpdateNeuralNetwork{
		ValidatorAddress: valAddr.String(),
		NetworkId:       networkID,
		Architecture:    architecture,
		Layers:          layers,
		Weights:         weights,
		Metadata:        metadata,
	}
}

// Route implements Msg
func (msg MsgUpdateNeuralNetwork) Route() string { return RouterKey }

// Type implements Msg
func (msg MsgUpdateNeuralNetwork) Type() string { return TypeMsgUpdateNeuralNetwork }

// ValidateBasic implements Msg
func (msg MsgUpdateNeuralNetwork) ValidateBasic() error {
	// Validate validator address
	_, err := sdk.ValAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validator address: %s", err)
	}

	// Validate network ID
	if msg.NetworkId == "" {
		return sdkerrors.Wrap(ErrNoNeuralNetworkFound, "network ID cannot be empty")
	}

	// Validate architecture
	valid := false
	validArchitectures := []string{
		NeuralNetworkArchitectureMLP,
		NeuralNetworkArchitectureCNN,
		NeuralNetworkArchitectureRNN,
		NeuralNetworkArchitectureLSTM,
		NeuralNetworkArchitectureGRU,
		NeuralNetworkArchitectureTransformer,
		NeuralNetworkArchitectureAutoencoder,
		NeuralNetworkArchitectureGAN,
	}

	for _, arch := range validArchitectures {
		if msg.Architecture == arch {
			valid = true
			break
		}
	}

	if !valid {
		return sdkerrors.Wrap(ErrInvalidNeuralNetworkArchitecture, msg.Architecture)
	}

	// Validate layers
	if len(msg.Layers) == 0 {
		return sdkerrors.Wrap(ErrInvalidNeuralNetworkArchitecture, "layers cannot be empty")
	}

	// Validate weights
	if len(msg.Weights) == 0 {
		return sdkerrors.Wrap(ErrInvalidNeuralNetworkWeights, "weights cannot be empty")
	}

	return nil
}

// GetSignBytes implements Msg
func (msg MsgUpdateNeuralNetwork) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg
func (msg MsgUpdateNeuralNetwork) GetSigners() []sdk.AccAddress {
	valAddr, err := sdk.ValAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sdk.AccAddress(valAddr)}
}

var _ sdk.Msg = &MsgTrainNeuralNetwork{}

// NewMsgTrainNeuralNetwork creates a new MsgTrainNeuralNetwork instance
func NewMsgTrainNeuralNetwork(
	valAddr sdk.ValAddress,
	networkID string,
	features json.RawMessage,
	labels json.RawMessage,
	epochs uint64,
	learningRate sdk.Dec,
	metadata []byte,
) *MsgTrainNeuralNetwork {
	return &MsgTrainNeuralNetwork{
		ValidatorAddress: valAddr.String(),
		NetworkId:       networkID,
		Features:        features,
		Labels:          labels,
		Epochs:          epochs,
		LearningRate:    learningRate,
		Metadata:        metadata,
	}
}

// Route implements Msg
func (msg MsgTrainNeuralNetwork) Route() string { return RouterKey }

// Type implements Msg
func (msg MsgTrainNeuralNetwork) Type() string { return TypeMsgTrainNeuralNetwork }

// ValidateBasic implements Msg
func (msg MsgTrainNeuralNetwork) ValidateBasic() error {
	// Validate validator address
	_, err := sdk.ValAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validator address: %s", err)
	}

	// Validate network ID
	if msg.NetworkId == "" {
		return sdkerrors.Wrap(ErrNoNeuralNetworkFound, "network ID cannot be empty")
	}

	// Validate features
	if len(msg.Features) == 0 {
		return sdkerrors.Wrap(ErrInvalidTrainingData, "features cannot be empty")
	}

	// Validate labels
	if len(msg.Labels) == 0 {
		return sdkerrors.Wrap(ErrInvalidTrainingData, "labels cannot be empty")
	}

	// Validate epochs
	if msg.Epochs == 0 {
		return sdkerrors.Wrap(ErrInvalidEpochs, "epochs cannot be zero")
	}

	// Validate learning rate
	if msg.LearningRate.IsNegative() || msg.LearningRate.GT(sdk.OneDec()) {
		return sdkerrors.Wrap(ErrInvalidLearningRate, "learning rate must be between 0 and 1")
	}

	return nil
}

// GetSignBytes implements Msg
func (msg MsgTrainNeuralNetwork) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg
func (msg MsgTrainNeuralNetwork) GetSigners() []sdk.AccAddress {
	valAddr, err := sdk.ValAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sdk.AccAddress(valAddr)}
}

var _ sdk.Msg = &MsgSubmitNeuralPrediction{}

// NewMsgSubmitNeuralPrediction creates a new MsgSubmitNeuralPrediction instance
func NewMsgSubmitNeuralPrediction(
	valAddr sdk.ValAddress,
	networkID string,
	input json.RawMessage,
	output json.RawMessage,
	confidence sdk.Dec,
	metadata []byte,
) *MsgSubmitNeuralPrediction {
	return &MsgSubmitNeuralPrediction{
		ValidatorAddress: valAddr.String(),
		NetworkId:       networkID,
		Input:           input,
		Output:          output,
		Confidence:      confidence,
		Metadata:        metadata,
	}
}

// Route implements Msg
func (msg MsgSubmitNeuralPrediction) Route() string { return RouterKey }

// Type implements Msg
func (msg MsgSubmitNeuralPrediction) Type() string { return TypeMsgSubmitNeuralPrediction }

// ValidateBasic implements Msg
func (msg MsgSubmitNeuralPrediction) ValidateBasic() error {
	// Validate validator address
	_, err := sdk.ValAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validator address: %s", err)
	}

	// Validate network ID
	if msg.NetworkId == "" {
		return sdkerrors.Wrap(ErrNoNeuralNetworkFound, "network ID cannot be empty")
	}

	// Validate input
	if len(msg.Input) == 0 {
		return sdkerrors.Wrap(ErrInvalidPredictionInput, "input cannot be empty")
	}

	// Validate output
	if len(msg.Output) == 0 {
		return sdkerrors.Wrap(ErrInvalidPredictionInput, "output cannot be empty")
	}

	// Validate confidence
	if msg.Confidence.IsNegative() || msg.Confidence.GT(sdk.OneDec()) {
		return sdkerrors.Wrap(ErrInvalidPredictionInput, "confidence must be between 0 and 1")
	}

	return nil
}

// GetSignBytes implements Msg
func (msg MsgSubmitNeuralPrediction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg
func (msg MsgSubmitNeuralPrediction) GetSigners() []sdk.AccAddress {
	valAddr, err := sdk.ValAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sdk.AccAddress(valAddr)}
}

var _ sdk.Msg = &MsgUpdateValidatorReputation{}

// NewMsgUpdateValidatorReputation creates a new MsgUpdateValidatorReputation instance
func NewMsgUpdateValidatorReputation(
	adminAddr sdk.AccAddress,
	valAddr sdk.ValAddress,
	reputationChange sdk.Dec,
	reason string,
) *MsgUpdateValidatorReputation {
	return &MsgUpdateValidatorReputation{
		AdminAddress:      adminAddr.String(),
		ValidatorAddress:  valAddr.String(),
		ReputationChange:  reputationChange,
		Reason:           reason,
	}
}

// Route implements Msg
func (msg MsgUpdateValidatorReputation) Route() string { return RouterKey }

// Type implements Msg
func (msg MsgUpdateValidatorReputation) Type() string { return TypeMsgUpdateValidatorReputation }

// ValidateBasic implements Msg
func (msg MsgUpdateValidatorReputation) ValidateBasic() error {
	// Validate admin address
	_, err := sdk.AccAddressFromBech32(msg.AdminAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address: %s", err)
	}

	// Validate validator address
	_, err = sdk.ValAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validator address: %s", err)
	}

	// Validate reputation change
	if msg.ReputationChange.IsZero() {
		return sdkerrors.Wrap(ErrInvalidReputationChange, "reputation change cannot be zero")
	}

	// Validate reason
	if msg.Reason == "" {
		return sdkerrors.Wrap(ErrInvalidReputationChange, "reason cannot be empty")
	}

	return nil
}

// GetSignBytes implements Msg
func (msg MsgUpdateValidatorReputation) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg
func (msg MsgUpdateValidatorReputation) GetSigners() []sdk.AccAddress {
	adminAddr, err := sdk.AccAddressFromBech32(msg.AdminAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{adminAddr}
}