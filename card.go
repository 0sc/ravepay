package rave

import (
	"encoding/json"
	"log"
)

// Card is a type that encapsulates rave's card description
// It has all card attributes necessary for rave api card references
// It also implements the chargable interface required for making charge requests and validating them
type Card struct {
	Brand      string `json:"brand,omitempty"`
	CardBIN    string `json:"cardBIN,omitempty"`
	CardNo     string `json:"cardno"`
	CardTokens []struct {
		Embedtoken string `json:"embedtoken"`
		Shortcode  string `json:"shortcode"`
	} `json:"card_tokens,omitempty"`
	ChargeCardURL         string `json:"-"`
	Country               string `json:"country"`
	Currency              string `json:"currency"`
	Cvv                   int    `json:"cvv"`
	Expirymonth           string `json:"expirymonth"`
	Expiryyear            string `json:"expiryyear"`
	FirstName             string `json:"firstname,omitempty"`
	Last4digits           string `json:"last4digits,omitempty"`
	LastName              string `json:"lastname,omitempty"`
	Pin                   string `json:"pin"`
	ValidateCardChargeURL string `json:"-"`
}

// ChargeURL is an implemenation of the Chargeable interface
// it returns the url to be used for charging the given card
func (c *Card) ChargeURL() string {
	if c.ChargeCardURL == "" {
		c.ChargeCardURL = ChargeCardURL
	}
	return c.ChargeCardURL
}

// ValidateChargeURL is an implemenation of the Chargeable interface
// it returns the url to be used for charging the given card
func (c *Card) ValidateChargeURL() string {
	if c.ValidateCardChargeURL == "" {
		c.ValidateCardChargeURL = ValidateCardChargeURL
	}
	return c.ValidateCardChargeURL
}

// BuildChargeRequestPayload is an implemenation of the Chargeable interface
// it returns the byte representation of the charge request client
// at the ChargeRequest level, chargeables are merely interface objects
// so trying to compose a struct with an interface object results in go adding the interface name key to the result bytes
// see https://play.golang.com/p/MFfbuPLrjo6
// so here we upend it so the individual concrete types do the marshalling
func (c *Card) BuildChargeRequestPayload(creq *ChargeRequest) []byte {
	creq.PaymentType = "card"
	payload := struct {
		*Card
		*ChargeRequest
	}{c, creq}
	b, err := json.Marshal(payload)
	if err != nil {
		log.Println("couldn't marshal payload: ", err)
	}
	return b
}
