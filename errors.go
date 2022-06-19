package FAQBot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Error struct {
	Error string `json:"error"`
}

func CheckErrors(r *http.Response) string {
	data, err := ioutil.ReadAll(r.Body)

	r.Body.Close()

	r.Body = ioutil.NopCloser(bytes.NewBuffer(data))

	if err != nil {
		return err.Error()
	}

	var e *Error

	err = json.Unmarshal(data, &e)
	if err != nil {
		return err.Error()
	}

	if e.Error != "" {
		return fmt.Sprintf("Error: %s", e.Error)
	}

	if r.StatusCode != 200 {
		return fmt.Sprintf("Error: %s", string(data))
	}

	return "ok"
}
