package client

import (
	"context"

	"github.com/GrabItYourself/giys-backend/user/pkg/userproto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func NewClient(ctx context.Context, addr string, opts ...grpc.DialOption) (userproto.UserServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.DialContext(ctx, addr, opts...)
	if err != nil {
		return nil, nil, errors.Wrap(err, "Failed to connect to user service")
	}

	client := userproto.NewUserServiceClient(conn)
	return client, conn, nil
}
