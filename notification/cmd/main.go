package main

import (
	"context"
	"encoding/json"
	"os/signal"
	"syscall"

	"github.com/GrabItYourself/giys-backend/lib/logger"
	"github.com/GrabItYourself/giys-backend/lib/rabbitmq"
	"github.com/GrabItYourself/giys-backend/lib/rabbitmq/types"
	"github.com/GrabItYourself/giys-backend/notification/internal/config"
	"github.com/GrabItYourself/giys-backend/notification/internal/handler"
	"github.com/pkg/errors"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// Context
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	// Config
	conf := config.InitConfig()

	// Logger
	logger.InitLogger(&conf.Log)

	// Connect to RabbitMQ
	conn, err := amqp.Dial(conf.RabbitMQ.URL)
	if err != nil {
		logger.Fatal(errors.Wrap(err, "Failed to connect to RabbitMQ").Error())
	}
	defer conn.Close()

	emailConsumer, err := rabbitmq.NewConsumer(conn, "email", "email-consumer")
	if err != nil {
		logger.Fatal(errors.Wrap(err, "Can't initialize consumer").Error())
	}
	defer emailConsumer.Close()

	h := handler.NewHandler(&conf.EmailConfig)

	done := make(chan bool)

	go func() {
		for d := range emailConsumer.Messages {
			emailMessage := types.EmailMessage{}
			if err := json.Unmarshal([]byte(d.Body), &emailMessage); err != nil {
				logger.Panic(errors.Wrap(err, "Can't get email message").Error())
			}

			err := h.HandleEmailMessage(&emailMessage)
			if err != nil {
				logger.Panic(errors.Wrap(err, "Can't handle email message").Error())
			}
		}

		done <- true
	}()

	go func() {
		<-ctx.Done()
		logger.Info("Received shut down signal. Attempting graceful shutdown...")
		emailConsumer.Cancel()
		logger.Info("Stopped receiving message from queue")
	}()

	logger.Info("Waiting for messages")
	<-done
}
