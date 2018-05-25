package ravepay

import (
	"bytes"
	"crypto/des"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"strings"
)

func getEncryptionKey(seckey string) string {
	h := md5.New()
	io.WriteString(h, seckey)
	keyMD5 := fmt.Sprintf("%x", h.Sum(nil))
	keyMD5Last12 := keyMD5[len(keyMD5)-12:]

	adjustedSeckey := strings.Replace(seckey, "FLWSECK-", "", 1)
	adjustedSeckeyFirst12 := adjustedSeckey[:12]

	return adjustedSeckeyFirst12 + keyMD5Last12
}

// https://github.com/golang/go/issues/5597
func tripleDESEncrypt(payload, key []byte) string {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		log.Println("couldn't create 3DESC cipher: ", err)
		return ""
	}

	bs := block.BlockSize()
	if numStrandedBytes := len(payload) % bs; numStrandedBytes != 0 {
		paddingAmt := bs - numStrandedBytes
		padding := bytes.Repeat([]byte{byte(paddingAmt)}, paddingAmt)
		payload = append(payload, padding...)
	}

	cipher := make([]byte, len(payload))
	cipherDup := cipher
	for len(payload) > 0 {
		block.Encrypt(cipher, payload)
		payload = payload[bs:]
		cipher = cipher[bs:]
	}

	return base64.StdEncoding.EncodeToString(cipherDup)
}
