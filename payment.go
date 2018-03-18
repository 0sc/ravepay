package rave

import "fmt"

type Payment struct {
	// PBFPubKey         string `json:"PBFPubKey"`
	Amount            int    `json:"amount"`
	Country           string `json:"country"`
	Currency          string `json:"currency"`
	CustomDescription string `json:"custom_description"`
	CustomLogo        string `json:"custom_logo"`
	CustomTitle       string `json:"custom_title"`
	CustomerEmail     string `json:"customer_email"`
	CustomerFirstname string `json:"customer_firstname"`
	CustomerLastname  string `json:"customer_lastname"`
	CustomerPhone     string `json:"customer_phone"`
	PaymentMethod     string `json:"payment_method"`
	TxRef             string `json:"txref"`
}

const (
	PaymentVerificationRequestFailed               = ""
	PaymentVerificationStatusMismatch              = ""
	PaymentVerificationChargeResponseValueError    = ""
	PaymentVerificationTransactionRefMismatch      = ""
	PaymentVerificationTransactionCurrencyMismatch = ""
	PaymentVerificationAmountError                 = ""
)

// PaymentVerification encapsulates the payment verification response
// gotten from http://flw-pms-dev.eu-west-1.elasticbeanstalk.com/flwv3-pug/getpaidx/api/verify
type PaymentVerification struct {
	Data    paymentVerificationData `json:"data"`
	Message string                  `json:"message"`
	Status  string                  `json:"status"`
}

type paymentVerificationData struct {
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

// PaymentVerificationChecklist encapsulates the payment verification process
// It includes details of the payment to verify
// and the verification response
type PaymentVerificationChecklist struct {
	// Amount is the payment amount to be verified
	Amount int `json:"-"`
	// FlwRef is the unique transaction reference
	FlwRef    string `json:"flw_ref"`
	Normalize string `json:"normalize"`
	//secret key from pay button generated on rave dashboard
	SECKEY string `json:"SECKEY"`
	// PaymentVerification is the verification response object
	PaymentVerification    *PaymentVerification `json:"-"`
	PaymentVerificationURL string               `json:"-"`
	// TransactionCurrency is the payment currency to be verified
	TransactionCurrency string `json:"-"`
	// Done tracks whether verification has been attempted. It starts out false for new objects and changes to true after #verify is called on the object
	Done bool `json:"-"`
	// Errors is an array of any errors which occured with the verification
	Errors []string `json:"-"`
}

// NewPaymentVerificationChecklist returns a new paymentVerificationChecklist object with the given params
func NewPaymentVerificationChecklist(amount int, flwRef, transactionCurrency, secKey string) *PaymentVerificationChecklist {
	return &PaymentVerificationChecklist{
		Amount:                 amount,
		Done:                   false,
		FlwRef:                 flwRef,
		TransactionCurrency:    transactionCurrency,
		SECKEY:                 secKey,
		PaymentVerificationURL: PaymentVerificationURL,
	}
}

// Verify kicks off the verification process with the given payment checklist
// It returns true if verification is sucessful
// It returns false if verification is unsuccessful
// and marks the verification as Done
func (pvc *PaymentVerificationChecklist) Verify() bool {
	// TODO: Validate checklist???
	// Make request to endpoint
	pvc.PaymentVerification = &PaymentVerification{}
	err := sendRequestAndParseResponse("POST", pvc.PaymentVerificationURL, pvc, pvc.PaymentVerification)
	pvc.Done = true

	if err != nil {
		pvc.Errors = append(pvc.Errors, PaymentVerificationRequestFailed)
		return false
	}

	pvc.verifyTransactionReference()
	pvc.verifyTransactionStatus()
	pvc.verifyChargeResponseValue()
	pvc.verifyTransactionCurrency()
	pvc.verifyPaymentAmount()

	return len(pvc.Errors) == 0
}

func (pvc *PaymentVerificationChecklist) verifyTransactionReference() {
	if flwRef := pvc.PaymentVerification.Data.FlwRef; flwRef != pvc.FlwRef {
		msg := fmt.Sprintf("%s but got %s", PaymentVerificationTransactionRefMismatch, flwRef)
		pvc.Errors = append(pvc.Errors, msg)
	}
}

func (pvc *PaymentVerificationChecklist) verifyTransactionStatus() {
	if status := pvc.PaymentVerification.Status; status != "success" {
		msg := fmt.Sprintf("%s but got %s", PaymentVerificationStatusMismatch, status)
		pvc.Errors = append(pvc.Errors, msg)
	}
}

func (pvc *PaymentVerificationChecklist) verifyChargeResponseValue() {
	if respVal := pvc.PaymentVerification.Data.FlwMeta.ChargeResponse; respVal != "00" && respVal != "0" {
		msg := fmt.Sprintf("%s but got %s", PaymentVerificationChargeResponseValueError, respVal)
		pvc.Errors = append(pvc.Errors, msg)
	}
}

func (pvc *PaymentVerificationChecklist) verifyTransactionCurrency() {
	if currency := pvc.PaymentVerification.Data.TransactionCurrency; currency != pvc.TransactionCurrency {
		msg := fmt.Sprintf("%s but got %s", PaymentVerificationTransactionCurrencyMismatch, currency)
		pvc.Errors = append(pvc.Errors, msg)
	}
}

func (pvc *PaymentVerificationChecklist) verifyPaymentAmount() {
	if amt := pvc.PaymentVerification.Data.Amount; amt < pvc.Amount {
		msg := fmt.Sprintf("%s but got %d", PaymentVerificationAmountError, amt)
		pvc.Errors = append(pvc.Errors, msg)
	}
}
