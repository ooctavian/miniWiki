package security

// https://github.com/alexedwards/argon2id/blob/master/argon2id.go
import (
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

var (
	ErrInvalidHash         = errors.New("security is not in the correct format")
	ErrIncompatibleVariant = errors.New("incompatible variant of hash")
	ErrIncompatibleVersion = errors.New("incompatible version of hash")
)

type Argon2id struct {
	memory       uint32
	iterations   uint32
	parallelism  uint8
	saltLength   uint32
	keyLength    uint32
	generateSalt func(n uint32) ([]byte, error)
}

func NewArgon2id(
	memory uint32,
	iterations uint32,
	parallelism uint8,
	saltLength uint32,
	keyLength uint32,
	generateSalt func(n uint32) ([]byte, error),
) Argon2id {
	return Argon2id{
		memory:       memory,
		iterations:   iterations,
		parallelism:  parallelism,
		saltLength:   saltLength,
		keyLength:    keyLength,
		generateSalt: generateSalt,
	}
}

func (a Argon2id) GenerateFormatted(password string) (string, error) {
	key, salt, err := a.GenerateRaw(password)
	if err != nil {
		return "", err
	}

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Key := base64.RawStdEncoding.EncodeToString(key)

	hash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, a.memory, a.iterations, a.parallelism, b64Salt, b64Key)
	return hash, nil
}

func (a Argon2id) GenerateRaw(password string) ([]byte, []byte, error) {
	salt, err := a.generateSalt(a.saltLength)
	if err != nil {
		return []byte{}, []byte{}, err
	}

	key := argon2.IDKey([]byte(password), salt, a.iterations, a.memory, a.parallelism, a.keyLength)

	return key, salt, nil
}

func (a Argon2id) Equal(password, hash string) (bool, error) {
	ar, salt, key, err := decodeArgon2idHash(hash)
	if err != nil {
		return false, err
	}

	otherKey := argon2.IDKey([]byte(password), salt, ar.iterations, ar.memory, ar.parallelism, ar.keyLength)

	keyLen := int32(len(key))
	otherKeyLen := int32(len(otherKey))

	if subtle.ConstantTimeEq(keyLen, otherKeyLen) == 0 {
		return false, nil
	}
	if subtle.ConstantTimeCompare(key, otherKey) == 1 {
		return true, nil
	}
	return false, err
}

func decodeArgon2idHash(hash string) (*Argon2id, []byte, []byte, error) {
	values := strings.Split(hash, "$")
	if len(values) != 6 {
		return nil, nil, nil, ErrInvalidHash
	}

	if values[1] != "argon2id" {
		return nil, nil, nil, ErrIncompatibleVariant
	}

	var version int
	_, err := fmt.Sscanf(values[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, ErrIncompatibleVersion
	}

	var mem, iter uint32
	var parallelism uint8
	_, err = fmt.Sscanf(values[3], "m=%d,t=%d,p=%d", &mem, &iter, &parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err := base64.RawStdEncoding.Strict().DecodeString(values[4])
	if err != nil {
		return nil, nil, nil, err
	}

	key, err := base64.RawStdEncoding.Strict().DecodeString(values[5])
	if err != nil {
		return nil, nil, nil, err
	}

	extracted := NewArgon2id(mem, iter, parallelism, uint32(len(salt)), uint32(len(key)), nil)
	return &extracted, salt, key, nil
}
