package rave

import (
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestChargeRequest_Charge(t *testing.T) {
	handler := &testServer{}
	server := httptest.NewServer(handler)
	defer server.Close()

	type fields struct {
		PBFPubKey         string
		Amount            float64
		ChargeType        string
		DeviceFingerprint string
		Email             string
		IP                string
		TxRef             string
		PaymentType       string
		PhoneNumber       string
		RedirectURL       string
		SuggestedAuth     string
	}
	type args struct {
		chargeable Chargeable
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		serverResp string
		want       *ChargeResponse
		wantErr    bool
	}{
		{
			name: "returns the charge charge card request response",
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
				chargeable: &Card{
					CardNo:                "5438898014560229",
					ChargeCardURL:         server.URL,
					ValidateCardChargeURL: server.URL,
					Currency:              "NGN",
					Country:               "NG",
					Cvv:                   789,
					Expirymonth:           "09",
					Expiryyear:            "19",
					Pin:                   "3310",
				},
			},
			want: &ChargeResponse{
				Message: "V-COMP",
				Status:  "success",
				Data: chargeResponseData{
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
				ValidateChargeURL: server.URL,
			},
			serverResp: successfulCardChargeResponse,
			wantErr:    false,
		},
		{
			name: "returns response if charge card request has auth suggestion",
			fields: fields{
				PBFPubKey:         "FLWPUBK-e634d14d9ded04eaf05d5b63a0a06d2f-X",
				Amount:            300,
				Email:             "tester@flutter.co",
				IP:                "103.238.105.185",
				TxRef:             "'MXX-ASC-4578",
				DeviceFingerprint: "69e6b7f0sb72037aa8428b70fbe03986c",
			},
			args: args{
				chargeable: &Card{
					CardNo:                "5438898014560229",
					ChargeCardURL:         server.URL,
					ValidateCardChargeURL: server.URL,
					Currency:              "NGN",
					Country:               "NG",
					Cvv:                   789,
					Expirymonth:           "09",
					Expiryyear:            "19",
				},
			},
			want: &ChargeResponse{
				Message:           "AUTH_SUGGESTION",
				Status:            "success",
				Data:              chargeResponseData{SuggestedAuth: "PIN"},
				ValidateChargeURL: server.URL,
			},
			serverResp: suggestedAuthCardChargeResponse,
			wantErr:    false,
		},
		{
			name: "returns response if charge card request fails",
			args: args{
				chargeable: &Card{
					CardNo:                "5",
					ChargeCardURL:         server.URL,
					ValidateCardChargeURL: server.URL,
				},
			},
			want: &ChargeResponse{
				Message: "BIN not Found",
				Status:  "error",
				Data: chargeResponseData{
					Code:    "BIN_ERR",
					Message: "BIN not Found",
				},
				ValidateChargeURL: server.URL,
			},
			serverResp: failedCardChargeResponse,
			wantErr:    false,
		},
		{
			name: "returns response if charge account request succeeds",
			args: args{
				chargeable: &Account{
					AccountBank:              "044",
					AccountNumber:            "0690000031",
					Country:                  "NG",
					ChargeAccountURL:         server.URL,
					ValidateAccountChargeURL: server.URL,
				},
			},
			want: &ChargeResponse{
				Message: "V-COMP",
				Status:  "success",
				Data: chargeResponseData{
					AccountID:     134,
					IP:            "::ffff:127.0.0.1",
					Amount:        300,
					AuthModelUsed: "AUTH",
					Authurl:       "NO-URL",
					Customercandosubsequentnoauth: false,
					ChargeResponseCode:            "02",
					ChargeResponseMessage:         "Pending OTP validation",
					ChargeType:                    "normal",
					ChargedAmount:                 300,
					CreatedAt:                     "2018-03-18T19:41:50.000Z",
					Currency:                      "NGN",
					CustomerID:                    20385,
					Customer: Customer{
						AccountID: 134,
						CreatedAt: "2018-03-18T19:41:50.000Z",
						Email:     "tester@flutter.co",
						FullName:  "Anonymous customer",
						ID:        20385,
						UpdatedAt: "2018-03-18T19:41:50.000Z",
					},
					Cycle:               "one-time",
					DeviceFingerprint:   "69e6b7f0sb72037aa8428b70fbe03986c",
					FlwRef:              "ACHG-1521402110867",
					FraudStatus:         "ok",
					ID:                  113725,
					Merchantbearsfee:    1,
					Narration:           "Synergy Group",
					OrderRef:            "URF_1521402110125_5862635",
					PaymentID:           "2",
					PaymentType:         "account",
					RaveRef:             "RV31521402109467753A4A7081",
					RedirectURL:         "http://127.0.0",
					Status:              "success-pending-validation",
					TxRef:               "MXX-ASC-4578",
					UpdatedAt:           "2018-03-18T19:41:52.000Z",
					ValidateInstruction: "Please dial *901*4*1# to get your OTP. Enter the OTP gotten in the field below",
					ValidateInstructions: validateInstructions{
						Instruction: "Please validate with the OTP sent to your mobile or email",
						Valparams:   []string{"OTP"},
					},
					Vbvrespcode:    "N/A",
					Vbvrespmessage: "N/A",
				},
				ValidateChargeURL: server.URL,
			},
			serverResp: successfulChargeAccountResponse,
			wantErr:    false,
		},
		{
			name: "returns response if charge account request fails",
			args: args{
				chargeable: &Account{
					AccountBank:              "044",
					AccountNumber:            "",
					Country:                  "NG",
					ChargeAccountURL:         server.URL,
					ValidateAccountChargeURL: server.URL,
				},
			},
			want: &ChargeResponse{
				Message:           "accountnumber is required",
				Status:            "error",
				Data:              chargeResponseData{},
				ValidateChargeURL: server.URL,
			},
			serverResp: failedAccountChargeResponse,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cr := &ChargeRequest{
				PBFPubKey:         tt.fields.PBFPubKey,
				Amount:            tt.fields.Amount,
				ChargeType:        tt.fields.ChargeType,
				DeviceFingerprint: tt.fields.DeviceFingerprint,
				Email:             tt.fields.Email,
				IP:                tt.fields.IP,
				TxRef:             tt.fields.TxRef,
				PaymentType:       tt.fields.PaymentType,
				PhoneNumber:       tt.fields.PhoneNumber,
				RedirectURL:       tt.fields.RedirectURL,
				SuggestedAuth:     tt.fields.SuggestedAuth,
			}
			handler.resp = []byte(tt.serverResp)

			got, err := cr.Charge(tt.args.chargeable)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChargeRequest.Charge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChargeRequest.Charge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChargeResponse_Validate(t *testing.T) {
	handler := &testServer{}
	server := httptest.NewServer(handler)
	defer server.Close()

	type fields struct {
		Data              chargeResponseData
		ValidateChargeURL string
		PBFPubKey         string
	}
	type args struct {
		otp string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		serverResp string
		want       *ChargeValidationResponse
		wantErr    bool
	}{
		{
			name: "returns the charge card request response",
			fields: fields{
				PBFPubKey: "FLWPUBK-e634d14d9ded04eaf05d5b63a0a06d2f-X",
				Data: chargeResponseData{
					FlwRef: "FLW-MOCK-902e0201437d6b20e5b7e2b3ec140967",
				},
				ValidateChargeURL: server.URL,
			},
			args:       args{otp: "12345"},
			serverResp: successfulValidateChargeCardResponse,
			want: &ChargeValidationResponse{
				Status:  "success",
				Message: "Charge Complete",
			},
			wantErr: false,
		},
		{
			name: "returns the charge account request response",
			fields: fields{
				PBFPubKey: "FLWPUBK-e634d14d9ded04eaf05d5b63a0a06d2f-X",
				Data: chargeResponseData{
					FlwRef: "ACHG-1521409773573",
				},
				ValidateChargeURL: server.URL,
			},
			args:       args{otp: "12345"},
			serverResp: successfulValidateChargeAccountResponse,
			want: &ChargeValidationResponse{
				Status:  "success",
				Message: "Charge Complete",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cr := &ChargeResponse{
				Data:              tt.fields.Data,
				ValidateChargeURL: tt.fields.ValidateChargeURL,
				PBFPubKey:         tt.fields.PBFPubKey,
			}
			handler.resp = []byte(tt.serverResp)

			got, err := cr.Validate(tt.args.otp)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChargeResponse.Validate() error = %v, wantErr %v", err, tt.wantErr)
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

var successfulChargeAccountResponse = `{"status":"success","message":"V-COMP","data":{"id":113725,"txRef":"MXX-ASC-4578","orderRef":"URF_1521402110125_5862635","flwRef":"ACHG-1521402110867","redirectUrl":"http://127.0.0","device_fingerprint":"69e6b7f0sb72037aa8428b70fbe03986c","settlement_token":null,"cycle":"one-time","amount":300,"charged_amount":300,"appfee":0,"merchantfee":0,"merchantbearsfee":1,"chargeResponseCode":"02","raveRef":"RV31521402109467753A4A7081","chargeResponseMessage":"Pending OTP validation","authModelUsed":"AUTH","currency":"NGN","IP":"::ffff:127.0.0.1","narration":"Synergy Group","status":"success-pending-validation","vbvrespmessage":"N/A","authurl":"NO-URL","vbvrespcode":"N/A","acctvalrespmsg":null,"acctvalrespcode":null,"paymentType":"account","paymentPlan":null,"paymentPage":null,"paymentId":"2","fraud_status":"ok","charge_type":"normal","is_live":0,"createdAt":"2018-03-18T19:41:50.000Z","updatedAt":"2018-03-18T19:41:52.000Z","deletedAt":null,"customerId":20385,"AccountId":134,"customer":{"id":20385,"phone":null,"fullName":"Anonymous customer","customertoken":null,"email":"tester@flutter.co","createdAt":"2018-03-18T19:41:50.000Z","updatedAt":"2018-03-18T19:41:50.000Z","deletedAt":null,"AccountId":134},"validateInstructions":{"valparams":["OTP"],"instruction":"Please validate with the OTP sent to your mobile or email"},"validateInstruction":"Please dial *901*4*1# to get your OTP. Enter the OTP gotten in the field below"}}`

var failedCardChargeResponse = `{"status":"error","message":"BIN not Found","data":{"code":"BIN_ERR","message":"BIN not Found"}}`

var failedAccountChargeResponse = `{"status":"error","message":"accountnumber is required","data":{}}`

var successfulValidateChargeCardResponse = `{"status":"success","message":"Charge Complete","data":{"data":{"responsecode":"00","responsemessage":"successful"},"tx":{"id":113257,"txRef":"'MXX-ASC-4578","orderRef":"URF_1521341698263_4239435","flwRef":"FLW-MOCK-902e0201437d6b20e5b7e2b3ec140967","redirectUrl":"N/A","device_fingerprint":"69e6b7f0sb72037aa8428b70fbe03986c","settlement_token":null,"cycle":"one-time","amount":300,"charged_amount":300,"appfee":0,"merchantfee":0,"merchantbearsfee":1,"chargeResponseCode":"00","raveRef":"RV315213416971663EF7A02220","chargeResponseMessage":"Success-Pending-otp-validation","authModelUsed":"PIN","currency":"NGN","IP":"::ffff:127.0.0.1","narration":"FLW-PBF CARD Transaction ","status":"successful","vbvrespmessage":"successful","authurl":"http://flw-pms-dev.eu-west-1.elasticbeanstalk.com/mockvbvpage?ref=FLW-MOCK-902e0201437d6b20e5b7e2b3ec140967&code=00&message=Approved. Successful&receiptno=RN1521341698308","vbvrespcode":"00","acctvalrespmsg":null,"acctvalrespcode":null,"paymentType":"card","paymentPlan":null,"paymentPage":null,"paymentId":"861","fraud_status":"ok","charge_type":"normal","is_live":0,"createdAt":"2018-03-18T02:54:58.000Z","updatedAt":"2018-03-18T02:55:10.000Z","deletedAt":null,"customerId":20331,"AccountId":134,"customer":{"id":20331,"phone":null,"fullName":"Anonymous customer","customertoken":null,"email":"tester@flutter.co","createdAt":"2018-03-18T02:54:57.000Z","updatedAt":"2018-03-18T02:54:57.000Z","deletedAt":null,"AccountId":134},"chargeToken":{"user_token":"9b9c3","embed_token":"flw-t0-349218908feb91c1dda7ca991a4a4b3a-m03k"}}}}`

var successfulValidateChargeAccountResponse = `{"status":"success","message":"Charge Complete","data":{"id":113767,"txRef":"'MXX-ASC-4578","orderRef":"URF_1521409772957_2716735","flwRef":"ACHG-1521409773573","redirectUrl":"http://127.0.0","device_fingerprint":"69e6b7f0sb72037aa8428b70fbe03986c","settlement_token":null,"cycle":"one-time","amount":300,"charged_amount":300,"appfee":0,"merchantfee":0,"merchantbearsfee":1,"chargeResponseCode":"00","raveRef":"RV315214097725060FE00CEDBB","chargeResponseMessage":"Pending OTP validation","authModelUsed":"AUTH","currency":"NGN","IP":"::ffff:127.0.0.1","narration":"Synergy Group","status":"successful","vbvrespmessage":"N/A","authurl":"NO-URL","vbvrespcode":"N/A","acctvalrespmsg":"Approved Or Completed Successfully","acctvalrespcode":"00","paymentType":"account","paymentPlan":null,"paymentPage":null,"paymentId":"2","fraud_status":"ok","charge_type":"normal","is_live":0,"createdAt":"2018-03-18T21:49:32.000Z","updatedAt":"2018-03-18T21:49:50.000Z","deletedAt":null,"customerId":20391,"AccountId":134,"customer":{"id":20391,"phone":null,"fullName":"Anonymous customer","customertoken":null,"email":"tester@flutter.co","createdAt":"2018-03-18T21:49:32.000Z","updatedAt":"2018-03-18T21:49:32.000Z","deletedAt":null,"AccountId":134}}}`
