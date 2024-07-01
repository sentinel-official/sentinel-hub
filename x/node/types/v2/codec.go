package v2

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkmsgservice "github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgRegisterRequest{},
		&MsgUpdateDetailsRequest{},
		&MsgUpdateStatusRequest{},
		&MsgSubscribeRequest{},
	)

	sdkmsgservice.RegisterMsgServiceDesc(registry, &_MsgService_serviceDesc)
}
