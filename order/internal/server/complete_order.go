package server

import (
	"context"

	"github.com/GrabItYourself/giys-backend/lib/postgres"
	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
	"github.com/GrabItYourself/giys-backend/order/pkg/orderproto"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

func (s *Server) CompleteOrder(ctx context.Context, in *orderproto.CompleteOrderRequest) (*orderproto.OrderResponse, error) {
	var (
		orderId = in.GetOrderId()
		shopId  = in.GetShopId()
	)

	order := models.Order{}
	if err := s.pg.Model(&order).Where("id = ? AND shop_id = ?", orderId, shopId).Update("status", models.CompletedStatus).Error; err != nil {
		return nil, status.Errorf(postgres.InferCodeFromError(err), errors.Wrap(err, "Failed to cancel an order").Error())
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
