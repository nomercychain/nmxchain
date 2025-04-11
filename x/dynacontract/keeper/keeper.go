package keeper

import (
	"fmt"
	"crypto/sha256"
	"encoding/hex"

	"github.com/tendermint/tendermint/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/nomercychain/nmxchain/x/dynacontract/types"
	"github.com/nomercychain/nmxchain/x/deai/types" // For AI agent integration
)

// Keeper of the dynacontract store
type Keeper struct {
	cdc        codec.BinaryCodec
	storeKey   sdk.StoreKey
	memKey     sdk.StoreKey
	paramstore paramtypes.Subspace
	bankKeeper types.BankKeeper
	deaiKeeper types.DeAIKeeper
}

// NewKeeper creates a new dynacontract Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,
	bankKeeper types.BankKeeper,
	deaiKeeper types.DeAIKeeper,
) *Keeper {
	// Set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
		bankKeeper: bankKeeper,
		deaiKeeper: deaiKeeper,
	}
}

// Logger returns a module-specific logger
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetDynaContract returns a dynamic contract by ID
func (k Keeper) GetDynaContract(ctx sdk.Context, id string) (contract types.DynaContract, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.DynaContractKey(id)
	
	value := store.Get(key)
	if value == nil {
		return contract, false
	}
	
	k.cdc.MustUnmarshal(value, &contract)
	return contract, true
}

// SetDynaContract sets a dynamic contract in the store
func (k Keeper) SetDynaContract(ctx sdk.Context, contract types.DynaContract) {
	store := ctx.KVStore(k.storeKey)
	key := types.DynaContractKey(contract.Id)
	value := k.cdc.MustMarshal(&contract)
	
	store.Set(key, value)
}

// DeleteDynaContract deletes a dynamic contract from the store
func (k Keeper) DeleteDynaContract(ctx sdk.Context, id string) {
	store := ctx.KVStore(k.storeKey)
	key := types.DynaContractKey(id)
	store.Delete(key)
}

// GetAllDynaContracts returns all dynamic contracts
func (k Keeper) GetAllDynaContracts(ctx sdk.Context) (contracts []types.DynaContract) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.DynaContractKeyPrefix)
	defer iterator.Close()
	
	for ; iterator.Valid(); iterator.Next() {
		var contract types.DynaContract
		k.cdc.MustUnmarshal(iterator.Value(), &contract)
		contracts = append(contracts, contract)
	}
	
	return contracts
}

// GetDynaContractsByOwner returns all dynamic contracts owned by a specific address
func (k Keeper) GetDynaContractsByOwner(ctx sdk.Context, owner sdk.AccAddress) (contracts []types.DynaContract) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.DynaContractKeyPrefix)
	defer iterator.Close()
	
	for ; iterator.Valid(); iterator.Next() {
		var contract types.DynaContract
		k.cdc.MustUnmarshal(iterator.Value(), &contract)
		
		if contract.Owner == owner.String() {
			contracts = append(contracts, contract)
		}
	}
	
	return contracts
}

// GetDynaContractsByAgent returns all dynamic contracts associated with a specific AI agent
func (k Keeper) GetDynaContractsByAgent(ctx sdk.Context, agentID string) (contracts []types.DynaContract) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.DynaContractKeyPrefix)
	defer iterator.Close()
	
	for ; iterator.Valid(); iterator.Next() {
		var contract types.DynaContract
		k.cdc.MustUnmarshal(iterator.Value(), &contract)
		
		if contract.AgentId == agentID {
			contracts = append(contracts, contract)
		}
	}
	
	return contracts
}

// GetDynaContractsByTags returns all dynamic contracts with specific tags
func (k Keeper) GetDynaContractsByTags(ctx sdk.Context, tags []string, matchAll bool) (contracts []types.DynaContract) {
	allContracts := k.GetAllDynaContracts(ctx)
	
	for _, contract := range allContracts {
		if matchAll {
			// Contract must have all specified tags
			hasAllTags := true
			for _, tag := range tags {
				found := false
				for _, contractTag := range contract.Tags {
					if contractTag == tag {
						found = true
						break
					}
				}
				if !found {
					hasAllTags = false
					break
				}
			}
			if hasAllTags {
				contracts = append(contracts, contract)
			}
		} else {
			// Contract must have at least one of the specified tags
			for _, tag := range tags {
				for _, contractTag := range contract.Tags {
					if contractTag == tag {
						contracts = append(contracts, contract)
						goto nextContract
					}
				}
			}
		}
		nextContract:
	}
	
	return contracts
}

