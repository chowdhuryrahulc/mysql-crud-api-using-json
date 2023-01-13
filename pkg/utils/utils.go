package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"fmt"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		fmt.Println("Body is", body)
		if err := json.Unmarshal([]byte(body), x); err != nil {
			fmt.Println("Byte Body is", []byte(body))
			return
		}
	}
}
