// DO NOT COVER

package v2

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgStartRequest{},
		&MsgUpdateDetailsRequest{},
		&MsgEndRequest{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_MsgService_serviceDesc)
}
