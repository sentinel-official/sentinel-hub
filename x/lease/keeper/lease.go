package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	protobuf "github.com/gogo/protobuf/types"

	base "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/lease/types"
	"github.com/sentinel-official/hub/v12/x/lease/types/v1"
)

// SetLease stores a lease in the module's KVStore.
func (k *Keeper) SetLease(ctx sdk.Context, lease v1.Lease) {
	store := k.Store(ctx)
	key := types.LeaseKey(lease.ID)
	value := k.cdc.MustMarshal(&lease)

	store.Set(key, value)
}

// HasLease checks if a lease exists in the module's KVStore based on the lease ID.
func (k *Keeper) HasLease(ctx sdk.Context, id uint64) bool {
	store := k.Store(ctx)
	key := types.LeaseKey(id)

	return store.Has(key)
}

// GetLease retrieves a lease from the module's KVStore based on the lease ID.
// If the lease exists, it returns the lease and 'found' as true; otherwise, it returns 'found' as false.
func (k *Keeper) GetLease(ctx sdk.Context, id uint64) (lease v1.Lease, found bool) {
	store := k.Store(ctx)
	key := types.LeaseKey(id)
	value := store.Get(key)

	if value == nil {
		return lease, false
	}

	k.cdc.MustUnmarshal(value, &lease)
	return lease, true
}

// DeleteLease removes a lease from the module's KVStore based on the lease ID.
func (k *Keeper) DeleteLease(ctx sdk.Context, id uint64) {
	store := k.Store(ctx)
	key := types.LeaseKey(id)

	store.Delete(key)
}

// GetLeases retrieves all leases stored in the module's KVStore.
func (k *Keeper) GetLeases(ctx sdk.Context) (items []v1.Lease) {
	store := k.Store(ctx)
	iterator := sdk.KVStorePrefixIterator(store, types.LeaseKeyPrefix)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var item v1.Lease
		k.cdc.MustUnmarshal(iterator.Value(), &item)

		items = append(items, item)
	}

	return items
}

// IterateLeases iterates over all leases stored in the module's KVStore and calls the provided function for each lease.
// The iteration stops when the provided function returns 'true'.
func (k *Keeper) IterateLeases(ctx sdk.Context, fn func(index int, item v1.Lease) (stop bool)) {
	store := k.Store(ctx)
	iterator := sdk.KVStorePrefixIterator(store, types.LeaseKeyPrefix)

	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		var item v1.Lease
		k.cdc.MustUnmarshal(iterator.Value(), &item)

		if stop := fn(i, item); stop {
			break
		}
		i++
	}
}

// SetLeaseForNode stores a lease for a node in the module's KVStore.
func (k *Keeper) SetLeaseForNode(ctx sdk.Context, addr base.NodeAddress, id uint64) {
	store := k.Store(ctx)
	key := types.LeaseForNodeKey(addr, id)
	value := k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})

	store.Set(key, value)
}

// HasLeaseForNode checks if a lease for a node exists in the module's KVStore based on the node address and lease ID.
func (k *Keeper) HasLeaseForNode(ctx sdk.Context, addr base.NodeAddress, id uint64) bool {
	store := k.Store(ctx)
	key := types.LeaseForNodeKey(addr, id)

	return store.Has(key)
}

// DeleteLeaseForNode removes a lease for a node from the module's KVStore based on the node address and lease ID.
func (k *Keeper) DeleteLeaseForNode(ctx sdk.Context, addr base.NodeAddress, id uint64) {
	store := k.Store(ctx)
	key := types.LeaseForNodeKey(addr, id)

	store.Delete(key)
}

// IterateLeasesForNode iterates over all leases for a specific node and calls the provided function for each lease.
// The iteration stops when the provided function returns 'true'.
func (k *Keeper) IterateLeasesForNode(ctx sdk.Context, addr base.NodeAddress, fn func(index int, item v1.Lease) (stop bool)) {
	store := k.Store(ctx)
	iterator := store.Iterator(types.LeaseForNodeKeyPrefix, sdk.PrefixEndBytes(types.GetLeaseForNodeKeyPrefix(addr)))

	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		item, found := k.GetLease(ctx, types.IDFromLeaseForNodeKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("lease for node key %X does not exist", iterator.Key()))
		}

		if stop := fn(i, item); stop {
			break
		}
		i++
	}
}

