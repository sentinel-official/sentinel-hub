package v2

import (
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
)

func TestParams_Validate(t *testing.T) {
	type fields struct {
		Deposit                  sdk.Coin
		ActiveDuration           time.Duration
		MaxGigabytePrices        sdk.Coins
		MinGigabytePrices        sdk.Coins
		MaxHourlyPrices          sdk.Coins
		MinHourlyPrices          sdk.Coins
		MaxSubscriptionGigabytes int64
		MinSubscriptionGigabytes int64
		MaxSubscriptionHours     int64
		MinSubscriptionHours     int64
		StakingShare             sdkmath.LegacyDec
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"deposit empty",
			fields{
				Deposit: base.TestCoinEmptyZero,
			},
			true,
		},
		{
			"deposit empty denom",
			fields{
				Deposit: base.TestCoinEmptyPos,
			},
			true,
		},
		{
			"deposit invalid denom",
			fields{
				Deposit: base.TestCoinInvalidPos,
			},
			true,
		},
		{
			"deposit empty amount",
			fields{
				Deposit: base.TestCoinOneEmpty,
			},
			true,
		},
		{
			"deposit negative amount",
			fields{
				Deposit: base.TestCoinOneNeg,
			},
			true,
		},
		{
			"deposit zero amount",
			fields{
				Deposit: base.TestCoinOneZero,
			},
			true,
		},
		{
			"deposit positive amount",
			fields{
				Deposit:                  base.TestCoinOnePos,
				ActiveDuration:           1000,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"active_duration negative",
			fields{
				Deposit:        base.TestCoinOnePos,
				ActiveDuration: -1000,
			},
			true,
		},
		{
			"active_duration zero",
			fields{
				Deposit:        base.TestCoinOnePos,
				ActiveDuration: 0,
			},
			true,
		},
		{
			"active_duration positive",
			fields{
				Deposit:                  base.TestCoinOnePos,
				ActiveDuration:           1000,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"max_gigabyte_prices nil",
			fields{
				Deposit:                  base.TestCoinOnePos,
				ActiveDuration:           1000,
				MaxGigabytePrices:        nil,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"max_gigabyte_prices empty",
			fields{
				Deposit:                  base.TestCoinOnePos,
				ActiveDuration:           1000,
				MaxGigabytePrices:        base.TestCoinsEmpty,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"max_gigabyte_prices empty denom",
			fields{
				Deposit:           base.TestCoinOnePos,
				ActiveDuration:    1000,
				MaxGigabytePrices: base.TestCoinsEmptyPos,
			},
			true,
		},
		{
			"max_gigabyte_prices invalid denom",
			fields{
				Deposit:           base.TestCoinOnePos,
				ActiveDuration:    1000,
				MaxGigabytePrices: base.TestCoinsInvalidPos,
			},
			true,
		},
		{
			"max_gigabyte_prices empty amount",
			fields{
				Deposit:           base.TestCoinOnePos,
				ActiveDuration:    1000,
				MaxGigabytePrices: base.TestCoinsOneEmpty,
			},
			true,
		},
		{
			"max_gigabyte_prices negative amount",
			fields{
				Deposit:           base.TestCoinOnePos,
				ActiveDuration:    1000,
				MaxGigabytePrices: base.TestCoinsOneNeg,
			},
			true,
		},
		{
			"max_gigabyte_prices zero amount",
			fields{
				Deposit:           base.TestCoinOnePos,
				ActiveDuration:    1000,
				MaxGigabytePrices: base.TestCoinsOneZero,
			},
			true,
		},
		{
			"max_gigabyte_prices positive amount",
			fields{
				Deposit:                  base.TestCoinOnePos,
				ActiveDuration:           1000,
				MaxGigabytePrices:        base.TestCoinsOnePos,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"min_gigabyte_prices nil",
			fields{
				Deposit:                  base.TestCoinOnePos,
				ActiveDuration:           1000,
				MinGigabytePrices:        nil,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"min_gigabyte_prices empty",
			fields{
				Deposit:                  base.TestCoinOnePos,
				ActiveDuration:           1000,
				MinGigabytePrices:        base.TestCoinsEmpty,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"min_gigabyte_prices empty denom",
			fields{
				Deposit:           base.TestCoinOnePos,
				ActiveDuration:    1000,
				MinGigabytePrices: base.TestCoinsEmptyPos,
			},
			true,
		},
		{
			"min_gigabyte_prices invalid denom",
			fields{
				Deposit:           base.TestCoinOnePos,
				ActiveDuration:    1000,
				MinGigabytePrices: base.TestCoinsInvalidPos,
			},
			true,
		},
		{
			"min_gigabyte_prices empty amount",
			fields{
				Deposit:           base.TestCoinOnePos,
				ActiveDuration:    1000,
				MinGigabytePrices: base.TestCoinsOneEmpty,
			},
			true,
		},
		{
			"min_gigabyte_prices negative amount",
			fields{
				Deposit:           base.TestCoinOnePos,
				ActiveDuration:    1000,
				MinGigabytePrices: base.TestCoinsOneNeg,
			},
			true,
		},
		{
			"min_gigabyte_prices zero amount",
			fields{
				Deposit:           base.TestCoinOnePos,
				ActiveDuration:    1000,
				MinGigabytePrices: base.TestCoinsOneZero,
			},
			true,
		},
		{
			"min_gigabyte_prices positive amount",
			fields{
				Deposit:                  base.TestCoinOnePos,
				ActiveDuration:           1000,
				MinGigabytePrices:        base.TestCoinsOnePos,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"max_hourly_prices nil",
			fields{
				Deposit:                  base.TestCoinOnePos,
				ActiveDuration:           1000,
				MaxHourlyPrices:          nil,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"max_hourly_prices empty",
			fields{
				Deposit:                  base.TestCoinOnePos,
				ActiveDuration:           1000,
				MaxHourlyPrices:          base.TestCoinsEmpty,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"max_hourly_prices empty denom",
			fields{
				Deposit:         base.TestCoinOnePos,
				ActiveDuration:  1000,
				MaxHourlyPrices: base.TestCoinsEmptyPos,
			},
			true,
		},
		{
			"max_hourly_prices invalid denom",
			fields{
				Deposit:         base.TestCoinOnePos,
				ActiveDuration:  1000,
				MaxHourlyPrices: base.TestCoinsInvalidPos,
			},
			true,
		},
		{
			"max_hourly_prices empty amount",
			fields{
				Deposit:         base.TestCoinOnePos,
				ActiveDuration:  1000,
				MaxHourlyPrices: base.TestCoinsOneEmpty,
			},
			true,
		},
		{
			"max_hourly_prices negative amount",
			fields{
				Deposit:         base.TestCoinOnePos,
				ActiveDuration:  1000,
				MaxHourlyPrices: base.TestCoinsOneNeg,
			},
			true,
		},
		{
			"max_hourly_prices zero amount",
			fields{
				Deposit:         base.TestCoinOnePos,
				ActiveDuration:  1000,
				MaxHourlyPrices: base.TestCoinsOneZero,
			},
			true,
		},
		{
			"max_hourly_prices positive amount",
			fields{
				Deposit:                  base.TestCoinOnePos,
				ActiveDuration:           1000,
				MaxHourlyPrices:          base.TestCoinsOnePos,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"min_hourly_prices nil",
			fields{
				Deposit:                  base.TestCoinOnePos,
				ActiveDuration:           1000,
				MinHourlyPrices:          nil,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"min_hourly_prices empty",
			fields{
				Deposit:                  base.TestCoinOnePos,
				ActiveDuration:           1000,
				MinHourlyPrices:          base.TestCoinsEmpty,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"min_hourly_prices empty denom",
			fields{
				Deposit:         base.TestCoinOnePos,
				ActiveDuration:  1000,
				MinHourlyPrices: base.TestCoinsEmptyPos,
			},
			true,
		},
		{
			"min_hourly_prices invalid denom",
			fields{
				Deposit:         base.TestCoinOnePos,
				ActiveDuration:  1000,
				MinHourlyPrices: base.TestCoinsInvalidPos,
			},
			true,
		},
		{
			"min_hourly_prices empty amount",
			fields{
				Deposit:         base.TestCoinOnePos,
				ActiveDuration:  1000,
				MinHourlyPrices: base.TestCoinsOneEmpty,
			},
			true,
		},
		{
			"min_hourly_prices negative amount",
			fields{
				Deposit:         base.TestCoinOnePos,
				ActiveDuration:  1000,
				MinHourlyPrices: base.TestCoinsOneNeg,
			},
			true,
		},
		{
			"min_hourly_prices zero amount",
			fields{
				Deposit:         base.TestCoinOnePos,
				ActiveDuration:  1000,
				MinHourlyPrices: base.TestCoinsOneZero,
			},
			true,
		},
		{
			"min_hourly_prices positive amount",
			fields{
				Deposit:                  base.TestCoinOnePos,
				ActiveDuration:           1000,
				MinHourlyPrices:          base.TestCoinsOnePos,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"max_subscription_gigabytes negative",
			fields{
				Deposit:                  base.TestCoinOnePos,
				ActiveDuration:           1000,
				MaxSubscriptionGigabytes: -1000,
			},
			true,
		},
		{
			"max_subscription_gigabytes zero",
			fields{
				Deposit:                  base.TestCoinOnePos,
				ActiveDuration:           1000,
				MaxSubscriptionGigabytes: 0,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			true,
		},
		{
			"max_subscription_gigabytes positive",
			fields{
				Deposit:                  base.TestCoinOnePos,
				ActiveDuration:           1000,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"min_subscription_gigabytes negative",
			fields{
				Deposit:                  base.TestCoinOnePos,
				ActiveDuration:           1000,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: -1000,
			},
			true,
		},
		{
			"min_subscription_gigabytes zero",
			fields{
				Deposit:                  base.TestCoinOnePos,
				ActiveDuration:           1000,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 0,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			true,
		},
		{
			"min_subscription_gigabytes positive",
			fields{
				Deposit:                  base.TestCoinOnePos,
				ActiveDuration:           1000,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"max_subscription_hours negative",
			fields{
				Deposit:                  base.TestCoinOnePos,
				ActiveDuration:           1000,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     -1000,
			},
			true,
		},
		{
			"max_subscription_hours zero",
			fields{
				Deposit:                  base.TestCoinOnePos,
				ActiveDuration:           1000,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     0,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			true,
		},
		{
			"max_subscription_hours positive",
			fields{
				Deposit:                  base.TestCoinOnePos,
				ActiveDuration:           1000,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"min_subscription_hours negative",
			fields{
				Deposit:                  base.TestCoinOnePos,
				ActiveDuration:           1000,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     -1000,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			true,
		},
		{
			"min_subscription_hours zero",
			fields{
				Deposit:                  base.TestCoinOnePos,
				ActiveDuration:           1000,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     0,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			true,
		},
		{
			"min_subscription_hours positive",
			fields{
				Deposit:                  base.TestCoinOnePos,
				ActiveDuration:           1000,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"staking_share empty",
			fields{
				Deposit:        base.TestCoinOnePos,
				ActiveDuration: 1000,
				StakingShare:   sdkmath.LegacyDec{},
			},
			true,
		},
		{
			"staking_share -10",
			fields{
				Deposit:        base.TestCoinOnePos,
				ActiveDuration: 1000,
				StakingShare:   sdkmath.LegacyNewDecWithPrec(-10, 0),
			},
			true,
		},
		{
			"staking_share -1",
			fields{
				Deposit:        base.TestCoinOnePos,
				ActiveDuration: 1000,
				StakingShare:   sdkmath.LegacyNewDecWithPrec(-1, 0),
			},
			true,
		},
		{
			"staking_share -0.5",
			fields{
				Deposit:        base.TestCoinOnePos,
				ActiveDuration: 1000,
				StakingShare:   sdkmath.LegacyNewDecWithPrec(-5, 1),
			},
			true,
		},
		{
			"staking_share 0",
			fields{
				Deposit:                  base.TestCoinOnePos,
				ActiveDuration:           1000,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(0, 0),
			},
			false,
		},
		{
			"staking_share 0.5",
			fields{
				Deposit:                  base.TestCoinOnePos,
				ActiveDuration:           1000,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(5, 1),
			},
			false,
		},
		{
			"staking_share 1",
			fields{
				Deposit:                  base.TestCoinOnePos,
				ActiveDuration:           1000,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"staking_share 10",
			fields{
				Deposit:                  base.TestCoinOnePos,
				ActiveDuration:           1000,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(10, 0),
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Params{
				Deposit:                  tt.fields.Deposit,
				ActiveDuration:           tt.fields.ActiveDuration,
				MaxGigabytePrices:        tt.fields.MaxGigabytePrices,
				MinGigabytePrices:        tt.fields.MinGigabytePrices,
				MaxHourlyPrices:          tt.fields.MaxHourlyPrices,
				MinHourlyPrices:          tt.fields.MinHourlyPrices,
				MaxSubscriptionGigabytes: tt.fields.MaxSubscriptionGigabytes,
				MinSubscriptionGigabytes: tt.fields.MinSubscriptionGigabytes,
				MaxSubscriptionHours:     tt.fields.MaxSubscriptionHours,
				MinSubscriptionHours:     tt.fields.MinSubscriptionHours,
				StakingShare:             tt.fields.StakingShare,
			}
			if err := m.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
