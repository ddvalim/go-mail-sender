package client

import (
	"golang.org/x/oauth2"
	"net/http"
)

type Service interface {
	NewClient(config oauth2.Config, token *oauth2.Token) (*http.Client, error)
}
