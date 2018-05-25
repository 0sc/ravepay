package ravepay

import (
	"net/http/httptest"
	"testing"
)

func TestForexRate(t *testing.T) {
	handler := &testServer{}
	server := httptest.NewServer(handler)
	defer server.Close()
	fxURL = server.URL

	type args struct {
		fxp *ForexParams
	}
	tests := []struct {
		name       string
		args       args
		respBody   string
		want       *ForexResponse
		wantSecKey string
		wantErr    bool
	}{
		{
			name: "returns the forex response with amount",
			args: args{
				fxp: &ForexParams{},
			},
			respBody: forexRateWithAmountResponse,
			want: &ForexResponse{
				Message: "Rate Fetched",
				Status:  "success",
			},
			wantSecKey: SecretKey,
			wantErr:    false,
		},
		{
			name: "returns the forex response without amount",
			args: args{
				fxp: &ForexParams{},
			},
			respBody: forexRateWithoutAmountResponse,
			want: &ForexResponse{
				Message: "Rate Fetched",
				Status:  "success",
			},
			wantSecKey: SecretKey,
			wantErr:    false,
		},
		{
			name: "sets the secret key if empty",
			args: args{
				fxp: &ForexParams{},
			},
			respBody: forexRateWithAmountResponse,
			want: &ForexResponse{
				Message: "Rate Fetched",
				Status:  "success",
			},
			wantSecKey: SecretKey,
			wantErr:    false,
		},
		{
			name: "doesn't override the secret key if present",
			args: args{
				fxp: &ForexParams{SecKey: "my-sec-key"},
			},
			respBody: forexRateWithAmountResponse,
			want: &ForexResponse{
				Message: "Rate Fetched",
				Status:  "success",
			},
			wantSecKey: "my-sec-key",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler.resp = []byte(tt.respBody)

			got, err := ForexRate(tt.args.fxp)
			if (err != nil) != tt.wantErr {
				t.Errorf("ForexRate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if key := tt.args.fxp.SecKey; key != tt.wantSecKey {
				t.Errorf("PBFPubkey = %s, want %s", key, tt.wantSecKey)
			}
			if got.Message != tt.want.Message || got.Status != tt.want.Status {
				t.Errorf("ForexRate() = %v, want %v", got, tt.want)
			}
		})
	}
}
