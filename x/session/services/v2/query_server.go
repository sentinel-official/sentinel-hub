package v2

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/hub/v12/x/session/keeper"
	"github.com/sentinel-official/hub/v12/x/session/types/v2"
)

var (
	_ v2.QueryServiceServer = (*queryServer)(nil)
)

type queryServer struct {
	keeper.Keeper
}

func NewQueryServiceServer(k keeper.Keeper) v2.QueryServiceServer {
	return &queryServer{k}
}

func (q *queryServer) QueryParams(c context.Context, req *v2.QueryParamsRequest) (*v2.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return q.HandleQueryParams(ctx, req)
}

func (q *queryServer) QuerySession(_ context.Context, _ *v2.QuerySessionRequest) (*v2.QuerySessionResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (q *queryServer) QuerySessions(_ context.Context, _ *v2.QuerySessionsRequest) (*v2.QuerySessionsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (q *queryServer) QuerySessionsForAccount(_ context.Context, _ *v2.QuerySessionsForAccountRequest) (*v2.QuerySessionsForAccountResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (q *queryServer) QuerySessionsForNode(_ context.Context, _ *v2.QuerySessionsForNodeRequest) (*v2.QuerySessionsForNodeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (q *queryServer) QuerySessionsForSubscription(_ context.Context, _ *v2.QuerySessionsForSubscriptionRequest) (*v2.QuerySessionsForSubscriptionResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (q *queryServer) QuerySessionsForAllocation(_ context.Context, _ *v2.QuerySessionsForAllocationRequest) (*v2.QuerySessionsForAllocationResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
