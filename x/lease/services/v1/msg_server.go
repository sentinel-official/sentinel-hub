package v1

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/hub/v12/x/lease/keeper"
	"github.com/sentinel-official/hub/v12/x/lease/types/v1"
)

var (
	_ v1.MsgServiceServer = (*msgServer)(nil)
)

type msgServer struct {
	keeper.Keeper
}

func NewMsgServiceServer(k keeper.Keeper) v1.MsgServiceServer {
	return &msgServer{k}
}

func (m *msgServer) MsgEndLease(c context.Context, req *v1.MsgEndLeaseRequest) (*v1.MsgEndLeaseResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return m.HandleMsgEndLease(ctx, req)
}

func (m *msgServer) MsgRenewLease(_ context.Context, _ *v1.MsgRenewLeaseRequest) (*v1.MsgRenewLeaseResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (m *msgServer) MsgStartLease(c context.Context, req *v1.MsgStartLeaseRequest) (*v1.MsgStartLeaseResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return m.HandleMsgStartLease(ctx, req)
}

func (m *msgServer) MsgUpdateLease(c context.Context, req *v1.MsgUpdateLeaseRequest) (*v1.MsgUpdateLeaseResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return m.HandleMsgUpdateLease(ctx, req)
}

func (m *msgServer) MsgUpdateParams(c context.Context, req *v1.MsgUpdateParamsRequest) (*v1.MsgUpdateParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return m.HandleMsgUpdateParams(ctx, req)
}
