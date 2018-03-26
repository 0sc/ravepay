package rave

import (
	"encoding/json"
	"log"
)

// MobileMoneyGH is a type that encapsulates rave's ghana mobile money description
// It has all mpesa attributes necessary for rave api ghana mobile money description references
// It also implements the chargable interface required for making charge requests and validating them
type MobileMoneyGH struct {
	ChargeRequestURL string `json:"-"`
	Currency         string `json:"currency"`
	Country          string `json:"country"`
	LastName         string `json:"lastname,omitempty"`
	FirstName        string `json:"firstname,omitempty"`
	IsMobileMoneyGH  int    `json:"is_mobile_money_gh"`
	Network          string `json:"network"`
}

// ChargeURL is an implemenation of the Chargeable interface
// it returns the url to be used for charging the given card
func (gh *MobileMoneyGH) ChargeURL() string {
	if gh.ChargeRequestURL == "" {
		gh.ChargeRequestURL = MobileMoneyGHChargeURL
	}
	return gh.ChargeRequestURL
}

// ValidateChargeURL is an implemenation of the Chargeable interface
// it returns the url to be used for charging the given card
func (gh *MobileMoneyGH) ValidateChargeURL() string {
	return ""
}

// BuildChargeRequestPayload is an implemenation of the Chargeable interface
// it returns the byte representation of the charge request client
// at the ChargeRequest level, chargeables are merely interface objects
// so trying to compose a struct with an interface object results in go adding the interface name key to the result bytes
// see https://play.golang.com/p/MFfbuPLrjo6
// so here we upend it so the individual concrete types do the marshalling
func (gh *MobileMoneyGH) BuildChargeRequestPayload(cReq *ChargeRequest) []byte {
	cReq.PaymentType = "mobilemoneygh"
	payload := struct {
		*MobileMoneyGH
		*ChargeRequest
	}{gh, cReq}
	b, err := json.Marshal(payload)
	if err != nil {
		log.Println("couldn't marshal payload: ", err)
	}
	return b
}
