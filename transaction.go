package ravepay

import "fmt"

// TxnVerificationResponse is a type of rave response for transaction verification request
// it implements the verifiable interface to allow rave's recommended followup verification
type TxnVerificationResponse struct {
	Data    txnVerificationResponseData `json:"data"`
	Message string                      `json:"message"`
	Status  string                      `json:"status"`
}

// XRQTxnVerificationResponse is a type of rave response for xrquery transaction verification request
// it implements the verifiable interface to allow rave's recommended followup verification
type XRQTxnVerificationResponse struct {
	Data    xRQTxnVerificationResponseData `json:"data"`
	Message string                         `json:"message"`
	Status  string                         `json:"status"`
}

type txnVerificationResponseData struct {
	Account Account `json:"account"`

	AddonID int `json:"addon_id"`
	Amount  int `json:"amount"`
	Appfee  int `json:"appfee"`

	Card Card `json:"card"`

	ChargeType        string      `json:"charge_type"`
	ChargebackStatus  interface{} `json:"chargeback_status"`
	ChargedAmount     int         `json:"charged_amount"`
	CreatedAt         string      `json:"createdAt"`
	Code              string      `json:"code"`
	Customer          Customer    `json:"customer"`
	Cycle             string      `json:"cycle"`
	DeletedAt         interface{} `json:"deletedAt"`
	DeviceFingerprint string      `json:"device_fingerprint"`
	FlwMeta           FlwMeta     `json:"flwMeta"`
	FlwRef            string      `json:"flw_ref"`
	FraudStatus       string      `json:"fraud_status"`
	ID                int         `json:"id"`
	IP                string      `json:"ip"`
	IsLive            int         `json:"is_live"`
	MarkupFee         interface{} `json:"markupFee"`
	Message           string      `json:"message"`
	MerchantID        int         `json:"merchant_id"`
	Merchantbearsfee  int         `json:"merchantbearsfee"`
	Merchantfee       int         `json:"merchantfee"`
	Meta              []struct {
		CreatedAt            string      `json:"createdAt"`
		DeletedAt            interface{} `json:"deletedAt"`
		GetpaidTransactionID int         `json:"getpaidTransactionId"`
		ID                   int         `json:"id"`
		Metaname             string      `json:"metaname"`
		Metavalue            string      `json:"metavalue"`
		UpdatedAt            string      `json:"updatedAt"`
	} `json:"meta"`
	Narration            string      `json:"narration"`
	OrderRef             string      `json:"order_ref"`
	PaymentEntity        string      `json:"payment_entity"`
	PaymentID            string      `json:"payment_id"`
	RaveRef              string      `json:"rave_ref"`
	SettlementToken      interface{} `json:"settlement_token"`
	Status               string      `json:"status"`
	SystemType           interface{} `json:"system_type"`
	TransactionCurrency  string      `json:"transaction_currency"`
	TransactionProcessor string      `json:"transaction_processor"`
	TransactionType      string      `json:"transaction_type"`
	TxRef                string      `json:"tx_ref"`
	UpdatedAt            string      `json:"updatedAt"`
}

type xRQTxnVerificationResponseData struct {
	Accountid                       int         `json:"accountid"`
	Acctalias                       string      `json:"acctalias"`
	Acctbearsfeeattransactiontime   int         `json:"acctbearsfeeattransactiontime"`
	Acctbusinessname                string      `json:"acctbusinessname"`
	Acctcode                        interface{} `json:"acctcode"`
	Acctcontactperson               string      `json:"acctcontactperson"`
	Acctcountry                     string      `json:"acctcountry"`
	Acctisliveapproved              int         `json:"acctisliveapproved"`
	Acctmessage                     interface{} `json:"acctmessage"`
	Acctparent                      int         `json:"acctparent"`
	Acctvpcmerchant                 string      `json:"acctvpcmerchant"`
	Amount                          int         `json:"amount"`
	Amountsettledforthistransaction int         `json:"amountsettledforthistransaction"`
	Appfee                          int         `json:"appfee"`
	Authmodel                       string      `json:"authmodel"`
	Authurl                         string      `json:"authurl"`
	Chargecode                      string      `json:"chargecode"`
	Chargedamount                   int         `json:"chargedamount"`
	Chargemessage                   string      `json:"chargemessage"`
	Chargetype                      string      `json:"chargetype"`
	Created                         string      `json:"created"`
	Createdday                      int         `json:"createdday"`
	Createddayispublicholiday       int         `json:"createddayispublicholiday"`
	Createddayname                  string      `json:"createddayname"`
	Createdhour                     int         `json:"createdhour"`
	Createdminute                   int         `json:"createdminute"`
	Createdmonth                    int         `json:"createdmonth"`
	Createdmonthname                string      `json:"createdmonthname"`
	Createdpmam                     string      `json:"createdpmam"`
	Createdquarter                  int         `json:"createdquarter"`
	Createdweek                     int         `json:"createdweek"`
	Createdyear                     int         `json:"createdyear"`
	Createdyearisleap               bool        `json:"createdyearisleap"`
	Currency                        string      `json:"currency"`
	Custcreated                     string      `json:"custcreated"`
	Custemail                       string      `json:"custemail"`
	Custemailprovider               string      `json:"custemailprovider"`
	Custname                        string      `json:"custname"`
	Custnetworkprovider             string      `json:"custnetworkprovider"`
	Customerid                      int         `json:"customerid"`
	Custphone                       interface{} `json:"custphone"`
	Cycle                           string      `json:"cycle"`
	Devicefingerprint               string      `json:"devicefingerprint"`
	Flwref                          string      `json:"flwref"`
	Fraudstatus                     string      `json:"fraudstatus"`
	IP                              string      `json:"ip"`
	Merchantbearsfee                int         `json:"merchantbearsfee"`
	Merchantfee                     int         `json:"merchantfee"`
	Narration                       string      `json:"narration"`
	Orderref                        string      `json:"orderref"`
	Paymentid                       string      `json:"paymentid"`
	Paymentpage                     interface{} `json:"paymentpage"`
	Paymentplan                     interface{} `json:"paymentplan"`
	Paymenttype                     string      `json:"paymenttype"`
	Raveref                         interface{} `json:"raveref"`
	Status                          string      `json:"status"`
	Txid                            int         `json:"txid"`
	Txref                           string      `json:"txref"`
	Vbvcode                         string      `json:"vbvcode"`
	Vbvmessage                      string      `json:"vbvmessage"`
}

