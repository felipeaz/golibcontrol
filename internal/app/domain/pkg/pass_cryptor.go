package pkg

import (
	"golang.org/x/crypto/bcrypt"
)

type Cryptor struct{}

func (c *Cryptor) EncryptPassword(password string) (string, error) {
	return c.hashPassword(password)
}

func (c *Cryptor) hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 16)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePasswordAndHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
