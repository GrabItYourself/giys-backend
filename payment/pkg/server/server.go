package server

import (
	"context"

	pb "github.com/GrabItYourself/giys-backend/lib/proto/payment"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedPaymentServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) Pay(context.Context, *pb.PayRequest) (*pb.PayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pay not implemented")
}

func (s *server) AuthorizeCard(context.Context, *pb.AuthorizeCardRequest) (*pb.AuthorizeCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthorizeCard not implemented")
}

func (s *server) RegisterRecipient(context.Context, *pb.RegisterRecipientRequest) (*pb.RegisterRecipientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterRecipient not implemented")
}

func (s *server) GetTransactionHistory(context.Context, *pb.GetTransactionHistoryRequest) (*pb.GetTransactionHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionHistory not implemented")
}
