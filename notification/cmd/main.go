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
)

func main() {
	// Context
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	// Config
	conf := config.InitConfig()

	// Logger
	logger.InitLogger(&conf.Log)

	// Initialize email consumer
	emailConsumer, err := rabbitmq.NewConsumer(conf.RabbitMQ.URL, conf.EmailConsumer.QueueName, conf.EmailConsumer.ConsumerName)
	if err != nil {
		logger.Panic(errors.Wrap(err, "Can't initialize consumer").Error())
	}
	defer emailConsumer.Close()

	// Initialize handler
	h := handler.NewHandler(&conf.Email)

	logger.Info("Waiting for messages")

	// Graceful Shutdown
	go func() {
		<-ctx.Done()
		logger.Info("Received shut down signal. Attempting graceful shutdown...")
		emailConsumer.Cancel()
		logger.Info("Stopped receiving message from queue")
	}()

	for d := range emailConsumer.Messages {
		emailMessage := types.EmailMessage{}
		if err := json.Unmarshal(d.Body, &emailMessage); err != nil {
			logger.Panic(errors.Wrap(err, "Can't get email message").Error())
		}
		logger.Debug("Received a message: " + string(d.Body))
		err := h.HandleEmailMessage(&emailMessage)
		if err != nil {
			logger.Panic(errors.Wrap(err, "Can't handle email message").Error())
		}
	}
}
