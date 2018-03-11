package rave

type Payment struct {
	// PBFPubKey         string `json:"PBFPubKey"`
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
	Txref             string `json:"txref"`
}

func (p *Payment) createIntegerityHash() {

}
