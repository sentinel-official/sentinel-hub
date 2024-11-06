package v2

import (
	"reflect"
	"strings"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
)

func TestNode_GetAddress(t *testing.T) {
	type fields struct {
		Address string
	}
	tests := []struct {
		name   string
		fields fields
		want   base.NodeAddress
	}{
		{
			"address empty",
			fields{
				Address: base.TestAddrEmpty,
			},
			nil,
		},
		{
			"address 20 bytes",
			fields{
				Address: base.TestBech32NodeAddr20Bytes,
			},
			base.NodeAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Node{
				Address: tt.fields.Address,
			}
			if got := m.GetAddress(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_GigabytePrice(t *testing.T) {
	type fields struct {
		GigabytePrices sdk.Coins
	}
	type args struct {
		denom string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   sdk.Coin
		want1  bool
	}{
		{
			"gigabyte_prices nil and denom empty",
			fields{
				GigabytePrices: base.TestCoinsNil,
			},
			args{
				denom: base.TestDenomEmpty,
			},
			base.TestCoinEmptyZero,
			true,
		},
		{
			"gigabyte_prices empty and denom empty",
			fields{
				GigabytePrices: base.TestCoinsEmpty,
			},
			args{
				denom: base.TestDenomEmpty,
			},
			base.TestCoinEmptyZero,
			true,
		},
		{
			"gigabyte_prices 1000one and denom empty",
			fields{
				GigabytePrices: base.TestCoinsOnePos,
			},
			args{
				denom: base.TestDenomEmpty,
			},
			base.TestCoinEmptyZero,
			false,
		},
		{
			"gigabyte_prices nil and denom one",
			fields{
				GigabytePrices: base.TestCoinsNil,
			},
			args{
				denom: base.TestDenomOne,
			},
			base.TestCoinEmptyZero,
			false,
		},
		{
			"gigabyte_prices empty and denom one",
			fields{
				GigabytePrices: base.TestCoinsEmpty,
			},
			args{
				denom: base.TestDenomOne,
			},
			base.TestCoinEmptyZero,
			false,
		},
		{
			"gigabyte_prices 1000one and denom one",
			fields{
				GigabytePrices: base.TestCoinsOnePos,
			},
			args{
				denom: base.TestDenomOne,
			},
			base.TestCoinOnePos,
			true,
		},
		{
			"gigabyte_prices nil and denom two",
			fields{
				GigabytePrices: base.TestCoinsNil,
			},
			args{
				denom: base.TestDenomTwo,
			},
			base.TestCoinEmptyZero,
			false,
		},
		{
			"gigabyte_prices empty and denom two",
			fields{
				GigabytePrices: base.TestCoinsEmpty,
			},
			args{
				denom: base.TestDenomTwo,
			},
			base.TestCoinEmptyZero,
			false,
		},
		{
			"gigabyte_prices 1000one and denom two",
			fields{
				GigabytePrices: base.TestCoinsOnePos,
			},
			args{
				denom: base.TestDenomTwo,
			},
			base.TestCoinEmptyZero,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Node{
				GigabytePrices: tt.fields.GigabytePrices,
			}
			got, got1 := m.GigabytePrice(tt.args.denom)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GigabytePrice() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GigabytePrice() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNode_HourlyPrice(t *testing.T) {
	type fields struct {
		HourlyPrices sdk.Coins
	}
	type args struct {
		denom string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   sdk.Coin
		want1  bool
	}{
		{
			"hourly_prices nil and denom empty",
			fields{
				HourlyPrices: base.TestCoinsNil,
			},
			args{
				denom: base.TestDenomEmpty,
			},
			base.TestCoinEmptyZero,
			true,
		},
		{
			"hourly_prices empty and denom empty",
			fields{
				HourlyPrices: base.TestCoinsEmpty,
			},
			args{
				denom: base.TestDenomEmpty,
			},
			base.TestCoinEmptyZero,
			true,
		},
		{
			"hourly_prices 1000one and denom empty",
			fields{
				HourlyPrices: base.TestCoinsOnePos,
			},
			args{
				denom: base.TestDenomEmpty,
			},
			base.TestCoinEmptyZero,
			false,
		},
		{
			"hourly_prices nil and denom one",
			fields{
				HourlyPrices: nil,
			},
			args{
				denom: base.TestDenomOne,
			},
			base.TestCoinEmptyZero,
			false,
		},
		{
			"hourly_prices empty and denom one",
			fields{
				HourlyPrices: base.TestCoinsEmpty,
			},
			args{
				denom: base.TestDenomOne,
			},
			base.TestCoinEmptyZero,
			false,
		},
		{
			"hourly_prices 1000one and denom one",
			fields{
				HourlyPrices: base.TestCoinsOnePos,
			},
			args{
				denom: base.TestDenomOne,
			},
			base.TestCoinOnePos,
			true,
		},
		{
			"hourly_prices nil and denom two",
			fields{
				HourlyPrices: nil,
			},
			args{
				denom: base.TestDenomTwo,
			},
			base.TestCoinEmptyZero,
			false,
		},
		{
			"hourly_prices empty and denom two",
			fields{
				HourlyPrices: base.TestCoinsEmpty,
			},
			args{
				denom: base.TestDenomTwo,
			},
			base.TestCoinEmptyZero,
			false,
		},
		{
			"hourly_prices 1000one and denom two",
			fields{
				HourlyPrices: base.TestCoinsOnePos,
			},
			args{
				denom: base.TestDenomTwo,
			},
			base.TestCoinEmptyZero,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Node{
				HourlyPrices: tt.fields.HourlyPrices,
			}
			got, got1 := m.HourlyPrice(tt.args.denom)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HourlyPrice() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("HourlyPrice() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNode_Validate(t *testing.T) {
	type fields struct {
		Address        string
		GigabytePrices sdk.Coins
		HourlyPrices   sdk.Coins
		RemoteURL      string
		InactiveAt     time.Time
		Status         v1base.Status
		StatusAt       time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"address empty",
			fields{
				Address: base.TestAddrEmpty,
			},
			true,
		},
		{
			"address invalid",
			fields{
				Address: base.TestAddrInvalid,
			},
			true,
		},
		{
			"address invalid prefix",
			fields{
				Address: base.TestBech32AccAddr20Bytes,
			},
			true,
		},
		{
			"address 10 bytes",
			fields{
				Address:        base.TestBech32NodeAddr10Bytes,
				GigabytePrices: base.TestCoinsOnePos,
				HourlyPrices:   base.TestCoinsOnePos,
				RemoteURL:      "https://remote.url:443",
				InactiveAt:     base.TestTimeNow,
				Status:         v1base.StatusActive,
				StatusAt:       base.TestTimeNow,
			},
			false,
		},
		{
			"address 20 bytes",
			fields{
				Address:        base.TestBech32NodeAddr20Bytes,
				GigabytePrices: base.TestCoinsOnePos,
				HourlyPrices:   base.TestCoinsOnePos,
				RemoteURL:      "https://remote.url:443",
				InactiveAt:     base.TestTimeNow,
				Status:         v1base.StatusActive,
				StatusAt:       base.TestTimeNow,
			},
			false,
		},
		{
			"address 30 bytes",
			fields{
				Address:        base.TestBech32NodeAddr30Bytes,
				GigabytePrices: base.TestCoinsOnePos,
				HourlyPrices:   base.TestCoinsOnePos,
				RemoteURL:      "https://remote.url:443",
				InactiveAt:     base.TestTimeNow,
				Status:         v1base.StatusActive,
				StatusAt:       base.TestTimeNow,
			},
			false,
		},
		{
			"gigabyte_prices nil",
			fields{
				Address:        base.TestBech32NodeAddr20Bytes,
				GigabytePrices: base.TestCoinsNil,
				HourlyPrices:   base.TestCoinsOnePos,
				RemoteURL:      "https://remote.url:443",
				InactiveAt:     base.TestTimeNow,
				Status:         v1base.StatusActive,
				StatusAt:       base.TestTimeNow,
			},
			true,
		},
		{
			"gigabyte_prices empty",
			fields{
				Address:        base.TestBech32NodeAddr20Bytes,
				GigabytePrices: base.TestCoinsEmpty,
			},
			true,
		},
		{
			"gigabyte_prices empty denom",
			fields{
				Address:        base.TestBech32NodeAddr20Bytes,
				GigabytePrices: base.TestCoinsEmptyPos,
			},
			true,
		},
		{
			"gigabyte_prices invalid denom",
			fields{
				Address:        base.TestBech32NodeAddr20Bytes,
				GigabytePrices: base.TestCoinsInvalidPos,
			},
			true,
		},
		{
			"gigabyte_prices empty amount",
			fields{
				Address:        base.TestBech32NodeAddr20Bytes,
				GigabytePrices: base.TestCoinsOneEmpty,
			},
			true,
		},
		{
			"gigabyte_prices negative amount",
			fields{
				Address:        base.TestBech32NodeAddr20Bytes,
				GigabytePrices: base.TestCoinsOneNeg,
			},
			true,
		},
		{
			"gigabyte_prices zero amount",
			fields{
				Address:        base.TestBech32NodeAddr20Bytes,
				GigabytePrices: base.TestCoinsOneZero,
			},
			true,
		},
		{
			"gigabyte_prices positive amount",
			fields{
				Address:        base.TestBech32NodeAddr20Bytes,
				GigabytePrices: base.TestCoinsOnePos,
				HourlyPrices:   base.TestCoinsOnePos,
				RemoteURL:      "https://remote.url:443",
				InactiveAt:     base.TestTimeNow,
				Status:         v1base.StatusActive,
				StatusAt:       base.TestTimeNow,
			},
			false,
		},
		{
			"hourly_prices nil",
			fields{
				Address:        base.TestBech32NodeAddr20Bytes,
				GigabytePrices: base.TestCoinsOnePos,
				HourlyPrices:   base.TestCoinsNil,
				RemoteURL:      "https://remote.url:443",
				InactiveAt:     base.TestTimeNow,
				Status:         v1base.StatusActive,
				StatusAt:       base.TestTimeNow,
			},
			true,
		},
		{
			"hourly_prices empty",
			fields{
				Address:        base.TestBech32NodeAddr20Bytes,
				GigabytePrices: base.TestCoinsOnePos,
				HourlyPrices:   base.TestCoinsEmpty,
			},
			true,
		},
		{
			"hourly_prices empty denom",
			fields{
				Address:        base.TestBech32NodeAddr20Bytes,
				GigabytePrices: base.TestCoinsOnePos,
				HourlyPrices:   base.TestCoinsEmptyPos,
			},
			true,
		},
		{
			"hourly_prices invalid denom",
			fields{
				Address:        base.TestBech32NodeAddr20Bytes,
				GigabytePrices: base.TestCoinsOnePos,
				HourlyPrices:   base.TestCoinsInvalidPos,
			},
			true,
		},
		{
			"hourly_prices empty amount",
			fields{
				Address:        base.TestBech32NodeAddr20Bytes,
				GigabytePrices: base.TestCoinsOnePos,
				HourlyPrices:   base.TestCoinsOneEmpty,
			},
			true,
		},
		{
			"hourly_prices negative amount",
			fields{
				Address:        base.TestBech32NodeAddr20Bytes,
				GigabytePrices: base.TestCoinsOnePos,
				HourlyPrices:   base.TestCoinsOneNeg,
			},
			true,
		},
		{
			"hourly_prices zero amount",
			fields{
				Address:        base.TestBech32NodeAddr20Bytes,
				GigabytePrices: base.TestCoinsOnePos,
				HourlyPrices:   base.TestCoinsOneZero,
			},
			true,
		},
		{
			"hourly_prices positive amount",
			fields{
				Address:        base.TestBech32NodeAddr20Bytes,
				GigabytePrices: base.TestCoinsOnePos,
				HourlyPrices:   base.TestCoinsOnePos,
				RemoteURL:      "https://remote.url:443",
				InactiveAt:     base.TestTimeNow,
				Status:         v1base.StatusActive,
				StatusAt:       base.TestTimeNow,
			},
			false,
		},
		{
			"remote_url empty",
			fields{
				Address:        base.TestBech32NodeAddr20Bytes,
				GigabytePrices: base.TestCoinsOnePos,
				HourlyPrices:   base.TestCoinsOnePos,
				RemoteURL:      "",
			},
			true,
		},
		{
			"remote_url 72 chars",
			fields{
				Address:        base.TestBech32NodeAddr20Bytes,
				GigabytePrices: base.TestCoinsOnePos,
				HourlyPrices:   base.TestCoinsOnePos,
				RemoteURL:      strings.Repeat("r", 72),
			},
			true,
		},
		{
			"remote_url invalid",
			fields{
				Address:        base.TestBech32NodeAddr20Bytes,
				GigabytePrices: base.TestCoinsOnePos,
				HourlyPrices:   base.TestCoinsOnePos,
				RemoteURL:      "invalid",
			},
			true,
		},
		{
			"remote_url invalid scheme",
			fields{
				Address:        base.TestBech32NodeAddr20Bytes,
				GigabytePrices: base.TestCoinsOnePos,
				HourlyPrices:   base.TestCoinsOnePos,
				RemoteURL:      "tcp://remote.url:80",
			},
			true,
		},
		{
			"remote_url without port",
			fields{
				Address:        base.TestBech32NodeAddr20Bytes,
				GigabytePrices: base.TestCoinsOnePos,
				HourlyPrices:   base.TestCoinsOnePos,
				RemoteURL:      "https://remote.url",
			},
			true,
		},
		{
			"remote_url with port",
			fields{
				Address:        base.TestBech32NodeAddr20Bytes,
				GigabytePrices: base.TestCoinsOnePos,
				HourlyPrices:   base.TestCoinsOnePos,
				RemoteURL:      "https://remote.url:443",
				InactiveAt:     base.TestTimeNow,
				Status:         v1base.StatusActive,
				StatusAt:       base.TestTimeNow,
			},
			false,
		},
		{
			"inactive_at empty",
			fields{
				Address:        base.TestBech32NodeAddr20Bytes,
				GigabytePrices: base.TestCoinsOnePos,
				HourlyPrices:   base.TestCoinsOnePos,
				RemoteURL:      "https://remote.url:443",
				InactiveAt:     base.TestTimeZero,
			},
			true,
		},
		{
			"inactive_at now",
			fields{
				Address:        base.TestBech32NodeAddr20Bytes,
				GigabytePrices: base.TestCoinsOnePos,
				HourlyPrices:   base.TestCoinsOnePos,
				RemoteURL:      "https://remote.url:443",
				InactiveAt:     base.TestTimeNow,
				Status:         v1base.StatusActive,
				StatusAt:       base.TestTimeNow,
			},
			false,
		},
		{
			"status unspecified",
			fields{
				Address:        base.TestBech32NodeAddr20Bytes,
				GigabytePrices: base.TestCoinsOnePos,
				HourlyPrices:   base.TestCoinsOnePos,
				RemoteURL:      "https://remote.url:443",
				InactiveAt:     base.TestTimeNow,
				Status:         v1base.StatusUnspecified,
			},
			true,
		},
		{
			"status active",
			fields{
				Address:        base.TestBech32NodeAddr20Bytes,
				GigabytePrices: base.TestCoinsOnePos,
				HourlyPrices:   base.TestCoinsOnePos,
				RemoteURL:      "https://remote.url:443",
				InactiveAt:     base.TestTimeNow,
				Status:         v1base.StatusActive,
				StatusAt:       base.TestTimeNow,
			},
			false,
		},
		{
			"status inactive_pending",
			fields{
				Address:        base.TestBech32NodeAddr20Bytes,
				GigabytePrices: base.TestCoinsOnePos,
				HourlyPrices:   base.TestCoinsOnePos,
				RemoteURL:      "https://remote.url:443",
				InactiveAt:     base.TestTimeNow,
				Status:         v1base.StatusInactivePending,
			},
			true,
		},
		{
			"status inactive",
			fields{
				Address:        base.TestBech32NodeAddr20Bytes,
				GigabytePrices: base.TestCoinsOnePos,
				HourlyPrices:   base.TestCoinsOnePos,
				RemoteURL:      "https://remote.url:443",
				InactiveAt:     base.TestTimeZero,
				Status:         v1base.StatusInactive,
				StatusAt:       base.TestTimeNow,
			},
			false,
		},
		{
			"status_at empty",
			fields{
				Address:        base.TestBech32NodeAddr20Bytes,
				GigabytePrices: base.TestCoinsOnePos,
				HourlyPrices:   base.TestCoinsOnePos,
				RemoteURL:      "https://remote.url:443",
				InactiveAt:     base.TestTimeNow,
				Status:         v1base.StatusActive,
				StatusAt:       base.TestTimeZero,
			},
			true,
		},
		{
			"status_at now",
			fields{
				Address:        base.TestBech32NodeAddr20Bytes,
				GigabytePrices: base.TestCoinsOnePos,
				HourlyPrices:   base.TestCoinsOnePos,
				RemoteURL:      "https://remote.url:443",
				InactiveAt:     base.TestTimeNow,
				Status:         v1base.StatusActive,
				StatusAt:       base.TestTimeNow,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Node{
				Address:        tt.fields.Address,
				GigabytePrices: tt.fields.GigabytePrices,
				HourlyPrices:   tt.fields.HourlyPrices,
				RemoteURL:      tt.fields.RemoteURL,
				InactiveAt:     tt.fields.InactiveAt,
				Status:         tt.fields.Status,
				StatusAt:       tt.fields.StatusAt,
			}
			if err := m.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
