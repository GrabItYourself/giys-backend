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

func (s *Server) UpdateDefaultPaymentMethod(ctx context.Context, in *paymentproto.UpdateDefaultPaymentMethodRequest) (*paymentproto.UpdateDefaultPaymentMethodResponse, error) {
	identity, err := authutils.ExtractIdentityFromGrpcContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, errors.Wrap(err, "can't extract user from context").Error())
	}

	err = s.repo.UpdateDefaultPaymentMethodId(identity.UserId, in.PaymentMethodId)
	if err != nil {
		return nil, status.Error(postgres.InferCodeFromError(err), errors.Wrap(err, "can't update default payment method").Error())
	}

	return &paymentproto.UpdateDefaultPaymentMethodResponse{}, nil
}
