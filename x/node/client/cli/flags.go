package cli

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/pflag"
)

const (
	flagDenom          = "denom"
	flagGigabytePrices = "gigabyte-prices"
	flagGigabytes      = "gigabytes"
	flagHourlyPrices   = "hourly-prices"
	flagHours          = "hours"
	flagRemoteURL      = "remote-url"
)

func GetGigabytePrices(flags *pflag.FlagSet) (sdk.DecCoins, error) {
	s, err := flags.GetString(flagGigabytePrices)
	if err != nil {
		return nil, err
	}
	if s == "" {
		return nil, nil
	}

	return sdk.ParseDecCoins(s)
}

func GetHourlyPrices(flags *pflag.FlagSet) (sdk.DecCoins, error) {
	s, err := flags.GetString(flagHourlyPrices)
	if err != nil {
		return nil, err
	}
	if s == "" {
		return nil, nil
	}

	return sdk.ParseDecCoins(s)
}
