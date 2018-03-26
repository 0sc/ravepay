package rave

import (
	"encoding/json"
	"log"
)

// Account is a type that encapsulates rave's account description
// It has all account attributes necessary for rave api card references
// It also implements the chargable interface required for making charge requests and validating them
type Account struct {
	AccountBank          string `json:"account_bank"`
	AccountIsBlacklisted int    `json:"account_is_blacklisted"`
	AccountNumber        string `json:"account_number"`
	AccountToken         struct {
		Token string `json:"token"`
	} `json:"account_token"`
	ChargeAccountURL         string      `json:"-"`
	ValidateAccountChargeURL string      `json:"-"`
	Country                  string      `json:"country"`
	CreatedAt                string      `json:"createdAt"`
	Currency                 string      `json:"currency"`
	DeletedAt                interface{} `json:"deletedAt"`
	FirstName                string      `json:"first_name"`
	ID                       int         `json:"id"`
	LastName                 string      `json:"last_name"`
	Passcode                 string      `json:"passcode"`
	UpdatedAt                string      `json:"updatedAt"`
}

// ChargeURL is an implemenation of the Chargeable interface
// it returns the url to be used for charging the given account
func (a *Account) ChargeURL() string {
	if a.ChargeAccountURL == "" {
		a.ChargeAccountURL = buildURL(ChargeAccountURL)
	}
	return a.ChargeAccountURL
}

// ValidateChargeURL is an implemenation of the Chargeable interface
// it returns the url to be used for validating charge on the given account
func (a *Account) ValidateChargeURL() string {
	if a.ValidateAccountChargeURL == "" {
		a.ValidateAccountChargeURL = buildURL(ValidateAccountChargeURL)
	}
	return a.ValidateAccountChargeURL
}

// BuildChargeRequestPayload is an implemenation of the Chargeable interface
// it returns the byte representation of the charge request client
// at the ChargeRequest level, chargeables are merely interface objects
// so trying to compose a struct with an interface object results in go adding the interface name key to the result bytes
// see https://play.golang.com/p/MFfbuPLrjo6
// so here we upend it so the individual concrete types do the marshalling
func (a *Account) BuildChargeRequestPayload(creq *ChargeRequest) []byte {
	creq.PaymentType = "account"
	payload := struct {
		*Account
		*ChargeRequest
	}{a, creq}
	b, err := json.Marshal(payload)
	if err != nil {
		log.Println("couldn't marshal payload: ", err)
	}
	return b
}
