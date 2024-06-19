package v2

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	base "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/subscription/keeper"
	"github.com/sentinel-official/hub/v12/x/subscription/types"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v2"
)

var (
	_ v2.QueryServiceServer = (*queryServer)(nil)
)

type queryServer struct {
	codec.BinaryCodec
	keeper.Keeper
}

func NewQueryServiceServer(cdc codec.BinaryCodec, k keeper.Keeper) v2.QueryServiceServer {
	return &queryServer{
		BinaryCodec: cdc,
		Keeper:      k,
	}
}

func (k *queryServer) QuerySubscription(c context.Context, req *v2.QuerySubscriptionRequest) (*v2.QuerySubscriptionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	v, found := k.GetSubscription(ctx, req.Id)
	if !found {
		return nil, status.Errorf(codes.NotFound, "subscription does not exist for id %d", req.Id)
	}

	item, err := codectypes.NewAnyWithValue(v)
	if err != nil {
		return nil, err
	}

	return &v2.QuerySubscriptionResponse{Subscription: item}, nil
}

func (k *queryServer) QuerySubscriptions(c context.Context, req *v2.QuerySubscriptionsRequest) (*v2.QuerySubscriptionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var (
		items []*codectypes.Any
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(k.Store(ctx), types.SubscriptionKeyPrefix)
	)

	pagination, err := sdkquery.Paginate(store, req.Pagination, func(_, value []byte) error {
		var v v2.Subscription
		if err := k.UnmarshalInterface(value, &v); err != nil {
			return err
		}

		item, err := codectypes.NewAnyWithValue(v)
		if err != nil {
			return err
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v2.QuerySubscriptionsResponse{Subscriptions: items, Pagination: pagination}, nil
}

func (k *queryServer) QuerySubscriptionsForAccount(c context.Context, req *v2.QuerySubscriptionsForAccountRequest) (*v2.QuerySubscriptionsForAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	var (
		items []*codectypes.Any
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(k.Store(ctx), types.GetSubscriptionForAccountKeyPrefix(addr))
	)

	pagination, err := sdkquery.Paginate(store, req.Pagination, func(key, _ []byte) error {
		v, found := k.GetSubscription(ctx, sdk.BigEndianToUint64(key))
		if !found {
			return fmt.Errorf("subscription for key %X does not exist", key)
		}

		item, err := codectypes.NewAnyWithValue(v)
		if err != nil {
			return err
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v2.QuerySubscriptionsForAccountResponse{Subscriptions: items, Pagination: pagination}, nil
}

func (k *queryServer) QuerySubscriptionsForNode(c context.Context, req *v2.QuerySubscriptionsForNodeRequest) (*v2.QuerySubscriptionsForNodeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	addr, err := base.NodeAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	var (
		items []*codectypes.Any
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(k.Store(ctx), types.GetSubscriptionForNodeKeyPrefix(addr))
	)

	pagination, err := sdkquery.Paginate(store, req.Pagination, func(key, _ []byte) error {
		v, found := k.GetSubscription(ctx, sdk.BigEndianToUint64(key))
		if !found {
			return fmt.Errorf("subscription for key %X does not exist", key)
		}

		item, err := codectypes.NewAnyWithValue(v)
		if err != nil {
			return err
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v2.QuerySubscriptionsForNodeResponse{Subscriptions: items, Pagination: pagination}, nil
}

func (k *queryServer) QuerySubscriptionsForPlan(c context.Context, req *v2.QuerySubscriptionsForPlanRequest) (*v2.QuerySubscriptionsForPlanResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var (
		items []*codectypes.Any
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(k.Store(ctx), types.GetSubscriptionForPlanKeyPrefix(req.Id))
	)

	pagination, err := sdkquery.Paginate(store, req.Pagination, func(key, _ []byte) error {
		v, found := k.GetSubscription(ctx, sdk.BigEndianToUint64(key))
		if !found {
			return fmt.Errorf("subscription for key %X does not exist", key)
		}

		item, err := codectypes.NewAnyWithValue(v)
		if err != nil {
			return err
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v2.QuerySubscriptionsForPlanResponse{Subscriptions: items, Pagination: pagination}, nil
}

func (k *queryServer) QueryAllocation(c context.Context, req *v2.QueryAllocationRequest) (*v2.QueryAllocationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	ctx := sdk.UnwrapSDKContext(c)

	item, found := k.GetAllocation(ctx, req.Id, addr)
	if !found {
		return nil, status.Errorf(codes.NotFound, "allocation %d/%s does not exist", req.Id, req.Address)
	}

	return &v2.QueryAllocationResponse{Allocation: item}, nil
}

func (k *queryServer) QueryAllocations(c context.Context, req *v2.QueryAllocationsRequest) (*v2.QueryAllocationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var (
		items v2.Allocations
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(k.Store(ctx), types.GetAllocationForSubscriptionKeyPrefix(req.Id))
	)

	pagination, err := sdkquery.Paginate(store, req.Pagination, func(_, value []byte) error {
		var item v2.Allocation
		if err := k.Unmarshal(value, &item); err != nil {
			return err
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v2.QueryAllocationsResponse{Allocations: items, Pagination: pagination}, nil
}

func (k *queryServer) QueryPayout(c context.Context, req *v2.QueryPayoutRequest) (*v2.QueryPayoutResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	item, found := k.GetPayout(ctx, req.Id)
	if !found {
		return nil, status.Errorf(codes.NotFound, "payout does not exist for id %d", req.Id)
	}

	return &v2.QueryPayoutResponse{Payout: item}, nil
}

func (k *queryServer) QueryPayouts(c context.Context, req *v2.QueryPayoutsRequest) (res *v2.QueryPayoutsResponse, err error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var (
		items v2.Payouts
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(k.Store(ctx), types.PayoutKeyPrefix)
	)

	pagination, err := sdkquery.Paginate(store, req.Pagination, func(_, value []byte) error {
		var item v2.Payout
		if err := k.Unmarshal(value, &item); err != nil {
			return err
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v2.QueryPayoutsResponse{Payouts: items, Pagination: pagination}, nil
}

func (k *queryServer) QueryPayoutsForAccount(c context.Context, req *v2.QueryPayoutsForAccountRequest) (res *v2.QueryPayoutsForAccountResponse, err error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}

	var (
		items v2.Payouts
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(k.Store(ctx), types.GetPayoutForAccountKeyPrefix(addr))
	)

	pagination, err := sdkquery.Paginate(store, req.Pagination, func(key, _ []byte) error {
		item, found := k.GetPayout(ctx, sdk.BigEndianToUint64(key))
		if !found {
			return fmt.Errorf("payout for key %X does not exist", key)
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v2.QueryPayoutsForAccountResponse{Payouts: items, Pagination: pagination}, nil
}

func (k *queryServer) QueryPayoutsForNode(c context.Context, req *v2.QueryPayoutsForNodeRequest) (res *v2.QueryPayoutsForNodeResponse, err error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	addr, err := base.NodeAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}

	var (
		items v2.Payouts
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(k.Store(ctx), types.GetPayoutForNodeKeyPrefix(addr))
	)

	pagination, err := sdkquery.Paginate(store, req.Pagination, func(key, _ []byte) error {
		item, found := k.GetPayout(ctx, sdk.BigEndianToUint64(key))
		if !found {
			return fmt.Errorf("payout for key %X does not exist", key)
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v2.QueryPayoutsForNodeResponse{Payouts: items, Pagination: pagination}, nil
}

func (k *queryServer) QueryParams(c context.Context, _ *v2.QueryParamsRequest) (*v2.QueryParamsResponse, error) {
	var (
		ctx    = sdk.UnwrapSDKContext(c)
		params = k.GetParams(ctx)
	)

	return &v2.QueryParamsResponse{Params: params}, nil
}