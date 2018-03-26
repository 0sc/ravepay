package rave

// FIXME: Done to enable testing
var refundTxnURL = buildURL(RefundTxnURL)

// RefundTxnResponse is rave's response for refund txn request
type RefundTxnResponse struct {
	Data struct {
		AccountID      int    `json:"AccountId"`
		AmountRefunded int    `json:"AmountRefunded"`
		FlwRef         string `json:"FlwRef"`
		TransactionID  int    `json:"TransactionId"`
		CreatedAt      string `json:"createdAt"`
		ID             int    `json:"id"`
		Status         string `json:"status"`
		UpdatedAt      string `json:"updatedAt"`
		WalletID       int    `json:"walletId"`
	} `json:"data"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

// Refund makes a refund request for txn with the given ref
// it returns rave's response and any error that occurs
func Refund(ref string) (*RefundTxnResponse, error) {
	resp := &RefundTxnResponse{}
	payload := struct {
		SECKEY string `json:"SECKEY"`
		FlwRef string `json:"ref"`
	}{SecretKey, ref}

	err := sendRequestAndParseResponse("POST", refundTxnURL, payload, resp)
	return resp, err
}
