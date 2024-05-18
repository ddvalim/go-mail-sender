package client

import (
	"context"
	"golang.org/x/oauth2"
	"net/http"
)

type ServiceImpl struct {
}

func NewService() ServiceImpl {
	return ServiceImpl{}
}

func (s ServiceImpl) NewClient(config oauth2.Config, token *oauth2.Token) (*http.Client, error) {
	return config.Client(context.Background(), token), nil
}
