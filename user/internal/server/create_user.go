package server

import (
	"context"

	"github.com/GrabItYourself/giys-backend/user/internal/libproto"
)

func (*Server) CreateUser(ctx context.Context, in *libproto.CreateUserReq) (*libproto.CreateUserResp, error) {
	return nil, nil
}
