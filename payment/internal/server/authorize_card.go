package server

import (
	"context"

	"github.com/GrabItYourself/giys-backend/payment/internal/libproto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) AuthorizeCard(context.Context, *libproto.AuthorizeCardRequest) (*libproto.AuthorizeCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthorizeCard not implemented")
}
