package handler

import (
	"github.com/GrabItYourself/giys-backend/notification/internal/config"
)

type Handler struct {
	host     string
	port     string
	email    string
	password string
}

func NewHandler(emailConfig *config.EmailConfig) *Handler {
	return &Handler{
		host:     emailConfig.Host,
		port:     emailConfig.Port,
		email:    emailConfig.Email,
		password: emailConfig.Password,
	}
}
