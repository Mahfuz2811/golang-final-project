package models

type EmailMessage struct {
	Email             string `json:"email"`
	VerificationToken string `json:"verification_token"`
}
