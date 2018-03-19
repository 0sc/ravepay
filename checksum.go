package rave

import (
	"crypto/sha256"
	"fmt"
	"reflect"
	"sort"
)

// CalculateChecksum implements rave's checksum integrity check for inline js
// https://flutterwavedevelopers.readme.io/docs/checksum
// To use with a payment payload, provide the payment object as the payload interface
// add the payment PBFPubKey and your secret key as byte prefix and suffix respectively
func CalculateChecksum(payload interface{}, prefix, suffix []byte) string {
	value := reflect.Indirect(reflect.ValueOf(payload))

	sortedFields := sortStructFields(value)
	checksumPayload := ""

	for _, field := range sortedFields {
		fieldVal := value.FieldByName(field)
		if fieldVal.CanInterface() && !isZeroValue(fieldVal) {
			checksumPayload = fmt.Sprintf("%s%v", checksumPayload, fieldVal)
		}
	}

	h := sha256.New()
	payloadByt := append(prefix, []byte(checksumPayload)...)
	payloadByt = append(payloadByt, suffix...)
	h.Write(payloadByt)

	return fmt.Sprintf("%x", h.Sum(nil))
}

func sortStructFields(val reflect.Value) (fields []string) {
	numField := val.NumField()
	for i := 0; i < numField; i++ {
		fields = append(fields, val.Type().Field(i).Name)
	}
	sort.Strings(fields)
	return
}

func isZeroValue(val reflect.Value) bool {
	return reflect.Zero(val.Type()).Interface() == val.Interface()
}
