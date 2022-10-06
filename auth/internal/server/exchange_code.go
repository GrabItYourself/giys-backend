package server

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/GrabItYourself/giys-backend/auth/pkg/authproto"
	"github.com/GrabItYourself/giys-backend/lib/logger"
	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
	"github.com/GrabItYourself/giys-backend/user/pkg/userproto"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

type UserInfo struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

func (s *Server) ExchangeAuthCode(ctx context.Context, in *authproto.ExchangeAuthCodeReq) (*authproto.ExchangeAuthCodeResp, error) {
	token, err := s.oauthConfig.Exchange(ctx, in.AuthCode)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, errors.Wrap(err, "error during code exchange").Error())
	}

	userInfo, err := getUserInfoFromGoogle(token)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var user *userproto.User
	getUserResp, err := s.userClient.GetUserByGoogleId(ctx, &userproto.GetUserByGoogleIdReq{GoogleId: userInfo.Id})
	if err != nil {
		if status.Code(err) != codes.NotFound {
			return nil, status.Error(codes.Internal, errors.Wrap(err, "can't get user by google id").Error())
		}
		logger.Info("user not found, creating new one", zap.String("google_id", userInfo.Id), zap.String("email", userInfo.Email))
		createUserResp, err := s.userClient.CreateUser(ctx, &userproto.CreateUserReq{
			Email:    userInfo.Email,
			GoogleId: userInfo.Id,
		})
		user = createUserResp.User
		if err != nil {
			return nil, status.Error(codes.Internal, errors.Wrap(err, "error during creating user").Error())
		}
	} else {
		logger.Info("user found", zap.String("google_id", userInfo.Id), zap.String("email", userInfo.Email))
		user = getUserResp.User
	}

	accessToken, refreshToken, err := s.issueNewTokenPair(ctx, user.Id, models.RoleEnum(user.Role))
	if err != nil {
		return nil, status.Error(codes.Internal, errors.Wrap(err, "failed to issue new token pair").Error())
	}

	return &authproto.ExchangeAuthCodeResp{
		AccessToken:  accessToken.Token,
		RefreshToken: refreshToken.Token,
	}, nil
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
