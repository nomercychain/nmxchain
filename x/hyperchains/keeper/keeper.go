package keeper

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/nomercychain/nmxchain/x/hyperchains/types"
)

// Keeper of the hyperchains store
type Keeper struct {
	storeKey   sdk.StoreKey
	memKey     sdk.StoreKey
	cdc        codec.BinaryCodec
	paramstore paramtypes.Subspace

	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
}

// NewKeeper creates a new hyperchains Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
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
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// SetChain sets a chain
func (k Keeper) SetChain(ctx sdk.Context, chain types.Chain) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.ChainKey, []byte(chain.ID)...)
	value := k.cdc.MustMarshal(&chain)
	store.Set(key, value)
}

// GetChain returns a chain by ID
func (k Keeper) GetChain(ctx sdk.Context, id string) (types.Chain, bool) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.ChainKey, []byte(id)...)
	value := store.Get(key)
	if value == nil {
		return types.Chain{}, false
	}

	var chain types.Chain
	k.cdc.MustUnmarshal(value, &chain)
	return chain, true
}

// DeleteChain deletes a chain
func (k Keeper) DeleteChain(ctx sdk.Context, id string) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.ChainKey, []byte(id)...)
	store.Delete(key)
}

// GetAllChains returns all chains
func (k Keeper) GetAllChains(ctx sdk.Context) []types.Chain {
	var chains []types.Chain
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ChainKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var chain types.Chain
		k.cdc.MustUnmarshal(iterator.Value(), &chain)
		chains = append(chains, chain)
	}

	return chains
}

// SetChainTemplate sets a chain template
func (k Keeper) SetChainTemplate(ctx sdk.Context, template types.ChainTemplate) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.ChainTemplateKey, []byte(template.ID)...)
	value := k.cdc.MustMarshal(&template)
	store.Set(key, value)
}

// GetChainTemplate returns a chain template by ID
func (k Keeper) GetChainTemplate(ctx sdk.Context, id string) (types.ChainTemplate, bool) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.ChainTemplateKey, []byte(id)...)
	value := store.Get(key)
	if value == nil {
		return types.ChainTemplate{}, false
	}

	var template types.ChainTemplate
	k.cdc.MustUnmarshal(value, &template)
	return template, true
}

// GetAllChainTemplates returns all chain templates
func (k Keeper) GetAllChainTemplates(ctx sdk.Context) []types.ChainTemplate {
	var templates []types.ChainTemplate
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ChainTemplateKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var template types.ChainTemplate
		k.cdc.MustUnmarshal(iterator.Value(), &template)
		templates = append(templates, template)
	}

	return templates
}

// SetChainModule sets a chain module
func (k Keeper) SetChainModule(ctx sdk.Context, module types.ChainModule) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.ChainModuleKey, []byte(module.ID)...)
	value := k.cdc.MustMarshal(&module)
	store.Set(key, value)
}

// GetChainModule returns a chain module by ID
func (k Keeper) GetChainModule(ctx sdk.Context, id string) (types.ChainModule, bool) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.ChainModuleKey, []byte(id)...)
	value := store.Get(key)
	if value == nil {
		return types.ChainModule{}, false
	}

	var module types.ChainModule
	k.cdc.MustUnmarshal(value, &module)
	return module, true
}

// GetAllChainModules returns all chain modules
func (k Keeper) GetAllChainModules(ctx sdk.Context) []types.ChainModule {
	var modules []types.ChainModule
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ChainModuleKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var module types.ChainModule
		k.cdc.MustUnmarshal(iterator.Value(), &module)
		modules = append(modules, module)
	}

	return modules
}

// SetChainDeployment sets a chain deployment
func (k Keeper) SetChainDeployment(ctx sdk.Context, deployment types.ChainDeployment) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.ChainDeploymentKey, []byte(deployment.ID)...)
	value := k.cdc.MustMarshal(&deployment)
	store.Set(key, value)
}

// GetChainDeployment returns a chain deployment by ID
func (k Keeper) GetChainDeployment(ctx sdk.Context, id string) (types.ChainDeployment, bool) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.ChainDeploymentKey, []byte(id)...)
	value := store.Get(key)
	if value == nil {
		return types.ChainDeployment{}, false
	}

	var deployment types.ChainDeployment
	k.cdc.MustUnmarshal(value, &deployment)
	return deployment, true
}

