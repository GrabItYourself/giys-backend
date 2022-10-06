package server

import (
	"context"

	"github.com/GrabItYourself/giys-backend/lib/postgres"
	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
	"github.com/GrabItYourself/giys-backend/order/pkg/orderproto"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

func (s *Server) ReadyOrder(ctx context.Context, in *orderproto.ReadyOrderRequest) (*orderproto.OrderResponse, error) {
	var (
		orderId = int(in.GetOrderId())
		shopId  = int(in.GetShopId())
	)

	order, err := s.repo.UpdateOrderStatus(orderId, shopId, models.ReadyStatus)
	if err != nil {
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
