package token

import (
	"encoding/json"
	"golang.org/x/oauth2"
	"os"
)

type ServiceImpl struct {
}

func NewService() ServiceImpl {
	return ServiceImpl{}
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

func (s ServiceImpl) SaveTokenOnFile(path string, token *oauth2.Token) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer f.Close()

	err = json.NewEncoder(f).Encode(token)
	if err != nil {
		return err
	}

	return nil
}
