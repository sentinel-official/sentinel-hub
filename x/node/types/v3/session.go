package v3

import (
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	sessiontypes "github.com/sentinel-official/hub/v12/x/session/types/v3"
)

var _ sessiontypes.Session = (*Session)(nil)

// paymentAmountForBytes calculates the payment amount based on data usage.
func (m *Session) paymentAmountForBytes() sdkmath.Int {
	decPrice := m.Price.QuoteValue.ToLegacyDec()
	bytePrice := decPrice.QuoInt(base.Gigabyte)
	totalBytes := m.DownloadBytes.Add(m.UploadBytes)

	return bytePrice.MulInt(totalBytes).Ceil().TruncateInt()
}

// paymentAmountForDuration calculates the payment amount based on session duration.
func (m *Session) paymentAmountForDuration() sdkmath.Int {
	decPrice := m.Price.QuoteValue.ToLegacyDec()
	nsPrice := decPrice.QuoInt64(time.Hour.Nanoseconds())
	nsDuration := m.Duration.Nanoseconds()

	return nsPrice.MulInt64(nsDuration).Ceil().TruncateInt()
}

// DepositAmount calculates the deposit amount based on maximum allowed bytes and duration.
func (m *Session) DepositAmount() sdk.Coin {
	amount := sdkmath.ZeroInt()

	if !m.MaxBytes.IsZero() {
		gigabytes := m.MaxBytes.Quo(base.Gigabyte)
		amount = m.Price.QuoteValue.Mul(gigabytes)
	}
	if m.MaxDuration != 0 {
		hours := int64(m.MaxDuration / time.Hour)
		amount = m.Price.QuoteValue.MulRaw(hours)
	}

	return sdk.Coin{Denom: m.Price.Denom, Amount: amount}
}

// PaymentAmount calculates the total payment amount based on usage and duration.
func (m *Session) PaymentAmount() sdk.Coin {
	amount := sdkmath.ZeroInt()

	if !m.MaxBytes.IsZero() {
		amount = m.paymentAmountForBytes()
	}
	if m.MaxDuration != 0 {
		amount = m.paymentAmountForDuration()
	}

	return sdk.Coin{Denom: m.Price.Denom, Amount: amount}
}

// RefundAmount calculates the refund amount as the difference between deposit and payment.
func (m *Session) RefundAmount() sdk.Coin {
	deposit := m.DepositAmount()
	payment := m.PaymentAmount()

	if payment.IsLT(deposit) {
		return deposit.Sub(payment)
	}

	return sdk.Coin{Denom: m.Price.Denom, Amount: sdkmath.ZeroInt()}
}
