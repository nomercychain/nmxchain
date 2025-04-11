package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateHyperchain{}

// ValidateBasic implements Msg
func (msg *MsgCreateHyperchain) ValidateBasic() error {
	// Validate creator address
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	// Validate name
	if msg.Name == "" {
		return sdkerrors.Wrap(ErrInvalidName, "name cannot be empty")
	}
	if len(msg.Name) > 100 {
		return sdkerrors.Wrap(ErrInvalidName, "name too long")
	}

	// Validate description
	if len(msg.Description) > 1000 {
		return sdkerrors.Wrap(ErrInvalidDescription, "description too long")
	}

	// Validate chain type
	if msg.ChainType == HyperchainTypeUnspecified {
		return sdkerrors.Wrap(ErrInvalidChainType, "chain type cannot be unspecified")
	}

	// Validate consensus type
	if msg.ConsensusType == ConsensusTypeUnspecified {
		return sdkerrors.Wrap(ErrInvalidConsensusType, "consensus type cannot be unspecified")
	}

	// Validate max validators
	if msg.MaxValidators == 0 {
		return sdkerrors.Wrap(ErrInvalidMaxValidators, "max validators cannot be zero")
	}

	// Validate deposit
	if !msg.Deposit.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Deposit.String())
	}
	if msg.Deposit.IsZero() {
		return sdkerrors.Wrap(ErrInsufficientDeposit, "deposit cannot be zero")
	}

	return nil
}

// GetSigners implements Msg
func (msg *MsgCreateHyperchain) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

var _ sdk.Msg = &MsgUpdateHyperchain{}

// ValidateBasic implements Msg
func (msg *MsgUpdateHyperchain) ValidateBasic() error {
	// Validate admin address
	_, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}

	// Validate chain ID
	if msg.ChainId == "" {
		return sdkerrors.Wrap(ErrHyperchainNotFound, "chain ID cannot be empty")
	}

	// Validate name
	if msg.Name == "" {
		return sdkerrors.Wrap(ErrInvalidName, "name cannot be empty")
	}
	if len(msg.Name) > 100 {
		return sdkerrors.Wrap(ErrInvalidName, "name too long")
	}

	// Validate description
	if len(msg.Description) > 1000 {
		return sdkerrors.Wrap(ErrInvalidDescription, "description too long")
	}

	return nil
}

// GetSigners implements Msg
func (msg *MsgUpdateHyperchain) GetSigners() []sdk.AccAddress {
	admin, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{admin}
}

var _ sdk.Msg = &MsgJoinHyperchainAsValidator{}

// ValidateBasic implements Msg
func (msg *MsgJoinHyperchainAsValidator) ValidateBasic() error {
	// Validate validator address
	_, err := sdk.AccAddressFromBech32(msg.Validator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validator address (%s)", err)
	}

	// Validate chain ID
	if msg.ChainId == "" {
		return sdkerrors.Wrap(ErrHyperchainNotFound, "chain ID cannot be empty")
	}

	// Validate pubkey
	if msg.Pubkey == "" {
		return sdkerrors.Wrap(ErrInvalidPubkey, "pubkey cannot be empty")
	}

	// Validate stake
	if !msg.Stake.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Stake.String())
	}
	if msg.Stake.IsZero() {
		return sdkerrors.Wrap(ErrInsufficientStake, "stake cannot be zero")
	}

	return nil
}

