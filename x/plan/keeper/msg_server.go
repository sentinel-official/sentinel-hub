package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/plan/types"
	"github.com/sentinel-official/hub/v12/x/plan/types/v2"
)

// The following line asserts that the `msgServer` type implements the `types.MsgServiceServer` interface.
var (
	_ v2.MsgServiceServer = (*msgServer)(nil)
)

// msgServer is a message server that implements the `types.MsgServiceServer` interface.
type msgServer struct {
	Keeper // Keeper is an instance of the main keeper for the module.
}

// NewMsgServiceServer creates a new instance of `types.MsgServiceServer` using the provided Keeper.
func NewMsgServiceServer(k Keeper) v2.MsgServiceServer {
	return &msgServer{k}
}

// MsgCreate creates a new plan with the provided details and stores it in the Store.
// It validates the creation request, checks for provider existence, and assigns a unique ID to the plan.
func (k *msgServer) MsgCreate(c context.Context, msg *v2.MsgCreateRequest) (*v2.MsgCreateResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	// Convert the `msg.From` address (in Bech32 format) to a `base.ProvAddress`.
	provAddr, err := base.ProvAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	// Check if the provider with the given address exists in the network. If not, return an error.
	if !k.HasProvider(ctx, provAddr) {
		return nil, types.NewErrorProviderNotFound(provAddr)
	}

	// Get the current count of plans to assign a unique ID to the new plan.
	count := k.GetCount(ctx)
	plan := v2.Plan{
		ID:              count + 1,
		ProviderAddress: provAddr.String(),
		Duration:        msg.Duration,
		Gigabytes:       msg.Gigabytes,
		Prices:          msg.Prices,
		Status:          base.StatusInactive,
		StatusAt:        ctx.BlockTime(),
	}

	// Save the new plan in the Store and update the count for future plans.
	k.SetCount(ctx, count+1)
	k.SetPlan(ctx, plan)
	k.SetPlanForProvider(ctx, provAddr, plan.ID)

	// Emit an event to notify that a new plan has been created.
	ctx.EventManager().EmitTypedEvent(
		&v2.EventCreate{
			Address: plan.ProviderAddress,
			ID:      plan.ID,
		},
	)

	return &v2.MsgCreateResponse{}, nil
}

// MsgUpdateStatus updates the status of a plan.
// It validates the status update request, checks for plan existence, and updates the plan status and timestamp accordingly.
func (k *msgServer) MsgUpdateStatus(c context.Context, msg *v2.MsgUpdateStatusRequest) (*v2.MsgUpdateStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	// Get the plan from the Store based on the provided plan ID.
	plan, found := k.GetPlan(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorPlanNotFound(msg.ID)
	}

	// Check if the provided address (`msg.From`) matches the plan's provider address to verify authorization.
	if msg.From != plan.ProviderAddress {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	// If the current status of the plan is `Active`, handle the necessary updates for changing to `Inactive` status.
	if plan.Status.Equal(base.StatusActive) {
		if msg.Status.Equal(base.StatusInactive) {
			k.DeleteActivePlan(ctx, plan.ID)
		}
	}

	// If the current status of the plan is `Inactive`, handle the necessary updates for changing to `Active` status.
	if plan.Status.Equal(base.StatusInactive) {
		if msg.Status.Equal(base.StatusActive) {
			k.DeleteInactivePlan(ctx, plan.ID)
		}
	}

	// Update the plan's status and status timestamp with the provided new status and current block time.
	plan.Status = msg.Status
	plan.StatusAt = ctx.BlockTime()

	// Save the updated plan in the Store.
	k.SetPlan(ctx, plan)

	// Emit an event to notify that the plan status has been updated.
	ctx.EventManager().EmitTypedEvent(
		&v2.EventUpdateStatus{
			Status:  plan.Status,
			Address: plan.ProviderAddress,
			ID:      plan.ID,
		},
	)

	return &v2.MsgUpdateStatusResponse{}, nil
}

// MsgLinkNode links a node to a plan.
// It validates the link node request, checks for plan and node existence, and links the node to the plan.
func (k *msgServer) MsgLinkNode(c context.Context, msg *v2.MsgLinkNodeRequest) (*v2.MsgLinkNodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	// Get the plan from the Store based on the provided plan ID.
	plan, found := k.GetPlan(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorPlanNotFound(msg.ID)
	}

	// Check if the provided address (`msg.From`) matches the plan's provider address to verify authorization.
	if msg.From != plan.ProviderAddress {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	// Convert the `msg.NodeAddress` (node address) to a `base.NodeAddress`.
	nodeAddr, err := base.NodeAddressFromBech32(msg.NodeAddress)
	if err != nil {
		return nil, err
	}

	// Check if the node with the given address exists in the network. If not, return an error.
	if !k.HasNode(ctx, nodeAddr) {
		return nil, types.NewErrorNodeNotFound(nodeAddr)
	}

	// Link the node to the plan in the Store.
	k.SetNodeForPlan(ctx, plan.ID, nodeAddr)

	// Emit an event to notify that a node has been linked to the plan.
	ctx.EventManager().EmitTypedEvent(
		&v2.EventLinkNode{
			Address:     plan.ProviderAddress,
			NodeAddress: msg.NodeAddress,
			ID:          plan.ID,
		},
	)

	return &v2.MsgLinkNodeResponse{}, nil
}

// MsgUnlinkNode unlinks a node from a plan.
// It validates the unlink node request, checks for plan and node existence, and unlinks the node from the plan.
func (k *msgServer) MsgUnlinkNode(c context.Context, msg *v2.MsgUnlinkNodeRequest) (*v2.MsgUnlinkNodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	// Get the plan from the Store based on the provided plan ID.
	plan, found := k.GetPlan(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorPlanNotFound(msg.ID)
	}

	// Check if the provided address (`msg.From`) matches the plan's provider address to verify authorization.
	if msg.From != plan.ProviderAddress {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	// Convert the `msg.NodeAddress` (node address) to a `base.NodeAddress`.
	nodeAddr, err := base.NodeAddressFromBech32(msg.NodeAddress)
	if err != nil {
		return nil, err
	}

	// Unlink the node from the plan in the Store.
	k.DeleteNodeForPlan(ctx, plan.ID, nodeAddr)

	// Emit an event to notify that a node has been unlinked from the plan.
	ctx.EventManager().EmitTypedEvent(
		&v2.EventUnlinkNode{
			Address:     plan.ProviderAddress,
			NodeAddress: msg.NodeAddress,
			ID:          plan.ID,
		},
	)

	return &v2.MsgUnlinkNodeResponse{}, nil
}

// MsgSubscribe subscribes to a plan for a specific user account.
// It validates the subscription request and creates a new subscription for the provided plan and user account.
func (k *msgServer) MsgSubscribe(c context.Context, msg *v2.MsgSubscribeRequest) (*v2.MsgSubscribeResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	// Convert the `msg.From` address (in Bech32 format) to an `sdk.AccAddress`.
	accAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	// Create a new subscription for the provided plan, user account, and denom.
	subscription, err := k.CreateSubscriptionForPlan(ctx, accAddr, msg.ID, msg.Denom)
	if err != nil {
		return nil, err
	}

	// Emit an event to notify that a new subscription has been created.
	ctx.EventManager().EmitTypedEvent(
		&v2.EventCreateSubscription{
			Address:         subscription.Address,
			ProviderAddress: "",
			ID:              subscription.ID,
			PlanID:          subscription.PlanID,
		},
	)

	return &v2.MsgSubscribeResponse{}, nil
}
