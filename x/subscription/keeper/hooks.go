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
	// Retrieve the session by ID and check if it exists, return an error if not found.
	item, found := k.GetSession(ctx, id)
	if !found {
		return types.NewErrorSessionNotFound(id)
	}

	// Validate that the session status is "InactivePending", otherwise return an error.
	if !item.GetStatus().Equal(v1base.StatusInactivePending) {
		return types.NewErrorInvalidSessionStatus(item.GetID(), item.GetStatus())
	}

	// Assert the item to a v3.Session type and proceed only if the assertion is successful.
	session, ok := item.(*v3.Session)
	if !ok {
		return nil
	}

	// Retrieve the subscription linked to the session, and check if it exists, return an error if not found.
	subscription, found := k.GetSubscription(ctx, session.SubscriptionID)
	if !found {
		return types.NewErrorSubscriptionNotFound(session.SubscriptionID)
	}

	// Convert the account address from Bech32 format, and panic if conversion fails.
	accAddr, err := sdk.AccAddressFromBech32(item.GetAccAddress())
	if err != nil {
		panic(err)
	}

	// Convert the node address from Bech32 format, return an error if conversion fails.
	nodeAddr, err := base.NodeAddressFromBech32(item.GetNodeAddress())
	if err != nil {
		return err
	}

	// Retrieve the allocation for the subscription and account, return an error if not found.
	alloc, found := k.GetAllocation(ctx, subscription.ID, accAddr)
	if !found {
		return types.NewErrorAllocationNotFound(subscription.ID, accAddr)
	}

	// Calculate the total utilised bytes by adding download and upload bytes.
	utilisedBytes := session.DownloadBytes.Add(session.UploadBytes)

	// Update the utilised bytes in the allocation, ensuring it does not exceed the granted bytes.
	alloc.UtilisedBytes = alloc.UtilisedBytes.Add(utilisedBytes)
	if alloc.UtilisedBytes.GT(alloc.GrantedBytes) {
		alloc.UtilisedBytes = alloc.GrantedBytes
	}

	// Update the allocation in the store with the new utilised bytes.
	k.SetAllocation(ctx, alloc)

	// Emit an event indicating the allocation update.
	ctx.EventManager().EmitTypedEvent(
		&v3.EventAllocate{
			ID:            alloc.ID,
			AccAddress:    alloc.Address,
			GrantedBytes:  alloc.GrantedBytes.String(),
			UtilisedBytes: alloc.UtilisedBytes.String(),
		},
	)

	// Delete the session records from the store, including records for account, allocation, node, plan, and subscription.
	k.DeleteSession(ctx, item.GetID())
	k.DeleteSessionForAccount(ctx, accAddr, item.GetID())
	k.DeleteSessionForAllocation(ctx, subscription.ID, accAddr, item.GetID())
	k.DeleteSessionForNode(ctx, nodeAddr, item.GetID())
	k.DeleteSessionForPlanByNode(ctx, subscription.PlanID, nodeAddr, item.GetID())
	k.DeleteSessionForSubscription(ctx, subscription.ID, item.GetID())

	return nil
}
