package v1

import (
	"errors"
	"math/big"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (a *Asset) Multiplier() sdkmath.Int {
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
