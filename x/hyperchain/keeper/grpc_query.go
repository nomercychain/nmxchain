package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/nomercychain/nmxchain/x/hyperchain/types"
)

var _ types.QueryServer = Keeper{}

// Hyperchain returns information about a specific hyperchain
func (k Keeper) Hyperchain(c context.Context, req *types.QueryHyperchainRequest) (*types.QueryHyperchainResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	hyperchain, found := k.GetHyperchain(ctx, req.Id)
	if !found {
		return nil, status.Error(codes.NotFound, "hyperchain not found")
	}

	return &types.QueryHyperchainResponse{
		Hyperchain: hyperchain,
	}, nil
}

// Hyperchains returns all hyperchains
func (k Keeper) Hyperchains(c context.Context, req *types.QueryHyperchainsRequest) (*types.QueryHyperchainsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	hyperchains := k.GetAllHyperchains(ctx)
	var pageRes *query.PageResponse
	var err error

	// Apply pagination
	if req.Pagination != nil {
		start, end := query.Paginate(len(hyperchains), req.Pagination.Offset, req.Pagination.Limit, 100)
		if start < 0 || end < 0 {
			hyperchains = []types.Hyperchain{}
		} else {
			hyperchains = hyperchains[start:end]
		}
		pageRes, err = query.NewPaginationResponse(uint64(len(hyperchains)), req.Pagination.Offset, req.Pagination.Limit)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &types.QueryHyperchainsResponse{
		Hyperchains: hyperchains,
		Pagination:  pageRes,
	}, nil
}

// HyperchainsByCreator returns hyperchains created by a specific address
func (k Keeper) HyperchainsByCreator(c context.Context, req *types.QueryHyperchainsByCreatorRequest) (*types.QueryHyperchainsByCreatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	hyperchains := k.GetHyperchainsByCreator(ctx, req.Creator)
	var pageRes *query.PageResponse
	var err error

	// Apply pagination
	if req.Pagination != nil {
		start, end := query.Paginate(len(hyperchains), req.Pagination.Offset, req.Pagination.Limit, 100)
		if start < 0 || end < 0 {
			hyperchains = []types.Hyperchain{}
		} else {
			hyperchains = hyperchains[start:end]
		}
		pageRes, err = query.NewPaginationResponse(uint64(len(hyperchains)), req.Pagination.Offset, req.Pagination.Limit)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &types.QueryHyperchainsByCreatorResponse{
		Hyperchains: hyperchains,
		Pagination:  pageRes,
	}, nil
}

// HyperchainsByParent returns hyperchains with a specific parent chain
func (k Keeper) HyperchainsByParent(c context.Context, req *types.QueryHyperchainsByParentRequest) (*types.QueryHyperchainsByParentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	hyperchains := k.GetHyperchainsByParent(ctx, req.ParentChainId)
	var pageRes *query.PageResponse
	var err error

	// Apply pagination
	if req.Pagination != nil {
		start, end := query.Paginate(len(hyperchains), req.Pagination.Offset, req.Pagination.Limit, 100)
		if start < 0 || end < 0 {
			hyperchains = []types.Hyperchain{}
		} else {
			hyperchains = hyperchains[start:end]
		}
		pageRes, err = query.NewPaginationResponse(uint64(len(hyperchains)), req.Pagination.Offset, req.Pagination.Limit)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &types.QueryHyperchainsByParentResponse{
		Hyperchains: hyperchains,
		Pagination:  pageRes,
	}, nil
}

// HyperchainValidator returns information about a specific validator in a hyperchain
func (k Keeper) HyperchainValidator(c context.Context, req *types.QueryHyperchainValidatorRequest) (*types.QueryHyperchainValidatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	validator, found := k.GetHyperchainValidator(ctx, req.ChainId, req.Address)
	if !found {
		return nil, status.Error(codes.NotFound, "validator not found")
	}

	return &types.QueryHyperchainValidatorResponse{
		Validator: validator,
	}, nil
}

// HyperchainValidators returns all validators in a hyperchain
func (k Keeper) HyperchainValidators(c context.Context, req *types.QueryHyperchainValidatorsRequest) (*types.QueryHyperchainValidatorsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	validators := k.GetHyperchainValidatorsByChain(ctx, req.ChainId)
	var pageRes *query.PageResponse
	var err error

	// Apply pagination
	if req.Pagination != nil {
		start, end := query.Paginate(len(validators), req.Pagination.Offset, req.Pagination.Limit, 100)
		if start < 0 || end < 0 {
			validators = []types.HyperchainValidator{}
		} else {
			validators = validators[start:end]
		}
		pageRes, err = query.NewPaginationResponse(uint64(len(validators)), req.Pagination.Offset, req.Pagination.Limit)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &types.QueryHyperchainValidatorsResponse{
		Validators: validators,
		Pagination: pageRes,
	}, nil
}

// HyperchainBlock returns a specific block in a hyperchain
func (k Keeper) HyperchainBlock(c context.Context, req *types.QueryHyperchainBlockRequest) (*types.QueryHyperchainBlockResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	block, found := k.GetHyperchainBlock(ctx, req.ChainId, req.Height)
	if !found {
		return nil, status.Error(codes.NotFound, "block not found")
	}

	return &types.QueryHyperchainBlockResponse{
		Block: block,
	}, nil
}

// HyperchainBlocks returns blocks in a hyperchain
func (k Keeper) HyperchainBlocks(c context.Context, req *types.QueryHyperchainBlocksRequest) (*types.QueryHyperchainBlocksResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	blocks := k.GetHyperchainBlocksByChain(ctx, req.ChainId)
	var pageRes *query.PageResponse
	var err error

	// Apply pagination
	if req.Pagination != nil {
		start, end := query.Paginate(len(blocks), req.Pagination.Offset, req.Pagination.Limit, 100)
		if start < 0 || end < 0 {
			blocks = []types.HyperchainBlock{}
		} else {
			blocks = blocks[start:end]
		}
		pageRes, err = query.NewPaginationResponse(uint64(len(blocks)), req.Pagination.Offset, req.Pagination.Limit)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &types.QueryHyperchainBlocksResponse{
		Blocks:     blocks,
		Pagination: pageRes,
	}, nil
}

// HyperchainTransaction returns a specific transaction in a hyperchain
func (k Keeper) HyperchainTransaction(c context.Context, req *types.QueryHyperchainTransactionRequest) (*types.QueryHyperchainTransactionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	transaction, found := k.GetHyperchainTransaction(ctx, req.Id)
	if !found {
		return nil, status.Error(codes.NotFound, "transaction not found")
	}

	return &types.QueryHyperchainTransactionResponse{
		Transaction: transaction,
	}, nil
}

// HyperchainTransactions returns transactions in a hyperchain
func (k Keeper) HyperchainTransactions(c context.Context, req *types.QueryHyperchainTransactionsRequest) (*types.QueryHyperchainTransactionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	transactions := k.GetHyperchainTransactionsByChain(ctx, req.ChainId)
	var pageRes *query.PageResponse
	var err error

	// Apply pagination
	if req.Pagination != nil {
		start, end := query.Paginate(len(transactions), req.Pagination.Offset, req.Pagination.Limit, 100)
		if start < 0 || end < 0 {
			transactions = []types.HyperchainTransaction{}
		} else {
			transactions = transactions[start:end]
		}
		pageRes, err = query.NewPaginationResponse(uint64(len(transactions)), req.Pagination.Offset, req.Pagination.Limit)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &types.QueryHyperchainTransactionsResponse{
		Transactions: transactions,
		Pagination:   pageRes,
	}, nil
}

// HyperchainBridge returns information about a specific hyperchain bridge
func (k Keeper) HyperchainBridge(c context.Context, req *types.QueryHyperchainBridgeRequest) (*types.QueryHyperchainBridgeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	bridge, found := k.GetHyperchainBridge(ctx, req.Id)
	if !found {
		return nil, status.Error(codes.NotFound, "bridge not found")
	}

	return &types.QueryHyperchainBridgeResponse{
		Bridge: bridge,
	}, nil
}

// HyperchainBridges returns all hyperchain bridges
func (k Keeper) HyperchainBridges(c context.Context, req *types.QueryHyperchainBridgesRequest) (*types.QueryHyperchainBridgesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	bridges := k.GetAllHyperchainBridges(ctx)
	var pageRes *query.PageResponse
	var err error

	// Apply pagination
	if req.Pagination != nil {
		start, end := query.Paginate(len(bridges), req.Pagination.Offset, req.Pagination.Limit, 100)
		if start < 0 || end < 0 {
			bridges = []types.HyperchainBridge{}
		} else {
			bridges = bridges[start:end]
		}
		pageRes, err = query.NewPaginationResponse(uint64(len(bridges)), req.Pagination.Offset, req.Pagination.Limit)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &types.QueryHyperchainBridgesResponse{
		Bridges:    bridges,
		Pagination: pageRes,
	}, nil
}

// HyperchainBridgesByChain returns bridges for a specific hyperchain
func (k Keeper) HyperchainBridgesByChain(c context.Context, req *types.QueryHyperchainBridgesByChainRequest) (*types.QueryHyperchainBridgesByChainResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	bridges := k.GetHyperchainBridgesByChain(ctx, req.ChainId)
	var pageRes *query.PageResponse
	var err error

	// Apply pagination
	if req.Pagination != nil {
		start, end := query.Paginate(len(bridges), req.Pagination.Offset, req.Pagination.Limit, 100)
		if start < 0 || end < 0 {
			bridges = []types.HyperchainBridge{}
		} else {
			bridges = bridges[start:end]
		}
		pageRes, err = query.NewPaginationResponse(uint64(len(bridges)), req.Pagination.Offset, req.Pagination.Limit)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &types.QueryHyperchainBridgesByChainResponse{
		Bridges:    bridges,
		Pagination: pageRes,
	}, nil
}

// HyperchainBridgeTransaction returns a specific hyperchain bridge transaction
func (k Keeper) HyperchainBridgeTransaction(c context.Context, req *types.QueryHyperchainBridgeTransactionRequest) (*types.QueryHyperchainBridgeTransactionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	transaction, found := k.GetHyperchainBridgeTransaction(ctx, req.Id)
	if !found {
		return nil, status.Error(codes.NotFound, "bridge transaction not found")
	}

	return &types.QueryHyperchainBridgeTransactionResponse{
		Transaction: transaction,
	}, nil
}

// HyperchainBridgeTransactions returns hyperchain bridge transactions
func (k Keeper) HyperchainBridgeTransactions(c context.Context, req *types.QueryHyperchainBridgeTransactionsRequest) (*types.QueryHyperchainBridgeTransactionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	transactions := k.GetHyperchainBridgeTransactionsByBridge(ctx, req.BridgeId)
	var pageRes *query.PageResponse
	var err error

	// Apply pagination
	if req.Pagination != nil {
		start, end := query.Paginate(len(transactions), req.Pagination.Offset, req.Pagination.Limit, 100)
		if start < 0 || end < 0 {
			transactions = []types.HyperchainBridgeTransaction{}
		} else {
			transactions = transactions[start:end]
		}
		pageRes, err = query.NewPaginationResponse(uint64(len(transactions)), req.Pagination.Offset, req.Pagination.Limit)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &types.QueryHyperchainBridgeTransactionsResponse{
		Transactions: transactions,
		Pagination:   pageRes,
	}, nil
}

// HyperchainPermissions returns permissions for a specific hyperchain
func (k Keeper) HyperchainPermissions(c context.Context, req *types.QueryHyperchainPermissionsRequest) (*types.QueryHyperchainPermissionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	permissions := k.GetHyperchainPermissionsByChain(ctx, req.ChainId)
	var pageRes *query.PageResponse
	var err error

	// Apply pagination
	if req.Pagination != nil {
		start, end := query.Paginate(len(permissions), req.Pagination.Offset, req.Pagination.Limit, 100)
		if start < 0 || end < 0 {
			permissions = []types.HyperchainPermission{}
		} else {
			permissions = permissions[start:end]
		}
		pageRes, err = query.NewPaginationResponse(uint64(len(permissions)), req.Pagination.Offset, req.Pagination.Limit)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &types.QueryHyperchainPermissionsResponse{
		Permissions: permissions,
		Pagination:  pageRes,
	}, nil
}

// HyperchainPermission returns a specific permission for a hyperchain
func (k Keeper) HyperchainPermission(c context.Context, req *types.QueryHyperchainPermissionRequest) (*types.QueryHyperchainPermissionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	permission, found := k.GetHyperchainPermission(ctx, req.ChainId, req.Address, req.PermissionType)
	if !found {
		return nil, status.Error(codes.NotFound, "permission not found")
	}

	return &types.QueryHyperchainPermissionResponse{
		Permission: permission,
	}, nil
}