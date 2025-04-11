package types

import (
	"fmt"
)

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Hyperchains:        []Hyperchain{},
		Validators:         []HyperchainValidator{},
		Blocks:             []HyperchainBlock{},
		Transactions:       []HyperchainTransaction{},
		Bridges:            []HyperchainBridge{},
		BridgeTransactions: []HyperchainBridgeTransaction{},
		Permissions:        []HyperchainPermission{},
		Params:             DefaultParams(),
	}
}

// Validate performs basic genesis state validation
func (gs GenesisState) Validate() error {
	// Validate hyperchains
	hyperchainIDs := make(map[string]bool)
	for _, hyperchain := range gs.Hyperchains {
		if hyperchainIDs[hyperchain.Id] {
			return fmt.Errorf("duplicate hyperchain ID: %s", hyperchain.Id)
		}
		hyperchainIDs[hyperchain.Id] = true

		if err := hyperchain.Validate(); err != nil {
			return fmt.Errorf("invalid hyperchain: %w", err)
		}
	}

	// Validate validators
	validatorKeys := make(map[string]bool)
	for _, validator := range gs.Validators {
		key := fmt.Sprintf("%s/%s", validator.ChainId, validator.Address)
		if validatorKeys[key] {
			return fmt.Errorf("duplicate validator: %s", key)
		}
		validatorKeys[key] = true

		if err := validator.Validate(); err != nil {
			return fmt.Errorf("invalid validator: %w", err)
		}

		// Check if the hyperchain exists
		if !hyperchainIDs[validator.ChainId] {
			return fmt.Errorf("validator references non-existent hyperchain: %s", validator.ChainId)
		}
	}

	// Validate blocks
	blockKeys := make(map[string]bool)
	for _, block := range gs.Blocks {
		key := fmt.Sprintf("%s/%d", block.ChainId, block.Height)
		if blockKeys[key] {
			return fmt.Errorf("duplicate block: %s", key)
		}
		blockKeys[key] = true

		if err := block.Validate(); err != nil {
			return fmt.Errorf("invalid block: %w", err)
		}

		// Check if the hyperchain exists
		if !hyperchainIDs[block.ChainId] {
			return fmt.Errorf("block references non-existent hyperchain: %s", block.ChainId)
		}
	}

	// Validate transactions
	transactionIDs := make(map[string]bool)
	for _, transaction := range gs.Transactions {
		if transactionIDs[transaction.Id] {
			return fmt.Errorf("duplicate transaction ID: %s", transaction.Id)
		}
		transactionIDs[transaction.Id] = true

		if err := transaction.Validate(); err != nil {
			return fmt.Errorf("invalid transaction: %w", err)
		}

		// Check if the hyperchain exists
		if !hyperchainIDs[transaction.ChainId] {
			return fmt.Errorf("transaction references non-existent hyperchain: %s", transaction.ChainId)
		}
	}

	// Validate bridges
	bridgeIDs := make(map[string]bool)
	for _, bridge := range gs.Bridges {
		if bridgeIDs[bridge.Id] {
			return fmt.Errorf("duplicate bridge ID: %s", bridge.Id)
		}
		bridgeIDs[bridge.Id] = true

		if err := bridge.Validate(); err != nil {
			return fmt.Errorf("invalid bridge: %w", err)
		}

		// Check if the hyperchains exist
		if !hyperchainIDs[bridge.SourceChainId] {
			return fmt.Errorf("bridge references non-existent source hyperchain: %s", bridge.SourceChainId)
		}
		if !hyperchainIDs[bridge.TargetChainId] {
			return fmt.Errorf("bridge references non-existent target hyperchain: %s", bridge.TargetChainId)
		}
	}

	// Validate bridge transactions
	bridgeTransactionIDs := make(map[string]bool)
	for _, transaction := range gs.BridgeTransactions {
		if bridgeTransactionIDs[transaction.Id] {
			return fmt.Errorf("duplicate bridge transaction ID: %s", transaction.Id)
		}
		bridgeTransactionIDs[transaction.Id] = true

		if err := transaction.Validate(); err != nil {
			return fmt.Errorf("invalid bridge transaction: %w", err)
		}

		// Check if the bridge exists
		if !bridgeIDs[transaction.BridgeId] {
			return fmt.Errorf("bridge transaction references non-existent bridge: %s", transaction.BridgeId)
		}

		// Check if the hyperchains exist
		if !hyperchainIDs[transaction.SourceChainId] {
			return fmt.Errorf("bridge transaction references non-existent source hyperchain: %s", transaction.SourceChainId)
		}
		if !hyperchainIDs[transaction.TargetChainId] {
			return fmt.Errorf("bridge transaction references non-existent target hyperchain: %s", transaction.TargetChainId)
		}
	}

	// Validate permissions
	permissionKeys := make(map[string]bool)
	for _, permission := range gs.Permissions {
		key := fmt.Sprintf("%s/%s/%s", permission.ChainId, permission.Address, permission.PermissionType)
		if permissionKeys[key] {
			return fmt.Errorf("duplicate permission: %s", key)
		}
		permissionKeys[key] = true

		if err := permission.Validate(); err != nil {
			return fmt.Errorf("invalid permission: %w", err)
		}

		// Check if the hyperchain exists
		if !hyperchainIDs[permission.ChainId] {
			return fmt.Errorf("permission references non-existent hyperchain: %s", permission.ChainId)
		}
	}

	// Validate params
	if err := gs.Params.Validate(); err != nil {
		return fmt.Errorf("invalid params: %w", err)
	}

	return nil
}