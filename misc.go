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

func sendRequest(mtd, url string, payload, respObj interface{}) error {
	body, err := json.Marshal(payload)
	if err != nil {
		log.Println("Error marshalling request payload: ", err)
		return err
	}

	req, err := http.NewRequest(mtd, url, bytes.NewBuffer(body))
	if err != nil {
		log.Println("Error occured while creating request", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
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
