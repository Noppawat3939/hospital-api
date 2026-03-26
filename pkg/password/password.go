package password

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Hash(pw string) string {
	b, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)

	if err != nil {
		log.Printf("failed generating password hash %v:", err)
		return ""
	}

	return string(b)
}

func Compare(plain, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	if err != nil {
		if err != bcrypt.ErrMismatchedHashAndPassword {
			log.Printf("failed compare password hash and password %v:", err)
		}

		return false
	}

	return true
}
