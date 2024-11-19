package expected

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	base "github.com/sentinel-official/hub/v12/types"
	leasetypes "github.com/sentinel-official/hub/v12/x/lease/types/v1"
	nodetypes "github.com/sentinel-official/hub/v12/x/node/types/v2"
	subscriptiontypes "github.com/sentinel-official/hub/v12/x/subscription/types/v3"
)

type AccountKeeper interface {
	GetAccount(ctx sdk.Context, address sdk.AccAddress) authtypes.AccountI
}

type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, address sdk.AccAddress) sdk.Coins
}

type LeaseKeeper interface {
	GetLease(ctx sdk.Context, id uint64) (leasetypes.Lease, bool)
	GetLatestLeaseForNodeByProvider(ctx sdk.Context, nodeAddr base.NodeAddress, provAddr base.ProvAddress) (leasetypes.Lease, bool)
}

type NodeKeeper interface {
	GetNode(ctx sdk.Context, addr base.NodeAddress) (nodetypes.Node, bool)
	SetNodeForPlan(ctx sdk.Context, id uint64, addr base.NodeAddress)
	DeleteNodeForPlan(ctx sdk.Context, id uint64, addr base.NodeAddress)
	GetNodesForPlan(ctx sdk.Context, id uint64) []nodetypes.Node
	HasNodeForPlan(ctx sdk.Context, id uint64, addr base.NodeAddress) bool
}

type ProviderKeeper interface {
	HasProvider(ctx sdk.Context, addr base.ProvAddress) bool
}

type SessionKeeper interface {
	PlanUnlinkNodePreHook(ctx sdk.Context, id uint64, addr base.NodeAddress) error
}

type SubscriptionKeeper interface {
	HandleMsgStartSession(ctx sdk.Context, msg *subscriptiontypes.MsgStartSessionRequest) (*subscriptiontypes.MsgStartSessionResponse, error)
	HandleMsgStartSubscription(ctx sdk.Context, msg *subscriptiontypes.MsgStartSubscriptionRequest) (*subscriptiontypes.MsgStartSubscriptionResponse, error)
}
