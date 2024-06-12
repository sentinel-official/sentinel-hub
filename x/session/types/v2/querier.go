package v2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	base "github.com/sentinel-official/hub/v12/types"
)

func NewQuerySessionRequest(id uint64) *QuerySessionRequest {
	return &QuerySessionRequest{
		Id: id,
	}
}

func NewQuerySessionsRequest(pagination *query.PageRequest) *QuerySessionsRequest {
	return &QuerySessionsRequest{
		Pagination: pagination,
	}
}

func NewQuerySessionsForAccountRequest(addr sdk.AccAddress, pagination *query.PageRequest) *QuerySessionsForAccountRequest {
	return &QuerySessionsForAccountRequest{
		Address:    addr.String(),
		Pagination: pagination,
	}
}

func NewQuerySessionsForNodeRequest(addr base.NodeAddress, pagination *query.PageRequest) *QuerySessionsForNodeRequest {
	return &QuerySessionsForNodeRequest{
		Address:    addr.String(),
		Pagination: pagination,
	}
}

func NewQuerySessionsForSubscriptionRequest(id uint64, pagination *query.PageRequest) *QuerySessionsForSubscriptionRequest {
	return &QuerySessionsForSubscriptionRequest{
		Id:         id,
		Pagination: pagination,
	}
}

func NewQuerySessionsForAllocationRequest(id uint64, addr sdk.AccAddress, pagination *query.PageRequest) *QuerySessionsForAllocationRequest {
	return &QuerySessionsForAllocationRequest{
		Id:         id,
		Address:    addr.String(),
		Pagination: pagination,
	}
}

func NewQueryParamsRequest() *QueryParamsRequest {
	return &QueryParamsRequest{}
}
