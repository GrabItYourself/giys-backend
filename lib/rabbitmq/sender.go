package rabbitmq

import (
	"context"
	"encoding/json"
	"time"

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

func (s *Sender) SendMessage(queueName string, msg any) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := json.Marshal(msg)
	if err != nil {
		return err
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
	return err
}

func (s *Sender) Close() {
	s.conn.Close()
	s.ch.Close()
}
