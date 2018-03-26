package rave

// FIXME: Done to enable testing
var getFeeURL = buildURL(GetFeeURL)

// GetFeeRequest encapsulates the params need for requesting fee amount from the rave api
// https://flutterwavedevelopers.readme.io/v2.0/reference#get-fees
type GetFeeRequest struct {
	Amount    string `json:"amount"`
	PBFPubKey string `json:"PBFPubKey"`
	Currency  string `json:"currency"`
	PType     string `json:"ptype,omitempty"`
	Card6     string `json:"card6,omitempty"`
}

// GetFeeResponse is a type of rave's response to a get fee request
type GetFeeResponse struct {
	Data struct {
		ChargeAmount string  `json:"charge_amount"`
		Fee          float64 `json:"fee"`
		Merchantfee  string  `json:"merchantfee"`
		Ravefee      string  `json:"ravefee"`
	} `json:"data"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

// GetFee returns rave's fee response for the given get fee request
// it returns any error that occures
func GetFee(p *GetFeeRequest) (*GetFeeResponse, error) {
	// TODO: add request params validation
	if p.PBFPubKey == "" {
		p.PBFPubKey = PBFPubKey
	}

	resp := &GetFeeResponse{}

	err := sendRequestAndParseResponse("POST", getFeeURL, p, resp)
	return resp, err

}