// GetDynaContractExecution returns a dynamic contract execution by ID
func (k Keeper) GetDynaContractExecution(ctx sdk.Context, id string) (execution types.DynaContractExecution, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.DynaContractExecutionKey(id)
	
	value := store.Get(key)
	if value == nil {
		return execution, false
	}
	
	k.cdc.MustUnmarshal(value, &execution)
	return execution, true
}

// SetDynaContractExecution sets a dynamic contract execution in the store
func (k Keeper) SetDynaContractExecution(ctx sdk.Context, execution types.DynaContractExecution) {
	store := ctx.KVStore(k.storeKey)
	key := types.DynaContractExecutionKey(execution.Id)
	value := k.cdc.MustMarshal(&execution)
	
	store.Set(key, value)
}

// GetDynaContractExecutionsByContract returns all executions for a specific dynamic contract
func (k Keeper) GetDynaContractExecutionsByContract(ctx sdk.Context, contractID string) (executions []types.DynaContractExecution) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.DynaContractExecutionKeyPrefix)
	defer iterator.Close()
	
	for ; iterator.Valid(); iterator.Next() {
		var execution types.DynaContractExecution
		k.cdc.MustUnmarshal(iterator.Value(), &execution)
		
		if execution.ContractId == contractID {
			executions = append(executions, execution)
		}
	}
	
	return executions
}

// GetDynaContractTemplate returns a dynamic contract template by ID
func (k Keeper) GetDynaContractTemplate(ctx sdk.Context, id string) (template types.DynaContractTemplate, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.DynaContractTemplateKey(id)
	
	value := store.Get(key)
	if value == nil {
		return template, false
	}
	
	k.cdc.MustUnmarshal(value, &template)
	return template, true
}

// SetDynaContractTemplate sets a dynamic contract template in the store
func (k Keeper) SetDynaContractTemplate(ctx sdk.Context, template types.DynaContractTemplate) {
	store := ctx.KVStore(k.storeKey)
	key := types.DynaContractTemplateKey(template.Id)
	value := k.cdc.MustMarshal(&template)
	
	store.Set(key, value)
}

// GetAllDynaContractTemplates returns all dynamic contract templates
func (k Keeper) GetAllDynaContractTemplates(ctx sdk.Context) (templates []types.DynaContractTemplate) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.DynaContractTemplateKeyPrefix)
	defer iterator.Close()
	
	for ; iterator.Valid(); iterator.Next() {
		var template types.DynaContractTemplate
		k.cdc.MustUnmarshal(iterator.Value(), &template)
		templates = append(templates, template)
	}
	
	return templates
}

// GetDynaContractLearningData returns learning data for a specific dynamic contract
func (k Keeper) GetDynaContractLearningData(ctx sdk.Context, id string) (data types.DynaContractLearningData, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.DynaContractLearningDataKey(id)
	
	value := store.Get(key)
	if value == nil {
		return data, false
	}
	
	k.cdc.MustUnmarshal(value, &data)
	return data, true
}

// SetDynaContractLearningData sets learning data for a dynamic contract in the store
func (k Keeper) SetDynaContractLearningData(ctx sdk.Context, data types.DynaContractLearningData) {
	store := ctx.KVStore(k.storeKey)
	key := types.DynaContractLearningDataKey(data.Id)
	value := k.cdc.MustMarshal(&data)
	
	store.Set(key, value)
}

// GetDynaContractLearningDataByContract returns all learning data for a specific dynamic contract
func (k Keeper) GetDynaContractLearningDataByContract(ctx sdk.Context, contractID string) (dataList []types.DynaContractLearningData) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.DynaContractLearningDataKeyPrefix)
	defer iterator.Close()
	
	for ; iterator.Valid(); iterator.Next() {
		var data types.DynaContractLearningData
		k.cdc.MustUnmarshal(iterator.Value(), &data)
		
		if data.ContractId == contractID {
			dataList = append(dataList, data)
		}
	}
	
	return dataList
}

// GetDynaContractPermission returns a permission for a specific dynamic contract and address
func (k Keeper) GetDynaContractPermission(ctx sdk.Context, contractID string, address string, permissionType string) (permission types.DynaContractPermission, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.DynaContractPermissionKey(contractID, address, permissionType)
	
	value := store.Get(key)
	if value == nil {
		return permission, false
	}
	
	k.cdc.MustUnmarshal(value, &permission)
	return permission, true
}

