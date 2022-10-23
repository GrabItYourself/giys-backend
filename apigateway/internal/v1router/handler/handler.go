package v1handler

import (
	"github.com/GrabItYourself/giys-backend/apigateway/internal/config"
	"github.com/GrabItYourself/giys-backend/auth/pkg/authproto"
	"github.com/GrabItYourself/giys-backend/shop/pkg/shopproto"
	"github.com/GrabItYourself/giys-backend/user/pkg/userproto"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Handler struct {
	Grpc        *GrpcClients
	OAuthConfig *oauth2.Config
}

type GrpcClients struct {
	User userproto.UserServiceClient
	Auth authproto.AuthClient
	Shop shopproto.ShopServiceClient
}

func NewHandler(
	grpc *GrpcClients,
	oauthConf *config.OAuthConfig,
) *Handler {
	return &Handler{
		Grpc: grpc,
		OAuthConfig: &oauth2.Config{
			ClientID:     oauthConf.ClientId,
			ClientSecret: oauthConf.ClientSecret,
			RedirectURL:  oauthConf.RedirectURL,
			Endpoint:     google.Endpoint,
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "openid"},
		},
	}
}
