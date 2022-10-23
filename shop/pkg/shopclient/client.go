package shopclient

import (
	"context"

	"github.com/GrabItYourself/giys-backend/lib/logger"
	"github.com/GrabItYourself/giys-backend/shop/pkg/shopproto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func NewClient(ctx context.Context, addr string, opts ...grpc.DialOption) (shopproto.ShopServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.DialContext(ctx, addr, opts...)
	if err != nil {
		logger.Error("Failed to create Shop GRPC client: " + err.Error())
		return nil, nil, errors.Wrap(err, "Failed to create Shop GRPC client")
	}
	defer conn.Close()

	client := shopproto.NewShopServiceClient(conn)
	return client, conn, nil
}
