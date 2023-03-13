package security

type Hash interface {
	GenerateFormatted(password string) (string, error)
	GenerateRaw(password string) ([]byte, []byte, error)
	Equal(password string, hash string) (bool, error)
}
