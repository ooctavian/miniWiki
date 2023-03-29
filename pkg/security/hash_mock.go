package security

import "github.com/stretchr/testify/mock"

type HashMock struct {
	mock.Mock
}

func (m *HashMock) GenerateFormatted(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

func (m *HashMock) GenerateRaw(password string) ([]byte, []byte, error) {
	args := m.Called(password)
	return args.Get(0).([]byte), args.Get(1).([]byte), args.Error(2)
}

func (m *HashMock) Equal(password string, hash string) (bool, error) {
	args := m.Called(password, hash)
	return args.Bool(0), args.Error(1)
}
