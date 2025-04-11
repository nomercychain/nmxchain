package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgCreateContract{}
	_ sdk.Msg = &MsgUpdateContract{}
	_ sdk.Msg = &MsgDeleteContract{}
	_ sdk.Msg = &MsgExecuteContract{}
	_ sdk.Msg = &MsgUpgradeContract{}
)

// Message types for the dynacontract module
const (
	TypeMsgCreateContract  = "create_contract"
	TypeMsgUpdateContract  = "update_contract"
	TypeMsgDeleteContract  = "delete_contract"
	TypeMsgExecuteContract = "execute_contract"
	TypeMsgUpgradeContract = "upgrade_contract"
)

// NewMsgCreateContract creates a new MsgCreateContract instance
func NewMsgCreateContract(creator, name, version, code, initMsg string) *MsgCreateContract {
	return &MsgCreateContract{
		Creator: creator,
		Name:    name,
		Version: version,
		Code:    code,
		InitMsg: initMsg,
	}
}

// Route implements sdk.Msg
func (msg MsgCreateContract) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (msg MsgCreateContract) Type() string {
	return TypeMsgCreateContract
}

// GetSigners implements sdk.Msg
func (msg MsgCreateContract) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes implements sdk.Msg
func (msg MsgCreateContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements sdk.Msg
func (msg MsgCreateContract) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.Name == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "name cannot be empty")
	}

	if msg.Version == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "version cannot be empty")
	}

	if msg.Code == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "code cannot be empty")
	}

	return nil
}

// NewMsgUpdateContract creates a new MsgUpdateContract instance
func NewMsgUpdateContract(creator, id, name, version, code, updateMsg string) *MsgUpdateContract {
	return &MsgUpdateContract{
		Creator:   creator,
		Id:        id,
		Name:      name,
		Version:   version,
		Code:      code,
		UpdateMsg: updateMsg,
	}
}

// Route implements sdk.Msg
func (msg MsgUpdateContract) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (msg MsgUpdateContract) Type() string {
	return TypeMsgUpdateContract
}

// GetSigners implements sdk.Msg
func (msg MsgUpdateContract) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes implements sdk.Msg
func (msg MsgUpdateContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements sdk.Msg
func (msg MsgUpdateContract) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.Id == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "id cannot be empty")
	}

	if msg.Name == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "name cannot be empty")
	}

	if msg.Version == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "version cannot be empty")
	}

	if msg.Code == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "code cannot be empty")
	}

	return nil
}

// NewMsgDeleteContract creates a new MsgDeleteContract instance
func NewMsgDeleteContract(creator, id string) *MsgDeleteContract {
	return &MsgDeleteContract{
		Creator: creator,
		Id:      id,
	}
}

// Route implements sdk.Msg
func (msg MsgDeleteContract) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (msg MsgDeleteContract) Type() string {
	return TypeMsgDeleteContract
}

// GetSigners implements sdk.Msg
func (msg MsgDeleteContract) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes implements sdk.Msg
func (msg MsgDeleteContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements sdk.Msg
func (msg MsgDeleteContract) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.Id == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "id cannot be empty")
	}

	return nil
}

// NewMsgExecuteContract creates a new MsgExecuteContract instance
func NewMsgExecuteContract(sender, id, executeMsg string, coins sdk.Coins) *MsgExecuteContract {
	return &MsgExecuteContract{
		Sender:     sender,
		Id:         id,
		ExecuteMsg: executeMsg,
		Coins:      coins,
	}
}

// Route implements sdk.Msg
func (msg MsgExecuteContract) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (msg MsgExecuteContract) Type() string {
	return TypeMsgExecuteContract
}

// GetSigners implements sdk.Msg
func (msg MsgExecuteContract) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// GetSignBytes implements sdk.Msg
func (msg MsgExecuteContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements sdk.Msg
func (msg MsgExecuteContract) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	if msg.Id == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "id cannot be empty")
	}

	if msg.ExecuteMsg == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "execute message cannot be empty")
	}

	if !msg.Coins.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Coins.String())
	}

	return nil
}

// NewMsgUpgradeContract creates a new MsgUpgradeContract instance
func NewMsgUpgradeContract(creator, id, version, code, migrateMsg string) *MsgUpgradeContract {
	return &MsgUpgradeContract{
		Creator:    creator,
		Id:         id,
		Version:    version,
		Code:       code,
		MigrateMsg: migrateMsg,
	}
}

// Route implements sdk.Msg
func (msg MsgUpgradeContract) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (msg MsgUpgradeContract) Type() string {
	return TypeMsgUpgradeContract
}

// GetSigners implements sdk.Msg
func (msg MsgUpgradeContract) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes implements sdk.Msg
func (msg MsgUpgradeContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements sdk.Msg
func (msg MsgUpgradeContract) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.Id == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "id cannot be empty")
	}

	if msg.Version == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "version cannot be empty")
	}

	if msg.Code == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "code cannot be empty")
	}

	return nil
}