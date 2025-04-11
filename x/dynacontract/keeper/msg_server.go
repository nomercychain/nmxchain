package keeper

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/nomercychain/nmxchain/x/dynacontract/types"
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

// CreateDynaContract creates a new dynamic contract
func (k msgServer) CreateDynaContract(goCtx context.Context, msg *types.MsgCreateDynaContract) (*types.MsgCreateDynaContractResponse, error) {
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

	// Check code size
	maxCodeSize := k.MaxContractSize(ctx)
	if uint64(len(msg.Code)) > maxCodeSize {
		return nil, sdkerrors.Wrapf(types.ErrContractCodeTooLarge, "code size %d exceeds maximum %d", len(msg.Code), maxCodeSize)
	}

	// Check gas limit
	maxGas := k.MaxContractGas(ctx)
	if msg.GasLimit > maxGas {
		return nil, sdkerrors.Wrapf(types.ErrGasLimitTooHigh, "gas limit %d exceeds maximum %d", msg.GasLimit, maxGas)
	}

	// Check metadata size
	maxMetadataSize := k.MaxMetadataSize(ctx)
	if uint64(len(msg.Metadata)) > maxMetadataSize {
		return nil, sdkerrors.Wrapf(types.ErrMetadataTooLarge, "metadata size %d exceeds maximum %d", len(msg.Metadata), maxMetadataSize)
	}

	// Check deposit
	minDeposit := k.MinContractDeposit(ctx)
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

	// Generate a unique ID for the contract
	id := uuid.New().String()

	// Calculate the code hash
	codeHash := k.CalculateCodeHash(msg.Code)

	// Create the contract
	contract := types.DynaContract{
		Id:             id,
		Name:           msg.Name,
		Description:    msg.Description,
		Creator:        msg.Creator,
		Owner:          msg.Creator,
		ContractType:   msg.ContractType,
		Status:         types.DynaContractStatusActive,
		CodeHash:       codeHash,
		Code:           msg.Code,
		Abi:            msg.Abi,
		AgentId:        msg.AgentId,
		State:          []byte("{}"),
		Metadata:       msg.Metadata,
		CreatedAt:      ctx.BlockTime(),
		UpdatedAt:      ctx.BlockTime(),
		Tags:           msg.Tags,
		GasLimit:       msg.GasLimit,
		ExecutionCount: 0,
	}

	// Store the contract
	k.SetDynaContract(ctx, contract)

	// Grant execute permission to the creator
	permission := types.DynaContractPermission{
		ContractId:     id,
		Address:        msg.Creator,
		PermissionType: types.PermissionTypeExecute,
		GrantedAt:      ctx.BlockTime(),
		ExpiresAt:      time.Time{}, // Never expires
		GrantedBy:      msg.Creator,
		Metadata:       []byte("{}"),
	}
	k.SetDynaContractPermission(ctx, permission)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCreateDynaContract,
			sdk.NewAttribute(types.AttributeKeyContractID, id),
			sdk.NewAttribute(types.AttributeKeyCreator, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyName, msg.Name),
			sdk.NewAttribute(types.AttributeKeyContractType, msg.ContractType.String()),
			sdk.NewAttribute(types.AttributeKeyDeposit, msg.Deposit.String()),
		),
	)

	return &types.MsgCreateDynaContractResponse{
		Id: id,
	}, nil
}

