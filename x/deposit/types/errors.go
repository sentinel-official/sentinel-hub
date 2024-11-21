package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	ErrInvalidMessage = sdkerrors.Register(ModuleName, 101, "invalid message")

	ErrInsufficient = sdkerrors.Register(ModuleName, 201, "insufficient resources")
	ErrNotFound     = sdkerrors.Register(ModuleName, 202, "not found")
)

// NewErrorDepositNotFound wraps ErrNotFound with the provided address context for a missing deposit.
func NewErrorDepositNotFound(addr sdk.AccAddress) error {
	return sdkerrors.Wrapf(ErrNotFound, "deposit for address %s does not exist", addr)
}

// NewErrorInsufficientDeposit wraps ErrInsufficient with the provided address context for a deposit issue.
func NewErrorInsufficientDeposit(addr sdk.AccAddress) error {
	return sdkerrors.Wrapf(ErrInsufficient, "insufficient deposit for address %s", addr)
}

// NewErrorInsufficientFunds wraps ErrInsufficient with the provided address context for a funds issue.
func NewErrorInsufficientFunds(addr sdk.AccAddress) error {
	return sdkerrors.Wrapf(ErrInsufficient, "insufficient funds for address %s", addr)
}
