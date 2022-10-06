package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/GrabItYourself/giys-backend/shop/pkg/shopproto"
)

func (*Server) CreateShopItem(ctx context.Context, input *shopproto.CreateShopItemRequest) (*shopproto.ShopItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShopItem not implemented")
}

func (*Server) GetAllShopItems(ctx context.Context, input *shopproto.GetAllShopItemsRequest) (*shopproto.AllShopItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShopItem not implemented")
}

func (*Server) GetShopItem(ctx context.Context, input *shopproto.GetShopItemRequest) (*shopproto.ShopItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShopItem not implemented")
}

func (*Server) EditShopItem(ctx context.Context, input *shopproto.EditShopItemRequest) (*shopproto.ShopItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditShopItem not implemented")
}

func (*Server) DeleteShopItem(ctx context.Context, input *shopproto.DeleteShopItemRequest) (*shopproto.DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteShopItem not implemented")
}
