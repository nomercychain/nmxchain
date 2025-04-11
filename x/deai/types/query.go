package types

// Query endpoints
const (
	QueryAIAgent            = "ai_agent"
	QueryAIAgents           = "ai_agents"
	QueryAIAgentState       = "ai_agent_state"
	QueryAIAgentActions     = "ai_agent_actions"
	QueryAIAgentAction      = "ai_agent_action"
	QueryAIAgentModels      = "ai_agent_models"
	QueryAIAgentModel       = "ai_agent_model"
	QueryAIAgentTrainingData = "ai_agent_training_data"
	QueryMarketplaceListings = "marketplace_listings"
	QueryMarketplaceListing  = "marketplace_listing"
)

// QueryAIAgentRequest is the request type for the Query/AIAgent RPC method
type QueryAIAgentRequest struct {
	ID string `json:"id"`
}

// QueryAIAgentResponse is the response type for the Query/AIAgent RPC method
type QueryAIAgentResponse struct {
	Agent AIAgent `json:"agent"`
}

// QueryAIAgentsRequest is the request type for the Query/AIAgents RPC method
type QueryAIAgentsRequest struct {
	Owner string `json:"owner,omitempty"`
}

// QueryAIAgentsResponse is the response type for the Query/AIAgents RPC method
type QueryAIAgentsResponse struct {
	Agents []AIAgent `json:"agents"`
}

// QueryAIAgentStateRequest is the request type for the Query/AIAgentState RPC method
type QueryAIAgentStateRequest struct {
	AgentID string `json:"agent_id"`
}

// QueryAIAgentStateResponse is the response type for the Query/AIAgentState RPC method
type QueryAIAgentStateResponse struct {
	State AIAgentState `json:"state"`
}

// QueryAIAgentActionsRequest is the request type for the Query/AIAgentActions RPC method
type QueryAIAgentActionsRequest struct {
	AgentID string `json:"agent_id"`
}

// QueryAIAgentActionsResponse is the response type for the Query/AIAgentActions RPC method
type QueryAIAgentActionsResponse struct {
	Actions []AIAgentAction `json:"actions"`
}

// QueryAIAgentActionRequest is the request type for the Query/AIAgentAction RPC method
type QueryAIAgentActionRequest struct {
	ID string `json:"id"`
}

// QueryAIAgentActionResponse is the response type for the Query/AIAgentAction RPC method
type QueryAIAgentActionResponse struct {
	Action AIAgentAction `json:"action"`
}

// QueryAIAgentModelsRequest is the request type for the Query/AIAgentModels RPC method
type QueryAIAgentModelsRequest struct {}

// QueryAIAgentModelsResponse is the response type for the Query/AIAgentModels RPC method
type QueryAIAgentModelsResponse struct {
	Models []AIAgentModel `json:"models"`
}

// QueryAIAgentModelRequest is the request type for the Query/AIAgentModel RPC method
type QueryAIAgentModelRequest struct {
	ID string `json:"id"`
}

// QueryAIAgentModelResponse is the response type for the Query/AIAgentModel RPC method
type QueryAIAgentModelResponse struct {
	Model AIAgentModel `json:"model"`
}

// QueryAIAgentTrainingDataRequest is the request type for the Query/AIAgentTrainingData RPC method
type QueryAIAgentTrainingDataRequest struct {
	AgentID string `json:"agent_id"`
}

// QueryAIAgentTrainingDataResponse is the response type for the Query/AIAgentTrainingData RPC method
type QueryAIAgentTrainingDataResponse struct {
	TrainingData []AIAgentTrainingData `json:"training_data"`
}

// QueryMarketplaceListingsRequest is the request type for the Query/MarketplaceListings RPC method
type QueryMarketplaceListingsRequest struct {
	Status      string `json:"status,omitempty"`
	ListingType string `json:"listing_type,omitempty"`
}

// QueryMarketplaceListingsResponse is the response type for the Query/MarketplaceListings RPC method
type QueryMarketplaceListingsResponse struct {
	Listings []AIAgentMarketplaceListing `json:"listings"`
}

// QueryMarketplaceListingRequest is the request type for the Query/MarketplaceListing RPC method
type QueryMarketplaceListingRequest struct {
	ID string `json:"id"`
}

// QueryMarketplaceListingResponse is the response type for the Query/MarketplaceListing RPC method
type QueryMarketplaceListingResponse struct {
	Listing AIAgentMarketplaceListing `json:"listing"`
}