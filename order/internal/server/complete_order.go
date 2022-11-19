package server

import (
	"context"

	"github.com/GrabItYourself/giys-backend/lib/postgres"
	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
	"github.com/GrabItYourself/giys-backend/order/pkg/orderproto"
	"github.com/GrabItYourself/giys-backend/payment/pkg/paymentproto"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CompleteOrder(ctx context.Context, in *orderproto.CompleteOrderRequest) (*orderproto.OrderResponse, error) {
	var (
		orderId = in.GetOrderId()
		shopId  = in.GetShopId()
	)

	_, err := s.paymentClient.Pay(ctx, &paymentproto.PayRequest{
		ShopId:  shopId,
		OrderId: orderId,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, errors.Wrap(err, "Failed to request GRPC Pay").Error())
	}

	order, err := s.repo.UpdateOrderStatus(orderId, shopId, models.CompletedStatus)
	if err != nil {
		return nil, status.Errorf(postgres.InferCodeFromError(err), errors.Wrap(err, "Failed to cancel an order").Error())
	}

	return s.toProtoOrderResponse(order), nil
}
