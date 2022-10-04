package server

import (
	"github.com/GrabItYourself/giys-backend/auth/internal/config"
	"github.com/GrabItYourself/giys-backend/auth/internal/libproto"
	"github.com/GrabItYourself/giys-backend/auth/internal/repository"
	"golang.org/x/oauth2"
)

type Server struct {
	libproto.UnimplementedAuthServer
	repo        *repository.Repository
	oauthConfig *oauth2.Config
}

func NewServer(repo *repository.Repository, oauthConf *config.OAuthConfig) *Server {
	return &Server{
		repo: repo,
		oauthConfig: &oauth2.Config{
			ClientID:     oauthConf.ClientId,
			ClientSecret: oauthConf.ClientSecret,
		},
	}
}
