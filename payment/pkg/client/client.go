package client

import (
	"github.com/GrabItYourself/giys-backend/lib/logger"
	"github.com/GrabItYourself/giys-backend/payment/pkg/paymentproto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func NewClient(addr string, opts ...grpc.DialOption) (*paymentproto.PaymentServiceClient, error) {
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		logger.Error("Failed to connect to payment service", zap.Error(err))
		return nil, err
	}
	defer conn.Close()

	client := paymentproto.NewPaymentServiceClient(conn)
	return &client, nil
}
