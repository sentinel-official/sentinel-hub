package v3

import (
	"fmt"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	params "github.com/cosmos/cosmos-sdk/x/params/types"

	v1base "github.com/sentinel-official/hub/v12/types/v1"
)

var (
	DefaultDeposit                   = sdk.NewCoin(sdk.DefaultBondDenom, sdkmath.NewInt(10))
	DefaultActiveDuration            = 30 * time.Second
	DefaultMinGigabytePrices         = v1base.Prices{{Denom: sdk.DefaultBondDenom, BaseValue: sdkmath.LegacyMustNewDecFromStr("0.1"), QuoteValue: sdkmath.NewInt(10)}}
	DefaultMinHourlyPrices           = v1base.Prices{{Denom: sdk.DefaultBondDenom, BaseValue: sdkmath.LegacyMustNewDecFromStr("0.1"), QuoteValue: sdkmath.NewInt(10)}}
	DefaultMaxSessionGigabytes int64 = 10
	DefaultMinSessionGigabytes int64 = 1
	DefaultMaxSessionHours     int64 = 10
	DefaultMinSessionHours     int64 = 1
	DefaultStakingShare              = sdkmath.LegacyNewDecWithPrec(1, 1)
)

var (
	KeyDeposit             = []byte("Deposit")
	KeyActiveDuration      = []byte("ActiveDuration")
	KeyMinGigabytePrices   = []byte("MinGigabytePrices")
	KeyMinHourlyPrices     = []byte("MinHourlyPrices")
	KeyMaxSessionGigabytes = []byte("MaxSessionGigabytes")
	KeyMinSessionGigabytes = []byte("MinSessionGigabytes")
	KeyMaxSessionHours     = []byte("MaxSessionHours")
	KeyMinSessionHours     = []byte("MinSessionHours")
	KeyStakingShare        = []byte("StakingShare")
)

var (
	_ params.ParamSet = (*Params)(nil)
)

func (m *Params) GetMinGigabytePrices() v1base.Prices {
	return m.MinGigabytePrices
}

func (m *Params) GetMinHourlyPrices() v1base.Prices {
	return m.MinHourlyPrices
}

func (m *Params) Validate() error {
	if err := validateDeposit(m.Deposit); err != nil {
		return err
	}
	if err := validateActiveDuration(m.ActiveDuration); err != nil {
		return err
	}
	if err := validateMinGigabytePrices(m.MinGigabytePrices); err != nil {
		return err
	}
	if err := validateMinHourlyPrices(m.MinHourlyPrices); err != nil {
		return err
	}
	if err := validateMaxSessionGigabytes(m.MaxSessionGigabytes); err != nil {
		return err
	}
	if err := validateMinSessionGigabytes(m.MinSessionGigabytes); err != nil {
		return err
	}
	if err := validateMaxSessionHours(m.MaxSessionHours); err != nil {
		return err
	}
	if err := validateMinSessionHours(m.MinSessionHours); err != nil {
		return err
	}
	if err := validateStakingShare(m.StakingShare); err != nil {
		return err
	}

	return nil
}

func (m *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{
			Key:         KeyDeposit,
			Value:       &m.Deposit,
			ValidatorFn: validateDeposit,
		},
		{
			Key:         KeyActiveDuration,
			Value:       &m.ActiveDuration,
			ValidatorFn: validateActiveDuration,
		},
		{
			Key:         KeyMinGigabytePrices,
			Value:       &m.MinGigabytePrices,
			ValidatorFn: validateMinGigabytePrices,
		},
		{
			Key:         KeyMinHourlyPrices,
			Value:       &m.MinHourlyPrices,
			ValidatorFn: validateMinHourlyPrices,
		},
		{
			Key:         KeyMaxSessionGigabytes,
			Value:       &m.MaxSessionGigabytes,
			ValidatorFn: validateMaxSessionGigabytes,
		},
		{
			Key:         KeyMinSessionGigabytes,
			Value:       &m.MinSessionGigabytes,
			ValidatorFn: validateMinSessionGigabytes,
		},
		{
			Key:         KeyMaxSessionHours,
			Value:       &m.MaxSessionHours,
			ValidatorFn: validateMaxSessionHours,
		},
		{
			Key:         KeyMinSessionHours,
			Value:       &m.MinSessionHours,
			ValidatorFn: validateMinSessionHours,
		},
		{
			Key:         KeyStakingShare,
			Value:       &m.StakingShare,
			ValidatorFn: validateStakingShare,
		},
	}
}

func NewParams(
	deposit sdk.Coin, activeDuration time.Duration, minGigabytePrices, minHourlyPrices v1base.Prices,
	maxSessionGigabytes, minSessionGigabytes, maxSessionHours, minSessionHours int64, stakingShare sdkmath.LegacyDec,
) Params {
	return Params{
		Deposit:             deposit,
		ActiveDuration:      activeDuration,
		MinGigabytePrices:   minGigabytePrices,
		MinHourlyPrices:     minHourlyPrices,
		MaxSessionGigabytes: maxSessionGigabytes,
		MinSessionGigabytes: minSessionGigabytes,
		MaxSessionHours:     maxSessionHours,
		MinSessionHours:     minSessionHours,
		StakingShare:        stakingShare,
	}
}

func DefaultParams() Params {
	return NewParams(
		DefaultDeposit,
		DefaultActiveDuration,
		DefaultMinGigabytePrices,
		DefaultMinHourlyPrices,
		DefaultMaxSessionGigabytes,
		DefaultMinSessionGigabytes,
		DefaultMaxSessionHours,
		DefaultMinSessionHours,
		DefaultStakingShare,
	)
}

func ParamsKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
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

func validateActiveDuration(v interface{}) error {
	value, ok := v.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value < 0 {
		return fmt.Errorf("active_duration cannot be negative")
	}
	if value == 0 {
		return fmt.Errorf("active_duration cannot be zero")
	}

	return nil
}

func validateMinGigabytePrices(v interface{}) error {
	value, ok := v.(v1base.Prices)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value == nil {
		return nil
	}
	if !value.IsValid() {
		return fmt.Errorf("min_gigabyte_prices must be valid")
	}

	return nil
}

func validateMinHourlyPrices(v interface{}) error {
	value, ok := v.(v1base.Prices)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value == nil {
		return nil
	}
	if !value.IsValid() {
		return fmt.Errorf("min_hourly_prices must be valid")
	}

	return nil
}

func validateMaxSessionGigabytes(v interface{}) error {
	value, ok := v.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value < 0 {
		return fmt.Errorf("max_session_gigabytes cannot be negative")
	}
	if value == 0 {
		return fmt.Errorf("max_session_gigabytes cannot be zero")
	}

	return nil
}

func validateMinSessionGigabytes(v interface{}) error {
	value, ok := v.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value < 0 {
		return fmt.Errorf("min_session_gigabytes cannot be negative")
	}
	if value == 0 {
		return fmt.Errorf("min_session_gigabytes cannot be zero")
	}

	return nil
}

func validateMaxSessionHours(v interface{}) error {
	value, ok := v.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value < 0 {
		return fmt.Errorf("max_session_hours cannot be negative")
	}
	if value == 0 {
		return fmt.Errorf("max_session_hours cannot be zero")
	}

	return nil
}

func validateMinSessionHours(v interface{}) error {
	value, ok := v.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value < 0 {
		return fmt.Errorf("min_session_hours cannot be negative")
	}
	if value == 0 {
		return fmt.Errorf("min_session_hours cannot be zero")
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
