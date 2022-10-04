package server

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/GrabItYourself/giys-backend/auth/internal/libproto"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

type UserInfo struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

func (s *Server) ExchangeAuthCode(ctx context.Context, in *libproto.ExchangeAuthCodeReq) (*libproto.ExchangeAuthCodeResp, error) {
	token, err := s.oauthConfig.Exchange(ctx, in.AuthCode)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, errors.Wrap(err, "error during code exchange").Error())
	}

	userInfo, err := getUserInfoFromGoogle(token)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Print(userInfo)
	// TODO: wait for discussion about shared repository vs separate repositories

	return nil, nil
}

func getUserInfoFromGoogle(token *oauth2.Token) (*UserInfo, error) {
	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, errors.Wrap(err, "error during getting user info")
	}
	defer response.Body.Close()
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read userinfo response body")
	}
	var userInfo UserInfo
	err = json.Unmarshal(contents, &userInfo)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal userinfo response body")
	}
	return &userInfo, nil
}
