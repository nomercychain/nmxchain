package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/nomercychain/nmxchain/x/dynacontract/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// NewQuerier creates a new querier for dynacontract module
func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryGetDynaContract:
			return getDynaContract(ctx, path[1:], k, legacyQuerierCdc)
		case types.QueryListDynaContracts:
			return listDynaContracts(ctx, k, legacyQuerierCdc)
		case types.QueryGetDynaContractExecution:
			return getDynaContractExecution(ctx, path[1:], k, legacyQuerierCdc)
		case types.QueryListDynaContractExecutions:
			return listDynaContractExecutions(ctx, path[1:], k, legacyQuerierCdc)
		case types.QueryGetDynaContractTemplate:
			return getDynaContractTemplate(ctx, path[1:], k, legacyQuerierCdc)
		case types.QueryListDynaContractTemplates:
			return listDynaContractTemplates(ctx, k, legacyQuerierCdc)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown dynacontract query endpoint")
		}
	}
}

func getDynaContract(ctx sdk.Context, path []string, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	if len(path) != 1 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid path")
	}

	id := path[0]
	contract, found := k.GetDynaContract(ctx, id)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrContractNotFound, id)
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, contract)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func listDynaContracts(ctx sdk.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var contracts []types.DynaContract

	k.IterateContracts(ctx, func(contract types.DynaContract) (stop bool) {
		contracts = append(contracts, contract)
		return false
	})

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, contracts)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func getDynaContractExecution(ctx sdk.Context, path []string, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	if len(path) != 1 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid path")
	}

	id := path[0]
	execution, found := k.GetDynaContractExecution(ctx, id)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrExecutionNotFound, id)
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, execution)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func listDynaContractExecutions(ctx sdk.Context, path []string, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	if len(path) != 1 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid path")
	}

	contractID := path[0]
	var executions []types.DynaContractExecution

	k.IterateContractExecutions(ctx, contractID, func(execution types.DynaContractExecution) (stop bool) {
		executions = append(executions, execution)
		return false
	})

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, executions)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func getDynaContractTemplate(ctx sdk.Context, path []string, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	if len(path) != 1 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid path")
	}

	id := path[0]
	template, found := k.GetDynaContractTemplate(ctx, id)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrTemplateNotFound, id)
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, template)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func listDynaContractTemplates(ctx sdk.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var templates []types.DynaContractTemplate

	k.IterateTemplates(ctx, func(template types.DynaContractTemplate) (stop bool) {
		templates = append(templates, template)
		return false
	})

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, templates)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}