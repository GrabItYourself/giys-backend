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

	paymentMethods, err := s.repo.GetMyPaymentMethods(identity.UserId)
	if err != nil {
		return nil, status.Error(postgres.InferCodeFromError(err), errors.Wrap(err, "can't get all payment methods").Error())
	}

	paymentMethodsResponse := make([]*paymentproto.PaymentMethod, len(*paymentMethods))
	for index := range *paymentMethods {
		paymentMethodsResponse[index] = &paymentproto.PaymentMethod{
			LastFourDigits: "1234",
		}
	}

	return &paymentproto.GetMyPaymentMethodsResponse{
		PaymentMethods: paymentMethodsResponse,
	}, nil
}
