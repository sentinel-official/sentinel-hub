package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/node/types/v3"
)

func (k *Keeper) handleInactiveNodes(ctx sdk.Context) {
	k.IterateNodesForInactiveAt(ctx, ctx.BlockTime(), func(_ int, item v3.Node) bool {
		nodeAddr, err := base.NodeAddressFromBech32(item.Address)
		if err != nil {
			panic(err)
		}

		k.DeleteNodeForInactiveAt(ctx, item.InactiveAt, nodeAddr)

		msg := &v3.MsgUpdateNodeStatusRequest{
			From:   item.Address,
			Status: v1base.StatusInactive,
		}

		handler := k.router.Handler(msg)
		if handler == nil {
			panic(fmt.Errorf("nil handler for message route: %s", sdk.MsgTypeURL(msg)))
		}

		resp, err := handler(ctx, msg)
		if err != nil {
			panic(err)
		}

		ctx.EventManager().EmitEvents(resp.GetEvents())
		return false
	})
}

func (k *Keeper) BeginBlock(ctx sdk.Context) {
	k.handleInactiveNodes(ctx)
}
