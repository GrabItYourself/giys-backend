package server

import (
	"github.com/GrabItYourself/giys-backend/lib/postgres/repository"
	"github.com/GrabItYourself/giys-backend/order/pkg/orderproto"
)

type Server struct {
	orderproto.UnimplementedOrderServer
	repo *repository.Repository
}

func NewServer(repo *repository.Repository) *Server {
	return &Server{
		repo: repo,
	}
}
