package main

import (
	"encoding/json"

	"github.com/GrabItYourself/giys-backend/lib/logger"
	"github.com/GrabItYourself/giys-backend/lib/rabbitmq"
	"github.com/GrabItYourself/giys-backend/lib/rabbitmq/types"
	"github.com/GrabItYourself/giys-backend/notification/internal/config"
	"github.com/GrabItYourself/giys-backend/notification/internal/handler"
	"github.com/pkg/errors"
)

func main() {
	// Config
	conf := config.InitConfig()

	// Logger
	logger.InitLogger(&conf.Log)

	emailConsumer, err := rabbitmq.NewConsumer(conf.RabbitMQ.URL, "email")
	if err != nil {
		logger.Panic(errors.Wrap(err, "Can't initialize consumer").Error())
	}
	defer emailConsumer.Close()

	h := handler.NewHandler(&conf.EmailConfig)

	var forever chan struct{}

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
	}()

	logger.Info("Waiting for messages")
	<-forever
}
