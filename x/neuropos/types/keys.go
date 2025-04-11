package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "neuropos"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_neuropos"
)

var (
	// ValidatorKeyPrefix is the prefix for validator keys
	ValidatorKeyPrefix = []byte{0x01}

	// DelegationKeyPrefix is the prefix for delegation keys
	DelegationKeyPrefix = []byte{0x02}

	// UnbondingDelegationKeyPrefix is the prefix for unbonding delegation keys
	UnbondingDelegationKeyPrefix = []byte{0x03}

	// RedelegationKeyPrefix is the prefix for redelegation keys
	RedelegationKeyPrefix = []byte{0x04}

	// ValidatorQueueKeyPrefix is the prefix for the validator queue
	ValidatorQueueKeyPrefix = []byte{0x05}

	// HistoricalInfoKeyPrefix is the prefix for historical info
	HistoricalInfoKeyPrefix = []byte{0x06}

	// NeuralNetworkKeyPrefix is the prefix for neural network keys
	NeuralNetworkKeyPrefix = []byte{0x07}

	// NeuralNetworkWeightKeyPrefix is the prefix for neural network weight keys
	NeuralNetworkWeightKeyPrefix = []byte{0x08}

	// NeuralNetworkTrainingDataKeyPrefix is the prefix for neural network training data keys
	NeuralNetworkTrainingDataKeyPrefix = []byte{0x09}

	// NeuralNetworkPredictionKeyPrefix is the prefix for neural network prediction keys
	NeuralNetworkPredictionKeyPrefix = []byte{0x0A}

	// ValidatorPerformanceKeyPrefix is the prefix for validator performance keys
	ValidatorPerformanceKeyPrefix = []byte{0x0B}

	// ValidatorReputationKeyPrefix is the prefix for validator reputation keys
	ValidatorReputationKeyPrefix = []byte{0x0C}

	// ValidatorSlashEventKeyPrefix is the prefix for validator slash event keys
	ValidatorSlashEventKeyPrefix = []byte{0x0D}

	// ValidatorMissedBlocksKeyPrefix is the prefix for validator missed blocks keys
	ValidatorMissedBlocksKeyPrefix = []byte{0x0E}

	// ValidatorSigningInfoKeyPrefix is the prefix for validator signing info keys
	ValidatorSigningInfoKeyPrefix = []byte{0x0F}

	// ValidatorAIModelKey is the prefix for validator AI model keys
	ValidatorAIModelKey = []byte{0x30}

	// AnomalyReportKey is the prefix for anomaly report keys
	AnomalyReportKey = []byte{0x31}

	// NetworkStateKey is the key for the network state
	NetworkStateKey = []byte{0x40}
)

// Parameter store keys
var (
	KeyUnbondingTime               = []byte("UnbondingTime")
	KeyMaxValidators               = []byte("MaxValidators")
	KeyMinSelfDelegation           = []byte("MinSelfDelegation")
	KeyHistoricalEntries           = []byte("HistoricalEntries")
	KeyNeuralNetworkUpdateInterval = []byte("NeuralNetworkUpdateInterval")
	KeyNeuralNetworkLearningRate   = []byte("NeuralNetworkLearningRate")
	KeyNeuralNetworkArchitecture   = []byte("NeuralNetworkArchitecture")
	KeyReputationDecayRate         = []byte("ReputationDecayRate")
	KeyPerformanceAssessmentWindow = []byte("PerformanceAssessmentWindow")
	KeyMinValidatorReputation      = []byte("MinValidatorReputation")
	KeyMaxMissedBlocks             = []byte("MaxMissedBlocks")
	KeySignedBlocksWindow          = []byte("SignedBlocksWindow")
	KeyMinSignedPerWindow          = []byte("MinSignedPerWindow")
	KeyDowntimeJailDuration        = []byte("DowntimeJailDuration")
	KeySlashFractionDoubleSign     = []byte("SlashFractionDoubleSign")
	KeySlashFractionDowntime       = []byte("SlashFractionDowntime")
	KeyReputationBonusRate         = []byte("ReputationBonusRate")
	KeyReputationPenaltyRate       = []byte("ReputationPenaltyRate")
	KeyNeuralNetworkInfluenceRate  = []byte("NeuralNetworkInfluenceRate")
)

