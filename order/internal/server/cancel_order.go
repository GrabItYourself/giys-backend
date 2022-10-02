package server

import (
	"context"

	"github.com/GrabItYourself/giys-backend/order/libproto"
)

func (*Server) CancelOrder(ctx context.Context, in *libproto.CancelOrderRequest) (*libproto.OrderResponse, error) {
	return nil, nil
}
