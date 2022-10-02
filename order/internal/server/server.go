package server

import (
	"github.com/GrabItYourself/giys-backend/order/libproto"
	"gorm.io/gorm"
)

type Server struct {
	libproto.UnimplementedOrderServer
	pg *gorm.DB
}

func NewServer(pg *gorm.DB) *Server {
	return &Server{
		pg: pg,
	}
}
