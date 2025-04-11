package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nomercychain/nmxchain/x/deai/types"
)

// RegisterInvariants registers all deai invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, "valid-agent-states", ValidAgentStatesInvariant(k))
	ir.RegisterRoute(types.ModuleName, "valid-marketplace-listings", ValidMarketplaceListingsInvariant(k))
}

// ValidAgentStatesInvariant checks that all AI agents have valid states
func ValidAgentStatesInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var invalidAgents []string
		var msg string
		var broken bool

		agents := k.GetAllAIAgents(ctx)
		for _, agent := range agents {
			// Check if agent has a state
			state, found := k.GetAIAgentState(ctx, agent.ID)
			if !found {
				invalidAgents = append(invalidAgents, agent.ID)
				broken = true
				continue
			}

			// Check if state is valid
			if state.AgentID != agent.ID {
				invalidAgents = append(invalidAgents, agent.ID)
				broken = true
			}
		}

		if broken {
			msg = fmt.Sprintf("found %d agents with invalid states: %v", len(invalidAgents), invalidAgents)
		}

		return sdk.FormatInvariant(
			types.ModuleName,
			"valid-agent-states",
			msg,
		), broken
	}
}

// ValidMarketplaceListingsInvariant checks that all marketplace listings are valid
func ValidMarketplaceListingsInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var invalidListings []string
		var msg string
		var broken bool

		listings := k.GetAllAIAgentMarketplaceListings(ctx)
		for _, listing := range listings {
			// Check if listing references a valid agent
			agent, found := k.GetAIAgent(ctx, listing.AgentID)
			if !found {
				invalidListings = append(invalidListings, listing.ID)
				broken = true
				continue
			}

			// Check if active listing has correct agent status
			if listing.Status == "active" {
				if listing.ListingType == "sale" && agent.Status != types.AIAgentStatusForSale {
					invalidListings = append(invalidListings, listing.ID)
					broken = true
				} else if listing.ListingType == "rent" && agent.Status != types.AIAgentStatusForRent {
					invalidListings = append(invalidListings, listing.ID)
					broken = true
				} else if listing.ListingType == "both" && agent.Status != types.AIAgentStatusForSale && agent.Status != types.AIAgentStatusForRent {
					invalidListings = append(invalidListings, listing.ID)
					broken = true
				}
			}

			// Check if listing has expired
			if listing.Status == "active" && !listing.ExpiresAt.IsZero() && listing.ExpiresAt.Before(ctx.BlockTime()) {
				invalidListings = append(invalidListings, listing.ID)
				broken = true
			}
		}

		if broken {
			msg = fmt.Sprintf("found %d invalid marketplace listings: %v", len(invalidListings), invalidListings)
		}

		return sdk.FormatInvariant(
			types.ModuleName,
			"valid-marketplace-listings",
			msg,
		), broken
	}
}