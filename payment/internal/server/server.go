package server

import (
	"github.com/GrabItYourself/giys-backend/payment/internal/repository"
	"github.com/GrabItYourself/giys-backend/payment/pkg/paymentproto"
	"github.com/omise/omise-go"
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
