package response

import (
	"encoding/json"
	"github.com/ddvalim/go-mail-sender/core/ports"
	"log"
	"net/http"
)

func Write(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}
}

func Error(w http.ResponseWriter, statusCode int, err error) {
	var response ports.Response

	response.Message = err.Error()
	response.StatusCode = statusCode

	Write(w, statusCode, response)
}
