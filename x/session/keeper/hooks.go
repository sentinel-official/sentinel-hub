package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/session/types/v3"
)

func (k *Keeper) LeaseInactivePreHook(_ sdk.Context, _ uint64) error {
	return nil
}

func (k *Keeper) NodeInactivePreHook(ctx sdk.Context, addr base.NodeAddress) error {
	k.IterateSessionsForNode(ctx, addr, func(_ int, item v3.Session) bool {
		if !item.GetStatus().Equal(v1base.StatusActive) {
			return false
		}

		msg := &v3.MsgCancelSessionRequest{
			From: item.GetAccAddress(),
			ID:   item.GetID(),
		}

		if _, err := k.HandleMsgCancelSession(ctx, msg); err != nil {
			panic(err)
		}

		return false
	})

	return nil
}

func (k *Keeper) SubscriptionInactivePendingPreHook(ctx sdk.Context, id uint64) error {
	k.IterateSessionsForSubscription(ctx, id, func(_ int, item v3.Session) bool {
		if !item.GetStatus().Equal(v1base.StatusActive) {
			return false
		}

		msg := &v3.MsgCancelSessionRequest{
			From: item.GetAccAddress(),
			ID:   item.GetID(),
		}

		if _, err := k.HandleMsgCancelSession(ctx, msg); err != nil {
			panic(err)
		}

		return false
	})

	return nil
}
