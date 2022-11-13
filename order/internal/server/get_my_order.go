package server

import (
	"context"

	"github.com/GrabItYourself/giys-backend/auth/pkg/authutils"
	"github.com/GrabItYourself/giys-backend/lib/postgres"
	"github.com/GrabItYourself/giys-backend/order/pkg/orderproto"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetMyOrders(ctx context.Context, in *orderproto.GetMyOrdersRequest) (*orderproto.OrderListResponse, error) {
	identity, err := authutils.ExtractIdentityFromGrpcContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, errors.Wrap(err, "can't extract user from context").Error())
	}
	order, err := s.repo.GetMyOrders(identity.UserId)
	if err != nil {
		return nil, status.Errorf(postgres.InferCodeFromError(err), errors.Wrap(err, "Failed to get my orders").Error())
	}
	return &orderproto.OrderListResponse{
		Result: s.toProtoOrderListResponse(order),
	}, nil
}
