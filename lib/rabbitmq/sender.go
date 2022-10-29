package rabbitmq

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Sender struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewSender(url string) (*Sender, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to RabbitMQ")
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to open a channel")
	}

	return &Sender{
		conn: conn,
		ch:   ch,
	}, nil
}

func (s *Sender) SendMessage(ctx context.Context, queueName string, msg any) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return errors.Wrap(err, "Can't convert to json")
	}

	err = s.ch.PublishWithContext(ctx,
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		return errors.Wrap(err, "Can't publish message")
	}

	return nil
}

func (s *Sender) Close() {
	s.ch.Close()
	s.conn.Close()
}
