package server

import (
	"context"

	"github.com/GrabItYourself/giys-backend/order/internal/libproto"
)

func (*Server) CancelOrder(ctx context.Context, in *libproto.CancelOrderRequest) (*libproto.OrderResponse, error) {
	return nil, nil
}
