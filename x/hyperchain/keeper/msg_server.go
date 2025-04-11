package keeper

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/nomercychain/nmxchain/x/hyperchain/types"
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

// CreateHyperchain creates a new hyperchain
func (k msgServer) CreateHyperchain(goCtx context.Context, msg *types.MsgCreateHyperchain) (*types.MsgCreateHyperchainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// Parse the creator address
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	// Check if the creator has reached the maximum number of hyperchains
	creatorHyperchains := k.GetHyperchainsByCreator(ctx, msg.Creator)
	maxHyperchainsPerAccount := k.MaxHyperchainsPerAccount(ctx)
	if uint64(len(creatorHyperchains)) >= maxHyperchainsPerAccount {
		return nil, sdkerrors.Wrapf(types.ErrMaxHyperchainsReached, "creator has reached the maximum number of hyperchains (%d)", maxHyperchainsPerAccount)
	}

	// Check deposit
	minDeposit := k.MinHyperchainCreationDeposit(ctx)
	if !msg.Deposit.IsGTE(minDeposit) {
		return nil, sdkerrors.Wrapf(types.ErrInsufficientDeposit, "deposit %s is less than minimum %s", msg.Deposit, minDeposit)
	}

	// If an agent ID is provided, check if it exists
	if msg.AgentId != "" {
		_, found := k.deaiKeeper.GetAIAgent(ctx, msg.AgentId)
		if !found {
			return nil, sdkerrors.Wrapf(types.ErrAgentNotFound, "agent ID %s not found", msg.AgentId)
		}
	}

	// Collect the deposit
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, sdk.NewCoins(msg.Deposit))
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to collect deposit")
	}

	// Generate a unique ID for the hyperchain
	id := uuid.New().String()

	// Create the hyperchain
	hyperchain := types.Hyperchain{
		Id:               id,
		Name:             msg.Name,
		Description:      msg.Description,
		Creator:          msg.Creator,
		Admin:            msg.Creator,
		ChainType:        msg.ChainType,
		ConsensusType:    msg.ConsensusType,
		Status:           types.HyperchainStatusInitializing,
		BlockHeight:      0,
		GenesisHash:      "",
		GenesisConfig:    msg.GenesisConfig,
		ChainConfig:      msg.ChainConfig,
		Metadata:         msg.Metadata,
		CreatedAt:        ctx.BlockTime(),
		UpdatedAt:        ctx.BlockTime(),
		Validators:       []string{},
		MaxValidators:    msg.MaxValidators,
		MinStake:         msg.MinStake,
		ParentChainId:    msg.ParentChainId,
		ChildChainIds:    []string{},
		SupportedTokens:  msg.SupportedTokens,
		SupportedModules: msg.SupportedModules,
		AgentId:          msg.AgentId,
	}

	// Store the hyperchain
	k.SetHyperchain(ctx, hyperchain)

	// Grant admin permission to the creator
	permission := types.HyperchainPermission{
		ChainId:        id,
		Address:        msg.Creator,
		PermissionType: types.PermissionTypeAdmin,
		GrantedAt:      ctx.BlockTime(),
		ExpiresAt:      time.Time{}, // Never expires
		GrantedBy:      msg.Creator,
		Metadata:       []byte("{}"),
	}
	k.SetHyperchainPermission(ctx, permission)

	// If this is a child chain, update the parent chain
	if msg.ParentChainId != "" {
		parentChain, found := k.GetHyperchain(ctx, msg.ParentChainId)
		if found {
			parentChain.ChildChainIds = append(parentChain.ChildChainIds, id)
			k.SetHyperchain(ctx, parentChain)
		}
	}

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCreateHyperchain,
			sdk.NewAttribute(types.AttributeKeyHyperchainID, id),
			sdk.NewAttribute(types.AttributeKeyCreator, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyName, msg.Name),
			sdk.NewAttribute(types.AttributeKeyChainType, msg.ChainType.String()),
			sdk.NewAttribute(types.AttributeKeyConsensusType, msg.ConsensusType.String()),
			sdk.NewAttribute(types.AttributeKeyDeposit, msg.Deposit.String()),
		),
	)

	return &types.MsgCreateHyperchainResponse{
		Id: id,
	}, nil
}

