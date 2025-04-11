package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
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
	// Keys for store prefixes
	ValidatorAIModelKey = []byte{0x01} // key for storing AI models for validators
	AnomalyReportKey    = []byte{0x02} // key for storing anomaly reports
	NetworkStateKey     = []byte{0x03} // key for storing network state metrics
)

// AccountKeeper defines the expected account keeper
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// other methods from the interface you are implementing
}

// BankKeeper defines the expected bank keeper
type BankKeeper interface {
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	// other methods from the interface you are implementing
}

// StakingKeeper defines the expected staking keeper
type StakingKeeper interface {
	GetValidator(ctx sdk.Context, addr sdk.ValAddress) (validator stakingtypes.Validator, found bool)
	GetValidators(ctx sdk.Context, maxRetrieve uint32) (validators []stakingtypes.Validator)
	Slash(ctx sdk.Context, consAddr sdk.ConsAddress, infractionHeight int64, power int64, slashFactor sdk.Dec)
	// other methods from the interface you are implementing
}

// AIModel represents a neural network model used by validators
type AIModel struct {
	ValidatorAddress sdk.ValAddress `json:"validator_address"`
	ModelHash        string         `json:"model_hash"`
	ModelURL         string         `json:"model_url"`
	Version          uint64         `json:"version"`
	LastUpdated      int64          `json:"last_updated"`
}

// AnomalyReport represents a detected anomaly in the network
type AnomalyReport struct {
	ID               uint64         `json:"id"`
	ValidatorAddress sdk.ValAddress `json:"validator_address"`
	BlockHeight      int64          `json:"block_height"`
	AnomalyType      string         `json:"anomaly_type"`
	Confidence       sdk.Dec        `json:"confidence"`
	Description      string         `json:"description"`
	Evidence         string         `json:"evidence"`
	Timestamp        int64          `json:"timestamp"`
}

// NetworkState represents the current state of the network
type NetworkState struct {
	BlockHeight       int64    `json:"block_height"`
	TPS               sdk.Dec  `json:"tps"`
	AverageFeeRate    sdk.Dec  `json:"average_fee_rate"`
	NetworkCongestion sdk.Dec  `json:"network_congestion"`
	AnomalyScore      sdk.Dec  `json:"anomaly_score"`
	ValidatorCount    uint64   `json:"validator_count"`
	ActiveValidators  uint64   `json:"active_validators"`
	Timestamp         int64    `json:"timestamp"`
}