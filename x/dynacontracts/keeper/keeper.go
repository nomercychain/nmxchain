package keeper

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/nomercychain/nmxchain/x/dynacontracts/types"
)

// Keeper of the dynacontracts store
type Keeper struct {
	storeKey   sdk.StoreKey
	memKey     sdk.StoreKey
	cdc        codec.BinaryCodec
	paramstore paramtypes.Subspace

	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
	govKeeper     types.GovKeeper
}

// NewKeeper creates a new dynacontracts Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	govKeeper types.GovKeeper,
) Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		storeKey:      storeKey,
		memKey:        memKey,
		cdc:           cdc,
		paramstore:    ps,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
		govKeeper:     govKeeper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// SetContract sets a contract
func (k Keeper) SetContract(ctx sdk.Context, contract types.Contract) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.ContractKey, []byte(contract.ID)...)
	value := k.cdc.MustMarshal(&contract)
	store.Set(key, value)
}

// GetContract returns a contract by ID
func (k Keeper) GetContract(ctx sdk.Context, id string) (types.Contract, bool) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.ContractKey, []byte(id)...)
	value := store.Get(key)
	if value == nil {
		return types.Contract{}, false
	}

	var contract types.Contract
	k.cdc.MustUnmarshal(value, &contract)
	return contract, true
}

// DeleteContract deletes a contract
func (k Keeper) DeleteContract(ctx sdk.Context, id string) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.ContractKey, []byte(id)...)
	store.Delete(key)
}

// GetAllContracts returns all contracts
func (k Keeper) GetAllContracts(ctx sdk.Context) []types.Contract {
	var contracts []types.Contract
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ContractKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var contract types.Contract
		k.cdc.MustUnmarshal(iterator.Value(), &contract)
		contracts = append(contracts, contract)
	}

	return contracts
}

// SetContractCode sets a contract code
func (k Keeper) SetContractCode(ctx sdk.Context, code types.ContractCode) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.ContractCodeKey, []byte(fmt.Sprintf("%s-%d", code.ContractID, code.Version))...)
	value := k.cdc.MustMarshal(&code)
	store.Set(key, value)
}

// GetContractCode returns a contract code by ID and version
func (k Keeper) GetContractCode(ctx sdk.Context, id string, version uint64) (types.ContractCode, bool) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.ContractCodeKey, []byte(fmt.Sprintf("%s-%d", id, version))...)
	value := store.Get(key)
	if value == nil {
		return types.ContractCode{}, false
	}

	var code types.ContractCode
	k.cdc.MustUnmarshal(value, &code)
	return code, true
}

// GetLatestContractCode returns the latest version of a contract code
func (k Keeper) GetLatestContractCode(ctx sdk.Context, id string) (types.ContractCode, bool) {
	contract, found := k.GetContract(ctx, id)
	if !found {
		return types.ContractCode{}, false
	}

	return k.GetContractCode(ctx, id, contract.Version)
}

// SetContractState sets a contract state
func (k Keeper) SetContractState(ctx sdk.Context, state types.ContractState) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.ContractStateKey, []byte(fmt.Sprintf("%s-%s", state.ContractID, state.Key))...)
	value := k.cdc.MustMarshal(&state)
	store.Set(key, value)
}

// GetContractState returns a contract state by ID and key
func (k Keeper) GetContractState(ctx sdk.Context, id string, key string) (types.ContractState, bool) {
	store := ctx.KVStore(k.storeKey)
	storeKey := append(types.ContractStateKey, []byte(fmt.Sprintf("%s-%s", id, key))...)
	value := store.Get(storeKey)
	if value == nil {
		return types.ContractState{}, false
	}

	var state types.ContractState
	k.cdc.MustUnmarshal(value, &state)
	return state, true
}

