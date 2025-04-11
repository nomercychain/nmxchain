package types

import (
	"encoding/json"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "dynacontracts"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_dynacontracts"
)

var (
	// Keys for store prefixes
	ContractKey           = []byte{0x01} // key for storing contracts
	ContractCodeKey       = []byte{0x02} // key for storing contract code
	ContractStateKey      = []byte{0x03} // key for storing contract state
	ContractVersionKey    = []byte{0x04} // key for storing contract versions
	ContractParameterKey  = []byte{0x05} // key for storing contract parameters
	ContractAIModelKey    = []byte{0x06} // key for storing contract AI models
	ContractProposalKey   = []byte{0x07} // key for storing contract proposals
	ContractExecutionKey  = []byte{0x08} // key for storing contract executions
	ExternalDataSourceKey = []byte{0x09} // key for storing external data sources
)

// AccountKeeper defines the expected account keeper
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
	// other methods from the interface you are implementing
}

// BankKeeper defines the expected bank keeper
type BankKeeper interface {
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	// other methods from the interface you are implementing
}

// GovKeeper defines the expected governance keeper
type GovKeeper interface {
	GetProposal(ctx sdk.Context, proposalID uint64) (govtypes.Proposal, bool)
	// other methods from the interface you are implementing
}

// ContractLanguage represents the programming language of a contract
type ContractLanguage string

const (
	ContractLanguageRust    ContractLanguage = "rust"
	ContractLanguageSolidity ContractLanguage = "solidity"
)

// ContractStatus represents the status of a contract
type ContractStatus string

const (
	ContractStatusActive    ContractStatus = "active"
	ContractStatusInactive  ContractStatus = "inactive"
	ContractStatusProposed  ContractStatus = "proposed"
	ContractStatusRejected  ContractStatus = "rejected"
	ContractStatusUpgrading ContractStatus = "upgrading"
)

// Contract represents a DynaContract
type Contract struct {
	ID              string          `json:"id"`
	Name            string          `json:"name"`
	Description     string          `json:"description"`
	Creator         sdk.AccAddress  `json:"creator"`
	Owner           sdk.AccAddress  `json:"owner"`
	Language        ContractLanguage `json:"language"`
	Version         uint64          `json:"version"`
	Status          ContractStatus  `json:"status"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	AIEnabled       bool            `json:"ai_enabled"`
	AIModelID       string          `json:"ai_model_id,omitempty"`
	DataSourceIDs   []string        `json:"data_source_ids,omitempty"`
	GovernanceEnabled bool           `json:"governance_enabled"`
}

// ContractCode represents the code of a contract
type ContractCode struct {
	ContractID string `json:"contract_id"`
	Version    uint64 `json:"version"`
	Code       []byte `json:"code"`
	Checksum   string `json:"checksum"`
	Metadata   string `json:"metadata"`
}

// ContractState represents the state of a contract
type ContractState struct {
	ContractID string          `json:"contract_id"`
	Key        string          `json:"key"`
	Value      json.RawMessage `json:"value"`
	UpdatedAt  time.Time       `json:"updated_at"`
}

// ContractParameter represents a dynamic parameter of a contract
type ContractParameter struct {
	ContractID  string          `json:"contract_id"`
	Name        string          `json:"name"`
	Type        string          `json:"type"`
	Value       json.RawMessage `json:"value"`
	MinValue    json.RawMessage `json:"min_value,omitempty"`
	MaxValue    json.RawMessage `json:"max_value,omitempty"`
	Description string          `json:"description"`
	AIControlled bool           `json:"ai_controlled"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

// AIModel represents an AI model used by a contract
type AIModel struct {
	ID          string    `json:"id"`
	ContractID  string    `json:"contract_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ModelType   string    `json:"model_type"`
	ModelURL    string    `json:"model_url"`
	ModelHash   string    `json:"model_hash"`
	Version     uint64    `json:"version"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ContractProposal represents a proposal to update a contract
type ContractProposal struct {
	ID          uint64         `json:"id"`
	ContractID  string         `json:"contract_id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Proposer    sdk.AccAddress `json:"proposer"`
	NewCode     []byte         `json:"new_code,omitempty"`
	NewParams   []ContractParameter `json:"new_params,omitempty"`
	Status      string         `json:"status"`
	VotesYes    sdk.Int        `json:"votes_yes"`
	VotesNo     sdk.Int        `json:"votes_no"`
	CreatedAt   time.Time      `json:"created_at"`
	EndTime     time.Time      `json:"end_time"`
}

// ContractExecution represents an execution of a contract
type ContractExecution struct {
	ID         string         `json:"id"`
	ContractID string         `json:"contract_id"`
	Caller     sdk.AccAddress `json:"caller"`
	Method     string         `json:"method"`
	Params     json.RawMessage `json:"params"`
	Result     json.RawMessage `json:"result"`
	GasUsed    uint64         `json:"gas_used"`
	Timestamp  time.Time      `json:"timestamp"`
	Success    bool           `json:"success"`
	Error      string         `json:"error,omitempty"`
}

// ExternalDataSource represents an external data source for contracts
type ExternalDataSource struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	Type        string    `json:"type"`
	AuthType    string    `json:"auth_type,omitempty"`
	AuthParams  string    `json:"auth_params,omitempty"`
	UpdateFreq  uint64    `json:"update_freq"`
	Reliability sdk.Dec   `json:"reliability"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}