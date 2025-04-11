package types

import (
	"fmt"
	"time"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NeuroPOS parameters
const (
	// DefaultUnbondingTime reflects three weeks in seconds
	DefaultUnbondingTime = time.Hour * 24 * 7 * 3

	// DefaultMaxValidators is the default maximum number of validators
	DefaultMaxValidators = 100

	// DefaultMinSelfDelegation is the default minimum self delegation
	DefaultMinSelfDelegation = 1000000 // 1 NMX

	// DefaultHistoricalEntries is the default number of historical entries
	DefaultHistoricalEntries = 10000

	// DefaultNeuralNetworkUpdateInterval is the default interval for neural network updates
	DefaultNeuralNetworkUpdateInterval = time.Hour * 24 // 1 day

	// DefaultNeuralNetworkLearningRate is the default learning rate for neural networks
	DefaultNeuralNetworkLearningRate = "0.001"

	// DefaultReputationDecayRate is the default rate at which reputation decays
	DefaultReputationDecayRate = "0.01"

	// DefaultPerformanceAssessmentWindow is the default window for performance assessment
	DefaultPerformanceAssessmentWindow = 10000

	// DefaultMinValidatorReputation is the default minimum validator reputation
	DefaultMinValidatorReputation = "0.5"

	// DefaultMaxMissedBlocks is the default maximum number of blocks a validator can miss
	DefaultMaxMissedBlocks = 100

	// DefaultSignedBlocksWindow is the default window for signed blocks
	DefaultSignedBlocksWindow = 10000

	// DefaultMinSignedPerWindow is the default minimum signed blocks per window
	DefaultMinSignedPerWindow = "0.05"

	// DefaultDowntimeJailDuration is the default downtime jail duration
	DefaultDowntimeJailDuration = time.Hour * 24 // 1 day

	// DefaultSlashFractionDoubleSign is the default slash fraction for double signing
	DefaultSlashFractionDoubleSign = "0.05"

	// DefaultSlashFractionDowntime is the default slash fraction for downtime
	DefaultSlashFractionDowntime = "0.01"

	// DefaultReputationBonusRate is the default rate for reputation bonuses
	DefaultReputationBonusRate = "0.01"

	// DefaultReputationPenaltyRate is the default rate for reputation penalties
	DefaultReputationPenaltyRate = "0.02"

	// DefaultNeuralNetworkInfluenceRate is the default influence rate of neural networks
	DefaultNeuralNetworkInfluenceRate = "0.3"
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

// ParamKeyTable returns the parameter key table
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// Params defines the parameters for the NeuroPOS module
type Params struct {
	UnbondingTime               time.Duration `json:"unbonding_time"`
	MaxValidators               uint32        `json:"max_validators"`
	MinSelfDelegation           int64         `json:"min_self_delegation"`
	HistoricalEntries           uint32        `json:"historical_entries"`
	NeuralNetworkUpdateInterval time.Duration `json:"neural_network_update_interval"`
	NeuralNetworkLearningRate   sdk.Dec       `json:"neural_network_learning_rate"`
	NeuralNetworkArchitecture   string        `json:"neural_network_architecture"`
	ReputationDecayRate         sdk.Dec       `json:"reputation_decay_rate"`
	PerformanceAssessmentWindow uint64        `json:"performance_assessment_window"`
	MinValidatorReputation      sdk.Dec       `json:"min_validator_reputation"`
	MaxMissedBlocks             uint64        `json:"max_missed_blocks"`
	SignedBlocksWindow          int64         `json:"signed_blocks_window"`
	MinSignedPerWindow          sdk.Dec       `json:"min_signed_per_window"`
	DowntimeJailDuration        time.Duration `json:"downtime_jail_duration"`
	SlashFractionDoubleSign     sdk.Dec       `json:"slash_fraction_double_sign"`
	SlashFractionDowntime       sdk.Dec       `json:"slash_fraction_downtime"`
	ReputationBonusRate         sdk.Dec       `json:"reputation_bonus_rate"`
	ReputationPenaltyRate       sdk.Dec       `json:"reputation_penalty_rate"`
	NeuralNetworkInfluenceRate  sdk.Dec       `json:"neural_network_influence_rate"`
}

// DefaultParams returns default parameters
func DefaultParams() Params {
	return Params{
		UnbondingTime:               DefaultUnbondingTime,
		MaxValidators:               DefaultMaxValidators,
		MinSelfDelegation:           DefaultMinSelfDelegation,
		HistoricalEntries:           DefaultHistoricalEntries,
		NeuralNetworkUpdateInterval: DefaultNeuralNetworkUpdateInterval,
		NeuralNetworkLearningRate:   sdk.MustNewDecFromStr(DefaultNeuralNetworkLearningRate),
		NeuralNetworkArchitecture:   NeuralNetworkArchitectureMLP,
		ReputationDecayRate:         sdk.MustNewDecFromStr(DefaultReputationDecayRate),
		PerformanceAssessmentWindow: DefaultPerformanceAssessmentWindow,
		MinValidatorReputation:      sdk.MustNewDecFromStr(DefaultMinValidatorReputation),
		MaxMissedBlocks:             DefaultMaxMissedBlocks,
		SignedBlocksWindow:          DefaultSignedBlocksWindow,
		MinSignedPerWindow:          sdk.MustNewDecFromStr(DefaultMinSignedPerWindow),
		DowntimeJailDuration:        DefaultDowntimeJailDuration,
		SlashFractionDoubleSign:     sdk.MustNewDecFromStr(DefaultSlashFractionDoubleSign),
		SlashFractionDowntime:       sdk.MustNewDecFromStr(DefaultSlashFractionDowntime),
		ReputationBonusRate:         sdk.MustNewDecFromStr(DefaultReputationBonusRate),
		ReputationPenaltyRate:       sdk.MustNewDecFromStr(DefaultReputationPenaltyRate),
		NeuralNetworkInfluenceRate:  sdk.MustNewDecFromStr(DefaultNeuralNetworkInfluenceRate),
	}
}

// ParamSetPairs implements the ParamSet interface
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyUnbondingTime, &p.UnbondingTime, validateUnbondingTime),
		paramtypes.NewParamSetPair(KeyMaxValidators, &p.MaxValidators, validateMaxValidators),
		paramtypes.NewParamSetPair(KeyMinSelfDelegation, &p.MinSelfDelegation, validateMinSelfDelegation),
		paramtypes.NewParamSetPair(KeyHistoricalEntries, &p.HistoricalEntries, validateHistoricalEntries),
		paramtypes.NewParamSetPair(KeyNeuralNetworkUpdateInterval, &p.NeuralNetworkUpdateInterval, validateNeuralNetworkUpdateInterval),
		paramtypes.NewParamSetPair(KeyNeuralNetworkLearningRate, &p.NeuralNetworkLearningRate, validateNeuralNetworkLearningRate),
		paramtypes.NewParamSetPair(KeyNeuralNetworkArchitecture, &p.NeuralNetworkArchitecture, validateNeuralNetworkArchitecture),
		paramtypes.NewParamSetPair(KeyReputationDecayRate, &p.ReputationDecayRate, validateReputationDecayRate),
		paramtypes.NewParamSetPair(KeyPerformanceAssessmentWindow, &p.PerformanceAssessmentWindow, validatePerformanceAssessmentWindow),
		paramtypes.NewParamSetPair(KeyMinValidatorReputation, &p.MinValidatorReputation, validateMinValidatorReputation),
		paramtypes.NewParamSetPair(KeyMaxMissedBlocks, &p.MaxMissedBlocks, validateMaxMissedBlocks),
		paramtypes.NewParamSetPair(KeySignedBlocksWindow, &p.SignedBlocksWindow, validateSignedBlocksWindow),
		paramtypes.NewParamSetPair(KeyMinSignedPerWindow, &p.MinSignedPerWindow, validateMinSignedPerWindow),
		paramtypes.NewParamSetPair(KeyDowntimeJailDuration, &p.DowntimeJailDuration, validateDowntimeJailDuration),
		paramtypes.NewParamSetPair(KeySlashFractionDoubleSign, &p.SlashFractionDoubleSign, validateSlashFractionDoubleSign),
		paramtypes.NewParamSetPair(KeySlashFractionDowntime, &p.SlashFractionDowntime, validateSlashFractionDowntime),
		paramtypes.NewParamSetPair(KeyReputationBonusRate, &p.ReputationBonusRate, validateReputationBonusRate),
		paramtypes.NewParamSetPair(KeyReputationPenaltyRate, &p.ReputationPenaltyRate, validateReputationPenaltyRate),
		paramtypes.NewParamSetPair(KeyNeuralNetworkInfluenceRate, &p.NeuralNetworkInfluenceRate, validateNeuralNetworkInfluenceRate),
	}
}

