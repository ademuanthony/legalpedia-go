package web

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
)

func RenderErrorfJSON(res http.ResponseWriter, errorMessage string, args ...interface{}) {
	var response = struct {
		Success             bool        `json:"success"`
		Result              interface{} `json:"result"`
		UnAuthorizedRequest bool        `json:"unAuthorizedRequest"`
		Error               interface{} `json:"error"`
	}{
		Success:             false,
		Result:              nil,
		UnAuthorizedRequest: false,
		Error:               fmt.Sprintf(errorMessage, args...),
	}

	d, err := json.Marshal(response)
	if err != nil {
		log.Errorf("Error marshalling data: %s", err.Error())
	}

	res.Header().Set("Content-Type", "application/json")
	_, _ = res.Write(d)
}

func RenderJSON(res http.ResponseWriter, data interface{}) {
	var response = struct {
		Success             bool        `json:"success"`
		Result              interface{} `json:"result"`
		UnAuthorizedRequest bool        `json:"unAuthorizedRequest"`
		Error               interface{} `json:"error"`
	}{
		Success:             true,
		Result:              data,
		UnAuthorizedRequest: false,
		Error:               nil,
	}

	d, err := json.Marshal(response)
	if err != nil {
		log.Errorf("Error marshalling data: %s", err.Error())
	}

	res.Header().Set("Content-Type", "application/json")
	_, _ = res.Write(d)
}

// RenderJSONBytes prepares the headers for pre-encoded JSON and writes the JSON
// bytes.
func RenderJSONBytes(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err := w.Write(data)
	if err != nil {
		// Filter out broken pipe (user pressed "stop") errors
		if _, ok := err.(*net.OpError); ok {
			if strings.Contains(err.Error(), "broken pipe") {
				return
			}
		}
		log.Warnf("ResponseWriter.Write error: %v", err)
	}
}
