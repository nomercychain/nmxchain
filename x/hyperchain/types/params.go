package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ParamKeyTable returns the parameter key table
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// DefaultParams returns default parameters
func DefaultParams() Params {
	return Params{
		MaxHyperchainsPerAccount:     5,
		MaxValidatorsPerHyperchain:   100,
		MaxBridgesPerHyperchain:      10,
		MinHyperchainCreationDeposit: sdk.NewCoin("unmx", sdk.NewInt(1000000000)), // 1000 NMX
		MinValidatorStake:            sdk.NewCoin("unmx", sdk.NewInt(100000000)),  // 100 NMX
		BridgeFeeRate:                sdk.NewDecWithPrec(1, 2), // 1%
	}
}

// ParamSetPairs implements the ParamSet interface
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMaxHyperchainsPerAccount, &p.MaxHyperchainsPerAccount, validateUint64),
		paramtypes.NewParamSetPair(KeyMaxValidatorsPerHyperchain, &p.MaxValidatorsPerHyperchain, validateUint64),
		paramtypes.NewParamSetPair(KeyMaxBridgesPerHyperchain, &p.MaxBridgesPerHyperchain, validateUint64),
		paramtypes.NewParamSetPair(KeyMinHyperchainCreationDeposit, &p.MinHyperchainCreationDeposit, validateCoin),
		paramtypes.NewParamSetPair(KeyMinValidatorStake, &p.MinValidatorStake, validateCoin),
		paramtypes.NewParamSetPair(KeyBridgeFeeRate, &p.BridgeFeeRate, validateBridgeFeeRate),
	}
}

// Validate validates the parameters
func (p Params) Validate() error {
	if err := validateUint64(p.MaxHyperchainsPerAccount); err != nil {
		return err
	}
	if err := validateUint64(p.MaxValidatorsPerHyperchain); err != nil {
		return err
	}
	if err := validateUint64(p.MaxBridgesPerHyperchain); err != nil {
		return err
	}
	if err := validateCoin(p.MinHyperchainCreationDeposit); err != nil {
		return err
	}
	if err := validateCoin(p.MinValidatorStake); err != nil {
		return err
	}
	if err := validateBridgeFeeRate(p.BridgeFeeRate); err != nil {
		return err
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

func validateCoin(i interface{}) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	
	if !v.IsValid() {
		return fmt.Errorf("invalid coin: %s", v)
	}
	
	if v.IsZero() {
		return fmt.Errorf("coin amount cannot be zero")
	}
	
	return nil
}

func validateBridgeFeeRate(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	
	if v.IsNegative() {
		return fmt.Errorf("bridge fee rate cannot be negative")
	}
	
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("bridge fee rate cannot be greater than 1")
	}
	
	return nil
}