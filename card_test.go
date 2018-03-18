package rave

import (
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestChargeCardRequest_Charge(t *testing.T) {
	handler := &testServer{}
	server := httptest.NewServer(handler)
	defer server.Close()

	type fields struct {
		PBFPubKey         string
		Amount            float64
		DeviceFingerprint string
		Email             string
		IP                string
		TxRef             string
		PhoneNumber       string
		SuggestedAuth     string
	}

	type args struct {
		card *Card
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		serverResp string
		want       *ChargeCardResponse
		wantErr    bool
	}{
		{
			name: "returns the charge request response",
			fields: fields{
				PBFPubKey:         "FLWPUBK-e634d14d9ded04eaf05d5b63a0a06d2f-X",
				Amount:            300,
				Email:             "tester@flutter.co",
				IP:                "103.238.105.185",
				TxRef:             "'MXX-ASC-4578",
				DeviceFingerprint: "69e6b7f0sb72037aa8428b70fbe03986c",
				SuggestedAuth:     "pin",
			},
			args: args{
				card: &Card{
					CardNo:      "5438898014560229",
					Currency:    "NGN",
					Country:     "NG",
					Cvv:         789,
					Expirymonth: "09",
					Expiryyear:  "19",
					Pin:         "3310",
				},
			},
			want: &ChargeCardResponse{
				Message: "V-COMP",
				Status:  "success",
				Data: ChargeCardResponseData{
					AccountID:     134,
					IP:            "::ffff:127.0.0.1",
					Amount:        300,
					AuthModelUsed: "PIN",
					Authurl:       "http://flw-pms-dev.eu-west-1.elasticbeanstalk.com/mockvbvpage?ref=FLW-MOCK-0cd9a725cf2ad31303299840f5a0896a&code=00&message=Approved. Successful&receiptno=RN1521335722125",
					Customercandosubsequentnoauth: true,
					ChargeResponseCode:            "02",
					ChargeResponseMessage:         "Success-Pending-otp-validation",
					ChargeType:                    "normal",
					ChargedAmount:                 300,
					CreatedAt:                     "2018-03-18T01:15:22.000Z",
					Currency:                      "NGN",
					CustomerID:                    20322,
					Cycle:                         "one-time",
					DeviceFingerprint:             "69e6b7f0sb72037aa8428b70fbe03986c",
					FlwRef:                        "FLW-MOCK-0cd9a725cf2ad31303299840f5a0896a",
					FraudStatus:                   "ok",
					ID:                            113247,
					Merchantbearsfee:              1,
					Narration:                     "FLW-PBF CARD Transaction ",
					OrderRef:                      "URF_1521335722082_4111335",
					PaymentID:                     "861",
					PaymentType:                   "card",
					RaveRef:                       "RV31521335720694ADC5E86652",
					RedirectURL:                   "N/A",
					Status:                        "success-pending-validation",
					TxRef:                         "MXX-ASC-4578",
					UpdatedAt:                     "2018-03-18T01:15:24.000Z",
					Vbvrespcode:                   "00",
					Vbvrespmessage:                "Approved. Successful",
				},
			},
			serverResp: successfulCardChargeResponse,
			wantErr:    false,
		},
		{
			name: "returns response if request has auth suggestion",
			fields: fields{
				PBFPubKey:         "FLWPUBK-e634d14d9ded04eaf05d5b63a0a06d2f-X",
				Amount:            300,
				Email:             "tester@flutter.co",
				IP:                "103.238.105.185",
				TxRef:             "'MXX-ASC-4578",
				DeviceFingerprint: "69e6b7f0sb72037aa8428b70fbe03986c",
			},
			args: args{
				card: &Card{
					CardNo:      "5438898014560229",
					Currency:    "NGN",
					Country:     "NG",
					Cvv:         789,
					Expirymonth: "09",
					Expiryyear:  "19",
				},
			},
			want: &ChargeCardResponse{
				Message: "AUTH_SUGGESTION",
				Status:  "success",
				Data:    ChargeCardResponseData{SuggestedAuth: "PIN"},
			},
			serverResp: suggestedAuthCardChargeResponse,
			wantErr:    false,
		},
		{
			name: "returns response if request fails",
			args: args{
				card: &Card{CardNo: "5"},
			},
			want: &ChargeCardResponse{
				Message: "BIN not Found",
				Status:  "error",
				Data: ChargeCardResponseData{
					Code:    "BIN_ERR",
					Message: "BIN not Found",
				},
			},
			serverResp: failedCardChargeResponse,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ccr := &ChargeCardRequest{
				PBFPubKey:         tt.fields.PBFPubKey,
				Amount:            tt.fields.Amount,
				ChargeURL:         server.URL,
				DeviceFingerprint: tt.fields.DeviceFingerprint,
				Email:             tt.fields.Email,
				IP:                tt.fields.IP,
				TxRef:             tt.fields.TxRef,
				SuggestedAuth:     tt.fields.SuggestedAuth,
			}
			handler.resp = []byte(tt.serverResp)

			got, err := ccr.Charge(tt.args.card)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChargeCardRequest.Charge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChargeCardRequest.Charge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChargeCardResponse_Validate(t *testing.T) {
	handler := &testServer{}
	server := httptest.NewServer(handler)
	defer server.Close()

	type fields struct {
		Data      ChargeCardResponseData
		PBFPubKey string
	}
	type args struct {
		otp string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		serverResp string
		want       *ChargeCardValidationResponse
		wantErr    bool
	}{
		{
			name: "returns the request response",
			fields: fields{
				PBFPubKey: "FLWPUBK-e634d14d9ded04eaf05d5b63a0a06d2f-X",
				Data: ChargeCardResponseData{
					FlwRef: "FLW-MOCK-902e0201437d6b20e5b7e2b3ec140967",
				},
			},
			args:       args{otp: "12345"},
			serverResp: successfulValidateChargeCardResponse,
			want: &ChargeCardValidationResponse{
				Status:  "success",
				Message: "Charge Complete",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ccr := &ChargeCardResponse{
				Data:              tt.fields.Data,
				ValidateChargeURL: server.URL,
				PBFPubKey:         tt.fields.PBFPubKey,
			}

			handler.resp = []byte(tt.serverResp)

			got, err := ccr.Validate(tt.args.otp)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChargeCardResponse.Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.Message != tt.want.Message || got.Status != tt.want.Status {
				t.Errorf("ChargeCardResponse.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

var successfulCardChargeResponse = `{"status":"success","message":"V-COMP","data":{"id":113247,"txRef":"MXX-ASC-4578","orderRef":"URF_1521335722082_4111335","flwRef":"FLW-MOCK-0cd9a725cf2ad31303299840f5a0896a","redirectUrl":"N/A","device_fingerprint":"69e6b7f0sb72037aa8428b70fbe03986c","settlement_token":null,"cycle":"one-time","amount":300,"charged_amount":300,"appfee":0,"merchantfee":0,"merchantbearsfee":1,"chargeResponseCode":"02","raveRef":"RV31521335720694ADC5E86652","chargeResponseMessage":"Success-Pending-otp-validation","authModelUsed":"PIN","currency":"NGN","IP":"::ffff:127.0.0.1","narration":"FLW-PBF CARD Transaction ","status":"success-pending-validation","vbvrespmessage":"Approved. Successful","authurl":"http://flw-pms-dev.eu-west-1.elasticbeanstalk.com/mockvbvpage?ref=FLW-MOCK-0cd9a725cf2ad31303299840f5a0896a&code=00&message=Approved. Successful&receiptno=RN1521335722125","vbvrespcode":"00","acctvalrespmsg":null,"acctvalrespcode":null,"paymentType":"card","paymentPlan":null,"paymentPage":null,"paymentId":"861","fraud_status":"ok","charge_type":"normal","is_live":0,"createdAt":"2018-03-18T01:15:22.000Z","updatedAt":"2018-03-18T01:15:24.000Z","deletedAt":null,"customerId":20322,"AccountId":134,"customercandosubsequentnoauth":true}}`

var suggestedAuthCardChargeResponse = `{"status":"success","message":"AUTH_SUGGESTION","data":{"suggested_auth":"PIN"}}`

var failedCardChargeResponse = `{"status":"error","message":"BIN not Found","data":{"code":"BIN_ERR","message":"BIN not Found"}}`

var successfulValidateChargeCardResponse = `{"status":"success","message":"Charge Complete","data":{"data":{"responsecode":"00","responsemessage":"successful"},"tx":{"id":113257,"txRef":"'MXX-ASC-4578","orderRef":"URF_1521341698263_4239435","flwRef":"FLW-MOCK-902e0201437d6b20e5b7e2b3ec140967","redirectUrl":"N/A","device_fingerprint":"69e6b7f0sb72037aa8428b70fbe03986c","settlement_token":null,"cycle":"one-time","amount":300,"charged_amount":300,"appfee":0,"merchantfee":0,"merchantbearsfee":1,"chargeResponseCode":"00","raveRef":"RV315213416971663EF7A02220","chargeResponseMessage":"Success-Pending-otp-validation","authModelUsed":"PIN","currency":"NGN","IP":"::ffff:127.0.0.1","narration":"FLW-PBF CARD Transaction ","status":"successful","vbvrespmessage":"successful","authurl":"http://flw-pms-dev.eu-west-1.elasticbeanstalk.com/mockvbvpage?ref=FLW-MOCK-902e0201437d6b20e5b7e2b3ec140967&code=00&message=Approved. Successful&receiptno=RN1521341698308","vbvrespcode":"00","acctvalrespmsg":null,"acctvalrespcode":null,"paymentType":"card","paymentPlan":null,"paymentPage":null,"paymentId":"861","fraud_status":"ok","charge_type":"normal","is_live":0,"createdAt":"2018-03-18T02:54:58.000Z","updatedAt":"2018-03-18T02:55:10.000Z","deletedAt":null,"customerId":20331,"AccountId":134,"customer":{"id":20331,"phone":null,"fullName":"Anonymous customer","customertoken":null,"email":"tester@flutter.co","createdAt":"2018-03-18T02:54:57.000Z","updatedAt":"2018-03-18T02:54:57.000Z","deletedAt":null,"AccountId":134},"chargeToken":{"user_token":"9b9c3","embed_token":"flw-t0-349218908feb91c1dda7ca991a4a4b3a-m03k"}}}}`
