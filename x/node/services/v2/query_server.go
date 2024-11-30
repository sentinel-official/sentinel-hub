package v2

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/hub/v12/x/node/keeper"
	"github.com/sentinel-official/hub/v12/x/node/types/v2"
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

func (q *queryServer) QueryNode(_ context.Context, _ *v2.QueryNodeRequest) (*v2.QueryNodeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (q *queryServer) QueryNodes(_ context.Context, _ *v2.QueryNodesRequest) (res *v2.QueryNodesResponse, err error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (q *queryServer) QueryNodesForPlan(_ context.Context, _ *v2.QueryNodesForPlanRequest) (*v2.QueryNodesForPlanResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (q *queryServer) QueryParams(_ context.Context, _ *v2.QueryParamsRequest) (*v2.QueryParamsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
