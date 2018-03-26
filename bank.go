package rave

// Bank is a type of rave bank resources
type Bank struct {
	Code            string `json:"bankcode"`
	Name            string `json:"bankname"`
	Internetbanking bool   `json:"internetbanking"`
}

// FIXME: Done to enable testing
var banksURL = buildURL(listBanksURL)

// ListBanks returns list of banks from the rave api
// https://flutterwavedevelopers.readme.io/v1.0/reference#list-of-banks
func ListBanks() ([]Bank, error) {
	banks := []Bank{}

	err := sendRequestAndParseResponse("GET", banksURL, nil, &banks)
	return banks, err
}