// UpdateHyperchain updates an existing hyperchain
func (k msgServer) UpdateHyperchain(goCtx context.Context, msg *types.MsgUpdateHyperchain) (*types.MsgUpdateHyperchainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// Get the hyperchain
	hyperchain, found := k.GetHyperchain(ctx, msg.ChainId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrHyperchainNotFound, "hyperchain ID %s not found", msg.ChainId)
	}

	// Parse the admin address
	admin, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}

	// Check if the caller is the admin
	if hyperchain.Admin != admin.String() && !k.HasPermission(ctx, msg.ChainId, admin.String(), types.PermissionTypeAdmin) {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, "only the admin can update the hyperchain")
	}

	// If an agent ID is provided, check if it exists
	if msg.AgentId != "" && msg.AgentId != hyperchain.AgentId {
		_, found := k.deaiKeeper.GetAIAgent(ctx, msg.AgentId)
		if !found {
			return nil, sdkerrors.Wrapf(types.ErrAgentNotFound, "agent ID %s not found", msg.AgentId)
		}
		hyperchain.AgentId = msg.AgentId
	}

	// Update the hyperchain
	hyperchain.Name = msg.Name
	hyperchain.Description = msg.Description
	if len(msg.ChainConfig) > 0 {
		hyperchain.ChainConfig = msg.ChainConfig
	}
	if len(msg.Metadata) > 0 {
		hyperchain.Metadata = msg.Metadata
	}
	if msg.MaxValidators > 0 {
		hyperchain.MaxValidators = msg.MaxValidators
	}
	if msg.MinStake > 0 {
		hyperchain.MinStake = msg.MinStake
	}
	if len(msg.SupportedTokens) > 0 {
		hyperchain.SupportedTokens = msg.SupportedTokens
	}
	if len(msg.SupportedModules) > 0 {
		hyperchain.SupportedModules = msg.SupportedModules
	}
	hyperchain.UpdatedAt = ctx.BlockTime()

	// Store the updated hyperchain
	k.SetHyperchain(ctx, hyperchain)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeUpdateHyperchain,
			sdk.NewAttribute(types.AttributeKeyHyperchainID, msg.ChainId),
			sdk.NewAttribute(types.AttributeKeyAdmin, msg.Admin),
			sdk.NewAttribute(types.AttributeKeyName, msg.Name),
		),
	)

	return &types.MsgUpdateHyperchainResponse{}, nil
}

// JoinHyperchainAsValidator joins a hyperchain as a validator
func (k msgServer) JoinHyperchainAsValidator(goCtx context.Context, msg *types.MsgJoinHyperchainAsValidator) (*types.MsgJoinHyperchainAsValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// Get the hyperchain
	hyperchain, found := k.GetHyperchain(ctx, msg.ChainId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrHyperchainNotFound, "hyperchain ID %s not found", msg.ChainId)
	}

	// Check if the hyperchain is active
	if hyperchain.Status != types.HyperchainStatusActive && hyperchain.Status != types.HyperchainStatusInitializing {
		return nil, sdkerrors.Wrapf(types.ErrHyperchainNotActive, "hyperchain is not active, current status: %s", hyperchain.Status)
	}

	// Parse the validator address
	validator, err := sdk.AccAddressFromBech32(msg.Validator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validator address (%s)", err)
	}

	// Check if the validator is already registered
	_, found = k.GetHyperchainValidator(ctx, msg.ChainId, msg.Validator)
	if found {
		return nil, sdkerrors.Wrapf(types.ErrValidatorAlreadyExists, "validator %s is already registered for hyperchain %s", msg.Validator, msg.ChainId)
	}

	// Check if the hyperchain has reached the maximum number of validators
	validators := k.GetHyperchainValidatorsByChain(ctx, msg.ChainId)
	if uint64(len(validators)) >= hyperchain.MaxValidators {
		return nil, sdkerrors.Wrapf(types.ErrMaxValidatorsReached, "hyperchain has reached the maximum number of validators (%d)", hyperchain.MaxValidators)
	}

	// Check stake
	minStake := k.MinValidatorStake(ctx)
	if !msg.Stake.IsGTE(minStake) {
		return nil, sdkerrors.Wrapf(types.ErrInsufficientStake, "stake %s is less than minimum %s", msg.Stake, minStake)
	}

	// If the hyperchain has a custom minimum stake, check that too
	if hyperchain.MinStake > 0 {
		if msg.Stake.Amount.LT(sdk.NewInt(int64(hyperchain.MinStake))) {
			return nil, sdkerrors.Wrapf(types.ErrInsufficientStake, "stake %s is less than hyperchain minimum %d", msg.Stake, hyperchain.MinStake)
		}
	}

	// Collect the stake
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, validator, types.ModuleName, sdk.NewCoins(msg.Stake))
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to collect stake")
	}

	// Create the validator
	validatorObj := types.HyperchainValidator{
		ChainId:  msg.ChainId,
		Address:  msg.Validator,
		Pubkey:   msg.Pubkey,
		Power:    uint64(msg.Stake.Amount.Int64()),
		Status:   types.ValidatorStatusActive,
		Stake:    msg.Stake,
		JoinedAt: ctx.BlockTime(),
		Metadata: msg.Metadata,
	}

	// Store the validator
	k.SetHyperchainValidator(ctx, validatorObj)

	// Update the hyperchain validators list
	hyperchain.Validators = append(hyperchain.Validators, msg.Validator)
	k.SetHyperchain(ctx, hyperchain)

	// Grant validator permission
	permission := types.HyperchainPermission{
		ChainId:        msg.ChainId,
		Address:        msg.Validator,
		PermissionType: types.PermissionTypeValidator,
		GrantedAt:      ctx.BlockTime(),
		ExpiresAt:      time.Time{}, // Never expires
		GrantedBy:      hyperchain.Admin,
		Metadata:       []byte("{}"),
	}
	k.SetHyperchainPermission(ctx, permission)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeJoinHyperchainAsValidator,
			sdk.NewAttribute(types.AttributeKeyHyperchainID, msg.ChainId),
			sdk.NewAttribute(types.AttributeKeyValidator, msg.Validator),
			sdk.NewAttribute(types.AttributeKeyStake, msg.Stake.String()),
		),
	)

	return &types.MsgJoinHyperchainAsValidatorResponse{}, nil
}

