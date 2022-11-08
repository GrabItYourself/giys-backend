package server

import (
	"context"

	"github.com/GrabItYourself/giys-backend/auth/pkg/authutils"
	"github.com/GrabItYourself/giys-backend/lib/postgres"
	"github.com/GrabItYourself/giys-backend/order/pkg/orderproto"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetOrder(ctx context.Context, in *orderproto.GetOrderRequest) (*orderproto.OrderResponse, error) {
	var (
		orderId = in.GetOrderId()
		shopId  = in.GetShopId()
	)
	identity, err := authutils.ExtractIdentityFromGrpcContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, errors.Wrap(err, "Can't extract user from context").Error())
	}
	order, err := s.repo.GetOrderById(orderId, shopId, identity.UserId)
	if err != nil {
		return nil, status.Errorf(postgres.InferCodeFromError(err), errors.Wrap(err, "Failed to get an order").Error())
	}
	return s.toProtoOrderResponse(order), nil
}
