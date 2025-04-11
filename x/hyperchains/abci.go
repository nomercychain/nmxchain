package hyperchains

import (
	"encoding/json"

	"github.com/nomercychain/nmxchain/x/hyperchains/keeper"
	"github.com/nomercychain/nmxchain/x/hyperchains/types"
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeginBlocker is called at the beginning of every block
func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	// Process chain proposals
	k.ProcessChainProposals(ctx)
}

// EndBlocker is called at the end of every block
func EndBlocker(ctx sdk.Context, req abci.RequestEndBlock, k keeper.Keeper) []abci.ValidatorUpdate {
	// Process pending chain deployments
	processPendingDeployments(ctx, k)

	// Update chain metrics
	updateChainMetrics(ctx, k)

	// Return validator updates
	return []abci.ValidatorUpdate{}
}

// processPendingDeployments processes pending chain deployments
func processPendingDeployments(ctx sdk.Context, k keeper.Keeper) {
	deployments := k.GetAllChainDeployments(ctx)

	for _, deployment := range deployments {
		// Skip deployments that are not in progress
		if deployment.Status != "in_progress" && deployment.Status != "pending" {
			continue
		}

		// In a real implementation, this would check the status of the actual deployment
		// For now, we'll just simulate the deployment process

		// Update the deployment status
		if deployment.Status == "pending" {
			deployment.Status = "in_progress"
			k.UpdateChainDeployment(ctx, deployment.ID, "in_progress", "Deployment started", nil)
		} else {
			// Simulate a completed deployment
			endpoints := json.RawMessage(`{
				"rpc": "https://rpc.` + deployment.ChainID + `.nmxchain.io",
				"rest": "https://rest.` + deployment.ChainID + `.nmxchain.io",
				"explorer": "https://explorer.` + deployment.ChainID + `.nmxchain.io"
			}`)
			k.UpdateChainDeployment(ctx, deployment.ID, "completed", "Deployment completed successfully", endpoints)
		}
	}
}

// updateChainMetrics updates metrics for all active chains
func updateChainMetrics(ctx sdk.Context, k keeper.Keeper) {
	chains := k.GetAllChains(ctx)

	for _, chain := range chains {
		// Skip chains that are not active
		if chain.Status != types.ChainStatusActive {
			continue
		}

		// Get current metrics
		metrics, found := k.GetChainMetrics(ctx, chain.ID)
		if !found {
			// Initialize metrics for new chains
			metrics = types.ChainMetrics{
				ChainID:           chain.ID,
				BlockHeight:       0,
				TotalTransactions: 0,
				ActiveUsers:       0,
				TPS:               sdk.ZeroDec(),
				AverageFee:        sdk.ZeroDec(),
				TotalValueLocked:  sdk.NewCoins(),
				UpdatedAt:         ctx.BlockTime(),
			}
		}

		// In a real implementation, this would fetch actual metrics from the chain
		// For now, we'll just simulate some activity

		// Simulate block production (1 block per 2 seconds)
		timeSinceUpdate := ctx.BlockTime().Sub(metrics.UpdatedAt)
		newBlocks := int64(timeSinceUpdate.Seconds() / 2)
		if newBlocks > 0 {
			// Update block height
			metrics.BlockHeight += newBlocks

			// Simulate transactions (random number between 10-100 per block)
			newTx := uint64(newBlocks * (10 + ctx.BlockHeight()%91))
			metrics.TotalTransactions += newTx

			// Simulate TPS
			metrics.TPS = sdk.NewDec(int64(newTx)).Quo(sdk.NewDec(int64(timeSinceUpdate.Seconds())))

			// Simulate average fee (0.01 - 0.1 NMX)
			metrics.AverageFee = sdk.NewDecWithPrec(1+ctx.BlockHeight()%10, 2)

			// Update metrics
			metrics.UpdatedAt = ctx.BlockTime()
			k.SetChainMetrics(ctx, metrics)
		}
	}
}