// GetChainDeploymentsByChain returns all deployments for a chain
func (k Keeper) GetChainDeploymentsByChain(ctx sdk.Context, chainID string) []types.ChainDeployment {
	var deployments []types.ChainDeployment
	allDeployments := k.GetAllChainDeployments(ctx)

	for _, deployment := range allDeployments {
		if deployment.ChainID == chainID {
			deployments = append(deployments, deployment)
		}
	}

	return deployments
}

// GetAllChainDeployments returns all chain deployments
func (k Keeper) GetAllChainDeployments(ctx sdk.Context) []types.ChainDeployment {
	var deployments []types.ChainDeployment
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ChainDeploymentKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var deployment types.ChainDeployment
		k.cdc.MustUnmarshal(iterator.Value(), &deployment)
		deployments = append(deployments, deployment)
	}

	return deployments
}

// SetChainValidator sets a chain validator
func (k Keeper) SetChainValidator(ctx sdk.Context, validator types.ChainValidator) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.ChainValidatorKey, []byte(fmt.Sprintf("%s-%s", validator.ChainID, validator.ValidatorAddress))...)
	value := k.cdc.MustMarshal(&validator)
	store.Set(key, value)
}

// GetChainValidator returns a chain validator by chain ID and validator address
func (k Keeper) GetChainValidator(ctx sdk.Context, chainID string, validatorAddr sdk.ValAddress) (types.ChainValidator, bool) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.ChainValidatorKey, []byte(fmt.Sprintf("%s-%s", chainID, validatorAddr))...)
	value := store.Get(key)
	if value == nil {
		return types.ChainValidator{}, false
	}

	var validator types.ChainValidator
	k.cdc.MustUnmarshal(value, &validator)
	return validator, true
}

// GetChainValidatorsByChain returns all validators for a chain
func (k Keeper) GetChainValidatorsByChain(ctx sdk.Context, chainID string) []types.ChainValidator {
	var validators []types.ChainValidator
	store := ctx.KVStore(k.storeKey)
	prefix := append(types.ChainValidatorKey, []byte(fmt.Sprintf("%s-", chainID))...)
	iterator := sdk.KVStorePrefixIterator(store, prefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var validator types.ChainValidator
		k.cdc.MustUnmarshal(iterator.Value(), &validator)
		validators = append(validators, validator)
	}

	return validators
}

// SetChainProposal sets a chain proposal
func (k Keeper) SetChainProposal(ctx sdk.Context, proposal types.ChainProposal) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.ChainProposalKey, sdk.Uint64ToBigEndian(proposal.ID)...)
	value := k.cdc.MustMarshal(&proposal)
	store.Set(key, value)
}

// GetChainProposal returns a chain proposal by ID
func (k Keeper) GetChainProposal(ctx sdk.Context, id uint64) (types.ChainProposal, bool) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.ChainProposalKey, sdk.Uint64ToBigEndian(id)...)
	value := store.Get(key)
	if value == nil {
		return types.ChainProposal{}, false
	}

	var proposal types.ChainProposal
	k.cdc.MustUnmarshal(value, &proposal)
	return proposal, true
}

// GetChainProposalsByChain returns all proposals for a chain
func (k Keeper) GetChainProposalsByChain(ctx sdk.Context, chainID string) []types.ChainProposal {
	var proposals []types.ChainProposal
	allProposals := k.GetAllChainProposals(ctx)

	for _, proposal := range allProposals {
		if proposal.ChainID == chainID {
			proposals = append(proposals, proposal)
		}
	}

	return proposals
}

// GetAllChainProposals returns all chain proposals
func (k Keeper) GetAllChainProposals(ctx sdk.Context) []types.ChainProposal {
	var proposals []types.ChainProposal
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ChainProposalKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var proposal types.ChainProposal
		k.cdc.MustUnmarshal(iterator.Value(), &proposal)
		proposals = append(proposals, proposal)
	}

	return proposals
}

// SetChainMetrics sets chain metrics
func (k Keeper) SetChainMetrics(ctx sdk.Context, metrics types.ChainMetrics) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.ChainMetricsKey, []byte(metrics.ChainID)...)
	value := k.cdc.MustMarshal(&metrics)
	store.Set(key, value)
}

