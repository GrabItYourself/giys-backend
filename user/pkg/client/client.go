package client

import (
	"context"

	"github.com/GrabItYourself/giys-backend/user/pkg/libproto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func NewClient(ctx context.Context, addr string, opts ...grpc.DialOption) (libproto.UserServiceClient, error) {
	conn, err := grpc.DialContext(ctx, addr, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to user service")
	}
	defer conn.Close()

	client := libproto.NewUserServiceClient(conn)
	return client, nil
}