// UpdateDynaContract updates an existing dynamic contract
func (k msgServer) UpdateDynaContract(goCtx context.Context, msg *types.MsgUpdateDynaContract) (*types.MsgUpdateDynaContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// Get the contract
	contract, found := k.GetDynaContract(ctx, msg.ContractId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrContractNotFound, "contract ID %s not found", msg.ContractId)
	}

	// Parse the owner address
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}

	// Check if the caller is the owner
	if contract.Owner != owner.String() {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, "only the owner can update the contract")
	}

	// Check code size if provided
	if len(msg.Code) > 0 {
		maxCodeSize := k.MaxContractSize(ctx)
		if uint64(len(msg.Code)) > maxCodeSize {
			return nil, sdkerrors.Wrapf(types.ErrContractCodeTooLarge, "code size %d exceeds maximum %d", len(msg.Code), maxCodeSize)
		}
		// Update code and code hash
		contract.Code = msg.Code
		contract.CodeHash = k.CalculateCodeHash(msg.Code)
	}

	// Check gas limit if provided
	if msg.GasLimit > 0 {
		maxGas := k.MaxContractGas(ctx)
		if msg.GasLimit > maxGas {
			return nil, sdkerrors.Wrapf(types.ErrGasLimitTooHigh, "gas limit %d exceeds maximum %d", msg.GasLimit, maxGas)
		}
		contract.GasLimit = msg.GasLimit
	}

	// Check metadata size if provided
	if len(msg.Metadata) > 0 {
		maxMetadataSize := k.MaxMetadataSize(ctx)
		if uint64(len(msg.Metadata)) > maxMetadataSize {
			return nil, sdkerrors.Wrapf(types.ErrMetadataTooLarge, "metadata size %d exceeds maximum %d", len(msg.Metadata), maxMetadataSize)
		}
		contract.Metadata = msg.Metadata
	}

	// If an agent ID is provided, check if it exists
	if msg.AgentId != "" && msg.AgentId != contract.AgentId {
		_, found := k.deaiKeeper.GetAIAgent(ctx, msg.AgentId)
		if !found {
			return nil, sdkerrors.Wrapf(types.ErrAgentNotFound, "agent ID %s not found", msg.AgentId)
		}
		contract.AgentId = msg.AgentId
	}

	// Update the contract
	contract.Name = msg.Name
	contract.Description = msg.Description
	if len(msg.Abi) > 0 {
		contract.Abi = msg.Abi
	}
	if len(msg.Tags) > 0 {
		contract.Tags = msg.Tags
	}
	contract.UpdatedAt = ctx.BlockTime()

	// Store the updated contract
	k.SetDynaContract(ctx, contract)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeUpdateDynaContract,
			sdk.NewAttribute(types.AttributeKeyContractID, msg.ContractId),
			sdk.NewAttribute(types.AttributeKeyOwner, msg.Owner),
			sdk.NewAttribute(types.AttributeKeyName, msg.Name),
		),
	)

	return &types.MsgUpdateDynaContractResponse{}, nil
}

// ExecuteDynaContract executes a dynamic contract
func (k msgServer) ExecuteDynaContract(goCtx context.Context, msg *types.MsgExecuteDynaContract) (*types.MsgExecuteDynaContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// Parse the caller address
	caller, err := sdk.AccAddressFromBech32(msg.Caller)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid caller address (%s)", err)
	}

	// Get the contract
	contract, found := k.GetDynaContract(ctx, msg.ContractId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrContractNotFound, "contract ID %s not found", msg.ContractId)
	}

	// Check if the contract is active
	if contract.Status != types.DynaContractStatusActive {
		return nil, sdkerrors.Wrapf(types.ErrContractNotActive, "contract is not active, current status: %s", contract.Status)
	}

	// Check if the caller has permission to execute the contract
	if !k.HasPermission(ctx, msg.ContractId, msg.Caller, types.PermissionTypeExecute) && contract.Owner != msg.Caller {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, "caller does not have execute permission")
	}

	// Check gas limit
	gasLimit := msg.GasLimit
	if gasLimit == 0 {
		gasLimit = contract.GasLimit
	}
	if gasLimit > contract.GasLimit {
		return nil, sdkerrors.Wrapf(types.ErrGasLimitTooHigh, "gas limit %d exceeds contract limit %d", gasLimit, contract.GasLimit)
	}

	// Collect the fee
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, caller, types.ModuleName, sdk.NewCoins(msg.Fee))
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to collect fee")
	}

	// Execute the contract
	output, gasUsed, err := k.ExecuteDynaContract(ctx, msg.ContractId, caller, msg.Input, gasLimit)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to execute contract")
	}

	// Calculate the fee to distribute to the contract owner
	feeRate := k.ExecutionFeeRate(ctx)
	ownerFee := sdk.NewCoin(msg.Fee.Denom, msg.Fee.Amount.MulRaw(feeRate.MulInt64(100).TruncateInt64()).QuoRaw(100))

	// Distribute the fee to the contract owner
	ownerAddr, err := sdk.AccAddressFromBech32(contract.Owner)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, ownerAddr, sdk.NewCoins(ownerFee))
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to distribute fee to owner")
	}

	// Get the execution ID
	executionID := fmt.Sprintf("%s-%d", msg.ContractId, contract.ExecutionCount)

	return &types.MsgExecuteDynaContractResponse{
		ExecutionId: executionID,
		Output:      output,
		GasUsed:     gasUsed,
	}, nil
}

