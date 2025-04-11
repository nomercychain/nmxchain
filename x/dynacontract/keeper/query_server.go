package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/nomercychain/nmxchain/x/dynacontract/types"
)

type queryServer struct {
	Keeper
}

// NewQueryServerImpl returns an implementation of the QueryServer interface
// for the provided Keeper.
func NewQueryServerImpl(k Keeper) types.QueryServer {
	return &queryServer{Keeper: k}
}

var _ types.QueryServer = queryServer{}

// DynaContract returns information about a specific dynamic contract
func (q queryServer) DynaContract(goCtx context.Context, req *types.QueryDynaContractRequest) (*types.QueryDynaContractResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	contract, found := q.GetDynaContract(ctx, req.Id)
	if !found {
		return nil, status.Errorf(codes.NotFound, "contract with ID %s not found", req.Id)
	}

	return &types.QueryDynaContractResponse{
		Contract: contract,
	}, nil
}

// DynaContracts returns all dynamic contracts
func (q queryServer) DynaContracts(goCtx context.Context, req *types.QueryDynaContractsRequest) (*types.QueryDynaContractsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var contracts []types.DynaContract
	var pageRes *query.PageResponse
	var err error

	store := ctx.KVStore(q.storeKey)
	contractStore := prefix.NewStore(store, types.KeyPrefix(types.DynaContractKey))

	pageRes, err = query.Paginate(contractStore, req.Pagination, func(key []byte, value []byte) error {
		var contract types.DynaContract
		if err := q.cdc.Unmarshal(value, &contract); err != nil {
			return err
		}

		contracts = append(contracts, contract)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryDynaContractsResponse{
		Contracts:  contracts,
		Pagination: pageRes,
	}, nil
}

// DynaContractExecutions returns executions for a specific dynamic contract
func (q queryServer) DynaContractExecutions(goCtx context.Context, req *types.QueryDynaContractExecutionsRequest) (*types.QueryDynaContractExecutionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the contract exists
	_, found := q.GetDynaContract(ctx, req.ContractId)
	if !found {
		return nil, status.Errorf(codes.NotFound, "contract with ID %s not found", req.ContractId)
	}

	var executions []types.DynaContractExecution
	var pageRes *query.PageResponse
	var err error

	store := ctx.KVStore(q.storeKey)
	executionStore := prefix.NewStore(store, types.KeyPrefix(types.DynaContractExecutionByContractKey))
	prefixKey := []byte(req.ContractId)
	executionByContractStore := prefix.NewStore(executionStore, prefixKey)

	pageRes, err = query.Paginate(executionByContractStore, req.Pagination, func(key []byte, value []byte) error {
		var executionID string
		executionID = string(key)

		execution, found := q.GetDynaContractExecution(ctx, executionID)
		if !found {
			return nil
		}

		executions = append(executions, execution)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryDynaContractExecutionsResponse{
		Executions: executions,
		Pagination: pageRes,
	}, nil
}

// DynaContractExecution returns a specific dynamic contract execution
func (q queryServer) DynaContractExecution(goCtx context.Context, req *types.QueryDynaContractExecutionRequest) (*types.QueryDynaContractExecutionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	execution, found := q.GetDynaContractExecution(ctx, req.Id)
	if !found {
		return nil, status.Errorf(codes.NotFound, "execution with ID %s not found", req.Id)
	}

	return &types.QueryDynaContractExecutionResponse{
		Execution: execution,
	}, nil
}

// DynaContractTemplates returns all dynamic contract templates
func (q queryServer) DynaContractTemplates(goCtx context.Context, req *types.QueryDynaContractTemplatesRequest) (*types.QueryDynaContractTemplatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var templates []types.DynaContractTemplate
	var pageRes *query.PageResponse
	var err error

	store := ctx.KVStore(q.storeKey)
	templateStore := prefix.NewStore(store, types.KeyPrefix(types.DynaContractTemplateKey))

	pageRes, err = query.Paginate(templateStore, req.Pagination, func(key []byte, value []byte) error {
		var template types.DynaContractTemplate
		if err := q.cdc.Unmarshal(value, &template); err != nil {
			return err
		}

		templates = append(templates, template)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryDynaContractTemplatesResponse{
		Templates:  templates,
		Pagination: pageRes,
	}, nil
}

// DynaContractTemplate returns a specific dynamic contract template
func (q queryServer) DynaContractTemplate(goCtx context.Context, req *types.QueryDynaContractTemplateRequest) (*types.QueryDynaContractTemplateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	template, found := q.GetDynaContractTemplate(ctx, req.Id)
	if !found {
		return nil, status.Errorf(codes.NotFound, "template with ID %s not found", req.Id)
	}

	return &types.QueryDynaContractTemplateResponse{
		Template: template,
	}, nil
}

// DynaContractLearningData returns learning data for a specific dynamic contract
func (q queryServer) DynaContractLearningData(goCtx context.Context, req *types.QueryDynaContractLearningDataRequest) (*types.QueryDynaContractLearningDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the contract exists
	_, found := q.GetDynaContract(ctx, req.ContractId)
	if !found {
		return nil, status.Errorf(codes.NotFound, "contract with ID %s not found", req.ContractId)
	}

	var learningData []types.DynaContractLearningData
	var pageRes *query.PageResponse
	var err error

	store := ctx.KVStore(q.storeKey)
	learningDataStore := prefix.NewStore(store, types.KeyPrefix(types.DynaContractLearningDataByContractKey))
	prefixKey := []byte(req.ContractId)
	learningDataByContractStore := prefix.NewStore(learningDataStore, prefixKey)

	pageRes, err = query.Paginate(learningDataByContractStore, req.Pagination, func(key []byte, value []byte) error {
		var dataID string
		dataID = string(key)

		data, found := q.GetDynaContractLearningData(ctx, dataID)
		if !found {
			return nil
		}

		learningData = append(learningData, data)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryDynaContractLearningDataResponse{
		LearningData: learningData,
		Pagination:   pageRes,
	}, nil
}

// DynaContractPermissions returns permissions for a specific dynamic contract
func (q queryServer) DynaContractPermissions(goCtx context.Context, req *types.QueryDynaContractPermissionsRequest) (*types.QueryDynaContractPermissionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the contract exists
	_, found := q.GetDynaContract(ctx, req.ContractId)
	if !found {
		return nil, status.Errorf(codes.NotFound, "contract with ID %s not found", req.ContractId)
	}

	var permissions []types.DynaContractPermission
	var pageRes *query.PageResponse
	var err error

	store := ctx.KVStore(q.storeKey)
	permissionStore := prefix.NewStore(store, types.KeyPrefix(types.DynaContractPermissionByContractKey))
	prefixKey := []byte(req.ContractId)
	permissionByContractStore := prefix.NewStore(permissionStore, prefixKey)

	pageRes, err = query.Paginate(permissionByContractStore, req.Pagination, func(key []byte, value []byte) error {
		var permission types.DynaContractPermission
		if err := q.cdc.Unmarshal(value, &permission); err != nil {
			return err
		}

		permissions = append(permissions, permission)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryDynaContractPermissionsResponse{
		Permissions: permissions,
		Pagination:  pageRes,
	}, nil
}

// DynaContractsByOwner returns dynamic contracts owned by a specific address
func (q queryServer) DynaContractsByOwner(goCtx context.Context, req *types.QueryDynaContractsByOwnerRequest) (*types.QueryDynaContractsByOwnerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var contracts []types.DynaContract
	var pageRes *query.PageResponse
	var err error

	store := ctx.KVStore(q.storeKey)
	contractByOwnerStore := prefix.NewStore(store, types.KeyPrefix(types.DynaContractByOwnerKey))
	prefixKey := []byte(req.Owner)
	contractByOwnerStore = prefix.NewStore(contractByOwnerStore, prefixKey)

	pageRes, err = query.Paginate(contractByOwnerStore, req.Pagination, func(key []byte, value []byte) error {
		var contractID string
		contractID = string(key)

		contract, found := q.GetDynaContract(ctx, contractID)
		if !found {
			return nil
		}

		contracts = append(contracts, contract)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryDynaContractsByOwnerResponse{
		Contracts:  contracts,
		Pagination: pageRes,
	}, nil
}

// DynaContractsByAgent returns dynamic contracts associated with a specific AI agent
func (q queryServer) DynaContractsByAgent(goCtx context.Context, req *types.QueryDynaContractsByAgentRequest) (*types.QueryDynaContractsByAgentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var contracts []types.DynaContract
	var pageRes *query.PageResponse
	var err error

	store := ctx.KVStore(q.storeKey)
	contractByAgentStore := prefix.NewStore(store, types.KeyPrefix(types.DynaContractByAgentKey))
	prefixKey := []byte(req.AgentId)
	contractByAgentStore = prefix.NewStore(contractByAgentStore, prefixKey)

	pageRes, err = query.Paginate(contractByAgentStore, req.Pagination, func(key []byte, value []byte) error {
		var contractID string
		contractID = string(key)

		contract, found := q.GetDynaContract(ctx, contractID)
		if !found {
			return nil
		}

		contracts = append(contracts, contract)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryDynaContractsByAgentResponse{
		Contracts:  contracts,
		Pagination: pageRes,
	}, nil
}

// DynaContractsByTags returns dynamic contracts with specific tags
func (q queryServer) DynaContractsByTags(goCtx context.Context, req *types.QueryDynaContractsByTagsRequest) (*types.QueryDynaContractsByTagsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if len(req.Tags) == 0 {
		return nil, status.Error(codes.InvalidArgument, "tags cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var contracts []types.DynaContract
	var contractIDs = make(map[string]bool)

	// For each tag, get all contracts with that tag
	for _, tag := range req.Tags {
		store := ctx.KVStore(q.storeKey)
		contractByTagStore := prefix.NewStore(store, types.KeyPrefix(types.DynaContractByTagKey))
		prefixKey := []byte(tag)
		contractByTagStore = prefix.NewStore(contractByTagStore, prefixKey)

		iterator := contractByTagStore.Iterator(nil, nil)
		defer iterator.Close()

		for ; iterator.Valid(); iterator.Next() {
			contractID := string(iterator.Key())
			contractIDs[contractID] = true
		}
	}

	// If match_all is true, we need to filter out contracts that don't have all tags
	if req.MatchAll && len(req.Tags) > 1 {
		// For each contract, check if it has all tags
		for contractID := range contractIDs {
			hasAllTags := true
			for _, tag := range req.Tags {
				store := ctx.KVStore(q.storeKey)
				contractByTagStore := prefix.NewStore(store, types.KeyPrefix(types.DynaContractByTagKey))
				prefixKey := []byte(tag)
				contractByTagStore = prefix.NewStore(contractByTagStore, prefixKey)

				if !contractByTagStore.Has([]byte(contractID)) {
					hasAllTags = false
					break
				}
			}

			if !hasAllTags {
				delete(contractIDs, contractID)
			}
		}
	}

	// Get the contracts
	for contractID := range contractIDs {
		contract, found := q.GetDynaContract(ctx, contractID)
		if found {
			contracts = append(contracts, contract)
		}
	}

	// Apply pagination
	start, end := client.Paginate(len(contracts), req.Pagination.Offset, req.Pagination.Limit, 100)
	if start < 0 || end < 0 {
		contracts = []types.DynaContract{}
	} else {
		contracts = contracts[start:end]
	}

	return &types.QueryDynaContractsByTagsResponse{
		Contracts: contracts,
		Pagination: &query.PageResponse{
			NextKey: nil,
			Total:   uint64(len(contractIDs)),
		},
	}, nil
}