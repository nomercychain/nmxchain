package types

import (
	"encoding/json"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AI agent status constants
const (
	AIAgentStatusInactive = "inactive"
	AIAgentStatusActive   = "active"
	AIAgentStatusTraining = "training"
	AIAgentStatusForSale  = "for_sale"
	AIAgentStatusForRent  = "for_rent"
)

// AIAgentType defines the type of AI agent
type AIAgentType string

// AI agent type constants
const (
	AIAgentTypeGeneral     AIAgentType = "general"
	AIAgentTypeText        AIAgentType = "text"
	AIAgentTypeImage       AIAgentType = "image"
	AIAgentTypeAudio       AIAgentType = "audio"
	AIAgentTypeVideo       AIAgentType = "video"
	AIAgentTypeMultimodal  AIAgentType = "multimodal"
	AIAgentTypeSpecialized AIAgentType = "specialized"
)

// AIAgent defines the structure for an AI agent
type AIAgent struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Owner       sdk.AccAddress  `json:"owner"`
	Creator     sdk.AccAddress  `json:"creator"`
	AgentType   AIAgentType     `json:"agent_type"`
	Status      string          `json:"status"`
	ModelID     string          `json:"model_id"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	Permissions json.RawMessage `json:"permissions,omitempty"`
	Metadata    json.RawMessage `json:"metadata,omitempty"`
}

// Validate performs basic validation of the AI agent
func (a AIAgent) Validate() error {
	if a.ID == "" {
		return fmt.Errorf("agent ID cannot be empty")
	}
	if a.Name == "" {
		return fmt.Errorf("agent name cannot be empty")
	}
	if a.Owner.Empty() {
		return fmt.Errorf("owner address cannot be empty")
	}
	if a.Creator.Empty() {
		return fmt.Errorf("creator address cannot be empty")
	}
	if a.ModelID == "" {
		return fmt.Errorf("model ID cannot be empty")
	}

	// Validate status
	validStatus := map[string]bool{
		AIAgentStatusInactive: true,
		AIAgentStatusActive:   true,
		AIAgentStatusTraining: true,
		AIAgentStatusForSale:  true,
		AIAgentStatusForRent:  true,
	}
	if !validStatus[a.Status] {
		return fmt.Errorf("invalid agent status: %s", a.Status)
	}

	// Validate agent type
	validTypes := map[AIAgentType]bool{
		AIAgentTypeGeneral:     true,
		AIAgentTypeText:        true,
		AIAgentTypeImage:       true,
		AIAgentTypeAudio:       true,
		AIAgentTypeVideo:       true,
		AIAgentTypeMultimodal:  true,
		AIAgentTypeSpecialized: true,
	}
	if !validTypes[a.AgentType] {
		return fmt.Errorf("invalid agent type: %s", a.AgentType)
	}

	return nil
}

// AIAgentState defines the state of an AI agent
type AIAgentState struct {
	AgentID     string    `json:"agent_id"`
	StateData   []byte    `json:"state_data"`
	StorageType string    `json:"storage_type"` // "chain", "ipfs", "arweave", etc.
	UpdatedAt   time.Time `json:"updated_at"`
}

// AIAgentAction defines an action performed by an AI agent
type AIAgentAction struct {
	ID         string          `json:"id"`
	AgentID    string          `json:"agent_id"`
	ActionType string          `json:"action_type"`
	Timestamp  time.Time       `json:"timestamp"`
	Data       json.RawMessage `json:"data,omitempty"`
	Result     json.RawMessage `json:"result,omitempty"`
	Status     string          `json:"status"` // "pending", "processing", "completed", "failed"
	GasUsed    uint64          `json:"gas_used"`
}

// AIAgentTrainingData defines training data for an AI agent
type AIAgentTrainingData struct {
	ID          string          `json:"id"`
	AgentID     string          `json:"agent_id"`
	DataType    string          `json:"data_type"`
	Data        json.RawMessage `json:"data"`
	Source      string          `json:"source,omitempty"`
	Timestamp   time.Time       `json:"timestamp"`
	StorageType string          `json:"storage_type"` // "chain", "ipfs", "arweave", etc.
}

// AIAgentModel defines an AI model that can be used by agents
type AIAgentModel struct {
	ID           string          `json:"id"`
	Name         string          `json:"name"`
	Description  string          `json:"description"`
	Version      string          `json:"version"`
	Creator      sdk.AccAddress  `json:"creator"`
	ModelType    string          `json:"model_type"`
	Capabilities []string        `json:"capabilities"`
	Parameters   json.RawMessage `json:"parameters,omitempty"`
	Metadata     json.RawMessage `json:"metadata,omitempty"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
}

// AIAgentMarketplaceListing defines a marketplace listing for an AI agent
type AIAgentMarketplaceListing struct {
	ID             string         `json:"id"`
	AgentID        string         `json:"agent_id"`
	Seller         sdk.AccAddress `json:"seller"`
	Price          sdk.Coins      `json:"price,omitempty"`
	RentalPrice    sdk.Coins      `json:"rental_price,omitempty"`
	RentalDuration uint64         `json:"rental_duration,omitempty"`
	ListingType    string         `json:"listing_type"` // "sale", "rent", "both"
	Status         string         `json:"status"`       // "active", "completed", "cancelled", "expired"
	CreatedAt      time.Time      `json:"created_at"`
	ExpiresAt      time.Time      `json:"expires_at"`
}
