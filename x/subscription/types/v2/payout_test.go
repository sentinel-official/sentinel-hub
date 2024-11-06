package v2

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
)

func TestPayout_Validate(t *testing.T) {
	type fields struct {
		ID     uint64
		Hours  int64
		Price  sdk.Coin
		NextAt time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"id zero",
			fields{
				ID: 0,
			},
			true,
		},
		{
			"id positive",
			fields{
				ID:     1000,
				Hours:  1000,
				Price:  base.TestCoinOnePos,
				NextAt: base.TestTimeNow,
			},
			false,
		},
		{
			"hours negative",
			fields{
				ID:    1000,
				Hours: -1000,
			},
			true,
		},
		{
			"hours zero",
			fields{
				ID:    1000,
				Hours: 0,
			},
			true,
		},
		{
			"hours positive",
			fields{
				ID:    1000,
				Hours: 1000,
				Price: base.TestCoinOnePos,
			},
			false,
		},
		{
			"price empty",
			fields{
				ID:    1000,
				Hours: 1000,
				Price: base.TestCoinEmptyZero,
			},
			true,
		},
		{
			"price empty denom",
			fields{
				ID:    1000,
				Hours: 1000,
				Price: base.TestCoinEmptyZero,
			},
			true,
		},
		{
			"price invalid denom",
			fields{
				ID:    1000,
				Hours: 1000,
				Price: base.TestCoinInvalidPos,
			},
			true,
		},
		{
			"price empty amount",
			fields{
				ID:    1000,
				Hours: 1000,
				Price: base.TestCoinOneEmpty,
			},
			true,
		},
		{
			"price negative amount",
			fields{
				ID:    1000,
				Hours: 1000,
				Price: base.TestCoinOneNeg,
			},
			true,
		},
		{
			"price zero amount",
			fields{
				ID:    1000,
				Hours: 1000,
				Price: base.TestCoinOneZero,
			},
			true,
		},
		{
			"price positive amount",
			fields{
				ID:    1000,
				Hours: 1000,
				Price: base.TestCoinOnePos,
			},
			false,
		},
		{
			"next_at zero",
			fields{
				ID:     1000,
				Hours:  1000,
				Price:  base.TestCoinOnePos,
				NextAt: base.TestTimeZero,
			},
			false,
		},
		{
			"next_at positive",
			fields{
				ID:     1000,
				Hours:  1000,
				Price:  base.TestCoinOnePos,
				NextAt: base.TestTimeNow,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Payout{
				ID:     tt.fields.ID,
				Hours:  tt.fields.Hours,
				Price:  tt.fields.Price,
				NextAt: tt.fields.NextAt,
			}
			if err := l.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