// Event types
const (
	EventTypeCreateValidator           = "create_validator"
	EventTypeEditValidator             = "edit_validator"
	EventTypeDelegate                  = "delegate"
	EventTypeUnbond                    = "unbond"
	EventTypeRedelegate                = "redelegate"
	EventTypeSlash                     = "slash"
	EventTypeJail                      = "jail"
	EventTypeUnjail                    = "unjail"
	EventTypeCreateNeuralNetwork       = "create_neural_network"
	EventTypeUpdateNeuralNetwork       = "update_neural_network"
	EventTypeTrainNeuralNetwork        = "train_neural_network"
	EventTypeSubmitNeuralPrediction    = "submit_neural_prediction"
	EventTypeUpdateValidatorReputation = "update_validator_reputation"
	EventTypeValidatorPerformance      = "validator_performance"
	EventTypeAnomalyDetected           = "anomaly_detected"
)

// Neural network architectures
const (
	// NeuralNetworkArchitectureMLP is the architecture for a multi-layer perceptron
	NeuralNetworkArchitectureMLP = "mlp"

	// NeuralNetworkArchitectureCNN is the architecture for a convolutional neural network
	NeuralNetworkArchitectureCNN = "cnn"

	// NeuralNetworkArchitectureRNN is the architecture for a recurrent neural network
	NeuralNetworkArchitectureRNN = "rnn"

	// NeuralNetworkArchitectureLSTM is the architecture for a long short-term memory network
	NeuralNetworkArchitectureLSTM = "lstm"

	// NeuralNetworkArchitectureGRU is the architecture for a gated recurrent unit network
	NeuralNetworkArchitectureGRU = "gru"

	// NeuralNetworkArchitectureTransformer is the architecture for a transformer network
	NeuralNetworkArchitectureTransformer = "transformer"

	// NeuralNetworkArchitectureAutoencoder is the architecture for an autoencoder
	NeuralNetworkArchitectureAutoencoder = "autoencoder"

	// NeuralNetworkArchitectureGAN is the architecture for a generative adversarial network
	NeuralNetworkArchitectureGAN = "gan"
)

// Neural network statuses
const (
	// NeuralNetworkStatusActive is the status for an active neural network
	NeuralNetworkStatusActive = "active"

	// NeuralNetworkStatusUpdating is the status for a neural network being updated
	NeuralNetworkStatusUpdating = "updating"

	// NeuralNetworkStatusTraining is the status for a neural network being trained
	NeuralNetworkStatusTraining = "training"

	// NeuralNetworkStatusInactive is the status for an inactive neural network
	NeuralNetworkStatusInactive = "inactive"
)

// Validator statuses
const (
	// BondStatusUnbonded is the status for an unbonded validator
	BondStatusUnbonded = "unbonded"

	// BondStatusUnbonding is the status for a validator that is unbonding
	BondStatusUnbonding = "unbonding"

	// BondStatusBonded is the status for a bonded validator
	BondStatusBonded = "bonded"
)

// Key functions

// ValidatorKey returns the key for a validator
func ValidatorKey(operatorAddr string) []byte {
	return append(ValidatorKeyPrefix, []byte(operatorAddr)...)
}

// DelegationKey returns the key for a delegation
func DelegationKey(delegatorAddr, validatorAddr string) []byte {
	return append(append(DelegationKeyPrefix, []byte(delegatorAddr+"/")...), []byte(validatorAddr)...)
}

// UnbondingDelegationKey returns the key for an unbonding delegation
func UnbondingDelegationKey(delegatorAddr, validatorAddr string) []byte {
	return append(append(UnbondingDelegationKeyPrefix, []byte(delegatorAddr+"/")...), []byte(validatorAddr)...)
}

