package v3

import (
	"fmt"
	"time"

	sdkerrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
)

func (m *Plan) GetBytes() sdkmath.Int {
	return base.Gigabyte.MulRaw(m.Gigabytes)
}

func (m *Plan) GetDuration() time.Duration {
	return time.Duration(m.Hours) * time.Hour
}

func (m *Plan) Price(denom string) (sdk.DecCoin, bool) {
	for _, v := range m.Prices {
		if v.Denom == denom {
			return v, true
		}
	}

	// If there are no prices and denom is empty, return a zero amount coin and true
	return sdk.DecCoin{Amount: sdkmath.LegacyZeroDec()}, m.Prices.Len() == 0 && denom == ""
}

func (m *Plan) Validate() error {
	if m.ID == 0 {
		return fmt.Errorf("id cannot be zero")
	}
	if m.ProvAddress == "" {
		return fmt.Errorf("prov_address cannot be empty")
	}
	if _, err := base.ProvAddressFromBech32(m.ProvAddress); err != nil {
		return sdkerrors.Wrapf(err, "invalid prov_address %s", m.ProvAddress)
	}
	if m.Gigabytes < 0 {
		return fmt.Errorf("gigabytes cannot be negative")
	}
	if m.Gigabytes == 0 {
		return fmt.Errorf("gigabytes cannot be zero")
	}
	if m.Hours < 0 {
		return fmt.Errorf("hours cannot be negative")
	}
	if m.Hours == 0 {
		return fmt.Errorf("hours cannot be zero")
	}
	if m.Prices == nil {
		return fmt.Errorf("prices cannot be nil")
	}
	if !m.Prices.IsValid() {
		return fmt.Errorf("prices must be valid")
	}
	if !m.Status.IsOneOf(v1base.StatusActive, v1base.StatusInactive) {
		return fmt.Errorf("status must be one of [active, inactive]")
	}
	if m.StatusAt.IsZero() {
		return fmt.Errorf("status_at cannot be zero")
	}

	return nil
}