// SetLeaseForProvider stores a lease for a provider in the module's KVStore.
func (k *Keeper) SetLeaseForProvider(ctx sdk.Context, addr base.ProvAddress, id uint64) {
	store := k.Store(ctx)
	key := types.LeaseForProviderKey(addr, id)
	value := k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})

	store.Set(key, value)
}

// HasLeaseForProvider checks if a lease for a provider exists in the module's KVStore based on the provider address and lease ID.
func (k *Keeper) HasLeaseForProvider(ctx sdk.Context, addr base.ProvAddress, id uint64) bool {
	store := k.Store(ctx)
	key := types.LeaseForProviderKey(addr, id)

	return store.Has(key)
}

// DeleteLeaseForProvider removes a lease for a provider from the module's KVStore based on the provider address and lease ID.
func (k *Keeper) DeleteLeaseForProvider(ctx sdk.Context, addr base.ProvAddress, id uint64) {
	store := k.Store(ctx)
	key := types.LeaseForProviderKey(addr, id)

	store.Delete(key)
}

// IterateLeasesForProvider iterates over all leases for a specific provider and calls the provided function for each lease.
// The iteration stops when the provided function returns 'true'.
func (k *Keeper) IterateLeasesForProvider(ctx sdk.Context, addr base.ProvAddress, fn func(index int, item v1.Lease) (stop bool)) {
	store := k.Store(ctx)
	iterator := store.Iterator(types.LeaseForProviderKeyPrefix, sdk.PrefixEndBytes(types.GetLeaseForProviderKeyPrefix(addr)))

	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		item, found := k.GetLease(ctx, types.IDFromLeaseForProviderKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("lease for provider key %X does not exist", iterator.Key()))
		}

		if stop := fn(i, item); stop {
			break
		}
		i++
	}
}

// SetLeaseForProviderByNode stores a lease for a provider by node in the module's KVStore.
func (k *Keeper) SetLeaseForProviderByNode(ctx sdk.Context, provAddr base.ProvAddress, nodeAddr base.NodeAddress, id uint64) {
	store := k.Store(ctx)
	key := types.LeaseForProviderByNodeKey(provAddr, nodeAddr, id)
	value := k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})

	store.Set(key, value)
}

// HasLeaseForProviderByNode checks if a lease for a provider by node exists in the module's KVStore based on the provider and node addresses and lease ID.
func (k *Keeper) HasLeaseForProviderByNode(ctx sdk.Context, provAddr base.ProvAddress, nodeAddr base.NodeAddress, id uint64) bool {
	store := k.Store(ctx)
	key := types.LeaseForProviderByNodeKey(provAddr, nodeAddr, id)

	return store.Has(key)
}

// DeleteLeaseForProviderByNode removes a lease for a provider by node from the module's KVStore based on the provider and node addresses and lease ID.
func (k *Keeper) DeleteLeaseForProviderByNode(ctx sdk.Context, provAddr base.ProvAddress, nodeAddr base.NodeAddress, id uint64) {
	store := k.Store(ctx)
	key := types.LeaseForProviderByNodeKey(provAddr, nodeAddr, id)

	store.Delete(key)
}

// GetLatestLeaseForProviderByNode retrieves the latest lease for a provider by node from the module's KVStore.
// If the lease exists, it returns the lease and 'found' as true; otherwise, it returns 'found' as false.
func (k *Keeper) GetLatestLeaseForProviderByNode(ctx sdk.Context, provAddr base.ProvAddress, nodeAddr base.NodeAddress) (lease v1.Lease, found bool) {
	store := k.Store(ctx)
	iterator := sdk.KVStoreReversePrefixIterator(store, types.GetLeaseForProviderByNodeKeyPrefix(provAddr, nodeAddr))

	defer iterator.Close()

	if iterator.Valid() {
		lease, found = k.GetLease(ctx, types.IDFromLeaseForProviderByNodeKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("lease for provider by node key %X does not exist", iterator.Key()))
		}
	}

	return lease, found
}

// SetLeaseForInactiveAt stores a lease for inactive status at a specific time in the module's KVStore.
func (k *Keeper) SetLeaseForInactiveAt(ctx sdk.Context, at time.Time, id uint64) {
	if at.IsZero() {
		return
	}

	store := k.Store(ctx)
	key := types.LeaseForInactiveAtKey(at, id)
	value := k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})

	store.Set(key, value)
}

