package v1

import (
	"crypto/rand"
	"testing"

	sdkquery "github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"

	"github.com/sentinel-official/hub/v12/x/swap/types"
)

func TestNewQueryParamsRequest(t *testing.T) {
	require.Equal(
		t,
		&QueryParamsRequest{},
		NewQueryParamsRequest(),
	)
}

func TestNewQuerySwapRequest(t *testing.T) {
	var (
		bytes []byte
	)

	for i := 0; i < 20; i++ {
		bytes = make([]byte, i)
		_, _ = rand.Read(bytes)

		require.Equal(
			t,
			&QuerySwapRequest{
				TxHash: types.BytesToHash(bytes).Bytes(),
			},
			NewQuerySwapRequest(types.BytesToHash(bytes)),
		)
	}
}

func TestNewQuerySwapsRequest(t *testing.T) {
	var (
		pagination *sdkquery.PageRequest
	)

	for i := 0; i < 20; i++ {
		pagination = &sdkquery.PageRequest{
			Key:        make([]byte, i),
			Offset:     uint64(i),
			Limit:      uint64(i),
			CountTotal: i/2 == 0,
		}

		require.Equal(
			t,
			&QuerySwapsRequest{
				Pagination: pagination,
			},
			NewQuerySwapsRequest(pagination),
		)
	}
}
