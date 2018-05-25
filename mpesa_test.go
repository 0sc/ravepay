package ravepay

import (
	"reflect"
	"testing"
)

func TestMpesaPaymentInstruction(t *testing.T) {
	type args struct {
		cr *ChargeResponse
	}
	tests := []struct {
		name string
		args args
		want *MpesaPaymentInfo
	}{
		{
			name: "returns the mpesa payment info-1",
			args: args{
				cr: &ChargeResponse{
					Data: chargeResponseData{
						Amount:         900,
						OrderRef:       "some-ref",
						BusinessNumber: "some-biz-num",
					},
				},
			},
			want: &MpesaPaymentInfo{
				Amount:         900,
				AccountNumber:  "some-ref",
				BusinessNumber: "some-biz-num",
			},
		},
		{
			name: "returns the mpesa payment info-2",
			args: args{
				cr: &ChargeResponse{},
			},
			want: &MpesaPaymentInfo{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MpesaPaymentInstruction(tt.args.cr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MpesaPaymentInstruction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMpesa_ChargeURL(t *testing.T) {
	type fields struct {
		ChargeMpesaURL string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "returns the ChargeMpesaURL in the mpesa object if present",
			fields: fields{"https://charge.mpesa.url"},
			want:   "https://charge.mpesa.url",
		},
		{
			name: "set's the object ChargeMpesaURL to config's defaultChargeURL and returns it",
			want: baseURL + defaultChargeURL,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mpesa{
				ChargeMpesaURL: tt.fields.ChargeMpesaURL,
			}
			if got := m.ChargeURL(); got != tt.want {
				t.Errorf("Mpesa.ChargeURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
