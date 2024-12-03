package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/subscription/types"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v3"
)

// SessionInactivePreHook handles the necessary operations when a session becomes inactive.
func (k *Keeper) SessionInactivePreHook(ctx sdk.Context, id uint64) error {
	k.Logger(ctx).Info("Running session inactive pre-hook", "id", id)

	// Retrieve the session by ID; return an error if not found.
	item, found := k.GetSession(ctx, id)
	if !found {
		return types.NewErrorSessionNotFound(id)
	}

	// Assert the retrieved session to the v3.Session type; return nil if the assertion fails.
	session, ok := item.(*v3.Session)
	if !ok {
		return nil
	}

	// Ensure the session status is "InactivePending"; return an error if it has a different status.
	if !session.Status.Equal(v1base.StatusInactivePending) {
		return types.NewErrorInvalidSessionStatus(session.ID, session.Status)
	}

	// Retrieve the subscription associated with the session; return an error if not found.
	subscription, found := k.GetSubscription(ctx, session.SubscriptionID)
	if !found {
		return types.NewErrorSubscriptionNotFound(session.SubscriptionID)
	}

	// Convert the session's account and node addresses from Bech32 format.
	accAddr, err := sdk.AccAddressFromBech32(session.AccAddress)
	if err != nil {
		return err
	}

	nodeAddr, err := base.NodeAddressFromBech32(session.NodeAddress)
	if err != nil {
		return err
	}

	// Retrieve the allocation for the subscription and account; return an error if not found.
	alloc, found := k.GetAllocation(ctx, subscription.ID, accAddr)
	if !found {
		return types.NewErrorAllocationNotFound(subscription.ID, accAddr)
	}

	// Calculate the total utilised bytes as the sum of download and upload bytes.
	utilisedBytes := session.DownloadBytes.Add(session.UploadBytes)

	// Update the utilised bytes in the allocation; cap it at the granted bytes if it exceeds the limit.
	alloc.UtilisedBytes = alloc.UtilisedBytes.Add(utilisedBytes)
	if alloc.UtilisedBytes.GT(alloc.GrantedBytes) {
		alloc.UtilisedBytes = alloc.GrantedBytes
	}

	// Save the updated allocation in the store.
	k.SetAllocation(ctx, alloc)

	// Emit an event to log the allocation update.
	ctx.EventManager().EmitTypedEvent(
		&v3.EventAllocate{
			ID:            alloc.ID,
			AccAddress:    alloc.Address,
			GrantedBytes:  alloc.GrantedBytes.String(),
			UtilisedBytes: alloc.UtilisedBytes.String(),
		},
	)

	// Delete the session records associated with allocation, node, plan, and subscription from the store.
	k.DeleteSessionForAllocation(ctx, subscription.ID, accAddr, session.ID)
	k.DeleteSessionForPlanByNode(ctx, subscription.PlanID, nodeAddr, session.ID)
	k.DeleteSessionForSubscription(ctx, subscription.ID, session.ID)

	return nil
}
