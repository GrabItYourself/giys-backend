package server

import (
	"context"
	"io"
	"net/http"

	"github.com/GrabItYourself/giys-backend/auth/internal/libproto"
	"github.com/GrabItYourself/giys-backend/lib/logger"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func (s *Server) ExchangeAuthCode(ctx context.Context, in *libproto.ExchangeAuthCodeReq) (*libproto.ExchangeAuthCodeResp, error) {
	token, err := s.oauthConfig.Exchange(ctx, in.AuthCode)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, errors.Wrap(err, "error during code exchange").Error())
	}
	content, err := getUserInfoFromGoogle(token)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	logger.Info(content)
	return nil, nil
}

func getUserInfoFromGoogle(token *oauth2.Token) (string, error) {
	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return "", errors.Wrap(err, "error during getting user info")
	}
	defer response.Body.Close()
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return "", errors.Wrap(err, "failed to read userinfo response body")
	}
	return string(contents), nil
}
