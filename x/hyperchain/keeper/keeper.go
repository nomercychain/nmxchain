package keeper

import (
	"fmt"
	"crypto/sha256"
	"encoding/hex"

	"github.com/tendermint/tendermint/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/nomercychain/nmxchain/x/hyperchain/types"
	"github.com/nomercychain/nmxchain/x/deai/types" // For AI agent integration
)

// Keeper of the hyperchain store
type Keeper struct {
	cdc        codec.BinaryCodec
	storeKey   sdk.StoreKey
	memKey     sdk.StoreKey
	paramstore paramtypes.Subspace
	bankKeeper types.BankKeeper
	deaiKeeper types.DeAIKeeper
}

// NewKeeper creates a new hyperchain Keeper instance
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

// GetHyperchain returns a hyperchain by ID
func (k Keeper) GetHyperchain(ctx sdk.Context, id string) (hyperchain types.Hyperchain, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.HyperchainKey(id)
	
	value := store.Get(key)
	if value == nil {
		return hyperchain, false
	}
	
	k.cdc.MustUnmarshal(value, &hyperchain)
	return hyperchain, true
}

// SetHyperchain sets a hyperchain in the store
func (k Keeper) SetHyperchain(ctx sdk.Context, hyperchain types.Hyperchain) {
	store := ctx.KVStore(k.storeKey)
	key := types.HyperchainKey(hyperchain.Id)
	value := k.cdc.MustMarshal(&hyperchain)
	
	store.Set(key, value)
}

// DeleteHyperchain deletes a hyperchain from the store
func (k Keeper) DeleteHyperchain(ctx sdk.Context, id string) {
	store := ctx.KVStore(k.storeKey)
	key := types.HyperchainKey(id)
	store.Delete(key)
}

// GetAllHyperchains returns all hyperchains
func (k Keeper) GetAllHyperchains(ctx sdk.Context) (hyperchains []types.Hyperchain) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.HyperchainKeyPrefix)
	defer iterator.Close()
	
	for ; iterator.Valid(); iterator.Next() {
		var hyperchain types.Hyperchain
		k.cdc.MustUnmarshal(iterator.Value(), &hyperchain)
		hyperchains = append(hyperchains, hyperchain)
	}
	
	return hyperchains
}

// GetHyperchainsByCreator returns all hyperchains created by a specific address
func (k Keeper) GetHyperchainsByCreator(ctx sdk.Context, creator string) (hyperchains []types.Hyperchain) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.HyperchainKeyPrefix)
	defer iterator.Close()
	
	for ; iterator.Valid(); iterator.Next() {
		var hyperchain types.Hyperchain
		k.cdc.MustUnmarshal(iterator.Value(), &hyperchain)
		
		if hyperchain.Creator == creator {
			hyperchains = append(hyperchains, hyperchain)
		}
	}
	
	return hyperchains
}

// GetHyperchainsByParent returns all hyperchains with a specific parent chain
func (k Keeper) GetHyperchainsByParent(ctx sdk.Context, parentChainID string) (hyperchains []types.Hyperchain) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.HyperchainKeyPrefix)
	defer iterator.Close()
	
	for ; iterator.Valid(); iterator.Next() {
		var hyperchain types.Hyperchain
		k.cdc.MustUnmarshal(iterator.Value(), &hyperchain)
		
		if hyperchain.ParentChainId == parentChainID {
			hyperchains = append(hyperchains, hyperchain)
		}
	}
	
	return hyperchains
}

// GetHyperchainValidator returns a validator for a specific hyperchain
func (k Keeper) GetHyperchainValidator(ctx sdk.Context, chainID string, address string) (validator types.HyperchainValidator, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.HyperchainValidatorKey(chainID, address)
	
	value := store.Get(key)
	if value == nil {
		return validator, false
	}
	
	k.cdc.MustUnmarshal(value, &validator)
	return validator, true
}

