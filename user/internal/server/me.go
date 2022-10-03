package server

import (
	"context"

	"github.com/GrabItYourself/giys-backend/lib/authutils"
	"github.com/GrabItYourself/giys-backend/user/internal/libproto"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (s *Server) Me(ctx context.Context, in *libproto.MeReq) (*libproto.MeResp, error) {
	userId, _, err := authutils.ExtractUserFromGrpcContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, errors.Wrap(err, "can't extract user from context").Error())
	}
	user, err := s.repo.GetUserById(userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, errors.Wrap(err, "can't get user").Error())
		}
		return nil, status.Error(codes.Internal, errors.Wrap(err, "can't get user").Error())
	}
	resp := &libproto.MeResp{
		User: libproto.ConvertUserToProto(user),
	}
	return resp, nil
}
