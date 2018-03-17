package rave

type Account struct {
	AccountBank          string `json:"account_bank"`
	AccountIsBlacklisted int    `json:"account_is_blacklisted"`
	AccountNumber        string `json:"account_number"`
	AccountToken         struct {
		Token string `json:"token"`
	} `json:"account_token"`
	CreatedAt string      `json:"createdAt"`
	DeletedAt interface{} `json:"deletedAt"`
	FirstName string      `json:"first_name"`
	ID        int         `json:"id"`
	LastName  string      `json:"last_name"`
	UpdatedAt string      `json:"updatedAt"`
}
