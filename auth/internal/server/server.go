package server

import (
	"github.com/GrabItYourself/giys-backend/auth/internal/libproto"
	"github.com/GrabItYourself/giys-backend/auth/internal/repository"
)

type Server struct {
	libproto.UnimplementedAuthServer
	repo *repository.Repository
}

func NewServer(repo *repository.Repository) *Server {
	return &Server{
		repo: repo,
	}
}