// AddLearningData adds learning data to a dynamic contract
func (k msgServer) AddLearningData(goCtx context.Context, msg *types.MsgAddLearningData) (*types.MsgAddLearningDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// Get the contract
	contract, found := k.GetDynaContract(ctx, msg.ContractId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrContractNotFound, "contract ID %s not found", msg.ContractId)
	}

	// Parse the owner address
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}

	// Check if the caller is the owner
	if contract.Owner != owner.String() {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, "only the owner can add learning data")
	}

	// Check if the contract type supports learning
	if contract.ContractType != types.DynaContractTypeLearning && contract.ContractType != types.DynaContractTypeAdaptive {
		return nil, sdkerrors.Wrapf(types.ErrContractTypeNotSupported, "contract type %s does not support learning", contract.ContractType.String())
	}

	// Check data size
	maxDataSize := k.MaxLearningDataSize(ctx)
	if uint64(len(msg.Data)) > maxDataSize {
		return nil, sdkerrors.Wrapf(types.ErrLearningDataTooLarge, "data size %d exceeds maximum %d", len(msg.Data), maxDataSize)
	}

	// Generate a unique ID for the learning data
	dataID := uuid.New().String()

	// Create the learning data
	learningData := types.DynaContractLearningData{
		Id:         dataID,
		ContractId: msg.ContractId,
		DataType:   msg.DataType,
		Data:       msg.Data,
		Source:     msg.Source,
		Timestamp:  ctx.BlockTime(),
		Metadata:   msg.Metadata,
	}

	// Store the learning data
	k.SetDynaContractLearningData(ctx, learningData)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeAddLearningData,
			sdk.NewAttribute(types.AttributeKeyContractID, msg.ContractId),
			sdk.NewAttribute(types.AttributeKeyOwner, msg.Owner),
			sdk.NewAttribute(types.AttributeKeyDataID, dataID),
			sdk.NewAttribute(types.AttributeKeyDataType, msg.DataType),
		),
	)

	return &types.MsgAddLearningDataResponse{
		DataId: dataID,
	}, nil
}

// CreateDynaContractTemplate creates a new dynamic contract template
func (k msgServer) CreateDynaContractTemplate(goCtx context.Context, msg *types.MsgCreateDynaContractTemplate) (*types.MsgCreateDynaContractTemplateResponse, error) {
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

	// Check code size
	maxCodeSize := k.MaxContractSize(ctx)
	if uint64(len(msg.Code)) > maxCodeSize {
		return nil, sdkerrors.Wrapf(types.ErrContractCodeTooLarge, "code size %d exceeds maximum %d", len(msg.Code), maxCodeSize)
	}

	// Check metadata size
	maxMetadataSize := k.MaxMetadataSize(ctx)
	if uint64(len(msg.Metadata)) > maxMetadataSize {
		return nil, sdkerrors.Wrapf(types.ErrMetadataTooLarge, "metadata size %d exceeds maximum %d", len(msg.Metadata), maxMetadataSize)
	}

	// Generate a unique ID for the template
	id := uuid.New().String()

	// Create the template
	template := types.DynaContractTemplate{
		Id:           id,
		Name:         msg.Name,
		Description:  msg.Description,
		Creator:      msg.Creator,
		ContractType: msg.ContractType,
		Code:         msg.Code,
		Abi:          msg.Abi,
		Metadata:     msg.Metadata,
		CreatedAt:    ctx.BlockTime(),
		UpdatedAt:    ctx.BlockTime(),
		Tags:         msg.Tags,
		UsageCount:   0,
	}

	// Store the template
	k.SetDynaContractTemplate(ctx, template)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCreateDynaContractTemplate,
			sdk.NewAttribute(types.AttributeKeyTemplateID, id),
			sdk.NewAttribute(types.AttributeKeyCreator, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyName, msg.Name),
			sdk.NewAttribute(types.AttributeKeyContractType, msg.ContractType.String()),
		),
	)

	return &types.MsgCreateDynaContractTemplateResponse{
		Id: id,
	}, nil
}

