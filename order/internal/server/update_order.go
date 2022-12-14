package server

import (
	"context"

	"github.com/GrabItYourself/giys-backend/auth/pkg/authutils"
	"github.com/GrabItYourself/giys-backend/lib/postgres"
	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
	"github.com/GrabItYourself/giys-backend/order/pkg/orderproto"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateOrder(ctx context.Context, in *orderproto.UpdateOrderRequest) (*orderproto.OrderResponse, error) {
	var (
		orderId = in.GetOrderId()
		shopId  = in.GetShopId()
		items   = in.GetItems()
	)

	identity, err := authutils.ExtractIdentityFromGrpcContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, errors.Wrap(err, "can't extract user from context").Error())
	}

	orderItems := make([]models.OrderItem, len(items))
	for index, item := range items {
		orderItems[index] = models.OrderItem{
			OrderId:    orderId,
			ShopId:     shopId,
			ShopItemId: item.ShopItemId,
			Quantity:   item.Quantity,
		}
	}

	order := models.Order{
		Id:     orderId,
		UserId: identity.UserId,
		ShopId: shopId,
		Status: models.InQueueStatus,
		Items:  orderItems,
	}

	if err := s.repo.UpdateOrder(&order); err != nil {
		return nil, status.Errorf(postgres.InferCodeFromError(err), errors.Wrap(err, "Failed to save an order").Error())
	}

	return s.toProtoOrderResponse(&order), nil
}
