package email

import "github.com/ddvalim/go-mail-sender/core/ports"

type ServiceImpl struct {
}

func NewService() ServiceImpl {
	return ServiceImpl{}
}

func (s ServiceImpl) Send(email ports.Email) error {
	return nil
}
