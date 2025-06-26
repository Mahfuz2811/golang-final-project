package services

type EmailSender interface {
	SendVerificationEmail(email, verificationToken string) error
}
