package ravepay

import (
	"testing"
)

func TestTxnVerificationResponse_VerifyStatus(t *testing.T) {
	type fields struct {
		Data    txnVerificationResponseData
		Message string
		Status  string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "returns an error is status is not success",
			fields: fields{
				Status: "not-success",
			},
			wantErr: true,
		},
		{
			name: "does not return an error is status is success",
			fields: fields{
				Status: "success",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &TxnVerificationResponse{
				Data:    tt.fields.Data,
				Message: tt.fields.Message,
				Status:  tt.fields.Status,
			}
			if err := resp.VerifyStatus(); (err != nil) != tt.wantErr {
				t.Errorf("TxnVerificationResponse.VerifyStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTxnVerificationResponse_VerifyCurrency(t *testing.T) {
	type fields struct {
		Data    txnVerificationResponseData
		Message string
		Status  string
	}
	type args struct {
		currency string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "returns an error if currency doesn't match the given currency",
			fields: fields{
				Data: txnVerificationResponseData{TransactionCurrency: "NGN"},
			},
			args:    args{currency: "USD"},
			wantErr: true,
		},
		{
			name: "doesn't return an error if currency matches the given currency",
			fields: fields{
				Data: txnVerificationResponseData{TransactionCurrency: "NGN"},
			},
			args:    args{currency: "NGN"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &TxnVerificationResponse{
				Data:    tt.fields.Data,
				Message: tt.fields.Message,
				Status:  tt.fields.Status,
			}
			if err := resp.VerifyCurrency(tt.args.currency); (err != nil) != tt.wantErr {
				t.Errorf("TxnVerificationResponse.VerifyCurrency() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTxnVerificationResponse_VerifyAmount(t *testing.T) {
	type fields struct {
		Data    txnVerificationResponseData
		Message string
		Status  string
	}
	type args struct {
		amt int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "returns an error if amount is lesser than the given amount",
			fields: fields{
				Data: txnVerificationResponseData{Amount: 999},
			},
			args:    args{amt: 1000},
			wantErr: true,
		},
		{
			name: "doesn't return an error if amount equals the given amount",
			fields: fields{
				Data: txnVerificationResponseData{Amount: 1000},
			},
			args:    args{amt: 1000},
			wantErr: false,
		},
		{
			name: "doesn't return an error if amount is greater than the given amount",
			fields: fields{
				Data: txnVerificationResponseData{Amount: 2000},
			},
			args:    args{amt: 1000},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &TxnVerificationResponse{
				Data:    tt.fields.Data,
				Message: tt.fields.Message,
				Status:  tt.fields.Status,
			}
			if err := resp.VerifyAmount(tt.args.amt); (err != nil) != tt.wantErr {
				t.Errorf("TxnVerificationResponse.VerifyAmount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTxnVerificationResponse_VerifyChargeResponseValue(t *testing.T) {
	type fields struct {
		Data    txnVerificationResponseData
		Message string
		Status  string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "returns an error if charge response is not 00 or 0",
			fields: fields{
				Data: txnVerificationResponseData{
					FlwMeta: FlwMeta{ChargeResponse: "404"},
				},
			},
			wantErr: true,
		},
		{
			name: "doesn't return an error if charge response is 0",
			fields: fields{
				Data: txnVerificationResponseData{
					FlwMeta: FlwMeta{ChargeResponse: "0"},
				},
			},
			wantErr: false,
		},
		{
			name: "doesn't return an error if charge response is 00",
			fields: fields{
				Data: txnVerificationResponseData{
					FlwMeta: FlwMeta{ChargeResponse: "00"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &TxnVerificationResponse{
				Data:    tt.fields.Data,
				Message: tt.fields.Message,
				Status:  tt.fields.Status,
			}
			if err := resp.VerifyChargeResponseValue(); (err != nil) != tt.wantErr {
				t.Errorf("TxnVerificationResponse.VerifyChargeResponseValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTxnVerificationResponse_VerifyReference(t *testing.T) {
	type fields struct {
		Data    txnVerificationResponseData
		Message string
		Status  string
	}
	type args struct {
		ref string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "returns an error if flwref doesn't match the given ref",
			fields: fields{
				Data: txnVerificationResponseData{FlwRef: "another-ref"},
			},
			args:    args{ref: "my-ref"},
			wantErr: true,
		},
		{
			name: "doesn't return an error if flwref matches the given ref",
			fields: fields{
				Data: txnVerificationResponseData{FlwRef: "my-ref"},
			},
			args:    args{ref: "my-ref"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &TxnVerificationResponse{
				Data:    tt.fields.Data,
				Message: tt.fields.Message,
				Status:  tt.fields.Status,
			}
			if err := resp.VerifyReference(tt.args.ref); (err != nil) != tt.wantErr {
				t.Errorf("TxnVerificationResponse.VerifyReference() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestXRQTxnVerificationResponse_VerifyStatus(t *testing.T) {
	type fields struct {
		Data    xRQTxnVerificationResponseData
		Message string
		Status  string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "returns an error is status is not success",
			fields: fields{
				Status: "not-success",
			},
			wantErr: true,
		},
		{
			name: "does not return an error is status is success",
			fields: fields{
				Status: "success",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &XRQTxnVerificationResponse{
				Data:    tt.fields.Data,
				Message: tt.fields.Message,
				Status:  tt.fields.Status,
			}
			if err := resp.VerifyStatus(); (err != nil) != tt.wantErr {
				t.Errorf("XRQTxnVerificationResponse.VerifyStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestXRQTxnVerificationResponse_VerifyCurrency(t *testing.T) {
	type fields struct {
		Data    xRQTxnVerificationResponseData
		Message string
		Status  string
	}
	type args struct {
		currency string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "returns an error if currency doesn't match the given currency",
			fields: fields{
				Data: xRQTxnVerificationResponseData{Currency: "NGN"},
			},
			args:    args{currency: "USD"},
			wantErr: true,
		},
		{
			name: "doesn't return an error if currency matches the given currency",
			fields: fields{
				Data: xRQTxnVerificationResponseData{Currency: "NGN"},
			},
			args:    args{currency: "NGN"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &XRQTxnVerificationResponse{
				Data:    tt.fields.Data,
				Message: tt.fields.Message,
				Status:  tt.fields.Status,
			}
			if err := resp.VerifyCurrency(tt.args.currency); (err != nil) != tt.wantErr {
				t.Errorf("XRQTxnVerificationResponse.VerifyCurrency() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestXRQTxnVerificationResponse_VerifyAmount(t *testing.T) {
	type fields struct {
		Data    xRQTxnVerificationResponseData
		Message string
		Status  string
	}
	type args struct {
		amt int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "returns an error if amount is lesser than the given amount",
			fields: fields{
				Data: xRQTxnVerificationResponseData{Amount: 999},
			},
			args:    args{amt: 1000},
			wantErr: true,
		},
		{
			name: "doesn't return an error if amount equals the given amount",
			fields: fields{
				Data: xRQTxnVerificationResponseData{Amount: 1000},
			},
			args:    args{amt: 1000},
			wantErr: false,
		},
		{
			name: "doesn't return an error if amount is greater than the given amount",
			fields: fields{
				Data: xRQTxnVerificationResponseData{Amount: 2000},
			},
			args:    args{amt: 1000},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &XRQTxnVerificationResponse{
				Data:    tt.fields.Data,
				Message: tt.fields.Message,
				Status:  tt.fields.Status,
			}
			if err := resp.VerifyAmount(tt.args.amt); (err != nil) != tt.wantErr {
				t.Errorf("XRQTxnVerificationResponse.VerifyAmount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestXRQTxnVerificationResponse_VerifyChargeResponseValue(t *testing.T) {
	type fields struct {
		Data    xRQTxnVerificationResponseData
		Message string
		Status  string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "returns an error if charge response is not 00 or 0",
			fields: fields{
				Data: xRQTxnVerificationResponseData{Chargecode: "404"},
			},
			wantErr: true,
		},
		{
			name: "doesn't return an error if charge response is 0",
			fields: fields{
				Data: xRQTxnVerificationResponseData{Chargecode: "00"},
			},
			wantErr: false,
		},
		{
			name: "doesn't return an error if charge response is 00",
			fields: fields{
				Data: xRQTxnVerificationResponseData{Chargecode: "00"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &XRQTxnVerificationResponse{
				Data:    tt.fields.Data,
				Message: tt.fields.Message,
				Status:  tt.fields.Status,
			}
			if err := resp.VerifyChargeResponseValue(); (err != nil) != tt.wantErr {
				t.Errorf("XRQTxnVerificationResponse.VerifyChargeResponseValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestXRQTxnVerificationResponse_VerifyReference(t *testing.T) {
	type fields struct {
		Data    xRQTxnVerificationResponseData
		Message string
		Status  string
	}
	type args struct {
		ref string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "returns an error if flwref doesn't match the given ref",
			fields: fields{
				Data: xRQTxnVerificationResponseData{Flwref: "another-ref"},
			},
			args:    args{ref: "my-ref"},
			wantErr: true,
		},
		{
			name: "doesn't return an error if flwref matches the given ref",
			fields: fields{
				Data: xRQTxnVerificationResponseData{Flwref: "my-ref"},
			},
			args:    args{ref: "my-ref"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &XRQTxnVerificationResponse{
				Data:    tt.fields.Data,
				Message: tt.fields.Message,
				Status:  tt.fields.Status,
			}
			if err := resp.VerifyReference(tt.args.ref); (err != nil) != tt.wantErr {
				t.Errorf("XRQTxnVerificationResponse.VerifyReference() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
