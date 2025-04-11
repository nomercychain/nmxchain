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
		MaxContractSize:     1048576,  // 1MB
		MaxContractGas:      10000000, // 10M gas
		MaxLearningDataSize: 5242880,  // 5MB
		MaxMetadataSize:     102400,   // 100KB
		MinContractDeposit:  sdk.NewCoin("unmx", sdk.NewInt(100000000)), // 100 NMX
		ExecutionFeeRate:    sdk.NewDecWithPrec(5, 2), // 5%
	}
}

// ParamSetPairs implements the ParamSet interface
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMaxContractSize, &p.MaxContractSize, validateUint64),
		paramtypes.NewParamSetPair(KeyMaxContractGas, &p.MaxContractGas, validateUint64),
		paramtypes.NewParamSetPair(KeyMaxLearningDataSize, &p.MaxLearningDataSize, validateUint64),
		paramtypes.NewParamSetPair(KeyMaxMetadataSize, &p.MaxMetadataSize, validateUint64),
		paramtypes.NewParamSetPair(KeyMinContractDeposit, &p.MinContractDeposit, validateMinContractDeposit),
		paramtypes.NewParamSetPair(KeyExecutionFeeRate, &p.ExecutionFeeRate, validateExecutionFeeRate),
	}
}

// Validate validates the parameters
func (p Params) Validate() error {
	if err := validateUint64(p.MaxContractSize); err != nil {
		return err
	}
	if err := validateUint64(p.MaxContractGas); err != nil {
		return err
	}
	if err := validateUint64(p.MaxLearningDataSize); err != nil {
		return err
	}
	if err := validateUint64(p.MaxMetadataSize); err != nil {
		return err
	}
	if err := validateMinContractDeposit(p.MinContractDeposit); err != nil {
		return err
	}
	if err := validateExecutionFeeRate(p.ExecutionFeeRate); err != nil {
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

func validateMinContractDeposit(i interface{}) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	
	if !v.IsValid() {
		return fmt.Errorf("invalid min contract deposit: %s", v)
	}
	
	if v.IsZero() {
		return fmt.Errorf("min contract deposit cannot be zero")
	}
	
	return nil
}

func validateExecutionFeeRate(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	
	if v.IsNegative() {
		return fmt.Errorf("execution fee rate cannot be negative")
	}
	
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("execution fee rate cannot be greater than 1")
	}
	
	return nil
}