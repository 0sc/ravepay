package rave

import (
	"reflect"
	"testing"
)

func TestNewTxnVerificationChecklist(t *testing.T) {
	type args struct {
		amount   int
		flwRef   string
		currency string
		secKey   string
	}
	tests := []struct {
		name string
		args args
		want *TxnVerificationChecklist
	}{
		{
			name: "returns a new txn verification checklist",
			args: args{
				amount:   1000,
				flwRef:   "some-flw-ref",
				currency: "NGN",
				secKey:   "some-sec-key",
			},
			want: &TxnVerificationChecklist{
				Amount:              1000,
				FlwRef:              "some-flw-ref",
				TransactionCurrency: "NGN",
				SECKEY:              "some-sec-key",
				VerificationURL:     TransactionVerificationURL,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTxnVerificationChecklist(tt.args.amount, tt.args.flwRef, tt.args.currency, tt.args.secKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTxnVerificationChecklist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewXRQTxnVerificationChecklist(t *testing.T) {
	type args struct {
		amount   int
		flwRef   string
		txRef    string
		currency string
		secKey   string
	}
	tests := []struct {
		name string
		args args
		want *TxnVerificationChecklist
	}{
		{
			name: "returns a new txn verification checklist",
			args: args{
				amount:   1000,
				flwRef:   "some-flw-ref",
				currency: "NGN",
				secKey:   "some-sec-key",
			},
			want: &TxnVerificationChecklist{
				Amount:              1000,
				Flwref:              "some-flw-ref",
				FlwRef:              "some-flw-ref",
				TransactionCurrency: "NGN",
				LastAttempt:         "1",
				OnlySuccessful:      "1",
				SECKEY:              "some-sec-key",
				VerificationURL:     TransactionVerificationRequeryURL,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewXRQTxnVerificationChecklist(tt.args.amount, tt.args.flwRef, tt.args.txRef, tt.args.currency, tt.args.secKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewXRQTxnVerificationChecklist() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
