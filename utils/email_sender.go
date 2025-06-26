package utils

import (
	"final-golang-project/models"
	"final-golang-project/rabbitmq"
)

type EmailSender struct {
	Publisher *rabbitmq.RabbitMQ
}

func NewEmailSender(publisher *rabbitmq.RabbitMQ) *EmailSender {
	return &EmailSender{
		Publisher: publisher,
	}
}

func (es *EmailSender) SendVerificationEmail(email, verificationToken string) error {
	message := models.EmailMessage{
		Email:             email,
		VerificationToken: verificationToken,
	}

	err := es.Publisher.Publish(message)
	if err != nil {
		return err
	}

	return nil
}
