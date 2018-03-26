package rave

import (
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestCapturePreAuthPayment(t *testing.T) {
	handler := &testServer{}
	server := httptest.NewServer(handler)
	defer server.Close()

	type args struct {
		ref string
	}
	tests := []struct {
		name     string
		args     args
		respBody string
		want     *ChargeResponse
		wantErr  bool
	}{
		{
			name:     "returns the capture response",
			respBody: successfulPreAuthPaymentCaptureResponse,
			args: args{
				ref: "FLW-MOCK-839c1abc23b6a4bbb9da807d54c5bbda",
			},
			want: &ChargeResponse{
				Status:  "success",
				Message: "Capture complete",
				Data: chargeResponseData{
					AccountID:     48,
					Amount:        20,
					Appfee:        0.25,
					AuthModelUsed: "NOAUTH",
					Authurl:       "N/A",
					// ChargedAmount:         20.25,
					ChargeResponseMessage: "Approved",
					ChargeResponseCode:    "00",
					ChargeType:            "preauth",
					CreatedAt:             "2017-10-26T03:56:42.000Z",
					Currency:              "NGN",
					CustomerID:            88067,
					Cycle:                 "one-time",
					DeviceFingerprint:     "69e6b7f0b72037aa8428b70fbe03986c",
					FlwRef:                "FLW-MOCK-839c1abc23b6a4bbb9da807d54c5bbda",
					FraudStatus:           "ok",
					ID:                    194211,
					IP:                    "::ffff:10.37.225.74",
					Narration:             "FLW-PBF CARD Transaction ",
					PaymentID:             "47433",
					PaymentType:           "card",
					RedirectURL:           "http://127.0.0",
					Status:                "successful",
					TxRef:                 "MC-1508990174050",
					UpdatedAt:             "2017-10-26T04:20:21.000Z",
					Vbvrespcode:           "00",
					Vbvrespmessage:        "Approved",
					Customer: Customer{
						AccountID: 48,
						CreatedAt: "2017-10-26T03:39:45.000Z",
						Email:     "tester@flutter.co",
						FullName:  "temi desola",
						ID:        88067,
						Phone:     "08056552980",
						UpdatedAt: "2017-10-26T03:39:45.000Z",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			capturePreAuthURL = server.URL
			handler.resp = []byte(tt.respBody)

			got, err := CapturePreAuthPayment(tt.args.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("CapturePreAuthPayment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CapturePreAuthPayment() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestRefundPreAuthPayment(t *testing.T) {
	handler := &testServer{}
	server := httptest.NewServer(handler)
	defer server.Close()

	type args struct {
		ref string
	}
	tests := []struct {
		name     string
		args     args
		respBody string
		want     *PreAuthResponse
		wantErr  bool
	}{
		{
			name: "returns the refund preauth payment response if successful",
			args: args{
				ref: "FLW-MOCK-839c1abc23b6a4bbb9da807d54c5bbda",
			},
			respBody: successfulPreAuthPaymentRefundResponse,
			want: &PreAuthResponse{
				Message: "Refund or void complete",
				Status:  "success",
			},
			wantErr: false,
		},
		{
			name: "returns the refund preauth payment response if unsuccessful",
			args: args{
				ref: "FLW-MOCK-839c1abc23b6a4bbb9da807d54c5bbda",
			},
			respBody: failedPreAuthPaymentRefundResponse,
			want: &PreAuthResponse{
				Message: "No transaction found",
				Status:  "error",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			voidorRefundPreAuthURL = server.URL
			handler.resp = []byte(tt.respBody)

			got, err := RefundPreAuthPayment(tt.args.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("RefundPreAuthPayment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.Message != tt.want.Message || got.Status != tt.want.Status {
				t.Errorf("RefundPreAuthPayment() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestVoidPreAuthPayment(t *testing.T) {
	handler := &testServer{}
	server := httptest.NewServer(handler)
	defer server.Close()

	type args struct {
		ref string
	}
	tests := []struct {
		name     string
		args     args
		respBody string
		want     *PreAuthResponse
		wantErr  bool
	}{
		{
			name: "returns the refund preauth payment response if successful",
			args: args{
				ref: "FLW-MOCK-839c1abc23b6a4bbb9da807d54c5bbda",
			},
			respBody: successfulPreAuthPaymentRefundResponse,
			want: &PreAuthResponse{
				Message: "Refund or void complete",
				Status:  "success",
			},
			wantErr: false,
		},
		{
			name: "returns the refund preauth payment response if unsuccessful",
			args: args{
				ref: "FLW-MOCK-839c1abc23b6a4bbb9da807d54c5bbda",
			},
			respBody: failedPreAuthPaymentRefundResponse,
			want: &PreAuthResponse{
				Message: "No transaction found",
				Status:  "error",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			voidorRefundPreAuthURL = server.URL
			handler.resp = []byte(tt.respBody)

			got, err := VoidPreAuthPayment(tt.args.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("RefundPreAuthPayment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.Message != tt.want.Message || got.Status != tt.want.Status {
				t.Errorf("RefundPreAuthPayment() = %v, want %v", got, tt.want)
			}
		})
	}
}
