package server

import (
	"context"

	"github.com/GrabItYourself/giys-backend/order/internal/libproto"
)

func (*Server) CompleteOrder(ctx context.Context, in *libproto.CompleteOrderRequest) (*libproto.OrderResponse, error) {
	// _, _, err := authutils.ExtractUserFromGrpcContext(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	// return nil, nil
	return nil, nil
}
