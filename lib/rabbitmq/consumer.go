package rabbitmq

import (
	"github.com/pkg/errors"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	name     string
	ch       *amqp.Channel
	Messages <-chan amqp.Delivery
}

func NewConsumer(conn *amqp.Connection, queueName string, consumerName string) (*Consumer, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to open a channel")
	}

	queue, err := ch.QueueDeclare("email", true, false, false, false, nil)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to declare a queue")
	}

	msgs, err := ch.Consume(
		queue.Name,   // queue
		consumerName, // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to register a consumer")
	}

	return &Consumer{
		name:     consumerName,
		ch:       ch,
		Messages: msgs,
	}, nil
}

func (c *Consumer) Close() {
	c.ch.Close()
}

func (c *Consumer) Cancel() {
	c.ch.Cancel(c.name, false)
}
