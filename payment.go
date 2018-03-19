package rave

// Payment is a type that encapsulates rave's concept of payment
// together with PaymentVerfication it contains required implementations for interacting with the payment APIs
type Payment struct {
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

// NewTxnVerificationChecklist returns a new paymentVerificationChecklist object with the given params
func NewTxnVerificationChecklist(amount int, flwRef, currency, secKey string) *TxnVerificationChecklist {
	return &TxnVerificationChecklist{
		Amount:              amount,
		Done:                false,
		FlwRef:              flwRef,
		TransactionCurrency: currency,
		SECKEY:              secKey,
		VerificationURL:     TransactionVerificationURL,
	}
}

// NewXRQTxnVerificationChecklist returns a new paymentVerificationChecklist object with the given params
// It sets up the checklist to use the rave's xrequery transaction verification endpoint
func NewXRQTxnVerificationChecklist(amount int, flwRef, txRef, currency, secKey string) *TxnVerificationChecklist {
	return &TxnVerificationChecklist{
		Amount:              amount,
		Done:                false,
		LastAttempt:         "1",
		OnlySuccessful:      "1",
		FlwRef:              flwRef,
		Flwref:              flwRef,
		TransactionCurrency: currency,
		TxRef:               txRef,
		Txref:               txRef,
		SECKEY:              secKey,
		VerificationURL:     TransactionVerificationRequeryURL,
	}
}
