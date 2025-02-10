package keeper

import (
	"time"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/plan/types/v2"
	"github.com/sentinel-official/hub/v12/x/plan/types/v3"
)

type Migrator struct {
	Keeper
}

func NewMigrator(k Keeper) Migrator {
	return Migrator{k}
}

func (k *Migrator) Migrate(ctx sdk.Context) error {
	prefixList := [][]byte{
		{0x11}, // PlanForProviderKeyPrefix
	}

	for _, buf := range prefixList {
		k.deleteKeys(ctx, buf)
	}

	k.migratePlans(ctx)

	return nil
}

func (k *Migrator) deleteKeys(ctx sdk.Context, keyPrefix []byte) {
	store := prefix.NewStore(k.Store(ctx), keyPrefix)

	it := store.Iterator(nil, nil)
	defer it.Close()

	for ; it.Valid(); it.Next() {
		store.Delete(it.Key())
	}
}

func (k *Migrator) migratePlans(ctx sdk.Context) {
	store := prefix.NewStore(k.Store(ctx), []byte{0x10})

	it := store.Iterator(nil, nil)
	defer it.Close()

	for ; it.Valid(); it.Next() {
		var item v2.Plan
		k.cdc.MustUnmarshal(it.Value(), &item)
		store.Delete(it.Key())

		plan := v3.Plan{
			ID:          item.ID,
			ProvAddress: item.ProviderAddress,
			Gigabytes:   item.Gigabytes,
			Hours:       int64(item.Duration / time.Hour),
			Prices:      []v1base.Price{},
			Status:      item.Status,
			StatusAt:    item.StatusAt,
		}

		addr, err := base.ProvAddressFromBech32(plan.ProvAddress)
		if err != nil {
			panic(err)
		}

		k.SetPlan(ctx, plan)
		k.SetPlanForProvider(ctx, addr, plan.ID)
	}
}
