package server

import (
	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
	"github.com/GrabItYourself/giys-backend/order/pkg/orderproto"
)

func (s *Server) toProtoOrderResponse(order *models.Order) *orderproto.OrderResponse {
	orderResponseItems := make([]*orderproto.OrderItem, len(order.Items))
	for index, item := range order.Items {
		orderResponseItems[index] = &orderproto.OrderItem{
			ShopId:     item.ShopId,
			ShopItemId: item.ShopItemId,
			Quantity:   item.Quantity,
		}
	}

	return &orderproto.OrderResponse{
		OrderId: order.Id,
		UserId:  order.UserId,
		ShopId:  order.ShopId,
		Status:  string(order.Status),
		Items:   orderResponseItems,
	}
}
