package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func ParseBodyRaw(r *http.Request) map[string]interface{} {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return nil
	}

	// Unmarshal
	var x map[string]interface{}
	err = json.Unmarshal(body, &x)
	if err != nil {
		return nil
	}
	return x
}
