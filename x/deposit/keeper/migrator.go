package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/deposit/types/v1"
)

type Migrator struct {
	Keeper
}

func NewMigrator(k Keeper) Migrator {
	return Migrator{k}
}

func (k Migrator) Migrate(ctx sdk.Context) error {
	if err := k.refund(ctx); err != nil {
		return err
	}

	return nil
}

func (k Migrator) refund(ctx sdk.Context) error {
	k.IterateDeposits(ctx, func(_ int, item v1.Deposit) (stop bool) {
		addr, err := sdk.AccAddressFromBech32(item.Address)
		if err != nil {
			panic(err)
		}

		if err := k.SubtractDeposit(ctx, addr, item.Coins); err != nil {
			panic(err)
		}

		return false
	})

	return nil
}
