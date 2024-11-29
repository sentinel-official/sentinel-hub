package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/session/types/v3"
)

func (k *Keeper) NodeInactivePreHook(ctx sdk.Context, addr base.NodeAddress) error {
	return k.IterateSessionsForNode(ctx, addr, func(_ int, item v3.Session) (bool, error) {
		if !item.GetStatus().Equal(v1base.StatusActive) {
			return false, nil
		}

		msg := &v3.MsgCancelSessionRequest{
			From: item.GetAccAddress(),
			ID:   item.GetID(),
		}

		handler := k.router.Handler(msg)
		if handler == nil {
			return false, fmt.Errorf("nil handler for message route: %s", sdk.MsgTypeURL(msg))
		}

		resp, err := handler(ctx, msg)
		if err != nil {
			return false, err
		}

		ctx.EventManager().EmitEvents(resp.GetEvents())
		return false, nil
	})
}

func (k *Keeper) SubscriptionInactivePendingPreHook(ctx sdk.Context, id uint64) error {
	return k.IterateSessionsForSubscription(ctx, id, func(_ int, item v3.Session) (bool, error) {
		if !item.GetStatus().Equal(v1base.StatusActive) {
			return false, nil
		}

		msg := &v3.MsgCancelSessionRequest{
			From: item.GetAccAddress(),
			ID:   item.GetID(),
		}

		handler := k.router.Handler(msg)
		if handler == nil {
			return false, fmt.Errorf("nil handler for message route: %s", sdk.MsgTypeURL(msg))
		}

		resp, err := handler(ctx, msg)
		if err != nil {
			return false, err
		}

		ctx.EventManager().EmitEvents(resp.GetEvents())
		return false, nil
	})
}

func (k *Keeper) PlanUnlinkNodePreHook(ctx sdk.Context, id uint64, addr base.NodeAddress) error {
	return k.IterateSessionsForPlanByNode(ctx, id, addr, func(_ int, item v3.Session) (bool, error) {
		if !item.GetStatus().Equal(v1base.StatusActive) {
			return false, nil
		}

		msg := &v3.MsgCancelSessionRequest{
			From: item.GetAccAddress(),
			ID:   item.GetID(),
		}

		handler := k.router.Handler(msg)
		if handler == nil {
			return false, fmt.Errorf("nil handler for message route: %s", sdk.MsgTypeURL(msg))
		}

		resp, err := handler(ctx, msg)
		if err != nil {
			return false, err
		}

		ctx.EventManager().EmitEvents(resp.GetEvents())
		return false, nil
	})
}
