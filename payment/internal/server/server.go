package server

import (
	"github.com/GrabItYourself/giys-backend/lib/postgres/repository"
	"github.com/GrabItYourself/giys-backend/lib/rabbitmq"
	"github.com/GrabItYourself/giys-backend/payment/pkg/paymentproto"
	"github.com/omise/omise-go"
)

type Server struct {
	paymentproto.UnimplementedPaymentServiceServer
	omiseClient  *omise.Client
	repo         *repository.Repository
	rabbitSender *rabbitmq.Sender
}

func NewServer(omiseClient *omise.Client, repo *repository.Repository, rabbitSender *rabbitmq.Sender) (*Server, error) {
	return &Server{
		omiseClient:  omiseClient,
		repo:         repo,
		rabbitSender: rabbitSender,
	}, nil
}
