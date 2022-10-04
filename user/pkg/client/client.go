package client

import (
	"github.com/GrabItYourself/giys-backend/user/libproto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func NewClient(addr string, opts ...grpc.DialOption) (libproto.UserServiceClient, error) {
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to user service")
	}
	defer conn.Close()

	client := libproto.NewUserServiceClient(conn)
	return client, nil
}
