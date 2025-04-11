package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRegisterDataSource{}, "truthgpt/RegisterDataSource", nil)
	cdc.RegisterConcrete(&MsgUpdateDataSource{}, "truthgpt/UpdateDataSource", nil)
	cdc.RegisterConcrete(&MsgRemoveDataSource{}, "truthgpt/RemoveDataSource", nil)
	cdc.RegisterConcrete(&MsgCreateOracleQuery{}, "truthgpt/CreateOracleQuery", nil)
	cdc.RegisterConcrete(&MsgSubmitOracleResponse{}, "truthgpt/SubmitOracleResponse", nil)
	cdc.RegisterConcrete(&MsgRegisterAIModel{}, "truthgpt/RegisterAIModel", nil)
	cdc.RegisterConcrete(&MsgUpdateAIModel{}, "truthgpt/UpdateAIModel", nil)
	cdc.RegisterConcrete(&MsgRemoveAIModel{}, "truthgpt/RemoveAIModel", nil)
	cdc.RegisterConcrete(&MsgReportMisinformation{}, "truthgpt/ReportMisinformation", nil)
	cdc.RegisterConcrete(&MsgCreateVerificationTask{}, "truthgpt/CreateVerificationTask", nil)
	cdc.RegisterConcrete(&MsgCompleteVerificationTask{}, "truthgpt/CompleteVerificationTask", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterDataSource{},
		&MsgUpdateDataSource{},
		&MsgRemoveDataSource{},
		&MsgCreateOracleQuery{},
		&MsgSubmitOracleResponse{},
		&MsgRegisterAIModel{},
		&MsgUpdateAIModel{},
		&MsgRemoveAIModel{},
		&MsgReportMisinformation{},
		&MsgCreateVerificationTask{},
		&MsgCompleteVerificationTask{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)