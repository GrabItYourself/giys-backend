package server

import (
	"context"

	"github.com/GrabItYourself/giys-backend/lib/postgres"
	"github.com/GrabItYourself/giys-backend/order/pkg/orderproto"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteOrder(ctx context.Context, in *orderproto.DeleteOrderRequest) (*orderproto.OrderResponse, error) {
	var (
		shopId  = in.GetShopId()
		orderId = in.GetOrderId()
		userId  = in.GetUserId()
	)

	if err := s.repo.DeleteOrder(shopId, orderId, userId); err != nil {
		return nil, status.Errorf(postgres.InferCodeFromError(err), errors.Wrap(err, "Failed to create an order").Error())
	}

	return &orderproto.OrderResponse{}, nil
}