// SetDynaContractPermission sets a permission for a dynamic contract in the store
func (k Keeper) SetDynaContractPermission(ctx sdk.Context, permission types.DynaContractPermission) {
	store := ctx.KVStore(k.storeKey)
	key := types.DynaContractPermissionKey(permission.ContractId, permission.Address, permission.PermissionType)
	value := k.cdc.MustMarshal(&permission)
	
	store.Set(key, value)
}

// DeleteDynaContractPermission deletes a permission for a dynamic contract from the store
func (k Keeper) DeleteDynaContractPermission(ctx sdk.Context, contractID string, address string, permissionType string) {
	store := ctx.KVStore(k.storeKey)
	key := types.DynaContractPermissionKey(contractID, address, permissionType)
	store.Delete(key)
}

// GetDynaContractPermissionsByContract returns all permissions for a specific dynamic contract
func (k Keeper) GetDynaContractPermissionsByContract(ctx sdk.Context, contractID string) (permissions []types.DynaContractPermission) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.DynaContractPermissionKeyPrefixByContract(contractID))
	defer iterator.Close()
	
	for ; iterator.Valid(); iterator.Next() {
		var permission types.DynaContractPermission
		k.cdc.MustUnmarshal(iterator.Value(), &permission)
		permissions = append(permissions, permission)
	}
	
	return permissions
}

// HasPermission checks if an address has a specific permission for a dynamic contract
func (k Keeper) HasPermission(ctx sdk.Context, contractID string, address string, permissionType string) bool {
	permission, found := k.GetDynaContractPermission(ctx, contractID, address, permissionType)
	if !found {
		return false
	}
	
	// Check if the permission has expired
	if !permission.ExpiresAt.IsZero() && permission.ExpiresAt.Before(ctx.BlockTime()) {
		return false
	}
	
	return true
}

// CalculateCodeHash calculates the hash of the contract code
func (k Keeper) CalculateCodeHash(code []byte) string {
	hash := sha256.Sum256(code)
	return hex.EncodeToString(hash[:])
}

