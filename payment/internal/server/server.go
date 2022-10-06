package server

import (
	"context"

	"github.com/GrabItYourself/giys-backend/payment/internal/config"
	"github.com/GrabItYourself/giys-backend/payment/internal/libproto"
	"github.com/GrabItYourself/giys-backend/payment/internal/repository"
	"github.com/omise/omise-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	libproto.UnimplementedPaymentServiceServer
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

func (s *Server) Pay(context.Context, *libproto.PayRequest) (*libproto.PayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pay not implemented")
}

func (s *Server) RegisterRecipient(context.Context, *libproto.RegisterRecipientRequest) (*libproto.RegisterRecipientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterRecipient not implemented")
}

func (s *Server) GetTransactionHistory(context.Context, *libproto.GetTransactionHistoryRequest) (*libproto.GetTransactionHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionHistory not implemented")
}