// GetAllContractStates returns all states for a contract
func (k Keeper) GetAllContractStates(ctx sdk.Context, id string) []types.ContractState {
	var states []types.ContractState
	store := ctx.KVStore(k.storeKey)
	prefix := append(types.ContractStateKey, []byte(fmt.Sprintf("%s-", id))...)
	iterator := sdk.KVStorePrefixIterator(store, prefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var state types.ContractState
		k.cdc.MustUnmarshal(iterator.Value(), &state)
		states = append(states, state)
	}

	return states
}

// SetContractParameter sets a contract parameter
func (k Keeper) SetContractParameter(ctx sdk.Context, param types.ContractParameter) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.ContractParameterKey, []byte(fmt.Sprintf("%s-%s", param.ContractID, param.Name))...)
	value := k.cdc.MustMarshal(&param)
	store.Set(key, value)
}

// GetContractParameter returns a contract parameter by ID and name
func (k Keeper) GetContractParameter(ctx sdk.Context, id string, name string) (types.ContractParameter, bool) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.ContractParameterKey, []byte(fmt.Sprintf("%s-%s", id, name))...)
	value := store.Get(key)
	if value == nil {
		return types.ContractParameter{}, false
	}

	var param types.ContractParameter
	k.cdc.MustUnmarshal(value, &param)
	return param, true
}

// GetAllContractParameters returns all parameters for a contract
func (k Keeper) GetAllContractParameters(ctx sdk.Context, id string) []types.ContractParameter {
	var params []types.ContractParameter
	store := ctx.KVStore(k.storeKey)
	prefix := append(types.ContractParameterKey, []byte(fmt.Sprintf("%s-", id))...)
	iterator := sdk.KVStorePrefixIterator(store, prefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var param types.ContractParameter
		k.cdc.MustUnmarshal(iterator.Value(), &param)
		params = append(params, param)
	}

	return params
}

// SetAIModel sets an AI model
func (k Keeper) SetAIModel(ctx sdk.Context, model types.AIModel) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.ContractAIModelKey, []byte(model.ID)...)
	value := k.cdc.MustMarshal(&model)
	store.Set(key, value)
}

// GetAIModel returns an AI model by ID
func (k Keeper) GetAIModel(ctx sdk.Context, id string) (types.AIModel, bool) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.ContractAIModelKey, []byte(id)...)
	value := store.Get(key)
	if value == nil {
		return types.AIModel{}, false
	}

	var model types.AIModel
	k.cdc.MustUnmarshal(value, &model)
	return model, true
}

// GetAllAIModels returns all AI models
func (k Keeper) GetAllAIModels(ctx sdk.Context) []types.AIModel {
	var models []types.AIModel
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ContractAIModelKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var model types.AIModel
		k.cdc.MustUnmarshal(iterator.Value(), &model)
		models = append(models, model)
	}

	return models
}

// SetContractProposal sets a contract proposal
func (k Keeper) SetContractProposal(ctx sdk.Context, proposal types.ContractProposal) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.ContractProposalKey, sdk.Uint64ToBigEndian(proposal.ID)...)
	value := k.cdc.MustMarshal(&proposal)
	store.Set(key, value)
}

// GetContractProposal returns a contract proposal by ID
func (k Keeper) GetContractProposal(ctx sdk.Context, id uint64) (types.ContractProposal, bool) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.ContractProposalKey, sdk.Uint64ToBigEndian(id)...)
	value := store.Get(key)
	if value == nil {
		return types.ContractProposal{}, false
	}

	var proposal types.ContractProposal
	k.cdc.MustUnmarshal(value, &proposal)
	return proposal, true
}

// GetAllContractProposals returns all contract proposals
func (k Keeper) GetAllContractProposals(ctx sdk.Context) []types.ContractProposal {
	var proposals []types.ContractProposal
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ContractProposalKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var proposal types.ContractProposal
		k.cdc.MustUnmarshal(iterator.Value(), &proposal)
		proposals = append(proposals, proposal)
	}

	return proposals
}

