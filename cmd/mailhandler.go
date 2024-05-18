package cmd

import (
	"encoding/json"
	"github.com/ddvalim/go-mail-sender/cmd/response"
	"github.com/ddvalim/go-mail-sender/core/ports"
	"github.com/ddvalim/go-mail-sender/internal/client"
	"github.com/ddvalim/go-mail-sender/internal/credentials"
	"github.com/ddvalim/go-mail-sender/internal/email"
	"github.com/ddvalim/go-mail-sender/internal/token"
	"io/ioutil"
	"net/http"
)

type MailHandler struct {
	mailService email.Service
}

func NewMailHandler() MailHandler {
	clientService := client.NewService()

	credentialsService := credentials.NewService()

	tokenService := token.NewService()

	mailService := email.NewService(clientService, credentialsService, tokenService)

	return MailHandler{
		mailService: mailService,
	}
}

func (h MailHandler) Send(w http.ResponseWriter, r *http.Request) {
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
