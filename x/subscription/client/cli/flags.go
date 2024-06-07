package cli

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/pflag"

	base "github.com/sentinel-official/hub/v12/types"
)

const (
	flagAccountAddress = "account-addr"
	flagNodeAddress    = "node-addr"
	flagPlanID         = "plan-id"
)

func GetAccountAddress(flags *pflag.FlagSet) (sdk.AccAddress, error) {
	s, err := flags.GetString(flagAccountAddress)
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
