package keeper

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/oracle/types"
	"github.com/sentinel-official/hub/v12/x/oracle/types/v1"
)

func (k *Keeper) HandleMsgAddAsset(ctx sdk.Context, msg *v1.MsgAddAssetRequest) (*v1.MsgAddAssetResponse, error) {
	authority := k.GetAuthority()
	if msg.From != authority {
		return nil, types.NewErrorInvalidSigner(msg.From, authority)
	}

	if k.HasAsset(ctx, msg.Denom) {
		return nil, types.NewErrorDuplicateAsset(msg.Denom)
	}

	asset := v1.Asset{
		Denom:           msg.Denom,
		Decimals:        msg.Decimals,
		PoolID:          0,
		BaseAssetDenom:  msg.BaseAssetDenom,
		QuoteAssetDenom: msg.QuoteAssetDenom,
		Price:           sdkmath.ZeroInt(),
	}

	k.SetAsset(ctx, asset)

	return &v1.MsgAddAssetResponse{}, nil
}

func (k *Keeper) HandleMsgDeleteAsset(ctx sdk.Context, msg *v1.MsgDeleteAssetRequest) (*v1.MsgDeleteAssetResponse, error) {
	authority := k.GetAuthority()
	if msg.From != authority {
		return nil, types.NewErrorInvalidSigner(msg.From, authority)
	}

	k.DeleteAsset(ctx, msg.Denom)

	return &v1.MsgDeleteAssetResponse{}, nil
}

func (k *Keeper) HandleMsgUpdateAsset(ctx sdk.Context, msg *v1.MsgUpdateAssetRequest) (*v1.MsgUpdateAssetResponse, error) {
	authority := k.GetAuthority()
	if msg.From != authority {
		return nil, types.NewErrorInvalidSigner(msg.From, authority)
	}

	asset, found := k.GetAsset(ctx, msg.Denom)
	if !found {
		return nil, types.NewErrorAssetNotFound(msg.Denom)
	}

	asset.Decimals = msg.Decimals
	asset.BaseAssetDenom = msg.BaseAssetDenom
	asset.QuoteAssetDenom = msg.QuoteAssetDenom

	k.SetAsset(ctx, asset)

	return &v1.MsgUpdateAssetResponse{}, nil
}

func (k *Keeper) HandleMsgUpdateParams(ctx sdk.Context, msg *v1.MsgUpdateParamsRequest) (*v1.MsgUpdateParamsResponse, error) {
	authority := k.GetAuthority()
	if msg.From != authority {
		return nil, types.NewErrorInvalidSigner(msg.From, authority)
	}

	k.SetParams(ctx, msg.Params)
	return &v1.MsgUpdateParamsResponse{}, nil
}
