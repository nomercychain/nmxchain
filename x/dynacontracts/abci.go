package dynacontracts

import (
	"github.com/nomercychain/nmxchain/x/dynacontracts/keeper"
	"github.com/nomercychain/nmxchain/x/dynacontracts/types"
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeginBlocker is called at the beginning of every block
func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	// Process any pending contract proposals
	k.ProcessContractProposals(ctx)
}

// EndBlocker is called at the end of every block
func EndBlocker(ctx sdk.Context, req abci.RequestEndBlock, k keeper.Keeper) []abci.ValidatorUpdate {
	// Update contract parameters using AI
	k.UpdateContractParameters(ctx)

	// Return validator updates
	return []abci.ValidatorUpdate{}
}