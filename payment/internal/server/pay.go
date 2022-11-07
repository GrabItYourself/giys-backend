package server

import (
	"context"

	"github.com/GrabItYourself/giys-backend/lib/postgres"
	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
	"github.com/GrabItYourself/giys-backend/payment/pkg/paymentproto"
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Pay(ctx context.Context, in *paymentproto.PayRequest) (*paymentproto.PayResponse, error) {
	var totalAmountTHB int64 = 0
	order, err := s.repo.GetOrderById(in.OrderId, in.ShopId)
	if err != nil {
		return nil, status.Error(postgres.InferCodeFromError(err), errors.Wrap(err, "can't get order").Error())
	}

	for _, orderItem := range order.Items {
		totalAmountTHB += int64(orderItem.Quantity) * int64(orderItem.ShopItem.Price)
	}

	totalAmount := totalAmountTHB * 100

	user, err := s.repo.GetUserById(order.UserId)
	if err != nil {
		return nil, status.Error(postgres.InferCodeFromError(err), errors.Wrap(err, "can't get user").Error())
	}

	if user.DefaultPaymentMethodId == nil {
		return nil, status.Error(codes.FailedPrecondition, errors.Wrap(err, "no default payment method").Error())
	}

	paymentMethod, err := s.repo.GetPaymentMethodById(*user.DefaultPaymentMethodId)
	if err != nil {
		return nil, status.Error(postgres.InferCodeFromError(err), errors.Wrap(err, "can't get payment method").Error())
	}

	shop, err := s.repo.GetShopById(in.ShopId)
	if err != nil {
		return nil, status.Error(postgres.InferCodeFromError(err), errors.Wrap(err, "can't get shop").Error())
	}

	if shop.OmiseResipientId == nil {
		return nil, status.Error(codes.FailedPrecondition, errors.Wrap(err, "shop is unregistered").Error())
	}

	charge, createCharge := &omise.Charge{}, &operations.CreateCharge{
		Amount:   totalAmount,
		Currency: "thb",
		Customer: *user.OmiseCustomerId,
		Card:     paymentMethod.OmiseCardId,
	}
	if err := s.omiseClient.Do(charge, createCharge); err != nil {
		return nil, status.Error(InferCodeFromOmiseError(err), errors.Wrap(err, "can't charge").Error())
	}

	transfer, createTransfer := &omise.Transfer{}, &operations.CreateTransfer{
		Amount:    totalAmount,
		Recipient: *shop.OmiseResipientId,
	}
	if err := s.omiseClient.Do(transfer, createTransfer); err != nil {
		return nil, status.Error(InferCodeFromOmiseError(err), errors.Wrap(err, "can't transfer to shop").Error())
	}

	err = s.repo.CreatePaymentTransaction(&models.PaymentTransaction{
		ShopId:  in.ShopId,
		OrderId: in.OrderId,
		Amount:  int(totalAmount),
	})
	if err != nil {
		return nil, status.Error(postgres.InferCodeFromError(err), errors.Wrap(err, "can't create payment transaction").Error())
	}

	return &paymentproto.PayResponse{}, nil
}
