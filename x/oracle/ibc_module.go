package oracle

import (
	sdkerrors "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	ibcchanneltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	ibcporttypes "github.com/cosmos/ibc-go/v7/modules/core/05-port/types"
	ibchost "github.com/cosmos/ibc-go/v7/modules/core/24-host"
	ibcerrors "github.com/cosmos/ibc-go/v7/modules/core/errors"
	ibcexported "github.com/cosmos/ibc-go/v7/modules/core/exported"

	"github.com/sentinel-official/hub/v12/x/oracle/keeper"
	"github.com/sentinel-official/hub/v12/x/oracle/types"
)

var (
	_ ibcporttypes.IBCModule = IBCModule{}
)

type IBCModule struct {
	cdc    codec.Codec
	keeper keeper.Keeper
}

func (im IBCModule) OnChanOpenInit(
	ctx sdk.Context, _ ibcchanneltypes.Order, _ []string, portID, channelID string,
	channelCap *capabilitytypes.Capability, _ ibcchanneltypes.Counterparty, version string,
) (string, error) {
	if version != types.Version {
		return "", types.NewErrorInvalidVersion(version, types.Version)
	}

	boundPortID := im.keeper.GetPortID(ctx)
	if portID != boundPortID {
		return "", types.NewErrorInvalidPort(portID, boundPortID)
	}

	capPath := ibchost.ChannelCapabilityPath(portID, channelID)
	if err := im.keeper.ClaimCapability(ctx, channelCap, capPath); err != nil {
		return "", err
	}

	return types.Version, nil
}

func (im IBCModule) OnChanOpenTry(
	ctx sdk.Context, _ ibcchanneltypes.Order, _ []string, portID, channelID string,
	channelCap *capabilitytypes.Capability, _ ibcchanneltypes.Counterparty, counterpartyVersion string,
) (string, error) {
	if counterpartyVersion != types.Version {
		return "", types.NewErrorInvalidCounterpartyVersion(counterpartyVersion, types.Version)
	}

	boundPortID := im.keeper.GetPortID(ctx)
	if portID != boundPortID {
		return "", types.NewErrorInvalidPort(portID, boundPortID)
	}

	capPath := ibchost.ChannelCapabilityPath(portID, channelID)
	if !im.keeper.AuthenticateCapability(ctx, channelCap, capPath) {
		if err := im.keeper.ClaimCapability(ctx, channelCap, capPath); err != nil {
			return "", err
		}
	}

	return types.Version, nil
}

func (im IBCModule) OnChanOpenAck(_ sdk.Context, _, _, _ string, counterpartyVersion string) error {
	if counterpartyVersion != types.Version {
		return types.NewErrorInvalidCounterpartyVersion(counterpartyVersion, types.Version)
	}

	return nil
}

func (im IBCModule) OnChanOpenConfirm(_ sdk.Context, _, _ string) error {
	return nil
}

func (im IBCModule) OnChanCloseInit(_ sdk.Context, _, _ string) error {
	return sdkerrors.Wrap(ibcerrors.ErrInvalidRequest, "user cannot close the channel")
}

func (im IBCModule) OnChanCloseConfirm(_ sdk.Context, _, _ string) error {
	return nil
}

func (im IBCModule) OnRecvPacket(_ sdk.Context, _ ibcchanneltypes.Packet, _ sdk.AccAddress) ibcexported.Acknowledgement {
	return ibcchanneltypes.NewErrorAcknowledgement(sdkerrors.Wrap(ibcerrors.ErrInvalidRequest, "oracle module can not receive the packets"))
}

func (im IBCModule) OnAcknowledgementPacket(
	ctx sdk.Context, packet ibcchanneltypes.Packet, acknowledgement []byte, _ sdk.AccAddress,
) error {
	var ack ibcchanneltypes.Acknowledgement
	if err := im.cdc.UnmarshalJSON(acknowledgement, &ack); err != nil {
		return err
	}

	return im.keeper.OnAcknowledgementPacket(ctx, packet, ack)
}

func (im IBCModule) OnTimeoutPacket(ctx sdk.Context, packet ibcchanneltypes.Packet, _ sdk.AccAddress) error {
	return im.keeper.OnTimeoutPacket(ctx, packet)
}
