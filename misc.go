package rave

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }

func sendRequestAndParseResponse(mtd, url string, payload, respObj interface{}) error {
	resp, err := sendRequest(mtd, url, payload)
	if err != nil {
		log.Println("Error occured while making request", err)
		return err
	}

	err = json.NewDecoder(resp.Body).Decode(respObj)
	if err != nil {
		log.Println("Error occured while parsing response body", err)
	}

	return err
}

func sendRequest(mtd, url string, payload interface{}) (*http.Response, error) {
	var req *http.Request
	var err error

	if payload != nil {
		body, err := json.Marshal(payload)
		if err != nil {
			log.Println("Error marshalling request payload: ", err)
			return nil, err
		}
		req, err = http.NewRequest(mtd, url, bytes.NewBuffer(body))
	} else {
		req, err = http.NewRequest(mtd, url, nil)
	}

	if err != nil {
		log.Println("Error occured while creating request", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	return client.Do(req)
}
