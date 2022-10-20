package orderclient

import (
	"context"

	"github.com/GrabItYourself/giys-backend/order/pkg/orderproto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func NewClient(ctx context.Context, addr string, opts ...grpc.DialOption) (orderproto.OrderClient, *grpc.ClientConn, error) {
	conn, err := grpc.DialContext(ctx, addr, opts...)
	if err != nil {
		return nil, nil, errors.Wrap(err, "Failed to create Order GRPC client")
	}

	client := orderproto.NewOrderClient(conn)
	return client, conn, nil
}
