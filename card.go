package rave

type Card struct {
	Brand      string `json:"brand"`
	CardBIN    string `json:"cardBIN"`
	CardTokens []struct {
		Embedtoken string `json:"embedtoken"`
		Shortcode  string `json:"shortcode"`
	} `json:"card_tokens"`
	Expirymonth string `json:"expirymonth"`
	Expiryyear  string `json:"expiryyear"`
	Last4digits string `json:"last4digits"`
}
