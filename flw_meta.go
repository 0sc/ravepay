package rave

type FlwMeta struct {
	ACCOUNTVALIDATIONRESPMESSAGE  interface{} `json:"ACCOUNTVALIDATIONRESPMESSAGE"`
	ACCOUNTVALIDATIONRESPONSECODE string      `json:"ACCOUNTVALIDATIONRESPONSECODE"`
	VBVRESPONSECODE               string      `json:"VBVRESPONSECODE"`
	VBVRESPONSEMESSAGE            string      `json:"VBVRESPONSEMESSAGE"`
	ChargeResponse                string      `json:"chargeResponse"`
	ChargeResponseMessage         string      `json:"chargeResponseMessage"`
}
