package server

import (
	"context"

	"google.golang.org/grpc/status"

	"github.com/GrabItYourself/giys-backend/lib/postgres"
	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
	"github.com/GrabItYourself/giys-backend/shop/pkg/shopproto"
	"github.com/pkg/errors"
)

func (s *Server) CreateShop(ctx context.Context, input *shopproto.CreateShopRequest) (*shopproto.ShopResponse, error) {
	shop := &models.Shop{
		Name:        input.Name,
		Image:       *input.Image,
		Description: *input.Description,
		Location:    *input.Location,
		Contact:     *input.Contact,
	}
	err := s.repo.CreateShop(shop)
	if err != nil {
		return nil, status.Error(postgres.InferCodeFromError(err), errors.Wrap(err, "can't create shop").Error())
	}
	return &shopproto.ShopResponse{
		Shop: &shopproto.Shop{
			Id:          shop.Id,
			Image:       &shop.Image,
			Description: &shop.Description,
			Location:    &shop.Location,
			Contact:     &shop.Contact,
		},
	}, nil
}

func (s *Server) GetShop(ctx context.Context, input *shopproto.GetShopRequest) (*shopproto.ShopResponse, error) {
	shop, err := s.repo.GetShopById(input.Id)
	if err != nil {
		return nil, status.Error(postgres.InferCodeFromError(err), errors.Wrap(err, "can't get shop").Error())
	}
	return &shopproto.ShopResponse{
		Shop: &shopproto.Shop{
			Id:          shop.Id,
			Image:       &shop.Image,
			Description: &shop.Description,
			Location:    &shop.Location,
			Contact:     &shop.Contact,
		},
	}, nil
}

func (s *Server) EditShop(ctx context.Context, input *shopproto.EditShopRequest) (*shopproto.ShopResponse, error) {
	shop := &models.Shop{
		Id:          input.EditedShop.Id,
		Name:        input.EditedShop.Name,
		Image:       *input.EditedShop.Image,
		Description: *input.EditedShop.Description,
		Location:    *input.EditedShop.Location,
		Contact:     *input.EditedShop.Contact,
	}
	err := s.repo.EditShop(shop)
	if err != nil {
		return nil, status.Error(postgres.InferCodeFromError(err), errors.Wrap(err, "can't edit shop").Error())
	}
	return &shopproto.ShopResponse{
		Shop: input.EditedShop,
	}, nil
}

func (s *Server) DeleteShop(ctx context.Context, input *shopproto.DeleteShopRequest) (*shopproto.ShopResponse, error) {
	err := s.repo.DeleteShop(input.Id)
	if err != nil {
		return nil, status.Error(postgres.InferCodeFromError(err), errors.Wrap(err, "can't delete shop").Error())
	}
	return &shopproto.ShopResponse{
		Shop: nil,
	}, nil
}
