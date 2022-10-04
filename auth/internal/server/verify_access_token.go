package server

import (
	"context"

	"github.com/GrabItYourself/giys-backend/auth/internal/libproto"
	"github.com/GrabItYourself/giys-backend/auth/internal/repository"
	"github.com/GrabItYourself/giys-backend/lib/redis"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

func (s *Server) VerifyAccessToken(ctx context.Context, in *libproto.VerifyAccessTokenReq) (*libproto.VerifyAccessTokenResp, error) {
	accessTokenKey := &repository.AccessTokenKey{Token: in.AccessToken}
	accessToken, err := s.repo.GetAccessToken(ctx, accessTokenKey)
	if err != nil {
		return nil, status.Errorf(redis.InferCodeFromError(err), errors.Wrap(err, "error during getting access token").Error())
	}
	return &libproto.VerifyAccessTokenResp{
		UserId: accessToken.UserId,
		Role:   string(accessToken.Role),
	}, nil
}
