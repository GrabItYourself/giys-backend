package server

import (
	"github.com/GrabItYourself/giys-backend/shop/pkg/shopproto"
	"gorm.io/gorm"
)

type Server struct {
	shopproto.UnimplementedShopServiceServer
	pg *gorm.DB
}

func NewServer(pg *gorm.DB) *Server {
	return &Server{
		pg: pg,
	}
}
