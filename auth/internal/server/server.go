package server

import (
	"github.com/GrabItYourself/giys-backend/auth/internal/config"
	"github.com/GrabItYourself/giys-backend/auth/internal/repository"
	"github.com/GrabItYourself/giys-backend/auth/pkg/authproto"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Server struct {
	authproto.UnimplementedAuthServer
	repo        *repository.Repository
	oauthConfig *oauth2.Config
}

func NewServer(repo *repository.Repository, oauthConf *config.OAuthConfig) *Server {
	return &Server{
		repo: repo,
		oauthConfig: &oauth2.Config{
			ClientID:     oauthConf.ClientId,
			ClientSecret: oauthConf.ClientSecret,
			RedirectURL:  oauthConf.RedirectURL,
			Endpoint:     google.Endpoint,
		},
	}
}
