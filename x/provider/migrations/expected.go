package migrations

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/provider/types/v2"
)

type ProviderKeeper interface {
	SetParams(ctx sdk.Context, params v2.Params)
}
