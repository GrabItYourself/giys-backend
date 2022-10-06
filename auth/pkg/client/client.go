package client

import (
	"context"

	"github.com/GrabItYourself/giys-backend/auth/pkg/authproto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func NewClient(ctx context.Context, addr string, opts ...grpc.DialOption) (authproto.AuthClient, *grpc.ClientConn, error) {
	conn, err := grpc.DialContext(ctx, addr, opts...)
	if err != nil {
		return nil, nil, errors.Wrap(err, "Failed to connect to auth service")
	}

	client := authproto.NewAuthClient(conn)
	return client, conn, nil
}
