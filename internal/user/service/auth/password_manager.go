package auth

import (
	"crypto/sha1"
	"fmt"
)

type Salt string

type PasswordManager struct {
	Salt Salt
}

func (mgr *PasswordManager) HashPassword(password string) string {
	fmt.Println(password)
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(mgr.Salt)))
}

func (mgr *PasswordManager) VerifyHashedPassword(password, expectedHash string) bool {
	return mgr.HashPassword(password) == expectedHash
}
