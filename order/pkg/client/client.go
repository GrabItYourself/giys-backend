package client

import (
	"github.com/GrabItYourself/giys-backend/lib/logger"
	"github.com/GrabItYourself/giys-backend/order/libproto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func NewClient(addr string, opts ...grpc.DialOption) (*libproto.OrderClient, error) {
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		logger.Error("Failed to connect to order service", zap.Error(err))
		return nil, err
	}
	defer conn.Close()

	client := libproto.NewOrderClient(conn)
	return &client, nil
}
