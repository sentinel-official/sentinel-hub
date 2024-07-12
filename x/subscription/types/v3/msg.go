package v3

import (
	sdkerrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/subscription/types"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v2"
)

var (
	_ sdk.Msg = (*MsgCancelSubscriptionRequest)(nil)
	_ sdk.Msg = (*MsgRenewSubscriptionRequest)(nil)
	_ sdk.Msg = (*MsgShareSubscriptionRequest)(nil)
	_ sdk.Msg = (*MsgStartSubscriptionRequest)(nil)
	_ sdk.Msg = (*MsgUpdateSubscriptionRequest)(nil)
	_ sdk.Msg = (*MsgStartSessionRequest)(nil)
	_ sdk.Msg = (*MsgUpdateParamsRequest)(nil)
)

func NewMsgCancelSubscriptionRequest(from sdk.AccAddress, id uint64) *MsgCancelSubscriptionRequest {
	return &MsgCancelSubscriptionRequest{
		From: from.String(),
		ID:   id,
	}
}

func (m *MsgCancelSubscriptionRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrapf(types.ErrorInvalidMessage, "invalid from %s", err)
	}
	if m.ID == 0 {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "id cannot be zero")
	}

	return nil
}

func (m *MsgCancelSubscriptionRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgRenewSubscriptionRequest(from sdk.AccAddress, id uint64, denom string) *MsgRenewSubscriptionRequest {
	return &MsgRenewSubscriptionRequest{
		From:  from.String(),
		ID:    id,
		Denom: denom,
	}
}

func (m *MsgRenewSubscriptionRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}
	if m.ID == 0 {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "id cannot be zero")
	}
	if m.Denom == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "denom cannot be empty")
	}
	if err := sdk.ValidateDenom(m.Denom); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}

	return nil
}

func (m *MsgRenewSubscriptionRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgShareSubscriptionRequest(from sdk.AccAddress, id uint64, accAddr sdk.AccAddress, bytes sdkmath.Int) *MsgShareSubscriptionRequest {
	return &MsgShareSubscriptionRequest{
		From:       from.String(),
		ID:         id,
		AccAddress: accAddr.String(),
		Bytes:      bytes,
	}
}

func (m *MsgShareSubscriptionRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrapf(types.ErrorInvalidMessage, "invalid from %s", err)
	}
	if m.ID == 0 {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "id cannot be zero")
	}
	if m.AccAddress == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "acc_address cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.AccAddress); err != nil {
		return sdkerrors.Wrapf(types.ErrorInvalidMessage, "invalid acc_address %s", err)
	}
	if m.Bytes.IsNil() {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "bytes cannot be nil")
	}
	if m.Bytes.IsNegative() {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "bytes cannot be negative")
	}

	return nil
}

func (m *MsgShareSubscriptionRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgStartSubscriptionRequest(from sdk.AccAddress, id uint64, denom string, renewable bool) *MsgStartSubscriptionRequest {
	return &MsgStartSubscriptionRequest{
		From:      from.String(),
		ID:        id,
		Denom:     denom,
		Renewable: renewable,
	}
}

func (m *MsgStartSubscriptionRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}
	if m.ID == 0 {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "id cannot be zero")
	}
	if m.Denom == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "denom cannot be empty")
	}
	if err := sdk.ValidateDenom(m.Denom); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}

	return nil
}

func (m *MsgStartSubscriptionRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgUpdateSubscriptionRequest(from sdk.AccAddress, id uint64, renewable bool) *MsgUpdateSubscriptionRequest {
	return &MsgUpdateSubscriptionRequest{
		From:      from.String(),
		ID:        id,
		Renewable: renewable,
	}
}

func (m *MsgUpdateSubscriptionRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}
	if m.ID == 0 {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "id cannot be zero")
	}

	return nil
}

func (m *MsgUpdateSubscriptionRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgStartSessionRequest(from sdk.AccAddress, id uint64, nodeAddr base.NodeAddress) *MsgStartSessionRequest {
	return &MsgStartSessionRequest{
		From:        from.String(),
		ID:          id,
		NodeAddress: nodeAddr.String(),
	}
}

func (m *MsgStartSessionRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}
	if m.ID == 0 {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "id cannot be zero")
	}
	if m.NodeAddress == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "node_address cannot be empty")
	}
	if _, err := base.NodeAddressFromBech32(m.NodeAddress); err != nil {
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

func NewMsgUpdateParamsRequest(from sdk.AccAddress, params v2.Params) *MsgUpdateParamsRequest {
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
