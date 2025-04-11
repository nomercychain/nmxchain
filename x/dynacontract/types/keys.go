package types

const (
	// ModuleName defines the module name
	ModuleName = "dynacontract"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_dynacontract"
)

var (
	// DynaContractKeyPrefix is the prefix for dynamic contract keys
	DynaContractKeyPrefix = []byte{0x01}

	// DynaContractExecutionKeyPrefix is the prefix for dynamic contract execution keys
	DynaContractExecutionKeyPrefix = []byte{0x02}

	// DynaContractTemplateKeyPrefix is the prefix for dynamic contract template keys
	DynaContractTemplateKeyPrefix = []byte{0x03}

	// DynaContractLearningDataKeyPrefix is the prefix for dynamic contract learning data keys
	DynaContractLearningDataKeyPrefix = []byte{0x04}

	// DynaContractPermissionKeyPrefix is the prefix for dynamic contract permission keys
	DynaContractPermissionKeyPrefix = []byte{0x05}
)

// Parameter store keys
var (
	KeyMaxContractSize     = []byte("MaxContractSize")
	KeyMaxContractGas      = []byte("MaxContractGas")
	KeyMaxLearningDataSize = []byte("MaxLearningDataSize")
	KeyMaxMetadataSize     = []byte("MaxMetadataSize")
	KeyMinContractDeposit  = []byte("MinContractDeposit")
	KeyExecutionFeeRate    = []byte("ExecutionFeeRate")
)

// Event types
const (
	EventTypeCreateDynaContract         = "create_dyna_contract"
	EventTypeUpdateDynaContract         = "update_dyna_contract"
	EventTypeExecuteDynaContract        = "execute_dyna_contract"
	EventTypeAddLearningData            = "add_learning_data"
	EventTypeCreateDynaContractTemplate = "create_dyna_contract_template"
	EventTypeInstantiateFromTemplate    = "instantiate_from_template"
	EventTypeGrantContractPermission    = "grant_contract_permission"
	EventTypeRevokeContractPermission   = "revoke_contract_permission"
)

// Event attribute keys
const (
	AttributeKeyContractID     = "contract_id"
	AttributeKeyCreator        = "creator"
	AttributeKeyOwner          = "owner"
	AttributeKeyName           = "name"
	AttributeKeyContractType   = "contract_type"
	AttributeKeyDeposit        = "deposit"
	AttributeKeyCaller         = "caller"
	AttributeKeyExecutionID    = "execution_id"
	AttributeKeyGasUsed        = "gas_used"
	AttributeKeyDataID         = "data_id"
	AttributeKeyDataType       = "data_type"
	AttributeKeyTemplateID     = "template_id"
	AttributeKeyAddress        = "address"
	AttributeKeyPermissionType = "permission_type"
	AttributeKeyExpirationDays = "expiration_days"
)
