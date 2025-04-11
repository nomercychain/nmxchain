package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nomercychain/nmxchain/x/dynacontract/types"
)

// RegisterInvariants registers all dynacontract invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, "valid-contracts",
		ValidContractsInvariant(k))
}

// ValidContractsInvariant checks that all stored contracts are valid
func ValidContractsInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var invalidContracts []string
		var msg string
		var count int

		k.IterateContracts(ctx, func(contract types.DynaContract) (stop bool) {
			count++

			// Check if contract has valid fields
			if contract.Id == "" {
				invalidContracts = append(invalidContracts, fmt.Sprintf("Contract has empty ID: %v", contract))
			}

			if contract.Creator == "" {
				invalidContracts = append(invalidContracts, fmt.Sprintf("Contract has empty Creator: %s", contract.Id))
			}

			if contract.Owner == "" {
				invalidContracts = append(invalidContracts, fmt.Sprintf("Contract has empty Owner: %s", contract.Id))
			}

			if len(contract.Code) == 0 {
				invalidContracts = append(invalidContracts, fmt.Sprintf("Contract has empty Code: %s", contract.Id))
			}

			return false
		})

		if len(invalidContracts) > 0 {
			msg = fmt.Sprintf("Found %d invalid contracts:\n%s", len(invalidContracts), invalidContracts)
			return sdk.FormatInvariant(types.ModuleName, "valid-contracts", msg), true
		}

		return sdk.FormatInvariant(types.ModuleName, "valid-contracts",
			fmt.Sprintf("All %d contracts are valid", count)), false
	}
}
