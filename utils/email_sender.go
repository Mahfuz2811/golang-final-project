// utils/email_sender.go
package utils

import "fmt"

type DefaultEmailSender struct{}

func (e *DefaultEmailSender) SendVerificationEmail(email, verificationToken string) {
	fmt.Printf("Send verification token %s to email %s", verificationToken, email)
}
