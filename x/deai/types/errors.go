package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// DeAI module sentinel errors
var (
	ErrAgentNotFound          = sdkerrors.Register(ModuleName, 1, "AI agent not found")
	ErrAgentAlreadyExists     = sdkerrors.Register(ModuleName, 2, "AI agent already exists")
	ErrUnauthorized           = sdkerrors.Register(ModuleName, 3, "unauthorized")
	ErrInvalidAgentStatus     = sdkerrors.Register(ModuleName, 4, "invalid agent status")
	ErrInvalidAgentType       = sdkerrors.Register(ModuleName, 5, "invalid agent type")
	ErrInvalidModelID         = sdkerrors.Register(ModuleName, 6, "invalid model ID")
	ErrModelNotFound          = sdkerrors.Register(ModuleName, 7, "AI model not found")
	ErrInvalidTrainingData    = sdkerrors.Register(ModuleName, 8, "invalid training data")
	ErrTrainingDataTooLarge   = sdkerrors.Register(ModuleName, 9, "training data too large")
	ErrAgentStateNotFound     = sdkerrors.Register(ModuleName, 10, "agent state not found")
	ErrActionNotFound         = sdkerrors.Register(ModuleName, 11, "action not found")
	ErrInvalidActionType      = sdkerrors.Register(ModuleName, 12, "invalid action type")
	ErrInsufficientFee        = sdkerrors.Register(ModuleName, 13, "insufficient fee")
	ErrListingNotFound        = sdkerrors.Register(ModuleName, 14, "marketplace listing not found")
	ErrListingNotActive       = sdkerrors.Register(ModuleName, 15, "marketplace listing not active")
	ErrInvalidListingType     = sdkerrors.Register(ModuleName, 16, "invalid listing type")
	ErrAgentNotForSale        = sdkerrors.Register(ModuleName, 17, "agent not for sale")
	ErrAgentNotForRent        = sdkerrors.Register(ModuleName, 18, "agent not for rent")
	ErrInvalidRentalDuration  = sdkerrors.Register(ModuleName, 19, "invalid rental duration")
	ErrTooManyListings        = sdkerrors.Register(ModuleName, 20, "too many marketplace listings")
	ErrInvalidAgentName       = sdkerrors.Register(ModuleName, 21, "invalid agent name")
	ErrInvalidAgentDescription = sdkerrors.Register(ModuleName, 22, "invalid agent description")
	ErrInvalidDeposit         = sdkerrors.Register(ModuleName, 23, "invalid deposit")
	ErrInsufficientDeposit    = sdkerrors.Register(ModuleName, 24, "insufficient deposit")
)