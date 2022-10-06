package server

import (
	"context"

	"github.com/GrabItYourself/giys-backend/lib/postgres"
	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
	"github.com/GrabItYourself/giys-backend/order/pkg/orderproto"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateOrder(ctx context.Context, in *orderproto.CreateOrderRequest) (*orderproto.OrderResponse, error) {
	var (
		shopId = in.GetShopId()
		userId = in.GetUserId()
		items  = in.GetItems()
	)

	orderItems := make([]models.OrderItem, len(items))
	for index, item := range items {
		orderItems[index] = models.OrderItem{
			ShopId:     int(item.ShopId),
			ShopItemId: int(item.ShopItemId),
			Quantity:   int(item.Quantity),
		}
	}

	order := models.Order{
		UserId: userId,
		ShopId: int(shopId),
		Status: models.InQueueStatus,
		Items:  orderItems,
	}

	if err := s.pg.Create(&order).Error; err != nil {
		return nil, status.Errorf(postgres.InferCodeFromError(err), errors.Wrap(err, "Failed to create an order").Error())
	}

	orderResponseItems := make([]*orderproto.OrderItem, len(order.Items))
	for index, item := range order.Items {
		orderResponseItems[index] = &orderproto.OrderItem{
			ShopId:     int32(item.ShopId),
			ShopItemId: int32(item.ShopItemId),
			Quantity:   int32(item.Quantity),
		}
	}

	return &orderproto.OrderResponse{
		OrderId: int32(order.Id),
		UserId:  order.UserId,
		ShopId:  int32(order.ShopId),
		Status:  string(order.Status),
		Items:   orderResponseItems,
	}, nil
}
