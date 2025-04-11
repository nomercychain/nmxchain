package types

import (
	"encoding/json"
)

// Message response types
type MsgCreateAIAgentResponse struct {
	ID string `json:"id"`
}

type MsgUpdateAIAgentResponse struct {}

type MsgTrainAIAgentResponse struct {
	TrainingDataID string `json:"training_data_id"`
}

type MsgExecuteAIAgentResponse struct {
	ActionID string          `json:"action_id"`
	Result   json.RawMessage `json:"result"`
}

type MsgListAIAgentForSaleResponse struct {
	ListingID string `json:"listing_id"`
}

type MsgBuyAIAgentResponse struct {
	AgentID string `json:"agent_id"`
}

type MsgRentAIAgentResponse struct {
	AgentID  string `json:"agent_id"`
	Duration uint64 `json:"duration"`
}

type MsgCancelMarketListingResponse struct {}