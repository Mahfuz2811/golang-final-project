package main

import (
	"encoding/json"
	"final-golang-project/models"
	"fmt"

	"github.com/streadway/amqp"
)

func sendEmail(email, token string) {
	// Replace with real email logic (SMTP, Mailgun, etc.)
	fmt.Printf("📧 Sending verification token '%s' to %s\n", token, email)
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("email_queue", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	msgs, err := ch.Consume(
		q.Name, "", true, false, false, false, nil,
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Worker started. Waiting for messages...")

	for d := range msgs {
		var msg models.EmailMessage
		if err := json.Unmarshal(d.Body, &msg); err != nil {
			fmt.Println("Failed to parse message:", err)
			continue
		}
		sendEmail(msg.Email, msg.VerificationToken)
	}
}
