package main

import "github.com/streadway/amqp"

func main() {
	// this is a subscriber worker for the message queue
	// it will listen for messages from the channel email_queue and process them

	// connect to the rabbitmq server

	amqpConn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer amqpConn.Close()

	// create a channel
	ch, err := amqpConn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	// declare a queue
	q, err := ch.QueueDeclare(
		"email_queue", // name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		panic(err)
	}
	// consume messages from the queue
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // arguments
	)

	if err != nil {
		panic(err)
	}

	for msg := range msgs {
		// process the message
		// in this case, we will just print the message body
		// in a real application, you would send an email or perform some other action
		println("Received a message:", string(msg.Body))
	}
}
