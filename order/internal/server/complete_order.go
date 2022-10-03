package server

import (
	"context"

	"github.com/GrabItYourself/giys-backend/order/internal/libproto"
)

func (*Server) CompleteOrder(ctx context.Context, in *libproto.CompleteOrderRequest) (*libproto.OrderResponse, error) {
	return nil, nil
}
