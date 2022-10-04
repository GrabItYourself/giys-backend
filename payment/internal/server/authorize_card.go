package server

import (
	"context"
	"time"

	"github.com/GrabItYourself/giys-backend/lib/authutils"
	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
	"github.com/GrabItYourself/giys-backend/payment/internal/libproto"
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

func (s *Server) AuthorizeCard(ctx context.Context, in *libproto.AuthorizeCardRequest) (*libproto.AuthorizeCardResponse, error) {
	userId, _, err := authutils.ExtractUserFromGrpcContext(ctx)
	if err != nil {
		return nil, err
	}

	user, err := s.repo.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	token, createToken := &omise.Token{}, &operations.CreateToken{
		Name:            in.Name,
		Number:          in.CardNumber,
		ExpirationMonth: time.Month(in.ExpirationMonth),
		ExpirationYear:  int(in.ExpirationYear),
	}
	err = s.omiseClient.Do(token, createToken)
	if err != nil {
		return nil, err
	}

	if user.OmiseCustomerId == nil {
		customer, createCustomer := &omise.Customer{}, &operations.CreateCustomer{
			Email: user.Email,
			Card:  token.ID,
		}
		err = s.omiseClient.Do(customer, createCustomer)
		if err != nil {
			return nil, err
		}

		err = s.repo.UpdateOmiseCustomerId(userId, customer.ID)
		if err != nil {
			return nil, err
		}

	} else {
		customer, updateCustomer := &omise.Customer{}, &operations.UpdateCustomer{
			CustomerID: *user.OmiseCustomerId,
			Card:       token.ID,
		}
		err = s.omiseClient.Do(customer, updateCustomer)
		if err != nil {
			return nil, err
		}
	}

	err = s.repo.CreatePaymentMethod(&models.PaymentMethod{
		UserId:      userId,
		OmiseCardId: token.Card.ID,
	})
	if err != nil {
		return nil, err
	}

	return &libproto.AuthorizeCardResponse{}, nil
}
