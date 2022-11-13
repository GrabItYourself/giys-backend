package server

import (
	"context"

	"github.com/GrabItYourself/giys-backend/lib/postgres"
	"github.com/GrabItYourself/giys-backend/order/pkg/orderproto"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

func (s *Server) GetShopOrders(ctx context.Context, in *orderproto.GetShopOrdersRequest) (*orderproto.OrderListResponse, error) {
	var (
		shopId = in.GetShopId()
	)
	order, err := s.repo.GetShopOrders(shopId)
	if err != nil {
		return nil, status.Errorf(postgres.InferCodeFromError(err), errors.Wrap(err, "Failed to shop orders").Error())
	}
	return &orderproto.OrderListResponse{
		Result: s.toProtoOrderListResponse(order),
	}, nil
}
