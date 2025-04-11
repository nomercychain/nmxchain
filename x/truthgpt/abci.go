package truthgpt

import (
	"github.com/nomercychain/nmxchain/x/truthgpt/keeper"
	"github.com/nomercychain/nmxchain/x/truthgpt/types"
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeginBlocker is called at the beginning of every block
func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	// Process pending oracle queries
	k.ProcessPendingQueries(ctx)
}

// EndBlocker is called at the end of every block
func EndBlocker(ctx sdk.Context, req abci.RequestEndBlock, k keeper.Keeper) []abci.ValidatorUpdate {
	// Process verification tasks
	k.ProcessVerificationTasks(ctx)

	// Update data source rankings periodically
	if ctx.BlockHeight()%100 == 0 { // Every 100 blocks
		k.UpdateDataSourceRankings(ctx)
	}

	// Return validator updates
	return []abci.ValidatorUpdate{}
}