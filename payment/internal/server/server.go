package server

import (
	"context"

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

func NewServer(omiseClient *omise.Client, repo *repository.Repository) (*Server, error) {
	return &Server{
		omiseClient: omiseClient,
		repo:        repo,
	}, nil
}

func (s *Server) GetTransactionHistory(context.Context, *paymentproto.GetTransactionHistoryRequest) (*paymentproto.GetTransactionHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionHistory not implemented")
}
