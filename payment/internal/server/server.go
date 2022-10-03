package server

import (
	"context"

	"github.com/GrabItYourself/giys-backend/payment/internal/config"
	"github.com/GrabItYourself/giys-backend/payment/internal/libproto"
	"github.com/omise/omise-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	libproto.UnimplementedPaymentServiceServer
	omiseClient *omise.Client
}

func NewServer(omiseConfig *config.OmiseConfig) (*server, error) {
	client, err := omise.NewClient(omiseConfig.PublicKey, omiseConfig.SecretKey)
	if err != nil {
		return nil, err
	}

	return &server{omiseClient: client}, nil
}

func (s *server) Pay(context.Context, *libproto.PayRequest) (*libproto.PayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pay not implemented")
}

func (s *server) AuthorizeCard(context.Context, *libproto.AuthorizeCardRequest) (*libproto.AuthorizeCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthorizeCard not implemented")
}

func (s *server) RegisterRecipient(context.Context, *libproto.RegisterRecipientRequest) (*libproto.RegisterRecipientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterRecipient not implemented")
}

func (s *server) GetTransactionHistory(context.Context, *libproto.GetTransactionHistoryRequest) (*libproto.GetTransactionHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionHistory not implemented")
}
