package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/nomercychain/nmxchain/x/neuropos/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// CreateValidator defines a method for creating a new validator
func (k msgServer) CreateValidator(goCtx context.Context, msg *types.MsgCreateValidator) (*types.MsgCreateValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	valAddr, err := sdk.ValAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		return nil, err
	}

	// Check if the validator already exists
	_, found := k.stakingKeeper.GetValidator(ctx, valAddr)
	if found {
		return nil, sdkerrors.Wrap(types.ErrValidatorExists, "validator already exists")
	}

	// Create validator
	// In a real implementation, this would interact with the staking module
	// For now, we'll create a simplified validator in our own store
	validator := types.Validator{
		OperatorAddress:          msg.ValidatorAddress,
		ConsensusPubkey:          msg.Pubkey,
		Jailed:                   false,
		Status:                   types.BondStatusBonded,
		Tokens:                   msg.SelfDelegation.Amount,
		DelegatorShares:          sdk.OneDec(),
		Description:              msg.Description,
		UnbondingHeight:          0,
		UnbondingTime:            ctx.BlockTime(),
		Commission:               types.Commission{CommissionRates: msg.CommissionRates},
		MinSelfDelegation:        msg.MinSelfDelegation,
		Reputation:               sdk.OneDec(),
		PerformanceScore:         sdk.OneDec(),
		NeuralNetworkContribution: sdk.ZeroDec(),
	}

	// Set the validator in the store
	k.SetValidator(ctx, validator)

	// Initialize validator reputation
	k.InitializeValidatorReputation(ctx, msg.ValidatorAddress)

	// Create self-delegation
	delegation := types.Delegation{
		DelegatorAddress: sdk.AccAddress(valAddr).String(),
		ValidatorAddress: msg.ValidatorAddress,
		Shares:           sdk.OneDec(),
	}

	// Set the delegation in the store
	k.SetDelegation(ctx, delegation)

	// Transfer tokens from the delegator's account to the module account
	delAddr := sdk.AccAddress(valAddr)
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, delAddr, types.ModuleName, sdk.NewCoins(msg.SelfDelegation)); err != nil {
		return nil, err
	}

	// Emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateValidator,
			sdk.NewAttribute(types.AttributeKeyValidator, msg.ValidatorAddress),
			sdk.NewAttribute(types.AttributeKeyAmount, msg.SelfDelegation.Amount.String()),
			sdk.NewAttribute(types.AttributeKeyDelegator, delAddr.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, delAddr.String()),
		),
	})

	return &types.MsgCreateValidatorResponse{}, nil
}

// EditValidator defines a method for editing an existing validator
func (k msgServer) EditValidator(goCtx context.Context, msg *types.MsgEditValidator) (*types.MsgEditValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	valAddr, err := sdk.ValAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		return nil, err
	}

	// Check if the validator exists
	validator, found := k.GetValidator(ctx, msg.ValidatorAddress)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrNoValidatorFound, "validator not found")
	}

	// Update description
	validator.Description = msg.Description

	// Update commission rate if provided
	if msg.CommissionRate != nil {
		// Check if commission rate is valid
		if msg.CommissionRate.GT(validator.Commission.CommissionRates.MaxRate) {
			return nil, sdkerrors.Wrap(types.ErrInvalidCommissionRate, "commission rate cannot be greater than max rate")
		}

		// Check if commission rate change is valid
		if validator.Commission.CommissionRates.Rate.Sub(*msg.CommissionRate).Abs().GT(validator.Commission.CommissionRates.MaxChangeRate) {
			return nil, sdkerrors.Wrap(types.ErrInvalidCommissionRate, "commission rate change is greater than max change rate")
		}

		// Update commission rate
		validator.Commission.CommissionRates.Rate = *msg.CommissionRate
	}

	// Update min self delegation if provided
	if msg.MinSelfDelegation != nil {
		// Check if min self delegation is valid
		if msg.MinSelfDelegation.LT(validator.MinSelfDelegation) {
			return nil, sdkerrors.Wrap(types.ErrInvalidInput, "min self delegation cannot be decreased")
		}

		// Update min self delegation
		validator.MinSelfDelegation = *msg.MinSelfDelegation
	}

	// Set the updated validator in the store
	k.SetValidator(ctx, validator)

	// Emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeEditValidator,
			sdk.NewAttribute(types.AttributeKeyValidator, msg.ValidatorAddress),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, sdk.AccAddress(valAddr).String()),
		),
	})

	return &types.MsgEditValidatorResponse{}, nil
}

