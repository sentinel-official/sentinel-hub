package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	protobuf "github.com/gogo/protobuf/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/plan/types"
	"github.com/sentinel-official/hub/v12/x/plan/types/v2"
)

func (k *Keeper) SetActivePlan(ctx sdk.Context, plan v2.Plan) {
	var (
		store = k.Store(ctx)
		key   = types.ActivePlanKey(plan.ID)
		value = k.cdc.MustMarshal(&plan)
	)

	store.Set(key, value)
}

func (k *Keeper) HasActivePlan(ctx sdk.Context, id uint64) bool {
	var (
		store = k.Store(ctx)
		key   = types.ActivePlanKey(id)
	)

	return store.Has(key)
}

func (k *Keeper) GetActivePlan(ctx sdk.Context, id uint64) (plan v2.Plan, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.ActivePlanKey(id)
		value = store.Get(key)
	)

	if value == nil {
		return plan, false
	}

	k.cdc.MustUnmarshal(value, &plan)
	return plan, true
}

func (k *Keeper) DeleteActivePlan(ctx sdk.Context, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.ActivePlanKey(id)
	)

	store.Delete(key)
}

func (k *Keeper) SetInactivePlan(ctx sdk.Context, plan v2.Plan) {
	var (
		store = k.Store(ctx)
		key   = types.InactivePlanKey(plan.ID)
		value = k.cdc.MustMarshal(&plan)
	)

	store.Set(key, value)
}

func (k *Keeper) HasInactivePlan(ctx sdk.Context, id uint64) bool {
	var (
		store = k.Store(ctx)
		key   = types.InactivePlanKey(id)
	)

	return store.Has(key)
}

func (k *Keeper) GetInactivePlan(ctx sdk.Context, id uint64) (plan v2.Plan, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.InactivePlanKey(id)
		value = store.Get(key)
	)

	if value == nil {
		return plan, false
	}

	k.cdc.MustUnmarshal(value, &plan)
	return plan, true
}

func (k *Keeper) DeleteInactivePlan(ctx sdk.Context, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.InactivePlanKey(id)
	)

	store.Delete(key)
}

func (k *Keeper) SetPlan(ctx sdk.Context, plan v2.Plan) {
	switch plan.Status {
	case v1base.StatusActive:
		k.SetActivePlan(ctx, plan)
	case v1base.StatusInactive:
		k.SetInactivePlan(ctx, plan)
	default:
		panic(fmt.Errorf("failed to set the plan %v", plan))
	}
}

func (k *Keeper) HasPlan(ctx sdk.Context, id uint64) bool {
	return k.HasActivePlan(ctx, id) ||
		k.HasInactivePlan(ctx, id)
}

func (k *Keeper) GetPlan(ctx sdk.Context, id uint64) (plan v2.Plan, found bool) {
	plan, found = k.GetActivePlan(ctx, id)
	if found {
		return
	}

	plan, found = k.GetInactivePlan(ctx, id)
	if found {
		return
	}

	return plan, false
}

func (k *Keeper) GetPlans(ctx sdk.Context) (items v2.Plans) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.PlanKeyPrefix)
	)

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var item v2.Plan
		k.cdc.MustUnmarshal(iter.Value(), &item)
		items = append(items, item)
	}

	return items
}

func (k *Keeper) SetPlanForProvider(ctx sdk.Context, addr base.ProvAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.PlanForProviderKey(addr, id)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) DeletePlanForProvider(ctx sdk.Context, addr base.ProvAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.PlanForProviderKey(addr, id)
	)

	store.Delete(key)
}

func (k *Keeper) GetPlansForProvider(ctx sdk.Context, addr base.ProvAddress) (items v2.Plans) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.GetPlanForProviderKeyPrefix(addr))
	)

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		item, found := k.GetPlan(ctx, types.IDFromPlanForProviderKey(iter.Key()))
		if !found {
			panic(fmt.Errorf("plan for provider key %X does not exist", iter.Key()))
		}

		items = append(items, item)
	}

	return items
}
