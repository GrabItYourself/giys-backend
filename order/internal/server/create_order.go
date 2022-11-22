package server

import (
	"context"

	"github.com/GrabItYourself/giys-backend/auth/pkg/authutils"
	"github.com/GrabItYourself/giys-backend/lib/logger"
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

	shop, err := s.repo.GetShopById(shopId)
	if err != nil {
		return nil, status.Errorf(postgres.InferCodeFromError(err), errors.Wrap(err, "Failed to get shop").Error())
	}

	for _, owner := range shop.Owners {
		emailMessage := s.toOrderEmailMessage(owner.User.Email, shop.Name, &order)
		err := s.rabbitSender.SendMessage(ctx, "email", emailMessage)
		if err != nil {
			logger.Error(errors.Wrap(err, "Failed to send email message").Error())
		}
	}

	user, err := s.repo.GetUserById(identity.UserId)
	if err != nil {
		return nil, status.Errorf(postgres.InferCodeFromError(err), errors.Wrap(err, "Failed to get user").Error())
	}
	emailMessage := s.toOrderEmailMessage(user.Email, shop.Name, &order)
	err = s.rabbitSender.SendMessage(ctx, "email", emailMessage)
	if err != nil {
		logger.Error(errors.Wrap(err, "Failed to send email message").Error())
	}

	return s.toProtoOrderResponse(&order), nil
}
