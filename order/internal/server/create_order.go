package server

import (
	"context"

	"github.com/GrabItYourself/giys-backend/order/libproto"
)

func (*Server) CreateOrder(ctx context.Context, in *libproto.CreateOrderRequest) (*libproto.OrderResponse, error) {
	return nil, nil
}
