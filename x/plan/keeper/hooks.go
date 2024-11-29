package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/plan/types/v3"
)

func (k *Keeper) LeaseInactivePreHook(ctx sdk.Context, id uint64) error {
	lease, found := k.GetLease(ctx, id)
	if !found {
		return fmt.Errorf("lease %d does not exist", id)
	}

	nodeAddr, err := base.NodeAddressFromBech32(lease.NodeAddress)
	if err != nil {
		return err
	}

	provAddr, err := base.ProvAddressFromBech32(lease.ProvAddress)
	if err != nil {
		return err
	}

	return k.IteratePlansForNodeByProvider(ctx, nodeAddr, provAddr, func(_ int, item v3.Plan) (bool, error) {
		msg := &v3.MsgUnlinkNodeRequest{
			From:        item.ProvAddress,
			ID:          item.ID,
			NodeAddress: lease.NodeAddress,
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
	return k.IteratePlansForProvider(ctx, addr, func(_ int, item v3.Plan) (bool, error) {
		if !item.Status.Equal(v1base.StatusActive) {
			return false, nil
		}

		msg := &v3.MsgUpdatePlanStatusRequest{
			From:   item.ProvAddress,
			ID:     item.ID,
			Status: v1base.StatusInactive,
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
