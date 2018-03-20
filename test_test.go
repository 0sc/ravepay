package rave

import (
	"log"
	"net/http"
)

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

var successfulCardVerifyPaymentResponse = `{
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
      "brand": "VISA  CREDIT"
    },
    
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

var successfulAccountVerifyPaymentResponse = `{
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
      "deletedAt": null
    },
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

var xRQSuccessfulVerificationResponse = `{"status":"success","message":"Tx Fetched","data":{"txid":32458,"txref":"OH-AAED44","flwref":"FLW-MOCK-5980e4f35eb54158fc296a1f22b1ff88","devicefingerprint":"ee09e26e8a76cfa35e64688c95145c2f","cycle":"one-time","amount":8150,"currency":"NGN","chargedamount":8150,"appfee":0,"merchantfee":0,"merchantbearsfee":0,"chargecode":"00","chargemessage":"Success-Pending-otp-validation","authmodel":"PIN","ip":"197.210.25.161","narration":"FLW-PBF CARD Transaction ","status":"successful","vbvcode":"00","vbvmessage":"successful","authurl":"http://flw-pms-dev.eu-west-1.elasticbeanstalk.com/mockvbvpage?ref=FLW-MOCK-5980e4f35eb54158fc296a1f22b1ff88&code=00&message=Approved. Successful&receiptno=RN1504378687808","acctcode":null,"acctmessage":null,"paymenttype":"card","paymentid":"2","fraudstatus":"ok","chargetype":"normal","createdday":6,"createddayname":"SATURDAY","createdweek":35,"createdmonth":8,"createdmonthname":"SEPTEMBER","createdquarter":3,"createdyear":2017,"createdyearisleap":false,"createddayispublicholiday":0,"createdhour":18,"createdminute":58,"createdpmam":"pm","created":"2017-09-02T18:58:07.000Z","customerid":2777,"custphone":null,"custnetworkprovider":"N/A","custname":"Anonymous customer","custemail":"bakare.wilmot@thegiggroupng.com","custemailprovider":"COMPANY EMAIL","custcreated":"2017-09-02T18:58:07.000Z","accountid":134,"acctbusinessname":"Synergy Group","acctcontactperson":"Desola Ade","acctcountry":"NG","acctbearsfeeattransactiontime":1,"acctparent":1,"acctvpcmerchant":"N/A","acctalias":"temi","acctisliveapproved":0,"orderref":"URF_1504378687803_68035","paymentplan":null,"paymentpage":null,"raveref":null,"amountsettledforthistransaction":8150}}`

var successfulPreAuthPaymentCaptureResponse = `{"status":"success","message":"Capture complete","data":{"id":194211,"txRef":"MC-1508990174050","orderRef":null,"flwRef":"FLW-MOCK-839c1abc23b6a4bbb9da807d54c5bbda","redirectUrl":"http://127.0.0","device_fingerprint":"69e6b7f0b72037aa8428b70fbe03986c","settlement_token":null,"cycle":"one-time","amount":20,"charged_amount":20.25,"appfee":0.25,"merchantfee":0,"merchantbearsfee":0,"chargeResponseCode":"00","chargeResponseMessage":"Approved","authModelUsed":"NOAUTH","currency":"NGN","IP":"::ffff:10.37.225.74","narration":"FLW-PBF CARD Transaction ","status":"successful","vbvrespmessage":"Approved","authurl":"N/A","vbvrespcode":"00","acctvalrespmsg":null,"acctvalrespcode":null,"paymentType":"card","paymentPlan":null,"paymentPage":null,"paymentId":"47433","fraud_status":"ok","charge_type":"preauth","is_live":0,"createdAt":"2017-10-26T03:56:42.000Z","updatedAt":"2017-10-26T04:20:21.000Z","deletedAt":null,"customerId":88067,"AccountId":48,"customer":{"id":88067,"phone":"08056552980","fullName":"temi desola","customertoken":null,"email":"tester@flutter.co","createdAt":"2017-10-26T03:39:45.000Z","updatedAt":"2017-10-26T03:39:45.000Z","deletedAt":null,"AccountId":48}}}`

var successfulPreAuthPaymentRefundResponse = `{"status":"success","message":"Refund or void complete","data":{"data":{"responsecode":"00","redirecturl":null,"avsresponsemessage":null,"avsresponsecode":null,"authorizeId":"","responsemessage":"Approved","otptransactionidentifier":null,"transactionreference":"FLW59293657","responsehtml":null,"responsetoken":null},"status":"success"}}`

var failedPreAuthPaymentRefundResponse = `{"status":"error","message":"No transaction found","data":{"code":"NO_TX","message":"No transaction found"}}`

var getFeeResponse = `{"status":"success","message":"Charged fee","data":{"charge_amount":"1052.50","fee":52.5,"merchantfee":"0","ravefee":"52.5"}}`
