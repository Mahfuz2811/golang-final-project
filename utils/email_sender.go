package utils

import (
	"final-golang-project/models"
	"final-golang-project/rabbitmq"
	"fmt"
)

type DefaultEmailSender struct {
	Publisher *rabbitmq.RabbitMQPublisher
}

func NewDefaultEmailSender(p *rabbitmq.RabbitMQPublisher) *DefaultEmailSender {
	return &DefaultEmailSender{Publisher: p}
}

func (e *DefaultEmailSender) SendVerificationEmail(email, verificationToken string) {
	msg := models.EmailMessage{
		Email:             email,
		VerificationToken: verificationToken,
	}

	if err := e.Publisher.PublishEmail(msg); err != nil {
		fmt.Println("Failed to publish email message:", err)
	}
}