// GetSigners implements Msg
func (msg *MsgJoinHyperchainAsValidator) GetSigners() []sdk.AccAddress {
	validator, err := sdk.AccAddressFromBech32(msg.Validator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{validator}
}

var _ sdk.Msg = &MsgLeaveHyperchain{}

// ValidateBasic implements Msg
func (msg *MsgLeaveHyperchain) ValidateBasic() error {
	// Validate validator address
	_, err := sdk.AccAddressFromBech32(msg.Validator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validator address (%s)", err)
	}

	// Validate chain ID
	if msg.ChainId == "" {
		return sdkerrors.Wrap(ErrHyperchainNotFound, "chain ID cannot be empty")
	}

	return nil
}

// GetSigners implements Msg
func (msg *MsgLeaveHyperchain) GetSigners() []sdk.AccAddress {
	validator, err := sdk.AccAddressFromBech32(msg.Validator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{validator}
}

var _ sdk.Msg = &MsgCreateHyperchainBridge{}

// ValidateBasic implements Msg
func (msg *MsgCreateHyperchainBridge) ValidateBasic() error {
	// Validate creator address
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	// Validate source chain ID
	if msg.SourceChainId == "" {
		return sdkerrors.Wrap(ErrHyperchainNotFound, "source chain ID cannot be empty")
	}

	// Validate target chain ID
	if msg.TargetChainId == "" {
		return sdkerrors.Wrap(ErrHyperchainNotFound, "target chain ID cannot be empty")
	}

	// Validate min relayers
	if msg.MinRelayers == 0 {
		return sdkerrors.Wrap(ErrInvalidMinRelayers, "min relayers cannot be zero")
	}

	// Validate supported tokens
	if len(msg.SupportedTokens) == 0 {
		return sdkerrors.Wrap(ErrInvalidSupportedTokens, "supported tokens cannot be empty")
	}

	return nil
}

// GetSigners implements Msg
func (msg *MsgCreateHyperchainBridge) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

var _ sdk.Msg = &MsgUpdateHyperchainBridge{}

// ValidateBasic implements Msg
func (msg *MsgUpdateHyperchainBridge) ValidateBasic() error {
	// Validate admin address
	_, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}

	// Validate bridge ID
	if msg.BridgeId == "" {
		return sdkerrors.Wrap(ErrBridgeNotFound, "bridge ID cannot be empty")
	}

	return nil
}

// GetSigners implements Msg
func (msg *MsgUpdateHyperchainBridge) GetSigners() []sdk.AccAddress {
	admin, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{admin}
}

var _ sdk.Msg = &MsgRegisterHyperchainBridgeRelayer{}

// ValidateBasic implements Msg
func (msg *MsgRegisterHyperchainBridgeRelayer) ValidateBasic() error {
	// Validate admin address
	_, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}

	// Validate relayer address
	_, err = sdk.AccAddressFromBech32(msg.Relayer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid relayer address (%s)", err)
	}

	// Validate bridge ID
	if msg.BridgeId == "" {
		return sdkerrors.Wrap(ErrBridgeNotFound, "bridge ID cannot be empty")
	}

	return nil
}

// GetSigners implements Msg
func (msg *MsgRegisterHyperchainBridgeRelayer) GetSigners() []sdk.AccAddress {
	admin, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{admin}
}

var _ sdk.Msg = &MsgRemoveHyperchainBridgeRelayer{}

// ValidateBasic implements Msg
func (msg *MsgRemoveHyperchainBridgeRelayer) ValidateBasic() error {
	// Validate admin address
	_, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}

	// Validate relayer address
	_, err = sdk.AccAddressFromBech32(msg.Relayer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid relayer address (%s)", err)
	}

	// Validate bridge ID
	if msg.BridgeId == "" {
		return sdkerrors.Wrap(ErrBridgeNotFound, "bridge ID cannot be empty")
	}

	return nil
}

// GetSigners implements Msg
func (msg *MsgRemoveHyperchainBridgeRelayer) GetSigners() []sdk.AccAddress {
	admin, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{admin}
}

var _ sdk.Msg = &MsgInitiateHyperchainBridgeTransaction{}

// ValidateBasic implements Msg
func (msg *MsgInitiateHyperchainBridgeTransaction) ValidateBasic() error {
	// Validate sender address
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	// Validate recipient address
	_, err = sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid recipient address (%s)", err)
	}

	// Validate bridge ID
	if msg.BridgeId == "" {
		return sdkerrors.Wrap(ErrBridgeNotFound, "bridge ID cannot be empty")
	}

	// Validate amount
	if !msg.Amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}
	if msg.Amount.IsZero() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "amount cannot be zero")
	}

	return nil
}

// GetSigners implements Msg
func (msg *MsgInitiateHyperchainBridgeTransaction) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

var _ sdk.Msg = &MsgApproveHyperchainBridgeTransaction{}

