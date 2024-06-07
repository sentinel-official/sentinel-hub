package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/session/types"
)

func (k *Keeper) InitGenesis(ctx sdk.Context, state *types.GenesisState) {
	k.SetParams(ctx, state.Params)

	for _, item := range state.Sessions {
		var (
			accAddr  = item.GetAddress()
			nodeAddr = item.GetNodeAddress()
		)

		k.SetSession(ctx, item)
		k.SetSessionForAccount(ctx, accAddr, item.ID)
		k.SetSessionForNode(ctx, nodeAddr, item.ID)
		k.SetSessionForSubscription(ctx, item.SubscriptionID, item.ID)
		k.SetSessionForAllocation(ctx, item.SubscriptionID, accAddr, item.ID)
		k.SetSessionForInactiveAt(ctx, item.InactiveAt, item.ID)
	}

	count := uint64(0)
	for _, item := range state.Sessions {
		if item.ID > count {
			count = item.ID
		}
	}

	k.SetCount(ctx, count)
}

func (k *Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return types.NewGenesisState(
		k.GetSessions(ctx),
		k.GetParams(ctx),
	)
}
