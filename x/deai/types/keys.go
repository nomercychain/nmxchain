package types

// Module name and store keys
const (
	ModuleName   = "deai"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	QuerierRoute = ModuleName
	MemStoreKey  = "mem_" + ModuleName
)

// Key prefixes for store keys
var (
	AIAgentKey                    = []byte{0x01} // prefix for AI agents
	AIAgentByOwnerKey             = []byte{0x02} // prefix for AI agents by owner
	AIAgentStateKey               = []byte{0x03} // prefix for AI agent states
	AIAgentActionKey              = []byte{0x04} // prefix for AI agent actions
	AIAgentActionByAgentKey       = []byte{0x05} // prefix for AI agent actions by agent
	AIAgentModelKey               = []byte{0x06} // prefix for AI agent models
	AIAgentTrainingKey            = []byte{0x07} // prefix for AI agent training data
	AIAgentTrainingByAgentKey     = []byte{0x08} // prefix for AI agent training data by agent
	AIAgentMarketplaceKey         = []byte{0x09} // prefix for AI agent marketplace listings
	AIAgentMarketplaceByAgentKey  = []byte{0x0A} // prefix for AI agent marketplace listings by agent
	AIAgentMarketplaceBySellerKey = []byte{0x0B} // prefix for AI agent marketplace listings by seller
)

// GetAIAgentKey returns the store key to retrieve an AI agent by ID
func GetAIAgentKey(id string) []byte {
	return append(AIAgentKey, []byte(id)...)
}

// GetAIAgentByOwnerKey returns the store key to retrieve AI agents by owner
func GetAIAgentByOwnerKey(owner []byte, id string) []byte {
	return append(append(AIAgentByOwnerKey, owner...), []byte(id)...)
}

// GetAIAgentStateKey returns the store key to retrieve an AI agent state by agent ID
func GetAIAgentStateKey(agentID string) []byte {
	return append(AIAgentStateKey, []byte(agentID)...)
}

// GetAIAgentActionKey returns the store key to retrieve an AI agent action by ID
func GetAIAgentActionKey(id string) []byte {
	return append(AIAgentActionKey, []byte(id)...)
}

// GetAIAgentActionByAgentKey returns the store key to retrieve AI agent actions by agent ID
func GetAIAgentActionByAgentKey(agentID string, actionID string) []byte {
	return append(append(AIAgentActionByAgentKey, []byte(agentID)...), []byte(actionID)...)
}

// GetAIAgentModelKey returns the store key to retrieve an AI agent model by ID
func GetAIAgentModelKey(id string) []byte {
	return append(AIAgentModelKey, []byte(id)...)
}

// GetAIAgentTrainingKey returns the store key to retrieve AI agent training data by ID
func GetAIAgentTrainingKey(id string) []byte {
	return append(AIAgentTrainingKey, []byte(id)...)
}

// GetAIAgentTrainingByAgentKey returns the store key to retrieve AI agent training data by agent ID
func GetAIAgentTrainingByAgentKey(agentID string, trainingID string) []byte {
	return append(append(AIAgentTrainingByAgentKey, []byte(agentID)...), []byte(trainingID)...)
}

// GetAIAgentMarketplaceKey returns the store key to retrieve an AI agent marketplace listing by ID
func GetAIAgentMarketplaceKey(id string) []byte {
	return append(AIAgentMarketplaceKey, []byte(id)...)
}

// GetAIAgentMarketplaceByAgentKey returns the store key to retrieve AI agent marketplace listings by agent ID
func GetAIAgentMarketplaceByAgentKey(agentID string, listingID string) []byte {
	return append(append(AIAgentMarketplaceByAgentKey, []byte(agentID)...), []byte(listingID)...)
}

// GetAIAgentMarketplaceBySellerKey returns the store key to retrieve AI agent marketplace listings by seller
func GetAIAgentMarketplaceBySellerKey(seller []byte, listingID string) []byte {
	return append(append(AIAgentMarketplaceBySellerKey, seller...), []byte(listingID)...)
}
