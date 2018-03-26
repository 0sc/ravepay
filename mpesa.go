package rave

import (
	"encoding/json"
	"log"
)

// Mpesa is a type that encapsulates rave's mpesa description
// It has all mpesa attributes necessary for rave api mpesa references
// It also implements the chargable interface required for making charge requests and validating them
type Mpesa struct {
	ChargeMpesaURL string `json:"-"`
	Currency       string `json:"currency"`
	Country        string `json:"country"`
	LastName       string `json:"lastname,omitempty"`
	FirstName      string `json:"firstname,omitempty"`
	IsMpesa        string `json:"is_mpesa,omitempty"`
}

// MpesaPaymentInfo is the information necessary for completing mpesa payment
type MpesaPaymentInfo struct {
	AccountNumber  string
	Amount         int
	BusinessNumber string
}

// ChargeURL is an implemenation of the Chargeable interface
// it returns the url to be used for charging the given card
func (m *Mpesa) ChargeURL() string {
	if m.ChargeMpesaURL == "" {
		m.ChargeMpesaURL = buildURL(defaultChargeURL)
	}
	return m.ChargeMpesaURL
}

// ValidateChargeURL is an implemenation of the Chargeable interface
// it returns the url to be used for charging the given card
func (m *Mpesa) ValidateChargeURL() string {
	return ""
}

// BuildChargeRequestPayload is an implemenation of the Chargeable interface
// it returns the byte representation of the charge request client
// at the ChargeRequest level, chargeables are merely interface objects
// so trying to compose a struct with an interface object results in go adding the interface name key to the result bytes
// see https://play.golang.com/p/MFfbuPLrjo6
// so here we upend it so the individual concrete types do the marshalling
func (m *Mpesa) BuildChargeRequestPayload(cReq *ChargeRequest) []byte {
	cReq.PaymentType = "mpesa"
	payload := struct {
		*Mpesa
		*ChargeRequest
	}{m, cReq}
	b, err := json.Marshal(payload)
	if err != nil {
		log.Println("couldn't marshal payload: ", err)
	}
	return b
}

// MpesaPaymentInstruction parses the given charge response
// and returns the payment info: business number and account number for completing the payment
func MpesaPaymentInstruction(cr *ChargeResponse) *MpesaPaymentInfo {
	return &MpesaPaymentInfo{
		Amount:         cr.Data.Amount,
		AccountNumber:  cr.Data.OrderRef,
		BusinessNumber: cr.Data.BusinessNumber,
	}
}
