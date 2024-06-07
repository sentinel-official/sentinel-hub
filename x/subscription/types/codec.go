// DO NOT COVER

package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"

	"github.com/sentinel-official/hub/v12/x/subscription/types/v1"
)

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	v1types.RegisterInterfaces(registry)

	registry.RegisterInterface(
		"sentinel.subscription.v2.Subscription",
		(*Subscription)(nil),
		&NodeSubscription{},
		&PlanSubscription{},
	)

	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgCancelRequest{},
		&MsgAllocateRequest{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_MsgService_serviceDesc)
}
