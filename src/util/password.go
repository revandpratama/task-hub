package util

import "golang.org/x/crypto/bcrypt"

func VerifyPassword(dbPassword, requestPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(requestPassword))

	return err
}