// Delegate defines a method for performing a delegation of coins from a delegator to a validator
func (k msgServer) Delegate(goCtx context.Context, msg *types.MsgDelegate) (*types.MsgDelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	delAddr, err := sdk.AccAddressFromBech32(msg.DelegatorAddress)
	if err != nil {
		return nil, err
	}

	valAddr, err := sdk.ValAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		return nil, err
	}

	// Check if the validator exists
	validator, found := k.GetValidator(ctx, msg.ValidatorAddress)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrNoValidatorFound, "validator not found")
	}

	// Check if the validator is bonded
	if validator.Status != types.BondStatusBonded {
		return nil, sdkerrors.Wrap(types.ErrValidatorNotBonded, "validator is not bonded")
	}

	// Get or create delegation
	delegation, found := k.GetDelegation(ctx, msg.DelegatorAddress, msg.ValidatorAddress)
	if !found {
		// Create new delegation
		delegation = types.Delegation{
			DelegatorAddress: msg.DelegatorAddress,
			ValidatorAddress: msg.ValidatorAddress,
			Shares:           sdk.ZeroDec(),
		}
	}

	// Calculate new shares
	shares := sdk.NewDecFromInt(msg.Amount.Amount)
	delegation.Shares = delegation.Shares.Add(shares)

	// Update validator tokens
	validator.Tokens = validator.Tokens.Add(msg.Amount.Amount)

	// Set the updated delegation in the store
	k.SetDelegation(ctx, delegation)

	// Set the updated validator in the store
	k.SetValidator(ctx, validator)

	// Transfer tokens from the delegator's account to the module account
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, delAddr, types.ModuleName, sdk.NewCoins(msg.Amount)); err != nil {
		return nil, err
	}

	// Emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeDelegate,
			sdk.NewAttribute(types.AttributeKeyValidator, msg.ValidatorAddress),
			sdk.NewAttribute(types.AttributeKeyDelegator, msg.DelegatorAddress),
			sdk.NewAttribute(types.AttributeKeyAmount, msg.Amount.Amount.String()),
			sdk.NewAttribute(types.AttributeKeyShares, shares.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.DelegatorAddress),
		),
	})

	return &types.MsgDelegateResponse{}, nil
}

