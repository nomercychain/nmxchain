package types

import (
	"encoding/json"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "hyperchains"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_hyperchains"
)

var (
	// Keys for store prefixes
	ChainKey           = []byte{0x01} // key for storing chains
	ChainTemplateKey   = []byte{0x02} // key for storing chain templates
	ChainModuleKey     = []byte{0x03} // key for storing chain modules
	ChainDeploymentKey = []byte{0x04} // key for storing chain deployments
	ChainValidatorKey  = []byte{0x05} // key for storing chain validators
	ChainProposalKey   = []byte{0x06} // key for storing chain proposals
	ChainMetricsKey    = []byte{0x07} // key for storing chain metrics
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

// ChainType represents the type of chain
type ChainType string

const (
	ChainTypeZkEVM     ChainType = "zkevm"
	ChainTypeOptimistic ChainType = "optimistic"
	ChainTypeRollup    ChainType = "rollup"
	ChainTypeAppChain  ChainType = "appchain"
	ChainTypeCustom    ChainType = "custom"
)

// ChainStatus represents the status of a chain
type ChainStatus string

const (
	ChainStatusProposed  ChainStatus = "proposed"
	ChainStatusApproved  ChainStatus = "approved"
	ChainStatusDeploying ChainStatus = "deploying"
	ChainStatusActive    ChainStatus = "active"
	ChainStatusPaused    ChainStatus = "paused"
	ChainStatusUpgrading ChainStatus = "upgrading"
	ChainStatusRetired   ChainStatus = "retired"
)

// Chain represents a Layer 3 chain
type Chain struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Creator     sdk.AccAddress `json:"creator"`
	ChainType   ChainType   `json:"chain_type"`
	Status      ChainStatus `json:"status"`
	Version     uint64      `json:"version"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	Modules     []string    `json:"modules"`
	Config      json.RawMessage `json:"config"`
	Metadata    json.RawMessage `json:"metadata"`
}

// ChainTemplate represents a template for creating chains
type ChainTemplate struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ChainType   ChainType `json:"chain_type"`
	Creator     sdk.AccAddress `json:"creator"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Modules     []string  `json:"modules"`
	Config      json.RawMessage `json:"config"`
	Metadata    json.RawMessage `json:"metadata"`
}

// ChainModule represents a module that can be included in a chain
type ChainModule struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Version     string    `json:"version"`
	Creator     sdk.AccAddress `json:"creator"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Code        string    `json:"code"`
	Config      json.RawMessage `json:"config"`
	Dependencies []string  `json:"dependencies"`
	Metadata    json.RawMessage `json:"metadata"`
}

// ChainDeployment represents a deployment of a chain
type ChainDeployment struct {
	ID          string    `json:"id"`
	ChainID     string    `json:"chain_id"`
	Version     uint64    `json:"version"`
	Deployer    sdk.AccAddress `json:"deployer"`
	Status      string    `json:"status"`
	StartedAt   time.Time `json:"started_at"`
	CompletedAt time.Time `json:"completed_at,omitempty"`
	Logs        string    `json:"logs,omitempty"`
	Endpoints   json.RawMessage `json:"endpoints,omitempty"`
	Config      json.RawMessage `json:"config"`
}

// ChainValidator represents a validator for a chain
type ChainValidator struct {
	ChainID          string         `json:"chain_id"`
	ValidatorAddress sdk.ValAddress `json:"validator_address"`
	OperatorAddress  sdk.AccAddress `json:"operator_address"`
	Status           string         `json:"status"`
	Power            int64          `json:"power"`
	JoinedAt         time.Time      `json:"joined_at"`
	LastActive       time.Time      `json:"last_active"`
}

// ChainProposal represents a proposal for a chain
type ChainProposal struct {
	ID          uint64         `json:"id"`
	ChainID     string         `json:"chain_id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Proposer    sdk.AccAddress `json:"proposer"`
	ProposalType string        `json:"proposal_type"`
	Status      string         `json:"status"`
	VotesYes    sdk.Int        `json:"votes_yes"`
	VotesNo     sdk.Int        `json:"votes_no"`
	CreatedAt   time.Time      `json:"created_at"`
	EndTime     time.Time      `json:"end_time"`
	Content     json.RawMessage `json:"content"`
}

// ChainMetrics represents metrics for a chain
type ChainMetrics struct {
	ChainID           string    `json:"chain_id"`
	BlockHeight       int64     `json:"block_height"`
	TotalTransactions uint64    `json:"total_transactions"`
	ActiveUsers       uint64    `json:"active_users"`
	TPS               sdk.Dec   `json:"tps"`
	AverageFee        sdk.Dec   `json:"average_fee"`
	TotalValueLocked  sdk.Coins `json:"total_value_locked"`
	UpdatedAt         time.Time `json:"updated_at"`
}