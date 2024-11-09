package v1

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/oracle/keeper"
	"github.com/sentinel-official/hub/v12/x/oracle/types/v1"
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

func (m *msgServer) MsgAddAsset(c context.Context, req *v1.MsgAddAssetRequest) (*v1.MsgAddAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return m.HandleMsgAddAsset(ctx, req)
}

func (m *msgServer) MsgDeleteAsset(c context.Context, req *v1.MsgDeleteAssetRequest) (*v1.MsgDeleteAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return m.HandleMsgDeleteAsset(ctx, req)
}

func (m *msgServer) MsgUpdateAsset(c context.Context, req *v1.MsgUpdateAssetRequest) (*v1.MsgUpdateAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return m.HandleMsgUpdateAsset(ctx, req)
}

func (m *msgServer) MsgUpdateParams(c context.Context, req *v1.MsgUpdateParamsRequest) (*v1.MsgUpdateParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return m.HandleMsgUpdateParams(ctx, req)
}
