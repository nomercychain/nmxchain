package types

import (
	"fmt"
)

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Contracts:    []DynaContract{},
		Executions:   []DynaContractExecution{},
		Templates:    []DynaContractTemplate{},
		LearningData: []DynaContractLearningData{},
		Permissions:  []DynaContractPermission{},
		Params:       DefaultParams(),
	}
}

// Validate performs basic genesis state validation
func (gs GenesisState) Validate() error {
	// Validate contracts
	contractIDs := make(map[string]bool)
	for _, contract := range gs.Contracts {
		if contractIDs[contract.Id] {
			return fmt.Errorf("duplicate contract ID: %s", contract.Id)
		}
		contractIDs[contract.Id] = true

		if err := contract.Validate(); err != nil {
			return fmt.Errorf("invalid contract: %w", err)
		}
	}

	// Validate executions
	executionIDs := make(map[string]bool)
	for _, execution := range gs.Executions {
		if executionIDs[execution.Id] {
			return fmt.Errorf("duplicate execution ID: %s", execution.Id)
		}
		executionIDs[execution.Id] = true

		if err := execution.Validate(); err != nil {
			return fmt.Errorf("invalid execution: %w", err)
		}

		// Check if the contract exists
		if !contractIDs[execution.ContractId] {
			return fmt.Errorf("execution references non-existent contract: %s", execution.ContractId)
		}
	}

	// Validate templates
	templateIDs := make(map[string]bool)
	for _, template := range gs.Templates {
		if templateIDs[template.Id] {
			return fmt.Errorf("duplicate template ID: %s", template.Id)
		}
		templateIDs[template.Id] = true

		if err := template.Validate(); err != nil {
			return fmt.Errorf("invalid template: %w", err)
		}
	}

	// Validate learning data
	learningDataIDs := make(map[string]bool)
	for _, data := range gs.LearningData {
		if learningDataIDs[data.Id] {
			return fmt.Errorf("duplicate learning data ID: %s", data.Id)
		}
		learningDataIDs[data.Id] = true

		if err := data.Validate(); err != nil {
			return fmt.Errorf("invalid learning data: %w", err)
		}

		// Check if the contract exists
		if !contractIDs[data.ContractId] {
			return fmt.Errorf("learning data references non-existent contract: %s", data.ContractId)
		}
	}

	// Validate permissions
	permissionKeys := make(map[string]bool)
	for _, permission := range gs.Permissions {
		key := fmt.Sprintf("%s/%s/%s", permission.ContractId, permission.Address, permission.PermissionType)
		if permissionKeys[key] {
			return fmt.Errorf("duplicate permission: %s", key)
		}
		permissionKeys[key] = true

		if err := permission.Validate(); err != nil {
			return fmt.Errorf("invalid permission: %w", err)
		}

		// Check if the contract exists
		if !contractIDs[permission.ContractId] {
			return fmt.Errorf("permission references non-existent contract: %s", permission.ContractId)
		}
	}

	// Validate params
	if err := gs.Params.Validate(); err != nil {
		return fmt.Errorf("invalid params: %w", err)
	}

	return nil
}