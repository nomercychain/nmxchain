package types

import (
	"encoding/json"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "deai"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_deai"
)

var (
	// Keys for store prefixes
	AIAgentKey           = []byte{0x01} // key for storing AI agents
	AIAgentModelKey      = []byte{0x02} // key for storing AI agent models
	AIAgentStateKey      = []byte{0x03} // key for storing AI agent states
	AIAgentActionKey     = []byte{0x04} // key for storing AI agent actions
	AIAgentTrainingKey   = []byte{0x05} // key for storing AI agent training data
	AIAgentMarketplaceKey = []byte{0x06} // key for storing AI agent marketplace listings
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

// NFTKeeper defines the expected NFT keeper
type NFTKeeper interface {
	// Methods for NFT operations
}

// AIAgentType represents the type of AI agent
type AIAgentType string

const (
	AIAgentTypeTrading    AIAgentType = "trading"
	AIAgentTypeGovernance AIAgentType = "governance"
	AIAgentTypeStaking    AIAgentType = "staking"
	AIAgentTypeGaming     AIAgentType = "gaming"
	AIAgentTypeCustom     AIAgentType = "custom"
)

// AIAgentStatus represents the status of an AI agent
type AIAgentStatus string

const (
	AIAgentStatusActive    AIAgentStatus = "active"
	AIAgentStatusInactive  AIAgentStatus = "inactive"
	AIAgentStatusTraining  AIAgentStatus = "training"
	AIAgentStatusLocked    AIAgentStatus = "locked"
	AIAgentStatusForSale   AIAgentStatus = "for_sale"
	AIAgentStatusForRent   AIAgentStatus = "for_rent"
)

// AIAgent represents a Digital Twin AI agent
type AIAgent struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Owner       sdk.AccAddress `json:"owner"`
	Creator     sdk.AccAddress `json:"creator"`
	AgentType   AIAgentType    `json:"agent_type"`
	Status      AIAgentStatus  `json:"status"`
	ModelID     string         `json:"model_id"`
	NFTID       string         `json:"nft_id,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	Permissions json.RawMessage `json:"permissions"`
	Metadata    json.RawMessage `json:"metadata"`
}

// AIAgentModel represents an AI model used by an agent
type AIAgentModel struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ModelType   string    `json:"model_type"`
	ModelURL    string    `json:"model_url"`
	ModelHash   string    `json:"model_hash"`
	Version     uint64    `json:"version"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Metadata    json.RawMessage `json:"metadata"`
}

// AIAgentState represents the state of an AI agent
type AIAgentState struct {
	AgentID     string          `json:"agent_id"`
	StateData   json.RawMessage `json:"state_data"`
	StorageType string          `json:"storage_type"` // e.g., "ipfs", "arweave", "chain"
	StorageRef  string          `json:"storage_ref,omitempty"` // Reference to off-chain storage
	UpdatedAt   time.Time       `json:"updated_at"`
}

// AIAgentAction represents an action taken by an AI agent
type AIAgentAction struct {
	ID          string         `json:"id"`
	AgentID     string         `json:"agent_id"`
	ActionType  string         `json:"action_type"`
	Timestamp   time.Time      `json:"timestamp"`
	Data        json.RawMessage `json:"data"`
	Result      json.RawMessage `json:"result,omitempty"`
	Status      string         `json:"status"`
	GasUsed     uint64         `json:"gas_used,omitempty"`
	Error       string         `json:"error,omitempty"`
}

// AIAgentTrainingData represents training data for an AI agent
type AIAgentTrainingData struct {
	ID          string          `json:"id"`
	AgentID     string          `json:"agent_id"`
	DataType    string          `json:"data_type"`
	Data        json.RawMessage `json:"data"`
	Source      string          `json:"source"`
	Timestamp   time.Time       `json:"timestamp"`
	StorageType string          `json:"storage_type"` // e.g., "ipfs", "arweave", "chain"
	StorageRef  string          `json:"storage_ref,omitempty"` // Reference to off-chain storage
}

// AIAgentMarketplaceListing represents a marketplace listing for an AI agent
type AIAgentMarketplaceListing struct {
	ID          string         `json:"id"`
	AgentID     string         `json:"agent_id"`
	Seller      sdk.AccAddress `json:"seller"`
	Price       sdk.Coins      `json:"price"`
	RentalPrice sdk.Coins      `json:"rental_price,omitempty"`
	RentalDuration uint64      `json:"rental_duration,omitempty"`
	ListingType string         `json:"listing_type"` // "sale", "rent", "both"
	Status      string         `json:"status"`
	CreatedAt   time.Time      `json:"created_at"`
	ExpiresAt   time.Time      `json:"expires_at,omitempty"`
}