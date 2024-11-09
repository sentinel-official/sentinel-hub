package keeper

import (
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

func (k *Keeper) OnAcknowledgementPacket(
	ctx sdk.Context, packet ibcchanneltypes.Packet, ack ibcchanneltypes.Acknowledgement,
) error {
	if !ack.Success() {
		return nil
	}

	var packetData ibcicqtypes.InterchainQueryPacketData
	if err := k.cdc.UnmarshalJSON(packet.GetData(), &packetData); err != nil {
		return err
	}

	reqs, err := ibcicqtypes.DeserializeCosmosQuery(packetData.Data)
	if err != nil {
		return err
	}

	var packetAck ibcicqtypes.InterchainQueryPacketAck
	if err := k.cdc.UnmarshalJSON(ack.GetResult(), &packetAck); err != nil {
		return err
	}

	resps, err := ibcicqtypes.DeserializeCosmosResponse(packetAck.Data)
	if err != nil {
		return err
	}

	portID := packet.GetSourcePort()
	channelID := packet.GetSourceChannel()
	sequence := packet.GetSequence()

	for i := 0; i < len(reqs); i++ {
		switch reqs[i].Path {
		case "/osmosis.poolmanager.v1beta1.Query/SpotPrice":
			if err := k.processSpotPriceResponse(ctx, portID, channelID, sequence, resps[i].GetValue()); err != nil {
				return err
			}
		case "/osmosis.protorev.v1beta1.Query/GetProtoRevPool":
			if err := k.processProtoRevPoolResponse(ctx, portID, channelID, sequence, resps[i].GetValue()); err != nil {
				return err
			}
		}
	}

	k.DeleteDenomForPacket(ctx, portID, channelID, sequence)
	return nil
}

func (k *Keeper) processSpotPriceResponse(ctx sdk.Context, portID, channelID string, sequence uint64, buf []byte) error {
	var res queryproto.SpotPriceResponse
	if err := k.cdc.Unmarshal(buf, &res); err != nil {
		return err
	}

	spotPrice, err := sdkmath.LegacyNewDecFromStr(res.GetSpotPrice())
	if err != nil {
		return err
	}

	asset, err := k.GetAssetForPacket(ctx, portID, channelID, sequence)
	if err != nil {
		return err
	}

	asset.Price = spotPrice.MulInt(asset.Exponent()).TruncateInt()
	k.SetAsset(ctx, asset)

	return nil
}

func (k *Keeper) processProtoRevPoolResponse(ctx sdk.Context, portID, channelID string, sequence uint64, buf []byte) error {
	var res protorevtypes.QueryGetProtoRevPoolResponse
	if err := k.cdc.Unmarshal(buf, &res); err != nil {
		return err
	}

	asset, err := k.GetAssetForPacket(ctx, portID, channelID, sequence)
	if err != nil {
		return err
	}

	asset.PoolID = res.GetPoolId()
	k.SetAsset(ctx, asset)

	return nil
}
