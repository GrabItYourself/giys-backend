package server

import (
	"context"

	"github.com/GrabItYourself/giys-backend/lib/postgres"
	"github.com/GrabItYourself/giys-backend/payment/pkg/paymentproto"
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

func (s *Server) RegisterRecipient(ctx context.Context, in *paymentproto.RegisterRecipientRequest) (*paymentproto.RegisterRecipientResponse, error) {
	recipient, create := &omise.Recipient{}, &operations.CreateRecipient{
		Name: in.BankAccount.Name,
		Type: omise.RecipientType(in.BankAccount.Type),
		BankAccount: &omise.BankAccount{
			Brand:  in.BankAccount.Brand,
			Number: in.BankAccount.Number,
			Name:   in.BankAccount.Name,
		},
	}
	if err := s.omiseClient.Do(recipient, create); err != nil {
		return nil, status.Error(InferCodeFromOmiseError(err), errors.Wrap(err, "can't create omise recipient").Error())
	}

	if err := s.repo.UpdateOmiseRecipientId(in.ShopId, recipient.ID); err != nil {
		return nil, status.Error(postgres.InferCodeFromError(err), errors.Wrap(err, "can't update omise recipient id").Error())
	}

	return &paymentproto.RegisterRecipientResponse{}, nil
}
