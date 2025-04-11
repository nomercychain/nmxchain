package types

import (
	"fmt"
)

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Agents:              []AIAgent{},
		Models:              []AIAgentModel{},
		States:              []AIAgentState{},
		Actions:             []AIAgentAction{},
		TrainingData:        []AIAgentTrainingData{},
		MarketplaceListings: []AIAgentMarketplaceListing{},
		Params:              DefaultParams(),
	}
}

// Validate performs basic genesis state validation
func (gs GenesisState) Validate() error {
	// Validate agents
	agentIDs := make(map[string]bool)
	for _, agent := range gs.Agents {
		if agentIDs[agent.ID] {
			return fmt.Errorf("duplicate agent ID: %s", agent.ID)
		}
		agentIDs[agent.ID] = true

		if err := agent.Validate(); err != nil {
			return err
		}
	}

	// Validate models
	modelIDs := make(map[string]bool)
	for _, model := range gs.Models {
		if modelIDs[model.ID] {
			return fmt.Errorf("duplicate model ID: %s", model.ID)
		}
		modelIDs[model.ID] = true
	}

	// Validate states
	for _, state := range gs.States {
		if !agentIDs[state.AgentID] {
			return fmt.Errorf("state references non-existent agent: %s", state.AgentID)
		}
	}

	// Validate actions
	actionIDs := make(map[string]bool)
	for _, action := range gs.Actions {
		if actionIDs[action.ID] {
			return fmt.Errorf("duplicate action ID: %s", action.ID)
		}
		actionIDs[action.ID] = true

		if !agentIDs[action.AgentID] {
			return fmt.Errorf("action references non-existent agent: %s", action.AgentID)
		}
	}

	// Validate training data
	trainingDataIDs := make(map[string]bool)
	for _, data := range gs.TrainingData {
		if trainingDataIDs[data.ID] {
			return fmt.Errorf("duplicate training data ID: %s", data.ID)
		}
		trainingDataIDs[data.ID] = true

		if !agentIDs[data.AgentID] {
			return fmt.Errorf("training data references non-existent agent: %s", data.AgentID)
		}
	}

	// Validate marketplace listings
	listingIDs := make(map[string]bool)
	for _, listing := range gs.MarketplaceListings {
		if listingIDs[listing.ID] {
			return fmt.Errorf("duplicate listing ID: %s", listing.ID)
		}
		listingIDs[listing.ID] = true

		if !agentIDs[listing.AgentID] {
			return fmt.Errorf("listing references non-existent agent: %s", listing.AgentID)
		}
	}

	return gs.Params.Validate()
}

// GenesisState defines the deai module's genesis state
type GenesisState struct {
	Agents              []AIAgent                   `json:"agents"`
	Models              []AIAgentModel              `json:"models"`
	States              []AIAgentState              `json:"states"`
	Actions             []AIAgentAction             `json:"actions"`
	TrainingData        []AIAgentTrainingData       `json:"training_data"`
	MarketplaceListings []AIAgentMarketplaceListing `json:"marketplace_listings"`
	Params              Params                      `json:"params"`
}
