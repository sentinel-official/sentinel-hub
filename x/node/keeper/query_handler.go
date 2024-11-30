package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/node/types"
	"github.com/sentinel-official/hub/v12/x/node/types/v3"
)

func (k *Keeper) HandleQueryNode(ctx sdk.Context, req *v3.QueryNodeRequest) (*v3.QueryNodeResponse, error) {
	addr, err := base.NodeAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	item, found := k.GetNode(ctx, addr)
	if !found {
		return nil, status.Errorf(codes.NotFound, "node does not exist for address %s", req.Address)
	}

	return &v3.QueryNodeResponse{Node: item}, nil
}

func (k *Keeper) HandleQueryNodes(ctx sdk.Context, req *v3.QueryNodesRequest) (res *v3.QueryNodesResponse, err error) {
	var (
		items     []v3.Node
		keyPrefix []byte
	)

	switch req.Status {
	case v1base.StatusActive:
		keyPrefix = types.ActiveNodeKeyPrefix
	case v1base.StatusInactive:
		keyPrefix = types.InactiveNodeKeyPrefix
	default:
		keyPrefix = types.NodeKeyPrefix
	}

	store := prefix.NewStore(k.Store(ctx), keyPrefix)
	pagination, err := sdkquery.Paginate(store, req.Pagination, func(_, value []byte) error {
		var item v3.Node
		if err := k.cdc.Unmarshal(value, &item); err != nil {
			return err
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v3.QueryNodesResponse{Nodes: items, Pagination: pagination}, nil
}

func (k *Keeper) HandleQueryNodesForPlan(ctx sdk.Context, req *v3.QueryNodesForPlanRequest) (*v3.QueryNodesForPlanResponse, error) {
	var (
		items []v3.Node
		store = prefix.NewStore(k.Store(ctx), types.GetNodeForPlanKeyPrefix(req.Id))
	)

	pagination, err := sdkquery.FilteredPaginate(store, req.Pagination, func(key, _ []byte, accumulate bool) (bool, error) {
		if !accumulate {
			return false, nil
		}

		item, found := k.GetNode(ctx, key[1:])
		if !found {
			return false, fmt.Errorf("node for key %X does not exist", key)
		}

		if req.Status.Equal(v1base.StatusUnspecified) || item.Status.Equal(req.Status) {
			items = append(items, item)
			return true, nil
		}

		return false, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v3.QueryNodesForPlanResponse{Nodes: items, Pagination: pagination}, nil
}

func (k *Keeper) HandleQueryParams(ctx sdk.Context, _ *v3.QueryParamsRequest) (*v3.QueryParamsResponse, error) {
	params := k.GetParams(ctx)
	return &v3.QueryParamsResponse{Params: params}, nil
}
