package rave

import (
	"fmt"
)

const (
	txnVerificationURL        = "/flwv3-pug/getpaidx/api/verify"
	txnVerificationRequeryURL = "/flwv3-pug/getpaidx/api/xrequery"

	defaultChargeURL = "/flwv3-pug/getpaidx/api/charge"

	validateCardChargeURL    = "/flwv3-pug/getpaidx/api/validatecharge"
	validateAccountChargeURL = "/flwv3-pug/getpaidx/api/validate"
	listBanksURL             = "/flwv3-pug/getpaidx/api/flwpbf-banks.js?json=1"
	capturePreAuthPaymentURL = "/flwv3-pug/getpaidx/api/capture"
	voidOrRefundPreAuthURL   = "/flwv3-pug/getpaidx/api/refundorvoid"
	getFeeURL                = "/flwv3-pug/getpaidx/api/fee"
	refundTxnURL             = "/gpx/merchant/transactions/refund"
	forexURL                 = "/flwv3-pug/getpaidx/api/forex"

	testModeBaseURL = "http://flw-pms-dev.eu-west-1.elasticbeanstalk.com"
	liveModeBaseURL = "https://api.ravepay.co"
)

var (
	currentMode = "test"
	baseURL     = testModeBaseURL

	// PBFPubKey is your rave secret key
	PBFPubKey = "FLWPUBK-e634d14d9ded04eaf05d5b63a0a06d2f-X"
	// SecretKey is your rave secret key
	SecretKey = "FLWSECK-bb971402072265fb156e90a3578fe5e6-X"
)

// CurrentMode returns the current mode of operation, live or test
// This actual variable itself, currentMode is not exposed to prevent direct (external) modification
// (external) modification should only be done via the helper methods below
func CurrentMode() string {
	return currentMode
}

// SwitchToLiveMode changes to current operation mode to live
// Rave api requests in the live mode are made to the real live rave api servers and not the test servers
func SwitchToLiveMode() {
	currentMode = "live"
	baseURL = liveModeBaseURL
}

// SwitchToTestMode changes to current operation mode to test
// Rave api requests in the live mode are made to the test rave api servers and not the live servers
func SwitchToTestMode() {
	currentMode = "test"
	baseURL = testModeBaseURL
}

func buildURL(path string) string {
	return fmt.Sprintf("%s%s", baseURL, path)
}
