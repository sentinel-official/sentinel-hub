package migrations

import (
	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/provider/types/v2"
)

type Migrator struct {
	cdc      codec.BinaryCodec
	provider ProviderKeeper
}

func NewMigrator(cdc codec.BinaryCodec, provider ProviderKeeper) Migrator {
	return Migrator{
		cdc:      cdc,
		provider: provider,
	}
}

func (k *Migrator) Migrate(ctx sdk.Context) error {
	k.setParams(ctx)

	return nil
}

func (k *Migrator) setParams(ctx sdk.Context) {
	params := v2.Params{
		Deposit:      sdk.NewInt64Coin("udvpn", 0),
		StakingShare: math.LegacyMustNewDecFromStr("0.2"),
	}

	k.provider.SetParams(ctx, params)
}
