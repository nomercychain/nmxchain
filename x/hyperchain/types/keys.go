package types

const (
	// ModuleName defines the module name
	ModuleName = "hyperchain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_hyperchain"
)

var (
	// HyperchainKeyPrefix is the prefix for hyperchain keys
	HyperchainKeyPrefix = []byte{0x01}

	// HyperchainValidatorKeyPrefix is the prefix for hyperchain validator keys
	HyperchainValidatorKeyPrefix = []byte{0x02}

	// HyperchainBlockKeyPrefix is the prefix for hyperchain block keys
	HyperchainBlockKeyPrefix = []byte{0x03}

	// HyperchainTransactionKeyPrefix is the prefix for hyperchain transaction keys
	HyperchainTransactionKeyPrefix = []byte{0x04}

	// HyperchainBridgeKeyPrefix is the prefix for hyperchain bridge keys
	HyperchainBridgeKeyPrefix = []byte{0x05}

	// HyperchainBridgeTransactionKeyPrefix is the prefix for hyperchain bridge transaction keys
	HyperchainBridgeTransactionKeyPrefix = []byte{0x06}

	// HyperchainPermissionKeyPrefix is the prefix for hyperchain permission keys
	HyperchainPermissionKeyPrefix = []byte{0x07}
)

// Parameter store keys
var (
	KeyMaxHyperchainsPerAccount     = []byte("MaxHyperchainsPerAccount")
	KeyMaxValidatorsPerHyperchain   = []byte("MaxValidatorsPerHyperchain")
	KeyMaxBridgesPerHyperchain      = []byte("MaxBridgesPerHyperchain")
	KeyMinHyperchainCreationDeposit = []byte("MinHyperchainCreationDeposit")
	KeyMinValidatorStake            = []byte("MinValidatorStake")
	KeyBridgeFeeRate                = []byte("BridgeFeeRate")
)

// Event types
const (
	EventTypeCreateHyperchain                  = "create_hyperchain"
	EventTypeUpdateHyperchain                  = "update_hyperchain"
	EventTypeJoinHyperchainAsValidator         = "join_hyperchain_as_validator"
	EventTypeLeaveHyperchain                   = "leave_hyperchain"
	EventTypeCreateHyperchainBridge            = "create_hyperchain_bridge"
	EventTypeUpdateHyperchainBridge            = "update_hyperchain_bridge"
	EventTypeRegisterHyperchainBridgeRelayer   = "register_hyperchain_bridge_relayer"
	EventTypeRemoveHyperchainBridgeRelayer     = "remove_hyperchain_bridge_relayer"
	EventTypeInitiateHyperchainBridgeTransaction = "initiate_hyperchain_bridge_transaction"
	EventTypeApproveHyperchainBridgeTransaction = "approve_hyperchain_bridge_transaction"
	EventTypeCompleteHyperchainBridgeTransaction = "complete_hyperchain_bridge_transaction"
	EventTypeSubmitHyperchainBlock             = "submit_hyperchain_block"
	EventTypeSubmitHyperchainTransaction       = "submit_hyperchain_transaction"
	EventTypeGrantHyperchainPermission         = "grant_hyperchain_permission"
	EventTypeRevokeHyperchainPermission        = "revoke_hyperchain_permission"
)

// Event attribute keys
const (
	AttributeKeyHyperchainID        = "hyperchain_id"
	AttributeKeyCreator             = "creator"
	AttributeKeyAdmin               = "admin"
	AttributeKeyName                = "name"
	AttributeKeyChainType           = "chain_type"
	AttributeKeyConsensusType       = "consensus_type"
	AttributeKeyDeposit             = "deposit"
	AttributeKeyValidator           = "validator"
	AttributeKeyStake               = "stake"
	AttributeKeyBridgeID            = "bridge_id"
	AttributeKeySourceChainID       = "source_chain_id"
	AttributeKeyTargetChainID       = "target_chain_id"
	AttributeKeyRelayer             = "relayer"
	AttributeKeyTransactionID       = "transaction_id"
	AttributeKeyTargetTransactionID = "target_transaction_id"
	AttributeKeySender              = "sender"
	AttributeKeyRecipient           = "recipient"
	AttributeKeyAmount              = "amount"
	AttributeKeyBlockHeight         = "block_height"
	AttributeKeyBlockHash           = "block_hash"
	AttributeKeyProposer            = "proposer"
	AttributeKeyAddress             = "address"
	AttributeKeyPermissionType      = "permission_type"
	AttributeKeyExpirationDays      = "expiration_days"
)