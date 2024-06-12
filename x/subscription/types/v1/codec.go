// DO NOT COVER

package v1

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgSubscribeToNodeRequest{},
		&MsgSubscribeToPlanRequest{},
		&MsgCancelRequest{},
		&MsgAddQuotaRequest{},
		&MsgUpdateQuotaRequest{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_MsgService_serviceDesc)
}
