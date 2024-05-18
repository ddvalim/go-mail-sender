package token

import "golang.org/x/oauth2"

type Service interface {
	GetTokenFromFile(file string) (*oauth2.Token, error)
	SaveTokenOnFile(path string, token *oauth2.Token) error
}
