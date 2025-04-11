package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "truthgpt"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_truthgpt"
)

// Store prefixes
var (
	// OracleProviderKeyPrefix is the prefix for oracle provider keys
	OracleProviderKeyPrefix = []byte{0x01}

	// OracleRequestKeyPrefix is the prefix for oracle request keys
	OracleRequestKeyPrefix = []byte{0x02}

	// OracleResponseKeyPrefix is the prefix for oracle response keys
	OracleResponseKeyPrefix = []byte{0x03}

	// ProviderResponseHistoryKeyPrefix is the prefix for provider response history keys
	ProviderResponseHistoryKeyPrefix = []byte{0x04}

	// DataSourceKeyPrefix is the prefix for data source keys
	DataSourceKeyPrefix = []byte{0x05}

	// OracleScriptKeyPrefix is the prefix for oracle script keys
	OracleScriptKeyPrefix = []byte{0x06}

	// ProviderRewardsKeyPrefix is the prefix for provider rewards keys
	ProviderRewardsKeyPrefix = []byte{0x07}

	// ProviderStatsKeyPrefix is the prefix for provider stats keys
	ProviderStatsKeyPrefix = []byte{0x08}

	// ResultKeyPrefix is the prefix for result keys
	ResultKeyPrefix = []byte{0x09}

	// LastOracleRequestIDKey is the key for the last oracle request ID
	LastOracleRequestIDKey = []byte{0x10}

	// LastOracleResponseIDKey is the key for the last oracle response ID
	LastOracleResponseIDKey = []byte{0x11}

	// LastDataSourceIDKey is the key for the last data source ID
	LastDataSourceIDKey = []byte{0x12}

	// LastOracleScriptIDKey is the key for the last oracle script ID
	LastOracleScriptIDKey = []byte{0x13}
)

// Parameter store keys
var (
	// KeyMinProviderStake is the key for the minimum provider stake parameter
	KeyMinProviderStake = []byte("MinProviderStake")

	// KeyMaxProviderCount is the key for the maximum provider count parameter
	KeyMaxProviderCount = []byte("MaxProviderCount")

	// KeyMinRequestFee is the key for the minimum request fee parameter
	KeyMinRequestFee = []byte("MinRequestFee")

	// KeyDefaultTimeout is the key for the default timeout parameter
	KeyDefaultTimeout = []byte("DefaultTimeout")

	// KeyMaxRawRequestCount is the key for the maximum raw request count parameter
	KeyMaxRawRequestCount = []byte("MaxRawRequestCount")

	// KeyMaxCalldataSize is the key for the maximum calldata size parameter
	KeyMaxCalldataSize = []byte("MaxCalldataSize")

	// KeyMaxResultSize is the key for the maximum result size parameter
	KeyMaxResultSize = []byte("MaxResultSize")

	// KeyProviderRewardPercentage is the key for the provider reward percentage parameter
	KeyProviderRewardPercentage = []byte("ProviderRewardPercentage")

	// KeyProviderReputationDecayRate is the key for the provider reputation decay rate parameter
	KeyProviderReputationDecayRate = []byte("ProviderReputationDecayRate")

	// KeyMinProviderReputation is the key for the minimum provider reputation parameter
	KeyMinProviderReputation = []byte("MinProviderReputation")

	// KeyReputationBonusRate is the key for the reputation bonus rate parameter
	KeyReputationBonusRate = []byte("ReputationBonusRate")

	// KeyReputationPenaltyRate is the key for the reputation penalty rate parameter
	KeyReputationPenaltyRate = []byte("ReputationPenaltyRate")

	// KeyMaxHistorySize is the key for the maximum history size parameter
	KeyMaxHistorySize = []byte("MaxHistorySize")

	// KeyMaxRequestsPerBlock is the key for the maximum requests per block parameter
	KeyMaxRequestsPerBlock = []byte("MaxRequestsPerBlock")

	// KeyMaxResponsesPerBlock is the key for the maximum responses per block parameter
	KeyMaxResponsesPerBlock = []byte("MaxResponsesPerBlock")
)

