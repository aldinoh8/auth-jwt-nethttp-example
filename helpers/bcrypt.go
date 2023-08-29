package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	byteHashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(byteHashed)
}

func ValidatePassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	// if err != nil {
	// 	return false
	// } else {
	// 	return true
	// }

	return err == nil
}

// CompareHashAndPassword(hashedPassword, password []byte) error

// func GenerateFromPassword(password []byte, cost int) ([]byte, error)
