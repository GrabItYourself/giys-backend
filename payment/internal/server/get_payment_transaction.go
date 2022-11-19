package server

import (
	"context"

	"github.com/GrabItYourself/giys-backend/lib/postgres"
	"github.com/GrabItYourself/giys-backend/payment/pkg/paymentproto"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

func (s *Server) GetPaymentTransaction(ctx context.Context, in *paymentproto.GetPaymentTransactionRequest) (*paymentproto.GetPaymentTransactionResponse, error) {
	paymentTransaction, err := s.repo.GetPaymentTransactionByOrderIdAndShopId(in.OrderId, in.ShopId)
	if err != nil {
		return nil, status.Error(postgres.InferCodeFromError(err), errors.Wrap(err, "can't get payment transaction").Error())
	}

	return &paymentproto.GetPaymentTransactionResponse{
		Amount:    int32(paymentTransaction.Amount),
		CreatedAt: paymentTransaction.CreatedAt.UnixMilli(),
	}, nil
}
