package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nomercychain/nmxchain/x/neuropos/types"
)

// RegisterInvariants registers all neuropos invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, "module-accounts",
		ModuleAccountInvariant(k))
	ir.RegisterRoute(types.ModuleName, "delegator-shares",
		DelegatorSharesInvariant(k))
	ir.RegisterRoute(types.ModuleName, "validator-tokens",
		ValidatorTokensInvariant(k))
	ir.RegisterRoute(types.ModuleName, "validator-reputation",
		ValidatorReputationInvariant(k))
}

// ModuleAccountInvariant checks that the module account's balance matches the sum of all delegations
func ModuleAccountInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)
		moduleBalance := k.bankKeeper.GetAllBalances(ctx, moduleAddr)

		// Calculate total delegated tokens
		totalDelegated := sdk.ZeroInt()
		delegations := k.GetAllDelegations(ctx)
		for _, del := range delegations {
			val, found := k.GetValidator(ctx, del.ValidatorAddress)
			if !found {
				continue
			}

			totalDelegated = totalDelegated.Add(val.Tokens.Mul(del.Shares.TruncateInt()).Quo(val.DelegatorShares.TruncateInt()))
		}

		// Add unbonding delegations
		ubds := k.GetAllUnbondingDelegations(ctx)
		for _, ubd := range ubds {
			for _, entry := range ubd.Entries {
				totalDelegated = totalDelegated.Add(entry.Balance)
			}
		}

		// Check if module balance matches total delegated tokens
		expectedBalance := sdk.NewCoins(sdk.NewCoin("unomx", totalDelegated))
		if !moduleBalance.IsEqual(expectedBalance) {
			return fmt.Sprintf("module account balance (%s) does not match total delegated tokens (%s)", moduleBalance, expectedBalance), true
		}

		return "", false
	}
}

// DelegatorSharesInvariant checks that the sum of delegator shares equals validator's delegator shares
func DelegatorSharesInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		validators := k.GetAllValidators(ctx)
		for _, val := range validators {
			totalShares := sdk.ZeroDec()
			delegations := k.GetAllDelegations(ctx)
			for _, del := range delegations {
				if del.ValidatorAddress == val.OperatorAddress {
					totalShares = totalShares.Add(del.Shares)
				}
			}

			if !totalShares.Equal(val.DelegatorShares) {
				return fmt.Sprintf("validator %s delegator shares (%s) does not match sum of delegations (%s)", val.OperatorAddress, val.DelegatorShares, totalShares), true
			}
		}

		return "", false
	}
}

// ValidatorTokensInvariant checks that validator tokens are valid
func ValidatorTokensInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		validators := k.GetAllValidators(ctx)
		for _, val := range validators {
			if val.Tokens.IsNegative() {
				return fmt.Sprintf("validator %s has negative tokens: %s", val.OperatorAddress, val.Tokens), true
			}

			if val.DelegatorShares.IsNegative() {
				return fmt.Sprintf("validator %s has negative delegator shares: %s", val.OperatorAddress, val.DelegatorShares), true
			}

			if val.Tokens.IsZero() && !val.DelegatorShares.IsZero() {
				return fmt.Sprintf("validator %s has zero tokens but non-zero delegator shares: %s", val.OperatorAddress, val.DelegatorShares), true
			}
		}

		return "", false
	}
}

// ValidatorReputationInvariant checks that validator reputations are valid
func ValidatorReputationInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		reputations := k.GetAllValidatorReputations(ctx)
		for _, rep := range reputations {
			if rep.Reputation.IsNegative() {
				return fmt.Sprintf("validator %s has negative reputation: %s", rep.ValidatorAddress, rep.Reputation), true
			}

			if rep.Reputation.GT(sdk.OneDec()) {
				return fmt.Sprintf("validator %s has reputation greater than 1: %s", rep.ValidatorAddress, rep.Reputation), true
			}
		}

		return "", false
	}
}