// LeaveHyperchain leaves a hyperchain as a validator
func (k msgServer) LeaveHyperchain(goCtx context.Context, msg *types.MsgLeaveHyperchain) (*types.MsgLeaveHyperchainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// Get the hyperchain
	hyperchain, found := k.GetHyperchain(ctx, msg.ChainId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrHyperchainNotFound, "hyperchain ID %s not found", msg.ChainId)
	}

	// Parse the validator address
	validator, err := sdk.AccAddressFromBech32(msg.Validator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validator address (%s)", err)
	}

	// Check if the validator is registered
	validatorObj, found := k.GetHyperchainValidator(ctx, msg.ChainId, msg.Validator)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrValidatorNotFound, "validator %s is not registered for hyperchain %s", msg.Validator, msg.ChainId)
	}

	// Return the stake
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, validator, sdk.NewCoins(validatorObj.Stake))
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to return stake")
	}

	// Delete the validator
	k.DeleteHyperchainValidator(ctx, msg.ChainId, msg.Validator)

	// Update the hyperchain validators list
	for i, v := range hyperchain.Validators {
		if v == msg.Validator {
			hyperchain.Validators = append(hyperchain.Validators[:i], hyperchain.Validators[i+1:]...)
			break
		}
	}
	k.SetHyperchain(ctx, hyperchain)

	// Revoke validator permission
	k.DeleteHyperchainPermission(ctx, msg.ChainId, msg.Validator, types.PermissionTypeValidator)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeLeaveHyperchain,
			sdk.NewAttribute(types.AttributeKeyHyperchainID, msg.ChainId),
			sdk.NewAttribute(types.AttributeKeyValidator, msg.Validator),
		),
	)

	return &types.MsgLeaveHyperchainResponse{}, nil
}

// CreateHyperchainBridge creates a new bridge between hyperchains
func (k msgServer) CreateHyperchainBridge(goCtx context.Context, msg *types.MsgCreateHyperchainBridge) (*types.MsgCreateHyperchainBridgeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// Parse the creator address
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	// Check if the source hyperchain exists
	sourceChain, found := k.GetHyperchain(ctx, msg.SourceChainId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrHyperchainNotFound, "source hyperchain ID %s not found", msg.SourceChainId)
	}

	// Check if the target hyperchain exists
	targetChain, found := k.GetHyperchain(ctx, msg.TargetChainId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrHyperchainNotFound, "target hyperchain ID %s not found", msg.TargetChainId)
	}

	// Check if the creator has admin permission on both hyperchains
	if sourceChain.Admin != creator.String() && !k.HasPermission(ctx, msg.SourceChainId, creator.String(), types.PermissionTypeAdmin) {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, "creator does not have admin permission on source hyperchain")
	}
	if targetChain.Admin != creator.String() && !k.HasPermission(ctx, msg.TargetChainId, creator.String(), types.PermissionTypeAdmin) {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, "creator does not have admin permission on target hyperchain")
	}

	// Check if the hyperchains have reached the maximum number of bridges
	sourceBridges := k.GetHyperchainBridgesByChain(ctx, msg.SourceChainId)
	maxBridgesPerHyperchain := k.MaxBridgesPerHyperchain(ctx)
	if uint64(len(sourceBridges)) >= maxBridgesPerHyperchain {
		return nil, sdkerrors.Wrapf(types.ErrMaxBridgesReached, "source hyperchain has reached the maximum number of bridges (%d)", maxBridgesPerHyperchain)
	}

	// Generate a unique ID for the bridge
	id := uuid.New().String()

	// Create the bridge
	bridge := types.HyperchainBridge{
		Id:              id,
		SourceChainId:   msg.SourceChainId,
		TargetChainId:   msg.TargetChainId,
		Status:          types.BridgeStatusActive,
		Creator:         msg.Creator,
		Admin:           msg.Creator,
		CreatedAt:       ctx.BlockTime(),
		UpdatedAt:       ctx.BlockTime(),
		Relayers:        []string{},
		MinRelayers:     msg.MinRelayers,
		SupportedTokens: msg.SupportedTokens,
		Metadata:        msg.Metadata,
	}

	// Store the bridge
	k.SetHyperchainBridge(ctx, bridge)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCreateHyperchainBridge,
			sdk.NewAttribute(types.AttributeKeyBridgeID, id),
			sdk.NewAttribute(types.AttributeKeyCreator, msg.Creator),
			sdk.NewAttribute(types.AttributeKeySourceChainID, msg.SourceChainId),
			sdk.NewAttribute(types.AttributeKeyTargetChainID, msg.TargetChainId),
		),
	)

	return &types.MsgCreateHyperchainBridgeResponse{
		Id: id,
	}, nil
}

