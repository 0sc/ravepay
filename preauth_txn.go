package rave

// FIXME: Done to enable testing
var (
	capturePreAuthURL      = buildURL(CapturePreAuthPaymentURL)
	voidorRefundPreAuthURL = buildURL(VoidorRefundPreAuthURL)
)

// CapturePreAuthPayment makes request to rave's capture endpoint to claim preauth payments
// It takes the flwRef as param and returns the capture response and any error that occures
// https://flutterwavedevelopers.readme.io/v2.0/reference#capture
func CapturePreAuthPayment(ref string) (*ChargeResponse, error) {
	resp := &ChargeResponse{}
	payload := struct {
		SECKEY string `json:"SECKEY"`
		FlwRef string `json:"flwRef"`
	}{SecretKey, ref}

	err := sendRequestAndParseResponse("POST", capturePreAuthURL, payload, resp)

	return resp, err
}

// PreAuthResponse is a type of rave response for refund or void preauth payment response
type PreAuthResponse struct {
	Data struct {
		Data struct {
			AuthorizeID              string      `json:"authorizeId"`
			Avsresponsecode          interface{} `json:"avsresponsecode"`
			Avsresponsemessage       interface{} `json:"avsresponsemessage"`
			Otptransactionidentifier interface{} `json:"otptransactionidentifier"`
			Redirecturl              interface{} `json:"redirecturl"`
			Responsecode             string      `json:"responsecode"`
			Responsehtml             interface{} `json:"responsehtml"`
			Responsemessage          string      `json:"responsemessage"`
			Responsetoken            interface{} `json:"responsetoken"`
			Transactionreference     string      `json:"transactionreference"`
		} `json:"data"`
		Status string `json:"status"`
	} `json:"data"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

// RefundPreAuthPayment fullfiles the raves preauth feature by refunding the preauthorized paymemnt
// It takes the flwRef and makes a request to refund the txn
// It returns the response and any error that occurs
// https://flutterwavedevelopers.readme.io/v2.0/reference#refund-or-void
func RefundPreAuthPayment(ref string) (*PreAuthResponse, error) {
	resp := &PreAuthResponse{}
	payload := struct {
		Action string `json:"action"`
		FlwRef string `json:"ref"`
		SECKEY string `json:"SECKEY"`
	}{"refund", SecretKey, ref}

	err := sendRequestAndParseResponse("POST", voidorRefundPreAuthURL, payload, resp)

	return resp, err
}

// VoidPreAuthPayment fullfiles the raves preauth feature by voiding the preauthorized paymemnt
// It takes the flwRef and makes a request to void the txn
// It returns the response and any error that occurs
// https://flutterwavedevelopers.readme.io/v2.0/reference#refund-or-void
func VoidPreAuthPayment(ref string) (*PreAuthResponse, error) {
	resp := &PreAuthResponse{}
	payload := struct {
		Action string `json:"action"`
		FlwRef string `json:"ref"`
		SECKEY string `json:"SECKEY"`
	}{"void", SecretKey, ref}

	err := sendRequestAndParseResponse("POST", voidorRefundPreAuthURL, payload, resp)

	return resp, err
}
