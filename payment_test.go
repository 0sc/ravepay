package rave

import (
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestPaymentVerificationChecklist_verifyPaymentAmount(t *testing.T) {
	type fields struct {
		Amount              int
		PaymentVerification *PaymentVerification
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "it adds a PaymentVerificationAmountError if payment amount is less than checklist amount",
			fields: fields{
				Amount: 100,
				PaymentVerification: &PaymentVerification{
					Data: paymentVerificationData{Amount: 90},
				},
			},
			want: []string{PaymentVerificationAmountError},
		},
		{
			name: "it doesn't add a PaymentVerificationAmountError if payment amount is equal to checklist amount",
			fields: fields{
				Amount: 100,
				PaymentVerification: &PaymentVerification{
					Data: paymentVerificationData{Amount: 100},
				},
			},
			want: []string{},
		},
		{
			name: "it doesn't add a PaymentVerificationAmountError if payment amount is more than checklist amount",
			fields: fields{
				Amount: 100,
				PaymentVerification: &PaymentVerification{
					Data: paymentVerificationData{Amount: 110},
				},
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pvc := &PaymentVerificationChecklist{
				Amount:              tt.fields.Amount,
				PaymentVerification: tt.fields.PaymentVerification,
			}
			pvc.verifyPaymentAmount()
			if reflect.DeepEqual(pvc.Errors, tt.want) {
				t.Errorf("expected %+v, got %+v", tt.want, pvc.Errors)
			}
		})
	}
}

func TestPaymentVerificationChecklist_verifyTransactionCurrency(t *testing.T) {
	type fields struct {
		PaymentVerification *PaymentVerification
		TransactionCurrency string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "it does not add a PaymentVerificationTransactionCurrencyMismatch if payment currency is the same as the checklist transaction currency",
			fields: fields{
				TransactionCurrency: "NGN",
				PaymentVerification: &PaymentVerification{
					Data: paymentVerificationData{TransactionCurrency: "NGN"},
				},
			},
			want: []string{},
		},
		{
			name: "it adds a PaymentVerificationTransactionCurrencyMismatch if payment currency is different from the checklist transaction currency",
			fields: fields{
				TransactionCurrency: "NGN",
				PaymentVerification: &PaymentVerification{
					Data: paymentVerificationData{TransactionCurrency: "USD"},
				},
			},
			want: []string{PaymentVerificationTransactionCurrencyMismatch},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pvc := &PaymentVerificationChecklist{
				PaymentVerification: tt.fields.PaymentVerification,
				TransactionCurrency: tt.fields.TransactionCurrency,
			}
			pvc.verifyTransactionCurrency()
			if reflect.DeepEqual(pvc.Errors, tt.want) {
				t.Errorf("expected %+v, got %+v", tt.want, pvc.Errors)
			}
		})
	}
}

func TestPaymentVerificationChecklist_verifyChargeResponseValue(t *testing.T) {
	type fields struct {
		PaymentVerification *PaymentVerification
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "it adds a PaymentVerificationChargeResponseValueError if charge response is not 00 or 0",
			fields: fields{
				PaymentVerification: &PaymentVerification{
					Data: paymentVerificationData{
						FlwMeta: FlwMeta{ChargeResponse: "404"},
					},
				},
			},
			want: []string{PaymentVerificationChargeResponseValueError},
		},
		{
			name: "it doesn't add a PaymentVerificationChargeResponseValueError if charge response is 00",
			fields: fields{
				PaymentVerification: &PaymentVerification{
					Data: paymentVerificationData{
						FlwMeta: FlwMeta{ChargeResponse: "00"},
					},
				},
			},
			want: []string{},
		},
		{
			name: "it doesn't add a PaymentVerificationChargeResponseValueError if charge response is 0",
			fields: fields{
				PaymentVerification: &PaymentVerification{
					Data: paymentVerificationData{
						FlwMeta: FlwMeta{ChargeResponse: "0"},
					},
				},
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pvc := &PaymentVerificationChecklist{
				PaymentVerification: tt.fields.PaymentVerification,
			}
			pvc.verifyChargeResponseValue()
			if reflect.DeepEqual(pvc.Errors, tt.want) {
				t.Errorf("expected %+v, got %+v", tt.want, pvc.Errors)
			}
		})
	}
}

func TestPaymentVerificationChecklist_verifyTransactionStatus(t *testing.T) {
	type fields struct {
		PaymentVerification *PaymentVerification
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "it adds a PaymentVerificationStatusMismatch error if status is not 'success''",
			fields: fields{
				PaymentVerification: &PaymentVerification{Status: "not-success"},
			},
			want: []string{PaymentVerificationStatusMismatch},
		},
		{
			name: "it doesn't add a PaymentVerificationStatusMismatch error if status is success",
			fields: fields{
				PaymentVerification: &PaymentVerification{Status: "success"},
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pvc := &PaymentVerificationChecklist{
				PaymentVerification: tt.fields.PaymentVerification,
			}
			pvc.verifyTransactionStatus()
			if reflect.DeepEqual(pvc.Errors, tt.want) {
				t.Errorf("expected %+v, got %+v", tt.want, pvc.Errors)
			}
		})
	}
}