// Event types
const (
	// EventTypeRegisterOracleProvider is the event type for registering an oracle provider
	EventTypeRegisterOracleProvider = "register_oracle_provider"

	// EventTypeUpdateOracleProvider is the event type for updating an oracle provider
	EventTypeUpdateOracleProvider = "update_oracle_provider"

	// EventTypeDeregisterOracleProvider is the event type for deregistering an oracle provider
	EventTypeDeregisterOracleProvider = "deregister_oracle_provider"

	// EventTypeStakeOracleProvider is the event type for staking to an oracle provider
	EventTypeStakeOracleProvider = "stake_oracle_provider"

	// EventTypeUnstakeOracleProvider is the event type for unstaking from an oracle provider
	EventTypeUnstakeOracleProvider = "unstake_oracle_provider"

	// EventTypeCreateOracleRequest is the event type for creating an oracle request
	EventTypeCreateOracleRequest = "create_oracle_request"

	// EventTypeSubmitOracleResponse is the event type for submitting an oracle response
	EventTypeSubmitOracleResponse = "submit_oracle_response"

	// EventTypeCancelOracleRequest is the event type for canceling an oracle request
	EventTypeCancelOracleRequest = "cancel_oracle_request"

	// EventTypeUpdateProviderReputation is the event type for updating a provider's reputation
	EventTypeUpdateProviderReputation = "update_provider_reputation"

	// EventTypeCreateDataSource is the event type for creating a data source
	EventTypeCreateDataSource = "create_data_source"

	// EventTypeUpdateDataSource is the event type for updating a data source
	EventTypeUpdateDataSource = "update_data_source"

	// EventTypeRemoveDataSource is the event type for removing a data source
	EventTypeRemoveDataSource = "remove_data_source"

	// EventTypeCreateOracleScript is the event type for creating an oracle script
	EventTypeCreateOracleScript = "create_oracle_script"

	// EventTypeUpdateOracleScript is the event type for updating an oracle script
	EventTypeUpdateOracleScript = "update_oracle_script"

	// EventTypeRemoveOracleScript is the event type for removing an oracle script
	EventTypeRemoveOracleScript = "remove_oracle_script"

	// EventTypeResolveOracleRequest is the event type for resolving an oracle request
	EventTypeResolveOracleRequest = "resolve_oracle_request"

	// EventTypeTimeoutOracleRequest is the event type for timing out an oracle request
	EventTypeTimeoutOracleRequest = "timeout_oracle_request"

	// EventTypeRewardOracleProvider is the event type for rewarding an oracle provider
	EventTypeRewardOracleProvider = "reward_oracle_provider"
)

