package client

import (
	"context"
	"encoding/json"
	"golang.org/x/oauth2"
	"net/http"
	"os"
)

type ServiceImpl struct {
}

func NewService() ServiceImpl {
	return ServiceImpl{}
}

func (s ServiceImpl) NewClient(config oauth2.Config, token *oauth2.Token) (*http.Client, error) {
	return config.Client(context.Background(), token), nil
}

func (s ServiceImpl) GetTokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var token oauth2.Token

	err = json.NewDecoder(f).Decode(&token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}