// ExecuteDynaContract executes a dynamic contract
func (k Keeper) ExecuteDynaContract(ctx sdk.Context, contractID string, caller sdk.AccAddress, input []byte, gasLimit uint64) (output []byte, gasUsed uint64, err error) {
	// Get the contract
	contract, found := k.GetDynaContract(ctx, contractID)
	if !found {
		return nil, 0, types.ErrContractNotFound
	}
	
	// Check if the contract is active
	if contract.Status != types.DynaContractStatusActive {
		return nil, 0, types.ErrContractNotActive
	}
	
	// Check if the caller has permission to execute the contract
	if !k.HasPermission(ctx, contractID, caller.String(), types.PermissionTypeExecute) && contract.Owner != caller.String() {
		return nil, 0, types.ErrUnauthorized
	}
	
	// Get the contract state
	stateBefore := contract.State
	
	// If the contract has an associated AI agent, use it for execution
	if contract.AgentId != "" {
		// Get the AI agent
		agent, found := k.deaiKeeper.GetAIAgent(ctx, contract.AgentId)
		if !found {
			return nil, 0, types.ErrAgentNotFound
		}
		
		// Check if the agent is active
		if agent.Status != deaitypes.AIAgentStatusActive {
			return nil, 0, types.ErrAgentNotActive
		}
		
		// Execute the contract using the AI agent
		// This is a simplified implementation
		// In a real implementation, we would:
		// 1. Prepare the execution environment
		// 2. Load the contract code and ABI
		// 3. Execute the contract code with the input
		// 4. Update the contract state based on the execution result
		// 5. Return the output and gas used
		
		// For now, we'll just simulate the execution
		output = []byte(`{"status": "success", "result": "Contract executed successfully"}`)
		gasUsed = 1000 // Dummy value
		
		// Update the contract state
		contract.State = []byte(`{"updated": true, "last_execution": "` + ctx.BlockTime().String() + `"}`)
		contract.ExecutionCount++
		contract.UpdatedAt = ctx.BlockTime()
		
		// Save the updated contract
		k.SetDynaContract(ctx, contract)
		
		// Create an execution record
		executionID := fmt.Sprintf("%s-%d", contractID, contract.ExecutionCount)
		execution := types.DynaContractExecution{
			Id:          executionID,
			ContractId:  contractID,
			Caller:      caller.String(),
			Input:       input,
			Output:      output,
			Status:      "success",
			GasUsed:     gasUsed,
			Error:       "",
			Timestamp:   ctx.BlockTime(),
			StateBefore: stateBefore,
			StateAfter:  contract.State,
			Metadata:    []byte{},
		}
		
		// Save the execution record
		k.SetDynaContractExecution(ctx, execution)
		
		// Emit an event
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeExecuteDynaContract,
				sdk.NewAttribute(types.AttributeKeyContractID, contractID),
				sdk.NewAttribute(types.AttributeKeyCaller, caller.String()),
				sdk.NewAttribute(types.AttributeKeyExecutionID, executionID),
				sdk.NewAttribute(types.AttributeKeyGasUsed, fmt.Sprintf("%d", gasUsed)),
			),
		)
		
		return output, gasUsed, nil
	}
	
	// If the contract doesn't have an associated AI agent, execute it directly
	// This is a simplified implementation
	// In a real implementation, we would:
	// 1. Prepare the execution environment
	// 2. Load the contract code and ABI
	// 3. Execute the contract code with the input
	// 4. Update the contract state based on the execution result
	// 5. Return the output and gas used
	
	// For now, we'll just simulate the execution
	output = []byte(`{"status": "success", "result": "Contract executed successfully"}`)
	gasUsed = 1000 // Dummy value
	
	// Update the contract state
	contract.State = []byte(`{"updated": true, "last_execution": "` + ctx.BlockTime().String() + `"}`)
	contract.ExecutionCount++
	contract.UpdatedAt = ctx.BlockTime()
	
	// Save the updated contract
	k.SetDynaContract(ctx, contract)
	
	// Create an execution record
	executionID := fmt.Sprintf("%s-%d", contractID, contract.ExecutionCount)
	execution := types.DynaContractExecution{
		Id:          executionID,
		ContractId:  contractID,
		Caller:      caller.String(),
		Input:       input,
		Output:      output,
		Status:      "success",
		GasUsed:     gasUsed,
		Error:       "",
		Timestamp:   ctx.BlockTime(),
		StateBefore: stateBefore,
		StateAfter:  contract.State,
		Metadata:    []byte{},
	}
	
	// Save the execution record
	k.SetDynaContractExecution(ctx, execution)
	
	// Emit an event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeExecuteDynaContract,
			sdk.NewAttribute(types.AttributeKeyContractID, contractID),
			sdk.NewAttribute(types.AttributeKeyCaller, caller.String()),
			sdk.NewAttribute(types.AttributeKeyExecutionID, executionID),
			sdk.NewAttribute(types.AttributeKeyGasUsed, fmt.Sprintf("%d", gasUsed)),
		),
	)
	
	return output, gasUsed, nil
}

// GetParams gets the dynacontract module parameters
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramstore.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the dynacontract module parameters
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// MaxContractSize returns the maximum contract size parameter
func (k Keeper) MaxContractSize(ctx sdk.Context) uint64 {
	var size uint64
	k.paramstore.Get(ctx, types.KeyMaxContractSize, &size)
	return size
}

// MaxContractGas returns the maximum contract gas parameter
func (k Keeper) MaxContractGas(ctx sdk.Context) uint64 {
	var gas uint64
	k.paramstore.Get(ctx, types.KeyMaxContractGas, &gas)
	return gas
}

// MaxLearningDataSize returns the maximum learning data size parameter
func (k Keeper) MaxLearningDataSize(ctx sdk.Context) uint64 {
	var size uint64
	k.paramstore.Get(ctx, types.KeyMaxLearningDataSize, &size)
	return size
}

// MaxMetadataSize returns the maximum metadata size parameter
func (k Keeper) MaxMetadataSize(ctx sdk.Context) uint64 {
	var size uint64
	k.paramstore.Get(ctx, types.KeyMaxMetadataSize, &size)
	return size
}

// MinContractDeposit returns the minimum contract deposit parameter
func (k Keeper) MinContractDeposit(ctx sdk.Context) sdk.Coin {
	var deposit sdk.Coin
	k.paramstore.Get(ctx, types.KeyMinContractDeposit, &deposit)
	return deposit
}

// ExecutionFeeRate returns the execution fee rate parameter
func (k Keeper) ExecutionFeeRate(ctx sdk.Context) sdk.Dec {
	var rate sdk.Dec
	k.paramstore.Get(ctx, types.KeyExecutionFeeRate, &rate)
	return rate
}