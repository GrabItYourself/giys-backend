package shopclient

import (
	"context"

	"github.com/GrabItYourself/giys-backend/shop/pkg/shopproto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func NewClient(ctx context.Context, addr string, opts ...grpc.DialOption) (shopproto.ShopServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.DialContext(ctx, addr, opts...)
	if err != nil {
		return nil, nil, errors.Wrap(err, "Failed to connect to shop service")
	}

	client := shopproto.NewShopServiceClient(conn)
	return client, conn, nil
}
