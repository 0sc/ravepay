package rave

import "testing"

func TestCard_ChargeURL(t *testing.T) {
	type fields struct {
		ChargeCardURL string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "returns the ChargeCardURL in the card object if present",
			fields: fields{"https://charge.card.url"},
			want:   "https://charge.card.url",
		},
		{
			name: "set's the object ChargeCardURL to config's ChargeCardURL and returns it",
			want: baseURL + ChargeCardURL,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Card{
				ChargeCardURL: tt.fields.ChargeCardURL,
			}
			if got := c.ChargeURL(); got != tt.want {
				t.Errorf("Card.ChargeURL() = %v, want %v", got, tt.want)
			}

			if c.ChargeCardURL != tt.want {
				t.Errorf("Card.ChargeURL() = %v, want %v", c.ChargeCardURL, tt.want)
			}
		})
	}
}

func TestCard_ValidateChargeURL(t *testing.T) {
	type fields struct {
		ValidateCardChargeURL string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "returns ValidateCardChargeURL in card object if present",
			fields: fields{"https://validate.card.charge.url"},
			want:   "https://validate.card.charge.url",
		},
		{
			name: "set's the object ValidateCardChargeURL to the config ValidateCardChargeURL and returns it",
			want: ValidateCardChargeURL,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Card{
				ValidateCardChargeURL: tt.fields.ValidateCardChargeURL,
			}
			if got := c.ValidateChargeURL(); got != tt.want {
				t.Errorf("Card.ValidateChargeURL() = %v, want %v", got, tt.want)
			}

			if c.ValidateCardChargeURL != tt.want {
				t.Errorf("Card.ValidateChargeURL() = %v, want %v", c.ValidateCardChargeURL, tt.want)
			}
		})
	}
}
