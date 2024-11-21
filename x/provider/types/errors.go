package types

import (
	sdkerrors "cosmossdk.io/errors"

	base "github.com/sentinel-official/hub/v12/types"
)

var (
	ErrInvalidMessage = sdkerrors.Register(ModuleName, 101, "invalid message")

	ErrDuplicateEntry = sdkerrors.Register(ModuleName, 201, "duplicate entry")
	ErrNotFound       = sdkerrors.Register(ModuleName, 202, "not found")
	ErrUnauthorized   = sdkerrors.Register(ModuleName, 203, "unauthorized")
)

// NewErrorDuplicateProvider returns an error indicating that the specified provider already exists.
func NewErrorDuplicateProvider(addr base.ProvAddress) error {
	return sdkerrors.Wrapf(ErrDuplicateEntry, "provider %s already exists", addr)
}

// NewErrorProviderNotFound returns an error indicating that the specified provider does not exist.
func NewErrorProviderNotFound(addr base.ProvAddress) error {
	return sdkerrors.Wrapf(ErrNotFound, "provider %s does not exist", addr)
}

// NewErrorUnauthorized returns an error indicating that the specified address is not authorized.
func NewErrorUnauthorized(addr string) error {
	return sdkerrors.Wrapf(ErrUnauthorized, "address %s is not authorized", addr)
}
