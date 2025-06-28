package rabbitmq

import (
	"encoding/json"
	"final-golang-project/models"
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	Channel *amqp.Channel
	Queue   amqp.Queue
}

func NewRabbitMQ() (*RabbitMQ, error) {
	host := getEnv("RABBITMQ_HOST", "localhost")
	port := getEnv("RABBITMQ_PORT", "5672")
	user := getEnv("RABBITMQ_USER", "guest")
	password := getEnv("RABBITMQ_PASSWORD", "guest")
	queueName := getEnv("RABBITMQ_QUEUE", "email_queue")

	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, host, port)

	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	queue, err := channel.QueueDeclare(
		queueName,
		true,  // durable
		false, // auto-delete
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return nil, err
	}
	return &RabbitMQ{
		Channel: channel,
		Queue:   queue,
	}, nil
}

func (r *RabbitMQ) Publish(message models.EmailMessage) error {
	rawMessage, error := json.Marshal(message)
	if error != nil {
		return error
	}

	return r.Channel.Publish(
		"",           // exchange
		r.Queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        rawMessage,
		},
	)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
