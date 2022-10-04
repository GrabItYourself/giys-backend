package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/GrabItYourself/giys-backend/shop/pkg/shopproto"
)

func (*Server) AddBankAccount(ctx context.Context, input *shopproto.AddBankAccountRequest) (*shopproto.AddBankAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBankAccount not implemented")
}
