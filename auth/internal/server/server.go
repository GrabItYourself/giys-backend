package server

import (
	"github.com/GrabItYourself/giys-backend/auth/internal/config"
	"github.com/GrabItYourself/giys-backend/auth/internal/repository"
	"github.com/GrabItYourself/giys-backend/auth/pkg/authproto"
	"github.com/GrabItYourself/giys-backend/user/pkg/userproto"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Server struct {
	authproto.UnimplementedAuthServer
	repo         *repository.Repository
	oauthIOS     *oauth2.Config
	oauthAndroid *oauth2.Config
	userClient   userproto.UserServiceClient
}

func NewServer(repo *repository.Repository, conf *config.Config, userClient userproto.UserServiceClient) *Server {
	return &Server{
		repo: repo,
		oauthIOS: &oauth2.Config{
			ClientID:    conf.OAuth.IOS.ClientId,
			RedirectURL: conf.OAuth.IOS.RedirectURL,
			Endpoint:    google.Endpoint,
		},
		oauthAndroid: &oauth2.Config{
			ClientID:    conf.OAuth.Android.ClientId,
			RedirectURL: conf.OAuth.Android.RedirectURL,
			Endpoint:    google.Endpoint,
		},
		userClient: userClient,
	}
}
