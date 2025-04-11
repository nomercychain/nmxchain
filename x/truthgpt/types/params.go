package types

import (
	"fmt"
	"time"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TruthGPT parameters
const (
	// DefaultMinProviderStake is the default minimum provider stake
	DefaultMinProviderStake = 1000000 // 1 NMX

	// DefaultMaxProviderCount is the default maximum provider count
	DefaultMaxProviderCount = 100

	// DefaultMinRequestFee is the default minimum request fee
	DefaultMinRequestFee = 10000 // 0.01 NMX

	// DefaultDefaultTimeout is the default timeout in blocks
	DefaultDefaultTimeout = 100

	// DefaultMaxRawRequestCount is the default maximum raw request count
	DefaultMaxRawRequestCount = 16

	// DefaultMaxCalldataSize is the default maximum calldata size in bytes
	DefaultMaxCalldataSize = 1024

	// DefaultMaxResultSize is the default maximum result size in bytes
	DefaultMaxResultSize = 1024

	// DefaultProviderRewardPercentage is the default provider reward percentage
	DefaultProviderRewardPercentage = "0.9" // 90%

	// DefaultProviderReputationDecayRate is the default provider reputation decay rate
	DefaultProviderReputationDecayRate = "0.01" // 1% per block

	// DefaultMinProviderReputation is the default minimum provider reputation
	DefaultMinProviderReputation = "0.5" // 50%

	// DefaultReputationBonusRate is the default reputation bonus rate
	DefaultReputationBonusRate = "0.01" // 1% per successful response

	// DefaultReputationPenaltyRate is the default reputation penalty rate
	DefaultReputationPenaltyRate = "0.02" // 2% per failed response

	// DefaultMaxHistorySize is the default maximum history size
	DefaultMaxHistorySize = 100

	// DefaultMaxRequestsPerBlock is the default maximum requests per block
	DefaultMaxRequestsPerBlock = 10

	// DefaultMaxResponsesPerBlock is the default maximum responses per block
	DefaultMaxResponsesPerBlock = 20
)

// Parameter store keys
var (
	KeyMinProviderStake           = []byte("MinProviderStake")
	KeyMaxProviderCount           = []byte("MaxProviderCount")
	KeyMinRequestFee              = []byte("MinRequestFee")
	KeyDefaultTimeout             = []byte("DefaultTimeout")
	KeyMaxRawRequestCount         = []byte("MaxRawRequestCount")
	KeyMaxCalldataSize            = []byte("MaxCalldataSize")
	KeyMaxResultSize              = []byte("MaxResultSize")
	KeyProviderRewardPercentage   = []byte("ProviderRewardPercentage")
	KeyProviderReputationDecayRate = []byte("ProviderReputationDecayRate")
	KeyMinProviderReputation      = []byte("MinProviderReputation")
	KeyReputationBonusRate        = []byte("ReputationBonusRate")
	KeyReputationPenaltyRate      = []byte("ReputationPenaltyRate")
	KeyMaxHistorySize             = []byte("MaxHistorySize")
	KeyMaxRequestsPerBlock        = []byte("MaxRequestsPerBlock")
	KeyMaxResponsesPerBlock       = []byte("MaxResponsesPerBlock")
)

// ParamKeyTable returns the parameter key table
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// Params defines the parameters for the TruthGPT module
type Params struct {
	MinProviderStake           int64    `json:"min_provider_stake"`
	MaxProviderCount           uint32   `json:"max_provider_count"`
	MinRequestFee              int64    `json:"min_request_fee"`
	DefaultTimeout             int64    `json:"default_timeout"`
	MaxRawRequestCount         uint32   `json:"max_raw_request_count"`
	MaxCalldataSize            uint32   `json:"max_calldata_size"`
	MaxResultSize              uint32   `json:"max_result_size"`
	ProviderRewardPercentage   sdk.Dec  `json:"provider_reward_percentage"`
	ProviderReputationDecayRate sdk.Dec  `json:"provider_reputation_decay_rate"`
	MinProviderReputation      sdk.Dec  `json:"min_provider_reputation"`
	ReputationBonusRate        sdk.Dec  `json:"reputation_bonus_rate"`
	ReputationPenaltyRate      sdk.Dec  `json:"reputation_penalty_rate"`
	MaxHistorySize             uint32   `json:"max_history_size"`
	MaxRequestsPerBlock        uint32   `json:"max_requests_per_block"`
	MaxResponsesPerBlock       uint32   `json:"max_responses_per_block"`
}

// DefaultParams returns default parameters
func DefaultParams() Params {
	return Params{
		MinProviderStake:           DefaultMinProviderStake,
		MaxProviderCount:           DefaultMaxProviderCount,
		MinRequestFee:              DefaultMinRequestFee,
		DefaultTimeout:             DefaultDefaultTimeout,
		MaxRawRequestCount:         DefaultMaxRawRequestCount,
		MaxCalldataSize:            DefaultMaxCalldataSize,
		MaxResultSize:              DefaultMaxResultSize,
		ProviderRewardPercentage:   sdk.MustNewDecFromStr(DefaultProviderRewardPercentage),
		ProviderReputationDecayRate: sdk.MustNewDecFromStr(DefaultProviderReputationDecayRate),
		MinProviderReputation:      sdk.MustNewDecFromStr(DefaultMinProviderReputation),
		ReputationBonusRate:        sdk.MustNewDecFromStr(DefaultReputationBonusRate),
		ReputationPenaltyRate:      sdk.MustNewDecFromStr(DefaultReputationPenaltyRate),
		MaxHistorySize:             DefaultMaxHistorySize,
		MaxRequestsPerBlock:        DefaultMaxRequestsPerBlock,
		MaxResponsesPerBlock:       DefaultMaxResponsesPerBlock,
	}
}

// ParamSetPairs implements the ParamSet interface
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMinProviderStake, &p.MinProviderStake, validateMinProviderStake),
		paramtypes.NewParamSetPair(KeyMaxProviderCount, &p.MaxProviderCount, validateMaxProviderCount),
		paramtypes.NewParamSetPair(KeyMinRequestFee, &p.MinRequestFee, validateMinRequestFee),
		paramtypes.NewParamSetPair(KeyDefaultTimeout, &p.DefaultTimeout, validateDefaultTimeout),
		paramtypes.NewParamSetPair(KeyMaxRawRequestCount, &p.MaxRawRequestCount, validateMaxRawRequestCount),
		paramtypes.NewParamSetPair(KeyMaxCalldataSize, &p.MaxCalldataSize, validateMaxCalldataSize),
		paramtypes.NewParamSetPair(KeyMaxResultSize, &p.MaxResultSize, validateMaxResultSize),
		paramtypes.NewParamSetPair(KeyProviderRewardPercentage, &p.ProviderRewardPercentage, validateProviderRewardPercentage),
		paramtypes.NewParamSetPair(KeyProviderReputationDecayRate, &p.ProviderReputationDecayRate, validateProviderReputationDecayRate),
		paramtypes.NewParamSetPair(KeyMinProviderReputation, &p.MinProviderReputation, validateMinProviderReputation),
		paramtypes.NewParamSetPair(KeyReputationBonusRate, &p.ReputationBonusRate, validateReputationBonusRate),
		paramtypes.NewParamSetPair(KeyReputationPenaltyRate, &p.ReputationPenaltyRate, validateReputationPenaltyRate),
		paramtypes.NewParamSetPair(KeyMaxHistorySize, &p.MaxHistorySize, validateMaxHistorySize),
		paramtypes.NewParamSetPair(KeyMaxRequestsPerBlock, &p.MaxRequestsPerBlock, validateMaxRequestsPerBlock),
		paramtypes.NewParamSetPair(KeyMaxResponsesPerBlock, &p.MaxResponsesPerBlock, validateMaxResponsesPerBlock),
	}
}

