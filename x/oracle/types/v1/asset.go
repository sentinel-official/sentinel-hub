package v1

import (
	"errors"
	"math/big"

	sdkmath "cosmossdk.io/math"
	abcitypes "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/third_party/osmosis/x/poolmanager/client/queryproto"
	protorevtypes "github.com/sentinel-official/hub/v12/third_party/osmosis/x/protorev/types"
)

func (a *Asset) ProtoRevPoolRequest(cdc codec.Codec) abcitypes.RequestQuery {
	return abcitypes.RequestQuery{
		Data: cdc.MustMarshal(
			&protorevtypes.QueryGetProtoRevPoolRequest{
				BaseDenom:  a.BaseAssetDenom,
				OtherDenom: a.QuoteAssetDenom,
			},
		),
		Path: "/osmosis.protorev.v1beta1.Query/GetProtoRevPool",
	}
}

func (a *Asset) SpotPriceRequest(cdc codec.Codec) abcitypes.RequestQuery {
	return abcitypes.RequestQuery{
		Data: cdc.MustMarshal(
			&queryproto.SpotPriceRequest{
				PoolId:          a.PoolID,
				BaseAssetDenom:  a.BaseAssetDenom,
				QuoteAssetDenom: a.QuoteAssetDenom,
			},
		),
		Path: "/osmosis.poolmanager.v1beta1.Query/SpotPrice",
	}
}

func (a *Asset) Exponent() sdkmath.Int {
	i := new(big.Int).Exp(big.NewInt(10), big.NewInt(a.Decimals), nil)
	return sdkmath.NewIntFromBigInt(i)
}

func (a *Asset) Validate() error {
	if err := sdk.ValidateDenom(a.Denom); err != nil {
		return err
	}
	if a.Decimals < 0 {
		return errors.New("decimals must be non-negative")
	}
	if len(a.BaseAssetDenom) == 0 {
		return errors.New("base_asset_denom cannot be empty")
	}
	if len(a.QuoteAssetDenom) == 0 {
		return errors.New("quote_asset_denom cannot be empty")
	}
	if a.BaseAssetDenom == a.QuoteAssetDenom {
		return errors.New("base_asset_denom and quote_asset_denom must be different")
	}

	return nil
}
