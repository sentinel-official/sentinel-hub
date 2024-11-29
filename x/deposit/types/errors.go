package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	ErrInvalidMessage = sdkerrors.Register(ModuleName, 101, "invalid message")

	ErrDepositNotFound     = sdkerrors.Register(ModuleName, 201, "deposit not found")
	ErrInsufficientDeposit = sdkerrors.Register(ModuleName, 202, "insufficient deposit")
	ErrInsufficientFunds   = sdkerrors.Register(ModuleName, 203, "insufficient funds")
)

// NewErrorDepositNotFound wraps ErrDepositNotFound with the provided address context for a missing deposit.
func NewErrorDepositNotFound(addr sdk.AccAddress) error {
	return sdkerrors.Wrapf(ErrDepositNotFound, "deposit for address %s does not exist", addr)
}

// NewErrorInsufficientDeposit wraps ErrInsufficientDeposit with the provided address context for a deposit issue.
func NewErrorInsufficientDeposit(addr sdk.AccAddress) error {
	return sdkerrors.Wrapf(ErrInsufficientDeposit, "insufficient deposit for address %s", addr)
}

// NewErrorInsufficientFunds wraps ErrInsufficientFunds with the provided address context for a funds issue.
func NewErrorInsufficientFunds(addr sdk.AccAddress) error {
	return sdkerrors.Wrapf(ErrInsufficientFunds, "insufficient funds for address %s", addr)
}
