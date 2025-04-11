package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateHyperchain{}, "hyperchain/CreateHyperchain", nil)
	cdc.RegisterConcrete(&MsgUpdateHyperchain{}, "hyperchain/UpdateHyperchain", nil)
	cdc.RegisterConcrete(&MsgJoinHyperchainAsValidator{}, "hyperchain/JoinHyperchainAsValidator", nil)
	cdc.RegisterConcrete(&MsgLeaveHyperchain{}, "hyperchain/LeaveHyperchain", nil)
	cdc.RegisterConcrete(&MsgCreateHyperchainBridge{}, "hyperchain/CreateHyperchainBridge", nil)
	cdc.RegisterConcrete(&MsgUpdateHyperchainBridge{}, "hyperchain/UpdateHyperchainBridge", nil)
	cdc.RegisterConcrete(&MsgRegisterHyperchainBridgeRelayer{}, "hyperchain/RegisterHyperchainBridgeRelayer", nil)
	cdc.RegisterConcrete(&MsgRemoveHyperchainBridgeRelayer{}, "hyperchain/RemoveHyperchainBridgeRelayer", nil)
	cdc.RegisterConcrete(&MsgInitiateHyperchainBridgeTransaction{}, "hyperchain/InitiateHyperchainBridgeTransaction", nil)
	cdc.RegisterConcrete(&MsgApproveHyperchainBridgeTransaction{}, "hyperchain/ApproveHyperchainBridgeTransaction", nil)
	cdc.RegisterConcrete(&MsgSubmitHyperchainBlock{}, "hyperchain/SubmitHyperchainBlock", nil)
	cdc.RegisterConcrete(&MsgSubmitHyperchainTransaction{}, "hyperchain/SubmitHyperchainTransaction", nil)
	cdc.RegisterConcrete(&MsgGrantHyperchainPermission{}, "hyperchain/GrantHyperchainPermission", nil)
	cdc.RegisterConcrete(&MsgRevokeHyperchainPermission{}, "hyperchain/RevokeHyperchainPermission", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateHyperchain{},
		&MsgUpdateHyperchain{},
		&MsgJoinHyperchainAsValidator{},
		&MsgLeaveHyperchain{},
		&MsgCreateHyperchainBridge{},
		&MsgUpdateHyperchainBridge{},
		&MsgRegisterHyperchainBridgeRelayer{},
		&MsgRemoveHyperchainBridgeRelayer{},
		&MsgInitiateHyperchainBridgeTransaction{},
		&MsgApproveHyperchainBridgeTransaction{},
		&MsgSubmitHyperchainBlock{},
		&MsgSubmitHyperchainTransaction{},
		&MsgGrantHyperchainPermission{},
		&MsgRevokeHyperchainPermission{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

func init() {
	RegisterCodec(Amino)
	sdk.RegisterLegacyAminoCodec(Amino)
}