// UpdateHyperchainBridge updates an existing hyperchain bridge
func (k msgServer) UpdateHyperchainBridge(goCtx context.Context, msg *types.MsgUpdateHyperchainBridge) (*types.MsgUpdateHyperchainBridgeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// Get the bridge
	bridge, found := k.GetHyperchainBridge(ctx, msg.BridgeId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrBridgeNotFound, "bridge ID %s not found", msg.BridgeId)
	}

	// Parse the admin address
	admin, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}

	// Check if the caller is the admin
	if bridge.Admin != admin.String() {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, "only the admin can update the bridge")
	}

	// Update the bridge
	if msg.MinRelayers > 0 {
		bridge.MinRelayers = msg.MinRelayers
	}
	if len(msg.SupportedTokens) > 0 {
		bridge.SupportedTokens = msg.SupportedTokens
	}
	if len(msg.Metadata) > 0 {
		bridge.Metadata = msg.Metadata
	}
	bridge.UpdatedAt = ctx.BlockTime()

	// Store the updated bridge
	k.SetHyperchainBridge(ctx, bridge)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeUpdateHyperchainBridge,
			sdk.NewAttribute(types.AttributeKeyBridgeID, msg.BridgeId),
			sdk.NewAttribute(types.AttributeKeyAdmin, msg.Admin),
		),
	)

	return &types.MsgUpdateHyperchainBridgeResponse{}, nil
}

// RegisterHyperchainBridgeRelayer registers a relayer for a hyperchain bridge
func (k msgServer) RegisterHyperchainBridgeRelayer(goCtx context.Context, msg *types.MsgRegisterHyperchainBridgeRelayer) (*types.MsgRegisterHyperchainBridgeRelayerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// Get the bridge
	bridge, found := k.GetHyperchainBridge(ctx, msg.BridgeId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrBridgeNotFound, "bridge ID %s not found", msg.BridgeId)
	}

	// Parse the admin address
	admin, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}

	// Check if the caller is the admin
	if bridge.Admin != admin.String() {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, "only the admin can register relayers")
	}

	// Parse the relayer address
	_, err = sdk.AccAddressFromBech32(msg.Relayer)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid relayer address (%s)", err)
	}

	// Check if the relayer is already registered
	for _, r := range bridge.Relayers {
		if r == msg.Relayer {
			return nil, sdkerrors.Wrapf(types.ErrRelayerAlreadyExists, "relayer %s is already registered for bridge %s", msg.Relayer, msg.BridgeId)
		}
	}

	// Add the relayer
	bridge.Relayers = append(bridge.Relayers, msg.Relayer)
	bridge.UpdatedAt = ctx.BlockTime()

	// Store the updated bridge
	k.SetHyperchainBridge(ctx, bridge)

	// Grant relayer permission
	permission := types.HyperchainPermission{
		ChainId:        bridge.SourceChainId,
		Address:        msg.Relayer,
		PermissionType: types.PermissionTypeRelayer,
		GrantedAt:      ctx.BlockTime(),
		ExpiresAt:      time.Time{}, // Never expires
		GrantedBy:      msg.Admin,
		Metadata:       []byte("{}"),
	}
	k.SetHyperchainPermission(ctx, permission)

	permission = types.HyperchainPermission{
		ChainId:        bridge.TargetChainId,
		Address:        msg.Relayer,
		PermissionType: types.PermissionTypeRelayer,
		GrantedAt:      ctx.BlockTime(),
		ExpiresAt:      time.Time{}, // Never expires
		GrantedBy:      msg.Admin,
		Metadata:       []byte("{}"),
	}
	k.SetHyperchainPermission(ctx, permission)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRegisterHyperchainBridgeRelayer,
			sdk.NewAttribute(types.AttributeKeyBridgeID, msg.BridgeId),
			sdk.NewAttribute(types.AttributeKeyAdmin, msg.Admin),
			sdk.NewAttribute(types.AttributeKeyRelayer, msg.Relayer),
		),
	)

	return &types.MsgRegisterHyperchainBridgeRelayerResponse{}, nil
}

