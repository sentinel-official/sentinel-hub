package v1

import (
	"fmt"

	sdkmath "cosmossdk.io/math"
)

var (
	DefaultMaxLeaseHours int64 = 10
	DefaultMinLeaseHours int64 = 1
	DefaultStakingShare        = sdkmath.LegacyNewDecWithPrec(1, 1)
)

func (m *Params) Validate() error {
	if err := validateMaxLeaseHours(m.MaxLeaseHours); err != nil {
		return err
	}
	if err := validateMinLeaseHours(m.MinLeaseHours); err != nil {
		return err
	}
	if err := validateStakingShare(m.StakingShare); err != nil {
		return err
	}

	return nil
}

func NewParams(maxLeaseHours, minLeaseHours int64, stakingShare sdkmath.LegacyDec) Params {
	return Params{
		MaxLeaseHours: maxLeaseHours,
		MinLeaseHours: minLeaseHours,
		StakingShare:  stakingShare,
	}
}

func DefaultParams() Params {
	return NewParams(
		DefaultMaxLeaseHours,
		DefaultMinLeaseHours,
		DefaultStakingShare,
	)
}

func validateMaxLeaseHours(v interface{}) error {
	value, ok := v.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value < 0 {
		return fmt.Errorf("max_lease_hours cannot be negative")
	}
	if value == 0 {
		return fmt.Errorf("max_lease_hours cannot be zero")
	}

	return nil
}

func validateMinLeaseHours(v interface{}) error {
	value, ok := v.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value < 0 {
		return fmt.Errorf("min_lease_hours cannot be negative")
	}
	if value == 0 {
		return fmt.Errorf("min_lease_hours cannot be zero")
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
