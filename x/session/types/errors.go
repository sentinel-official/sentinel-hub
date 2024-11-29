package types

import (
	sdkerrors "cosmossdk.io/errors"

	v1base "github.com/sentinel-official/hub/v12/types/v1"
)

var (
	ErrInvalidMessage = sdkerrors.Register(ModuleName, 101, "invalid message")

	ErrInvalidSessionStatus = sdkerrors.Register(ModuleName, 201, "invalid session status")
	ErrInvalidSignature     = sdkerrors.Register(ModuleName, 202, "invalid signature")
	ErrSessionNotFound      = sdkerrors.Register(ModuleName, 203, "session not found")
	ErrUnauthorized         = sdkerrors.Register(ModuleName, 204, "unauthorized")
)

// NewErrorInvalidSessionStatus returns an error indicating that the provided status is invalid for the session.
func NewErrorInvalidSessionStatus(id uint64, status v1base.Status) error {
	return sdkerrors.Wrapf(ErrInvalidSessionStatus, "invalid status %s for session %d", status, id)
}

// NewErrorInvalidSignature returns an error indicating that the provided signature is invalid.
func NewErrorInvalidSignature(signature []byte) error {
	return sdkerrors.Wrapf(ErrInvalidSignature, "invalid signature %X", signature)
}

// NewErrorSessionNotFound returns an error indicating that the specified session does not exist.
func NewErrorSessionNotFound(id uint64) error {
	return sdkerrors.Wrapf(ErrSessionNotFound, "session %d does not exist", id)
}

// NewErrorUnauthorized returns an error indicating that the specified address is not authorized.
func NewErrorUnauthorized(addr string) error {
	return sdkerrors.Wrapf(ErrUnauthorized, "address %s is not authorized", addr)
}