// ValidateBasic implements Msg
func (msg *MsgApproveHyperchainBridgeTransaction) ValidateBasic() error {
	// Validate relayer address
	_, err := sdk.AccAddressFromBech32(msg.Relayer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid relayer address (%s)", err)
	}

	// Validate bridge ID
	if msg.BridgeId == "" {
		return sdkerrors.Wrap(ErrBridgeNotFound, "bridge ID cannot be empty")
	}

	// Validate transaction ID
	if msg.TxId == "" {
		return sdkerrors.Wrap(ErrTransactionNotFound, "transaction ID cannot be empty")
	}

	return nil
}

// GetSigners implements Msg
func (msg *MsgApproveHyperchainBridgeTransaction) GetSigners() []sdk.AccAddress {
	relayer, err := sdk.AccAddressFromBech32(msg.Relayer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{relayer}
}

var _ sdk.Msg = &MsgSubmitHyperchainBlock{}

// ValidateBasic implements Msg
func (msg *MsgSubmitHyperchainBlock) ValidateBasic() error {
	// Validate proposer address
	_, err := sdk.AccAddressFromBech32(msg.Proposer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid proposer address (%s)", err)
	}

	// Validate chain ID
	if msg.ChainId == "" {
		return sdkerrors.Wrap(ErrHyperchainNotFound, "chain ID cannot be empty")
	}

	// Validate height
	if msg.Height == 0 {
		return sdkerrors.Wrap(ErrInvalidBlockHeight, "height cannot be zero")
	}

	return nil
}

// GetSigners implements Msg
func (msg *MsgSubmitHyperchainBlock) GetSigners() []sdk.AccAddress {
	proposer, err := sdk.AccAddressFromBech32(msg.Proposer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{proposer}
}

var _ sdk.Msg = &MsgSubmitHyperchainTransaction{}

// ValidateBasic implements Msg
func (msg *MsgSubmitHyperchainTransaction) ValidateBasic() error {
	// Validate sender address
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	// Validate recipient address
	_, err = sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid recipient address (%s)", err)
	}

	// Validate chain ID
	if msg.ChainId == "" {
		return sdkerrors.Wrap(ErrHyperchainNotFound, "chain ID cannot be empty")
	}

	// Validate amount
	if !msg.Amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}
	if msg.Amount.IsZero() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "amount cannot be zero")
	}

	return nil
}

// GetSigners implements Msg
func (msg *MsgSubmitHyperchainTransaction) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

var _ sdk.Msg = &MsgGrantHyperchainPermission{}

// ValidateBasic implements Msg
func (msg *MsgGrantHyperchainPermission) ValidateBasic() error {
	// Validate admin address
	_, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}

	// Validate address
	_, err = sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address (%s)", err)
	}

	// Validate chain ID
	if msg.ChainId == "" {
		return sdkerrors.Wrap(ErrHyperchainNotFound, "chain ID cannot be empty")
	}

	// Validate permission type
	if !IsValidPermissionType(msg.PermissionType) {
		return sdkerrors.Wrap(ErrInvalidPermissionType, msg.PermissionType)
	}

	// Validate expiration days
	if msg.ExpirationDays > 3650 { // Max 10 years
		return sdkerrors.Wrap(ErrInvalidExpirationDays, "expiration days too large")
	}

	return nil
}

// GetSigners implements Msg
func (msg *MsgGrantHyperchainPermission) GetSigners() []sdk.AccAddress {
	admin, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{admin}
}

var _ sdk.Msg = &MsgRevokeHyperchainPermission{}

// ValidateBasic implements Msg
func (msg *MsgRevokeHyperchainPermission) ValidateBasic() error {
	// Validate admin address
	_, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}

	// Validate address
	_, err = sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address (%s)", err)
	}

	// Validate chain ID
	if msg.ChainId == "" {
		return sdkerrors.Wrap(ErrHyperchainNotFound, "chain ID cannot be empty")
	}

	// Validate permission type
	if !IsValidPermissionType(msg.PermissionType) {
		return sdkerrors.Wrap(ErrInvalidPermissionType, msg.PermissionType)
	}

	return nil
}

// GetSigners implements Msg
func (msg *MsgRevokeHyperchainPermission) GetSigners() []sdk.AccAddress {
	admin, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{admin}
}