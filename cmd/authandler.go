package cmd

import (
	"context"
	"fmt"
	"github.com/ddvalim/go-mail-sender/cmd/response"
	"github.com/ddvalim/go-mail-sender/internal/credentials"
	"github.com/ddvalim/go-mail-sender/internal/token"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"net/http"
)

type AuthHandler struct {
	credentials credentials.Service
	token       token.Service
}

func NewAuthHandler() AuthHandler {
	credentialsService := credentials.NewService()

	tokenService := token.NewService()

	return AuthHandler{
		credentials: credentialsService,
		token:       tokenService,
	}
}

func (h AuthHandler) Auth(w http.ResponseWriter, r *http.Request) {
	credentials, err := h.credentials.GetCredentials()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
	}

	oauthConfig, err := google.ConfigFromJSON(credentials, gmail.GmailSendScope)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
	}

	url := oauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	fmt.Println(url)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (h AuthHandler) Callback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	credentials, err := h.credentials.GetCredentials()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
	}

	oauthConfig, err := google.ConfigFromJSON(credentials, gmail.GmailSendScope)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
	}

	tok, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
	}

	err = h.token.SaveTokenOnFile("config/token.json", tok)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
	}

	response.Write(w, http.StatusOK, "ok")
}
