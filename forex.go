package rave

// FIXME: Done to enable testing
var forexURL = ForexURL

// ForexParams type represents allowed params for querying rave's forex endpoint
type ForexParams struct {
	Amount              string `json:"amount"`
	OriginCurrency      string `json:"origin_currency"`
	DestinationCurrency string `json:"destination_currency"`
	SecKey              string `json:"SECKEY"`
}

// ForexResponse is raves response for forex rate request
type ForexResponse struct {
	Data struct {
		ConvertedAmount     int    `json:"converted_amount"`
		Destinationcurrency string `json:"destinationcurrency"`
		Lastupdated         string `json:"lastupdated"`
		OriginalAmount      string `json:"original_amount"`
		Origincurrency      string `json:"origincurrency"`
		Rate                int    `json:"rate"`
	} `json:"data"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

// ForexRate queries rave forex endpoint for the current rate of the currencies in the params
// it returns the request response and any error that occurs
func ForexRate(fxp *ForexParams) (*ForexResponse, error) {
	if fxp.SecKey == "" {
		fxp.SecKey = SecretKey
	}

	resp := &ForexResponse{}
	err := sendRequestAndParseResponse("POST", forexURL, fxp, resp)
	return resp, err
}
