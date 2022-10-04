package server

import (
	"github.com/GrabItYourself/giys-backend/user/internal/repository"
	"github.com/GrabItYourself/giys-backend/user/pkg/libproto"
)

type Server struct {
	libproto.UnimplementedUserServiceServer
	repo *repository.Repository
}

func NewServer(repo *repository.Repository) *Server {
	return &Server{
		repo: repo,
	}
}
