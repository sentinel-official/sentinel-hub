package types

import (
	"reflect"
	"strings"
	"testing"

	base "github.com/sentinel-official/hub/v12/types"
)

func TestProvider_GetAddress(t *testing.T) {
	type fields struct {
		Address     string
		Name        string
		Identity    string
		Website     string
		Description string
	}
	tests := []struct {
		name   string
		fields fields
		want   base.ProvAddress
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
				Address: base.TestBech32ProvAddr20Bytes,
			},
			base.ProvAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Provider{
				Address:     tt.fields.Address,
				Name:        tt.fields.Name,
				Identity:    tt.fields.Identity,
				Website:     tt.fields.Website,
				Description: tt.fields.Description,
			}
			if got := p.GetAddress(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProvider_Validate(t *testing.T) {
	type fields struct {
		Address     string
		Name        string
		Identity    string
		Website     string
		Description string
		Status      base.Status
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
				Address: base.TestBech32AccAddr20Bytes,
			},
			true,
		},
		{
			"10 bytes address",
			fields{
				Address: base.TestBech32ProvAddr10Bytes,
			},
			true,
		},
		{
			"20 bytes address",
			fields{
				Address: base.TestBech32ProvAddr20Bytes,
			},
			true,
		},
		{
			"30 bytes address",
			fields{
				Address: base.TestBech32ProvAddr30Bytes,
			},
			true,
		},
		{
			"empty name",
			fields{
				Address: base.TestBech32ProvAddr20Bytes,
				Name:    "",
			},
			true,
		},
		{
			"non-empty name",
			fields{
				Address: base.TestBech32ProvAddr20Bytes,
				Name:    "name",
				Status:  base.StatusActive,
			},
			false,
		},
		{
			"length 72 name",
			fields{
				Address: base.TestBech32ProvAddr20Bytes,
				Name:    strings.Repeat("n", 72),
			},
			true,
		},
		{
			"empty identity",
			fields{
				Address:  base.TestBech32ProvAddr20Bytes,
				Name:     "name",
				Identity: "",
				Status:   base.StatusActive,
			},
			false,
		},
		{
			"non-empty identity",
			fields{
				Address:  base.TestBech32ProvAddr20Bytes,
				Name:     "name",
				Identity: "identity",
				Status:   base.StatusActive,
			},
			false,
		},
		{
			"length 72 identity",
			fields{
				Address:  base.TestBech32ProvAddr20Bytes,
				Name:     "name",
				Identity: strings.Repeat("i", 72),
			},
			true,
		},
		{
			"empty website",
			fields{
				Address:  base.TestBech32ProvAddr20Bytes,
				Name:     "name",
				Identity: "identity",
				Website:  "",
				Status:   base.StatusActive,
			},
			false,
		},
		{
			"non-empty website",
			fields{
				Address:  base.TestBech32ProvAddr20Bytes,
				Name:     "name",
				Identity: "identity",
				Website:  "https://website",
				Status:   base.StatusActive,
			},
			false,
		},
		{
			"length 72 website",
			fields{
				Address:  base.TestBech32ProvAddr20Bytes,
				Name:     "name",
				Identity: "identity",
				Website:  strings.Repeat("w", 72),
			},
			true,
		},
		{
			"invalid website",
			fields{
				Address:  base.TestBech32ProvAddr20Bytes,
				Name:     "name",
				Identity: "identity",
				Website:  "invalid",
			},
			true,
		},
		{
			"empty description",
			fields{
				Address:     base.TestBech32ProvAddr20Bytes,
				Name:        "name",
				Identity:    "identity",
				Website:     "https://website",
				Description: "",
				Status:      base.StatusActive,
			},
			false,
		},
		{
			"non-empty description",
			fields{
				Address:     base.TestBech32ProvAddr20Bytes,
				Name:        "name",
				Identity:    "identity",
				Website:     "https://website",
				Description: "description",
				Status:      base.StatusActive,
			},
			false,
		},
		{
			"length 264 description",
			fields{
				Address:     base.TestBech32ProvAddr20Bytes,
				Name:        "name",
				Identity:    "identity",
				Website:     "https://website",
				Description: strings.Repeat("d", 264),
			},
			true,
		},
		{
			"unspecified status",
			fields{
				Address:     base.TestBech32ProvAddr20Bytes,
				Name:        "name",
				Identity:    "identity",
				Website:     "https://website",
				Description: strings.Repeat("d", 256),
				Status:      base.StatusUnspecified,
			},
			true,
		},
		{
			"active status",
			fields{
				Address:     base.TestBech32ProvAddr20Bytes,
				Name:        "name",
				Identity:    "identity",
				Website:     "https://website",
				Description: strings.Repeat("d", 256),
				Status:      base.StatusActive,
			},
			false,
		},
		{
			"inactive status",
			fields{
				Address:     base.TestBech32ProvAddr20Bytes,
				Name:        "name",
				Identity:    "identity",
				Website:     "https://website",
				Description: strings.Repeat("d", 256),
				Status:      base.StatusInactive,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Provider{
				Address:     tt.fields.Address,
				Name:        tt.fields.Name,
				Identity:    tt.fields.Identity,
				Website:     tt.fields.Website,
				Description: tt.fields.Description,
				Status:      tt.fields.Status,
			}
			if err := p.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
