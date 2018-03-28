# Rave

```go
import "github.com/0sc/rave"
```

## Setup
Sign up for a rave account; [here](http://rave.frontendpwc.com) for a sandbox account and [here](https://rave.flutterwave.com) for a live account. Follow the instructions [here](https://flutterwavedevelopers.readme.io/v1.0/docs/prerequisite) to generate relevant public and private keys. Set your keys as enviroment variables like so:

```bash
export RAVE_PUBLIC_KEY=FLWPUBK-your-rave-public-key
export RAVE_SECRET_KEY=FLWSECK-your-rave-secret-key
export RAVE_MODE=test # can be either of test or live
```

Be sure that the set mode matches the keys provided; use your **sandbox account** keys for the `test` mode and your **live account** keys for the `live` mode. Rave will be sad if you do otherwise.

## Usage

### Card
```go
package main

import (
	"fmt"
	"log"

	"github.com/0sc/rave"
)

func main(){
  card := &rave.Card{
		CardNo:      "5438898014560229",
		Currency:    "NGN",
		Country:     "NG",
		Cvv:         "789",
		Expirymonth: "09",
		Expiryyear:  "19",
		Pin:         "3310",
  }
  
  chargeRequest := rave.ChargeRequest{
		Amount:            300,
		Email:             "tester@flutter.co",
		IP:                "103.238.105.185",
		TxRef:             "'MXX-ASC-4579",
		PhoneNumber:       "0926420185",
		DeviceFingerprint: "69e6b7f0sb72037aa8428b70fbe03986c",
  }
  
  chargeResponse, err := chargeRequest.Charge(card)
	if err != nil {
		log.Println(err)
	}

	validationResponse, err := chargeResponse.OTPValidation("123")
	if err != nil {
		log.Println(err)
  }
  
  fmt.Println(validationResponse.Message)
}
```

### Account
```go
package main

import (
	"fmt"
	"log"

	"github.com/0sc/rave"
)

func main(){
  account := &rave.Account{
		AccountBank:   "044",
		AccountNumber: "0690000031",
		Country:       "NG",
  }
  
  chargeRequest := rave.ChargeRequest{
		Amount:            300,
		Email:             "tester@flutter.co",
		IP:                "103.238.105.185",
		TxRef:             "'MXX-ASC-4579",
		PhoneNumber:       "0926420185",
		DeviceFingerprint: "69e6b7f0sb72037aa8428b70fbe03986c",
  }
  
  chargeResponse, err := chargeRequest.Charge(account)
	if err != nil {
		log.Println(err)
	}

	validationResponse, err := chargeResponse.OTPValidation("123")
	if err != nil {
		log.Println(err)
  }
  
  fmt.Println(validationResponse.Message)
}
```

### Transaction 
#### Status check
```go
package main

import (
	"fmt"

	"github.com/0sc/rave"
)

func main(){
  ref := "FLW-MOCK-a08e154ad6bcd97fba7fa66e4438614c"
  expAmt := 300
  currency := "NGN"

  txnChecklist := rave.NewTxnVerificationChecklist(expAmt, ref, currency)
  _, errs := txnChecklist.VerifyTransaction()

  if len(errs) == 0 {
    fmt.Println("Transaction checks out 🎉")
  } else {
    fmt.Println(errs)
  }
}
```

#### Status check using XRequery
```go
package main

import (
	"fmt"

	"github.com/0sc/rave"
)

func main(){
  flwRef := "FLW-MOCK-a08e154ad6bcd97fba7fa66e4438614c"
  txRef := "'MXX-ASC-4579"
  expAmt := 300
  currency := "NGN"

  txnChecklist := rave.NewXRQTxnVerificationChecklist(expAmt, ref, currency)
  _, errs := txnChecklist.VerifyXRequeryTransaction()

  if len(errs) == 0 {
    fmt.Println("Transaction checks out 🎉")
  } else {
    fmt.Println(errs)
  }
}
```

### Preauth
```go
package main

import (
  "fmt"
  "log"

	"github.com/0sc/rave"
)

func main(){
  card := &rave.Card{
		CardNo:      "5840406187553286",
		Currency:    "NGN",
		Country:     "NG",
		Cvv:         "116",
		Expirymonth: "09",
		Expiryyear:  "18",
		Pin:         "1111",
  }
  
  chargeRequest := rave.ChargeRequest{
		Amount:            300,
		ChargeType:        "preauth",
		Email:             "tester@flutter.co",
		IP:                "103.238.105.185",
		TxRef:             "'MXX-ASC-4579",
		PhoneNumber:       "0926420185",
		DeviceFingerprint: "69e6b7f0sb72037aa8428b70fbe03986c",
  }
  
  chargeResponse, err := chargeRequest.Charge(card)
	if err != nil {
		log.Println(err)
  }
  ref := chargeResponse.Data.FlwRef
  
  /// CAPTURE PAYMENT
  resp, err := rave.CapturePreAuthPayment(ref)
  if err != nil {
		log.Println(err)
  }
  fmt.Println(resp.Status)

  /// VOID PAYMENT
  res, err := rave.VoidPreAuthPayment(ref)
  if err != nil {
		log.Println(err)
  }
  fmt.Println(res.Status)

  /// REFUND PAYMENT
  res, err = rave.RefundPreAuthPayment(chargeRef)
  if err != nil {
		log.Println(err)
  }
  fmt.Println(res.Status)
}
```

### List Banks
```go
package main

import (
  "fmt"
  "log"

	"github.com/0sc/rave"
)

func main(){
  banks, err := rave.ListBanks()
  if err != nil {
    log.Println(err)
  }

  fmt.Println(banks)
}
```

### Get Fees
```go
package main

import (
  "fmt"
  "log"

	"github.com/0sc/rave"
)

func main(){
  feeReq := &rave.GetFeeRequest{
    Amount:   "100",
		Currency: "USD",
  }

  resp, err := rave.GetFee(feeReq)
  if err != nil {
		log.Println(err)
  }
  fmt.Println(resp.Data.Fee)
}
```

### Exchange Rates
```go
package main

import (
  "fmt"
  "log"

	"github.com/0sc/rave"
)

func main(){
  feeReq := &rave.ForexParams{
    Amount:   "100",
    OriginCurrency: "USD",
    DestinationCurrency: "NGN",
  }

  resp, err := rave.GetFee(feeReq)

  if err != nil {
    log.Println(err)
  }
  fmt.Printf("Converted amount: %d, rate %d \n", resp.Data.ConvertedAmount, resp.Data.Rate)
}
```

### Direct Charge Refunds
```go
package main

import (
  "fmt"
  "log"

	"github.com/0sc/rave"
)

func main(){
  ref := "FLW-MOCK-a08e154ad6bcd97fba7fa66e4438614c"
  resp, err := rave.Refund(ref)
  if err != nil {
    log.Println(err)
  }
  fmt.Println(resp.Status)
}
```

### Alternative Payments
#### USSD

```go
package main

import (
  "fmt"
  "log"

	"github.com/0sc/rave"
)

func main(){
  ussd := &rave.USSD{
		AccountBank:   "044",
		AccountNumber: "0690000031",
		Currency:      "NGN",
		Country:       "NG",
		FirstName:     "jsksk",
		LastName:      "ioeoe",
  }
  
  chargeRequest := rave.ChargeRequest{
    Amount:            300,
    PaymentType:       "ussd",
		Email:             "tester@flutter.co",
		IP:                "103.238.105.185",
		TxRef:             "'MXX-ASC-4579",
		PhoneNumber:       "0926420185",
		DeviceFingerprint: "69e6b7f0sb72037aa8428b70fbe03986c",
  }
  
  chargeResponse, err := chargeRequest.Charge(ussd)
	if err != nil {
		log.Println(err)
  }
  
  fmt.Printf("%+v\n", rave.USSDPaymentInstruction(chargeResponse))

	validationResponse, err := chargeResponse.OTPValidation("123")
	if err != nil {
		log.Println(err)
  }
  
  fmt.Println(validationResponse.Message)
}
```

#### Mpesa
```go
package main

import (
  "fmt"
  "log"

	"github.com/0sc/rave"
)

func main(){
  mpesa := &rave.Mpesa{
		Currency:  "KES",
		Country:   "NG",
		FirstName: "jsksk",
		LastName:  "ioeoe",
		IsMpesa:   "1",
	}
  
  chargeRequest := rave.ChargeRequest{
    Amount:            300,
    PaymentType:       "mpesa",
		Email:             "tester@flutter.co",
		IP:                "103.238.105.185",
		TxRef:             "'MXX-ASC-4579",
		PhoneNumber:       "0926420185",
		DeviceFingerprint: "69e6b7f0sb72037aa8428b70fbe03986c",
  }
  
  chargeResponse, err := chargeRequest.Charge(mpesa)
	if err != nil {
		log.Println(err)
	}

  instructions := rave.MpesaPaymentInstruction(chargeResponse)
	fmt.Printf("Complete transaction by sending %d to %s\n", instructions.Amount, instructions.BusinessNumber)
}
```
#### Mobile Money Ghana
```go
package main

import (
  "fmt"
  "log"

	"github.com/0sc/rave"
)

func main(){
  mm := &rave.MobileMoneyGH{
		Currency:        "GHS",
		Country:         "GH",
		Network:         "MTN",
		IsMobileMoneyGH: 1,
	}
  
  chargeRequest := rave.ChargeRequest{
    Amount:            300,
    PaymentType:       "mobilemoneygh",
		Email:             "tester@flutter.co",
		IP:                "103.238.105.185",
		TxRef:             "'MXX-ASC-4579",
		PhoneNumber:       "0926420185",
		DeviceFingerprint: "69e6b7f0sb72037aa8428b70fbe03986c",
  }
  
  chargeResponse, err := chargeRequest.Charge(mpesa)
	if err != nil {
		log.Println(err)
	}

  fmt.Println(chargeResponse.Status)
}
```

### Checksum
```go
  package main

  import (
  "fmt"

	"github.com/0sc/rave"
)

  func main(){
    payload := rave.Payment{
      Amount:            20,
      PaymentMethod:     "both",
      CustomDescription: "Pay Internet",
      CustomLogo:        "http://localhost/payporte-3/skin/frontend/ultimo/shoppy/custom/images/logo.svg",
      CustomTitle:       "Shoppy Global systems",
      Country:           "NG",
      Currency:          "NGN",
      CustomerEmail:     "user@example.com",
      CustomerFirstname: "Temi",
      CustomerLastname:  "Adelewa",
      CustomerPhone:     "234099940409",
      TxRef:             "MG-1500041286295",
    }
    prefix := []byte(rave.PublicKey)
    suffix := []byte(rave.Secret)

    checksum := rave.CalculateChecksum(payload, prefix, suffix)
    fmt.Println(checksum)
  }
```

### Utils
```go
package main

import (
  "fmt"
  "log"

	"github.com/0sc/rave"
)

 func main(){
   // check currentMode
   rave.CurrentMode() // => live or test

   // manually set public key
   rave.PublicKey = FLWPUBK-your-rave-public-key
 
   // manually set secret key
   rave.SecretKey = FLWSECK-your-rave-secret-key

   // switch to live mode
   rave.SwitchToLiveMode()

   // switch to test mode
   rave.SwitchToTestMode()
 }
```

## Todo
[ ] Add pre request validations where applicable

[ ] Implement a cleaner way for reasonably handling the inconsistences with Rave API request/response parameters

[ ] Switch to pointers for json struct fields to better manage Go's default type