// SetContractExecution records a contract execution
func (k Keeper) SetContractExecution(ctx sdk.Context, execution types.ContractExecution) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.ContractExecutionKey, []byte(execution.ID)...)
	value := k.cdc.MustMarshal(&execution)
	store.Set(key, value)
}

// GetContractExecution returns a contract execution by ID
func (k Keeper) GetContractExecution(ctx sdk.Context, id string) (types.ContractExecution, bool) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.ContractExecutionKey, []byte(id)...)
	value := store.Get(key)
	if value == nil {
		return types.ContractExecution{}, false
	}

	var execution types.ContractExecution
	k.cdc.MustUnmarshal(value, &execution)
	return execution, true
}

// SetExternalDataSource sets an external data source
func (k Keeper) SetExternalDataSource(ctx sdk.Context, source types.ExternalDataSource) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.ExternalDataSourceKey, []byte(source.ID)...)
	value := k.cdc.MustMarshal(&source)
	store.Set(key, value)
}

// GetExternalDataSource returns an external data source by ID
func (k Keeper) GetExternalDataSource(ctx sdk.Context, id string) (types.ExternalDataSource, bool) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.ExternalDataSourceKey, []byte(id)...)
	value := store.Get(key)
	if value == nil {
		return types.ExternalDataSource{}, false
	}

	var source types.ExternalDataSource
	k.cdc.MustUnmarshal(value, &source)
	return source, true
}

// GetAllExternalDataSources returns all external data sources
func (k Keeper) GetAllExternalDataSources(ctx sdk.Context) []types.ExternalDataSource {
	var sources []types.ExternalDataSource
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ExternalDataSourceKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var source types.ExternalDataSource
		k.cdc.MustUnmarshal(iterator.Value(), &source)
		sources = append(sources, source)
	}

	return sources
}

// CreateContract creates a new contract
func (k Keeper) CreateContract(ctx sdk.Context, creator sdk.AccAddress, name string, description string, language types.ContractLanguage, code []byte, aiEnabled bool) (string, error) {
	// Generate a unique ID for the contract
	id := fmt.Sprintf("%s-%d", creator.String(), ctx.BlockHeight())

	// Create the contract
	contract := types.Contract{
		ID:               id,
		Name:             name,
		Description:      description,
		Creator:          creator,
		Owner:            creator,
		Language:         language,
		Version:          1,
		Status:           types.ContractStatusActive,
		CreatedAt:        ctx.BlockTime(),
		UpdatedAt:        ctx.BlockTime(),
		AIEnabled:        aiEnabled,
		GovernanceEnabled: false,
	}

	// Create the contract code
	contractCode := types.ContractCode{
		ContractID: id,
		Version:    1,
		Code:       code,
		Checksum:   "", // In a real implementation, we would compute a checksum
		Metadata:   "",
	}

	// Store the contract and code
	k.SetContract(ctx, contract)
	k.SetContractCode(ctx, contractCode)

	return id, nil
}

// UpdateContract updates an existing contract
func (k Keeper) UpdateContract(ctx sdk.Context, id string, owner sdk.AccAddress, newCode []byte) error {
	// Get the contract
	contract, found := k.GetContract(ctx, id)
	if !found {
		return fmt.Errorf("contract not found: %s", id)
	}

	// Check if the caller is the owner
	if !contract.Owner.Equals(owner) {
		return fmt.Errorf("only the owner can update the contract")
	}

	// Increment the version
	contract.Version++
	contract.UpdatedAt = ctx.BlockTime()

	// Create the new contract code
	contractCode := types.ContractCode{
		ContractID: id,
		Version:    contract.Version,
		Code:       newCode,
		Checksum:   "", // In a real implementation, we would compute a checksum
		Metadata:   "",
	}

	// Store the updated contract and code
	k.SetContract(ctx, contract)
	k.SetContractCode(ctx, contractCode)

	return nil
}

