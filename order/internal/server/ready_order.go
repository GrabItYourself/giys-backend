package server

import (
	"context"

	"github.com/GrabItYourself/giys-backend/order/libproto"
)

func (*Server) ReadyOrder(ctx context.Context, in *libproto.ReadyOrderRequest) (*libproto.OrderResponse, error) {
	return nil, nil
}