// RemoveHyperchainBridgeRelayer removes a relayer from a hyperchain bridge
func (k msgServer) RemoveHyperchainBridgeRelayer(goCtx context.Context, msg *types.MsgRemoveHyperchainBridgeRelayer) (*types.MsgRemoveHyperchainBridgeRelayerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// Get the bridge
	bridge, found := k.GetHyperchainBridge(ctx, msg.BridgeId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrBridgeNotFound, "bridge ID %s not found", msg.BridgeId)
	}

	// Parse the admin address
	admin, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}

	// Check if the caller is the admin
	if bridge.Admin != admin.String() {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, "only the admin can remove relayers")
	}

	// Parse the relayer address
	_, err = sdk.AccAddressFromBech32(msg.Relayer)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid relayer address (%s)", err)
	}

	// Check if the relayer is registered
	found = false
	for i, r := range bridge.Relayers {
		if r == msg.Relayer {
			bridge.Relayers = append(bridge.Relayers[:i], bridge.Relayers[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrRelayerNotFound, "relayer %s is not registered for bridge %s", msg.Relayer, msg.BridgeId)
	}

	bridge.UpdatedAt = ctx.BlockTime()

	// Store the updated bridge
	k.SetHyperchainBridge(ctx, bridge)

	// Revoke relayer permission
	k.DeleteHyperchainPermission(ctx, bridge.SourceChainId, msg.Relayer, types.PermissionTypeRelayer)
	k.DeleteHyperchainPermission(ctx, bridge.TargetChainId, msg.Relayer, types.PermissionTypeRelayer)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRemoveHyperchainBridgeRelayer,
			sdk.NewAttribute(types.AttributeKeyBridgeID, msg.BridgeId),
			sdk.NewAttribute(types.AttributeKeyAdmin, msg.Admin),
			sdk.NewAttribute(types.AttributeKeyRelayer, msg.Relayer),
		),
	)

	return &types.MsgRemoveHyperchainBridgeRelayerResponse{}, nil
}

// InitiateHyperchainBridgeTransaction initiates a transaction through a hyperchain bridge
func (k msgServer) InitiateHyperchainBridgeTransaction(goCtx context.Context, msg *types.MsgInitiateHyperchainBridgeTransaction) (*types.MsgInitiateHyperchainBridgeTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// Get the bridge
	bridge, found := k.GetHyperchainBridge(ctx, msg.BridgeId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrBridgeNotFound, "bridge ID %s not found", msg.BridgeId)
	}

	// Check if the bridge is active
	if bridge.Status != types.BridgeStatusActive {
		return nil, sdkerrors.Wrapf(types.ErrBridgeNotActive, "bridge is not active, current status: %s", bridge.Status)
	}

	// Parse the sender address
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	// Parse the recipient address
	_, err = sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid recipient address (%s)", err)
	}

	// Check if the token is supported
	tokenSupported := false
	for _, token := range bridge.SupportedTokens {
		if token == msg.Amount.Denom {
			tokenSupported = true
			break
		}
	}
	if !tokenSupported {
		return nil, sdkerrors.Wrapf(types.ErrUnsupportedToken, "token %s is not supported by the bridge", msg.Amount.Denom)
	}

	// Collect the amount
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.NewCoins(msg.Amount))
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to collect amount")
	}

	// Generate a unique ID for the transaction
	id := uuid.New().String()

	// Create the transaction
	tx := types.HyperchainBridgeTransaction{
		Id:            id,
		BridgeId:      msg.BridgeId,
		SourceChainId: bridge.SourceChainId,
		TargetChainId: bridge.TargetChainId,
		Sender:        msg.Sender,
		Recipient:     msg.Recipient,
		Amount:        msg.Amount,
		Status:        types.BridgeTransactionStatusPending,
		CreatedAt:     ctx.BlockTime(),
		CompletedAt:   nil,
		SourceTxId:    msg.SourceTxId,
		TargetTxId:    "",
		Approvals:     []string{},
		Metadata:      msg.Metadata,
	}

	// Store the transaction
	k.SetHyperchainBridgeTransaction(ctx, tx)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeInitiateHyperchainBridgeTransaction,
			sdk.NewAttribute(types.AttributeKeyBridgeID, msg.BridgeId),
			sdk.NewAttribute(types.AttributeKeyTransactionID, id),
			sdk.NewAttribute(types.AttributeKeySender, msg.Sender),
			sdk.NewAttribute(types.AttributeKeyRecipient, msg.Recipient),
			sdk.NewAttribute(types.AttributeKeyAmount, msg.Amount.String()),
		),
	)

	return &types.MsgInitiateHyperchainBridgeTransactionResponse{
		Id: id,
	}, nil
}

