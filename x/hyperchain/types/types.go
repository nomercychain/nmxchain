package types

import (
	"fmt"
	"time"
)

// Hyperchain status constants
const (
	HyperchainStatusInitializing = "initializing"
	HyperchainStatusActive       = "active"
	HyperchainStatusInactive     = "inactive"
	HyperchainStatusPaused       = "paused"
	HyperchainStatusUpgrading    = "upgrading"
)

// Validator status constants
const (
	ValidatorStatusActive   = "active"
	ValidatorStatusInactive = "inactive"
	ValidatorStatusJailed   = "jailed"
)

// Bridge status constants
const (
	BridgeStatusActive   = "active"
	BridgeStatusInactive = "inactive"
	BridgeStatusPaused   = "paused"
)

// Bridge transaction status constants
const (
	BridgeTransactionStatusPending   = "pending"
	BridgeTransactionStatusCompleted = "completed"
	BridgeTransactionStatusFailed    = "failed"
)

// Transaction status constants
const (
	TransactionStatusSuccess = "success"
	TransactionStatusFailed  = "failed"
	TransactionStatusPending = "pending"
)

// Permission type constants
const (
	PermissionTypeAdmin     = "admin"
	PermissionTypeValidator = "validator"
	PermissionTypeRelayer   = "relayer"
	PermissionTypeUser      = "user"
	PermissionTypeRead      = "read"
	PermissionTypeWrite     = "write"
)

// IsValidPermissionType checks if a permission type is valid
func IsValidPermissionType(permissionType string) bool {
	switch permissionType {
	case PermissionTypeAdmin, PermissionTypeValidator, PermissionTypeRelayer, PermissionTypeUser, PermissionTypeRead, PermissionTypeWrite:
		return true
	default:
		return false
	}
}

// Validate validates the Hyperchain
func (h Hyperchain) Validate() error {
	if h.Id == "" {
		return fmt.Errorf("hyperchain ID cannot be empty")
	}
	if h.Name == "" {
		return fmt.Errorf("hyperchain name cannot be empty")
	}
	if h.Creator == "" {
		return fmt.Errorf("creator cannot be empty")
	}
	if h.Admin == "" {
		return fmt.Errorf("admin cannot be empty")
	}
	if h.Status == "" {
		return fmt.Errorf("status cannot be empty")
	}
	return nil
}

// Validate validates the HyperchainValidator
func (v HyperchainValidator) Validate() error {
	if v.ChainId == "" {
		return fmt.Errorf("chain ID cannot be empty")
	}
	if v.Address == "" {
		return fmt.Errorf("address cannot be empty")
	}
	if v.Pubkey == "" {
		return fmt.Errorf("pubkey cannot be empty")
	}
	if v.Status == "" {
		return fmt.Errorf("status cannot be empty")
	}
	return nil
}

// Validate validates the HyperchainBlock
func (b HyperchainBlock) Validate() error {
	if b.ChainId == "" {
		return fmt.Errorf("chain ID cannot be empty")
	}
	if b.Hash == "" {
		return fmt.Errorf("hash cannot be empty")
	}
	if b.Proposer == "" {
		return fmt.Errorf("proposer cannot be empty")
	}
	return nil
}

// Validate validates the HyperchainTransaction
func (t HyperchainTransaction) Validate() error {
	if t.Id == "" {
		return fmt.Errorf("transaction ID cannot be empty")
	}
	if t.ChainId == "" {
		return fmt.Errorf("chain ID cannot be empty")
	}
	if t.Sender == "" {
		return fmt.Errorf("sender cannot be empty")
	}
	if t.Recipient == "" {
		return fmt.Errorf("recipient cannot be empty")
	}
	if t.Status == "" {
		return fmt.Errorf("status cannot be empty")
	}
	return nil
}

// Validate validates the HyperchainBridge
func (b HyperchainBridge) Validate() error {
	if b.Id == "" {
		return fmt.Errorf("bridge ID cannot be empty")
	}
	if b.SourceChainId == "" {
		return fmt.Errorf("source chain ID cannot be empty")
	}
	if b.TargetChainId == "" {
		return fmt.Errorf("target chain ID cannot be empty")
	}
	if b.Status == "" {
		return fmt.Errorf("status cannot be empty")
	}
	if b.Creator == "" {
		return fmt.Errorf("creator cannot be empty")
	}
	if b.Admin == "" {
		return fmt.Errorf("admin cannot be empty")
	}
	if b.MinRelayers == 0 {
		return fmt.Errorf("min relayers cannot be zero")
	}
	return nil
}

