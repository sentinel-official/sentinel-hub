package keeper

import (
	abcitypes "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ibchost "github.com/cosmos/ibc-go/v7/modules/core/24-host"

	protorevtypes "github.com/sentinel-official/hub/v12/third_party/osmosis/x/protorev/types"
	"github.com/sentinel-official/hub/v12/x/oracle/types/v1"
)

// EndBlock is called at the end of each block to trigger IBC query packets for relevant assets.
func (k *Keeper) EndBlock(ctx sdk.Context) []abcitypes.ValidatorUpdate {
	// Retrieve necessary information for the IBC packet.
	portID := k.GetPortID(ctx)
	channelID := k.GetChannelID(ctx)
	timeout := k.GetQueryTimeout(ctx)

	// Get the channel capability to ensure we have the authority to send packets.
	channelCap, found := k.capability.GetCapability(ctx, ibchost.ChannelCapabilityPath(portID, channelID))
	if !found {
		return nil
	}

	// Iterate over each asset and send a ProtoRevPool query for each.
	k.IterateAssets(ctx, func(_ int, item v1.Asset) bool {
		// Create a request for the GetProtoRevPool query using asset details.
		req := abcitypes.RequestQuery{
			Data: k.cdc.MustMarshal(
				&protorevtypes.QueryGetProtoRevPoolRequest{
					BaseDenom:  item.BaseAssetDenom,
					OtherDenom: item.QuoteAssetDenom,
				},
			),
			Path: "/osmosis.protorev.v1beta1.Query/GetProtoRevPool",
		}

		// Send the GetProtoRevPool query packet over IBC.
		sequence, err := k.SendQueryPacket(ctx, channelCap, portID, channelID, uint64(timeout), req)
		if err != nil {
			panic(err)
		}

		// Map the sequence number to the asset denom for tracking.
		k.SetDenomForPacket(ctx, portID, channelID, sequence, item.Denom)
		return false
	})

	// No validator updates in this function.
	return nil
}
