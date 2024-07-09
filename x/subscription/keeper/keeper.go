package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/subscription/expected"
	"github.com/sentinel-official/hub/v12/x/subscription/types"
)

type Keeper struct {
	authority        string
	feeCollectorName string
	cdc              codec.BinaryCodec
	key              storetypes.StoreKey
	bank             expected.BankKeeper
	node             expected.NodeKeeper
	plan             expected.PlanKeeper
	provider         expected.ProviderKeeper
	session          expected.SessionKeeper
}

func NewKeeper(cdc codec.BinaryCodec, key storetypes.StoreKey, authority, feeCollectorName string) Keeper {
	return Keeper{
		authority:        authority,
		feeCollectorName: feeCollectorName,
		cdc:              cdc,
		key:              key,
	}
}

func (k *Keeper) WithBankKeeper(keeper expected.BankKeeper) {
	k.bank = keeper
}

func (k *Keeper) WithProviderKeeper(keeper expected.ProviderKeeper) {
	k.provider = keeper
}

func (k *Keeper) WithNodeKeeper(keeper expected.NodeKeeper) {
	k.node = keeper
}

func (k *Keeper) WithPlanKeeper(keeper expected.PlanKeeper) {
	k.plan = keeper
}

func (k *Keeper) WithSessionKeeper(keeper expected.SessionKeeper) {
	k.session = keeper
}

func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

func (k *Keeper) Store(ctx sdk.Context) sdk.KVStore {
	child := fmt.Sprintf("%s/", types.ModuleName)
	return prefix.NewStore(ctx.KVStore(k.key), []byte(child))
}
