package types

import (
	"encoding/json"
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:              DefaultParams(),
		DataSources:         []DataSource{},
		OracleQueries:       []OracleQuery{},
		OracleResponses:     []OracleResponse{},
		AIModels:            []AIModel{},
		DataSourceRanks:     []DataSourceRank{},
		MisinformationList:  []Misinformation{},
		VerificationTasks:   []VerificationTask{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Validate params
	if err := gs.Params.Validate(); err != nil {
		return err
	}

	// Validate data sources
	dataSourceIDs := make(map[string]bool)
	for _, dataSource := range gs.DataSources {
		if dataSourceIDs[dataSource.ID] {
			return fmt.Errorf("duplicate data source ID: %s", dataSource.ID)
		}
		dataSourceIDs[dataSource.ID] = true
	}

	// Validate oracle queries
	queryIDs := make(map[string]bool)
	for _, query := range gs.OracleQueries {
		if queryIDs[query.ID] {
			return fmt.Errorf("duplicate oracle query ID: %s", query.ID)
		}
		queryIDs[query.ID] = true
	}

	// Validate oracle responses
	responseIDs := make(map[string]bool)
	for _, response := range gs.OracleResponses {
		if responseIDs[response.ID] {
			return fmt.Errorf("duplicate oracle response ID: %s", response.ID)
		}
		responseIDs[response.ID] = true

		// Check that the query exists for this response
		if !queryIDs[response.QueryID] {
			return fmt.Errorf("oracle response %s references non-existent query: %s", response.ID, response.QueryID)
		}
	}

	// Validate AI models
	modelIDs := make(map[string]bool)
	for _, model := range gs.AIModels {
		if modelIDs[model.ID] {
			return fmt.Errorf("duplicate AI model ID: %s", model.ID)
		}
		modelIDs[model.ID] = true
	}

	// Validate data source ranks
	for _, rank := range gs.DataSourceRanks {
		if !dataSourceIDs[rank.SourceID] {
			return fmt.Errorf("data source rank references non-existent data source: %s", rank.SourceID)
		}
	}

	// Validate misinformation
	misinfoIDs := make(map[string]bool)
	for _, misinfo := range gs.MisinformationList {
		if misinfoIDs[misinfo.ID] {
			return fmt.Errorf("duplicate misinformation ID: %s", misinfo.ID)
		}
		misinfoIDs[misinfo.ID] = true

		// Check that the AI model exists if specified
		if misinfo.DetectedBy != "" && !modelIDs[misinfo.DetectedBy] {
			return fmt.Errorf("misinformation %s references non-existent AI model: %s", misinfo.ID, misinfo.DetectedBy)
		}
	}

	// Validate verification tasks
	taskIDs := make(map[string]bool)
	for _, task := range gs.VerificationTasks {
		if taskIDs[task.ID] {
			return fmt.Errorf("duplicate verification task ID: %s", task.ID)
		}
		taskIDs[task.ID] = true
	}

	return nil
}

// GetGenesisStateFromAppState returns x/truthgpt GenesisState given raw application
// genesis state.
func GetGenesisStateFromAppState(appState map[string]json.RawMessage) *GenesisState {
	var genesisState GenesisState

	if appState[ModuleName] != nil {
		ModuleCdc.MustUnmarshalJSON(appState[ModuleName], &genesisState)
	}

	return &genesisState
}