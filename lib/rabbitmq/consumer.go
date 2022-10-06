package rabbitmq

import (
	"github.com/pkg/errors"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn     *amqp.Connection
	ch       *amqp.Channel
	Messages <-chan amqp.Delivery
}

func NewConsumer(url string, queueName string) (*Consumer, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to RabbitMQ")
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to open a channel")
	}

	queue, err := ch.QueueDeclare("email", true, false, false, false, nil)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to declare a queue")
	}

	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to register a consumer")
	}

	return &Consumer{
		conn:     conn,
		ch:       ch,
		Messages: msgs,
	}, nil
}

func (c *Consumer) Close() {
	c.conn.Close()
	c.ch.Close()
}
