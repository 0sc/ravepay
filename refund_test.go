package rave

import (
	"net/http/httptest"
	"testing"
)

func TestRefund(t *testing.T) {
	handler := &testServer{}
	server := httptest.NewServer(handler)
	defer server.Close()
	rTxnURL = server.URL

	type args struct {
		ref string
	}
	tests := []struct {
		name     string
		args     args
		respBody string
		want     *RefundTxnResponse
		wantErr  bool
	}{
		{
			name: "returns the refund response if successful",
			args: args{
				ref: "some-txn-ref",
			},
			respBody: successfulRefundTxnResponse,
			want: &RefundTxnResponse{
				Status:  "success",
				Message: "Refunded",
			},
			wantErr: false,
		},
		{
			name: "returns the refund response if unsuccessful",
			args: args{
				ref: "some-txn-ref",
			},
			respBody: failedRefundTxnResponse,
			want: &RefundTxnResponse{
				Status:  "error",
				Message: "No transaction found",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler.resp = []byte(tt.respBody)

			got, err := Refund(tt.args.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("Refund() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Message != tt.want.Message || got.Status != tt.want.Status {
				t.Errorf("RefundPreAuthPayment() = %v, want %v", got, tt.want)
			}
		})
	}
}
