package v3

import (
	"time"

	sdkerrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/plan/types"
)

var (
	_ sdk.Msg = (*MsgCreatePlanRequest)(nil)
	_ sdk.Msg = (*MsgLinkNodeRequest)(nil)
	_ sdk.Msg = (*MsgUnlinkNodeRequest)(nil)
	_ sdk.Msg = (*MsgUpdatePlanStatusRequest)(nil)
	_ sdk.Msg = (*MsgStartSessionRequest)(nil)
)

func NewMsgCreatePlanRequest(from base.ProvAddress, bytes sdkmath.Int, duration time.Duration, prices sdk.DecCoins, private bool) *MsgCreatePlanRequest {
	return &MsgCreatePlanRequest{
		From:     from.String(),
		Bytes:    bytes,
		Duration: duration,
		Prices:   prices,
		Private:  private,
	}
}

func (m *MsgCreatePlanRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "from cannot be empty")
	}
	if _, err := base.ProvAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
	}
	if m.Bytes.IsNil() {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "bytes cannot be nil")
	}
	if m.Bytes.IsNegative() {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "bytes cannot be negative")
	}
	if m.Bytes.IsZero() {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "bytes cannot be zero")
	}
	if m.Duration < 0 {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "duration cannot be negative")
	}
	if m.Duration == 0 {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "duration cannot be zero")
	}
	if m.Prices == nil {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "prices cannot be nil")
	}
	if !m.Prices.IsValid() {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "prices must be valid")
	}

	return nil
}

func (m *MsgCreatePlanRequest) GetSigners() []sdk.AccAddress {
	from, err := base.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgLinkNodeRequest(from base.ProvAddress, id uint64, addr base.NodeAddress) *MsgLinkNodeRequest {
	return &MsgLinkNodeRequest{
		From:        from.String(),
		ID:          id,
		NodeAddress: addr.String(),
	}
}

func (m *MsgLinkNodeRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "from cannot be empty")
	}
	if _, err := base.ProvAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
	}
	if m.ID == 0 {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "id cannot be zero")
	}
	if m.NodeAddress == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "node_address cannot be empty")
	}
	if _, err := base.NodeAddressFromBech32(m.NodeAddress); err != nil {
		return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
	}

	return nil
}

func (m *MsgLinkNodeRequest) GetSigners() []sdk.AccAddress {
	from, err := base.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgUnlinkNodeRequest(from base.ProvAddress, id uint64, addr base.NodeAddress) *MsgUnlinkNodeRequest {
	return &MsgUnlinkNodeRequest{
		From:        from.String(),
		ID:          id,
		NodeAddress: addr.String(),
	}
}

func (m *MsgUnlinkNodeRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "from cannot be empty")
	}
	if _, err := base.ProvAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
	}
	if m.ID == 0 {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "id cannot be zero")
	}
	if m.NodeAddress == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "node_address cannot be empty")
	}
	if _, err := base.NodeAddressFromBech32(m.NodeAddress); err != nil {
		return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
	}

	return nil
}

func (m *MsgUnlinkNodeRequest) GetSigners() []sdk.AccAddress {
	from, err := base.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgUpdatePlanStatusRequest(from base.ProvAddress, id uint64, status v1base.Status) *MsgUpdatePlanStatusRequest {
	return &MsgUpdatePlanStatusRequest{
		From:   from.String(),
		ID:     id,
		Status: status,
	}
}

func (m *MsgUpdatePlanStatusRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "from cannot be empty")
	}
	if _, err := base.ProvAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
	}
	if m.ID == 0 {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "id cannot be zero")
	}
	if !m.Status.IsOneOf(v1base.StatusActive, v1base.StatusInactive) {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "status must be one of [active, inactive]")
	}

	return nil
}

func (m *MsgUpdatePlanStatusRequest) GetSigners() []sdk.AccAddress {
	from, err := base.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgStartSessionRequest(from sdk.AccAddress, id uint64, denom string, renewable bool, nodeAddr base.NodeAddress) *MsgStartSessionRequest {
	return &MsgStartSessionRequest{
		From:        from.String(),
		ID:          id,
		Denom:       denom,
		Renewable:   renewable,
		NodeAddress: nodeAddr.String(),
	}
}

func (m *MsgStartSessionRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
	}
	if m.ID == 0 {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "id cannot be zero")
	}
	if m.Denom != "" {
		if err := sdk.ValidateDenom(m.Denom); err != nil {
			return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
		}
	}
	if m.NodeAddress == "" {
		return sdkerrors.Wrap(types.ErrInvalidMessage, "node_address cannot be empty")
	}
	if _, err := base.NodeAddressFromBech32(m.NodeAddress); err != nil {
		return sdkerrors.Wrap(types.ErrInvalidMessage, err.Error())
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