// Undelegate defines a method for performing an undelegation from a delegate and a validator
func (k msgServer) Undelegate(goCtx context.Context, msg *types.MsgUndelegate) (*types.MsgUndelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	delAddr, err := sdk.AccAddressFromBech32(msg.DelegatorAddress)
	if err != nil {
		return nil, err
	}

	valAddr, err := sdk.ValAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		return nil, err
	}

	// Check if the delegation exists
	delegation, found := k.GetDelegation(ctx, msg.DelegatorAddress, msg.ValidatorAddress)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrNoDelegation, "delegation not found")
	}

	// Check if the validator exists
	validator, found := k.GetValidator(ctx, msg.ValidatorAddress)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrNoValidatorFound, "validator not found")
	}

	// Calculate shares to remove
	shares := sdk.NewDecFromInt(msg.Amount.Amount)
	if shares.GT(delegation.Shares) {
		return nil, sdkerrors.Wrap(types.ErrInvalidDelegation, "unbonding amount exceeds delegation")
	}

	// Update delegation shares
	delegation.Shares = delegation.Shares.Sub(shares)

	// Update validator tokens
	validator.Tokens = validator.Tokens.Sub(msg.Amount.Amount)

	// Create unbonding delegation
	unbondingDelegation, found := k.GetUnbondingDelegation(ctx, msg.DelegatorAddress, msg.ValidatorAddress)
	if !found {
		unbondingDelegation = types.UnbondingDelegation{
			DelegatorAddress: msg.DelegatorAddress,
			ValidatorAddress: msg.ValidatorAddress,
			Entries:          []types.UnbondingDelegationEntry{},
		}
	}

	// Create new entry
	completionTime := ctx.BlockTime().Add(k.UnbondingTime(ctx))
	entry := types.UnbondingDelegationEntry{
		CreationHeight:  ctx.BlockHeight(),
		CompletionTime:  completionTime,
		InitialBalance:  msg.Amount.Amount,
		Balance:         msg.Amount.Amount,
	}

	// Add entry to unbonding delegation
	unbondingDelegation.Entries = append(unbondingDelegation.Entries, entry)

	// Set the updated delegation in the store
	if delegation.Shares.IsZero() {
		// Remove delegation if shares are zero
		store := ctx.KVStore(k.storeKey)
		key := types.DelegationKey(msg.DelegatorAddress, msg.ValidatorAddress)
		store.Delete(key)
	} else {
		k.SetDelegation(ctx, delegation)
	}

	// Set the updated validator in the store
	k.SetValidator(ctx, validator)

	// Set the unbonding delegation in the store
	k.SetUnbondingDelegation(ctx, unbondingDelegation)

	// Emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUnbond,
			sdk.NewAttribute(types.AttributeKeyValidator, msg.ValidatorAddress),
			sdk.NewAttribute(types.AttributeKeyDelegator, msg.DelegatorAddress),
			sdk.NewAttribute(types.AttributeKeyAmount, msg.Amount.Amount.String()),
			sdk.NewAttribute(types.AttributeKeyCompletionTime, completionTime.Format(time.RFC3339)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.DelegatorAddress),
		),
	})

	return &types.MsgUndelegateResponse{
		CompletionTime: completionTime,
	}, nil
}

