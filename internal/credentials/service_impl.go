package credentials

import "os"

const credentialsFile = "config/credentials.json"

type ServiceImpl struct {
}

func NewService() ServiceImpl {
	return ServiceImpl{}
}

func (s ServiceImpl) GetCredentials() ([]byte, error) {
	credentials, err := os.ReadFile(credentialsFile)
	if err != nil {
		return nil, err
	}

	return credentials, nil
}
