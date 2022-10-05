package client

import (
	"fmt"

	"github.com/GrabItYourself/giys-backend/lib/logger"
	"github.com/GrabItYourself/giys-backend/user/pkg/userproto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func NewClient(addr string, opts ...grpc.DialOption) (userproto.UserServiceClient, error) {
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to user service")
	}
	defer conn.Close()

	client := userproto.NewUserServiceClient(conn)
	logger.Info(fmt.Sprintf("gRPC client connected to user service at %s", addr))
	return client, nil
}
