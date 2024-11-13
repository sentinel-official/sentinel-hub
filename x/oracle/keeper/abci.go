package keeper

import (
	abcitypes "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ibchost "github.com/cosmos/ibc-go/v7/modules/core/24-host"

	"github.com/sentinel-official/hub/v12/x/oracle/types/v1"
)

func (k *Keeper) EndBlock(ctx sdk.Context) []abcitypes.ValidatorUpdate {
	portID := k.GetPortID(ctx)
	channelID := k.GetChannelID(ctx)
	timeout := k.GetQueryTimeout(ctx)

	channelCap, found := k.capability.GetCapability(ctx, ibchost.ChannelCapabilityPath(portID, channelID))
	if !found {
		return nil
	}

	k.IterateAssets(ctx, func(_ int, item v1.Asset) bool {
		sequence, err := k.SendQueryPacket(
			ctx, channelCap, portID, channelID, uint64(timeout),
			item.ProtoRevPoolRequest(k.cdc),
			item.SpotPriceRequest(k.cdc),
		)
		if err != nil {
			panic(err)
		}

		k.SetDenomForPacket(ctx, portID, channelID, sequence, item.Denom)
		return false
	})

	return nil
}