// Validate validates the HyperchainBridgeTransaction
func (t HyperchainBridgeTransaction) Validate() error {
	if t.Id == "" {
		return fmt.Errorf("transaction ID cannot be empty")
	}
	if t.BridgeId == "" {
		return fmt.Errorf("bridge ID cannot be empty")
	}
	if t.SourceChainId == "" {
		return fmt.Errorf("source chain ID cannot be empty")
	}
	if t.TargetChainId == "" {
		return fmt.Errorf("target chain ID cannot be empty")
	}
	if t.Sender == "" {
		return fmt.Errorf("sender cannot be empty")
	}
	if t.Recipient == "" {
		return fmt.Errorf("recipient cannot be empty")
	}
	if t.Status == "" {
		return fmt.Errorf("status cannot be empty")
	}
	return nil
}

// Validate validates the HyperchainPermission
func (p HyperchainPermission) Validate() error {
	if p.ChainId == "" {
		return fmt.Errorf("chain ID cannot be empty")
	}
	if p.Address == "" {
		return fmt.Errorf("address cannot be empty")
	}
	if p.PermissionType == "" {
		return fmt.Errorf("permission type cannot be empty")
	}
	if p.GrantedBy == "" {
		return fmt.Errorf("granted by cannot be empty")
	}
	return nil
}

// IsExpired checks if the permission has expired
func (p HyperchainPermission) IsExpired(now time.Time) bool {
	return p.ExpiresAt != nil && !p.ExpiresAt.IsZero() && p.ExpiresAt.Before(now)
}

// HyperchainKey returns the store key for a hyperchain
func HyperchainKey(id string) []byte {
	return append(HyperchainKeyPrefix, []byte(id)...)
}

// HyperchainValidatorKey returns the store key for a hyperchain validator
func HyperchainValidatorKey(chainID string, address string) []byte {
	key := append(HyperchainValidatorKeyPrefix, []byte(chainID)...)
	key = append(key, []byte("/")...)
	key = append(key, []byte(address)...)
	return key
}

// HyperchainValidatorKeyPrefixByChain returns the store key prefix for all validators of a hyperchain
func HyperchainValidatorKeyPrefixByChain(chainID string) []byte {
	return append(HyperchainValidatorKeyPrefix, []byte(chainID+"/")...)
}

// HyperchainBlockKey returns the store key for a hyperchain block
func HyperchainBlockKey(chainID string, height uint64) []byte {
	key := append(HyperchainBlockKeyPrefix, []byte(chainID)...)
	key = append(key, []byte("/")...)
	key = append(key, []byte(fmt.Sprintf("%d", height))...)
	return key
}

// HyperchainBlockKeyPrefixByChain returns the store key prefix for all blocks of a hyperchain
func HyperchainBlockKeyPrefixByChain(chainID string) []byte {
	return append(HyperchainBlockKeyPrefix, []byte(chainID+"/")...)
}

// HyperchainTransactionKey returns the store key for a hyperchain transaction
func HyperchainTransactionKey(id string) []byte {
	return append(HyperchainTransactionKeyPrefix, []byte(id)...)
}

// HyperchainBridgeKey returns the store key for a hyperchain bridge
func HyperchainBridgeKey(id string) []byte {
	return append(HyperchainBridgeKeyPrefix, []byte(id)...)
}

// HyperchainBridgeTransactionKey returns the store key for a hyperchain bridge transaction
func HyperchainBridgeTransactionKey(id string) []byte {
	return append(HyperchainBridgeTransactionKeyPrefix, []byte(id)...)
}

// HyperchainPermissionKey returns the store key for a hyperchain permission
func HyperchainPermissionKey(chainID string, address string, permissionType string) []byte {
	key := append(HyperchainPermissionKeyPrefix, []byte(chainID)...)
	key = append(key, []byte("/")...)
	key = append(key, []byte(address)...)
	key = append(key, []byte("/")...)
	key = append(key, []byte(permissionType)...)
	return key
}

// HyperchainPermissionKeyPrefixByChain returns the store key prefix for all permissions of a hyperchain
func HyperchainPermissionKeyPrefixByChain(chainID string) []byte {
	return append(HyperchainPermissionKeyPrefix, []byte(chainID+"/")...)
}