// GetChainMetrics returns chain metrics by chain ID
func (k Keeper) GetChainMetrics(ctx sdk.Context, chainID string) (types.ChainMetrics, bool) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.ChainMetricsKey, []byte(chainID)...)
	value := store.Get(key)
	if value == nil {
		return types.ChainMetrics{}, false
	}

	var metrics types.ChainMetrics
	k.cdc.MustUnmarshal(value, &metrics)
	return metrics, true
}

// GetAllChainMetrics returns all chain metrics
func (k Keeper) GetAllChainMetrics(ctx sdk.Context) []types.ChainMetrics {
	var metrics []types.ChainMetrics
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ChainMetricsKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var m types.ChainMetrics
		k.cdc.MustUnmarshal(iterator.Value(), &m)
		metrics = append(metrics, m)
	}

	return metrics
}

// CreateChainFromPrompt creates a new chain from a natural language prompt
func (k Keeper) CreateChainFromPrompt(ctx sdk.Context, creator sdk.AccAddress, prompt string) (string, error) {
	// In a real implementation, this would:
	// 1. Process the prompt using an AI model
	// 2. Determine the appropriate chain type, modules, and configuration
	// 3. Generate the chain code and configuration
	// 4. Create and deploy the chain

	// For now, we'll create a simple chain with default settings
	chainType := types.ChainTypeRollup
	if prompt == "Create a low-latency NFT game chain" {
		chainType = types.ChainTypeAppChain
	} else if prompt == "Create a privacy-focused DeFi chain" {
		chainType = types.ChainTypeZkEVM
	}

	// Generate a unique ID for the chain
	id := fmt.Sprintf("chain-%s-%d", creator.String(), ctx.BlockHeight())

	// Create the chain
	chain := types.Chain{
		ID:          id,
		Name:        fmt.Sprintf("Chain from prompt: %s", prompt),
		Description: prompt,
		Creator:     creator,
		ChainType:   chainType,
		Status:      types.ChainStatusProposed,
		Version:     1,
		CreatedAt:   ctx.BlockTime(),
		UpdatedAt:   ctx.BlockTime(),
		Modules:     []string{"core", "governance", "token"},
		Config:      json.RawMessage(`{"blockTime": 2, "maxValidators": 10}`),
		Metadata:    json.RawMessage(`{"prompt": "` + prompt + `"}`),
	}

	// Store the chain
	k.SetChain(ctx, chain)

	// Create a deployment
	deployment := types.ChainDeployment{
		ID:        fmt.Sprintf("deploy-%s-1", id),
		ChainID:   id,
		Version:   1,
		Deployer:  creator,
		Status:    "pending",
		StartedAt: ctx.BlockTime(),
		Config:    json.RawMessage(`{}`),
	}
	k.SetChainDeployment(ctx, deployment)

	return id, nil
}

// CreateChainTemplate creates a new chain template
func (k Keeper) CreateChainTemplate(ctx sdk.Context, creator sdk.AccAddress, name string, description string, chainType types.ChainType, modules []string, config json.RawMessage, metadata json.RawMessage) (string, error) {
	// Generate a unique ID for the template
	id := fmt.Sprintf("template-%s-%d", creator.String(), ctx.BlockHeight())

	// Create the template
	template := types.ChainTemplate{
		ID:          id,
		Name:        name,
		Description: description,
		ChainType:   chainType,
		Creator:     creator,
		CreatedAt:   ctx.BlockTime(),
		UpdatedAt:   ctx.BlockTime(),
		Modules:     modules,
		Config:      config,
		Metadata:    metadata,
	}

	// Store the template
	k.SetChainTemplate(ctx, template)

	return id, nil
}

