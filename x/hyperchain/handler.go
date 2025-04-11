package hyperchain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/nomercychain/nmxchain/x/hyperchain/keeper"
	"github.com/nomercychain/nmxchain/x/hyperchain/types"
)

// NewHandler creates a new handler for the hyperchain module
func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgCreateHyperchain:
			res, err := msgServer.CreateHyperchain(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgUpdateHyperchain:
			res, err := msgServer.UpdateHyperchain(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgJoinHyperchainAsValidator:
			res, err := msgServer.JoinHyperchainAsValidator(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgLeaveHyperchain:
			res, err := msgServer.LeaveHyperchain(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgCreateHyperchainBridge:
			res, err := msgServer.CreateHyperchainBridge(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgUpdateHyperchainBridge:
			res, err := msgServer.UpdateHyperchainBridge(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgRegisterHyperchainBridgeRelayer:
			res, err := msgServer.RegisterHyperchainBridgeRelayer(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgRemoveHyperchainBridgeRelayer:
			res, err := msgServer.RemoveHyperchainBridgeRelayer(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgInitiateHyperchainBridgeTransaction:
			res, err := msgServer.InitiateHyperchainBridgeTransaction(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgApproveHyperchainBridgeTransaction:
			res, err := msgServer.ApproveHyperchainBridgeTransaction(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgSubmitHyperchainBlock:
			res, err := msgServer.SubmitHyperchainBlock(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgSubmitHyperchainTransaction:
			res, err := msgServer.SubmitHyperchainTransaction(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgGrantHyperchainPermission:
			res, err := msgServer.GrantHyperchainPermission(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgRevokeHyperchainPermission:
			res, err := msgServer.RevokeHyperchainPermission(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s message type: %T", types.ModuleName, msg)
		}
	}
}