package rabbitmq

import (
	"context"
	"encoding/json"
	"time"

	"github.com/pkg/errors"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Sender struct {
	ch *amqp.Channel
}

func NewSender(conn *amqp.Connection) (*Sender, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to open a channel")
	}

	return &Sender{
		ch: ch,
	}, nil
}

func (s *Sender) SendMessage(queueName string, msg any) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

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
}
