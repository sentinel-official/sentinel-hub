package v1

import (
	"reflect"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
)

func TestDeposit_GetAddress(t *testing.T) {
	type fields struct {
		Address string
	}
	tests := []struct {
		name   string
		fields fields
		want   sdk.AccAddress
	}{
		{
			"empty",
			fields{
				Address: base.TestAddrEmpty,
			},
			nil,
		},
		{
			"20 bytes",
			fields{
				Address: base.TestBech32AccAddr20Bytes,
			},
			sdk.AccAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Deposit{
				Address: tt.fields.Address,
			}
			if got := d.GetAddress(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeposit_Validate(t *testing.T) {
	type fields struct {
		Address string
		Coins   sdk.Coins
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"empty address",
			fields{
				Address: base.TestAddrEmpty,
			},
			true,
		},
		{
			"invalid address",
			fields{
				Address: base.TestAddrInvalid,
			},
			true,
		},
		{
			"invalid prefix address",
			fields{
				Address: base.TestBech32NodeAddr20Bytes,
			},
			true,
		},
		{
			"10 bytes address",
			fields{
				Address: base.TestBech32AccAddr10Bytes,
				Coins:   base.TestCoinsOnePos,
			},
			false,
		},
		{
			"20 bytes address",
			fields{
				Address: base.TestBech32AccAddr20Bytes,
				Coins:   base.TestCoinsOnePos,
			},
			false,
		},
		{
			"30 bytes address",
			fields{
				Address: base.TestBech32AccAddr30Bytes,
				Coins:   base.TestCoinsOnePos,
			},
			false,
		},
		{
			"nil coins",
			fields{
				Address: base.TestBech32AccAddr20Bytes,
				Coins:   nil,
			},
			true,
		},
		{
			"empty coins",
			fields{
				Address: base.TestBech32AccAddr20Bytes,
				Coins:   base.TestCoinsEmpty,
			},
			true,
		},
		{
			"empty denom coins",
			fields{
				Address: base.TestBech32AccAddr20Bytes,
				Coins:   base.TestCoinsEmptyPos,
			},
			true,
		},
		{
			"invalid denom coins",
			fields{
				Address: base.TestBech32AccAddr20Bytes,
				Coins:   base.TestCoinsInvalidPos,
			},
			true,
		},
		{
			"nil amount coins",
			fields{
				Address: base.TestBech32AccAddr20Bytes,
				Coins:   base.TestCoinsOneEmpty,
			},
			true,
		},
		{
			"negative amount coins",
			fields{
				Address: base.TestBech32AccAddr20Bytes,
				Coins:   base.TestCoinsOneNeg,
			},
			true,
		},
		{
			"zero amount coins",
			fields{
				Address: base.TestBech32AccAddr20Bytes,
				Coins:   base.TestCoinsOneZero,
			},
			true,
		},
		{
			"positive amount coins",
			fields{
				Address: base.TestBech32AccAddr20Bytes,
				Coins:   base.TestCoinsOnePos,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Deposit{
				Address: tt.fields.Address,
				Coins:   tt.fields.Coins,
			}
			if err := d.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
