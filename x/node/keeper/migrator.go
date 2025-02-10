package keeper

import (
	"time"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/node/types/v2"
	"github.com/sentinel-official/hub/v12/x/node/types/v3"
)

type Migrator struct {
	Keeper
}

func NewMigrator(k Keeper) Migrator {
	return Migrator{k}
}

func (k *Migrator) Migrate(ctx sdk.Context) error {
	prefixList := [][]byte{
		{0x11}, // NodeForInactiveAtKeyPrefix
		{0x12}, // NodeForPlanKeyPrefix
	}

	for _, buf := range prefixList {
		k.deleteKeys(ctx, buf)
	}

	k.setParams(ctx)
	k.migrateNodes(ctx)

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
	params := v3.Params{
		Deposit:             sdk.NewInt64Coin("udvpn", 0),
		ActiveDuration:      1 * time.Hour,
		MinGigabytePrices:   []v1base.Price{},
		MinHourlyPrices:     []v1base.Price{},
		MaxSessionGigabytes: 1_000_000,
		MinSessionGigabytes: 1,
		MaxSessionHours:     720,
		MinSessionHours:     1,
		StakingShare:        sdkmath.LegacyMustNewDecFromStr("0.2"),
	}

	k.SetParams(ctx, params)
}

func (k *Migrator) migrateNodes(ctx sdk.Context) {
	store := prefix.NewStore(k.Store(ctx), []byte{0x10})

	it := store.Iterator(nil, nil)
	defer it.Close()

	for ; it.Valid(); it.Next() {
		var item v2.Node
		k.cdc.MustUnmarshal(it.Value(), &item)
		store.Delete(it.Key())

		node := v3.Node{
			Address:        item.Address,
			GigabytePrices: nil,
			HourlyPrices:   nil,
			RemoteURL:      item.RemoteURL,
			InactiveAt:     item.InactiveAt,
			Status:         item.Status,
			StatusAt:       item.StatusAt,
		}

		addr, err := base.NodeAddressFromBech32(node.Address)
		if err != nil {
			panic(err)
		}

		k.SetNode(ctx, node)
		k.SetNodeForInactiveAt(ctx, node.InactiveAt, addr)
	}
}
