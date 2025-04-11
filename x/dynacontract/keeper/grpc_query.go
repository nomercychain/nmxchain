package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/nomercychain/nmxchain/x/dynacontract/types"
)

var _ types.QueryServer = Keeper{}

// DynaContract returns information about a specific dynamic contract
func (k Keeper) DynaContract(c context.Context, req *types.QueryDynaContractRequest) (*types.QueryDynaContractResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	contract, found := k.GetDynaContract(ctx, req.Id)
	if !found {
		return nil, status.Error(codes.NotFound, "dynamic contract not found")
	}

	return &types.QueryDynaContractResponse{
		Contract: contract,
	}, nil
}

// DynaContracts returns all dynamic contracts
func (k Keeper) DynaContracts(c context.Context, req *types.QueryDynaContractsRequest) (*types.QueryDynaContractsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	contracts := k.GetAllDynaContracts(ctx)
	var pageRes *query.PageResponse
	var err error

	// Apply pagination
	if req.Pagination != nil {
		start, end := query.Paginate(len(contracts), req.Pagination.Offset, req.Pagination.Limit, 100)
		if start < 0 || end < 0 {
			contracts = []types.DynaContract{}
		} else {
			contracts = contracts[start:end]
		}
		pageRes, err = query.NewPaginationResponse(uint64(len(contracts)), req.Pagination.Offset, req.Pagination.Limit)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &types.QueryDynaContractsResponse{
		Contracts:  contracts,
		Pagination: pageRes,
	}, nil
}

// DynaContractExecutions returns executions for a specific dynamic contract
func (k Keeper) DynaContractExecutions(c context.Context, req *types.QueryDynaContractExecutionsRequest) (*types.QueryDynaContractExecutionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	executions := k.GetDynaContractExecutionsByContract(ctx, req.ContractId)
	var pageRes *query.PageResponse
	var err error

	// Apply pagination
	if req.Pagination != nil {
		start, end := query.Paginate(len(executions), req.Pagination.Offset, req.Pagination.Limit, 100)
		if start < 0 || end < 0 {
			executions = []types.DynaContractExecution{}
		} else {
			executions = executions[start:end]
		}
		pageRes, err = query.NewPaginationResponse(uint64(len(executions)), req.Pagination.Offset, req.Pagination.Limit)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &types.QueryDynaContractExecutionsResponse{
		Executions: executions,
		Pagination: pageRes,
	}, nil
}

// DynaContractExecution returns a specific dynamic contract execution
func (k Keeper) DynaContractExecution(c context.Context, req *types.QueryDynaContractExecutionRequest) (*types.QueryDynaContractExecutionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	execution, found := k.GetDynaContractExecution(ctx, req.Id)
	if !found {
		return nil, status.Error(codes.NotFound, "dynamic contract execution not found")
	}

	return &types.QueryDynaContractExecutionResponse{
		Execution: execution,
	}, nil
}

// DynaContractTemplates returns all dynamic contract templates
func (k Keeper) DynaContractTemplates(c context.Context, req *types.QueryDynaContractTemplatesRequest) (*types.QueryDynaContractTemplatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	templates := k.GetAllDynaContractTemplates(ctx)
	var pageRes *query.PageResponse
	var err error

	// Apply pagination
	if req.Pagination != nil {
		start, end := query.Paginate(len(templates), req.Pagination.Offset, req.Pagination.Limit, 100)
		if start < 0 || end < 0 {
			templates = []types.DynaContractTemplate{}
		} else {
			templates = templates[start:end]
		}
		pageRes, err = query.NewPaginationResponse(uint64(len(templates)), req.Pagination.Offset, req.Pagination.Limit)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &types.QueryDynaContractTemplatesResponse{
		Templates:  templates,
		Pagination: pageRes,
	}, nil
}

// DynaContractTemplate returns a specific dynamic contract template
func (k Keeper) DynaContractTemplate(c context.Context, req *types.QueryDynaContractTemplateRequest) (*types.QueryDynaContractTemplateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	template, found := k.GetDynaContractTemplate(ctx, req.Id)
	if !found {
		return nil, status.Error(codes.NotFound, "dynamic contract template not found")
	}

	return &types.QueryDynaContractTemplateResponse{
		Template: template,
	}, nil
}

// DynaContractLearningData returns learning data for a specific dynamic contract
func (k Keeper) DynaContractLearningData(c context.Context, req *types.QueryDynaContractLearningDataRequest) (*types.QueryDynaContractLearningDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	learningData := k.GetDynaContractLearningDataByContract(ctx, req.ContractId)
	var pageRes *query.PageResponse
	var err error

	// Apply pagination
	if req.Pagination != nil {
		start, end := query.Paginate(len(learningData), req.Pagination.Offset, req.Pagination.Limit, 100)
		if start < 0 || end < 0 {
			learningData = []types.DynaContractLearningData{}
		} else {
			learningData = learningData[start:end]
		}
		pageRes, err = query.NewPaginationResponse(uint64(len(learningData)), req.Pagination.Offset, req.Pagination.Limit)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &types.QueryDynaContractLearningDataResponse{
		LearningData: learningData,
		Pagination:   pageRes,
	}, nil
}

// DynaContractPermissions returns permissions for a specific dynamic contract
func (k Keeper) DynaContractPermissions(c context.Context, req *types.QueryDynaContractPermissionsRequest) (*types.QueryDynaContractPermissionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	permissions := k.GetDynaContractPermissionsByContract(ctx, req.ContractId)
	var pageRes *query.PageResponse
	var err error

	// Apply pagination
	if req.Pagination != nil {
		start, end := query.Paginate(len(permissions), req.Pagination.Offset, req.Pagination.Limit, 100)
		if start < 0 || end < 0 {
			permissions = []types.DynaContractPermission{}
		} else {
			permissions = permissions[start:end]
		}
		pageRes, err = query.NewPaginationResponse(uint64(len(permissions)), req.Pagination.Offset, req.Pagination.Limit)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &types.QueryDynaContractPermissionsResponse{
		Permissions: permissions,
		Pagination:  pageRes,
	}, nil
}

// DynaContractsByOwner returns dynamic contracts owned by a specific address
func (k Keeper) DynaContractsByOwner(c context.Context, req *types.QueryDynaContractsByOwnerRequest) (*types.QueryDynaContractsByOwnerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	owner, err := sdk.AccAddressFromBech32(req.Owner)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid owner address")
	}

	contracts := k.GetDynaContractsByOwner(ctx, owner)
	var pageRes *query.PageResponse

	// Apply pagination
	if req.Pagination != nil {
		start, end := query.Paginate(len(contracts), req.Pagination.Offset, req.Pagination.Limit, 100)
		if start < 0 || end < 0 {
			contracts = []types.DynaContract{}
		} else {
			contracts = contracts[start:end]
		}
		pageRes, err = query.NewPaginationResponse(uint64(len(contracts)), req.Pagination.Offset, req.Pagination.Limit)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &types.QueryDynaContractsByOwnerResponse{
		Contracts:  contracts,
		Pagination: pageRes,
	}, nil
}

// DynaContractsByAgent returns dynamic contracts associated with a specific AI agent
func (k Keeper) DynaContractsByAgent(c context.Context, req *types.QueryDynaContractsByAgentRequest) (*types.QueryDynaContractsByAgentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	contracts := k.GetDynaContractsByAgent(ctx, req.AgentId)
	var pageRes *query.PageResponse
	var err error

	// Apply pagination
	if req.Pagination != nil {
		start, end := query.Paginate(len(contracts), req.Pagination.Offset, req.Pagination.Limit, 100)
		if start < 0 || end < 0 {
			contracts = []types.DynaContract{}
		} else {
			contracts = contracts[start:end]
		}
		pageRes, err = query.NewPaginationResponse(uint64(len(contracts)), req.Pagination.Offset, req.Pagination.Limit)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &types.QueryDynaContractsByAgentResponse{
		Contracts:  contracts,
		Pagination: pageRes,
	}, nil
}

// DynaContractsByTags returns dynamic contracts with specific tags
func (k Keeper) DynaContractsByTags(c context.Context, req *types.QueryDynaContractsByTagsRequest) (*types.QueryDynaContractsByTagsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	contracts := k.GetDynaContractsByTags(ctx, req.Tags, req.MatchAll)
	var pageRes *query.PageResponse
	var err error

	// Apply pagination
	if req.Pagination != nil {
		start, end := query.Paginate(len(contracts), req.Pagination.Offset, req.Pagination.Limit, 100)
		if start < 0 || end < 0 {
			contracts = []types.DynaContract{}
		} else {
			contracts = contracts[start:end]
		}
		pageRes, err = query.NewPaginationResponse(uint64(len(contracts)), req.Pagination.Offset, req.Pagination.Limit)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &types.QueryDynaContractsByTagsResponse{
		Contracts:  contracts,
		Pagination: pageRes,
	}, nil
}