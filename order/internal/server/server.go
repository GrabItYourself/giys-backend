package server

import (
	"github.com/GrabItYourself/giys-backend/order/pkg/orderproto"
	"gorm.io/gorm"
)

type Server struct {
	orderproto.UnimplementedOrderServer
	pg *gorm.DB
}

func NewServer(pg *gorm.DB) *Server {
	return &Server{
		pg: pg,
	}
}
