package util

import "golang.org/x/crypto/bcrypt"

func ValidatePassword(hashedPassword, password string) error {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err

	// Extract salt from hashedPassword
	// salt, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// if err != nil {
	// 	return err
	// }
	// // Compare hash with salted password
	// return bcrypt.CompareHashAndPassword(salt, []byte(password))
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(bytes), err
}
