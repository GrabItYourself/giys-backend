package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/GrabItYourself/giys-backend/shop/internal/libproto"
)

func (*Server) AddBankAccount(ctx context.Context, input *libproto.AddBankAccountRequest) (*libproto.AddBankAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBankAccount not implemented")
}
