package keeper

import (
	"fmt"

	sdkmath "cosmossdk.io/math"
	abcitypes "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	ibcicqtypes "github.com/cosmos/ibc-apps/modules/async-icq/v7/types"
	ibcclienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	ibcchanneltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"

	"github.com/sentinel-official/hub/v12/third_party/osmosis/x/poolmanager/client/queryproto"
	protorevtypes "github.com/sentinel-official/hub/v12/third_party/osmosis/x/protorev/types"
)

// SendQueryPacket serializes query requests and sends them as an IBC packet to a destination chain.
func (k *Keeper) SendQueryPacket(
	ctx sdk.Context, channelCap *capabilitytypes.Capability, portID, channelID string, timeout uint64,
	reqs ...abcitypes.RequestQuery,
) (uint64, error) {
	// Serialize the Cosmos query requests into binary format.
	data, err := ibcicqtypes.SerializeCosmosQuery(reqs)
	if err != nil {
		return 0, err
	}

	// Create packet data with the serialized queries and validate it.
	packetData := ibcicqtypes.InterchainQueryPacketData{Data: data}
	if err := packetData.ValidateBasic(); err != nil {
		return 0, err
	}

	// Use the ICS-04 interface to send the packet over IBC.
	return k.ics4.SendPacket(
		ctx, channelCap, portID, channelID, ibcclienttypes.ZeroHeight(), timeout, packetData.GetBytes(),
	)
}

// OnAcknowledgementPacket processes the acknowledgement packet received after sending an IBC query packet.
// It deserializes the packet data and acknowledgement, validates the response count, and updates the relevant asset.
func (k *Keeper) OnAcknowledgementPacket(
	ctx sdk.Context, packet ibcchanneltypes.Packet, ack ibcchanneltypes.Acknowledgement,
) error {
	// Retrieve the source port, channel, and sequence number from the packet.
	portID := packet.GetSourcePort()
	channelID := packet.GetSourceChannel()
	sequence := packet.GetSequence()

	// Ensure the denom mapping for the packet is deleted after processing the acknowledgement.
	defer k.DeleteDenomForPacket(ctx, portID, channelID, sequence)

	// If the acknowledgement indicates failure, there's no further processing required.
	if !ack.Success() {
		return nil
	}

	// Unmarshal the packet data to get the interchain query packet details.
	var packetData ibcicqtypes.InterchainQueryPacketData
	if err := k.cdc.UnmarshalJSON(packet.GetData(), &packetData); err != nil {
		return err
	}

	// Deserialize the Cosmos queries from the packet data.
	reqs, err := ibcicqtypes.DeserializeCosmosQuery(packetData.Data)
	if err != nil {
		return err
	}

	// Unmarshal the acknowledgement result to obtain the query responses.
	var packetAck ibcicqtypes.InterchainQueryPacketAck
	if err := k.cdc.UnmarshalJSON(ack.GetResult(), &packetAck); err != nil {
		return err
	}

	// Deserialize the Cosmos responses from the acknowledgement data.
	resps, err := ibcicqtypes.DeserializeCosmosResponse(packetAck.Data)
	if err != nil {
		return err
	}

	// Verify that the number of responses matches the number of requests.
	if len(reqs) != len(resps) {
		return fmt.Errorf("invalid response count %d; expected %d", len(resps), len(reqs))
	}

	// Retrieve the asset associated with the packet using port, channel, and sequence information.
	asset, err := k.GetAssetForPacket(ctx, portID, channelID, sequence)
	if err != nil {
		return err
	}

	// Iterate through each request-response pair and update the asset accordingly.
	for i := 0; i < len(reqs); i++ {
		// Skip updates if the response height is older than the current asset height.
		if resps[i].GetHeight() < asset.Height {
			return nil
		}

		// Handle specific query paths to extract the required data and update the asset.
		switch reqs[i].Path {
		case "/osmosis.poolmanager.v1beta1.Query/SpotPrice":
			// Extract the spot price from the response and update the asset price.
			var res queryproto.SpotPriceResponse
			if err := k.cdc.Unmarshal(resps[i].GetValue(), &res); err != nil {
				return err
			}

			spotPrice, err := sdkmath.LegacyNewDecFromStr(res.GetSpotPrice())
			if err != nil {
				return err
			}

			// Update the asset price using the spot price and its exponent.
			asset.Price = spotPrice.MulInt(asset.Exponent()).TruncateInt()
		case "/osmosis.protorev.v1beta1.Query/GetProtoRevPool":
			// Extract the pool ID from the response and update the asset pool ID.
			var res protorevtypes.QueryGetProtoRevPoolResponse
			if err := k.cdc.Unmarshal(resps[i].GetValue(), &res); err != nil {
				return err
			}

			asset.PoolID = res.GetPoolId()
		}
	}

	// Persist the updated asset information in the store.
	k.SetAsset(ctx, asset)

	return nil
}
