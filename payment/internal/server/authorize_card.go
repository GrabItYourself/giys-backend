package server

import (
	"context"
	"time"

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

func (s *Server) AuthorizeCard(ctx context.Context, in *paymentproto.AuthorizeCardRequest) (*paymentproto.AuthorizeCardResponse, error) {
	identity, err := authutils.ExtractIdentityFromGrpcContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, errors.Wrap(err, "can't extract user from context").Error())
	}

	user, err := s.repo.GetUserById(identity.UserId)
	if err != nil {
		return nil, status.Error(postgres.InferCodeFromError(err), errors.Wrap(err, "can't get user").Error())
	}

	token, createToken := &omise.Token{}, &operations.CreateToken{
		Name:            in.Name,
		Number:          in.CardNumber,
		ExpirationMonth: time.Month(in.ExpirationMonth),
		ExpirationYear:  int(in.ExpirationYear),
	}
	err = s.omiseClient.Do(token, createToken)
	if err != nil {
		return nil, status.Error(InferCodeFromOmiseError(err), errors.Wrap(err, "can't create token").Error())
	}

	if user.OmiseCustomerId == nil {
		customer, createCustomer := &omise.Customer{}, &operations.CreateCustomer{
			Email: user.Email,
			Card:  token.ID,
		}
		err = s.omiseClient.Do(customer, createCustomer)
		if err != nil {
			return nil, status.Error(InferCodeFromOmiseError(err), errors.Wrap(err, "can't create omise customer").Error())
		}

		err = s.repo.UpdateOmiseCustomerId(identity.UserId, customer.ID)
		if err != nil {
			return nil, status.Error(postgres.InferCodeFromError(err), errors.Wrap(err, "can't update omise customer id").Error())
		}

	} else {
		customer, updateCustomer := &omise.Customer{}, &operations.UpdateCustomer{
			CustomerID: *user.OmiseCustomerId,
			Card:       token.ID,
		}
		err = s.omiseClient.Do(customer, updateCustomer)
		if err != nil {
			return nil, status.Error(InferCodeFromOmiseError(err), errors.Wrap(err, "can't add card").Error())
		}
	}

	paymentMethod := &models.PaymentMethod{
		UserId:         identity.UserId,
		OmiseCardId:    token.Card.ID,
		LastFourDigits: in.CardNumber[len(in.CardNumber)-4:],
	}
	err = s.repo.CreatePaymentMethod(paymentMethod)
	if err != nil {
		return nil, status.Error(postgres.InferCodeFromError(err), errors.Wrap(err, "can't create payment method").Error())
	}

	if user.DefaultPaymentMethodId == nil {
		err = s.repo.UpdateDefaultPaymentMethodId(identity.UserId, paymentMethod.Id)
		if err != nil {
			return nil, status.Error(postgres.InferCodeFromError(err), errors.Wrap(err, "can't update default payment method").Error())
		}
	}

	return &paymentproto.AuthorizeCardResponse{
		PaymentMethod: &paymentproto.PaymentMethod{
			Id:             int64(paymentMethod.Id),
			LastFourDigits: paymentMethod.LastFourDigits,
		},
	}, nil
}
