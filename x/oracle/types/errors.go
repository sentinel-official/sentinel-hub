package types

import (
	sdkerrors "cosmossdk.io/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	ibcchanneltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	ibcporttypes "github.com/cosmos/ibc-go/v7/modules/core/05-port/types"
	ibcerrors "github.com/cosmos/ibc-go/v7/modules/core/errors"
)

var (
	ErrorInvalidMessage = sdkerrors.Register(ModuleName, 101, "invalid message")

	ErrorAssetNotFound  = sdkerrors.Register(ModuleName, 201, "asset not found")
	ErrorDenomNotFound  = sdkerrors.Register(ModuleName, 202, "denom not found")
	ErrorDuplicateAsset = sdkerrors.Register(ModuleName, 203, "duplicate asset")
)

func NewErrorInvalidVersion(version, expected string) error {
	return sdkerrors.Wrapf(ibcerrors.ErrInvalidVersion, "invalid version %s; expected %s", version, expected)
}

func NewErrorInvalidCounterpartyVersion(version, expected string) error {
	return sdkerrors.Wrapf(ibcerrors.ErrInvalidVersion, "invalid counteryparty version %s; expected %s", version, expected)
}

func NewErrorInvalidChannelOrdering(order, expected ibcchanneltypes.Order) error {
	return sdkerrors.Wrapf(ibcchanneltypes.ErrInvalidChannelOrdering, "invalid channel order %s; expected %s", order, expected)
}

func NewErrorInvalidPort(portID, expected string) error {
	return sdkerrors.Wrapf(ibcporttypes.ErrInvalidPort, "invalid port %s; expected %s", portID, expected)
}

func NewErrorAssetNotFound(denom string) error {
	return sdkerrors.Wrapf(ErrorAssetNotFound, "asset %s does not exist", denom)
}

func NewErrorDenomtNotFound(portID, channelID string, sequence uint64) error {
	return sdkerrors.Wrapf(ErrorAssetNotFound, "denom for packet %s/%s/%d does not exist", portID, channelID, sequence)
}

func NewErrorInvalidSigner(from, expected string) error {
	return sdkerrors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority %s; expected %s", from, expected)
}

func NewErrorDuplicateAsset(denom string) error {
	return sdkerrors.Wrapf(ErrorDuplicateAsset, "asset %s already exists", denom)
}
