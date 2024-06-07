package cli

import (
	"encoding/base64"

	base "github.com/sentinel-official/hub/v12/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/pflag"
)

const (
	flagAddress        = "account-addr"
	flagNodeAddress    = "node-addr"
	flagRating         = "rating"
	flagSignature      = "signature"
	flagSubscriptionID = "subscription-id"
)

func GetAddress(flags *pflag.FlagSet) (sdk.AccAddress, error) {
	s, err := flags.GetString(flagAddress)
	if err != nil {
		return nil, err
	}
	if s == "" {
		return nil, nil
	}

	return sdk.AccAddressFromBech32(s)
}

func GetNodeAddress(flags *pflag.FlagSet) (base.NodeAddress, error) {
	s, err := flags.GetString(flagNodeAddress)
	if err != nil {
		return nil, err
	}
	if s == "" {
		return nil, nil
	}

	return base.NodeAddressFromBech32(s)
}

func GetSignature(flags *pflag.FlagSet) ([]byte, error) {
	s, err := flags.GetString(flagSignature)
	if err != nil {
		return nil, err
	}
	if s == "" {
		return nil, nil
	}

	return base64.StdEncoding.DecodeString(s)
}
