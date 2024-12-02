package keeper

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ibchost "github.com/cosmos/ibc-go/v7/modules/core/24-host"
	protobuf "github.com/gogo/protobuf/types"

	"github.com/sentinel-official/hub/v12/x/oracle/types"
	"github.com/sentinel-official/hub/v12/x/oracle/types/v1"
)

// SetAsset stores an asset in the module's KVStore.
func (k *Keeper) SetAsset(ctx sdk.Context, asset v1.Asset) {
	store := k.Store(ctx)
	key := types.AssetKey(asset.Denom)
	value := k.cdc.MustMarshal(&asset)

	store.Set(key, value)
}

// GetAsset retrieves an asset from the module's KVStore based on the asset denomination.
// If the asset exists, it returns the asset and 'found' as true; otherwise, it returns 'found' as false.
func (k *Keeper) GetAsset(ctx sdk.Context, denom string) (v v1.Asset, found bool) {
	store := k.Store(ctx)
	key := types.AssetKey(denom)
	value := store.Get(key)

	if value == nil {
		return v, false
	}

	k.cdc.MustUnmarshal(value, &v)
	return v, true
}

// HasAsset checks if an asset with the given denomination exists in the module's KVStore.
func (k *Keeper) HasAsset(ctx sdk.Context, denom string) bool {
	store := k.Store(ctx)
	key := types.AssetKey(denom)

	return store.Has(key)
}

// DeleteAsset removes an asset from the module's KVStore based on the asset denomination.
func (k *Keeper) DeleteAsset(ctx sdk.Context, denom string) {
	store := k.Store(ctx)
	key := types.AssetKey(denom)

	store.Delete(key)
}

// IterateAssets iterates over all assets stored in the module's KVStore and calls the provided function for each asset.
// The iteration stops when the provided function returns 'true'.
func (k *Keeper) IterateAssets(ctx sdk.Context, fn func(int, v1.Asset) bool) {
	store := k.Store(ctx)
	iterator := sdk.KVStorePrefixIterator(store, types.AssetKeyPrefix)

	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		var item v1.Asset
		k.cdc.MustUnmarshal(iterator.Value(), &item)

		if stop := fn(i, item); stop {
			break
		}
		i++
	}
}

// SetDenomForPacket stores a denomination for a packet in the module's KVStore.
func (k *Keeper) SetDenomForPacket(ctx sdk.Context, portID, channelID string, sequence uint64, denom string) {
	store := k.Store(ctx)
	key := ibchost.PacketCommitmentKey(portID, channelID, sequence)
	value := k.cdc.MustMarshal(&protobuf.StringValue{Value: denom})

	store.Set(key, value)
}

// GetDenomForPacket retrieves a denomination for a packet from the module's KVStore based on the port ID, channel ID, and sequence.
// If the denomination exists, it returns the denomination and 'found' as true; otherwise, it returns 'found' as false.
func (k *Keeper) GetDenomForPacket(ctx sdk.Context, portID, channelID string, sequence uint64) (v string, found bool) {
	store := k.Store(ctx)
	key := ibchost.PacketCommitmentKey(portID, channelID, sequence)
	value := store.Get(key)

	if value == nil {
		return v, false
	}

	var denom protobuf.StringValue
	k.cdc.MustUnmarshal(value, &denom)

	return denom.GetValue(), true
}

// DeleteDenomForPacket removes a denomination for a packet from the module's KVStore based on the port ID, channel ID, and sequence.
func (k *Keeper) DeleteDenomForPacket(ctx sdk.Context, portID, channelID string, sequence uint64) {
	store := k.Store(ctx)
	key := ibchost.PacketCommitmentKey(portID, channelID, sequence)

	store.Delete(key)
}

// GetAssetForPacket retrieves an asset for a packet from the module's KVStore based on the port ID, channel ID, and sequence.
// If the asset exists, it returns the asset; otherwise, it returns an error.
func (k *Keeper) GetAssetForPacket(ctx sdk.Context, portID, channelID string, sequence uint64) (v v1.Asset, err error) {
	denom, found := k.GetDenomForPacket(ctx, portID, channelID, sequence)
	if !found {
		return v, types.NewErrorDenomForPacketNotFound(portID, channelID, sequence)
	}

	v, found = k.GetAsset(ctx, denom)
	if !found {
		return v, types.NewErrorAssetNotFound(denom)
	}

	return v, nil
}

// GetQuote returns the quote for a given base coin.
func (k *Keeper) GetQuote(ctx sdk.Context, coin sdk.DecCoin) (sdk.Coin, error) {
	if coin.Denom == "" {
		return sdk.Coin{Amount: sdkmath.ZeroInt()}, nil
	}

	asset, found := k.GetAsset(ctx, coin.Denom)
	if !found {
		return sdk.Coin{Amount: sdkmath.ZeroInt()}, types.NewErrorAssetNotFound(coin.Denom)
	}

	amount := coin.Amount.MulInt(asset.Price).TruncateInt()
	return sdk.Coin{Denom: coin.Denom, Amount: amount}, nil
}
