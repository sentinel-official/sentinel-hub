package v2

import (
	"fmt"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	DefaultDeposit      = sdk.NewCoin(sdk.DefaultBondDenom, sdkmath.NewInt(1000))
	DefaultStakingShare = sdkmath.LegacyNewDecWithPrec(1, 1)
)

func (m *Params) Validate() error {
	if err := validateDeposit(m.Deposit); err != nil {
		return err
	}
	if err := validateStakingShare(m.StakingShare); err != nil {
		return err
	}

	return nil
}

func NewParams(deposit sdk.Coin, stakingShare sdkmath.LegacyDec) Params {
	return Params{
		Deposit:      deposit,
		StakingShare: stakingShare,
	}
}

func DefaultParams() Params {
	return Params{
		Deposit:      DefaultDeposit,
		StakingShare: DefaultStakingShare,
	}
}

func validateDeposit(v interface{}) error {
	value, ok := v.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value.IsNil() {
		return fmt.Errorf("deposit cannot be nil")
	}
	if value.IsNegative() {
		return fmt.Errorf("deposit cannot be negative")
	}
	if !value.IsValid() {
		return fmt.Errorf("invalid deposit %s", value)
	}

	return nil
}

func validateStakingShare(v interface{}) error {
	value, ok := v.(sdkmath.LegacyDec)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value.IsNil() {
		return fmt.Errorf("staking_share cannot be nil")
	}
	if value.IsNegative() {
		return fmt.Errorf("staking_share cannot be negative")
	}
	if value.GT(sdkmath.LegacyOneDec()) {
		return fmt.Errorf("staking_share cannot be greater than 1")
	}

	return nil
}
