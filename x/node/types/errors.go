package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
)

var (
	ErrInvalidMessage = sdkerrors.Register(ModuleName, 101, "invalid message")

	ErrDuplicateEntry  = sdkerrors.Register(ModuleName, 201, "duplicate entry")
	ErrInvalidResource = sdkerrors.Register(ModuleName, 202, "invalid resource")
	ErrInvalidStatus   = sdkerrors.Register(ModuleName, 203, "invalid status")
	ErrNotFound        = sdkerrors.Register(ModuleName, 204, "not found")
	ErrUnauthorized    = sdkerrors.Register(ModuleName, 205, "unauthorized")
)

// NewErrorDuplicateNode returns an error indicating that the specified node already exists.
func NewErrorDuplicateNode(addr base.NodeAddress) error {
	return sdkerrors.Wrapf(ErrDuplicateEntry, "node %s already exists", addr)
}

// NewErrorInvalidGigabytes returns an error indicating that the provided gigabytes value is invalid.
func NewErrorInvalidGigabytes(gigabytes int64) error {
	return sdkerrors.Wrapf(ErrInvalidResource, "invalid gigabytes %d", gigabytes)
}

// NewErrorInvalidHours returns an error indicating that the provided hours value is invalid.
func NewErrorInvalidHours(hours int64) error {
	return sdkerrors.Wrapf(ErrInvalidResource, "invalid hours %d", hours)
}

// NewErrorInvalidPrices returns an error indicating that the provided prices are invalid.
func NewErrorInvalidPrices(prices sdk.Coins) error {
	return sdkerrors.Wrapf(ErrInvalidResource, "invalid prices %s", prices)
}

// NewErrorInvalidNodeStatus returns an error indicating that the provided status is invalid for the given node.
func NewErrorInvalidNodeStatus(addr base.NodeAddress, status v1base.Status) error {
	return sdkerrors.Wrapf(ErrInvalidStatus, "invalid status %s for node %s", status, addr)
}

// NewErrorNodeNotFound returns an error indicating that the specified node does not exist.
func NewErrorNodeNotFound(addr base.NodeAddress) error {
	return sdkerrors.Wrapf(ErrNotFound, "node %s does not exist", addr)
}

// NewErrorPriceNotFound returns an error indicating that the price for the specified denom does not exist.
func NewErrorPriceNotFound(denom string) error {
	return sdkerrors.Wrapf(ErrNotFound, "price for denom %s does not exist", denom)
}

// NewErrorUnauthorized returns an error indicating that the specified address is not authorized.
func NewErrorUnauthorized(addr string) error {
	return sdkerrors.Wrapf(ErrUnauthorized, "address %s is not authorized", addr)
}