// RedelegationKey returns the key for a redelegation
func RedelegationKey(delegatorAddr, validatorSrcAddr, validatorDstAddr string) []byte {
	return append(append(append(RedelegationKeyPrefix, []byte(delegatorAddr+"/")...), []byte(validatorSrcAddr+"/")...), []byte(validatorDstAddr)...)
}

// NeuralNetworkKey returns the key for a neural network
func NeuralNetworkKey(networkID string) []byte {
	return append(NeuralNetworkKeyPrefix, []byte(networkID)...)
}

// NeuralNetworkWeightKey returns the key for a neural network weight
func NeuralNetworkWeightKey(networkID string, version uint64) []byte {
	versionBytes := sdk.Uint64ToBigEndian(version)
	return append(append(NeuralNetworkWeightKeyPrefix, []byte(networkID+"/")...), versionBytes...)
}

// TrainingDataKey returns the key for training data
func TrainingDataKey(dataID string) []byte {
	return append(NeuralNetworkTrainingDataKeyPrefix, []byte(dataID)...)
}

// NeuralPredictionKey returns the key for a neural prediction
func NeuralPredictionKey(predictionID string) []byte {
	return append(NeuralNetworkPredictionKeyPrefix, []byte(predictionID)...)
}

// ValidatorPerformanceKey returns the key for a validator's performance
func ValidatorPerformanceKey(validatorAddr string) []byte {
	return append(ValidatorPerformanceKeyPrefix, []byte(validatorAddr)...)
}

// ValidatorReputationKey returns the key for a validator's reputation
func ValidatorReputationKey(validatorAddr string) []byte {
	return append(ValidatorReputationKeyPrefix, []byte(validatorAddr)...)
}

// ValidatorSlashEventKey returns the key for a validator's slash event
func ValidatorSlashEventKey(validatorAddr string, height int64) []byte {
	heightBytes := sdk.Uint64ToBigEndian(uint64(height))
	return append(append(ValidatorSlashEventKeyPrefix, []byte(validatorAddr+"/")...), heightBytes...)
}

// ValidatorSigningInfoKey returns the key for a validator's signing info
func ValidatorSigningInfoKey(validatorAddr string) []byte {
	return append(ValidatorSigningInfoKeyPrefix, []byte(validatorAddr)...)
}

// Event attribute keys
const (
	AttributeKeyValidator             = "validator"
	AttributeKeyDelegator             = "delegator"
	AttributeKeyDestination           = "destination"
	AttributeKeySource                = "source"
	AttributeKeyAmount                = "amount"
	AttributeKeyCompletionTime        = "completion_time"
	AttributeKeySlashReason           = "slash_reason"
	AttributeKeySlashFactor           = "slash_factor"
	AttributeKeyPeriod                = "period"
	AttributeKeyNeuralNetworkID       = "neural_network_id"
	AttributeKeyNeuralNetworkAccuracy = "neural_network_accuracy"
	AttributeKeyNeuralNetworkLoss     = "neural_network_loss"
	AttributeKeyTrainingDataSize      = "training_data_size"
	AttributeKeyEpochs                = "epochs"
	AttributeKeyValidatorReputation   = "validator_reputation"
	AttributeKeyReputationChange      = "reputation_change"
	AttributeKeyReputationReason      = "reputation_reason"
	AttributeKeyPerformanceScore      = "performance_score"
	AttributeKeyBlocksProposed        = "blocks_proposed"
	AttributeKeyBlocksValidated       = "blocks_validated"
	AttributeKeyMissedBlocks          = "missed_blocks"
	AttributeKeyPredictionAccuracy    = "prediction_accuracy"
	AttributeKeyPredictionConfidence  = "prediction_confidence"
	AttributeKeyPredictionTarget      = "prediction_target"
	AttributeKeyPredictionResult      = "prediction_result"
	AttributeKeyNeuralPredictionID    = "neural_prediction_id"
	AttributeKeyAnomalyID             = "anomaly_id"
	AttributeKeyAnomalyConfidence     = "anomaly_confidence"
	AttributeKeyAnomalyType           = "anomaly_type"
)
