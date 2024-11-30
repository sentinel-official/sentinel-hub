package cli

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/pflag"
)

const (
	flagPrices  = "prices"
	flagPrivate = "private"
)

func GetPrices(flags *pflag.FlagSet) (sdk.DecCoins, error) {
	s, err := flags.GetString(flagPrices)
	if err != nil {
		return nil, err
	}
	if s == "" {
		return nil, nil
	}

	return sdk.ParseDecCoins(s)
}

func GetPrivate(flags *pflag.FlagSet) (bool, error) {
	return flags.GetBool(flagPrivate)
}
