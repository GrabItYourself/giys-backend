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

func (s *Server) CreateOrder(ctx context.Context, in *orderproto.CreateOrderRequest) (*orderproto.OrderResponse, error) {
	var (
		shopId = in.GetShopId()
		items  = in.GetItems()
	)

	identity, err := authutils.ExtractIdentityFromGrpcContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, errors.Wrap(err, "can't extract user from context").Error())
	}

	orderItems := make([]models.OrderItem, len(items))
	for index, item := range items {
		orderItems[index] = models.OrderItem{
			ShopId:     shopId,
			ShopItemId: item.ShopItemId,
			Quantity:   item.Quantity,
		}
	}

	order := models.Order{
		UserId: identity.UserId,
		ShopId: shopId,
		Status: models.InQueueStatus,
		Items:  orderItems,
	}

	if err := s.repo.CreateOrder(&order); err != nil {
		return nil, status.Errorf(postgres.InferCodeFromError(err), errors.Wrap(err, "Failed to create an order").Error())
	}

	return s.toProtoOrderResponse(&order), nil
}
