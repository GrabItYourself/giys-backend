package server

import (
	"github.com/GrabItYourself/giys-backend/lib/postgres/repository"
	"github.com/GrabItYourself/giys-backend/lib/rabbitmq"
	"github.com/GrabItYourself/giys-backend/order/pkg/orderproto"
)

type Server struct {
	orderproto.UnimplementedOrderServer
	repo         *repository.Repository
	rabbitSender *rabbitmq.Sender
}

func NewServer(repo *repository.Repository, rabbitSender *rabbitmq.Sender) *Server {
	return &Server{
		repo:         repo,
		rabbitSender: rabbitSender,
	}
}