// ApproveHyperchainBridgeTransaction approves a transaction through a hyperchain bridge
func (k msgServer) ApproveHyperchainBridgeTransaction(goCtx context.Context, msg *types.MsgApproveHyperchainBridgeTransaction) (*types.MsgApproveHyperchainBridgeTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// Get the bridge
	bridge, found := k.GetHyperchainBridge(ctx, msg.BridgeId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrBridgeNotFound, "bridge ID %s not found", msg.BridgeId)
	}

	// Check if the bridge is active
	if bridge.Status != types.BridgeStatusActive {
		return nil, sdkerrors.Wrapf(types.ErrBridgeNotActive, "bridge is not active, current status: %s", bridge.Status)
	}

	// Parse the relayer address
	relayer, err := sdk.AccAddressFromBech32(msg.Relayer)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid relayer address (%s)", err)
	}

	// Check if the relayer is registered
	relayerFound := false
	for _, r := range bridge.Relayers {
		if r == msg.Relayer {
			relayerFound = true
			break
		}
	}
	if !relayerFound {
		return nil, sdkerrors.Wrapf(types.ErrRelayerNotFound, "relayer %s is not registered for bridge %s", msg.Relayer, msg.BridgeId)
	}

	// Get the transaction
	tx, found := k.GetHyperchainBridgeTransaction(ctx, msg.TxId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrTransactionNotFound, "transaction ID %s not found", msg.TxId)
	}

	// Check if the transaction is pending
	if tx.Status != types.BridgeTransactionStatusPending {
		return nil, sdkerrors.Wrapf(types.ErrTransactionNotPending, "transaction is not pending, current status: %s", tx.Status)
	}

	// Check if the relayer has already approved
	for _, approval := range tx.Approvals {
		if approval == msg.Relayer {
			return nil, sdkerrors.Wrapf(types.ErrAlreadyApproved, "relayer %s has already approved transaction %s", msg.Relayer, msg.TxId)
		}
	}

	// Add the approval
	tx.Approvals = append(tx.Approvals, msg.Relayer)

	// Check if the transaction has enough approvals
	if uint64(len(tx.Approvals)) >= bridge.MinRelayers {
		// Complete the transaction
		tx.Status = types.BridgeTransactionStatusCompleted
		completedAt := ctx.BlockTime()
		tx.CompletedAt = &completedAt

		// Generate a target transaction ID
		targetTxId := uuid.New().String()
		tx.TargetTxId = targetTxId

		// Transfer the amount to the recipient
		recipient, _ := sdk.AccAddressFromBech32(tx.Recipient)
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, recipient, sdk.NewCoins(tx.Amount))
		if err != nil {
			return nil, sdkerrors.Wrapf(err, "failed to transfer amount to recipient")
		}

		// Emit completion event
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeCompleteHyperchainBridgeTransaction,
				sdk.NewAttribute(types.AttributeKeyBridgeID, tx.BridgeId),
				sdk.NewAttribute(types.AttributeKeyTransactionID, tx.Id),
				sdk.NewAttribute(types.AttributeKeyTargetTransactionID, targetTxId),
				sdk.NewAttribute(types.AttributeKeyRecipient, tx.Recipient),
				sdk.NewAttribute(types.AttributeKeyAmount, tx.Amount.String()),
			),
		)
	}

	// Store the updated transaction
	k.SetHyperchainBridgeTransaction(ctx, tx)

	// Emit approval event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeApproveHyperchainBridgeTransaction,
			sdk.NewAttribute(types.AttributeKeyBridgeID, msg.BridgeId),
			sdk.NewAttribute(types.AttributeKeyTransactionID, msg.TxId),
			sdk.NewAttribute(types.AttributeKeyRelayer, msg.Relayer),
		),
	)

	return &types.MsgApproveHyperchainBridgeTransactionResponse{
		TargetTxId: tx.TargetTxId,
	}, nil
}

