package server

import (
	"context"

	"github.com/GrabItYourself/giys-backend/lib/postgres"
	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
	"github.com/GrabItYourself/giys-backend/order/pkg/orderproto"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

func (s *Server) CancelOrder(ctx context.Context, in *orderproto.CancelOrderRequest) (*orderproto.OrderResponse, error) {
	var (
		orderId = in.GetOrderId()
		shopId  = in.GetShopId()
	)

	order, err := s.repo.UpdateOrderStatus(orderId, shopId, models.CanceledStatus)
	if err != nil {
		return nil, status.Errorf(postgres.InferCodeFromError(err), errors.Wrap(err, "Failed to cancel an order").Error())
	}

	user, err := s.repo.GetUserById(order.UserId)
	if err != nil {
		return nil, status.Errorf(postgres.InferCodeFromError(err), errors.Wrap(err, "can't get user").Error())
	}

	shop, err := s.repo.GetShopById(shopId)
	if err != nil {
		return nil, status.Errorf(postgres.InferCodeFromError(err), errors.Wrap(err, "Failed to get shop").Error())
	}

	// Send email to customer
	emailMessage := s.toCancelOrderEmailMessage(user.Email, shop.Name, orderId)
	s.rabbitSender.SendMessage(ctx, "email", emailMessage)

	// Send email to shop owners
	for _, owner := range shop.Owners {
		emailMessage := s.toCancelOrderEmailMessage(owner.User.Email, shop.Name, orderId)
		s.rabbitSender.SendMessage(ctx, "email", emailMessage)
	}

	return s.toProtoOrderResponse(order), nil
}
