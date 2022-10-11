package server

import (
	"context"

	"github.com/GrabItYourself/giys-backend/lib/postgres"
	"github.com/GrabItYourself/giys-backend/order/pkg/orderproto"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

func (s *Server) GetOrder(ctx context.Context, in *orderproto.GetOrderRequest) (*orderproto.OrderResponse, error) {
	var (
		orderId = in.GetOrderId()
		shopId  = in.GetShopId()
	)
	order, err := s.repo.GetOrderById(orderId, shopId)
	if err != nil {
		return nil, status.Errorf(postgres.InferCodeFromError(err), errors.Wrap(err, "Failed to create an order").Error())
	}
	return s.toProtoOrderResponse(order), nil
}
