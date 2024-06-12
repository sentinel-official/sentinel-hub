package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	subscriptiontypes "github.com/sentinel-official/hub/v12/x/subscription/types/v2"
)

func (k *Keeper) FundCommunityPool(ctx sdk.Context, fromAddr sdk.AccAddress, coin sdk.Coin) error {
	if coin.IsZero() {
		return nil
	}

	return k.distribution.FundCommunityPool(ctx, sdk.NewCoins(coin), fromAddr)
}

func (k *Keeper) HasProvider(ctx sdk.Context, addr base.ProvAddress) bool {
	return k.provider.HasProvider(ctx, addr)
}

func (k *Keeper) CreateSubscriptionForNode(ctx sdk.Context, accAddr sdk.AccAddress, nodeAddr base.NodeAddress, gigabytes, hours int64, denom string) (*subscriptiontypes.NodeSubscription, error) {
	return k.subscription.CreateSubscriptionForNode(ctx, accAddr, nodeAddr, gigabytes, hours, denom)
}
