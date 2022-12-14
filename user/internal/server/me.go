package server

import (
	"context"

	"github.com/GrabItYourself/giys-backend/auth/pkg/authutils"
	"github.com/GrabItYourself/giys-backend/lib/postgres"
	"github.com/GrabItYourself/giys-backend/user/pkg/userproto"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Me(ctx context.Context, in *userproto.MeReq) (*userproto.MeResp, error) {
	identity, err := authutils.ExtractIdentityFromGrpcContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, errors.Wrap(err, "can't extract user from context").Error())
	}
	user, err := s.repo.GetUserById(identity.UserId)
	if err != nil {
		return nil, status.Error(postgres.InferCodeFromError(err), errors.Wrap(err, "can't get user").Error())
	}
	shops, err := s.repo.GetOwnedShops(user.Id)
	if err != nil {
		return nil, status.Error(postgres.InferCodeFromError(err), errors.Wrap(err, "can't get owned shops").Error())
	}
	userProto := &userproto.User{
		Id:       user.Id,
		Role:     string(user.Role),
		Email:    user.Email,
		GoogleId: user.GoogleId,
	}
	if len(shops) > 0 {
		userProto.ShopId = &shops[0].Id
	}
	resp := &userproto.MeResp{
		User: userProto,
	}
	return resp, nil
}
