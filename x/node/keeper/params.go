package keeper

import (
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/node/types"
	"github.com/sentinel-official/hub/v12/x/node/types/v3"
)

// SetParams stores the node module parameters in the module's KVStore.
func (k *Keeper) SetParams(ctx sdk.Context, params v3.Params) {
	store := k.Store(ctx)
	key := types.ParamsKey
	value := k.cdc.MustMarshal(&params)

	store.Set(key, value)
}

// GetParams retrieves the node module parameters from the module's KVStore.
func (k *Keeper) GetParams(ctx sdk.Context) (v v3.Params) {
	store := k.Store(ctx)
	key := types.ParamsKey
	value := store.Get(key)

	k.cdc.MustUnmarshal(value, &v)
	return v
}

// Deposit retrieves the deposit parameter from the module's parameters.
func (k *Keeper) Deposit(ctx sdk.Context) sdk.Coin {
	return k.GetParams(ctx).Deposit
}

// ActiveDuration retrieves the active duration parameter from the module's parameters.
func (k *Keeper) ActiveDuration(ctx sdk.Context) time.Duration {
	return k.GetParams(ctx).ActiveDuration
}

// MinGigabytePrices retrieves the minimum gigabyte prices parameter from the module's parameters.
func (k *Keeper) MinGigabytePrices(ctx sdk.Context) sdk.DecCoins {
	return k.GetParams(ctx).MinGigabytePrices
}

// MinHourlyPrices retrieves the minimum hourly prices parameter from the module's parameters.
func (k *Keeper) MinHourlyPrices(ctx sdk.Context) sdk.DecCoins {
	return k.GetParams(ctx).MinHourlyPrices
}

// MaxSessionGigabytes retrieves the maximum session gigabytes parameter from the module's parameters.
func (k *Keeper) MaxSessionGigabytes(ctx sdk.Context) int64 {
	return k.GetParams(ctx).MaxSessionGigabytes
}

// MinSessionGigabytes retrieves the minimum session gigabytes parameter from the module's parameters.
func (k *Keeper) MinSessionGigabytes(ctx sdk.Context) int64 {
	return k.GetParams(ctx).MinSessionGigabytes
}

// MaxSessionHours retrieves the maximum session hours parameter from the module's parameters.
func (k *Keeper) MaxSessionHours(ctx sdk.Context) int64 {
	return k.GetParams(ctx).MaxSessionHours
}

// MinSessionHours retrieves the minimum session hours parameter from the module's parameters.
func (k *Keeper) MinSessionHours(ctx sdk.Context) int64 {
	return k.GetParams(ctx).MinSessionHours
}

// StakingShare retrieves the staking share parameter from the module's parameters.
func (k *Keeper) StakingShare(ctx sdk.Context) sdkmath.LegacyDec {
	return k.GetParams(ctx).StakingShare
}

// IsValidGigabytePrices checks if the provided gigabyte prices are valid based on the minimum prices defined in the module's parameters.
func (k *Keeper) IsValidGigabytePrices(ctx sdk.Context, prices sdk.DecCoins) bool {
	minPrices := k.MinGigabytePrices(ctx)
	for _, coin := range minPrices {
		amount := prices.AmountOf(coin.Denom)
		if amount.LT(coin.Amount) {
			return false
		}
	}

	return true
}

// IsValidHourlyPrices checks if the provided hourly prices are valid based on the minimum prices defined in the module's parameters.
func (k *Keeper) IsValidHourlyPrices(ctx sdk.Context, prices sdk.DecCoins) bool {
	minPrices := k.MinHourlyPrices(ctx)
	for _, coin := range minPrices {
		amount := prices.AmountOf(coin.Denom)
		if amount.LT(coin.Amount) {
			return false
		}
	}

	return true
}

// IsValidSessionGigabytes checks if the provided session gigabytes are within the valid range defined by the module's parameters.
func (k *Keeper) IsValidSessionGigabytes(ctx sdk.Context, gigabytes int64) bool {
	if gigabytes < k.MinSessionGigabytes(ctx) {
		return false
	}
	if gigabytes > k.MaxSessionGigabytes(ctx) {
		return false
	}

	return true
}

// IsValidSessionHours checks if the provided session hours are within the valid range defined by the module's parameters.
func (k *Keeper) IsValidSessionHours(ctx sdk.Context, hours int64) bool {
	if hours < k.MinSessionHours(ctx) {
		return false
	}
	if hours > k.MaxSessionHours(ctx) {
		return false
	}

	return true
}

// GetInactiveAt returns the inactive time by adding ActiveDuration to the current block time.
func (k *Keeper) GetInactiveAt(ctx sdk.Context) time.Time {
	d := k.ActiveDuration(ctx)
	return ctx.BlockTime().Add(d)
}
