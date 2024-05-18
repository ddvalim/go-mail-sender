package credentials

type Service interface {
	GetCredentials() ([]byte, error)
}