// ExecuteContract executes a contract
func (k Keeper) ExecuteContract(ctx sdk.Context, caller sdk.AccAddress, contractID string, method string, params json.RawMessage) (json.RawMessage, error) {
	// Get the contract
	contract, found := k.GetContract(ctx, contractID)
	if !found {
		return nil, fmt.Errorf("contract not found: %s", contractID)
	}

	// Check if the contract is active
	if contract.Status != types.ContractStatusActive {
		return nil, fmt.Errorf("contract is not active")
	}

	// Get the latest contract code
	code, found := k.GetLatestContractCode(ctx, contractID)
	if !found {
		return nil, fmt.Errorf("contract code not found")
	}

	// In a real implementation, we would execute the contract code here
	// For now, we'll just record the execution and return a dummy result

	// Record the execution
	execution := types.ContractExecution{
		ID:         fmt.Sprintf("%s-%s-%d", contractID, method, ctx.BlockHeight()),
		ContractID: contractID,
		Caller:     caller,
		Method:     method,
		Params:     params,
		Result:     json.RawMessage(`{"status": "success"}`),
		GasUsed:    1000, // Dummy value
		Timestamp:  ctx.BlockTime(),
		Success:    true,
	}

	k.SetContractExecution(ctx, execution)

	return execution.Result, nil
}

// UpdateContractParameters updates contract parameters using AI
func (k Keeper) UpdateContractParameters(ctx sdk.Context) {
	// Get all contracts
	contracts := k.GetAllContracts(ctx)

	for _, contract := range contracts {
		// Skip contracts that don't have AI enabled
		if !contract.AIEnabled {
			continue
		}

		// Get the AI model for the contract
		if contract.AIModelID == "" {
			continue
		}

		model, found := k.GetAIModel(ctx, contract.AIModelID)
		if !found {
			continue
		}

		// Get all parameters for the contract
		params := k.GetAllContractParameters(ctx, contract.ID)

		// Update AI-controlled parameters
		for _, param := range params {
			if !param.AIControlled {
				continue
			}

			// In a real implementation, we would use the AI model to update the parameter
			// For now, we'll just update the timestamp
			param.UpdatedAt = ctx.BlockTime()
			k.SetContractParameter(ctx, param)
		}

		// Update the contract's last updated timestamp
		contract.UpdatedAt = ctx.BlockTime()
		k.SetContract(ctx, contract)
	}
}

// ProcessContractProposals processes pending contract proposals
func (k Keeper) ProcessContractProposals(ctx sdk.Context) {
	// Get all contract proposals
	proposals := k.GetAllContractProposals(ctx)

	for _, proposal := range proposals {
		// Skip proposals that are not pending or have already ended
		if proposal.Status != "pending" || proposal.EndTime.Before(ctx.BlockTime()) {
			continue
		}

		// Check if the proposal has passed
		if proposal.VotesYes.GT(proposal.VotesNo) {
			// Update the contract
			contract, found := k.GetContract(ctx, proposal.ContractID)
			if !found {
				continue
			}

			// Update the contract code if provided
			if len(proposal.NewCode) > 0 {
				contract.Version++
				contractCode := types.ContractCode{
					ContractID: contract.ID,
					Version:    contract.Version,
					Code:       proposal.NewCode,
					Checksum:   "", // In a real implementation, we would compute a checksum
					Metadata:   "",
				}
				k.SetContractCode(ctx, contractCode)
			}

			// Update the contract parameters if provided
			for _, param := range proposal.NewParams {
				k.SetContractParameter(ctx, param)
			}

			// Update the contract
			contract.UpdatedAt = ctx.BlockTime()
			k.SetContract(ctx, contract)

			// Update the proposal status
			proposal.Status = "approved"
		} else {
			// Update the proposal status
			proposal.Status = "rejected"
		}

		// Save the updated proposal
		k.SetContractProposal(ctx, proposal)
	}
}