package server

import (
	"context"

	"github.com/GrabItYourself/giys-backend/payment/internal/config"
	"github.com/GrabItYourself/giys-backend/payment/internal/repository"
	"github.com/GrabItYourself/giys-backend/payment/pkg/paymentproto"
	"github.com/omise/omise-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	paymentproto.UnimplementedPaymentServiceServer
	omiseClient *omise.Client
	repo        *repository.Repository
}

func NewServer(omiseConfig *config.OmiseConfig, repo *repository.Repository) (*Server, error) {
	client, err := omise.NewClient(omiseConfig.PublicKey, omiseConfig.SecretKey)
	if err != nil {
		return nil, err
	}

	return &Server{
		omiseClient: client,
		repo:        repo,
	}, nil
}

func (s *Server) Pay(context.Context, *paymentproto.PayRequest) (*paymentproto.PayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pay not implemented")
}

func (s *Server) RegisterRecipient(context.Context, *paymentproto.RegisterRecipientRequest) (*paymentproto.RegisterRecipientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterRecipient not implemented")
}

func (s *Server) GetTransactionHistory(context.Context, *paymentproto.GetTransactionHistoryRequest) (*paymentproto.GetTransactionHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionHistory not implemented")
}