// DeleteLeaseForInactiveAt removes a lease for inactive status at a specific time from the module's KVStore.
func (k *Keeper) DeleteLeaseForInactiveAt(ctx sdk.Context, at time.Time, id uint64) {
	if at.IsZero() {
		return
	}

	store := k.Store(ctx)
	key := types.LeaseForInactiveAtKey(at, id)

	store.Delete(key)
}

// IterateLeasesForInactiveAt iterates over all leases for inactive status at a specific time stored in the module's KVStore and calls the provided function for each lease.
// The iteration stops when the provided function returns 'true'.
func (k *Keeper) IterateLeasesForInactiveAt(ctx sdk.Context, at time.Time, fn func(index int, item v1.Lease) (stop bool)) {
	store := k.Store(ctx)
	iterator := store.Iterator(types.LeaseForInactiveAtKeyPrefix, sdk.PrefixEndBytes(types.GetLeaseForInactiveAtKeyPrefix(at)))

	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		item, found := k.GetLease(ctx, types.IDFromLeaseForInactiveAtKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("lease for inactive at key %X does not exist", iterator.Key()))
		}

		if stop := fn(i, item); stop {
			break
		}
		i++
	}
}

// SetLeaseForPayoutAt stores a lease for payout at a specific time in the module's KVStore.
func (k *Keeper) SetLeaseForPayoutAt(ctx sdk.Context, at time.Time, id uint64) {
	if at.IsZero() {
		return
	}

	store := k.Store(ctx)
	key := types.LeaseForPayoutAtKey(at, id)
	value := k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})

	store.Set(key, value)
}

// DeleteLeaseForPayoutAt removes a lease for payout at a specific time from the module's KVStore.
func (k *Keeper) DeleteLeaseForPayoutAt(ctx sdk.Context, at time.Time, id uint64) {
	if at.IsZero() {
		return
	}

	store := k.Store(ctx)
	key := types.LeaseForPayoutAtKey(at, id)

	store.Delete(key)
}

// IterateLeasesForPayoutAt iterates over all leases for payout at a specific time stored in the module's KVStore and calls the provided function for each lease.
// The iteration stops when the provided function returns 'true'.
func (k *Keeper) IterateLeasesForPayoutAt(ctx sdk.Context, at time.Time, fn func(index int, item v1.Lease) (stop bool)) {
	store := k.Store(ctx)
	iterator := store.Iterator(types.LeaseForPayoutAtKeyPrefix, sdk.PrefixEndBytes(types.GetLeaseForPayoutAtKeyPrefix(at)))

	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		lease, found := k.GetLease(ctx, types.IDFromLeaseForPayoutAtKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("lease for payout_at key %X does not exist", iterator.Key()))
		}

		if stop := fn(i, lease); stop {
			break
		}
		i++
	}
}

// SetLeaseForRenewalAt stores a lease for renewal at a specific time in the module's KVStore.
func (k *Keeper) SetLeaseForRenewalAt(ctx sdk.Context, at time.Time, id uint64) {
	if at.IsZero() {
		return
	}

	store := k.Store(ctx)
	key := types.LeaseForRenewalAtKey(at, id)
	value := k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})

	store.Set(key, value)
}

// DeleteLeaseForRenewalAt removes a lease for renewal at a specific time from the module's KVStore.
func (k *Keeper) DeleteLeaseForRenewalAt(ctx sdk.Context, at time.Time, id uint64) {
	if at.IsZero() {
		return
	}

	store := k.Store(ctx)
	key := types.LeaseForRenewalAtKey(at, id)

	store.Delete(key)
}

// IterateLeasesForRenewalAt iterates over all leases for renewal at a specific time stored in the module's KVStore and calls the provided function for each lease.
// The iteration stops when the provided function returns 'true'.
func (k *Keeper) IterateLeasesForRenewalAt(ctx sdk.Context, at time.Time, fn func(index int, item v1.Lease) (stop bool)) {
	store := k.Store(ctx)
	iterator := store.Iterator(types.LeaseForRenewalAtKeyPrefix, sdk.PrefixEndBytes(types.GetLeaseForRenewalAtKeyPrefix(at)))

	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		lease, found := k.GetLease(ctx, types.IDFromLeaseForRenewalAtKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("lease for renew key %X does not exist", iterator.Key()))
		}

		if stop := fn(i, lease); stop {
			break
		}
		i++
	}
}
