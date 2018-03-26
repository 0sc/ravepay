package rave

import (
	"reflect"
	"testing"
)

func TestUSSD_ChargeURL(t *testing.T) {
	type fields struct {
		ChargeRequestURL string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "returns the ChargeMpesaURL in the mpesa object if present",
			fields: fields{"https://charge.ussd.url"},
			want:   "https://charge.ussd.url",
		},
		{
			name: "set's the object ChargeMpesaURL to config's ChargeURL and returns it",
			want: ChargeUSSDURL,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &USSD{
				ChargeRequestURL: tt.fields.ChargeRequestURL,
			}
			if got := c.ChargeURL(); got != tt.want {
				t.Errorf("USSD.ChargeURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUSSDPaymentInstruction(t *testing.T) {
	type args struct {
		cr *ChargeResponse
	}
	tests := []struct {
		name string
		args args
		want *USSDPaymentInfo
	}{
		{
			name: "returns the ussd payment info-1",
			args: args{
				cr: &ChargeResponse{
					Data: chargeResponseData{
						Amount: 900,
						FlwRef: "some-ref",
					},
				},
			},
			want: &USSDPaymentInfo{
				Amount: 900,
				FlwRef: "some-ref",
			},
		},
		{
			name: "returns teh ussd payment info-2",
			args: args{
				cr: &ChargeResponse{},
			},
			want: &USSDPaymentInfo{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := USSDPaymentInstruction(tt.args.cr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("USSDPaymentInstruction() = %v, want %v", got, tt.want)
			}
		})
	}
}