// CreateChainFromTemplate creates a new chain from a template
func (k Keeper) CreateChainFromTemplate(ctx sdk.Context, creator sdk.AccAddress, templateID string, name string, description string, config json.RawMessage) (string, error) {
	// Get the template
	template, found := k.GetChainTemplate(ctx, templateID)
	if !found {
		return "", fmt.Errorf("template not found: %s", templateID)
	}

	// Generate a unique ID for the chain
	id := fmt.Sprintf("chain-%s-%d", creator.String(), ctx.BlockHeight())

	// Create the chain
	chain := types.Chain{
		ID:          id,
		Name:        name,
		Description: description,
		Creator:     creator,
		ChainType:   template.ChainType,
		Status:      types.ChainStatusProposed,
		Version:     1,
		CreatedAt:   ctx.BlockTime(),
		UpdatedAt:   ctx.BlockTime(),
		Modules:     template.Modules,
		Config:      config,
		Metadata:    json.RawMessage(`{"templateId": "` + templateID + `"}`),
	}

	// Store the chain
	k.SetChain(ctx, chain)

	// Create a deployment
	deployment := types.ChainDeployment{
		ID:        fmt.Sprintf("deploy-%s-1", id),
		ChainID:   id,
		Version:   1,
		Deployer:  creator,
		Status:    "pending",
		StartedAt: ctx.BlockTime(),
		Config:    json.RawMessage(`{}`),
	}
	k.SetChainDeployment(ctx, deployment)

	return id, nil
}

// DeployChain deploys a chain
func (k Keeper) DeployChain(ctx sdk.Context, chainID string, deployer sdk.AccAddress) error {
	// Get the chain
	chain, found := k.GetChain(ctx, chainID)
	if !found {
		return fmt.Errorf("chain not found: %s", chainID)
	}

	// Check if the chain is in a deployable state
	if chain.Status != types.ChainStatusProposed && chain.Status != types.ChainStatusApproved {
		return fmt.Errorf("chain is not in a deployable state")
	}

	// Check if the deployer is the creator or has permission
	if !chain.Creator.Equals(deployer) {
		return fmt.Errorf("only the creator can deploy the chain")
	}

	// Update the chain status
	chain.Status = types.ChainStatusDeploying
	chain.UpdatedAt = ctx.BlockTime()
	k.SetChain(ctx, chain)

	// Create a new deployment
	deployment := types.ChainDeployment{
		ID:        fmt.Sprintf("deploy-%s-%d", chainID, chain.Version),
		ChainID:   chainID,
		Version:   chain.Version,
		Deployer:  deployer,
		Status:    "in_progress",
		StartedAt: ctx.BlockTime(),
		Config:    chain.Config,
	}
	k.SetChainDeployment(ctx, deployment)

	// In a real implementation, this would initiate the actual deployment process
	// For now, we'll just update the deployment status in the next block

	return nil
}

// UpdateChainDeployment updates a chain deployment
func (k Keeper) UpdateChainDeployment(ctx sdk.Context, deploymentID string, status string, logs string, endpoints json.RawMessage) error {
	// Get the deployment
	deployment, found := k.GetChainDeployment(ctx, deploymentID)
	if !found {
		return fmt.Errorf("deployment not found: %s", deploymentID)
	}

	// Update the deployment
	deployment.Status = status
	if status == "completed" || status == "failed" {
		deployment.CompletedAt = ctx.BlockTime()
	}
	if logs != "" {
		deployment.Logs = logs
	}
	if endpoints != nil {
		deployment.Endpoints = endpoints
	}
	k.SetChainDeployment(ctx, deployment)

	// If the deployment is completed, update the chain status
	if status == "completed" {
		chain, found := k.GetChain(ctx, deployment.ChainID)
		if found {
			chain.Status = types.ChainStatusActive
			chain.UpdatedAt = ctx.BlockTime()
			k.SetChain(ctx, chain)
		}
	}

	return nil
}

// RegisterChainValidator registers a validator for a chain
func (k Keeper) RegisterChainValidator(ctx sdk.Context, chainID string, validatorAddr sdk.ValAddress, operatorAddr sdk.AccAddress) error {
	// Get the chain
	chain, found := k.GetChain(ctx, chainID)
	if !found {
		return fmt.Errorf("chain not found: %s", chainID)
	}

	// Check if the chain is active
	if chain.Status != types.ChainStatusActive {
		return fmt.Errorf("chain is not active")
	}

	// Check if the validator is already registered
	_, found = k.GetChainValidator(ctx, chainID, validatorAddr)
	if found {
		return fmt.Errorf("validator already registered for this chain")
	}

	// Create the validator
	validator := types.ChainValidator{
		ChainID:          chainID,
		ValidatorAddress: validatorAddr,
		OperatorAddress:  operatorAddr,
		Status:           "active",
		Power:            1,
		JoinedAt:         ctx.BlockTime(),
		LastActive:       ctx.BlockTime(),
	}

	// Store the validator
	k.SetChainValidator(ctx, validator)

	return nil
}