// BeginRedelegate defines a method for performing a redelegation of coins from a delegator and source validator to a destination validator
func (k msgServer) BeginRedelegate(goCtx context.Context, msg *types.MsgBeginRedelegate) (*types.MsgBeginRedelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	delAddr, err := sdk.AccAddressFromBech32(msg.DelegatorAddress)
	if err != nil {
		return nil, err
	}

	valSrcAddr, err := sdk.ValAddressFromBech32(msg.ValidatorSrcAddress)
	if err != nil {
		return nil, err
	}

	valDstAddr, err := sdk.ValAddressFromBech32(msg.ValidatorDstAddress)
	if err != nil {
		return nil, err
	}

	// Check if the source delegation exists
	srcDelegation, found := k.GetDelegation(ctx, msg.DelegatorAddress, msg.ValidatorSrcAddress)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrNoDelegation, "source delegation not found")
	}

	// Check if the source validator exists
	srcValidator, found := k.GetValidator(ctx, msg.ValidatorSrcAddress)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrNoValidatorFound, "source validator not found")
	}

	// Check if the destination validator exists
	dstValidator, found := k.GetValidator(ctx, msg.ValidatorDstAddress)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrNoValidatorFound, "destination validator not found")
	}

	// Check if the destination validator is bonded
	if dstValidator.Status != types.BondStatusBonded {
		return nil, sdkerrors.Wrap(types.ErrValidatorNotBonded, "destination validator is not bonded")
	}

	// Calculate shares to remove from source
	shares := sdk.NewDecFromInt(msg.Amount.Amount)
	if shares.GT(srcDelegation.Shares) {
		return nil, sdkerrors.Wrap(types.ErrInvalidDelegation, "redelegation amount exceeds delegation")
	}

	// Update source delegation shares
	srcDelegation.Shares = srcDelegation.Shares.Sub(shares)

	// Update source validator tokens
	srcValidator.Tokens = srcValidator.Tokens.Sub(msg.Amount.Amount)

	// Get or create destination delegation
	dstDelegation, found := k.GetDelegation(ctx, msg.DelegatorAddress, msg.ValidatorDstAddress)
	if !found {
		// Create new delegation
		dstDelegation = types.Delegation{
			DelegatorAddress: msg.DelegatorAddress,
			ValidatorAddress: msg.ValidatorDstAddress,
			Shares:           sdk.ZeroDec(),
		}
	}

	// Update destination delegation shares
	dstDelegation.Shares = dstDelegation.Shares.Add(shares)

	// Update destination validator tokens
	dstValidator.Tokens = dstValidator.Tokens.Add(msg.Amount.Amount)

	// Create redelegation
	redelegation, found := k.GetRedelegation(ctx, msg.DelegatorAddress, msg.ValidatorSrcAddress, msg.ValidatorDstAddress)
	if !found {
		redelegation = types.Redelegation{
			DelegatorAddress:    msg.DelegatorAddress,
			ValidatorSrcAddress: msg.ValidatorSrcAddress,
			ValidatorDstAddress: msg.ValidatorDstAddress,
			Entries:             []types.RedelegationEntry{},
		}
	}

	// Create new entry
	completionTime := ctx.BlockTime().Add(k.UnbondingTime(ctx))
	entry := types.RedelegationEntry{
		CreationHeight: ctx.BlockHeight(),
		CompletionTime: completionTime,
		InitialBalance: msg.Amount.Amount,
		SharesDst:      shares,
	}

	// Add entry to redelegation
	redelegation.Entries = append(redelegation.Entries, entry)

	// Set the updated source delegation in the store
	if srcDelegation.Shares.IsZero() {
		// Remove delegation if shares are zero
		store := ctx.KVStore(k.storeKey)
		key := types.DelegationKey(msg.DelegatorAddress, msg.ValidatorSrcAddress)
		store.Delete(key)
	} else {
		k.SetDelegation(ctx, srcDelegation)
	}

	// Set the updated destination delegation in the store
	k.SetDelegation(ctx, dstDelegation)

	// Set the updated validators in the store
	k.SetValidator(ctx, srcValidator)
	k.SetValidator(ctx, dstValidator)

	// Set the redelegation in the store
	k.SetRedelegation(ctx, redelegation)

	// Emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeRedelegate,
			sdk.NewAttribute(types.AttributeKeySrcValidator, msg.ValidatorSrcAddress),
			sdk.NewAttribute(types.AttributeKeyDstValidator, msg.ValidatorDstAddress),
			sdk.NewAttribute(types.AttributeKeyDelegator, msg.DelegatorAddress),
			sdk.NewAttribute(types.AttributeKeyAmount, msg.Amount.Amount.String()),
			sdk.NewAttribute(types.AttributeKeyCompletionTime, completionTime.Format(time.RFC3339)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.DelegatorAddress),
		),
	})

	return &types.MsgBeginRedelegateResponse{
		CompletionTime: completionTime,
	}, nil
}

// Unjail defines a method for unjailing a jailed validator
func (k msgServer) Unjail(goCtx context.Context, msg *types.MsgUnjail) (*types.MsgUnjailResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	valAddr, err := sdk.ValAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		return nil, err
	}

	// Check if the validator exists
	validator, found := k.GetValidator(ctx, msg.ValidatorAddress)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrNoValidatorFound, "validator not found")
	}

	// Check if the validator is jailed
	if !validator.Jailed {
		return nil, sdkerrors.Wrap(types.ErrValidatorNotJailed, "validator is not jailed")
	}

	// Check if the validator can be unjailed
	signingInfo, found := k.GetValidatorSigningInfo(ctx, msg.ValidatorAddress)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrNoValidatorSigningInfo, "no validator signing info found")
	}

	if signingInfo.JailedUntil.After(ctx.BlockTime()) {
		return nil, sdkerrors.Wrap(types.ErrValidatorJailed, "validator still jailed, cannot be unjailed until ")
	}

	// Unjail the validator
	validator.Jailed = false
	validator.Status = types.BondStatusBonded

	// Set the updated validator in the store
	k.SetValidator(ctx, validator)

	// Emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUnjail,
			sdk.NewAttribute(types.AttributeKeyValidator, msg.ValidatorAddress),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, sdk.AccAddress(valAddr).String()),
		),
	})

	return &types.MsgUnjailResponse{}, nil
}

