package server

import (
	"context"

	"github.com/GrabItYourself/giys-backend/payment/pkg/paymentproto"
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateRecipient(ctx context.Context, in *paymentproto.UpdateRecipientRequest) (*paymentproto.UpdateRecipientResponse, error) {
	recipient, update := &omise.Recipient{}, &operations.UpdateRecipient{
		RecipientID: in.RecipientId,
		Name:        in.BankAccount.Name,
		Type:        omise.RecipientType(in.BankAccount.Type),
		BankAccount: &omise.BankAccount{
			Brand:  in.BankAccount.Brand,
			Number: in.BankAccount.Number,
			Name:   in.BankAccount.Name,
		},
	}

	if err := s.omiseClient.Do(recipient, update); err != nil {
		return nil, status.Error(InferCodeFromOmiseError(err), errors.Wrap(err, "can't update omise recipient").Error())
	}

	return &paymentproto.UpdateRecipientResponse{}, nil
}
