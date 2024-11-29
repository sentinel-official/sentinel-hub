package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/lease/types/v1"
)

func (k *Keeper) NodeInactivePreHook(ctx sdk.Context, addr base.NodeAddress) error {
	return k.IterateLeasesForNode(ctx, addr, func(_ int, item v1.Lease) (bool, error) {
		msg := &v1.MsgEndLeaseRequest{
			From: item.ProvAddress,
			ID:   item.ID,
		}

		handler := k.router.Handler(msg)
		if handler == nil {
			return false, fmt.Errorf("nil handler for message route: %s", sdk.MsgTypeURL(msg))
		}

		resp, err := handler(ctx, msg)
		if err != nil {
			return false, err
		}

		ctx.EventManager().EmitEvents(resp.GetEvents())
		return false, nil
	})
}

func (k *Keeper) ProviderInactivePreHook(ctx sdk.Context, addr base.ProvAddress) error {
	return k.IterateLeasesForProvider(ctx, addr, func(_ int, item v1.Lease) (bool, error) {
		msg := &v1.MsgEndLeaseRequest{
			From: item.ProvAddress,
			ID:   item.ID,
		}

		handler := k.router.Handler(msg)
		if handler == nil {
			return false, fmt.Errorf("nil handler for message route: %s", sdk.MsgTypeURL(msg))
		}

		resp, err := handler(ctx, msg)
		if err != nil {
			return false, err
		}

		ctx.EventManager().EmitEvents(resp.GetEvents())
		return false, nil
	})
}