// UpdateNeuralNetwork defines a method for updating a neural network
func (k msgServer) UpdateNeuralNetwork(goCtx context.Context, msg *types.MsgUpdateNeuralNetwork) (*types.MsgUpdateNeuralNetworkResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	valAddr, err := sdk.ValAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		return nil, err
	}

	// Check if the validator exists
	_, found := k.GetValidator(ctx, msg.ValidatorAddress)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrNoValidatorFound, "validator not found")
	}

	// Convert layers from proto to domain type
	var layers []types.Layer
	for _, layer := range msg.Layers {
		layers = append(layers, types.Layer{
			Type:       layer.Type,
			InputSize:  layer.InputSize,
			OutputSize: layer.OutputSize,
			Activation: layer.Activation,
		})
	}

	// Update or create neural network
	if msg.NetworkId == "" {
		// Create new neural network
		network, err := k.CreateNeuralNetwork(ctx, msg.Architecture, layers, msg.Metadata)
		if err != nil {
			return nil, err
		}

		// Emit events
		ctx.EventManager().EmitEvents(sdk.Events{
			sdk.NewEvent(
				types.EventTypeCreateNeuralNetwork,
				sdk.NewAttribute(types.AttributeKeyValidator, msg.ValidatorAddress),
				sdk.NewAttribute(types.AttributeKeyNeuralNetworkID, network.ID),
				sdk.NewAttribute(types.AttributeKeyNeuralNetworkArchitecture, network.Architecture),
			),
			sdk.NewEvent(
				sdk.EventTypeMessage,
				sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
				sdk.NewAttribute(sdk.AttributeKeySender, sdk.AccAddress(valAddr).String()),
			),
		})

		return &types.MsgUpdateNeuralNetworkResponse{
			NetworkId: network.ID,
		}, nil
	} else {
		// Update existing neural network
		err := k.UpdateNeuralNetwork(ctx, msg.NetworkId, msg.Architecture, layers, msg.Weights, msg.Metadata)
		if err != nil {
			return nil, err
		}

		// Emit events
		ctx.EventManager().EmitEvents(sdk.Events{
			sdk.NewEvent(
				types.EventTypeUpdateNeuralNetwork,
				sdk.NewAttribute(types.AttributeKeyValidator, msg.ValidatorAddress),
				sdk.NewAttribute(types.AttributeKeyNeuralNetworkID, msg.NetworkId),
			),
			sdk.NewEvent(
				sdk.EventTypeMessage,
				sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
				sdk.NewAttribute(sdk.AttributeKeySender, sdk.AccAddress(valAddr).String()),
			),
		})

		return &types.MsgUpdateNeuralNetworkResponse{
			NetworkId: msg.NetworkId,
		}, nil
	}
}

// TrainNeuralNetwork defines a method for training a neural network
func (k msgServer) TrainNeuralNetwork(goCtx context.Context, msg *types.MsgTrainNeuralNetwork) (*types.MsgTrainNeuralNetworkResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	valAddr, err := sdk.ValAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		return nil, err
	}

	// Check if the validator exists
	_, found := k.GetValidator(ctx, msg.ValidatorAddress)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrNoValidatorFound, "validator not found")
	}

	// Train neural network
	err := k.TrainNeuralNetwork(ctx, msg.NetworkId, msg.Features, msg.Labels, msg.Epochs, msg.LearningRate, msg.Metadata)
	if err != nil {
		return nil, err
	}

	// Get the updated network
	network, found := k.GetNeuralNetwork(ctx, msg.NetworkId)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrNoNeuralNetworkFound, "neural network not found after training")
	}

	// Emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeTrainNeuralNetwork,
			sdk.NewAttribute(types.AttributeKeyValidator, msg.ValidatorAddress),
			sdk.NewAttribute(types.AttributeKeyNeuralNetworkID, msg.NetworkId),
			sdk.NewAttribute(types.AttributeKeyNeuralNetworkAccuracy, network.Accuracy.String()),
			sdk.NewAttribute(types.AttributeKeyNeuralNetworkLoss, network.Loss.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, sdk.AccAddress(valAddr).String()),
		),
	})

	return &types.MsgTrainNeuralNetworkResponse{
		Accuracy: network.Accuracy.String(),
		Loss:     network.Loss.String(),
	}, nil
}

