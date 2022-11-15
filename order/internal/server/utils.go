package server

import (
	"fmt"

	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
	"github.com/GrabItYourself/giys-backend/lib/rabbitmq/types"
	"github.com/GrabItYourself/giys-backend/order/pkg/orderproto"
)

func (s *Server) toProtoOrderResponse(order *models.Order) *orderproto.OrderResponse {
	orderResponseItems := make([]*orderproto.OrderItem, len(order.Items))
	for index, item := range order.Items {
		orderResponseItems[index] = &orderproto.OrderItem{
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

func (s *Server) toProtoOrderListResponse(order []models.Order) []*orderproto.OrderResponse {
	orderResponses := make([]*orderproto.OrderResponse, len(order))
	for index, item := range order {
		orderResponses[index] = s.toProtoOrderResponse(&item)
	}
	return orderResponses
}

func (s *Server) toOrderEmailMessage(toEmail string, shopName string, order *models.Order) *types.EmailMessage {
	var subject, body string

	switch order.Status {
	case models.InQueueStatus:
		subject = "GIYS: New order has been created"
		body = fmt.Sprintf("New order(Order ID %d) has been created.", order.Id)
	case models.ReadyStatus:
		subject = "GIYS: The order is ready"
		body = fmt.Sprintf("The order(Order ID %d) from Shop %s is now ready for picking up.", order.Id, shopName)
	case models.CanceledStatus:
		subject = "GIYS: The order has been canceled"
		body = fmt.Sprintf("The order(Order ID %d) from Shop %s has been canceled.", order.Id, shopName)
	}

	return &types.EmailMessage{
		To:      toEmail,
		Subject: subject,
		Body:    body,
	}
}
