package client

import (
	"github.com/GrabItYourself/giys-backend/lib/logger"
	"github.com/GrabItYourself/giys-backend/shop/pkg/shopproto"
	"google.golang.org/grpc"
)

func NewClient(addr string, opts ...grpc.DialOption) (shopproto.ShopServiceClient, error) {
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		logger.Error("Failed to connect to Shop service: " + err.Error())
		return nil, err
	}
	defer conn.Close()

	client := shopproto.NewShopServiceClient(conn)
	return client, nil
}
