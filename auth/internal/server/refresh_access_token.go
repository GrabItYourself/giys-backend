package server

import (
	"context"

	"github.com/GrabItYourself/giys-backend/auth/internal/repository"
	"github.com/GrabItYourself/giys-backend/auth/internal/types/accesstoken"
	"github.com/GrabItYourself/giys-backend/auth/internal/types/refreshtoken"
	"github.com/GrabItYourself/giys-backend/auth/internal/types/tokenmapping"
	"github.com/GrabItYourself/giys-backend/auth/pkg/authproto"
	"github.com/GrabItYourself/giys-backend/lib/logger"
	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
	"github.com/GrabItYourself/giys-backend/lib/redis"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

func (s *Server) RefreshAccessToken(ctx context.Context, in *authproto.RefreshAccessTokenReq) (*authproto.RefreshAccessTokenResp, error) {
	refreshTokenKey := &repository.RefreshTokenKey{Token: in.RefreshToken}
	refreshToken, err := s.repo.GetRefreshToken(ctx, refreshTokenKey)
	if err != nil {
		return nil, status.Errorf(redis.InferCodeFromError(err), err.Error())
	}
	accessToken, newRefreshToken, err := s.issueNewTokenPair(ctx, refreshToken.UserId, refreshToken.Role)
	if err != nil {
		// all errors are from redis, so we can infer code
		return nil, status.Errorf(redis.InferCodeFromError(err), err.Error())
	}
	return &authproto.RefreshAccessTokenResp{
		AccessToken:  accessToken.Token,
		RefreshToken: newRefreshToken.Token,
	}, nil
}

func (s *Server) issueNewTokenPair(ctx context.Context, userId string, role models.RoleEnum) (*accesstoken.AccessToken, *refreshtoken.RefreshToken, error) {
	// Retrieve user's token mapping
	tokenMapping, err := s.repo.GetTokenMapping(ctx, &repository.TokenMappingKey{UserId: userId})
	if err == nil {
		// Revoke old tokens
		err = s.repo.DeleteAccessToken(ctx, &repository.AccessTokenKey{Token: tokenMapping.AccessToken})
		if err != nil {
			logger.Warn(errors.Wrap(err, "failed to delete old access token").Error())
		}
		err = s.repo.DeleteRefreshToken(ctx, &repository.RefreshTokenKey{Token: tokenMapping.RefreshToken})
		if err != nil {
			logger.Warn(errors.Wrap(err, "failed to delete old refresh token").Error())
		}
	}

	// Issue new tokens with random UUIDs
	accessToken := accesstoken.New(userId, role)
	err = s.repo.SetAccessToken(ctx, &repository.AccessTokenKey{Token: accessToken.Token}, accessToken)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error during setting access token")
	}
	refreshToken := refreshtoken.New(userId, role)
	err = s.repo.SetRefreshToken(ctx, &repository.RefreshTokenKey{Token: refreshToken.Token}, refreshToken)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error during setting refresh token")
	}

	// Update user's token mapping
	err = s.repo.SetTokenMapping(ctx, &repository.TokenMappingKey{UserId: userId}, &tokenmapping.TokenMapping{
		AccessToken:  accessToken.Token,
		RefreshToken: refreshToken.Token,
	})
	if err != nil {
		return nil, nil, errors.Wrap(err, "error during setting token mapping")
	}

	return accessToken, refreshToken, nil
}
