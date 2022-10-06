package server

import (
	"context"

	"github.com/GrabItYourself/giys-backend/auth/pkg/authutils"
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
	userId, _, err := authutils.ExtractUserFromGrpcContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, errors.Wrap(err, "can't extract user from context").Error())
	}

	user, err := s.repo.GetUserById(userId)
	if err != nil {
		return nil, status.Error(postgres.InferCodeFromError(err), errors.Wrap(err, "can't get user").Error())
	}

	paymentMethod, err := s.repo.GetPaymentMethodById(in.CardId)
	if err != nil {
		return nil, status.Error(postgres.InferCodeFromError(err), errors.Wrap(err, "can't get payment method").Error())
	}

	shop, err := s.repo.GetShopById(in.ShopId)
	if err != nil {
		return nil, status.Error(postgres.InferCodeFromError(err), errors.Wrap(err, "can't get shop").Error())
	}

	charge, createCharge := &omise.Charge{}, &operations.CreateCharge{
		Amount:   in.Amount,
		Currency: "thb",
		Customer: *user.OmiseCustomerId,
		Card:     paymentMethod.OmiseCardId,
	}
	if e := s.omiseClient.Do(charge, createCharge); e != nil {
		return nil, status.Error(InferCodeFromOmiseError(err), errors.Wrap(err, "can't charge").Error())
	}

	transfer, createTransfer := &omise.Transfer{}, &operations.CreateTransfer{
		Amount:    in.Amount,
		Recipient: *shop.OmiseResipientId,
	}
	if e := s.omiseClient.Do(transfer, createTransfer); e != nil {
		return nil, status.Error(InferCodeFromOmiseError(err), errors.Wrap(err, "can't transfer to shop").Error())
	}

	err = s.repo.CreatePaymentTransaction(&models.PaymentTransaction{
		ShopId:  in.ShopId,
		OrderId: in.OrderId,
		Amount:  int(in.Amount),
	})
	if err != nil {
		return nil, status.Error(postgres.InferCodeFromError(err), errors.Wrap(err, "can't create payment transaction").Error())
	}

	return &paymentproto.PayResponse{}, nil
}
