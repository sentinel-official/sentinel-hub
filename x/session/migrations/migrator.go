package migrations

import (
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/session/types/v2"
)

type Migrator struct {
	cdc     codec.BinaryCodec
	session SessionKeeper
}

func NewMigrator(cdc codec.BinaryCodec, session SessionKeeper) Migrator {
	return Migrator{
		cdc:     cdc,
		session: session,
	}
}

func (k *Migrator) Migrate(ctx sdk.Context) error {
	k.setParams(ctx)

	_ = k.deleteKeys(ctx, []byte{0x10}) // sessionKeys
	_ = k.deleteKeys(ctx, []byte{0x11}) // sessionForInactiveAtKeys
	_ = k.deleteKeys(ctx, []byte{0x12}) // sessionForAccountKeys
	_ = k.deleteKeys(ctx, []byte{0x13}) // sessionForNodeKeys
	_ = k.deleteKeys(ctx, []byte{0x14}) // sessionForSubscriptionKeys
	_ = k.deleteKeys(ctx, []byte{0x15}) // sessionForAllocationKeys

	return nil
}

func (k *Migrator) deleteKeys(ctx sdk.Context, keyPrefix []byte) (keys [][]byte) {
	store := prefix.NewStore(k.session.Store(ctx), keyPrefix)

	it := store.Iterator(nil, nil)
	defer it.Close()

	for ; it.Valid(); it.Next() {
		store.Delete(it.Key())

		keys = append(keys, it.Key())
	}

	return keys
}

func (k *Migrator) setParams(ctx sdk.Context) {
	params := v2.Params{
		StatusChangeDelay:        2 * time.Hour,
		ProofVerificationEnabled: false,
	}

	k.session.SetParams(ctx, params)
}