// SetHyperchainValidator sets a validator for a hyperchain in the store
func (k Keeper) SetHyperchainValidator(ctx sdk.Context, validator types.HyperchainValidator) {
	store := ctx.KVStore(k.storeKey)
	key := types.HyperchainValidatorKey(validator.ChainId, validator.Address)
	value := k.cdc.MustMarshal(&validator)
	
	store.Set(key, value)
}

// DeleteHyperchainValidator deletes a validator for a hyperchain from the store
func (k Keeper) DeleteHyperchainValidator(ctx sdk.Context, chainID string, address string) {
	store := ctx.KVStore(k.storeKey)
	key := types.HyperchainValidatorKey(chainID, address)
	store.Delete(key)
}

// GetHyperchainValidatorsByChain returns all validators for a specific hyperchain
func (k Keeper) GetHyperchainValidatorsByChain(ctx sdk.Context, chainID string) (validators []types.HyperchainValidator) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.HyperchainValidatorKeyPrefixByChain(chainID))
	defer iterator.Close()
	
	for ; iterator.Valid(); iterator.Next() {
		var validator types.HyperchainValidator
		k.cdc.MustUnmarshal(iterator.Value(), &validator)
		validators = append(validators, validator)
	}
	
	return validators
}

// GetHyperchainBlock returns a block for a specific hyperchain
func (k Keeper) GetHyperchainBlock(ctx sdk.Context, chainID string, height uint64) (block types.HyperchainBlock, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.HyperchainBlockKey(chainID, height)
	
	value := store.Get(key)
	if value == nil {
		return block, false
	}
	
	k.cdc.MustUnmarshal(value, &block)
	return block, true
}

// SetHyperchainBlock sets a block for a hyperchain in the store
func (k Keeper) SetHyperchainBlock(ctx sdk.Context, block types.HyperchainBlock) {
	store := ctx.KVStore(k.storeKey)
	key := types.HyperchainBlockKey(block.ChainId, block.Height)
	value := k.cdc.MustMarshal(&block)
	
	store.Set(key, value)
}

// GetHyperchainBlocksByChain returns all blocks for a specific hyperchain
func (k Keeper) GetHyperchainBlocksByChain(ctx sdk.Context, chainID string) (blocks []types.HyperchainBlock) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.HyperchainBlockKeyPrefixByChain(chainID))
	defer iterator.Close()
	
	for ; iterator.Valid(); iterator.Next() {
		var block types.HyperchainBlock
		k.cdc.MustUnmarshal(iterator.Value(), &block)
		blocks = append(blocks, block)
	}
	
	return blocks
}

// GetHyperchainTransaction returns a transaction by ID
func (k Keeper) GetHyperchainTransaction(ctx sdk.Context, id string) (tx types.HyperchainTransaction, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.HyperchainTransactionKey(id)
	
	value := store.Get(key)
	if value == nil {
		return tx, false
	}
	
	k.cdc.MustUnmarshal(value, &tx)
	return tx, true
}

// SetHyperchainTransaction sets a transaction in the store
func (k Keeper) SetHyperchainTransaction(ctx sdk.Context, tx types.HyperchainTransaction) {
	store := ctx.KVStore(k.storeKey)
	key := types.HyperchainTransactionKey(tx.Id)
	value := k.cdc.MustMarshal(&tx)
	
	store.Set(key, value)
}

// GetHyperchainTransactionsByChain returns all transactions for a specific hyperchain
func (k Keeper) GetHyperchainTransactionsByChain(ctx sdk.Context, chainID string) (txs []types.HyperchainTransaction) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.HyperchainTransactionKeyPrefix)
	defer iterator.Close()
	
	for ; iterator.Valid(); iterator.Next() {
		var tx types.HyperchainTransaction
		k.cdc.MustUnmarshal(iterator.Value(), &tx)
		
		if tx.ChainId == chainID {
			txs = append(txs, tx)
		}
	}
	
	return txs
}

