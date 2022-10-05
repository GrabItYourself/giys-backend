package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/GrabItYourself/giys-backend/shop/pkg/shopproto"
)

func (*Server) CreateShop(ctx context.Context, input *shopproto.CreateShopRequest) (*shopproto.ShopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShop not implemented")
}

func (*Server) GetShop(ctx context.Context, input *shopproto.GetShopRequest) (*shopproto.ShopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShop not implemented")
}

func (*Server) EditShop(ctx context.Context, input *shopproto.EditShopRequest) (*shopproto.ShopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditShop not implemented")
}

func (*Server) DeleteShop(ctx context.Context, input *shopproto.DeleteShopRequest) (*shopproto.ShopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteShop not implemented")
}
