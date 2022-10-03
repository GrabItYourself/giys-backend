package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/GrabItYourself/giys-backend/shop/internal/libproto"
)

func (*Server) CreateShop(ctx context.Context, input *libproto.CreateShopRequest) (*libproto.ShopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShop not implemented")
}

func (*Server) GetShop(ctx context.Context, input *libproto.GetShopRequest) (*libproto.ShopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShop not implemented")
}

func (*Server) EditShop(ctx context.Context, input *libproto.EditShopRequest) (*libproto.ShopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditShop not implemented")
}

func (*Server) DeleteShop(ctx context.Context, input *libproto.DeleteShopRequest) (*libproto.ShopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteShop not implemented")
}
