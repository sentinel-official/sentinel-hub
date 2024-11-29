package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v2"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v3"
)

func (k *Keeper) handleInactivePendingSubscriptions(ctx sdk.Context) {
	k.IterateSubscriptionsForInactiveAt(ctx, ctx.BlockTime(), func(_ int, item v3.Subscription) bool {
		if !item.Status.Equal(v1base.StatusActive) {
			return false
		}

		k.DeleteSubscriptionForInactiveAt(ctx, item.InactiveAt, item.ID)

		msg := &v3.MsgCancelSubscriptionRequest{
			From: item.AccAddress,
			ID:   item.ID,
		}

		handler := k.router.Handler(msg)
		if handler == nil {
			panic(fmt.Errorf("nil handler for message route: %s", sdk.MsgTypeURL(msg)))
		}

		resp, err := handler(ctx, msg)
		if err != nil {
			panic(err)
		}

		ctx.EventManager().EmitEvents(resp.GetEvents())
		return false
	})
}

func (k *Keeper) handleInactiveSubscriptions(ctx sdk.Context) {
	k.IterateSubscriptionsForInactiveAt(ctx, ctx.BlockTime(), func(_ int, item v3.Subscription) bool {
		if !item.Status.Equal(v1base.StatusInactivePending) {
			return false
		}

		k.DeleteSubscriptionForInactiveAt(ctx, item.InactiveAt, item.ID)

		k.DeleteSubscription(ctx, item.ID)
		k.DeleteSubscriptionForPlan(ctx, item.PlanID, item.ID)

		k.IterateAllocationsForSubscription(ctx, item.ID, func(_ int, item v2.Allocation) bool {
			accAddr, err := sdk.AccAddressFromBech32(item.Address)
			if err != nil {
				panic(err)
			}

			k.DeleteAllocation(ctx, item.ID, accAddr)
			k.DeleteSubscriptionForAccount(ctx, accAddr, item.ID)

			return false
		})

		ctx.EventManager().EmitTypedEvent(
			&v3.EventUpdate{
				ID:         item.ID,
				PlanID:     item.PlanID,
				AccAddress: item.AccAddress,
				Status:     v1base.StatusInactive,
				StatusAt:   ctx.BlockTime().String(),
			},
		)

		return false
	})
}

func (k *Keeper) handleSubscriptionRenewals(ctx sdk.Context) {
	k.IterateSubscriptionsForRenewalAt(ctx, ctx.BlockTime(), func(_ int, item v3.Subscription) bool {
		k.DeleteSubscriptionForRenewalAt(ctx, item.RenewalAt(), item.ID)

		msg := &v3.MsgRenewSubscriptionRequest{
			From:  item.AccAddress,
			ID:    item.ID,
			Denom: item.Price.Denom,
		}

		handler := k.router.Handler(msg)
		if handler == nil {
			panic(fmt.Errorf("nil handler for message route: %s", sdk.MsgTypeURL(msg)))
		}

		resp, err := handler(ctx, msg)
		if err != nil {
			panic(err)
		}

		ctx.EventManager().EmitEvents(resp.GetEvents())
		return false
	})
}

func (k *Keeper) BeginBlock(ctx sdk.Context) {
	k.handleSubscriptionRenewals(ctx)
	k.handleInactivePendingSubscriptions(ctx)
	k.handleInactiveSubscriptions(ctx)
}
