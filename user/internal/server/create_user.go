package server

import (
	"context"

	"github.com/GrabItYourself/giys-backend/lib/postgres"
	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
	"github.com/GrabItYourself/giys-backend/user/pkg/userproto"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateUser(ctx context.Context, in *userproto.CreateUserReq) (*userproto.CreateUserResp, error) {
	user := &models.User{
		Email:    in.Email,
		GoogleId: in.GoogleId,
	}
	err := s.repo.CreateUser(user)
	if err != nil {
		return nil, status.Error(postgres.InferCodeFromError(err), errors.Wrap(err, "can't create user").Error())
	}
	return &userproto.CreateUserResp{
		User: &userproto.User{
			Id:       user.Id,
			Role:     string(user.Role),
			Email:    user.Email,
			GoogleId: user.GoogleId,
		},
	}, nil
}
