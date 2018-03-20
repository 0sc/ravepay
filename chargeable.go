package rave

// Chargeable is an abstract representation for chargeable resources
// like cards and accounts
type Chargeable interface {
	ChargeURL() string
	ValidateChargeURL() string
	BuildChargeRequestPayload(*ChargeRequest) []byte
}

// ChargeRequest is a holds information necessary charging a given card
// it has a charge method defined on it that takes a card and proceeds to charge it with the given information
type ChargeRequest struct {
	PBFPubKey         string  `json:"PBFPubKey"`
	Amount            float64 `json:"amount"`
	ChargeType        string  `json:"charge_type"`
	DeviceFingerprint string  `json:"device_fingerprint"`
	Email             string  `json:"email"`
	IP                string  `json:"IP"`
	TxRef             string  `json:"txRef"`
	PaymentType       string  `json:"payment_type"`
	PhoneNumber       string  `json:"phonenumber"`
	RedirectURL       string  `json:"redirect_url"`
	SuggestedAuth     string  `json:"suggested_auth"`
}

// ChargeResponse is a type of rave response to a charge card request
type ChargeResponse struct {
	Data              chargeResponseData `json:"data"`
	Message           string             `json:"message"`
	Status            string             `json:"status"`
	ValidateChargeURL string             `json:"-"`
	PBFPubKey         string             `json:"-"`
}

// chargeResponseData is a representation of the nested data attribute in a charge card request response
// These attributes will not be availabe in all request response
// Each request will only have attributes that are relevant to it
type chargeResponseData struct {
	AccountID                     int                  `json:"AccountId"`
	IP                            string               `json:"IP"`
	Acctvalrespcode               interface{}          `json:"acctvalrespcode"`
	Acctvalrespmsg                interface{}          `json:"acctvalrespmsg"`
	Amount                        int                  `json:"amount"`
	Appfee                        float64              `json:"appfee"`
	AuthModelUsed                 string               `json:"authModelUsed"`
	Authurl                       string               `json:"authurl"`
	ChargeResponseCode            string               `json:"chargeResponseCode"`
	ChargeResponseMessage         string               `json:"chargeResponseMessage"`
	ChargeType                    string               `json:"charge_type"`
	ChargedAmount                 float64              `json:"charged_amount"`
	Code                          string               `json:"code"`
	CreatedAt                     string               `json:"createdAt"`
	Currency                      string               `json:"currency"`
	Customer                      Customer             `json:"customer"`
	CustomerID                    int                  `json:"customerId"`
	Customercandosubsequentnoauth bool                 `json:"customercandosubsequentnoauth"`
	Cycle                         string               `json:"cycle"`
	DeletedAt                     interface{}          `json:"deletedAt"`
	DeviceFingerprint             string               `json:"device_fingerprint"`
	FlwRef                        string               `json:"flwRef"`
	FraudStatus                   string               `json:"fraud_status"`
	ID                            int                  `json:"id"`
	IsLive                        int                  `json:"is_live"`
	Message                       string               `json:"message"`
	Merchantbearsfee              int                  `json:"merchantbearsfee"`
	Merchantfee                   int                  `json:"merchantfee"`
	Narration                     string               `json:"narration"`
	OrderRef                      string               `json:"orderRef"`
	PaymentID                     string               `json:"paymentId"`
	PaymentPage                   interface{}          `json:"paymentPage"`
	PaymentPlan                   interface{}          `json:"paymentPlan"`
	PaymentType                   string               `json:"paymentType"`
	RaveRef                       string               `json:"raveRef"`
	RedirectURL                   string               `json:"redirectUrl"`
	SettlementToken               interface{}          `json:"settlement_token"`
	Status                        string               `json:"status"`
	SuggestedAuth                 string               `json:"suggested_auth"`
	TxRef                         string               `json:"txRef"`
	ValidateInstruction           string               `json:"validateInstruction"`
	ValidateInstructions          validateInstructions `json:"validateInstructions"`
	UpdatedAt                     string               `json:"updatedAt"`
	Vbvrespcode                   string               `json:"vbvrespcode"`
	Vbvrespmessage                string               `json:"vbvrespmessage"`
}

type validateInstructions struct {
	Instruction string   `json:"instruction"`
	Valparams   []string `json:"valparams"`
}

// ChargeValidationResponse is a type of the response from rave for charge validation request
// It's a hybrid of the response for both card and account validation requests
type ChargeValidationResponse struct {
	Data struct {
		Data struct {
			Responsecode    string `json:"responsecode"`
			Responsemessage string `json:"responsemessage"`
		} `json:"data"`
		Tx chargeResponseData `json:"tx"`
	} `json:"data"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

// Charge makes the request to charge the given card
// it returns the response from the server
func (cr *ChargeRequest) Charge(chargeable Chargeable) (*ChargeResponse, error) {
	encryptionKey := getEncryptionKey(SecretKey)
	reqPayload := chargeable.BuildChargeRequestPayload(cr)

	data := tripleDESEncrypt(reqPayload, []byte(encryptionKey))

	payload := struct {
		PBFPubKey string `json:"PBFPubKey"`
		Client    string `json:"client"`
		Alg       string `json:"alg"`
	}{
		PBFPubKey: cr.PBFPubKey,
		Client:    data,
		Alg:       "3DES-24",
	}

	resp := &ChargeResponse{}
	err := sendRequestAndParseResponse("POST", chargeable.ChargeURL(), payload, resp)
	resp.ValidateChargeURL = chargeable.ValidateChargeURL()
	// r, err := sendRequest("POST", chargeable.ChargeURL(), payload)
	// io.Copy(os.Stdout, r.Body)
	return resp, err
}

// Validate handles the final part to a resource charge using the provided otp
// returns the server response
func (cr *ChargeResponse) Validate(otp string) (*ChargeValidationResponse, error) {
	if cr.PBFPubKey == "" {
		cr.PBFPubKey = PBFPubKey
	}

	payload := struct {
		PBFPubKey            string `json:"PBFPubKey"`
		TransactionReference string `json:"transaction_reference"`
		// for some weird reason this field is named differently for account charge validation
		Transactionreference string `json:"transactionreference"`
		Otp                  string `json:"otp"`
	}{
		PBFPubKey:            cr.PBFPubKey,
		TransactionReference: cr.Data.FlwRef,
		Transactionreference: cr.Data.FlwRef,
		Otp:                  otp,
	}

	resp := &ChargeValidationResponse{}
	err := sendRequestAndParseResponse("POST", cr.ValidateChargeURL, payload, resp)

	// r, err := sendRequest("POST", cr.ValidateChargeURL, payload)
	// io.Copy(os.Stdout, r.Body)
	return resp, err
}
