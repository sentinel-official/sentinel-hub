package keeper

import (
	"context"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/subscription/types"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v2"
)

// The following line asserts that the `msgServer` type implements the `v2.MsgServiceServer` interface.
var (
	_ v2.MsgServiceServer = (*msgServer)(nil)
)

// msgServer is a message server that implements the `v2.MsgServiceServer` interface.
type msgServer struct {
	Keeper // Keeper is an instance of the main keeper for the module.
}

// NewMsgServiceServer creates a new instance of `v2.MsgServiceServer` using the provided Keeper.
func NewMsgServiceServer(k Keeper) v2.MsgServiceServer {
	return &msgServer{k}
}

// MsgCancel cancels an active subscription.
// It validates the cancel request and performs necessary operations to set the subscription to the inactive state.
func (k *msgServer) MsgCancel(c context.Context, msg *v2.MsgCancelRequest) (*v2.MsgCancelResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	// Convert the `msg.From` address (in Bech32 format) to an `sdk.AccAddress`.
	fromAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	// Get the subscription from the Store based on the provided subscription ID (msg.ID).
	subscription, found := k.GetSubscription(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorSubscriptionNotFound(msg.ID)
	}

	// Check if the subscription is in an active state. If not, return an error.
	if !subscription.GetStatus().Equal(v1base.StatusActive) {
		return nil, types.NewErrorInvalidSubscriptionStatus(subscription.GetID(), subscription.GetStatus())
	}

	// Check if the `msg.From` address matches the owner address of the subscription. If not, return an error.
	if !fromAddr.Equals(subscription.GetAddress()) {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	// Get the status change delay from the store.
	statusChangeDelay := k.StatusChangeDelay(ctx)

	// Delete the subscription from the Store for the time it becomes inactive.
	k.DeleteSubscriptionForInactiveAt(ctx, subscription.GetInactiveAt(), subscription.GetID())

	// Run the SubscriptionInactivePendingHook to perform custom actions before setting the subscription to inactive pending state.
	if err = k.SubscriptionInactivePendingHook(ctx, subscription.GetID()); err != nil {
		return nil, err
	}

	// Calculate the duration for which the subscription will be in the inactive state.
	subscription.SetInactiveAt(ctx.BlockTime().Add(statusChangeDelay))
	subscription.SetStatus(v1base.StatusInactivePending)
	subscription.SetStatusAt(ctx.BlockTime())

	// Update the subscription in the Store.
	k.SetSubscription(ctx, subscription)

	// Add the subscription back to the Store with the new inactive time.
	k.SetSubscriptionForInactiveAt(ctx, subscription.GetInactiveAt(), subscription.GetID())

	// Emit an event to notify that the subscription status has been updated.
	ctx.EventManager().EmitTypedEvent(
		&v2.EventUpdateStatus{
			Status:  v1base.StatusInactivePending,
			Address: subscription.GetAddress().String(),
			ID:      subscription.GetID(),
			PlanID:  0,
		},
	)

	// If the subscription is a NodeSubscription and the duration is specified in hours (non-zero), update the associated payout.
	if s, ok := subscription.(*v2.NodeSubscription); ok && s.Hours != 0 {
		payout, found := k.GetPayout(ctx, s.GetID())
		if !found {
			return nil, types.NewErrorPayoutNotFound(s.GetID())
		}

		var (
			accAddr  = payout.GetAddress()
			nodeAddr = payout.GetNodeAddress()
		)

		// Delete the payout from the Store for the given account and node.
		k.DeletePayoutForAccountByNode(ctx, accAddr, nodeAddr, payout.ID)
		k.DeletePayoutForNextAt(ctx, payout.NextAt, payout.ID)

		// Reset the `NextAt` field of the payout and update it in the Store.
		payout.NextAt = time.Time{}
		k.SetPayout(ctx, payout)
	}

	return &v2.MsgCancelResponse{}, nil
}

// MsgAllocate allocates bandwidth to another address.
// It validates the allocation request and updates the granted bytes for both the sender and receiver of the bandwidth.
func (k *msgServer) MsgAllocate(c context.Context, msg *v2.MsgAllocateRequest) (*v2.MsgAllocateResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	// Convert the `msg.From` address (in Bech32 format) to an `sdk.AccAddress`.
	fromAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	// Get the subscription from the Store based on the provided subscription ID (msg.ID).
	subscription, found := k.GetSubscription(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorSubscriptionNotFound(msg.ID)
	}

	// Check if the subscription type is a plan. If not, return an error.
	if _, ok := subscription.(*v2.PlanSubscription); !ok {
		return nil, types.NewErrorInvalidSubscription(subscription.GetID())
	}

	// Check if the `msg.From` address matches the owner address of the subscription. If not, return an error.
	if !fromAddr.Equals(subscription.GetAddress()) {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	// Get the existing allocation for the sender.
	fromAlloc, found := k.GetAllocation(ctx, subscription.GetID(), fromAddr)
	if !found {
		return nil, types.NewErrorAllocationNotFound(subscription.GetID(), fromAddr)
	}

	// Convert the `msg.Address` (receiver's address) from Bech32 format to an `sdk.AccAddress`.
	toAddr, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	// Get the existing allocation for the receiver.
	toAlloc, found := k.GetAllocation(ctx, subscription.GetID(), toAddr)
	if !found {
		// If the receiver has no existing allocation, create a new one.
		toAlloc = v2.Allocation{
			ID:            subscription.GetID(),
			Address:       toAddr.String(),
			GrantedBytes:  sdkmath.ZeroInt(),
			UtilisedBytes: sdkmath.ZeroInt(),
		}

		// Update the subscription in the Store to associate it with the new receiver.
		k.SetSubscriptionForAccount(ctx, toAddr, subscription.GetID())
	}

	// Calculate the available bytes for the sender and check if it is sufficient for the allocation.
	grantedBytes := fromAlloc.GrantedBytes.Add(toAlloc.GrantedBytes)
	utilisedBytes := fromAlloc.UtilisedBytes.Add(toAlloc.UtilisedBytes)
	availableBytes := grantedBytes.Sub(utilisedBytes)

	if msg.Bytes.GT(availableBytes) {
		return nil, types.NewErrorInsufficientBytes(subscription.GetID(), msg.Bytes)
	}

	// Update the allocation for the sender after deducting the allocated bytes.
	fromAlloc.GrantedBytes = availableBytes.Sub(msg.Bytes)
	if fromAlloc.GrantedBytes.LT(fromAlloc.UtilisedBytes) {
		return nil, types.NewErrorInvalidAllocation(subscription.GetID(), fromAddr)
	}

	// Update the sender's allocation in the Store.
	k.SetAllocation(ctx, fromAlloc)

	// Emit an event to notify that the sender's allocation has been updated.
	ctx.EventManager().EmitTypedEvent(
		&v2.EventAllocate{
			Address:       fromAlloc.Address,
			GrantedBytes:  fromAlloc.GrantedBytes,
			UtilisedBytes: fromAlloc.UtilisedBytes,
			ID:            fromAlloc.ID,
		},
	)

	// Update the allocation for the receiver after adding the allocated bytes.
	toAlloc.GrantedBytes = msg.Bytes
	if toAlloc.GrantedBytes.LT(toAlloc.UtilisedBytes) {
		return nil, types.NewErrorInvalidAllocation(subscription.GetID(), toAddr)
	}

	// Update the receiver's allocation in the Store.
	k.SetAllocation(ctx, toAlloc)

	// Emit an event to notify that the receiver's allocation has been updated.
	ctx.EventManager().EmitTypedEvent(
		&v2.EventAllocate{
			Address:       toAlloc.Address,
			GrantedBytes:  toAlloc.GrantedBytes,
			UtilisedBytes: toAlloc.UtilisedBytes,
			ID:            toAlloc.ID,
		},
	)

	return &v2.MsgAllocateResponse{}, nil
}
