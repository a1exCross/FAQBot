package FAQBot

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

const url string = "https://7319-94-241-204-46.eu.ngrok.io"

func request(method string, body []byte) ([]byte, error) {
	var Client http.Client

	req, err := http.NewRequest(http.MethodPost, url+method, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	res, err := Client.Do(req)

	check := CheckErrors(res)
	if check != "ok" {
		return nil, errors.New(check)
	}

	return ioutil.ReadAll(res.Body)
}
