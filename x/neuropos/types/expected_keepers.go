package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	deaitypes "github.com/nomercychain/nmxchain/x/deai/types"
)

// AccountKeeper defines the expected account keeper
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
	GetModuleAccount(ctx sdk.Context, moduleName string) authtypes.ModuleAccountI
	SetModuleAccount(ctx sdk.Context, macc authtypes.ModuleAccountI)
	NewAccount(ctx sdk.Context, acc authtypes.AccountI) authtypes.AccountI
	SetAccount(ctx sdk.Context, acc authtypes.AccountI)
}

// BankKeeper defines the expected bank keeper
type BankKeeper interface {
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	SetBalances(ctx sdk.Context, addr sdk.AccAddress, balances sdk.Coins) error
	IterateAllBalances(ctx sdk.Context, cb func(address sdk.AccAddress, coin sdk.Coin) bool)
}

// StakingKeeper defines the expected staking keeper
type StakingKeeper interface {
	GetValidator(ctx sdk.Context, addr sdk.ValAddress) (validator stakingtypes.Validator, found bool)
	GetValidatorByConsAddr(ctx sdk.Context, consAddr sdk.ConsAddress) (validator stakingtypes.Validator, found bool)
	GetValidatorDelegations(ctx sdk.Context, valAddr sdk.ValAddress) (delegations []stakingtypes.Delegation)
	GetDelegation(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) (delegation stakingtypes.Delegation, found bool)
	GetUnbondingDelegation(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) (ubd stakingtypes.UnbondingDelegation, found bool)
	GetValidators(ctx sdk.Context, maxRetrieve uint32) (validators []stakingtypes.Validator)
	GetAllValidators(ctx sdk.Context) (validators []stakingtypes.Validator)
	GetLastTotalPower(ctx sdk.Context) sdk.Int
	GetLastValidatorPower(ctx sdk.Context, valAddr sdk.ValAddress) int64
	IterateValidators(ctx sdk.Context, cb func(index int64, validator stakingtypes.ValidatorI) bool)
	IterateDelegations(ctx sdk.Context, delegator sdk.AccAddress, cb func(index int64, delegation stakingtypes.DelegationI) bool)
	SetValidator(ctx sdk.Context, validator stakingtypes.Validator)
	RemoveValidator(ctx sdk.Context, addr sdk.ValAddress)
	JailValidator(ctx sdk.Context, consAddr sdk.ConsAddress)
	UnjailValidator(ctx sdk.Context, consAddr sdk.ConsAddress)
	TombstoneValidator(ctx sdk.Context, consAddr sdk.ConsAddress)
	IsTombstoned(ctx sdk.Context, consAddr sdk.ConsAddress) bool
	SlashValidator(ctx sdk.Context, consAddr sdk.ConsAddress, infractionHeight int64, power int64, slashFactor sdk.Dec)
	GetParams(ctx sdk.Context) stakingtypes.Params
	SetParams(ctx sdk.Context, params stakingtypes.Params)
}

// SlashingKeeper defines the expected slashing keeper
type SlashingKeeper interface {
	GetValidatorSigningInfo(ctx sdk.Context, consAddr sdk.ConsAddress) (info slashingtypes.ValidatorSigningInfo, found bool)
	SetValidatorSigningInfo(ctx sdk.Context, consAddr sdk.ConsAddress, info slashingtypes.ValidatorSigningInfo)
	JailUntil(ctx sdk.Context, consAddr sdk.ConsAddress, jailTime time.Time)
	Tombstone(ctx sdk.Context, consAddr sdk.ConsAddress)
	IsTombstoned(ctx sdk.Context, consAddr sdk.ConsAddress) bool
	GetValidatorMissedBlocks(ctx sdk.Context, consAddr sdk.ConsAddress) (missedBlocks []slashingtypes.MissedBlock)
	SetValidatorMissedBlocks(ctx sdk.Context, consAddr sdk.ConsAddress, missedBlocks []slashingtypes.MissedBlock)
	Slash(ctx sdk.Context, consAddr sdk.ConsAddress, infractionHeight int64, power int64, slashFactor sdk.Dec)
	SlashWithInfractionReason(ctx sdk.Context, consAddr sdk.ConsAddress, infractionHeight int64, power int64, slashFactor sdk.Dec, reason string)
	GetParams(ctx sdk.Context) slashingtypes.Params
	SetParams(ctx sdk.Context, params slashingtypes.Params)
}

// DeAIKeeper defines the expected DeAI keeper
type DeAIKeeper interface {
	GetAIAgent(ctx sdk.Context, id string) (agent deaitypes.AIAgent, found bool)
	GetAIAgentState(ctx sdk.Context, agentID string) (state deaitypes.AIAgentState, found bool)
	SetAIAgentState(ctx sdk.Context, state deaitypes.AIAgentState)
	ExecuteAIAgent(ctx sdk.Context, agentID string, sender sdk.AccAddress, actionType string, data []byte, fee sdk.Coin) (actionID string, result []byte, err error)
}