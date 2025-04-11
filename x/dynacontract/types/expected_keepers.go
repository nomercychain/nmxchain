package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	deaitypes "github.com/nomercychain/nmxchain/x/deai/types"
)

// BankKeeper defines the expected bank keeper
type BankKeeper interface {
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToAccount(ctx sdk.Context, senderAddr sdk.AccAddress, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
}

// DeAIKeeper defines the expected DeAI keeper
type DeAIKeeper interface {
	GetAIAgent(ctx sdk.Context, id string) (agent deaitypes.AIAgent, found bool)
	GetAIAgentState(ctx sdk.Context, agentID string) (state deaitypes.AIAgentState, found bool)
	SetAIAgentState(ctx sdk.Context, state deaitypes.AIAgentState)
	ExecuteAIAgent(ctx sdk.Context, agentID string, sender sdk.AccAddress, actionType string, data []byte, fee sdk.Coin) (actionID string, result []byte, err error)
}