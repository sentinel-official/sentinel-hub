package keeper

import (
	"time"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/session/types/v2"
)

type Migrator struct {
	Keeper
}

func NewMigrator(k Keeper) Migrator {
	return Migrator{k}
}

func (k *Migrator) Migrate(ctx sdk.Context) error {
	prefixList := [][]byte{
		{0x10}, // SessionKeyPrefix

		{0x11}, // SessionForInactiveAtKeyPrefix
		{0x12}, // SessionForAccountKeyPrefix
		{0x13}, // SessionForNodeKeyPrefix
		{0x14}, // SessionForSubscriptionKeyPrefix
		{0x15}, // SessionForAllocationKeyPrefix
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
		StatusChangeDelay:        2 * time.Hour,
		ProofVerificationEnabled: false,
	}

	k.SetParams(ctx, params)
}
