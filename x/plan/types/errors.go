package types

import (
	sdkerrors "cosmossdk.io/errors"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
)

var (
	ErrorInvalidMessage = sdkerrors.Register(ModuleName, 101, "invalid message")

	ErrorDuplicateNode     = sdkerrors.Register(ModuleName, 201, "duplicate node")
	ErrorInvalidNodeStatus = sdkerrors.Register(ModuleName, 202, "invalid node status")
	ErrorLeaseNotFound     = sdkerrors.Register(ModuleName, 203, "lease not found")
	ErrorNodeNotFound      = sdkerrors.Register(ModuleName, 204, "node not found")
	ErrorPlanNotFound      = sdkerrors.Register(ModuleName, 205, "plan not found")
	ErrorProviderNotFound  = sdkerrors.Register(ModuleName, 206, "provider not found")
	ErrorUnauthorized      = sdkerrors.Register(ModuleName, 207, "unauthorized")
)

func NewErrorDuplicateNodeForPlan(id uint64, addr base.NodeAddress) error {
	return sdkerrors.Wrapf(ErrorDuplicateNode, "node %s for plan %d already exists", addr, id)
}

func NewErrorInvalidNodeStatus(addr base.NodeAddress, status v1base.Status) error {
	return sdkerrors.Wrapf(ErrorInvalidNodeStatus, "invalid status %s for node %s", status, addr)
}

func NewErrorLeaseNotFound(nodeAddr base.NodeAddress, provAddr base.ProvAddress) error {
	return sdkerrors.Wrapf(ErrorLeaseNotFound, "lease for node %s by provider %s does not exist", nodeAddr, provAddr)
}

func NewErrorNodeForPlanNotFound(id uint64, addr base.NodeAddress) error {
	return sdkerrors.Wrapf(ErrorNodeNotFound, "node %s for plan %d does not exist", addr, id)
}

func NewErrorNodeNotFound(addr base.NodeAddress) error {
	return sdkerrors.Wrapf(ErrorNodeNotFound, "node %s does not exist", addr)
}

func NewErrorPlanNotFound(id uint64) error {
	return sdkerrors.Wrapf(ErrorPlanNotFound, "plan %d does not exist", id)
}

func NewErrorProviderNotFound(addr base.ProvAddress) error {
	return sdkerrors.Wrapf(ErrorProviderNotFound, "provider %s does not exist", addr)
}

func NewErrorUnauthorized(addr string) error {
	return sdkerrors.Wrapf(ErrorUnauthorized, "address %s is not authorized", addr)
}
