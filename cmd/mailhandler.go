package cmd

import (
	"encoding/json"
	"github.com/ddvalim/go-mail-sender/cmd/response"
	"github.com/ddvalim/go-mail-sender/core/ports"
	"github.com/ddvalim/go-mail-sender/internal/email"
	"io/ioutil"
	"net/http"
)

type Handler struct {
	mailService email.Service
}

func NewHandler() Handler {
	mailService := email.NewService()

	return Handler{
		mailService: mailService,
	}
}

func (h Handler) Send(w http.ResponseWriter, r *http.Request) {
	requestBody, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		response.Error(w, http.StatusInternalServerError, readErr)
	}

	var mail ports.Email

	unmarshallErr := json.Unmarshal(requestBody, &mail)
	if unmarshallErr != nil {
		response.Error(w, http.StatusInternalServerError, unmarshallErr)
	}

	sendErr := h.mailService.Send(mail)
	if sendErr != nil {
		response.Error(w, http.StatusBadRequest, sendErr)
	}

	response.Write(w, http.StatusOK, nil)
}
