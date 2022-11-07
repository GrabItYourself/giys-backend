package client

import (
	"context"

	"github.com/GrabItYourself/giys-backend/payment/pkg/paymentproto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func NewClient(ctx context.Context, addr string, opts ...grpc.DialOption) (paymentproto.PaymentServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.DialContext(ctx, addr, opts...)
	if err != nil {
		return nil, nil, errors.Wrap(err, "Failed to connect to payment service")
	}

	client := paymentproto.NewPaymentServiceClient(conn)
	return client, conn, nil
}
