package rave

const (
	// TransactionVerificationURL is Rave's verification URL
	TransactionVerificationURL        = "http://flw-pms-dev.eu-west-1.elasticbeanstalk.com/flwv3-pug/getpaidx/api/verify"
	TransactionVerificationRequeryURL = "http://flw-pms-dev.eu-west-1.elasticbeanstalk.com/flwv3-pug/getpaidx/api/xrequery"
	ChargeCardURL                     = "http://flw-pms-dev.eu-west-1.elasticbeanstalk.com/flwv3-pug/getpaidx/api/charge"
	ValidateCardChargeURL             = "http://flw-pms-dev.eu-west-1.elasticbeanstalk.com/flwv3-pug/getpaidx/api/validatecharge"
	ChargeAccountURL                  = "http://flw-pms-dev.eu-west-1.elasticbeanstalk.com/flwv3-pug/getpaidx/api/charge"
	ValidateAccountChargeURL          = "http://flw-pms-dev.eu-west-1.elasticbeanstalk.com/flwv3-pug/getpaidx/api/validate"
	ListBanksURL                      = "http://flw-pms-dev.eu-west-1.elasticbeanstalk.com/flwv3-pug/getpaidx/api/flwpbf-banks.js?json=1"
	SecretKey                         = "FLWSECK-bb971402072265fb156e90a3578fe5e6-X"
	PBFPubKey                         = "FLWPUBK-e634d14d9ded04eaf05d5b63a0a06d2f-X"
	CapturePreAuthPaymentURL          = "http://flw-pms-dev.eu-west-1.elasticbeanstalk.com/flwv3-pug/getpaidx/api/capture"
	VoidorRefundPreAuthURL            = "http://flw-pms-dev.eu-west-1.elasticbeanstalk.com/flwv3-pug/getpaidx/api/refundorvoid"
	GetFeeURL                         = "http://flw-pms-dev.eu-west-1.elasticbeanstalk.com/flwv3-pug/getpaidx/api/fee"
	RefundTxnURL                      = "http://flw-pms-dev.eu-west-1.elasticbeanstalk.com/gpx/merchant/transactions/refund"
	ForexURL                          = "http://flw-pms-dev.eu-west-1.elasticbeanstalk.com/flwv3-pug/getpaidx/api/forex"
)
