package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
)

func (k *Keeper) PlanUnlinkNodePreHook(ctx sdk.Context, id uint64, addr base.NodeAddress) error {
	if err := k.session.PlanUnlinkNodePreHook(ctx, id, addr); err != nil {
		return err
	}

	return nil
}
