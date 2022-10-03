package server

import (
	"github.com/GrabItYourself/giys-backend/shop/internal/libproto"
	"gorm.io/gorm"
)

type Server struct {
	libproto.UnimplementedShopServiceServer
	pg *gorm.DB
}

func NewServer(pg *gorm.DB) *Server {
	return &Server{
		pg: pg,
	}
}
