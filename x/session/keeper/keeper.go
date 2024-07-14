package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/session/expected"
	"github.com/sentinel-official/hub/v12/x/session/types"
)

type Keeper struct {
	authority        string
	feeCollectorName string
	cdc              codec.BinaryCodec
	key              storetypes.StoreKey
	account          expected.AccountKeeper
	deposit          expected.DepositKeeper
	node             expected.NodeKeeper
	subscription     expected.SubscriptionKeeper
}

func NewKeeper(cdc codec.BinaryCodec, key storetypes.StoreKey, authority, feeCollectorName string) Keeper {
	return Keeper{
		authority:        authority,
		feeCollectorName: feeCollectorName,
		cdc:              cdc,
		key:              key,
	}
}

func (k *Keeper) WithAccountKeeper(keeper expected.AccountKeeper) {
	k.account = keeper
}

func (k *Keeper) WithDepositKeeper(keeper expected.DepositKeeper) {
	k.deposit = keeper
}

func (k *Keeper) WithNodeKeeper(keeper expected.NodeKeeper) {
	k.node = keeper
}

func (k *Keeper) WithSubscriptionKeeper(keeper expected.SubscriptionKeeper) {
	k.subscription = keeper
}

func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

func (k *Keeper) Store(ctx sdk.Context) sdk.KVStore {
	child := fmt.Sprintf("%s/", types.ModuleName)
	return prefix.NewStore(ctx.KVStore(k.key), []byte(child))
}