// SubmitHyperchainBlock submits a block to a hyperchain
func (k msgServer) SubmitHyperchainBlock(goCtx context.Context, msg *types.MsgSubmitHyperchainBlock) (*types.MsgSubmitHyperchainBlockResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// Get the hyperchain
	hyperchain, found := k.GetHyperchain(ctx, msg.ChainId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrHyperchainNotFound, "hyperchain ID %s not found", msg.ChainId)
	}

	// Check if the hyperchain is active
	if hyperchain.Status != types.HyperchainStatusActive {
		return nil, sdkerrors.Wrapf(types.ErrHyperchainNotActive, "hyperchain is not active, current status: %s", hyperchain.Status)
	}

	// Parse the proposer address
	proposer, err := sdk.AccAddressFromBech32(msg.Proposer)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid proposer address (%s)", err)
	}

	// Check if the proposer is a validator
	_, found = k.GetHyperchainValidator(ctx, msg.ChainId, msg.Proposer)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrNotValidator, "proposer %s is not a validator for hyperchain %s", msg.Proposer, msg.ChainId)
	}

	// Check if the block height is valid
	if msg.Height != hyperchain.BlockHeight+1 {
		return nil, sdkerrors.Wrapf(types.ErrInvalidBlockHeight, "expected block height %d, got %d", hyperchain.BlockHeight+1, msg.Height)
	}

	// Check if the parent hash is valid
	if hyperchain.BlockHeight > 0 {
		parentBlock, found := k.GetHyperchainBlock(ctx, msg.ChainId, hyperchain.BlockHeight)
		if !found {
			return nil, sdkerrors.Wrapf(types.ErrParentBlockNotFound, "parent block at height %d not found", hyperchain.BlockHeight)
		}
		if parentBlock.Hash != msg.ParentHash {
			return nil, sdkerrors.Wrapf(types.ErrInvalidParentHash, "expected parent hash %s, got %s", parentBlock.Hash, msg.ParentHash)
		}
	}

	// Create the block
	block := types.HyperchainBlock{
		ChainId:    msg.ChainId,
		Height:     msg.Height,
		Hash:       "",
		ParentHash: msg.ParentHash,
		Timestamp:  ctx.BlockTime(),
		NumTxs:     msg.NumTxs,
		Proposer:   msg.Proposer,
		Data:       msg.Data,
	}

	// Calculate the block hash
	block.Hash = k.CalculateBlockHash(block)

	// Store the block
	k.SetHyperchainBlock(ctx, block)

	// Update the hyperchain
	hyperchain.BlockHeight = msg.Height
	hyperchain.UpdatedAt = ctx.BlockTime()
	k.SetHyperchain(ctx, hyperchain)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSubmitHyperchainBlock,
			sdk.NewAttribute(types.AttributeKeyHyperchainID, msg.ChainId),
			sdk.NewAttribute(types.AttributeKeyBlockHeight, fmt.Sprintf("%d", msg.Height)),
			sdk.NewAttribute(types.AttributeKeyBlockHash, block.Hash),
			sdk.NewAttribute(types.AttributeKeyProposer, msg.Proposer),
		),
	)

	return &types.MsgSubmitHyperchainBlockResponse{
		Hash: block.Hash,
	}, nil
}

// SubmitHyperchainTransaction submits a transaction to a hyperchain
func (k msgServer) SubmitHyperchainTransaction(goCtx context.Context, msg *types.MsgSubmitHyperchainTransaction) (*types.MsgSubmitHyperchainTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// Get the hyperchain
	hyperchain, found := k.GetHyperchain(ctx, msg.ChainId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrHyperchainNotFound, "hyperchain ID %s not found", msg.ChainId)
	}

	// Check if the hyperchain is active
	if hyperchain.Status != types.HyperchainStatusActive {
		return nil, sdkerrors.Wrapf(types.ErrHyperchainNotActive, "hyperchain is not active, current status: %s", hyperchain.Status)
	}

	// Parse the sender address
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	// Parse the recipient address
	_, err = sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid recipient address (%s)", err)
	}

	// Check if the token is supported
	tokenSupported := false
	for _, token := range hyperchain.SupportedTokens {
		if token == msg.Amount.Denom {
			tokenSupported = true
			break
		}
	}
	if !tokenSupported {
		return nil, sdkerrors.Wrapf(types.ErrUnsupportedToken, "token %s is not supported by the hyperchain", msg.Amount.Denom)
	}

	// Generate a unique ID for the transaction
	id := uuid.New().String()

	// Create the transaction
	tx := types.HyperchainTransaction{
		Id:          id,
		ChainId:     msg.ChainId,
		BlockHeight: hyperchain.BlockHeight,
		Sender:      msg.Sender,
		Recipient:   msg.Recipient,
		Amount:      msg.Amount,
		Status:      types.TransactionStatusSuccess,
		Timestamp:   ctx.BlockTime(),
		Data:        msg.Data,
		GasUsed:     1000, // Dummy value
		Error:       "",
	}

	// Store the transaction
	k.SetHyperchainTransaction(ctx, tx)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSubmitHyperchainTransaction,
			sdk.NewAttribute(types.AttributeKeyHyperchainID, msg.ChainId),
			sdk.NewAttribute(types.AttributeKeyTransactionID, id),
			sdk.NewAttribute(types.AttributeKeySender, msg.Sender),
			sdk.NewAttribute(types.AttributeKeyRecipient, msg.Recipient),
			sdk.NewAttribute(types.AttributeKeyAmount, msg.Amount.String()),
		),
	)

	return &types.MsgSubmitHyperchainTransactionResponse{
		Id: id,
	}, nil
}

