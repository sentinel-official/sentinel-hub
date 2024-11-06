package v3

import (
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
)

func (m *Session) GetDownloadBytes() sdkmath.Int { return m.DownloadBytes }
func (m *Session) GetUploadBytes() sdkmath.Int   { return m.UploadBytes }

func (m *Session) SetDownloadBytes(v sdkmath.Int) { m.DownloadBytes = v }
func (m *Session) SetDuration(v time.Duration)    { m.Duration = v }
func (m *Session) SetInactiveAt(v time.Time)      { m.InactiveAt = v }
func (m *Session) SetStatusAt(v time.Time)        { m.StatusAt = v }
func (m *Session) SetStatus(v v1base.Status)      { m.Status = v }
func (m *Session) SetUploadBytes(v sdkmath.Int)   { m.UploadBytes = v }

func (m *Session) paymentAmountForBytes() sdkmath.Int {
	decPrice := m.Price.Amount.ToLegacyDec()
	bytePrice := decPrice.QuoInt(base.Gigabyte)
	totalBytes := m.DownloadBytes.Add(m.UploadBytes)
	return bytePrice.MulInt(totalBytes).Ceil().TruncateInt()
}

func (m *Session) paymentAmountForDuration() sdkmath.Int {
	decPrice := m.Price.Amount.ToLegacyDec()
	nsPrice := decPrice.QuoInt64(time.Hour.Nanoseconds())
	nsDuration := m.Duration.Nanoseconds()
	return nsPrice.MulInt64(nsDuration).Ceil().TruncateInt()
}

func (m *Session) DepositAmount() sdk.Coin {
	amount := sdk.ZeroInt()
	if m.MaxGigabytes != 0 {
		amount = m.Price.Amount.MulRaw(m.MaxGigabytes)
	}
	if m.MaxHours != 0 {
		amount = m.Price.Amount.MulRaw(m.MaxHours)
	}

	return sdk.NewCoin(m.Price.Denom, amount)
}

func (m *Session) PaymentAmount() sdk.Coin {
	amount := sdk.ZeroInt()
	if m.MaxGigabytes != 0 {
		amount = m.paymentAmountForBytes()
	}
	if m.MaxHours != 0 {
		amount = m.paymentAmountForDuration()
	}

	return sdk.NewCoin(m.Price.Denom, amount)
}

func (m *Session) RefundAmount() sdk.Coin {
	deposit := m.DepositAmount()
	payment := m.PaymentAmount()

	if payment.IsLT(deposit) {
		return deposit.Sub(payment)
	}

	return sdk.NewCoin(m.Price.Denom, sdk.ZeroInt())
}
