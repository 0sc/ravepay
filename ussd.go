package rave

import (
	"encoding/json"
	"log"
)

// USSD is a type that encapsulates rave's ussd description
// It has all ussd attributes necessary for rave api ussd referencess
// It also implements the chargabel interface required for making charge requests
type USSD struct {
	ChargeRequestURL      string `json:"-"`
	ValidateCardChargeURL string `json:"-"`
	Country               string `json:"country"`
	Currency              string `json:"currency"`
	FirstName             string `json:"firstname,omitempty"`
	LastName              string `json:"lastname,omitempty"`
}

// USSDPaymentInfo is the information necessary for completing mpesa payment
type USSDPaymentInfo struct {
	FlwRef string
	Amount int
}

// ChargeURL is an implemenation of the Chargeable interface
// it returns the url to be used for charging the given card
func (c *USSD) ChargeURL() string {
	if c.ChargeRequestURL == "" {
		c.ChargeRequestURL = ChargeUSSDURL
	}
	return c.ChargeRequestURL
}

// ValidateChargeURL is an implemenation of the Chargeable interface
// it returns the url to be used for charging the given card
func (c *USSD) ValidateChargeURL() string {
	return ""
}

// BuildChargeRequestPayload is an implemenation of the Chargeable interface
// it returns the byte representation of the charge request client
// at the ChargeRequest level, chargeables are merely interface objects
// so trying to compose a struct with an interface object results in go adding the interface name key to the result bytes
// see https://play.golang.com/p/MFfbuPLrjo6
// so here we upend it so the individual concrete types do the marshalling
func (c *USSD) BuildChargeRequestPayload(cReq *ChargeRequest) []byte {
	cReq.PaymentType = "ussd"
	payload := struct {
		*USSD
		*ChargeRequest
	}{c, cReq}
	b, err := json.Marshal(payload)
	if err != nil {
		log.Println("couldn't marshal payload: ", err)
	}
	return b
}

// USSDPaymentInstruction parses the given charge response
// and returns the payment info for completing the payment
func USSDPaymentInstruction(cr *ChargeResponse) *USSDPaymentInfo {
	return &USSDPaymentInfo{
		Amount: cr.Data.Amount,
		FlwRef: cr.Data.FlwRef,
	}
}
