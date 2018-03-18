package rave

import (
	"encoding/json"
	"log"
)

// Card is a type that encapsulates rave's card description
// It has all card attributes necessary for rave api card references
type Card struct {
	Brand      string `json:"brand,omitempty"`
	CardBIN    string `json:"cardBIN,omitempty"`
	CardNo     string `json:"cardno"`
	CardTokens []struct {
		Embedtoken string `json:"embedtoken"`
		Shortcode  string `json:"shortcode"`
	} `json:"card_tokens,omitempty"`
	Country     string `json:"country"`
	Currency    string `json:"currency"`
	Cvv         int    `json:"cvv"`
	Expirymonth string `json:"expirymonth"`
	Expiryyear  string `json:"expiryyear"`
	FirstName   string `json:"firstname,omitempty"`
	LastName    string `json:"lastname,omitempty"`
	Last4digits string `json:"last4digits,omitempty"`
	Pin         string `json:"pin"`
}

// ChargeCardRequest is a holds information necessary charging a given card
// it has a charge method defined on it that takes a card and proceeds to charge it with the given information
type ChargeCardRequest struct {
	PBFPubKey         string  `json:"PBFPubKey"`
	Amount            float64 `json:"amount"`
	ChargeType        string  `json:"charge_type"`
	ChargeURL         string  `json:"-"`
	DeviceFingerprint string  `json:"device_fingerprint"`
	Email             string  `json:"email"`
	IP                string  `json:"IP"`
	TxRef             string  `json:"txRef"`
	PhoneNumber       string  `json:"phonenumber"`
	RedirectURL       string  `json:"redirect_url"`
	SuggestedAuth     string  `json:"suggested_auth"`
}

// ChargeCardResponse is a type of rave response to a charge card request
type ChargeCardResponse struct {
	Data              ChargeCardResponseData `json:"data"`
	Message           string                 `json:"message"`
	Status            string                 `json:"status"`
	ValidateChargeURL string                 `json:"-"`
	PBFPubKey         string                 `json:"-"`
}

type ChargeCardValidationResponse struct {
	Data struct {
		Data struct {
			Responsecode    string `json:"responsecode"`
			Responsemessage string `json:"responsemessage"`
		} `json:"data"`
		Tx ChargeCardResponseData `json:"tx"`
	} `json:"data"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

// ChargeCardResponseData is a representation of the nested data attribute in a charge card request response
// These attributes will not be availabe in all request response
// Each request will only have attributes that are relevant to it
type ChargeCardResponseData struct {
	AccountID                     int         `json:"AccountId"`
	IP                            string      `json:"IP"`
	Acctvalrespcode               interface{} `json:"acctvalrespcode"`
	Acctvalrespmsg                interface{} `json:"acctvalrespmsg"`
	Amount                        int         `json:"amount"`
	Appfee                        int         `json:"appfee"`
	AuthModelUsed                 string      `json:"authModelUsed"`
	Authurl                       string      `json:"authurl"`
	ChargeResponseCode            string      `json:"chargeResponseCode"`
	ChargeResponseMessage         string      `json:"chargeResponseMessage"`
	ChargeType                    string      `json:"charge_type"`
	ChargedAmount                 int         `json:"charged_amount"`
	Code                          string      `json:"code"`
	CreatedAt                     string      `json:"createdAt"`
	Currency                      string      `json:"currency"`
	Customer                      Customer    `json:"customer"`
	CustomerID                    int         `json:"customerId"`
	Customercandosubsequentnoauth bool        `json:"customercandosubsequentnoauth"`
	Cycle                         string      `json:"cycle"`
	DeletedAt                     interface{} `json:"deletedAt"`
	DeviceFingerprint             string      `json:"device_fingerprint"`
	FlwRef                        string      `json:"flwRef"`
	FraudStatus                   string      `json:"fraud_status"`
	ID                            int         `json:"id"`
	IsLive                        int         `json:"is_live"`
	Message                       string      `json:"message"`
	Merchantbearsfee              int         `json:"merchantbearsfee"`
	Merchantfee                   int         `json:"merchantfee"`
	Narration                     string      `json:"narration"`
	OrderRef                      string      `json:"orderRef"`
	PaymentID                     string      `json:"paymentId"`
	PaymentPage                   interface{} `json:"paymentPage"`
	PaymentPlan                   interface{} `json:"paymentPlan"`
	PaymentType                   string      `json:"paymentType"`
	RaveRef                       string      `json:"raveRef"`
	RedirectURL                   string      `json:"redirectUrl"`
	SettlementToken               interface{} `json:"settlement_token"`
	Status                        string      `json:"status"`
	SuggestedAuth                 string      `json:"suggested_auth"`
	TxRef                         string      `json:"txRef"`
	UpdatedAt                     string      `json:"updatedAt"`
	Vbvrespcode                   string      `json:"vbvrespcode"`
	Vbvrespmessage                string      `json:"vbvrespmessage"`
}

// Charge makes the request to charge the given card
// it returns the response from the server
func (ccr *ChargeCardRequest) Charge(card *Card) (*ChargeCardResponse, error) {
	chargeData := struct {
		ChargeCardRequest
		Card
	}{*ccr, *card}

	encryptionKey := getEncryptionKey(SecretKey)
	b, err := json.Marshal(chargeData)
	if err != nil {
		log.Println("couldn't marshal payload: ", err)
	}

	data := tripleDESEncrypt(b, []byte(encryptionKey))

	payload := struct {
		PBFPubKey string `json:"PBFPubKey"`
		Client    string `json:"client"`
		Alg       string `json:"alg"`
	}{
		PBFPubKey: ccr.PBFPubKey,
		Client:    data,
		Alg:       "3DES-24",
	}

	if ccr.ChargeURL == "" {
		ccr.ChargeURL = ChargeCardURL
	}

	resp := &ChargeCardResponse{}
	err = sendRequestAndParseResponse("POST", ccr.ChargeURL, payload, resp)
	return resp, err
}

// Validate handles the final part to a card charge
// it validates the charge card request using the provided otp
// returns the server response
func (ccr *ChargeCardResponse) Validate(otp string) (*ChargeCardValidationResponse, error) {
	if ccr.PBFPubKey == "" {
		ccr.PBFPubKey = PBFPubKey
	}

	payload := struct {
		PBFPubKey            string `json:"PBFPubKey"`
		TransactionReference string `json:"transaction_reference"`
		Otp                  string `json:"otp"`
	}{
		PBFPubKey:            ccr.PBFPubKey,
		TransactionReference: ccr.Data.FlwRef,
		Otp:                  otp,
	}

	if ccr.ValidateChargeURL == "" {
		ccr.ValidateChargeURL = ValidateCardChargeURL
	}
	resp := &ChargeCardValidationResponse{}
	err := sendRequestAndParseResponse("POST", ccr.ValidateChargeURL, payload, resp)
	return resp, err
}
