package types

import (
	sdkerrors "cosmossdk.io/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	ibcporttypes "github.com/cosmos/ibc-go/v7/modules/core/05-port/types"
	ibcerrors "github.com/cosmos/ibc-go/v7/modules/core/errors"
)

var (
	ErrInvalidMessage = sdkerrors.Register(ModuleName, 101, "invalid message")

	ErrDuplicateEntry = sdkerrors.Register(ModuleName, 201, "duplicate entry")
	ErrNotFound       = sdkerrors.Register(ModuleName, 202, "not found")
)

// NewErrorAssetNotFound returns an error indicating that the specified asset does not exist.
func NewErrorAssetNotFound(denom string) error {
	return sdkerrors.Wrapf(ErrNotFound, "asset %s does not exist", denom)
}

// NewErrorDenomForPacketNotFound returns an error indicating that the denom for the specified packet does not exist.
func NewErrorDenomForPacketNotFound(portID, channelID string, sequence uint64) error {
	return sdkerrors.Wrapf(ErrNotFound, "denom for packet %s/%s/%d does not exist", portID, channelID, sequence)
}

// NewErrorDuplicateAsset returns an error indicating that the specified asset already exists.
func NewErrorDuplicateAsset(denom string) error {
	return sdkerrors.Wrapf(ErrDuplicateEntry, "asset %s already exists", denom)
}

// NewErrorInvalidCounterpartyVersion returns an error indicating that the counterparty version is invalid.
func NewErrorInvalidCounterpartyVersion(version, expected string) error {
	return sdkerrors.Wrapf(ibcerrors.ErrInvalidVersion, "invalid counterparty version %s; expected %s", version, expected)
}

// NewErrorInvalidPort returns an error indicating that the provided port is invalid.
func NewErrorInvalidPort(portID, expected string) error {
	return sdkerrors.Wrapf(ibcporttypes.ErrInvalidPort, "invalid port %s; expected %s", portID, expected)
}

// NewErrorInvalidSigner returns an error indicating that the provided signer is invalid.
func NewErrorInvalidSigner(from, expected string) error {
	return sdkerrors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority %s; expected %s", from, expected)
}

// NewErrorInvalidVersion returns an error indicating that the provided IBC version is invalid.
func NewErrorInvalidVersion(version, expected string) error {
	return sdkerrors.Wrapf(ibcerrors.ErrInvalidVersion, "invalid version %s; expected %s", version, expected)
}
