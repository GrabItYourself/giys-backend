package handler

import (
	"net/smtp"

	"github.com/GrabItYourself/giys-backend/lib/logger"
	"github.com/pkg/errors"

	"github.com/GrabItYourself/giys-backend/lib/rabbitmq/types"
)

func (h *Handler) HandleEmailMessage(emailMessage *types.EmailMessage) error {
	address := h.host + ":" + h.port
	to := []string{emailMessage.To}
	message := []byte("Subject: " + emailMessage.Subject + "\n" + emailMessage.Body)
	auth := smtp.PlainAuth("", h.email, h.password, h.host)

	logger.Debug("Sending email to " + emailMessage.To)
	err := smtp.SendMail(address, auth, h.email, to, message)
	if err != nil {
		return errors.Wrap(err, "Failed to send an email")
	}
	logger.Debug("Email sent to " + emailMessage.To + "successfully")
	return nil
}
