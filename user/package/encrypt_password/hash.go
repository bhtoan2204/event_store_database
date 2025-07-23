package encrypt_password

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type ArgonParam struct {
	Salt    []byte
	Time    uint32
	Memory  uint32
	Threads uint8
	KeyLen  uint32
}

func genSalt(keyLen uint32) ([]byte, error) {
	salt := make([]byte, keyLen)
	if _, err := rand.Read(salt); err != nil {
		return nil, fmt.Errorf("failed to generate salt: %w", err)
	}
	return salt, nil
}

func NewArgonParam() (*ArgonParam, error) {
	salt, err := genSalt(32)
	if err != nil {
		return nil, err
	}
	return &ArgonParam{
		Salt:    salt,
		Time:    1,
		Memory:  64 * 1024,
		Threads: 4,
		KeyLen:  32,
	}, nil
}

func (a *ArgonParam) HashPassword(password string) (string, error) {
	hashedBytes := argon2.IDKey([]byte(password), a.Salt, a.Time, a.Memory, a.Threads, a.KeyLen)

	encodedSalt := base64.RawStdEncoding.EncodeToString(a.Salt)
	encodedHash := base64.RawStdEncoding.EncodeToString(hashedBytes)

	return fmt.Sprintf("%s$%s", encodedSalt, encodedHash), nil
}

func (a *ArgonParam) VerifyPassword(storedHash, password string) (bool, error) {
	parts := strings.Split(storedHash, "$")
	if len(parts) != 2 {
		return false, fmt.Errorf("invalid stored hash format")
	}
	encodedSalt, encodedHash := parts[0], parts[1]

	salt, err := base64.RawStdEncoding.DecodeString(encodedSalt)
	if err != nil {
		return false, fmt.Errorf("failed to decode salt: %w", err)
	}

	expectedHash, err := base64.RawStdEncoding.DecodeString(encodedHash)
	if err != nil {
		return false, fmt.Errorf("failed to decode hash: %w", err)
	}

	computedHash := argon2.IDKey([]byte(password), salt, a.Time, a.Memory, a.Threads, a.KeyLen)

	if len(computedHash) != len(expectedHash) {
		return false, nil
	}
	for i := 0; i < len(computedHash); i++ {
		if computedHash[i] != expectedHash[i] {
			return false, nil
		}
	}
	return true, nil
}