func TestPaymentVerificationChecklist_verifyTransactionReference(t *testing.T) {
	type fields struct {
		FlwRef              string
		PaymentVerification *PaymentVerification
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "it adds a PaymentVerificationTransactionRefMismatch error if flwref doesn't match",
			fields: fields{
				FlwRef: "my-ref",
				PaymentVerification: &PaymentVerification{
					Data: paymentVerificationData{FlwRef: "another-ref"},
				},
			},
			want: []string{PaymentVerificationTransactionRefMismatch},
		},
		{
			name: "it doesn't add a PaymentVerificationTransactionRefMismatch error if flwref matches",
			fields: fields{
				FlwRef: "my-ref",
				PaymentVerification: &PaymentVerification{
					Data: paymentVerificationData{FlwRef: "my-ref"},
				},
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pvc := &PaymentVerificationChecklist{
				FlwRef:              tt.fields.FlwRef,
				PaymentVerification: tt.fields.PaymentVerification,
			}
			pvc.verifyTransactionReference()
			if reflect.DeepEqual(pvc.Errors, tt.want) {
				t.Errorf("expected %+v, got %+v", tt.want, pvc.Errors)
			}
		})
	}
}

func TestPaymentVerificationChecklist_Verify(t *testing.T) {
	handler := &testServer{}
	server := httptest.NewServer(handler)
	defer server.Close()

	type fields struct {
		Amount              int
		FlwRef              string
		TransactionCurrency string
	}
	tests := []struct {
		name     string
		fields   fields
		respBody string
		want     bool
	}{
		{
			name:     "it returns false if payment verification fails",
			respBody: noTransactionFoundVerifyPaymentResponse,
			fields: fields{
				Amount:              300,
				FlwRef:              "FLW-MOCK-09805abc71c5eebf80bb899183475fe3",
				TransactionCurrency: "NGN",
			},
			want: false,
		},
		{
			name:     "it returns true if account payment verification passes",
			respBody: successfulAccountVerifyPaymentResponse,
			fields: fields{
				Amount:              300,
				FlwRef:              "FLW-MOCK-09805abc71c5eebf80bb899183475fe3",
				TransactionCurrency: "NGN",
			},
			want: true,
		},
		{
			name:     "it returns true if card payment verification passes",
			respBody: successfulCardVerifyPaymentResponse,
			fields: fields{
				Amount:              300,
				FlwRef:              "FLW-MOCK-09805abc71c5eebf80bb899183475fe3",
				TransactionCurrency: "NGN",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pvc := &PaymentVerificationChecklist{
				Amount:                 tt.fields.Amount,
				FlwRef:                 tt.fields.FlwRef,
				PaymentVerificationURL: server.URL,
				TransactionCurrency:    tt.fields.TransactionCurrency,
			}
			handler.resp = []byte(tt.respBody)
			if got := pvc.Verify(); got != tt.want {
				t.Errorf("PaymentVerificationChecklist.Verify() = %v, want %v", got, tt.want)
			}
		})
	}
}

type testServer struct {
	resp []byte
}

func (ts *testServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	_, err := w.Write(ts.resp)
	if err != nil {
		log.Println("Error occurred encoding payload", err, ts.resp)
	}
}

