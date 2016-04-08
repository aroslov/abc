package util

import (
	"log"
	"github.com/streadway/amqp"
)

func RabbitMQConnect(rabbit_endpoint string, rabbit_exchange string) (*amqp.Connection, *amqp.Channel) {
	log.Printf("Connecting to RabbitMQ on %s\n", rabbit_endpoint)
	conn, err := amqp.Dial(rabbit_endpoint)
	FailOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")

	err = ch.ExchangeDeclare(
		rabbit_exchange, // name
		"fanout", // type
		true, // durable
		false, // auto-deleted
		false, // internal
		false, // no-wait
		nil, // arguments
	)
	FailOnError(err, "Failed to declare an exchange")

	return conn, ch
}

func RabbitMQGetMessages(rabbit_exchange string, ch *amqp.Channel) (<- chan amqp.Delivery) {
	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	FailOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name, // queue name
		"",     // routing key
		rabbit_exchange, // exchange
		false,
		nil)
	FailOnError(err, "Failed to bind a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	FailOnError(err, "Failed to register a consumer")
	return msgs
}