// InstantiateFromTemplate creates a new dynamic contract from a template
func (k msgServer) InstantiateFromTemplate(goCtx context.Context, msg *types.MsgInstantiateFromTemplate) (*types.MsgInstantiateFromTemplateResponse, error) {
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

	// Get the template
	template, found := k.GetDynaContractTemplate(ctx, msg.TemplateId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrTemplateNotFound, "template ID %s not found", msg.TemplateId)
	}

	// Check gas limit
	maxGas := k.MaxContractGas(ctx)
	if msg.GasLimit > maxGas {
		return nil, sdkerrors.Wrapf(types.ErrGasLimitTooHigh, "gas limit %d exceeds maximum %d", msg.GasLimit, maxGas)
	}

	// Check metadata size
	maxMetadataSize := k.MaxMetadataSize(ctx)
	if uint64(len(msg.Metadata)) > maxMetadataSize {
		return nil, sdkerrors.Wrapf(types.ErrMetadataTooLarge, "metadata size %d exceeds maximum %d", len(msg.Metadata), maxMetadataSize)
	}

	// Check deposit
	minDeposit := k.MinContractDeposit(ctx)
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

	// Generate a unique ID for the contract
	id := uuid.New().String()

	// Calculate the code hash
	codeHash := k.CalculateCodeHash(template.Code)

	// Create the contract
	contract := types.DynaContract{
		Id:             id,
		Name:           msg.Name,
		Description:    msg.Description,
		Creator:        msg.Creator,
		Owner:          msg.Creator,
		ContractType:   template.ContractType,
		Status:         types.DynaContractStatusActive,
		CodeHash:       codeHash,
		Code:           template.Code,
		Abi:            template.Abi,
		AgentId:        msg.AgentId,
		State:          []byte("{}"),
		Metadata:       msg.Metadata,
		CreatedAt:      ctx.BlockTime(),
		UpdatedAt:      ctx.BlockTime(),
		Tags:           msg.Tags,
		GasLimit:       msg.GasLimit,
		ExecutionCount: 0,
	}

	// Store the contract
	k.SetDynaContract(ctx, contract)

	// Grant execute permission to the creator
	permission := types.DynaContractPermission{
		ContractId:     id,
		Address:        msg.Creator,
		PermissionType: types.PermissionTypeExecute,
		GrantedAt:      ctx.BlockTime(),
		ExpiresAt:      time.Time{}, // Never expires
		GrantedBy:      msg.Creator,
		Metadata:       []byte("{}"),
	}
	k.SetDynaContractPermission(ctx, permission)

	// Update the template usage count
	template.UsageCount++
	k.SetDynaContractTemplate(ctx, template)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeInstantiateFromTemplate,
			sdk.NewAttribute(types.AttributeKeyContractID, id),
			sdk.NewAttribute(types.AttributeKeyTemplateID, msg.TemplateId),
			sdk.NewAttribute(types.AttributeKeyCreator, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyName, msg.Name),
			sdk.NewAttribute(types.AttributeKeyDeposit, msg.Deposit.String()),
		),
	)

	return &types.MsgInstantiateFromTemplateResponse{
		Id: id,
	}, nil
}

