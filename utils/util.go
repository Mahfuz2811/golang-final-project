package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	// Higher cost = more computation = more secure but slower.
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func SendVerificationEmail(email, verificationToken string) {
	fmt.Printf("Send verification token %s to email %s", verificationToken, email)
}