// GetHyperchainBridge returns a bridge by ID
func (k Keeper) GetHyperchainBridge(ctx sdk.Context, id string) (bridge types.HyperchainBridge, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.HyperchainBridgeKey(id)
	
	value := store.Get(key)
	if value == nil {
		return bridge, false
	}
	
	k.cdc.MustUnmarshal(value, &bridge)
	return bridge, true
}

// SetHyperchainBridge sets a bridge in the store
func (k Keeper) SetHyperchainBridge(ctx sdk.Context, bridge types.HyperchainBridge) {
	store := ctx.KVStore(k.storeKey)
	key := types.HyperchainBridgeKey(bridge.Id)
	value := k.cdc.MustMarshal(&bridge)
	
	store.Set(key, value)
}

// GetAllHyperchainBridges returns all hyperchain bridges
func (k Keeper) GetAllHyperchainBridges(ctx sdk.Context) (bridges []types.HyperchainBridge) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.HyperchainBridgeKeyPrefix)
	defer iterator.Close()
	
	for ; iterator.Valid(); iterator.Next() {
		var bridge types.HyperchainBridge
		k.cdc.MustUnmarshal(iterator.Value(), &bridge)
		bridges = append(bridges, bridge)
	}
	
	return bridges
}

// GetHyperchainBridgesByChain returns all bridges for a specific hyperchain
func (k Keeper) GetHyperchainBridgesByChain(ctx sdk.Context, chainID string) (bridges []types.HyperchainBridge) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.HyperchainBridgeKeyPrefix)
	defer iterator.Close()
	
	for ; iterator.Valid(); iterator.Next() {
		var bridge types.HyperchainBridge
		k.cdc.MustUnmarshal(iterator.Value(), &bridge)
		
		if bridge.SourceChainId == chainID || bridge.TargetChainId == chainID {
			bridges = append(bridges, bridge)
		}
	}
	
	return bridges
}

// GetHyperchainBridgeTransaction returns a bridge transaction by ID
func (k Keeper) GetHyperchainBridgeTransaction(ctx sdk.Context, id string) (tx types.HyperchainBridgeTransaction, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.HyperchainBridgeTransactionKey(id)
	
	value := store.Get(key)
	if value == nil {
		return tx, false
	}
	
	k.cdc.MustUnmarshal(value, &tx)
	return tx, true
}

// SetHyperchainBridgeTransaction sets a bridge transaction in the store
func (k Keeper) SetHyperchainBridgeTransaction(ctx sdk.Context, tx types.HyperchainBridgeTransaction) {
	store := ctx.KVStore(k.storeKey)
	key := types.HyperchainBridgeTransactionKey(tx.Id)
	value := k.cdc.MustMarshal(&tx)
	
	store.Set(key, value)
}

// GetHyperchainBridgeTransactionsByBridge returns all transactions for a specific bridge
func (k Keeper) GetHyperchainBridgeTransactionsByBridge(ctx sdk.Context, bridgeID string) (txs []types.HyperchainBridgeTransaction) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.HyperchainBridgeTransactionKeyPrefix)
	defer iterator.Close()
	
	for ; iterator.Valid(); iterator.Next() {
		var tx types.HyperchainBridgeTransaction
		k.cdc.MustUnmarshal(iterator.Value(), &tx)
		
		if tx.BridgeId == bridgeID {
			txs = append(txs, tx)
		}
	}
	
	return txs
}

// GetHyperchainPermission returns a permission for a specific hyperchain and address
func (k Keeper) GetHyperchainPermission(ctx sdk.Context, chainID string, address string, permissionType string) (permission types.HyperchainPermission, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.HyperchainPermissionKey(chainID, address, permissionType)
	
	value := store.Get(key)
	if value == nil {
		return permission, false
	}
	
	k.cdc.MustUnmarshal(value, &permission)
	return permission, true
}

