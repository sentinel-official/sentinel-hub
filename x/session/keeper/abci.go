package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/session/types/v3"
)

// handleInactivePendingSessions processes pending sessions that have become inactive at the current block time.
func (k *Keeper) handleInactivePendingSessions(ctx sdk.Context) {
	// Iterate through sessions that have become inactive at the current block time
	k.IterateSessionsForInactiveAt(ctx, ctx.BlockTime(), func(_ int, item v3.Session) bool {
		// Skip the session if its status is not active
		if !item.GetStatus().Equal(v1base.StatusActive) {
			return false
		}

		k.Logger(ctx).Info("Handling inactive pending session", "id", item.GetID())

		// Create a message to cancel the inactive pending session
		msg := &v3.MsgCancelSessionRequest{
			From: item.GetAccAddress(),
			ID:   item.GetID(),
		}

		// Get the appropriate handler for processing the cancel session message
		handler := k.router.Handler(msg)
		if handler == nil {
			panic(fmt.Errorf("nil handler for message route: %s", sdk.MsgTypeURL(msg)))
		}

		// Execute the handler to process the session cancel request
		resp, err := handler(ctx, msg)
		if err != nil {
			panic(err)
		}

		// Emit any events generated during the cancel process
		ctx.EventManager().EmitEvents(resp.GetEvents())
		return false
	})
}

// handleInactiveSessions processes sessions that are in the inactive pending state.
func (k *Keeper) handleInactiveSessions(ctx sdk.Context) {
	// Iterate through sessions that are inactive pending at the current block time
	k.IterateSessionsForInactiveAt(ctx, ctx.BlockTime(), func(_ int, item v3.Session) bool {
		// Skip the session if its status is not inactive pending
		if !item.GetStatus().Equal(v1base.StatusInactivePending) {
			return false
		}

		k.Logger(ctx).Info("Handling inactive session", "id", item.GetID())

		// Perform pre-hook processing for sessions that are transitioning to inactive
		if err := k.SessionInactivePreHook(ctx, item.GetID()); err != nil {
			panic(err)
		}

		// Delete the session from the inactive queue as it is now processed
		k.DeleteSessionForInactiveAt(ctx, item.GetInactiveAt(), item.GetID())

		// Emit an event indicating the update of the session status to inactive
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

// BeginBlock is called at the beginning of each block to handle session-related operations.
func (k *Keeper) BeginBlock(ctx sdk.Context) {
	// Handle inactive pending sessions at the beginning of each block
	k.handleInactivePendingSessions(ctx)

	// Handle sessions that have become inactive at the beginning of each block
	k.handleInactiveSessions(ctx)
}
