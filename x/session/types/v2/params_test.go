package v2

import (
	"testing"
	"time"
)

func TestParams_Validate(t *testing.T) {
	type fields struct {
		StatusChangeDelay        time.Duration
		ProofVerificationEnabled bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"negative status_change_delay",
			fields{
				StatusChangeDelay: -1000,
			},
			true,
		},
		{
			"zero status_change_delay",
			fields{
				StatusChangeDelay: 0,
			},
			true,
		},
		{
			"positive status_change_delay",
			fields{
				StatusChangeDelay: 1000,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Params{
				StatusChangeDelay:        tt.fields.StatusChangeDelay,
				ProofVerificationEnabled: tt.fields.ProofVerificationEnabled,
			}
			if err := m.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
