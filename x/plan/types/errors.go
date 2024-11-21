package types

import (
	sdkerrors "cosmossdk.io/errors"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
)

var (
	ErrInvalidMessage = sdkerrors.Register(ModuleName, 101, "invalid message")

	ErrDuplicateEntry = sdkerrors.Register(ModuleName, 201, "duplicate entry")
	ErrInvalidStatus  = sdkerrors.Register(ModuleName, 202, "invalid status")
	ErrNotFound       = sdkerrors.Register(ModuleName, 203, "not found")
	ErrUnauthorized   = sdkerrors.Register(ModuleName, 204, "unauthorized")
)

// NewErrorDuplicateNodeForPlan returns an error indicating that a node already exists for the specified plan.
func NewErrorDuplicateNodeForPlan(id uint64, addr base.NodeAddress) error {
	return sdkerrors.Wrapf(ErrDuplicateEntry, "node %s for plan %d already exists", addr, id)
}

// NewErrorInvalidNodeStatus returns an error indicating that the provided status is invalid for the given node.
func NewErrorInvalidNodeStatus(addr base.NodeAddress, status v1base.Status) error {
	return sdkerrors.Wrapf(ErrInvalidStatus, "invalid status %s for node %s", status, addr)
}

// NewErrorLeaseNotFound returns an error indicating that the specified lease does not exist.
func NewErrorLeaseNotFound(nodeAddr base.NodeAddress, provAddr base.ProvAddress) error {
	return sdkerrors.Wrapf(ErrNotFound, "lease for node %s by provider %s does not exist", nodeAddr, provAddr)
}

// NewErrorNodeForPlanNotFound returns an error indicating that the specified node does not exist for the plan.
func NewErrorNodeForPlanNotFound(id uint64, addr base.NodeAddress) error {
	return sdkerrors.Wrapf(ErrNotFound, "node %s for plan %d does not exist", addr, id)
}

// NewErrorNodeNotFound returns an error indicating that the specified node does not exist.
func NewErrorNodeNotFound(addr base.NodeAddress) error {
	return sdkerrors.Wrapf(ErrNotFound, "node %s does not exist", addr)
}

// NewErrorPlanNotFound returns an error indicating that the specified plan does not exist.
func NewErrorPlanNotFound(id uint64) error {
	return sdkerrors.Wrapf(ErrNotFound, "plan %d does not exist", id)
}

// NewErrorProviderNotFound returns an error indicating that the specified provider does not exist.
func NewErrorProviderNotFound(addr base.ProvAddress) error {
	return sdkerrors.Wrapf(ErrNotFound, "provider %s does not exist", addr)
}

// NewErrorUnauthorized returns an error indicating that the specified address is not authorized.
func NewErrorUnauthorized(addr string) error {
	return sdkerrors.Wrapf(ErrUnauthorized, "address %s is not authorized", addr)
}
