package server

import (
	"context"

	"github.com/GrabItYourself/giys-backend/lib/authutils"
	"github.com/GrabItYourself/giys-backend/lib/postgres"
	"github.com/GrabItYourself/giys-backend/user/pkg/libproto"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Me(ctx context.Context, in *libproto.MeReq) (*libproto.MeResp, error) {
	userId, _, err := authutils.ExtractUserFromGrpcContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, errors.Wrap(err, "can't extract user from context").Error())
	}
	user, err := s.repo.GetUserById(userId)
	if err != nil {
		return nil, status.Error(postgres.InferCodeFromError(err), errors.Wrap(err, "can't get user").Error())
	}
	resp := &libproto.MeResp{
		User: &libproto.User{
			Id:       user.Id,
			Role:     string(user.Role),
			Email:    user.Email,
			GoogleId: user.GoogleId,
		},
	}
	return resp, nil
}
