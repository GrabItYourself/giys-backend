package server

import (
	"context"

	"github.com/GrabItYourself/giys-backend/lib/logger"
	"github.com/GrabItYourself/giys-backend/lib/postgres"
	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
	"github.com/GrabItYourself/giys-backend/order/pkg/orderproto"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

func (s *Server) ReadyOrder(ctx context.Context, in *orderproto.ReadyOrderRequest) (*orderproto.OrderResponse, error) {
	var (
		orderId = in.GetOrderId()
		shopId  = in.GetShopId()
	)

	order, err := s.repo.UpdateOrderStatus(orderId, shopId, models.ReadyStatus)
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

	emailMessage := s.toOrderEmailMessage(user.Email, shop.Name, order)
	err = s.rabbitSender.SendMessage(ctx, "email", emailMessage)
	if err != nil {
		logger.Error(errors.Wrap(err, "Failed to send email message").Error())
	}

	return s.toProtoOrderResponse(order), nil
}