// Validate validates the parameters
func (p Params) Validate() error {
	if err := validateMinProviderStake(p.MinProviderStake); err != nil {
		return err
	}
	if err := validateMaxProviderCount(p.MaxProviderCount); err != nil {
		return err
	}
	if err := validateMinRequestFee(p.MinRequestFee); err != nil {
		return err
	}
	if err := validateDefaultTimeout(p.DefaultTimeout); err != nil {
		return err
	}
	if err := validateMaxRawRequestCount(p.MaxRawRequestCount); err != nil {
		return err
	}
	if err := validateMaxCalldataSize(p.MaxCalldataSize); err != nil {
		return err
	}
	if err := validateMaxResultSize(p.MaxResultSize); err != nil {
		return err
	}
	if err := validateProviderRewardPercentage(p.ProviderRewardPercentage); err != nil {
		return err
	}
	if err := validateProviderReputationDecayRate(p.ProviderReputationDecayRate); err != nil {
		return err
	}
	if err := validateMinProviderReputation(p.MinProviderReputation); err != nil {
		return err
	}
	if err := validateReputationBonusRate(p.ReputationBonusRate); err != nil {
		return err
	}
	if err := validateReputationPenaltyRate(p.ReputationPenaltyRate); err != nil {
		return err
	}
	if err := validateMaxHistorySize(p.MaxHistorySize); err != nil {
		return err
	}
	if err := validateMaxRequestsPerBlock(p.MaxRequestsPerBlock); err != nil {
		return err
	}
	if err := validateMaxResponsesPerBlock(p.MaxResponsesPerBlock); err != nil {
		return err
	}
	return nil
}

func validateMinProviderStake(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("minimum provider stake must be positive: %d", v)
	}

	return nil
}

func validateMaxProviderCount(i interface{}) error {
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("max provider count must be positive: %d", v)
	}

	return nil
}

func validateMinRequestFee(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v < 0 {
		return fmt.Errorf("minimum request fee cannot be negative: %d", v)
	}

	return nil
}

func validateDefaultTimeout(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("default timeout must be positive: %d", v)
	}

	return nil
}

func validateMaxRawRequestCount(i interface{}) error {
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("max raw request count must be positive: %d", v)
	}

	return nil
}

func validateMaxCalldataSize(i interface{}) error {
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("max calldata size must be positive: %d", v)
	}

	return nil
}

func validateMaxResultSize(i interface{}) error {
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("max result size must be positive: %d", v)
	}

	return nil
}

func validateProviderRewardPercentage(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("provider reward percentage cannot be negative: %s", v)
	}

	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("provider reward percentage cannot be greater than 1: %s", v)
	}

	return nil
}

func validateProviderReputationDecayRate(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("provider reputation decay rate cannot be negative: %s", v)
	}

	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("provider reputation decay rate cannot be greater than 1: %s", v)
	}

	return nil
}

func validateMinProviderReputation(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("minimum provider reputation cannot be negative: %s", v)
	}

	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("minimum provider reputation cannot be greater than 1: %s", v)
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

func validateMaxHistorySize(i interface{}) error {
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("max history size must be positive: %d", v)
	}

	return nil
}

func validateMaxRequestsPerBlock(i interface{}) error {
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("max requests per block must be positive: %d", v)
	}

	return nil
}

func validateMaxResponsesPerBlock(i interface{}) error {
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("max responses per block must be positive: %d", v)
	}

	return nil
}