// SubmitNeuralPrediction defines a method for submitting a neural prediction
func (k msgServer) SubmitNeuralPrediction(goCtx context.Context, msg *types.MsgSubmitNeuralPrediction) (*types.MsgSubmitNeuralPredictionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	valAddr, err := sdk.ValAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		return nil, err
	}

	// Check if the validator exists
	_, found := k.GetValidator(ctx, msg.ValidatorAddress)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrNoValidatorFound, "validator not found")
	}

	// Submit neural prediction
	prediction, err := k.SubmitNeuralPrediction(ctx, msg.NetworkId, msg.Input, msg.Output, msg.Confidence, valAddr, msg.Metadata)
	if err != nil {
		return nil, err
	}

	// Emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSubmitNeuralPrediction,
			sdk.NewAttribute(types.AttributeKeyValidator, msg.ValidatorAddress),
			sdk.NewAttribute(types.AttributeKeyNeuralNetworkID, msg.NetworkId),
			sdk.NewAttribute(types.AttributeKeyNeuralPredictionID, prediction.ID),
			sdk.NewAttribute(types.AttributeKeyNeuralPredictionConfidence, msg.Confidence.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, sdk.AccAddress(valAddr).String()),
		),
	})

	return &types.MsgSubmitNeuralPredictionResponse{
		PredictionId: prediction.ID,
	}, nil
}

// UpdateValidatorReputation defines a method for updating a validator's reputation
func (k msgServer) UpdateValidatorReputation(goCtx context.Context, msg *types.MsgUpdateValidatorReputation) (*types.MsgUpdateValidatorReputationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	adminAddr, err := sdk.AccAddressFromBech32(msg.AdminAddress)
	if err != nil {
		return nil, err
	}

	valAddr, err := sdk.ValAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		return nil, err
	}

	// Check if the sender is authorized to update reputations
	// In a real implementation, this would check against a list of authorized admins
	// For now, we'll use a simple check against the module account
	moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)
	if !adminAddr.Equals(moduleAddr) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "sender is not authorized to update validator reputations")
	}

	// Check if the validator exists
	_, found := k.GetValidator(ctx, msg.ValidatorAddress)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrNoValidatorFound, "validator not found")
	}

	// Update validator reputation
	err := k.UpdateValidatorReputation(ctx, msg.ValidatorAddress, msg.ReputationChange, msg.Reason)
	if err != nil {
		return nil, err
	}

	// Get the updated reputation
	reputation, found := k.GetValidatorReputation(ctx, msg.ValidatorAddress)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrNoValidatorFound, "validator reputation not found after update")
	}

	// Emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUpdateValidatorReputation,
			sdk.NewAttribute(types.AttributeKeyValidator, msg.ValidatorAddress),
			sdk.NewAttribute(types.AttributeKeyValidatorReputation, reputation.Reputation.String()),
			sdk.NewAttribute(types.AttributeKeyReputationChange, msg.ReputationChange.String()),
			sdk.NewAttribute(types.AttributeKeyReputationReason, msg.Reason),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.AdminAddress),
		),
	})

	return &types.MsgUpdateValidatorReputationResponse{
		NewReputation: reputation.Reputation.String(),
	}, nil
}