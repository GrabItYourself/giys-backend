package server

import (
	"context"

	"github.com/GrabItYourself/giys-backend/lib/postgres"
	"github.com/GrabItYourself/giys-backend/user/pkg/userproto"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

func (s *Server) GetUserByGoogleId(ctx context.Context, in *userproto.GetUserByGoogleIdReq) (*userproto.GetUserByGoogleIdResp, error) {
	user, err := s.repo.GetUserByGoogleId(in.GoogleId)
	if err != nil {
		return nil, status.Error(postgres.InferCodeFromError(err), errors.Wrap(err, "can't get user").Error())
	}
	resp := &userproto.GetUserByGoogleIdResp{
		User: &userproto.User{
			Id:       user.Id,
			Role:     string(user.Role),
			Email:    user.Email,
			GoogleId: user.GoogleId,
		},
	}
	return resp, nil
}
