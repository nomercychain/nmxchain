package types

import (
	context "context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// MsgServer defines the MsgServer interface for the deai module
type MsgServer interface {
	CreateAIAgent(context.Context, *MsgCreateAIAgent) (*MsgCreateAIAgentResponse, error)
	UpdateAIAgent(context.Context, *MsgUpdateAIAgent) (*MsgUpdateAIAgentResponse, error)
	TrainAIAgent(context.Context, *MsgTrainAIAgent) (*MsgTrainAIAgentResponse, error)
	ExecuteAIAgent(context.Context, *MsgExecuteAIAgent) (*MsgExecuteAIAgentResponse, error)
	ListAIAgentForSale(context.Context, *MsgListAIAgentForSale) (*MsgListAIAgentForSaleResponse, error)
	BuyAIAgent(context.Context, *MsgBuyAIAgent) (*MsgBuyAIAgentResponse, error)
	RentAIAgent(context.Context, *MsgRentAIAgent) (*MsgRentAIAgentResponse, error)
	CancelMarketListing(context.Context, *MsgCancelMarketListing) (*MsgCancelMarketListingResponse, error)
}

// QueryServer defines the QueryServer interface for the deai module
type QueryServer interface {
	AIAgent(context.Context, *QueryAIAgentRequest) (*QueryAIAgentResponse, error)
	AIAgents(context.Context, *QueryAIAgentsRequest) (*QueryAIAgentsResponse, error)
	AIAgentState(context.Context, *QueryAIAgentStateRequest) (*QueryAIAgentStateResponse, error)
	AIAgentActions(context.Context, *QueryAIAgentActionsRequest) (*QueryAIAgentActionsResponse, error)
	AIAgentAction(context.Context, *QueryAIAgentActionRequest) (*QueryAIAgentActionResponse, error)
	AIAgentModels(context.Context, *QueryAIAgentModelsRequest) (*QueryAIAgentModelsResponse, error)
	AIAgentModel(context.Context, *QueryAIAgentModelRequest) (*QueryAIAgentModelResponse, error)
	AIAgentTrainingData(context.Context, *QueryAIAgentTrainingDataRequest) (*QueryAIAgentTrainingDataResponse, error)
	MarketplaceListings(context.Context, *QueryMarketplaceListingsRequest) (*QueryMarketplaceListingsResponse, error)
	MarketplaceListing(context.Context, *QueryMarketplaceListingRequest) (*QueryMarketplaceListingResponse, error)
}

// AccountKeeper defines the expected account keeper
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// other methods from the interface you are implementing
}

// BankKeeper defines the expected bank keeper
type BankKeeper interface {
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoinsFromAccountToAccount(ctx sdk.Context, senderAddr sdk.AccAddress, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	// other methods from the interface you are implementing
}

// NFTKeeper defines the expected NFT keeper
type NFTKeeper interface {
	// Methods for NFT operations
}