// GrantContractPermission grants permission to a dynamic contract
func (k msgServer) GrantContractPermission(goCtx context.Context, msg *types.MsgGrantContractPermission) (*types.MsgGrantContractPermissionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// Get the contract
	contract, found := k.GetDynaContract(ctx, msg.ContractId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrContractNotFound, "contract ID %s not found", msg.ContractId)
	}

	// Parse the owner address
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}

	// Check if the caller is the owner
	if contract.Owner != owner.String() {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, "only the owner can grant permissions")
	}

	// Parse the address to grant permission to
	_, err = sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address (%s)", err)
	}

	// Check if the permission type is valid
	if msg.PermissionType != types.PermissionTypeExecute && msg.PermissionType != types.PermissionTypeAdmin {
		return nil, sdkerrors.Wrapf(types.ErrInvalidPermissionType, "invalid permission type: %s", msg.PermissionType)
	}

	// Calculate expiration time
	var expiresAt time.Time
	if msg.ExpirationDays > 0 {
		expiresAt = ctx.BlockTime().Add(time.Duration(msg.ExpirationDays) * 24 * time.Hour)
	}

	// Create the permission
	permission := types.DynaContractPermission{
		ContractId:     msg.ContractId,
		Address:        msg.Address,
		PermissionType: msg.PermissionType,
		GrantedAt:      ctx.BlockTime(),
		ExpiresAt:      expiresAt,
		GrantedBy:      msg.Owner,
		Metadata:       msg.Metadata,
	}

	// Store the permission
	k.SetDynaContractPermission(ctx, permission)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeGrantContractPermission,
			sdk.NewAttribute(types.AttributeKeyContractID, msg.ContractId),
			sdk.NewAttribute(types.AttributeKeyOwner, msg.Owner),
			sdk.NewAttribute(types.AttributeKeyAddress, msg.Address),
			sdk.NewAttribute(types.AttributeKeyPermissionType, msg.PermissionType),
			sdk.NewAttribute(types.AttributeKeyExpirationDays, fmt.Sprintf("%d", msg.ExpirationDays)),
		),
	)

	return &types.MsgGrantContractPermissionResponse{}, nil
}

// RevokeContractPermission revokes permission from a dynamic contract
func (k msgServer) RevokeContractPermission(goCtx context.Context, msg *types.MsgRevokeContractPermission) (*types.MsgRevokeContractPermissionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// Get the contract
	contract, found := k.GetDynaContract(ctx, msg.ContractId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrContractNotFound, "contract ID %s not found", msg.ContractId)
	}

	// Parse the owner address
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}

	// Check if the caller is the owner
	if contract.Owner != owner.String() {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, "only the owner can revoke permissions")
	}

	// Parse the address to revoke permission from
	_, err = sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address (%s)", err)
	}

	// Check if the permission type is valid
	if msg.PermissionType != types.PermissionTypeExecute && msg.PermissionType != types.PermissionTypeAdmin {
		return nil, sdkerrors.Wrapf(types.ErrInvalidPermissionType, "invalid permission type: %s", msg.PermissionType)
	}

	// Check if the permission exists
	_, found = k.GetDynaContractPermission(ctx, msg.ContractId, msg.Address, msg.PermissionType)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrPermissionNotFound, "permission not found for contract %s, address %s, type %s", msg.ContractId, msg.Address, msg.PermissionType)
	}

	// Delete the permission
	k.DeleteDynaContractPermission(ctx, msg.ContractId, msg.Address, msg.PermissionType)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRevokeContractPermission,
			sdk.NewAttribute(types.AttributeKeyContractID, msg.ContractId),
			sdk.NewAttribute(types.AttributeKeyOwner, msg.Owner),
			sdk.NewAttribute(types.AttributeKeyAddress, msg.Address),
			sdk.NewAttribute(types.AttributeKeyPermissionType, msg.PermissionType),
		),
	)

	return &types.MsgRevokeContractPermissionResponse{}, nil
}