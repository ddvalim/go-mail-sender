package email

import (
	"encoding/base64"
	"fmt"
	"github.com/ddvalim/go-mail-sender/core/ports"
	"github.com/ddvalim/go-mail-sender/internal/client"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"os"
)

const tokenFile = "config/token.json"
const credentialsFile = "config/credentials.json"

type ServiceImpl struct {
	client client.Service
}

func NewService(client client.Service) ServiceImpl {
	return ServiceImpl{
		client: client,
	}
}

func (s ServiceImpl) Send(email ports.Email) error {
	credentials, err := os.ReadFile(credentialsFile)
	if err != nil {
		return err
	}

	config, err := google.ConfigFromJSON(credentials, gmail.GmailSendScope)
	if err != nil {
		return err
	}

	token, err := s.client.GetTokenFromFile(tokenFile)
	if err != nil {
		return err
	}

	client, err := s.client.NewClient(*config, token)
	if err != nil {
		return err
	}

	svr, err := gmail.New(client)
	if err != nil {
		return err
	}

	var msg gmail.Message

	raw := fmt.Sprintf("From: 'me'\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", email.To, email.Subject, email.Text)

	msg.Raw = base64.URLEncoding.EncodeToString([]byte(raw))

	_, err = svr.Users.Messages.Send("", &msg).Do()
	if err != nil {
		return err
	}

	fmt.Println("email successfully sent")

	return nil
}