var successfulAccountVerifyPaymentResponse = `{
  "status": "success",
  "message": "Tx Fetched",
  "data": {
    "id": 56673,
    "tx_ref": "5f06e536-e981-4f52-9e0b-336600798dc5",
    "order_ref": "URF_1512654631908_3202535",
    "flw_ref": "FLW-MOCK-09805abc71c5eebf80bb899183475fe3",
    "transaction_type": "debit",
    "settlement_token": null,
    "rave_ref": "RV3151265463164150B7E2D76C",
    "transaction_processor": "FLW",
    "status": "successful",
    "chargeback_status": null,
    "ip": "::ffff:127.0.0.1",
    "device_fingerprint": "352693081974640",
    "cycle": "one-time",
    "narration": "FLW-PBF CARD Transaction ",
    "amount": 300,
    "appfee": 0,
    "merchantfee": 0,
    "markupFee": null,
    "merchantbearsfee": 1,
    "charged_amount": 300,
    "transaction_currency": "NGN",
    "system_type": null,
    "payment_entity": "card",
    "payment_id": "356",
    "fraud_status": "ok",
    "charge_type": "normal",
    "is_live": 0,
    "createdAt": "2017-12-07T13:50:33.000Z",
    "updatedAt": "2017-12-07T13:50:40.000Z",
    "deletedAt": null,
    "merchant_id": 134,
    "addon_id": 3,
    "customer": {
      "id": 9520,
      "phone": null,
      "fullName": "Hamza Fetuga",
      "customertoken": null,
      "email": "hfetuga@gmail.com",
      "createdAt": "2017-12-07T13:50:31.000Z",
      "updatedAt": "2017-12-07T13:50:31.000Z",
      "deletedAt": null,
      "AccountId": 134
    },
    "card": {
      "expirymonth": "09",
      "expiryyear": "20",
      "cardBIN": "424242",
      "last4digits": "4242",
      "brand": "VISA  CREDIT",
      "card_tokens": [
        {
          "shortcode": "ceccd",
          "embedtoken": "flw-t0-f6f915f53a094671d98560272572993e-m03k"
        }
      ]
    },
    "meta": [
      {
        "id": 4813,
        "metaname": "savings_id",
        "metavalue": "8ca5db60-db55-11e7-837e-bfb21bab8e55",
        "createdAt": "2017-12-07T13:50:33.000Z",
        "updatedAt": "2017-12-07T13:50:33.000Z",
        "deletedAt": null,
        "getpaidTransactionId": 58960
      }
    ],
    "flwMeta": {
      "chargeResponse": "00",
      "chargeResponseMessage": "Success-Pending-otp-validation",
      "VBVRESPONSEMESSAGE": "Approved. Successful",
      "VBVRESPONSECODE": "00",
      "ACCOUNTVALIDATIONRESPMESSAGE": null,
      "ACCOUNTVALIDATIONRESPONSECODE": "RN1512654631916"
    }
  }
}`

var successfulCardVerifyPaymentResponse = `{
  "status": "success",
  "message": "Tx Fetched",
  "data": {
    "id": 56465,
    "tx_ref": "BR-1512550521352-41424",
    "order_ref": null,
    "flw_ref": "ACHG-1512550576634",
    "transaction_type": "debit",
    "settlement_token": null,
    "rave_ref": null,
    "transaction_processor": "FLW",
    "status": "successful",
    "chargeback_status": null,
    "ip": "41.223.47.82",
    "device_fingerprint": "689de87638deca2ca28dc8bb16f39581",
    "cycle": "one-time",
    "narration": "Synergy Group",
    "amount": 100,
    "appfee": null,
    "merchantfee": null,
    "markupFee": null,
    "merchantbearsfee": 1,
    "charged_amount": 100,
    "transaction_currency": "NGN",
    "system_type": null,
    "payment_entity": "account",
    "payment_id": "16",
    "fraud_status": "ok",
    "charge_type": "normal",
    "is_live": 0,
    "createdAt": "2017-12-06T08:56:31.000Z",
    "updatedAt": "2017-12-06T08:56:38.000Z",
    "deletedAt": null,
    "merchant_id": 134,
    "addon_id": 3,
    "customer": {
      "id": 3096,
      "phone": "N/A",
      "fullName": "Somto ALI",
      "customertoken": null,
      "email": "alisomto@yahoo.com",
      "createdAt": "2017-09-06T15:48:07.000Z",
      "updatedAt": "2017-09-06T15:48:07.000Z",
      "deletedAt": null,
      "AccountId": 134
    },
    "account": {
      "id": 16,
      "account_number": "0690000004",
      "account_bank": "044",
      "first_name": "NO-NAME",
      "last_name": "NO-LNAME",
      "account_is_blacklisted": 0,
      "createdAt": "2017-01-27T09:41:58.000Z",
      "updatedAt": "2017-12-07T10:46:35.000Z",
      "deletedAt": null,
      "account_token": {
        "token": "flw-t0876849e016386b2d-k3n-mock"
      }
    },
    "meta": [
      {
        "id": 4786,
        "metaname": "brcrypt",
        "metavalue": "oNaBZqF8kPO9LL4EVtv4lIN00N7v9lUEePanT1sUAw7pyPcQFFt0bAbNUMS3Tmqx",
        "createdAt": "2017-12-06T08:56:16.000Z",
        "updatedAt": "2017-12-06T08:56:16.000Z",
        "deletedAt": null,
        "getpaidTransactionId": 58733
      }
    ],
    "flwMeta": {
      "chargeResponse": "00",
      "chargeResponseMessage": "Pending OTP validation",
      "VBVRESPONSEMESSAGE": "N/A",
      "VBVRESPONSECODE": "N/A",
      "ACCOUNTVALIDATIONRESPMESSAGE": "Approved Or Completed Successfully",
      "ACCOUNTVALIDATIONRESPONSECODE": "00"
    }
  }
}`

var noTransactionFoundVerifyPaymentResponse = `{
  "status": "error",
  "message": "No transaction found",
  "data": {
    "code": "NO TX",
    "message": "No transaction found"
  }
}`
