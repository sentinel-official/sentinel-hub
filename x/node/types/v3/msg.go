package v3

import (
	"net/url"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/node/types"
)

var (
	_ sdk.Msg = (*MsgRegisterNodeRequest)(nil)
	_ sdk.Msg = (*MsgUpdateNodeDetailsRequest)(nil)
	_ sdk.Msg = (*MsgUpdateNodeStatusRequest)(nil)
	_ sdk.Msg = (*MsgStartSessionRequest)(nil)
	_ sdk.Msg = (*MsgUpdateParamsRequest)(nil)
)

func NewMsgRegisterNodeRequest(from sdk.AccAddress, gigabytePrices, hourlyPrices sdk.Coins, remoteURL string) *MsgRegisterNodeRequest {
	return &MsgRegisterNodeRequest{
		From:           from.String(),
		GigabytePrices: gigabytePrices,
		HourlyPrices:   hourlyPrices,
		RemoteURL:      remoteURL,
	}
}

func (m *MsgRegisterNodeRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}
	if m.GigabytePrices == nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "gigabyte_prices cannot be nil")
	}
	if m.GigabytePrices.Len() == 0 {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "gigabyte_prices length cannot be zero")
	}
	if m.GigabytePrices.IsAnyNil() {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "gigabyte_prices cannot contain nil")
	}
	if !m.GigabytePrices.IsValid() {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "gigabyte_prices must be valid")
	}
	if m.HourlyPrices == nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "hourly_prices cannot be nil")
	}
	if m.HourlyPrices.Len() == 0 {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "hourly_prices length cannot be zero")
	}
	if m.HourlyPrices.IsAnyNil() {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "hourly_prices cannot contain nil")
	}
	if !m.HourlyPrices.IsValid() {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "hourly_prices must be valid")
	}
	if m.RemoteURL == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "remote_url cannot be empty")
	}
	if len(m.RemoteURL) > 64 {
		return sdkerrors.Wrapf(types.ErrorInvalidMessage, "remote_url length cannot be greater than %d chars", 64)
	}

	remoteURL, err := url.ParseRequestURI(m.RemoteURL)
	if err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}
	if remoteURL.Scheme != "https" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "remote_url scheme must be https")
	}
	if remoteURL.Port() == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "remote_url port cannot be empty")
	}

	return nil
}

func (m *MsgRegisterNodeRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgUpdateNodeDetailsRequest(from base.NodeAddress, gigabytePrices, hourlyPrices sdk.Coins, remoteURL string) *MsgUpdateNodeDetailsRequest {
	return &MsgUpdateNodeDetailsRequest{
		From:           from.String(),
		GigabytePrices: gigabytePrices,
		HourlyPrices:   hourlyPrices,
		RemoteURL:      remoteURL,
	}
}

func (m *MsgUpdateNodeDetailsRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := base.NodeAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}
	if m.GigabytePrices != nil {
		if m.GigabytePrices.Len() == 0 {
			return sdkerrors.Wrap(types.ErrorInvalidMessage, "gigabyte_prices length cannot be zero")
		}
		if m.GigabytePrices.IsAnyNil() {
			return sdkerrors.Wrap(types.ErrorInvalidMessage, "gigabyte_prices cannot contain nil")
		}
		if !m.GigabytePrices.IsValid() {
			return sdkerrors.Wrap(types.ErrorInvalidMessage, "gigabyte_prices must be valid")
		}
	}
	if m.HourlyPrices != nil {
		if m.HourlyPrices.Len() == 0 {
			return sdkerrors.Wrap(types.ErrorInvalidMessage, "hourly_prices length cannot be zero")
		}
		if m.HourlyPrices.IsAnyNil() {
			return sdkerrors.Wrap(types.ErrorInvalidMessage, "hourly_prices cannot contain nil")
		}
		if !m.HourlyPrices.IsValid() {
			return sdkerrors.Wrap(types.ErrorInvalidMessage, "hourly_prices must be valid")
		}
	}
	if m.RemoteURL != "" {
		if len(m.RemoteURL) > 64 {
			return sdkerrors.Wrapf(types.ErrorInvalidMessage, "remote_url length cannot be greater than %d chars", 64)
		}

		remoteURL, err := url.ParseRequestURI(m.RemoteURL)
		if err != nil {
			return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
		}
		if remoteURL.Scheme != "https" {
			return sdkerrors.Wrap(types.ErrorInvalidMessage, "remote_url scheme must be https")
		}
		if remoteURL.Port() == "" {
			return sdkerrors.Wrap(types.ErrorInvalidMessage, "remote_url port cannot be empty")
		}
	}

	return nil
}

func (m *MsgUpdateNodeDetailsRequest) GetSigners() []sdk.AccAddress {
	from, err := base.NodeAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgUpdateNodeStatusRequest(from base.NodeAddress, status v1base.Status) *MsgUpdateNodeStatusRequest {
	return &MsgUpdateNodeStatusRequest{
		From:   from.String(),
		Status: status,
	}
}

func (m *MsgUpdateNodeStatusRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := base.NodeAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}
	if !m.Status.IsOneOf(v1base.StatusActive, v1base.StatusInactive) {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "status must be one of [active, inactive]")
	}

	return nil
}

func (m *MsgUpdateNodeStatusRequest) GetSigners() []sdk.AccAddress {
	from, err := base.NodeAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgStartSessionRequest(from sdk.AccAddress, nodeAddr base.NodeAddress, gigabytes, hours int64, denom string) *MsgStartSessionRequest {
	return &MsgStartSessionRequest{
		From:        from.String(),
		NodeAddress: nodeAddr.String(),
		Gigabytes:   gigabytes,
		Hours:       hours,
		Denom:       denom,
	}
}

func (m *MsgStartSessionRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}
	if m.NodeAddress == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "node_address cannot be empty")
	}
	if _, err := base.NodeAddressFromBech32(m.NodeAddress); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}
	if m.Gigabytes == 0 && m.Hours == 0 {
		return sdkerrors.Wrapf(types.ErrorInvalidMessage, "[gigabytes, hours] cannot be empty")
	}
	if m.Gigabytes != 0 && m.Hours != 0 {
		return sdkerrors.Wrapf(types.ErrorInvalidMessage, "[gigabytes, hours] cannot be non-empty")
	}
	if m.Gigabytes != 0 {
		if m.Gigabytes < 0 {
			return sdkerrors.Wrap(types.ErrorInvalidMessage, "gigabytes cannot be negative")
		}
	}
	if m.Hours != 0 {
		if m.Hours < 0 {
			return sdkerrors.Wrap(types.ErrorInvalidMessage, "hours cannot be negative")
		}
	}
	if m.Denom == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "denom cannot be empty")
	}
	if err := sdk.ValidateDenom(m.Denom); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}

	return nil
}

func (m *MsgStartSessionRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgUpdateParamsRequest(from sdk.AccAddress, params Params) *MsgUpdateParamsRequest {
	return &MsgUpdateParamsRequest{
		From:   from.String(),
		Params: params,
	}
}

func (m *MsgUpdateParamsRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}
	if err := m.Params.Validate(); err != nil {
		return err
	}

	return nil
}

func (m *MsgUpdateParamsRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}