// Event attributes
const (
	// AttributeKeyProvider is the attribute key for a provider
	AttributeKeyProvider = "provider"

	// AttributeKeyName is the attribute key for a name
	AttributeKeyName = "name"

	// AttributeKeyDescription is the attribute key for a description
	AttributeKeyDescription = "description"

	// AttributeKeyWebsite is the attribute key for a website
	AttributeKeyWebsite = "website"

	// AttributeKeyIdentity is the attribute key for an identity
	AttributeKeyIdentity = "identity"

	// AttributeKeyStakedAmount is the attribute key for a staked amount
	AttributeKeyStakedAmount = "staked_amount"

	// AttributeKeyRequestID is the attribute key for a request ID
	AttributeKeyRequestID = "request_id"

	// AttributeKeyOracleScriptID is the attribute key for an oracle script ID
	AttributeKeyOracleScriptID = "oracle_script_id"

	// AttributeKeyCalldata is the attribute key for calldata
	AttributeKeyCalldata = "calldata"

	// AttributeKeyAskCount is the attribute key for an ask count
	AttributeKeyAskCount = "ask_count"

	// AttributeKeyMinCount is the attribute key for a min count
	AttributeKeyMinCount = "min_count"

	// AttributeKeyClientID is the attribute key for a client ID
	AttributeKeyClientID = "client_id"

	// AttributeKeyFeeLimit is the attribute key for a fee limit
	AttributeKeyFeeLimit = "fee_limit"

	// AttributeKeyPrepareGas is the attribute key for prepare gas
	AttributeKeyPrepareGas = "prepare_gas"

	// AttributeKeyExecuteGas is the attribute key for execute gas
	AttributeKeyExecuteGas = "execute_gas"

	// AttributeKeyTimeoutBlocks is the attribute key for timeout blocks
	AttributeKeyTimeoutBlocks = "timeout_blocks"

	// AttributeKeyResponseID is the attribute key for a response ID
	AttributeKeyResponseID = "response_id"

	// AttributeKeyResult is the attribute key for a result
	AttributeKeyResult = "result"

	// AttributeKeyConfidence is the attribute key for a confidence level
	AttributeKeyConfidence = "confidence"

	// AttributeKeyReputationChange is the attribute key for a reputation change
	AttributeKeyReputationChange = "reputation_change"

	// AttributeKeyReason is the attribute key for a reason
	AttributeKeyReason = "reason"

	// AttributeKeyDataSourceID is the attribute key for a data source ID
	AttributeKeyDataSourceID = "data_source_id"

	// AttributeKeyExecutable is the attribute key for an executable
	AttributeKeyExecutable = "executable"

	// AttributeKeyFee is the attribute key for a fee
	AttributeKeyFee = "fee"

	// AttributeKeyOwner is the attribute key for an owner
	AttributeKeyOwner = "owner"

	// AttributeKeySchema is the attribute key for a schema
	AttributeKeySchema = "schema"

	// AttributeKeySourceCodeURL is the attribute key for a source code URL
	AttributeKeySourceCodeURL = "source_code_url"

	// AttributeKeyCode is the attribute key for code
	AttributeKeyCode = "code"

	// AttributeKeyReward is the attribute key for a reward
	AttributeKeyReward = "reward"

	// AttributeKeyStatus is the attribute key for a status
	AttributeKeyStatus = "status"
)

// Request statuses
const (
	// RequestStatusPending is the status for a pending request
	RequestStatusPending = "pending"

	// RequestStatusActive is the status for an active request
	RequestStatusActive = "active"

	// RequestStatusSuccessful is the status for a successful request
	RequestStatusSuccessful = "successful"

	// RequestStatusFailed is the status for a failed request
	RequestStatusFailed = "failed"

	// RequestStatusExpired is the status for an expired request
	RequestStatusExpired = "expired"

	// RequestStatusCanceled is the status for a canceled request
	RequestStatusCanceled = "canceled"
)

// Key functions

// OracleProviderKey returns the key for an oracle provider
func OracleProviderKey(address string) []byte {
	return append(OracleProviderKeyPrefix, []byte(address)...)
}

// OracleRequestKey returns the key for an oracle request
func OracleRequestKey(id uint64) []byte {
	return append(OracleRequestKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

// OracleResponseKey returns the key for an oracle response
func OracleResponseKey(id uint64) []byte {
	return append(OracleResponseKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

// ProviderResponseHistoryKey returns the key for a provider's response history
func ProviderResponseHistoryKey(address string) []byte {
	return append(ProviderResponseHistoryKeyPrefix, []byte(address)...)
}

// DataSourceKey returns the key for a data source
func DataSourceKey(id uint64) []byte {
	return append(DataSourceKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

// OracleScriptKey returns the key for an oracle script
func OracleScriptKey(id uint64) []byte {
	return append(OracleScriptKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

// ProviderRewardsKey returns the key for a provider's rewards
func ProviderRewardsKey(address string) []byte {
	return append(ProviderRewardsKeyPrefix, []byte(address)...)
}

// ProviderStatsKey returns the key for a provider's stats
func ProviderStatsKey(address string) []byte {
	return append(ProviderStatsKeyPrefix, []byte(address)...)
}

// ResultKey returns the key for a result
func ResultKey(requestID uint64) []byte {
	return append(ResultKeyPrefix, sdk.Uint64ToBigEndian(requestID)...)
}
