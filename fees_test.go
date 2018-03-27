package rave

import (
	"net/http/httptest"
	"testing"
)

func TestGetFee(t *testing.T) {
	handler := &testServer{}
	server := httptest.NewServer(handler)
	defer server.Close()
	handler.resp = []byte(getFeeResponse)
	feeURL = server.URL

	type args struct {
		p *GetFeeRequest
	}
	tests := []struct {
		name       string
		args       args
		want       *GetFeeResponse
		wantPubkey string
		wantErr    bool
	}{
		{
			name: "returns the GetFee response",
			args: args{
				p: &GetFeeRequest{},
			},
			want: &GetFeeResponse{
				Status:  "success",
				Message: "Charged fee",
			},
			wantPubkey: PublicKey,
			wantErr:    false,
		},
		{
			name: "sets the PBFPubKey if empty",
			args: args{
				p: &GetFeeRequest{},
			},
			want: &GetFeeResponse{
				Status:  "success",
				Message: "Charged fee",
			},
			wantPubkey: PublicKey,
			wantErr:    false,
		},
		{
			name: "doesn't override the PBFPubKey if present",
			args: args{
				p: &GetFeeRequest{PBFPubKey: "my-pub-key"},
			},
			want: &GetFeeResponse{
				Status:  "success",
				Message: "Charged fee",
			},
			wantPubkey: "my-pub-key",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFee(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFee() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if key := tt.args.p.PBFPubKey; key != tt.wantPubkey {
				t.Errorf("PBFPubkey = %s, want %s", key, tt.wantPubkey)
			}
			if got.Message != tt.want.Message || got.Status != tt.want.Status {
				t.Errorf("GetFee() = %v, want %v", got, tt.want)
			}
		})
	}
}
