package email

import "github.com/ddvalim/go-mail-sender/core/ports"

type Service interface {
	Send(email ports.Email) error
}
