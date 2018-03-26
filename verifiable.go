package rave

// Verifiable is an abstract representation of any rave resources that can be verified
// verified here
type Verifiable interface {
	VerifyStatus() error
	VerifyCurrency(string) error
	VerifyAmount(int) error
	VerifyChargeResponseValue() error
	VerifyReference(string) error
}

// TxnVerificationChecklist encapsulates the payment verification process
// It includes details of the payment to verify
// and the verification response
type TxnVerificationChecklist struct {
	Amount              int    `json:"-"`
	FlwRef              string `json:"flw_ref,omitempty"` // for some weird reason, this just had to be different from below
	Flwref              string `json:"flwref,omitempty"`  // for some weird reason, this just had to be different from above
	LastAttempt         string `json:"last_attempt,omitempty"`
	Normalize           string `json:"normalize"`
	OnlySuccessful      string `json:"only_successful,omitempty"`
	VerificationURL     string `json:"-"`
	SECKEY              string `json:"SECKEY"`
	TransactionCurrency string `json:"-"`
	TxRef               string `json:"tx_ref,omitempty"` // for some weird reason, this just had to be different from above
	Txref               string `json:"txref,omitempty"`  // for some weird reason, this just had to be different from below
	// Done tracks whether verification has been attempted. It starts out false for new objects and changes to true after #verify is called on the object
	Done bool `json:"-"`
}

// VerifyTransaction sends a rave Transaction verfication request and then verfies the response
// It validates that verification checklist contains all the required information for making the request
// http://flw-pms-dev.eu-west-1.elasticbeanstalk.com/flwv3-pug/getpaidx/api/verify
// it also marks the verification as done
// If the verification URL is not set, it set's it to the default from config
func (tvc *TxnVerificationChecklist) VerifyTransaction() (*TxnVerificationResponse, []error) {
	// TODO: Validate checklist???
	// Make request to endpoint

	resp := &TxnVerificationResponse{}
	if tvc.VerificationURL == "" {
		tvc.VerificationURL = buildURL(TransactionVerificationURL)
	}

	err := sendRequestAndParseResponse("POST", tvc.VerificationURL, tvc, resp)
	tvc.Done = true

	if err != nil {
		return resp, []error{err}
	}

	errs := tvc.Verify(resp)
	return resp, errs
}

// VerifyXRequeryTransaction sends a rave XRequery transaction verification request and then verifies the response
// It validates that verification checklist contains all the required information for making the request
// it also marks the verification as done
// https://flutterwavedevelopers.readme.io/v1.0/reference#xrequery-transaction-verification
// If the verification URL is not set, it set's it to the default from config
func (tvc *TxnVerificationChecklist) VerifyXRequeryTransaction() (*XRQTxnVerificationResponse, []error) {
	// TODO: Validate checklist???
	// TODO: XRQT could return a data array depending on the query args. Handle that possibility
	resp := &XRQTxnVerificationResponse{}
	if tvc.VerificationURL == "" {
		tvc.VerificationURL = buildURL(TransactionVerificationRequeryURL)
	}

	err := sendRequestAndParseResponse("POST", tvc.VerificationURL, tvc, resp)
	tvc.Done = true

	if err != nil {
		return resp, []error{err}
	}

	errs := tvc.Verify(resp)
	return resp, errs
}

// Verify performs the rave's recommended verification check on the verifiable resource
// It returns an array of error for verfications that fail (if any)
// and marks the verification as Done
func (tvc *TxnVerificationChecklist) Verify(v Verifiable) []error {
	errs := []error{}

	if err := v.VerifyCurrency(tvc.TransactionCurrency); err != nil {
		errs = append(errs, err)
	}

	if err := v.VerifyAmount(tvc.Amount); err != nil {
		errs = append(errs, err)
	}

	// Todo solve Flwref vs FlwRef
	if err := v.VerifyReference(tvc.FlwRef); err != nil {
		errs = append(errs, err)
	}

	if err := v.VerifyStatus(); err != nil {
		errs = append(errs, err)
	}

	if err := v.VerifyChargeResponseValue(); err != nil {
		errs = append(errs, err)
	}

	return errs
}
