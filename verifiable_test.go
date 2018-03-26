package rave

import (
	"fmt"
	"net/http/httptest"
	"reflect"
	"testing"
)

type testVerifiable struct {
	statusErr, currencyErr, amountErr, rValueErr, refErr error
}

func (tv testVerifiable) VerifyStatus() error {
	return tv.statusErr
}
func (tv testVerifiable) VerifyCurrency(string) error {
	return tv.currencyErr
}
func (tv testVerifiable) VerifyAmount(int) error {
	return tv.amountErr
}
func (tv testVerifiable) VerifyChargeResponseValue() error {
	return tv.rValueErr
}
func (tv testVerifiable) VerifyReference(string) error {
	return tv.refErr
}

func TestTxnVerificationChecklist_Verify(t *testing.T) {
	err := fmt.Errorf("generic error :)")

	type args struct {
		v Verifiable
	}
	tests := []struct {
		name string
		args args
		want []error
	}{
		{
			name: "has error if all veriable verification returns errors",
			args: args{
				v: testVerifiable{
					statusErr:   err,
					currencyErr: err,
					amountErr:   err,
					rValueErr:   err,
					refErr:      err,
				},
			},
			want: []error{err, err, err, err, err},
		},
		{
			name: "has error if veriable's VerifyCurrency returns errors",
			args: args{
				v: testVerifiable{currencyErr: err},
			},
			want: []error{err},
		},
		{
			name: "has error if veriable's VerifyAmount returns errors",
			args: args{
				v: testVerifiable{amountErr: err},
			},
			want: []error{err},
		},
		{
			name: "has error if veriable's VerifyReference returns errors",
			args: args{
				v: testVerifiable{refErr: err},
			},
			want: []error{err},
		},
		{
			name: "has error if veriable's VerifyStatus returns errors",
			args: args{
				v: testVerifiable{statusErr: err},
			},
			want: []error{err},
		},
		{
			name: "has error if veriable's VerifyChargeResponseValue returns errors",
			args: args{
				v: testVerifiable{rValueErr: err},
			},
			want: []error{err},
		},
		{
			name: "has no error if no veriable verification returns an error",
			args: args{
				v: testVerifiable{},
			},
			want: []error{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tvc := &TxnVerificationChecklist{}
			if got := tvc.Verify(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TxnVerificationChecklist.Verify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTxnVerificationChecklist_VerifyTransaction(t *testing.T) {
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
		want     *TxnVerificationResponse
		want1    []error
	}{
		{
			name:     "returns error if verification fails",
			respBody: noTransactionFoundVerifyPaymentResponse,
			fields: fields{
				Amount:              300,
				FlwRef:              "FLW-MOCK-09805abc71c5eebf80bb899183475fe3",
				TransactionCurrency: "NGN",
			},
			want: &TxnVerificationResponse{
				Status:  "error",
				Message: "No transaction found",
				Data: txnVerificationResponseData{
					Code:    "NO TX",
					Message: "No transaction found",
				},
			},
			want1: []error{
				fmt.Errorf("CurrencyVerificationFailed: expected NGN but got "),
				fmt.Errorf("AmountVerificationFailed: expected 300 but got 0"),
				fmt.Errorf("FlwRefVerificationFailed: expected FLW-MOCK-09805abc71c5eebf80bb899183475fe3 but got "),
				fmt.Errorf("StatusVerificationFailed: expected success but got error"),
				fmt.Errorf("ChargeResponseVerificationFailed: expected 00 or 0 but got "),
			},
		},
		{
			name:     "returns no error if account verification succeeds",
			respBody: successfulAccountVerifyPaymentResponse,
			fields: fields{
				Amount:              100,
				FlwRef:              "ACHG-1512550576634",
				TransactionCurrency: "NGN",
			},
			want: &TxnVerificationResponse{
				Status:  "success",
				Message: "Tx Fetched",
				Data: txnVerificationResponseData{
					ID:                   56465,
					TxRef:                "BR-1512550521352-41424",
					FlwRef:               "ACHG-1512550576634",
					TransactionType:      "debit",
					TransactionProcessor: "FLW",
					Status:               "successful",
					IP:                   "41.223.47.82",
					DeviceFingerprint:    "689de87638deca2ca28dc8bb16f39581",
					Cycle:                "one-time",
					Narration:            "Synergy Group",
					Amount:               100,
					Merchantbearsfee:     1,
					ChargedAmount:        100,
					TransactionCurrency:  "NGN",
					PaymentEntity:        "account",
					PaymentID:            "16",
					FraudStatus:          "ok",
					ChargeType:           "normal",
					CreatedAt:            "2017-12-06T08:56:31.000Z",
					UpdatedAt:            "2017-12-06T08:56:38.000Z",
					MerchantID:           134,
					AddonID:              3,
					Customer: Customer{
						ID:        3096,
						FullName:  "Somto ALI",
						Phone:     "N/A",
						Email:     "alisomto@yahoo.com",
						CreatedAt: "2017-09-06T15:48:07.000Z",
						UpdatedAt: "2017-09-06T15:48:07.000Z",
						AccountID: 134,
					},
					Account: Account{
						ID:            16,
						AccountNumber: "0690000004",
						AccountBank:   "044",
						FirstName:     "NO-NAME",
						LastName:      "NO-LNAME",
						CreatedAt:     "2017-01-27T09:41:58.000Z",
						UpdatedAt:     "2017-12-07T10:46:35.000Z",
					},
					FlwMeta: FlwMeta{
						ChargeResponse:                "00",
						ChargeResponseMessage:         "Pending OTP validation",
						VBVRESPONSECODE:               "N/A",
						VBVRESPONSEMESSAGE:            "N/A",
						ACCOUNTVALIDATIONRESPMESSAGE:  "Approved Or Completed Successfully",
						ACCOUNTVALIDATIONRESPONSECODE: "00",
					},
				},
			},
			want1: []error{},
		},
		{
			name:     "returns no error if card verification succeeds",
			respBody: successfulCardVerifyPaymentResponse,
			fields: fields{
				Amount:              300,
				FlwRef:              "FLW-MOCK-09805abc71c5eebf80bb899183475fe3",
				TransactionCurrency: "NGN",
			},
			want: &TxnVerificationResponse{
				Status:  "success",
				Message: "Tx Fetched",
				Data: txnVerificationResponseData{
					ID:                   56673,
					TxRef:                "5f06e536-e981-4f52-9e0b-336600798dc5",
					OrderRef:             "URF_1512654631908_3202535",
					FlwRef:               "FLW-MOCK-09805abc71c5eebf80bb899183475fe3",
					TransactionType:      "debit",
					RaveRef:              "RV3151265463164150B7E2D76C",
					TransactionProcessor: "FLW",
					Status:               "successful",
					IP:                   "::ffff:127.0.0.1",
					DeviceFingerprint:    "352693081974640",
					Cycle:                "one-time",
					Narration:            "FLW-PBF CARD Transaction ",
					Amount:               300,
					Merchantbearsfee:     1,
					ChargedAmount:        300,
					TransactionCurrency:  "NGN",
					PaymentEntity:        "card",
					PaymentID:            "356",
					FraudStatus:          "ok",
					ChargeType:           "normal",
					CreatedAt:            "2017-12-07T13:50:33.000Z",
					UpdatedAt:            "2017-12-07T13:50:40.000Z",
					MerchantID:           134,
					AddonID:              3,
					Customer: Customer{
						ID:        9520,
						FullName:  "Hamza Fetuga",
						Email:     "hfetuga@gmail.com",
						CreatedAt: "2017-12-07T13:50:31.000Z",
						UpdatedAt: "2017-12-07T13:50:31.000Z",
						AccountID: 134,
					},
					Card: Card{
						Expirymonth: "09",
						Expiryyear:  "20",
						CardBIN:     "424242",
						Last4digits: "4242",
						Brand:       "VISA  CREDIT",
					},
					FlwMeta: FlwMeta{
						ChargeResponse:                "00",
						ChargeResponseMessage:         "Success-Pending-otp-validation",
						VBVRESPONSECODE:               "00",
						VBVRESPONSEMESSAGE:            "Approved. Successful",
						ACCOUNTVALIDATIONRESPONSECODE: "RN1512654631916",
					},
				},
			},
			want1: []error{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tvc := &TxnVerificationChecklist{
				Amount:              tt.fields.Amount,
				FlwRef:              tt.fields.FlwRef,
				VerificationURL:     server.URL,
				TransactionCurrency: tt.fields.TransactionCurrency,
			}

			handler.resp = []byte(tt.respBody)

			got, got1 := tvc.VerifyTransaction()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TxnVerificationChecklist.VerifyTransaction() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("TxnVerificationChecklist.VerifyTransaction() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}

	// Test it sets a verification URL if empty
	tvc := &TxnVerificationChecklist{VerificationURL: ""}
	baseURL = server.URL
	tvc.VerifyTransaction()

	if tvc.VerificationURL == "" {
		t.Errorf("expected TxnVerificationChecklist.VerifyTransaction().VerificationURL not to be mepty")
	}
}

func TestTxnVerificationChecklist_VerifyXRequeryTransaction(t *testing.T) {
	handler := &testServer{}
	server := httptest.NewServer(handler)
	defer server.Close()

	type fields struct {
		Amount              int
		Flwref              string
		LastAttempt         string
		OnlySuccessful      string
		VerificationURL     string
		TransactionCurrency string
		Txref               string
	}
	tests := []struct {
		name     string
		fields   fields
		respBody string
		want     *XRQTxnVerificationResponse
		want1    []error
	}{
		{
			name: "returns errors if verification fails",
			fields: fields{
				Amount:              10000,
				Flwref:              "FLW-MOCK-09805abc71c5eebf80bb899183475fe3",
				LastAttempt:         "1",
				OnlySuccessful:      "1",
				TransactionCurrency: "NGN",
				Txref:               "OH-AAED44",
				VerificationURL:     server.URL,
			},
			respBody: xRQSuccessfulVerificationResponse,
			want: &XRQTxnVerificationResponse{
				Status:  "success",
				Message: "Tx Fetched",
				Data: xRQTxnVerificationResponseData{
					Accountid:                     134,
					Acctalias:                     "temi",
					Acctbearsfeeattransactiontime: 1,
					Acctbusinessname:              "Synergy Group",
					Acctcontactperson:             "Desola Ade",
					Acctcountry:                   "NG",
					Acctisliveapproved:            0,
					Acctparent:                    1,
					Acctvpcmerchant:               "N/A",
					Amount:                        8150,
					Amountsettledforthistransaction: 8150,
					Authmodel:                       "PIN",
					Authurl:                         "http://flw-pms-dev.eu-west-1.elasticbeanstalk.com/mockvbvpage?ref=FLW-MOCK-5980e4f35eb54158fc296a1f22b1ff88&code=00&message=Approved. Successful&receiptno=RN1504378687808",
					Chargecode:                      "00",
					Chargedamount:                   8150,
					Chargemessage:                   "Success-Pending-otp-validation",
					Chargetype:                      "normal",
					Created:                         "2017-09-02T18:58:07.000Z",
					Createdday:                      6,
					Createddayispublicholiday:       0,
					Createddayname:                  "SATURDAY",
					Createdhour:                     18,
					Createdminute:                   58,
					Createdmonth:                    8,
					Createdmonthname:                "SEPTEMBER",
					Createdpmam:                     "pm",
					Createdquarter:                  3,
					Createdweek:                     35,
					Createdyear:                     2017,
					Createdyearisleap:               false,
					Currency:                        "NGN",
					Custcreated:                     "2017-09-02T18:58:07.000Z",
					Custemail:                       "bakare.wilmot@thegiggroupng.com",
					Custemailprovider:               "COMPANY EMAIL",
					Custname:                        "Anonymous customer",
					Custnetworkprovider:             "N/A",
					Customerid:                      2777,
					Cycle:                           "one-time",
					Devicefingerprint:               "ee09e26e8a76cfa35e64688c95145c2f",
					Flwref:                          "FLW-MOCK-5980e4f35eb54158fc296a1f22b1ff88",
					Fraudstatus:                     "ok",
					IP:                              "197.210.25.161",
					Narration:                       "FLW-PBF CARD Transaction ",
					Orderref:                        "URF_1504378687803_68035",
					Paymentid:                       "2",
					Paymenttype:                     "card",
					Status:                          "successful",
					Txid:                            32458,
					Txref:                           "OH-AAED44",
					Vbvcode:                         "00",
					Vbvmessage:                      "successful",
				},
			},
			want1: []error{
				fmt.Errorf("AmountVerificationFailed: expected 10000 but got 8150"),
				fmt.Errorf("FlwRefVerificationFailed: expected FLW-MOCK-09805abc71c5eebf80bb899183475fe3 but got FLW-MOCK-5980e4f35eb54158fc296a1f22b1ff88"),
			},
		},
		{
			name: "returns no errors if verification succeeds",
			fields: fields{
				Amount:              8150,
				Flwref:              "FLW-MOCK-5980e4f35eb54158fc296a1f22b1ff88",
				LastAttempt:         "1",
				OnlySuccessful:      "1",
				TransactionCurrency: "NGN",
				Txref:               "OH-AAED44",
				VerificationURL:     server.URL,
			},
			respBody: xRQSuccessfulVerificationResponse,
			want: &XRQTxnVerificationResponse{
				Status:  "success",
				Message: "Tx Fetched",
				Data: xRQTxnVerificationResponseData{
					Accountid:                     134,
					Acctalias:                     "temi",
					Acctbearsfeeattransactiontime: 1,
					Acctbusinessname:              "Synergy Group",
					Acctcontactperson:             "Desola Ade",
					Acctcountry:                   "NG",
					Acctisliveapproved:            0,
					Acctparent:                    1,
					Acctvpcmerchant:               "N/A",
					Amount:                        8150,
					Amountsettledforthistransaction: 8150,
					Authmodel:                       "PIN",
					Authurl:                         "http://flw-pms-dev.eu-west-1.elasticbeanstalk.com/mockvbvpage?ref=FLW-MOCK-5980e4f35eb54158fc296a1f22b1ff88&code=00&message=Approved. Successful&receiptno=RN1504378687808",
					Chargecode:                      "00",
					Chargedamount:                   8150,
					Chargemessage:                   "Success-Pending-otp-validation",
					Chargetype:                      "normal",
					Created:                         "2017-09-02T18:58:07.000Z",
					Createdday:                      6,
					Createddayispublicholiday:       0,
					Createddayname:                  "SATURDAY",
					Createdhour:                     18,
					Createdminute:                   58,
					Createdmonth:                    8,
					Createdmonthname:                "SEPTEMBER",
					Createdpmam:                     "pm",
					Createdquarter:                  3,
					Createdweek:                     35,
					Createdyear:                     2017,
					Createdyearisleap:               false,
					Currency:                        "NGN",
					Custcreated:                     "2017-09-02T18:58:07.000Z",
					Custemail:                       "bakare.wilmot@thegiggroupng.com",
					Custemailprovider:               "COMPANY EMAIL",
					Custname:                        "Anonymous customer",
					Custnetworkprovider:             "N/A",
					Customerid:                      2777,
					Cycle:                           "one-time",
					Devicefingerprint:               "ee09e26e8a76cfa35e64688c95145c2f",
					Flwref:                          "FLW-MOCK-5980e4f35eb54158fc296a1f22b1ff88",
					Fraudstatus:                     "ok",
					IP:                              "197.210.25.161",
					Narration:                       "FLW-PBF CARD Transaction ",
					Orderref:                        "URF_1504378687803_68035",
					Paymentid:                       "2",
					Paymenttype:                     "card",
					Status:                          "successful",
					Txid:                            32458,
					Txref:                           "OH-AAED44",
					Vbvcode:                         "00",
					Vbvmessage:                      "successful",
				},
			},
			want1: []error{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tvc := &TxnVerificationChecklist{
				Amount:              tt.fields.Amount,
				Flwref:              tt.fields.Flwref,
				FlwRef:              tt.fields.Flwref,
				LastAttempt:         tt.fields.LastAttempt,
				OnlySuccessful:      tt.fields.OnlySuccessful,
				VerificationURL:     tt.fields.VerificationURL,
				TransactionCurrency: tt.fields.TransactionCurrency,
				Txref:               tt.fields.Txref,
			}

			handler.resp = []byte(tt.respBody)

			got, got1 := tvc.VerifyXRequeryTransaction()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TxnVerificationChecklist.VerifyXRequeryTransaction() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("TxnVerificationChecklist.VerifyXRequeryTransaction() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}

	// Test it sets a verification URL if empty
	tvc := &TxnVerificationChecklist{VerificationURL: ""}
	baseURL = server.URL
	tvc.VerifyXRequeryTransaction()

	if tvc.VerificationURL == "" {
		t.Errorf("expected TxnVerificationChecklist.VerifyTransaction().VerificationURL not to be mepty")
	}
}
