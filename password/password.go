package password

import (
	"golang.org/x/crypto/bcrypt"
)

func CreatePasswordHash(password string) string {
	hash, error := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if error != nil {
		panic(error)
	}
	return string(hash)
}
