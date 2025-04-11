package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateValidator{}, "neuropos/CreateValidator", nil)
	cdc.RegisterConcrete(&MsgEditValidator{}, "neuropos/EditValidator", nil)
	cdc.RegisterConcrete(&MsgDelegate{}, "neuropos/Delegate", nil)
	cdc.RegisterConcrete(&MsgUndelegate{}, "neuropos/Undelegate", nil)
	cdc.RegisterConcrete(&MsgBeginRedelegate{}, "neuropos/BeginRedelegate", nil)
	cdc.RegisterConcrete(&MsgUnjail{}, "neuropos/Unjail", nil)
	cdc.RegisterConcrete(&MsgUpdateNeuralNetwork{}, "neuropos/UpdateNeuralNetwork", nil)
	cdc.RegisterConcrete(&MsgTrainNeuralNetwork{}, "neuropos/TrainNeuralNetwork", nil)
	cdc.RegisterConcrete(&MsgSubmitNeuralPrediction{}, "neuropos/SubmitNeuralPrediction", nil)
	cdc.RegisterConcrete(&MsgUpdateValidatorReputation{}, "neuropos/UpdateValidatorReputation", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateValidator{},
		&MsgEditValidator{},
		&MsgDelegate{},
		&MsgUndelegate{},
		&MsgBeginRedelegate{},
		&MsgUnjail{},
		&MsgUpdateNeuralNetwork{},
		&MsgTrainNeuralNetwork{},
		&MsgSubmitNeuralPrediction{},
		&MsgUpdateValidatorReputation{},
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