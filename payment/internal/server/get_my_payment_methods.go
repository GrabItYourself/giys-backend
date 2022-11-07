package server

import (
	"context"

	"github.com/GrabItYourself/giys-backend/auth/pkg/authutils"
	"github.com/GrabItYourself/giys-backend/lib/postgres"
	"github.com/GrabItYourself/giys-backend/payment/pkg/paymentproto"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetMyPaymentMethods(ctx context.Context, in *paymentproto.GetMyPaymentMethodsRequest) (*paymentproto.GetMyPaymentMethodsResponse, error) {
	identity, err := authutils.ExtractIdentityFromGrpcContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, errors.Wrap(err, "can't extract user from context").Error())
	}

	user, err := s.repo.GetUserById(identity.UserId)
	if err != nil {
		return nil, status.Error(postgres.InferCodeFromError(err), errors.Wrap(err, "can't get user").Error())
	}

	if user.DefaultPaymentMethodId == nil {
		return &paymentproto.GetMyPaymentMethodsResponse{
			PaymentMethods: make([]*paymentproto.PaymentMethod, 0),
		}, nil
	}

	paymentMethods, err := s.repo.GetMyPaymentMethods(identity.UserId)
	if err != nil {
		return nil, status.Error(postgres.InferCodeFromError(err), errors.Wrap(err, "can't get all payment methods").Error())
	}

	paymentMethodsResponse := make([]*paymentproto.PaymentMethod, len(*paymentMethods))
	for index, paymentMethod := range *paymentMethods {
		paymentMethodsResponse[index] = &paymentproto.PaymentMethod{
			Id:             paymentMethod.Id,
			LastFourDigits: paymentMethod.LastFourDigits,
			IsDefault:      paymentMethod.Id == *user.DefaultPaymentMethodId,
		}
	}

	return &paymentproto.GetMyPaymentMethodsResponse{
		PaymentMethods: paymentMethodsResponse,
	}, nil
}
