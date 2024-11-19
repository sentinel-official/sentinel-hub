package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/plan/types"
	"github.com/sentinel-official/hub/v12/x/plan/types/v3"
	subscriptiontypes "github.com/sentinel-official/hub/v12/x/subscription/types/v3"
)

func (k *Keeper) HandleMsgCreatePlan(ctx sdk.Context, msg *v3.MsgCreatePlanRequest) (*v3.MsgCreatePlanResponse, error) {
	provAddr, err := base.ProvAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	if !k.provider.HasProvider(ctx, provAddr) {
		return nil, types.NewErrorProviderNotFound(provAddr)
	}

	count := k.GetCount(ctx)
	plan := v3.Plan{
		ID:          count + 1,
		ProvAddress: provAddr.String(),
		Bytes:       msg.Bytes,
		Duration:    msg.Duration,
		Prices:      msg.Prices,
		Private:     msg.Private,
		Status:      v1base.StatusInactive,
		StatusAt:    ctx.BlockTime(),
	}

	k.SetCount(ctx, count+1)
	k.SetPlan(ctx, plan)
	k.SetPlanForProvider(ctx, provAddr, plan.ID)

	ctx.EventManager().EmitTypedEvent(
		&v3.EventCreate{
			ID:          plan.ID,
			ProvAddress: plan.ProvAddress,
			Bytes:       plan.Bytes.String(),
			Duration:    plan.Duration,
			Private:     plan.Private,
			Prices:      plan.Prices.String(),
		},
	)

	return &v3.MsgCreatePlanResponse{
		ID: plan.ID,
	}, nil
}

func (k *Keeper) HandleMsgLinkNode(ctx sdk.Context, msg *v3.MsgLinkNodeRequest) (*v3.MsgLinkNodeResponse, error) {
	plan, found := k.GetPlan(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorPlanNotFound(msg.ID)
	}
	if msg.From != plan.ProvAddress {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	nodeAddr, err := base.NodeAddressFromBech32(msg.NodeAddress)
	if err != nil {
		return nil, err
	}

	node, found := k.node.GetNode(ctx, nodeAddr)
	if !found {
		return nil, types.NewErrorNodeNotFound(nodeAddr)
	}
	if !node.Status.Equal(v1base.StatusActive) {
		return nil, types.NewErrorInvalidNodeStatus(nodeAddr, node.Status)
	}

	if k.node.HasNodeForPlan(ctx, plan.ID, nodeAddr) {
		return nil, types.NewErrorDuplicateNodeForPlan(plan.ID, nodeAddr)
	}

	provAddr, err := base.ProvAddressFromBech32(plan.ProvAddress)
	if err != nil {
		return nil, err
	}

	if _, found := k.lease.GetLatestLeaseForNodeByProvider(ctx, nodeAddr, provAddr); !found {
		return nil, types.NewErrorLeaseNotFound(nodeAddr, provAddr)
	}

	k.node.SetNodeForPlan(ctx, plan.ID, nodeAddr)
	k.SetPlanForNodeByProvider(ctx, nodeAddr, provAddr, plan.ID)

	ctx.EventManager().EmitTypedEvent(
		&v3.EventLinkNode{
			ID:          plan.ID,
			ProvAddress: plan.ProvAddress,
			NodeAddress: node.Address,
		},
	)

	return &v3.MsgLinkNodeResponse{}, nil
}

func (k *Keeper) HandleMsgUnlinkNode(ctx sdk.Context, msg *v3.MsgUnlinkNodeRequest) (*v3.MsgUnlinkNodeResponse, error) {
	plan, found := k.GetPlan(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorPlanNotFound(msg.ID)
	}
	if msg.From != plan.ProvAddress {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	nodeAddr, err := base.NodeAddressFromBech32(msg.NodeAddress)
	if err != nil {
		return nil, err
	}

	if !k.node.HasNodeForPlan(ctx, plan.ID, nodeAddr) {
		return nil, types.NewErrorNodeForPlanNotFound(plan.ID, nodeAddr)
	}

	if err := k.PlanUnlinkNodePreHook(ctx, plan.ID, nodeAddr); err != nil {
		return nil, err
	}

	provAddr, err := base.ProvAddressFromBech32(plan.ProvAddress)
	if err != nil {
		return nil, err
	}

	k.node.DeleteNodeForPlan(ctx, plan.ID, nodeAddr)
	k.DeletePlanForNodeByProvider(ctx, nodeAddr, provAddr, plan.ID)

	ctx.EventManager().EmitTypedEvent(
		&v3.EventUnlinkNode{
			ID:          plan.ID,
			ProvAddress: plan.ProvAddress,
			NodeAddress: nodeAddr.String(),
		},
	)

	return &v3.MsgUnlinkNodeResponse{}, nil
}

func (k *Keeper) HandleMsgUpdatePlanStatus(ctx sdk.Context, msg *v3.MsgUpdatePlanStatusRequest) (*v3.MsgUpdatePlanStatusResponse, error) {
	plan, found := k.GetPlan(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorPlanNotFound(msg.ID)
	}
	if msg.From != plan.ProvAddress {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	if msg.Status.Equal(v1base.StatusActive) {
		if plan.Status.Equal(v1base.StatusInactive) {
			k.DeleteInactivePlan(ctx, plan.ID)
		}
	}
	if msg.Status.Equal(v1base.StatusInactive) {
		if plan.Status.Equal(v1base.StatusActive) {
			k.DeleteActivePlan(ctx, plan.ID)
		}
	}

	plan.Status = msg.Status
	plan.StatusAt = ctx.BlockTime()

	k.SetPlan(ctx, plan)
	ctx.EventManager().EmitTypedEvent(
		&v3.EventUpdate{
			ID:          plan.ID,
			ProvAddress: plan.ProvAddress,
			Status:      plan.Status,
		},
	)

	return &v3.MsgUpdatePlanStatusResponse{}, nil
}

func (k *Keeper) HandleMsgStartSession(ctx sdk.Context, msg *v3.MsgStartSessionRequest) (*v3.MsgStartSessionResponse, error) {
	subscriptionReq := &subscriptiontypes.MsgStartSubscriptionRequest{
		From:      msg.From,
		ID:        msg.ID,
		Denom:     msg.Denom,
		Renewable: msg.Renewable,
	}

	subscriptionResp, err := k.subscription.HandleMsgStartSubscription(ctx, subscriptionReq)
	if err != nil {
		return nil, err
	}

	sessionReq := &subscriptiontypes.MsgStartSessionRequest{
		From:        msg.From,
		ID:          subscriptionResp.ID,
		NodeAddress: msg.NodeAddress,
	}

	sessionResp, err := k.subscription.HandleMsgStartSession(ctx, sessionReq)
	if err != nil {
		return nil, err
	}

	return &v3.MsgStartSessionResponse{
		ID: sessionResp.ID,
	}, nil
}
