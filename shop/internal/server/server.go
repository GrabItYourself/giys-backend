package server

import (
	"github.com/GrabItYourself/giys-backend/lib/postgres/repository"
	"github.com/GrabItYourself/giys-backend/payment/pkg/paymentproto"
	"github.com/GrabItYourself/giys-backend/shop/pkg/shopproto"
)

type Server struct {
	shopproto.UnimplementedShopServiceServer
	repo          *repository.Repository
	paymentClient paymentproto.PaymentServiceClient
}

func NewServer(repo *repository.Repository, paymentClient paymentproto.PaymentServiceClient) *Server {
	return &Server{
		repo:          repo,
		paymentClient: paymentClient,
	}
}
