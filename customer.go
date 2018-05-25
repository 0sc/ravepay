package ravepay

type Customer struct {
	AccountID     int         `json:"AccountId"`
	CreatedAt     string      `json:"createdAt"`
	Customertoken interface{} `json:"customertoken"`
	DeletedAt     interface{} `json:"deletedAt"`
	Email         string      `json:"email"`
	FullName      string      `json:"fullName"`
	ID            int         `json:"id"`
	Phone         interface{} `json:"phone"`
	UpdatedAt     string      `json:"updatedAt"`
}
