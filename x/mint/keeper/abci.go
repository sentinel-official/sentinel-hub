package keeper

import (
	abcitypes "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/mint/types/v1"
)

func (k *Keeper) BeginBlock(ctx sdk.Context) []abcitypes.ValidatorUpdate {
	k.IterateInflations(ctx, func(_ int, item v1.Inflation) bool {
		if item.Timestamp.After(ctx.BlockTime()) {
			return true
		}

		params := k.mint.GetParams(ctx)
		params.InflationMax = item.Max
		params.InflationMin = item.Min
		params.InflationRateChange = item.RateChange
		if err := k.mint.SetParams(ctx, params); err != nil {
			panic(err)
		}

		minter := k.mint.GetMinter(ctx)
		minter.Inflation = item.Min
		k.mint.SetMinter(ctx, minter)

		k.DeleteInflation(ctx, item.Timestamp)
		return false
	})

	return nil
}
