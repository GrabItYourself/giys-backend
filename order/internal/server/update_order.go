package server

import (
	"context"

	"github.com/GrabItYourself/giys-backend/order/internal/libproto"
)

func (*Server) UpdateOrder(ctx context.Context, in *libproto.UpdateOrderRequest) (*libproto.OrderResponse, error) {
	return nil, nil
}
