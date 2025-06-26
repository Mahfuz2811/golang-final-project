package rabbitmq

import (
	"encoding/json"
	"final-golang-project/models"

	"github.com/streadway/amqp"
)

type RabbitMQPublisher struct {
	Channel *amqp.Channel
	Queue   amqp.Queue
}

func NewRabbitMQPublisher() (*RabbitMQPublisher, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		"email_queue", true, false, false, false, nil,
	)
	if err != nil {
		return nil, err
	}

	return &RabbitMQPublisher{Channel: ch, Queue: q}, nil
}

func (p *RabbitMQPublisher) PublishEmail(msg models.EmailMessage) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return p.Channel.Publish(
		"",           // exchange
		p.Queue.Name, // routing key
		false, false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
