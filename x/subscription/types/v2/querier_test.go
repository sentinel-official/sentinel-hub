package v2

import (
	"crypto/rand"
	"testing"

	base "github.com/sentinel-official/hub/v12/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
)

func TestNewQueryParamsRequest(t *testing.T) {
	require.Equal(t,
		&QueryParamsRequest{},
		NewQueryParamsRequest(),
	)
}

func TestNewQueryAllocationRequest(t *testing.T) {
	var (
		addr []byte
	)

	for i := 0; i < 512; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		require.Equal(
			t,
			&QueryAllocationRequest{
				Id:      uint64(i),
				Address: sdk.AccAddress(addr).String(),
			},
			NewQueryAllocationRequest(uint64(i), addr),
		)
	}
}

func TestNewQueryAllocationsRequest(t *testing.T) {
	var (
		pagination *query.PageRequest
	)

	for i := 0; i < 512; i += 64 {
		pagination = &query.PageRequest{
			Key:        make([]byte, i),
			Offset:     uint64(i),
			Limit:      uint64(i),
			CountTotal: i/2 == 0,
		}

		require.Equal(
			t,
			&QueryAllocationsRequest{
				Id:         uint64(i),
				Pagination: pagination,
			},
			NewQueryAllocationsRequest(uint64(i), pagination),
		)
	}
}

func TestNewQuerySubscriptionRequest(t *testing.T) {
	for i := 0; i < 512; i += 64 {
		require.Equal(
			t,
			&QuerySubscriptionRequest{
				Id: uint64(i),
			},
			NewQuerySubscriptionRequest(uint64(i)),
		)
	}
}

func TestNewQuerySubscriptionsForAccountRequest(t *testing.T) {
	var (
		addr       []byte
		pagination *query.PageRequest
	)

	for i := 0; i < 512; i += 64 {
		addr = make([]byte, i)
		pagination = &query.PageRequest{
			Key:        make([]byte, i),
			Offset:     uint64(i),
			Limit:      uint64(i),
			CountTotal: i/2 == 0,
		}

		_, _ = rand.Read(addr)
		_, _ = rand.Read(pagination.Key)

		require.Equal(
			t,
			&QuerySubscriptionsForAccountRequest{
				Address:    sdk.AccAddress(addr).String(),
				Pagination: pagination,
			},
			NewQuerySubscriptionsForAccountRequest(addr, pagination),
		)
	}
}

func TestNewQuerySubscriptionsForNodeRequest(t *testing.T) {
	var (
		addr       []byte
		pagination *query.PageRequest
	)

	for i := 0; i < 512; i += 64 {
		addr = make([]byte, i)
		pagination = &query.PageRequest{
			Key:        make([]byte, i),
			Offset:     uint64(i),
			Limit:      uint64(i),
			CountTotal: i/2 == 0,
		}

		_, _ = rand.Read(addr)
		_, _ = rand.Read(pagination.Key)

		require.Equal(
			t,
			&QuerySubscriptionsForNodeRequest{
				Address:    base.NodeAddress(addr).String(),
				Pagination: pagination,
			},
			NewQuerySubscriptionsForNodeRequest(addr, pagination),
		)
	}
}

func TestNewQuerySubscriptionsForPlanRequest(t *testing.T) {
	var (
		pagination *query.PageRequest
	)

	for i := 0; i < 512; i += 64 {
		pagination = &query.PageRequest{
			Key:        make([]byte, i),
			Offset:     uint64(i),
			Limit:      uint64(i),
			CountTotal: i/2 == 0,
		}
		_, _ = rand.Read(pagination.Key)

		require.Equal(
			t,
			&QuerySubscriptionsForPlanRequest{
				Id:         uint64(i),
				Pagination: pagination,
			},
			NewQuerySubscriptionsForPlanRequest(uint64(i), pagination),
		)
	}
}

func TestNewQuerySubscriptionsRequest(t *testing.T) {
	var (
		pagination *query.PageRequest
	)

	for i := 0; i < 512; i += 64 {
		pagination = &query.PageRequest{
			Key:        make([]byte, i),
			Offset:     uint64(i),
			Limit:      uint64(i),
			CountTotal: i/2 == 0,
		}

		require.Equal(
			t,
			&QuerySubscriptionsRequest{
				Pagination: pagination,
			},
			NewQuerySubscriptionsRequest(pagination),
		)
	}
}
