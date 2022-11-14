package server

import (
	"fmt"

	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
	"github.com/GrabItYourself/giys-backend/lib/rabbitmq/types"
)

func (s *Server) createPaymentEmailMessage(toEmail string, shopName string, totalAmountTHB int, order *models.Order) *types.EmailMessage {
	orderDetails := "Order Details -\n"
	for _, orderItem := range order.Items {
		price := int64(orderItem.Quantity) * int64(orderItem.ShopItem.Price)
		orderDetails += fmt.Sprintf("x%d\t%s\t฿%d\n", orderItem.Quantity, orderItem.ShopItem.Name, price)
	}
	body := fmt.Sprintf("Shop - %s\nTotal Amount - ฿%d\n%s", shopName, totalAmountTHB, orderDetails)

	return &types.EmailMessage{
		To:      toEmail,
		Subject: "GIYS: Payment has been made",
		Body:    body,
	}
}