// Validate validates the parameters
func (p Params) Validate() error {
	if err := validateUnbondingTime(p.UnbondingTime); err != nil {
		return err
	}
	if err := validateMaxValidators(p.MaxValidators); err != nil {
		return err
	}
	if err := validateMinSelfDelegation(p.MinSelfDelegation); err != nil {
		return err
	}
	if err := validateHistoricalEntries(p.HistoricalEntries); err != nil {
		return err
	}
	if err := validateNeuralNetworkUpdateInterval(p.NeuralNetworkUpdateInterval); err != nil {
		return err
	}
	if err := validateNeuralNetworkLearningRate(p.NeuralNetworkLearningRate); err != nil {
		return err
	}
	if err := validateNeuralNetworkArchitecture(p.NeuralNetworkArchitecture); err != nil {
		return err
	}
	if err := validateReputationDecayRate(p.ReputationDecayRate); err != nil {
		return err
	}
	if err := validatePerformanceAssessmentWindow(p.PerformanceAssessmentWindow); err != nil {
		return err
	}
	if err := validateMinValidatorReputation(p.MinValidatorReputation); err != nil {
		return err
	}
	if err := validateMaxMissedBlocks(p.MaxMissedBlocks); err != nil {
		return err
	}
	if err := validateSignedBlocksWindow(p.SignedBlocksWindow); err != nil {
		return err
	}
	if err := validateMinSignedPerWindow(p.MinSignedPerWindow); err != nil {
		return err
	}
	if err := validateDowntimeJailDuration(p.DowntimeJailDuration); err != nil {
		return err
	}
	if err := validateSlashFractionDoubleSign(p.SlashFractionDoubleSign); err != nil {
		return err
	}
	if err := validateSlashFractionDowntime(p.SlashFractionDowntime); err != nil {
		return err
	}
	if err := validateReputationBonusRate(p.ReputationBonusRate); err != nil {
		return err
	}
	if err := validateReputationPenaltyRate(p.ReputationPenaltyRate); err != nil {
		return err
	}
	if err := validateNeuralNetworkInfluenceRate(p.NeuralNetworkInfluenceRate); err != nil {
		return err
	}
	return nil
}

