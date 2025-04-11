package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateDynaContract{}, "dynacontract/CreateDynaContract", nil)
	cdc.RegisterConcrete(&MsgUpdateDynaContract{}, "dynacontract/UpdateDynaContract", nil)
	cdc.RegisterConcrete(&MsgExecuteDynaContract{}, "dynacontract/ExecuteDynaContract", nil)
	cdc.RegisterConcrete(&MsgAddLearningData{}, "dynacontract/AddLearningData", nil)
	cdc.RegisterConcrete(&MsgCreateDynaContractTemplate{}, "dynacontract/CreateDynaContractTemplate", nil)
	cdc.RegisterConcrete(&MsgInstantiateFromTemplate{}, "dynacontract/InstantiateFromTemplate", nil)
	cdc.RegisterConcrete(&MsgGrantContractPermission{}, "dynacontract/GrantContractPermission", nil)
	cdc.RegisterConcrete(&MsgRevokeContractPermission{}, "dynacontract/RevokeContractPermission", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateDynaContract{},
		&MsgUpdateDynaContract{},
		&MsgExecuteDynaContract{},
		&MsgAddLearningData{},
		&MsgCreateDynaContractTemplate{},
		&MsgInstantiateFromTemplate{},
		&MsgGrantContractPermission{},
		&MsgRevokeContractPermission{},
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