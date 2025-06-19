// services/email_sender.go
package services

type EmailSender interface {
	SendVerificationEmail(email, token string)
}
