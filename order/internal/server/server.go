package server

import (
	"github.com/GrabItYourself/giys-backend/lib/postgres/repository"
	"github.com/GrabItYourself/giys-backend/lib/rabbitmq"
	"github.com/GrabItYourself/giys-backend/order/pkg/orderproto"
	"github.com/GrabItYourself/giys-backend/payment/pkg/paymentproto"
)

type Server struct {
	orderproto.UnimplementedOrderServer
	repo          *repository.Repository
	rabbitSender  *rabbitmq.Sender
	paymentClient paymentproto.PaymentServiceClient
}

func NewServer(repo *repository.Repository, rabbitSender *rabbitmq.Sender, paymentClient paymentproto.PaymentServiceClient) *Server {
	return &Server{
		repo:          repo,
		rabbitSender:  rabbitSender,
		paymentClient: paymentClient,
	}
}
