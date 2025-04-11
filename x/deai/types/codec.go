package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterLegacyAminoCodec registers the necessary interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateAIAgent{}, "deai/CreateAIAgent", nil)
	cdc.RegisterConcrete(&MsgUpdateAIAgent{}, "deai/UpdateAIAgent", nil)
	cdc.RegisterConcrete(&MsgTrainAIAgent{}, "deai/TrainAIAgent", nil)
	cdc.RegisterConcrete(&MsgExecuteAIAgent{}, "deai/ExecuteAIAgent", nil)
	cdc.RegisterConcrete(&MsgListAIAgentForSale{}, "deai/ListAIAgentForSale", nil)
	cdc.RegisterConcrete(&MsgBuyAIAgent{}, "deai/BuyAIAgent", nil)
	cdc.RegisterConcrete(&MsgRentAIAgent{}, "deai/RentAIAgent", nil)
	cdc.RegisterConcrete(&MsgCancelMarketListing{}, "deai/CancelMarketListing", nil)
}

// RegisterInterfaces registers the interfaces and implementations
func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateAIAgent{},
		&MsgUpdateAIAgent{},
		&MsgTrainAIAgent{},
		&MsgExecuteAIAgent{},
		&MsgListAIAgentForSale{},
		&MsgBuyAIAgent{},
		&MsgRentAIAgent{},
		&MsgCancelMarketListing{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

// ModuleCdc is the codec for the module
var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
}