// CreateChainProposal creates a new chain proposal
func (k Keeper) CreateChainProposal(ctx sdk.Context, chainID string, proposer sdk.AccAddress, title string, description string, proposalType string, content json.RawMessage) (uint64, error) {
	// Get the chain
	chain, found := k.GetChain(ctx, chainID)
	if !found {
		return 0, fmt.Errorf("chain not found: %s", chainID)
	}

	// Check if the chain is active
	if chain.Status != types.ChainStatusActive {
		return 0, fmt.Errorf("chain is not active")
	}

	// Generate a unique ID for the proposal
	id := uint64(ctx.BlockHeight())

	// Create the proposal
	proposal := types.ChainProposal{
		ID:           id,
		ChainID:      chainID,
		Title:        title,
		Description:  description,
		Proposer:     proposer,
		ProposalType: proposalType,
		Status:       "voting",
		VotesYes:     sdk.ZeroInt(),
		VotesNo:      sdk.ZeroInt(),
		CreatedAt:    ctx.BlockTime(),
		EndTime:      ctx.BlockTime().Add(time.Hour * 24 * 7), // 1 week voting period
		Content:      content,
	}

	// Store the proposal
	k.SetChainProposal(ctx, proposal)

	return id, nil
}

// VoteOnChainProposal votes on a chain proposal
func (k Keeper) VoteOnChainProposal(ctx sdk.Context, proposalID uint64, voter sdk.AccAddress, voteYes bool, votePower sdk.Int) error {
	// Get the proposal
	proposal, found := k.GetChainProposal(ctx, proposalID)
	if !found {
		return fmt.Errorf("proposal not found: %d", proposalID)
	}

	// Check if the proposal is still in voting period
	if proposal.Status != "voting" {
		return fmt.Errorf("proposal is not in voting period")
	}
	if ctx.BlockTime().After(proposal.EndTime) {
		return fmt.Errorf("voting period has ended")
	}

	// Update the vote count
	if voteYes {
		proposal.VotesYes = proposal.VotesYes.Add(votePower)
	} else {
		proposal.VotesNo = proposal.VotesNo.Add(votePower)
	}

	// Store the updated proposal
	k.SetChainProposal(ctx, proposal)

	return nil
}

// ProcessChainProposals processes chain proposals
func (k Keeper) ProcessChainProposals(ctx sdk.Context) {
	// Get all proposals
	proposals := k.GetAllChainProposals(ctx)
	currentTime := ctx.BlockTime()

	for _, proposal := range proposals {
		// Skip proposals that are not in voting period or haven't ended
		if proposal.Status != "voting" || currentTime.Before(proposal.EndTime) {
			continue
		}

		// Check if the proposal passed
		if proposal.VotesYes.GT(proposal.VotesNo) {
			// Update the proposal status
			proposal.Status = "passed"

			// Execute the proposal
			switch proposal.ProposalType {
			case "upgrade":
				// Handle chain upgrade
				chain, found := k.GetChain(ctx, proposal.ChainID)
				if found {
					chain.Version++
					chain.Status = types.ChainStatusUpgrading
					chain.UpdatedAt = currentTime
					k.SetChain(ctx, chain)
				}
			case "parameter":
				// Handle parameter change
				// In a real implementation, this would update specific parameters
			case "module":
				// Handle module addition/removal
				// In a real implementation, this would update the chain's modules
			}
		} else {
			// Update the proposal status
			proposal.Status = "rejected"
		}

		// Store the updated proposal
		k.SetChainProposal(ctx, proposal)
	}
}

// UpdateChainMetrics updates the metrics for a chain
func (k Keeper) UpdateChainMetrics(ctx sdk.Context, chainID string, blockHeight int64, totalTx uint64, activeUsers uint64, tps sdk.Dec, avgFee sdk.Dec, tvl sdk.Coins) error {
	// Create or update the metrics
	metrics := types.ChainMetrics{
		ChainID:           chainID,
		BlockHeight:       blockHeight,
		TotalTransactions: totalTx,
		ActiveUsers:       activeUsers,
		TPS:               tps,
		AverageFee:        avgFee,
		TotalValueLocked:  tvl,
		UpdatedAt:         ctx.BlockTime(),
	}

	// Store the metrics
	k.SetChainMetrics(ctx, metrics)

	return nil
}