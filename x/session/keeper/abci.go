package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/session/types/v3"
)

func (k *Keeper) handleInactivePendingSessions(ctx sdk.Context) {
	k.IterateSessionsForInactiveAt(ctx, ctx.BlockTime(), func(_ int, item v3.Session) bool {
		if !item.GetStatus().Equal(v1base.StatusActive) {
			return false
		}

		k.DeleteSessionForInactiveAt(ctx, item.GetInactiveAt(), item.GetID())

		msg := &v3.MsgCancelSessionRequest{
			From: item.GetAccAddress(),
			ID:   item.GetID(),
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

func (k *Keeper) handleInactiveSessions(ctx sdk.Context) {
	k.IterateSessionsForInactiveAt(ctx, ctx.BlockTime(), func(_ int, item v3.Session) bool {
		if !item.GetStatus().Equal(v1base.StatusInactivePending) {
			return false
		}

		k.DeleteSessionForInactiveAt(ctx, item.GetInactiveAt(), item.GetID())

		if err := k.HandleInactiveSession(ctx, item.GetID()); err != nil {
			panic(err)
		}

		ctx.EventManager().EmitTypedEvent(
			&v3.EventUpdateStatus{
				ID:          item.GetID(),
				AccAddress:  item.GetAccAddress(),
				NodeAddress: item.GetNodeAddress(),
				Status:      v1base.StatusInactive,
				StatusAt:    ctx.BlockTime().String(),
			},
		)

		return false
	})
}

func (k *Keeper) BeginBlock(ctx sdk.Context) {
	k.handleInactivePendingSessions(ctx)
	k.handleInactiveSessions(ctx)
}