// GrantHyperchainPermission grants permission to a hyperchain
func (k msgServer) GrantHyperchainPermission(goCtx context.Context, msg *types.MsgGrantHyperchainPermission) (*types.MsgGrantHyperchainPermissionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// Get the hyperchain
	hyperchain, found := k.GetHyperchain(ctx, msg.ChainId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrHyperchainNotFound, "hyperchain ID %s not found", msg.ChainId)
	}

	// Parse the admin address
	admin, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}

	// Check if the caller is the admin
	if hyperchain.Admin != admin.String() && !k.HasPermission(ctx, msg.ChainId, admin.String(), types.PermissionTypeAdmin) {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, "only the admin can grant permissions")
	}

	// Parse the address to grant permission to
	_, err = sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address (%s)", err)
	}

	// Check if the permission type is valid
	if !types.IsValidPermissionType(msg.PermissionType) {
		return nil, sdkerrors.Wrapf(types.ErrInvalidPermissionType, "invalid permission type: %s", msg.PermissionType)
	}

	// Calculate expiration time
	var expiresAt time.Time
	if msg.ExpirationDays > 0 {
		expiresAt = ctx.BlockTime().Add(time.Duration(msg.ExpirationDays) * 24 * time.Hour)
	}

	// Create the permission
	permission := types.HyperchainPermission{
		ChainId:        msg.ChainId,
		Address:        msg.Address,
		PermissionType: msg.PermissionType,
		GrantedAt:      ctx.BlockTime(),
		ExpiresAt:      &expiresAt,
		GrantedBy:      msg.Admin,
		Metadata:       msg.Metadata,
	}

	// Store the permission
	k.SetHyperchainPermission(ctx, permission)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeGrantHyperchainPermission,
			sdk.NewAttribute(types.AttributeKeyHyperchainID, msg.ChainId),
			sdk.NewAttribute(types.AttributeKeyAdmin, msg.Admin),
			sdk.NewAttribute(types.AttributeKeyAddress, msg.Address),
			sdk.NewAttribute(types.AttributeKeyPermissionType, msg.PermissionType),
			sdk.NewAttribute(types.AttributeKeyExpirationDays, fmt.Sprintf("%d", msg.ExpirationDays)),
		),
	)

	return &types.MsgGrantHyperchainPermissionResponse{}, nil
}

// RevokeHyperchainPermission revokes permission from a hyperchain
func (k msgServer) RevokeHyperchainPermission(goCtx context.Context, msg *types.MsgRevokeHyperchainPermission) (*types.MsgRevokeHyperchainPermissionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// Get the hyperchain
	hyperchain, found := k.GetHyperchain(ctx, msg.ChainId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrHyperchainNotFound, "hyperchain ID %s not found", msg.ChainId)
	}

	// Parse the admin address
	admin, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}

	// Check if the caller is the admin
	if hyperchain.Admin != admin.String() && !k.HasPermission(ctx, msg.ChainId, admin.String(), types.PermissionTypeAdmin) {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, "only the admin can revoke permissions")
	}

	// Parse the address to revoke permission from
	_, err = sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address (%s)", err)
	}

	// Check if the permission type is valid
	if !types.IsValidPermissionType(msg.PermissionType) {
		return nil, sdkerrors.Wrapf(types.ErrInvalidPermissionType, "invalid permission type: %s", msg.PermissionType)
	}

	// Check if the permission exists
	_, found = k.GetHyperchainPermission(ctx, msg.ChainId, msg.Address, msg.PermissionType)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrPermissionNotFound, "permission not found for hyperchain %s, address %s, type %s", msg.ChainId, msg.Address, msg.PermissionType)
	}

	// Delete the permission
	k.DeleteHyperchainPermission(ctx, msg.ChainId, msg.Address, msg.PermissionType)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRevokeHyperchainPermission,
			sdk.NewAttribute(types.AttributeKeyHyperchainID, msg.ChainId),
			sdk.NewAttribute(types.AttributeKeyAdmin, msg.Admin),
			sdk.NewAttribute(types.AttributeKeyAddress, msg.Address),
			sdk.NewAttribute(types.AttributeKeyPermissionType, msg.PermissionType),
		),
	)

	return &types.MsgRevokeHyperchainPermissionResponse{}, nil
}