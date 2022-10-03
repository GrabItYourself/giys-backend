package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/GrabItYourself/giys-backend/shop/internal/libproto"
)

func (*Server) CreateShopItem(ctx context.Context, input *libproto.CreateShopItemRequest) (*libproto.ShopItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShopItem not implemented")
}

func (*Server) GetAllShopItems(ctx context.Context, input *libproto.GetAllShopItemsRequest) (*libproto.AllShopItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShopItem not implemented")
}

func (*Server) GetShopItem(ctx context.Context, input *libproto.GetShopItemRequest) (*libproto.ShopItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShopItem not implemented")
}

func (*Server) EditShopItem(ctx context.Context, input *libproto.EditShopItemRequest) (*libproto.ShopItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditShopItem not implemented")
}

func (*Server) DeleteShopItem(ctx context.Context, input *libproto.DeleteShopItemRequest) (*libproto.ShopItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteShopItem not implemented")
}
