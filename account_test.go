package rave

import "testing"

func TestAccount_ChargeURL(t *testing.T) {
	type fields struct {
		ChargeAccountURL string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "returns the ChargeAccountURL in the account object if present",
			fields: fields{"https://charge.account.url"},
			want:   "https://charge.account.url",
		},
		{
			name: "set's the object ChargeAccountURL to config's defaultChargeURL and returns it",
			want: baseURL + defaultChargeURL,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				ChargeAccountURL: tt.fields.ChargeAccountURL,
			}
			if got := a.ChargeURL(); got != tt.want {
				t.Errorf("Account.ChargeURL() = %v, want %v", got, tt.want)
			}

			if a.ChargeAccountURL != tt.want {
				t.Errorf("Account.ChargeURL() = %v, want %v", a.ChargeAccountURL, tt.want)
			}
		})
	}
}

func TestAccount_ValidateChargeURL(t *testing.T) {
	type fields struct {
		ValidateAccountChargeURL string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "returns ValidateAccountChargeURL in the account object if present",
			fields: fields{"https://validate.account.charge.url"},
			want:   "https://validate.account.charge.url",
		},
		{
			name: "set's the object ValidateAccountChargeURL to the config ValidateAccountChargeURL and returns it",
			want: baseURL + ValidateAccountChargeURL,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				ValidateAccountChargeURL: tt.fields.ValidateAccountChargeURL,
			}
			if got := a.ValidateChargeURL(); got != tt.want {
				t.Errorf("Account.ValidateChargeURL() = %v, want %v", got, tt.want)
			}

			if a.ValidateAccountChargeURL != tt.want {
				t.Errorf("Account.ValidateChargeURL() = %v, want %v", a.ValidateAccountChargeURL, tt.want)
			}
		})
	}
}
