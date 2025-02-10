package keeper

import (
	"time"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/subscription/types/v2"
)

type Migrator struct {
	Keeper
}

func NewMigrator(k Keeper) Migrator {
	return Migrator{k}
}

func (k *Migrator) Migrate(ctx sdk.Context) error {
	prefixList := [][]byte{
		{0x10}, // SubscriptionKeyPrefix
		{0x20}, // AllocationKeyPrefix
		{0x30}, // PayoutKeyPrefix

		{0x11}, // SubscriptionForInactiveAtKeyPrefix
		{0x12}, // SubscriptionForAccountKeyPrefix
		{0x13}, // SubscriptionForNodeKeyPrefix
		{0x14}, // SubscriptionForPlanKeyPrefix

		{0x31}, // PayoutForNextAtKeyPrefix
		{0x32}, // PayoutForAccountKeyPrefix
		{0x33}, // PayoutForNodeKeyPrefix
		{0x34}, // PayoutForAccountByNodeKeyPrefix
	}

	for _, buf := range prefixList {
		k.deleteKeys(ctx, buf)
	}

	k.setParams(ctx)

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

func (k *Migrator) setParams(ctx sdk.Context) {
	params := v2.Params{
		StatusChangeDelay: 4 * time.Hour,
	}

	k.SetParams(ctx, params)
}
