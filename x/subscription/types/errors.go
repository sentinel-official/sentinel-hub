package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
)

var (
	ErrInvalidMessage = sdkerrors.Register(ModuleName, 101, "invalid message")

	ErrInsufficient    = sdkerrors.Register(ModuleName, 201, "insufficient resources")
	ErrInvalidResource = sdkerrors.Register(ModuleName, 202, "invalid resource")
	ErrInvalidStatus   = sdkerrors.Register(ModuleName, 203, "invalid status")
	ErrNotFound        = sdkerrors.Register(ModuleName, 204, "not found")
	ErrUnauthorized    = sdkerrors.Register(ModuleName, 205, "unauthorized")
)

// NewErrorAllocationNotFound returns an error indicating that the specified allocation does not exist.
func NewErrorAllocationNotFound(id uint64, addr sdk.AccAddress) error {
	return sdkerrors.Wrapf(ErrNotFound, "allocation %d/%s does not exist", id, addr)
}

// NewErrorInsufficientBytes returns an error indicating that there are insufficient bytes for the specified subscription.
func NewErrorInsufficientBytes(id uint64, bytes sdkmath.Int) error {
	return sdkerrors.Wrapf(ErrInsufficient, "insufficient bytes %s for subscription %d", bytes, id)
}

// NewErrorInvalidAllocation returns an error indicating that the allocation is invalid.
func NewErrorInvalidAllocation(id uint64, addr sdk.AccAddress) error {
	return sdkerrors.Wrapf(ErrInvalidResource, "invalid allocation %d/%s", id, addr)
}

// NewErrorInvalidNodeStatus returns an error indicating that the provided status is invalid for the node.
func NewErrorInvalidNodeStatus(addr base.NodeAddress, status v1base.Status) error {
	return sdkerrors.Wrapf(ErrInvalidStatus, "invalid status %s for node %s", status, addr)
}

// NewErrorInvalidPlanStatus returns an error indicating that the provided status is invalid for the plan.
func NewErrorInvalidPlanStatus(id uint64, status v1base.Status) error {
	return sdkerrors.Wrapf(ErrInvalidStatus, "invalid status %s for plan %d", status, id)
}

// NewErrorInvalidSubscriptionStatus returns an error indicating that the provided status is invalid for the subscription.
func NewErrorInvalidSubscriptionStatus(id uint64, status v1base.Status) error {
	return sdkerrors.Wrapf(ErrInvalidStatus, "invalid status %s for subscription %d", status, id)
}

// NewErrorNodeNotFound returns an error indicating that the specified node does not exist.
func NewErrorNodeNotFound(addr base.NodeAddress) error {
	return sdkerrors.Wrapf(ErrNotFound, "node %s does not exist", addr)
}

// NewErrorPlanNotFound returns an error indicating that the specified plan does not exist.
func NewErrorPlanNotFound(id uint64) error {
	return sdkerrors.Wrapf(ErrNotFound, "plan %d does not exist", id)
}

// NewErrorPriceNotFound returns an error indicating that the price for the specified denomination does not exist.
func NewErrorPriceNotFound(denom string) error {
	return sdkerrors.Wrapf(ErrNotFound, "price for denom %s does not exist", denom)
}

// NewErrorSubscriptionNotFound returns an error indicating that the specified subscription does not exist.
func NewErrorSubscriptionNotFound(id uint64) error {
	return sdkerrors.Wrapf(ErrNotFound, "subscription %d does not exist", id)
}

// NewErrorUnauthorized returns an error indicating that the specified address is not authorized.
func NewErrorUnauthorized(addr string) error {
	return sdkerrors.Wrapf(ErrUnauthorized, "address %s is not authorized", addr)
}