func validateUnbondingTime(i interface{}) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("unbonding time must be positive: %d", v)
	}

	return nil
}

func validateMaxValidators(i interface{}) error {
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("max validators must be positive: %d", v)
	}

	return nil
}

func validateMinSelfDelegation(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("minimum self delegation must be positive: %d", v)
	}

	return nil
}

func validateHistoricalEntries(i interface{}) error {
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("historical entries must be positive: %d", v)
	}

	return nil
}

func validateNeuralNetworkUpdateInterval(i interface{}) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("neural network update interval must be positive: %d", v)
	}

	return nil
}

func validateNeuralNetworkLearningRate(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("neural network learning rate cannot be negative: %s", v)
	}

	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("neural network learning rate cannot be greater than 1: %s", v)
	}

	return nil
}

func validateNeuralNetworkArchitecture(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	valid := false
	validArchitectures := []string{
		NeuralNetworkArchitectureMLP,
		NeuralNetworkArchitectureCNN,
		NeuralNetworkArchitectureRNN,
		NeuralNetworkArchitectureLSTM,
		NeuralNetworkArchitectureGRU,
		NeuralNetworkArchitectureTransformer,
		NeuralNetworkArchitectureAutoencoder,
		NeuralNetworkArchitectureGAN,
	}

	for _, arch := range validArchitectures {
		if v == arch {
			valid = true
			break
		}
	}

	if !valid {
		return fmt.Errorf("invalid neural network architecture: %s", v)
	}

	return nil
}

func validateReputationDecayRate(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("reputation decay rate cannot be negative: %s", v)
	}

	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("reputation decay rate cannot be greater than 1: %s", v)
	}

	return nil
}

func validatePerformanceAssessmentWindow(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("performance assessment window must be positive: %d", v)
	}

	return nil
}

func validateMinValidatorReputation(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("minimum validator reputation cannot be negative: %s", v)
	}

	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("minimum validator reputation cannot be greater than 1: %s", v)
	}

	return nil
}

func validateMaxMissedBlocks(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("max missed blocks must be positive: %d", v)
	}

	return nil
}

func validateSignedBlocksWindow(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("signed blocks window must be positive: %d", v)
	}

	return nil
}

func validateMinSignedPerWindow(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("min signed per window cannot be negative: %s", v)
	}

	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("min signed per window cannot be greater than 1: %s", v)
	}

	return nil
}

func validateDowntimeJailDuration(i interface{}) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("downtime jail duration must be positive: %d", v)
	}

	return nil
}

func validateSlashFractionDoubleSign(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("slash fraction double sign cannot be negative: %s", v)
	}

	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("slash fraction double sign cannot be greater than 1: %s", v)
	}

	return nil
}

func validateSlashFractionDowntime(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("slash fraction downtime cannot be negative: %s", v)
	}

	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("slash fraction downtime cannot be greater than 1: %s", v)
	}

	return nil
}

func validateReputationBonusRate(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("reputation bonus rate cannot be negative: %s", v)
	}

	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("reputation bonus rate cannot be greater than 1: %s", v)
	}

	return nil
}

func validateReputationPenaltyRate(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("reputation penalty rate cannot be negative: %s", v)
	}

	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("reputation penalty rate cannot be greater than 1: %s", v)
	}

	return nil
}

func validateNeuralNetworkInfluenceRate(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("neural network influence rate cannot be negative: %s", v)
	}

	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("neural network influence rate cannot be greater than 1: %s", v)
	}

	return nil
}