package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nomercychain/nmxchain/x/truthgpt/types"
)

// RegisterInvariants registers all truthgpt invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, "module-accounts",
		ModuleAccountInvariant(k))
	ir.RegisterRoute(types.ModuleName, "oracle-providers",
		OracleProviderInvariant(k))
	ir.RegisterRoute(types.ModuleName, "oracle-requests",
		OracleRequestInvariant(k))
	ir.RegisterRoute(types.ModuleName, "oracle-responses",
		OracleResponseInvariant(k))
	ir.RegisterRoute(types.ModuleName, "provider-reputation",
		ProviderReputationInvariant(k))
}

// ModuleAccountInvariant checks that the module account's balance matches the sum of all staked tokens
func ModuleAccountInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)
		moduleBalance := k.bankKeeper.GetAllBalances(ctx, moduleAddr)

		// Calculate total staked tokens
		totalStaked := sdk.ZeroInt()
		providers := k.GetAllOracleProviders(ctx)
		for _, provider := range providers {
			totalStaked = totalStaked.Add(provider.StakedAmount.Amount)
		}

		// Add pending rewards
		pendingRewards := sdk.ZeroInt()
		requests := k.GetAllOracleRequests(ctx)
		for _, request := range requests {
			if request.Status == types.RequestStatus_PENDING || request.Status == types.RequestStatus_ACTIVE {
				pendingRewards = pendingRewards.Add(request.Reward.Amount)
			}
		}

		totalExpected := totalStaked.Add(pendingRewards)
		expectedBalance := sdk.NewCoins(sdk.NewCoin("unomx", totalExpected))

		if !moduleBalance.IsEqual(expectedBalance) {
			return fmt.Sprintf("module account balance (%s) does not match total expected tokens (%s)", moduleBalance, expectedBalance), true
		}

		return "", false
	}
}

// OracleProviderInvariant checks that oracle providers have valid parameters
func OracleProviderInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		providers := k.GetAllOracleProviders(ctx)
		for _, provider := range providers {
			// Check that staked amount is positive
			if provider.StakedAmount.IsNegative() || provider.StakedAmount.IsZero() {
				return fmt.Sprintf("oracle provider %s has invalid staked amount: %s", provider.Address, provider.StakedAmount), true
			}

			// Check that reputation is between 0 and 1
			if provider.Reputation.IsNegative() || provider.Reputation.GT(sdk.OneDec()) {
				return fmt.Sprintf("oracle provider %s has invalid reputation: %s", provider.Address, provider.Reputation), true
			}

			// Check that success rate is between 0 and 1
			if provider.SuccessRate.IsNegative() || provider.SuccessRate.GT(sdk.OneDec()) {
				return fmt.Sprintf("oracle provider %s has invalid success rate: %s", provider.Address, provider.SuccessRate), true
			}
		}

		return "", false
	}
}

// OracleRequestInvariant checks that oracle requests have valid parameters
func OracleRequestInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		requests := k.GetAllOracleRequests(ctx)
		for _, request := range requests {
			// Check that reward is positive
			if request.Reward.IsNegative() || request.Reward.IsZero() {
				return fmt.Sprintf("oracle request %s has invalid reward: %s", request.Id, request.Reward), true
			}

			// Check that timeout height is greater than current height for active requests
			if request.Status == types.RequestStatus_ACTIVE && request.TimeoutHeight <= ctx.BlockHeight() {
				return fmt.Sprintf("oracle request %s has invalid timeout height: %d (current: %d)", request.Id, request.TimeoutHeight, ctx.BlockHeight()), true
			}

			// Check that min responses is positive and not greater than max responses
			if request.MinResponses <= 0 || request.MinResponses > request.MaxResponses {
				return fmt.Sprintf("oracle request %s has invalid min/max responses: %d/%d", request.Id, request.MinResponses, request.MaxResponses), true
			}
		}

		return "", false
	}
}

// OracleResponseInvariant checks that oracle responses have valid parameters
func OracleResponseInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		responses := k.GetAllOracleResponses(ctx)
		for _, response := range responses {
			// Check that request exists for this response
			request, found := k.GetOracleRequest(ctx, response.RequestId)
			if !found {
				return fmt.Sprintf("oracle response %s references non-existent request: %s", response.Id, response.RequestId), true
			}

			// Check that provider exists for this response
			provider, found := k.GetOracleProvider(ctx, response.ProviderAddress)
			if !found {
				return fmt.Sprintf("oracle response %s references non-existent provider: %s", response.Id, response.ProviderAddress), true
			}

			// Check that confidence is between 0 and 1
			if response.Confidence.IsNegative() || response.Confidence.GT(sdk.OneDec()) {
				return fmt.Sprintf("oracle response %s has invalid confidence: %s", response.Id, response.Confidence), true
			}
		}

		return "", false
	}
}

// ProviderReputationInvariant checks that provider reputations are consistent
func ProviderReputationInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		providers := k.GetAllOracleProviders(ctx)
		for _, provider := range providers {
			// Calculate expected reputation based on history
			history := k.GetProviderResponseHistory(ctx, provider.Address)
			if len(history.Responses) == 0 {
				// New provider with no history should have default reputation
				if !provider.Reputation.Equal(sdk.NewDecWithPrec(5, 1)) { // 0.5 default
					return fmt.Sprintf("new oracle provider %s has non-default reputation: %s", provider.Address, provider.Reputation), true
				}
				continue
			}

			// Check that success rate matches history
			totalResponses := len(history.Responses)
			successfulResponses := 0
			for _, responseID := range history.Responses {
				response, found := k.GetOracleResponse(ctx, responseID)
				if !found {
					continue
				}

				request, found := k.GetOracleRequest(ctx, response.RequestId)
				if !found {
					continue
				}

				if request.Status == types.RequestStatus_SUCCESSFUL && response.Accepted {
					successfulResponses++
				}
			}

			expectedSuccessRate := sdk.NewDec(int64(successfulResponses)).QuoInt64(int64(totalResponses))
			if !provider.SuccessRate.Equal(expectedSuccessRate) {
				return fmt.Sprintf("oracle provider %s has inconsistent success rate: %s (expected: %s)", provider.Address, provider.SuccessRate, expectedSuccessRate), true
			}
		}

		return "", false
	}
}
