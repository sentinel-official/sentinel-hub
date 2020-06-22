package node

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/dvpn/node/keeper"
	"github.com/sentinel-official/hub/x/dvpn/node/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state types.GenesisState) {
	for _, node := range state {
		k.SetNode(ctx, node)
		k.SetNodeAddressForProvider(ctx, node.Provider, node.Address)
	}
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	return k.GetNodes(ctx)
}

func ValidateGenesis(state types.GenesisState) error {
	return nil
}