package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nomercychain/nmxchain/x/deai/types"
)

// GetParams gets all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.Params{
		MinAgentDeposit:        k.MinAgentDeposit(ctx),
		MaxAgentNameLength:     k.MaxAgentNameLength(ctx),
		MaxAgentDescLength:     k.MaxAgentDescLength(ctx),
		MaxTrainingDataSize:    k.MaxTrainingDataSize(ctx),
		MaxMarketplaceListings: k.MaxMarketplaceListings(ctx),
		MarketplaceFeeRate:     k.MarketplaceFeeRate(ctx),
	}
}

// SetParams sets the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// MinAgentDeposit returns the MinAgentDeposit param
func (k Keeper) MinAgentDeposit(ctx sdk.Context) (res sdk.Coin) {
	k.paramstore.Get(ctx, types.KeyMinAgentDeposit, &res)
	return
}

// MaxAgentNameLength returns the MaxAgentNameLength param
func (k Keeper) MaxAgentNameLength(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyMaxAgentNameLength, &res)
	return
}

// MaxAgentDescLength returns the MaxAgentDescLength param
func (k Keeper) MaxAgentDescLength(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyMaxAgentDescLength, &res)
	return
}

// MaxTrainingDataSize returns the MaxTrainingDataSize param
func (k Keeper) MaxTrainingDataSize(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyMaxTrainingDataSize, &res)
	return
}

// MaxMarketplaceListings returns the MaxMarketplaceListings param
func (k Keeper) MaxMarketplaceListings(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyMaxMarketplaceListings, &res)
	return
}

// MarketplaceFeeRate returns the MarketplaceFeeRate param
func (k Keeper) MarketplaceFeeRate(ctx sdk.Context) (res sdk.Dec) {
	k.paramstore.Get(ctx, types.KeyMarketplaceFeeRate, &res)
	return
}