// SetHyperchainPermission sets a permission for a hyperchain in the store
func (k Keeper) SetHyperchainPermission(ctx sdk.Context, permission types.HyperchainPermission) {
	store := ctx.KVStore(k.storeKey)
	key := types.HyperchainPermissionKey(permission.ChainId, permission.Address, permission.PermissionType)
	value := k.cdc.MustMarshal(&permission)
	
	store.Set(key, value)
}

// DeleteHyperchainPermission deletes a permission for a hyperchain from the store
func (k Keeper) DeleteHyperchainPermission(ctx sdk.Context, chainID string, address string, permissionType string) {
	store := ctx.KVStore(k.storeKey)
	key := types.HyperchainPermissionKey(chainID, address, permissionType)
	store.Delete(key)
}

// GetHyperchainPermissionsByChain returns all permissions for a specific hyperchain
func (k Keeper) GetHyperchainPermissionsByChain(ctx sdk.Context, chainID string) (permissions []types.HyperchainPermission) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.HyperchainPermissionKeyPrefixByChain(chainID))
	defer iterator.Close()
	
	for ; iterator.Valid(); iterator.Next() {
		var permission types.HyperchainPermission
		k.cdc.MustUnmarshal(iterator.Value(), &permission)
		permissions = append(permissions, permission)
	}
	
	return permissions
}

// HasPermission checks if an address has a specific permission for a hyperchain
func (k Keeper) HasPermission(ctx sdk.Context, chainID string, address string, permissionType string) bool {
	permission, found := k.GetHyperchainPermission(ctx, chainID, address, permissionType)
	if !found {
		return false
	}
	
	// Check if the permission has expired
	if !permission.ExpiresAt.IsZero() && permission.ExpiresAt.Before(ctx.BlockTime()) {
		return false
	}
	
	return true
}

// CalculateBlockHash calculates the hash of a block
func (k Keeper) CalculateBlockHash(block types.HyperchainBlock) string {
	// Create a hash of the block data
	blockData := fmt.Sprintf("%s-%d-%s-%d-%s", block.ChainId, block.Height, block.ParentHash, block.NumTxs, block.Proposer)
	hash := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hash[:])
}

// GetParams gets the hyperchain module parameters
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramstore.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the hyperchain module parameters
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// MaxHyperchainsPerAccount returns the maximum number of hyperchains per account parameter
func (k Keeper) MaxHyperchainsPerAccount(ctx sdk.Context) uint64 {
	var max uint64
	k.paramstore.Get(ctx, types.KeyMaxHyperchainsPerAccount, &max)
	return max
}

// MaxValidatorsPerHyperchain returns the maximum number of validators per hyperchain parameter
func (k Keeper) MaxValidatorsPerHyperchain(ctx sdk.Context) uint64 {
	var max uint64
	k.paramstore.Get(ctx, types.KeyMaxValidatorsPerHyperchain, &max)
	return max
}

// MaxBridgesPerHyperchain returns the maximum number of bridges per hyperchain parameter
func (k Keeper) MaxBridgesPerHyperchain(ctx sdk.Context) uint64 {
	var max uint64
	k.paramstore.Get(ctx, types.KeyMaxBridgesPerHyperchain, &max)
	return max
}

// MinHyperchainCreationDeposit returns the minimum hyperchain creation deposit parameter
func (k Keeper) MinHyperchainCreationDeposit(ctx sdk.Context) sdk.Coin {
	var deposit sdk.Coin
	k.paramstore.Get(ctx, types.KeyMinHyperchainCreationDeposit, &deposit)
	return deposit
}

// MinValidatorStake returns the minimum validator stake parameter
func (k Keeper) MinValidatorStake(ctx sdk.Context) sdk.Coin {
	var stake sdk.Coin
	k.paramstore.Get(ctx, types.KeyMinValidatorStake, &stake)
	return stake
}

// BridgeFeeRate returns the bridge fee rate parameter
func (k Keeper) BridgeFeeRate(ctx sdk.Context) sdk.Dec {
	var rate sdk.Dec
	k.paramstore.Get(ctx, types.KeyBridgeFeeRate, &rate)
	return rate
}