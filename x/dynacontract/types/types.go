package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DynaContract status constants
const (
	DynaContractStatusActive    = "active"
	DynaContractStatusInactive  = "inactive"
	DynaContractStatusPaused    = "paused"
	DynaContractStatusLearning  = "learning"
	DynaContractStatusUpgrading = "upgrading"
)

// Permission type constants
const (
	PermissionTypeExecute = "execute"
	PermissionTypeAdmin   = "admin"
)

// Validate validates the DynaContract
func (c DynaContract) Validate() error {
	if c.Id == "" {
		return fmt.Errorf("contract ID cannot be empty")
	}
	if c.Name == "" {
		return fmt.Errorf("contract name cannot be empty")
	}
	if c.Creator == "" {
		return fmt.Errorf("creator cannot be empty")
	}
	if c.Owner == "" {
		return fmt.Errorf("owner cannot be empty")
	}
	if c.Status == "" {
		return fmt.Errorf("status cannot be empty")
	}
	if c.CodeHash == "" {
		return fmt.Errorf("code hash cannot be empty")
	}
	if len(c.Code) == 0 {
		return fmt.Errorf("code cannot be empty")
	}
	if c.GasLimit == 0 {
		return fmt.Errorf("gas limit cannot be zero")
	}
	return nil
}

// Validate validates the DynaContractExecution
func (e DynaContractExecution) Validate() error {
	if e.Id == "" {
		return fmt.Errorf("execution ID cannot be empty")
	}
	if e.ContractId == "" {
		return fmt.Errorf("contract ID cannot be empty")
	}
	if e.Caller == "" {
		return fmt.Errorf("caller cannot be empty")
	}
	if e.Status == "" {
		return fmt.Errorf("status cannot be empty")
	}
	return nil
}

// Validate validates the DynaContractTemplate
func (t DynaContractTemplate) Validate() error {
	if t.Id == "" {
		return fmt.Errorf("template ID cannot be empty")
	}
	if t.Name == "" {
		return fmt.Errorf("template name cannot be empty")
	}
	if t.Creator == "" {
		return fmt.Errorf("creator cannot be empty")
	}
	if len(t.Code) == 0 {
		return fmt.Errorf("code cannot be empty")
	}
	return nil
}

// Validate validates the DynaContractLearningData
func (d DynaContractLearningData) Validate() error {
	if d.Id == "" {
		return fmt.Errorf("data ID cannot be empty")
	}
	if d.ContractId == "" {
		return fmt.Errorf("contract ID cannot be empty")
	}
	if d.DataType == "" {
		return fmt.Errorf("data type cannot be empty")
	}
	if len(d.Data) == 0 {
		return fmt.Errorf("data cannot be empty")
	}
	return nil
}

// Validate validates the DynaContractPermission
func (p DynaContractPermission) Validate() error {
	if p.ContractId == "" {
		return fmt.Errorf("contract ID cannot be empty")
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
func (p DynaContractPermission) IsExpired(now time.Time) bool {
	return !p.ExpiresAt.IsZero() && p.ExpiresAt.Before(now)
}

// DynaContractKey returns the store key for a dynamic contract
func DynaContractKey(id string) []byte {
	return append(DynaContractKeyPrefix, []byte(id)...)
}

// DynaContractExecutionKey returns the store key for a dynamic contract execution
func DynaContractExecutionKey(id string) []byte {
	return append(DynaContractExecutionKeyPrefix, []byte(id)...)
}

// DynaContractTemplateKey returns the store key for a dynamic contract template
func DynaContractTemplateKey(id string) []byte {
	return append(DynaContractTemplateKeyPrefix, []byte(id)...)
}

// DynaContractLearningDataKey returns the store key for dynamic contract learning data
func DynaContractLearningDataKey(id string) []byte {
	return append(DynaContractLearningDataKeyPrefix, []byte(id)...)
}

// DynaContractPermissionKey returns the store key for a dynamic contract permission
func DynaContractPermissionKey(contractID string, address string, permissionType string) []byte {
	key := append(DynaContractPermissionKeyPrefix, []byte(contractID)...)
	key = append(key, []byte("/")...)
	key = append(key, []byte(address)...)
	key = append(key, []byte("/")...)
	key = append(key, []byte(permissionType)...)
	return key
}

// DynaContractPermissionKeyPrefixByContract returns the store key prefix for all permissions of a contract
func DynaContractPermissionKeyPrefixByContract(contractID string) []byte {
	return append(DynaContractPermissionKeyPrefix, []byte(contractID+"/")...)
}

// DynaContractPermissionKeyPrefixByAddress returns the store key prefix for all permissions of an address
func DynaContractPermissionKeyPrefixByAddress(address string) []byte {
	prefix := append(DynaContractPermissionKeyPrefix, []byte(".+/")...)
	return append(prefix, []byte(address+"/")...)
}