package keeper

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/provider/types/v2"
)

type Migrator struct {
	Keeper
}

func NewMigrator(k Keeper) Migrator {
	return Migrator{k}
}

func (k *Migrator) Migrate(ctx sdk.Context) error {
	k.setParams(ctx)

	return nil
}

func (k *Migrator) setParams(ctx sdk.Context) {
	params := v2.Params{
		Deposit:      sdk.NewInt64Coin("udvpn", 0),
		StakingShare: sdkmath.LegacyMustNewDecFromStr("0.2"),
	}

	k.SetParams(ctx, params)
}
