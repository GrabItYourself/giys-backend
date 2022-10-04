package server

import (
	"github.com/GrabItYourself/giys-backend/user/internal/repository"
	"github.com/GrabItYourself/giys-backend/user/pkg/userproto"
)

type Server struct {
	userproto.UnimplementedUserServiceServer
	repo *repository.Repository
}

func NewServer(repo *repository.Repository) *Server {
	return &Server{
		repo: repo,
	}
}
