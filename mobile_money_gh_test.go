package rave

import "testing"

func TestMobileMoneyGH_ChargeURL(t *testing.T) {
	type fields struct {
		ChargeRequestURL string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "returns the ChargeRequestURL in the mpesa object if present",
			fields: fields{"https://charge.ghMM.url"},
			want:   "https://charge.ghMM.url",
		},
		{
			name: "set's the object ChargeGHMobileMoneyURL to config's ChargeURL and returns it",
			want: baseURL + MobileMoneyGHChargeURL,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gh := &MobileMoneyGH{
				ChargeRequestURL: tt.fields.ChargeRequestURL,
			}
			if got := gh.ChargeURL(); got != tt.want {
				t.Errorf("GHMobileMoney.ChargeURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
