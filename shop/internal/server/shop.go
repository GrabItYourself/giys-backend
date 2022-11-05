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
		Image:       input.Image,
		Description: input.Description,
		Location:    input.Location,
		Contact:     input.Contact,
	}

	if err := s.repo.CreateShop(shop); err != nil {
		return nil, status.Error(postgres.InferCodeFromError(err), errors.Wrap(err, "can't create shop").Error())
	}
	return &shopproto.ShopResponse{
		Shop: &shopproto.Shop{
			Id:          shop.Id,
			Name:        shop.Name,
			Image:       shop.Image,
			Description: shop.Description,
			Location:    shop.Location,
			Contact:     shop.Contact,
		},
	}, nil
}

func (s *Server) GetAllShops(ctx context.Context, input *shopproto.GetAllShopsRequest) (*shopproto.AllShopsResponse, error) {
	shops, err := s.repo.GetAllShops()
	if err != nil {
		return nil, status.Error(postgres.InferCodeFromError(err), errors.Wrap(err, "can't get shops").Error())
	}
	items := make([]*shopproto.Shop, len(*shops))
	for index, item := range *shops {
		items[index] = &shopproto.Shop{
			Id:          item.Id,
			Name:        item.Name,
			Image:       item.Image,
			Description: item.Description,
			Location:    item.Location,
			Contact:     item.Contact,
		}
	}
	return &shopproto.AllShopsResponse{
		Shops: items,
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
			Name:        shop.Name,
			Image:       shop.Image,
			Description: shop.Description,
			Location:    shop.Location,
			Contact:     shop.Contact,
		},
	}, nil
}

func (s *Server) EditShop(ctx context.Context, input *shopproto.EditShopRequest) (*shopproto.ShopResponse, error) {
	shop := &models.Shop{
		Id:          input.EditedShop.Id,
		Name:        input.EditedShop.Name,
		Image:       input.EditedShop.Image,
		Description: input.EditedShop.Description,
		Location:    input.EditedShop.Location,
		Contact:     input.EditedShop.Contact,
	}
	editedShop, err := s.repo.EditShop(shop)
	if err != nil {
		return nil, status.Error(postgres.InferCodeFromError(err), errors.Wrap(err, "can't edit shop").Error())
	}
	return &shopproto.ShopResponse{
		Shop: &shopproto.Shop{
			Id:          editedShop.Id,
			Name:        editedShop.Name,
			Image:       editedShop.Image,
			Description: editedShop.Description,
			Location:    editedShop.Location,
			Contact:     editedShop.Contact,
		},
	}, nil
}

func (s *Server) DeleteShop(ctx context.Context, input *shopproto.DeleteShopRequest) (*shopproto.DeleteResponse, error) {
	rowsAffected, err := s.repo.DeleteShop(input.Id)
	if err != nil {
		return nil, status.Error(postgres.InferCodeFromError(err), errors.Wrap(err, "can't delete shop").Error())
	}
	return &shopproto.DeleteResponse{
		RowsAffected: rowsAffected,
	}, nil
}
