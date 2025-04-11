package types

import (
	"fmt"
	
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Parameter store keys
var (
	KeyMinAgentDeposit        = []byte("MinAgentDeposit")
	KeyMaxAgentNameLength     = []byte("MaxAgentNameLength")
	KeyMaxAgentDescLength     = []byte("MaxAgentDescLength")
	KeyMaxTrainingDataSize    = []byte("MaxTrainingDataSize")
	KeyMaxMarketplaceListings = []byte("MaxMarketplaceListings")
	KeyMarketplaceFeeRate     = []byte("MarketplaceFeeRate")
)

// ParamKeyTable returns the parameter key table
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// DefaultParams returns default parameters
func DefaultParams() Params {
	return Params{
		MinAgentDeposit:        sdk.NewCoin("unmx", sdk.NewInt(100000000)), // 100 NMX
		MaxAgentNameLength:     50,
		MaxAgentDescLength:     500,
		MaxTrainingDataSize:    1048576, // 1MB
		MaxMarketplaceListings: 100,     // Per account
		MarketplaceFeeRate:     sdk.NewDecWithPrec(25, 3), // 2.5%
	}
}

// ParamSetPairs implements the ParamSet interface
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMinAgentDeposit, &p.MinAgentDeposit, validateMinAgentDeposit),
		paramtypes.NewParamSetPair(KeyMaxAgentNameLength, &p.MaxAgentNameLength, validateUint64),
		paramtypes.NewParamSetPair(KeyMaxAgentDescLength, &p.MaxAgentDescLength, validateUint64),
		paramtypes.NewParamSetPair(KeyMaxTrainingDataSize, &p.MaxTrainingDataSize, validateUint64),
		paramtypes.NewParamSetPair(KeyMaxMarketplaceListings, &p.MaxMarketplaceListings, validateUint64),
		paramtypes.NewParamSetPair(KeyMarketplaceFeeRate, &p.MarketplaceFeeRate, validateMarketplaceFeeRate),
	}
}

// Validate validates the parameters
func (p Params) Validate() error {
	if err := validateMinAgentDeposit(p.MinAgentDeposit); err != nil {
		return err
	}
	if err := validateUint64(p.MaxAgentNameLength); err != nil {
		return err
	}
	if err := validateUint64(p.MaxAgentDescLength); err != nil {
		return err
	}
	if err := validateUint64(p.MaxTrainingDataSize); err != nil {
		return err
	}
	if err := validateUint64(p.MaxMarketplaceListings); err != nil {
		return err
	}
	if err := validateMarketplaceFeeRate(p.MarketplaceFeeRate); err != nil {
		return err
	}
	return nil
}

func validateMinAgentDeposit(i interface{}) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	
	if !v.IsValid() {
		return fmt.Errorf("invalid min agent deposit: %s", v)
	}
	
	if v.IsZero() {
		return fmt.Errorf("min agent deposit cannot be zero")
	}
	
	return nil
}

func validateUint64(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	
	if v == 0 {
		return fmt.Errorf("parameter cannot be zero")
	}
	
	return nil
}

func validateMarketplaceFeeRate(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	
	if v.IsNegative() {
		return fmt.Errorf("marketplace fee rate cannot be negative")
	}
	
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("marketplace fee rate cannot be greater than 1")
	}
	
	return nil
}

// Params defines the parameters for the deai module
type Params struct {
	MinAgentDeposit        sdk.Coin `json:"min_agent_deposit"`
	MaxAgentNameLength     uint64   `json:"max_agent_name_length"`
	MaxAgentDescLength     uint64   `json:"max_agent_desc_length"`
	MaxTrainingDataSize    uint64   `json:"max_training_data_size"`
	MaxMarketplaceListings uint64   `json:"max_marketplace_listings"`
	MarketplaceFeeRate     sdk.Dec  `json:"marketplace_fee_rate"`
}