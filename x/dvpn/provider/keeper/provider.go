package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/dvpn/provider/types"
)

func (k Keeper) SetProvider(ctx sdk.Context, provider types.Provider) {
	key := types.ProviderKey(provider.Address)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(provider)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) GetProvider(ctx sdk.Context, address hub.ProvAddress) (provider types.Provider, found bool) {
	store := k.Store(ctx)

	key := types.ProviderKey(address)
	value := store.Get(key)
	if value == nil {
		return provider, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &provider)
	return provider, true
}

func (k Keeper) GetProviders(ctx sdk.Context) (providers types.Providers) {
	store := k.Store(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.ProviderKeyPrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var provider types.Provider
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &provider)
		providers = append(providers, provider)
	}

	return providers
}

func (k Keeper) IterateProviders(ctx sdk.Context, f func(i int, provider types.Provider) (stop bool)) {
	store := k.Store(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.ProviderKeyPrefix)
	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		var provider types.Provider
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &provider)

		if stop := f(i, provider); stop {
			break
		}
		i++
	}
}