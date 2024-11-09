package v1

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/oracle/types"
)

var (
	_ sdk.Msg = (*MsgAddAssetRequest)(nil)
	_ sdk.Msg = (*MsgDeleteAssetRequest)(nil)
	_ sdk.Msg = (*MsgUpdateAssetRequest)(nil)
	_ sdk.Msg = (*MsgUpdateParamsRequest)(nil)
)

func (m *MsgAddAssetRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}
	if err := sdk.ValidateDenom(m.Denom); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}
	if m.Decimals < 0 {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "decimals cannot be negative")
	}
	if m.BaseAssetDenom == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "base_asset_denom cannot be empty")
	}
	if m.QuoteAssetDenom == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "quote_asset_denom cannot be empty")
	}

	return nil
}

func (m *MsgAddAssetRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func (m *MsgDeleteAssetRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}
	if err := sdk.ValidateDenom(m.Denom); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}

	return nil
}

func (m *MsgDeleteAssetRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func (m *MsgUpdateAssetRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}
	if err := sdk.ValidateDenom(m.Denom); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}
	if m.Decimals < 0 {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "decimals cannot be negative")
	}
	if m.BaseAssetDenom == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "base_asset_denom cannot be empty")
	}
	if m.QuoteAssetDenom == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "quote_asset_denom cannot be empty")
	}

	return nil
}

func (m *MsgUpdateAssetRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func (m *MsgUpdateParamsRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}

	return m.Params.Validate()
}

func (m *MsgUpdateParamsRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}