// VerifyStatus verifies that response status is success
// returns error otherwise
func (resp *TxnVerificationResponse) VerifyStatus() error {
	if resp.Status != "success" {
		return fmt.Errorf("StatusVerificationFailed: expected success but got %s", resp.Status)
	}
	return nil
}

// VerifyStatus verifies that response status is success
// returns error otherwise
func (resp *XRQTxnVerificationResponse) VerifyStatus() error {
	if resp.Status != "success" {
		return fmt.Errorf("StatusVerificationFailed: expected success but got %s", resp.Status)
	}
	return nil
}

// VerifyCurrency verifies that the currency in the txn verification matches the given currency
// returns error otherwise
func (resp *TxnVerificationResponse) VerifyCurrency(currency string) error {
	if got := resp.Data.TransactionCurrency; currency != got {
		return fmt.Errorf("CurrencyVerificationFailed: expected %s but got %s", currency, got)
	}
	return nil
}

// VerifyCurrency verifies that the currency in the txn verification matches the given currency
// returns error otherwise
func (resp *XRQTxnVerificationResponse) VerifyCurrency(currency string) error {
	if got := resp.Data.Currency; currency != got {
		return fmt.Errorf("CurrencyVerificationFailed: expected %s but got %s", currency, got)
	}
	return nil
}

// VerifyAmount verifies that the given amount matches that in the transaction
// returns error otherwise
func (resp *TxnVerificationResponse) VerifyAmount(amt int) error {
	if got := resp.Data.Amount; amt > got {
		return fmt.Errorf("AmountVerificationFailed: expected %d but got %d", amt, got)
	}
	return nil
}

// VerifyAmount verifies that the given amount matches that in the transaction
// returns error otherwise
func (resp *XRQTxnVerificationResponse) VerifyAmount(amt int) error {
	if got := resp.Data.Amount; amt > got {
		return fmt.Errorf("AmountVerificationFailed: expected %d but got %d", amt, got)
	}
	return nil
}

// VerifyChargeResponseValue verifies that the charge response value for the transaction is either '0' or '00'
func (resp *TxnVerificationResponse) VerifyChargeResponseValue() error {
	if respVal := resp.Data.FlwMeta.ChargeResponse; respVal != "00" && respVal != "0" {
		return fmt.Errorf("ChargeResponseVerificationFailed: expected 00 or 0 but got %s", respVal)
	}
	return nil
}

// VerifyChargeResponseValue verifies that the charge response value for the transaction is either '0' or '00'
func (resp *XRQTxnVerificationResponse) VerifyChargeResponseValue() error {
	if respVal := resp.Data.Chargecode; respVal != "00" && respVal != "0" {
		return fmt.Errorf("ChargeResponseVerificationFailed: expected 00 or 0 but got %s", respVal)
	}
	return nil
}

// VerifyReference verifies that the flw ref in the transaction matches the given ref
// returns error otherwise
func (resp *TxnVerificationResponse) VerifyReference(ref string) error {
	if got := resp.Data.FlwRef; ref != got {
		return fmt.Errorf("FlwRefVerificationFailed: expected %s but got %s", ref, got)
	}
	return nil
}

// VerifyReference verifies that the flw ref in the transaction matches the given ref
// returns error otherwise
func (resp *XRQTxnVerificationResponse) VerifyReference(ref string) error {
	// TODO: Update to compare txRef instead??
	if got := resp.Data.Flwref; ref != got {
		return fmt.Errorf("FlwRefVerificationFailed: expected %s but got %s", ref, got)
	}
	return nil
}
