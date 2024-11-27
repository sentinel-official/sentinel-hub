package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/session/types"
	"github.com/sentinel-official/hub/v12/x/session/types/v2"
)

// SetParams stores the parameters for the module in the KVStore.
func (k *Keeper) SetParams(ctx sdk.Context, params v2.Params) {
	store := k.Store(ctx)
	key := types.ParamsKey
	value := k.cdc.MustMarshal(&params)

	store.Set(key, value)
}

// GetParams retrieves the parameters from the module's KVStore.
func (k *Keeper) GetParams(ctx sdk.Context) (v v2.Params) {
	store := k.Store(ctx)
	key := types.ParamsKey
	value := store.Get(key)

	k.cdc.MustUnmarshal(value, &v)
	return v
}

// ProofVerificationEnabled returns whether proof verification is enabled from the module's parameters.
func (k *Keeper) ProofVerificationEnabled(ctx sdk.Context) bool {
	return k.GetParams(ctx).ProofVerificationEnabled
}

// StatusChangeDelay returns the delay for status changes from the module's parameters.
func (k *Keeper) StatusChangeDelay(ctx sdk.Context) time.Duration {
	return k.GetParams(ctx).StatusChangeDelay
}

// GetInactiveAt returns the inactive time by adding StatusChangeDelay to the current block time.
func (k *Keeper) GetInactiveAt(ctx sdk.Context) time.Time {
	d := k.StatusChangeDelay(ctx)
	return ctx.BlockTime().Add(d)
}
