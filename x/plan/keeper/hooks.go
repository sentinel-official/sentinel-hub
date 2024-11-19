package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/plan/types/v3"
)

func (k *Keeper) LeaseInactivePreHook(ctx sdk.Context, id uint64) error {
	lease, found := k.lease.GetLease(ctx, id)
	if !found {
		return nil
	}

	nodeAddr, err := base.NodeAddressFromBech32(lease.NodeAddress)
	if err != nil {
		return err
	}

	provAddr, err := base.ProvAddressFromBech32(lease.ProvAddress)
	if err != nil {
		return err
	}

	k.IteratePlansForNodeByProvider(ctx, nodeAddr, provAddr, func(_ int, item v3.Plan) (stop bool) {
		msg := &v3.MsgUnlinkNodeRequest{
			From:        item.ProvAddress,
			ID:          item.ID,
			NodeAddress: lease.NodeAddress,
		}

		handler := k.router.Handler(msg)
		if _, err := handler(ctx, msg); err != nil {
			panic(err)
		}

		return false
	})

	return nil
}

func (k *Keeper) ProviderInactivePreHook(ctx sdk.Context, addr base.ProvAddress) error {
	k.IteratePlansForProvider(ctx, addr, func(_ int, item v3.Plan) (stop bool) {
		msg := &v3.MsgUpdatePlanStatusRequest{
			From:   item.ProvAddress,
			ID:     item.ID,
			Status: v1base.StatusInactive,
		}

		handler := k.router.Handler(msg)
		if _, err := handler(ctx, msg); err != nil {
			panic(err)
		}

		return false
	})

	return nil
}
