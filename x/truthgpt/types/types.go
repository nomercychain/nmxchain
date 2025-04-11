package types

import (
	"encoding/json"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "truthgpt"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_truthgpt"
)

var (
	// Keys for store prefixes
	DataSourceKey       = []byte{0x01} // key for storing data sources
	OracleQueryKey      = []byte{0x02} // key for storing oracle queries
	OracleResponseKey   = []byte{0x03} // key for storing oracle responses
	AIModelKey          = []byte{0x04} // key for storing AI models
	DataSourceRankKey   = []byte{0x05} // key for storing data source rankings
	MisinformationKey   = []byte{0x06} // key for storing detected misinformation
	VerificationTaskKey = []byte{0x07} // key for storing verification tasks
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

// DataSourceType represents the type of data source
type DataSourceType string

const (
	DataSourceTypeAPI       DataSourceType = "api"
	DataSourceTypeWebsite   DataSourceType = "website"
	DataSourceTypeBlockchain DataSourceType = "blockchain"
	DataSourceTypeIPFS      DataSourceType = "ipfs"
	DataSourceTypeCustom    DataSourceType = "custom"
)

// DataSourceStatus represents the status of a data source
type DataSourceStatus string

const (
	DataSourceStatusActive   DataSourceStatus = "active"
	DataSourceStatusInactive DataSourceStatus = "inactive"
	DataSourceStatusPending  DataSourceStatus = "pending"
	DataSourceStatusBlocked  DataSourceStatus = "blocked"
)

// DataSource represents an external data source
type DataSource struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	SourceType  DataSourceType  `json:"source_type"`
	Endpoint    string          `json:"endpoint"`
	Status      DataSourceStatus `json:"status"`
	Owner       sdk.AccAddress  `json:"owner"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	Metadata    json.RawMessage `json:"metadata"`
}

// DataSourceRank represents the ranking of a data source
type DataSourceRank struct {
	SourceID       string    `json:"source_id"`
	Reliability    sdk.Dec   `json:"reliability"`
	Accuracy       sdk.Dec   `json:"accuracy"`
	Timeliness     sdk.Dec   `json:"timeliness"`
	Completeness   sdk.Dec   `json:"completeness"`
	TrustScore     sdk.Dec   `json:"trust_score"`
	LastEvaluated  time.Time `json:"last_evaluated"`
	EvaluationData string    `json:"evaluation_data"`
}

// OracleQueryStatus represents the status of an oracle query
type OracleQueryStatus string

const (
	OracleQueryStatusPending   OracleQueryStatus = "pending"
	OracleQueryStatusProcessing OracleQueryStatus = "processing"
	OracleQueryStatusCompleted OracleQueryStatus = "completed"
	OracleQueryStatusFailed    OracleQueryStatus = "failed"
	OracleQueryStatusDisputed  OracleQueryStatus = "disputed"
)

// OracleQuery represents a query to the oracle
type OracleQuery struct {
	ID           string           `json:"id"`
	Requester    sdk.AccAddress   `json:"requester"`
	QueryType    string           `json:"query_type"`
	Query        string           `json:"query"`
	DataSources  []string         `json:"data_sources,omitempty"`
	Status       OracleQueryStatus `json:"status"`
	Fee          sdk.Coins        `json:"fee"`
	CreatedAt    time.Time        `json:"created_at"`
	CompletedAt  time.Time        `json:"completed_at,omitempty"`
	ResponseID   string           `json:"response_id,omitempty"`
	CallbackData json.RawMessage  `json:"callback_data,omitempty"`
}

// OracleResponse represents a response from the oracle
type OracleResponse struct {
	ID              string          `json:"id"`
	QueryID         string          `json:"query_id"`
	Response        json.RawMessage `json:"response"`
	SourceResponses []SourceResponse `json:"source_responses"`
	Confidence      sdk.Dec         `json:"confidence"`
	ProcessedBy     string          `json:"processed_by"` // AI model ID
	CreatedAt       time.Time       `json:"created_at"`
	Metadata        json.RawMessage `json:"metadata"`
}

// SourceResponse represents a response from a specific data source
type SourceResponse struct {
	SourceID  string          `json:"source_id"`
	Response  json.RawMessage `json:"response"`
	Timestamp time.Time       `json:"timestamp"`
	Status    string          `json:"status"`
	Error     string          `json:"error,omitempty"`
}

// AIModel represents an AI model used by the oracle
type AIModel struct {
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

// Misinformation represents detected misinformation
type Misinformation struct {
	ID          string         `json:"id"`
	Content     string         `json:"content"`
	Source      string         `json:"source"`
	DetectedBy  string         `json:"detected_by"` // AI model ID
	Reporter    sdk.AccAddress `json:"reporter,omitempty"`
	Confidence  sdk.Dec        `json:"confidence"`
	Evidence    string         `json:"evidence"`
	CreatedAt   time.Time      `json:"created_at"`
	Status      string         `json:"status"`
	VerifiedBy  []string       `json:"verified_by,omitempty"`
}

// VerificationTask represents a task to verify information
type VerificationTask struct {
	ID          string         `json:"id"`
	Content     string         `json:"content"`
	Source      string         `json:"source"`
	Creator     sdk.AccAddress `json:"creator"`
	Status      string         `json:"status"`
	Priority    uint64         `json:"priority"`
	CreatedAt   time.Time      `json:"created_at"`
	CompletedAt time.Time      `json:"completed_at,omitempty"`
	Result      json.RawMessage `json:"result,omitempty"`
}