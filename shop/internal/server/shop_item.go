package server

import (
	"context"

	"google.golang.org/grpc/status"

	"github.com/GrabItYourself/giys-backend/lib/postgres"
	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
	"github.com/GrabItYourself/giys-backend/shop/pkg/shopproto"
	"github.com/pkg/errors"
)

func (s *Server) CreateShopItem(ctx context.Context, input *shopproto.CreateShopItemRequest) (*shopproto.ShopItemResponse, error) {
	shopItem := &models.ShopItem{
		ShopId: input.ShopId,
		Name:   input.Name,
		Image:  input.Image,
		Price:  input.Price,
	}
	if err := s.repo.CreateShopItem(shopItem); err != nil {
		return nil, status.Error(postgres.InferCodeFromError(err), errors.Wrap(err, "can't create shop item").Error())
	}
	return &shopproto.ShopItemResponse{
		Item: &shopproto.ShopItem{
			Id:     shopItem.Id,
			ShopId: shopItem.ShopId,
			Name:   shopItem.Name,
			Image:  shopItem.Image,
			Price:  shopItem.Price,
		},
	}, nil
}

func (s *Server) GetAllShopItems(ctx context.Context, input *shopproto.GetAllShopItemsRequest) (*shopproto.AllShopItemsResponse, error) {
	shopItems, err := s.repo.GetAllShopItems(input.ShopId)
	if err != nil {
		return nil, status.Error(postgres.InferCodeFromError(err), errors.Wrap(err, "can't get shop items").Error())
	}
	items := make([]*shopproto.ShopItem, len(*shopItems))
	for index, item := range *shopItems {
		items[index] = &shopproto.ShopItem{
			Id:     item.Id,
			ShopId: item.ShopId,
			Name:   item.Name,
			Image:  item.Image,
			Price:  item.Price,
		}
	}
	return &shopproto.AllShopItemsResponse{
		Items: items,
	}, nil
}

func (s *Server) GetShopItem(ctx context.Context, input *shopproto.GetShopItemRequest) (*shopproto.ShopItemResponse, error) {
	shopItem, err := s.repo.GetShopItemById(input.Id, input.ShopId)
	if err != nil {
		return nil, status.Error(postgres.InferCodeFromError(err), errors.Wrap(err, "can't get shop item").Error())
	}
	return &shopproto.ShopItemResponse{
		Item: &shopproto.ShopItem{
			Id:     shopItem.Id,
			ShopId: shopItem.ShopId,
			Name:   shopItem.Name,
			Image:  shopItem.Image,
			Price:  shopItem.Price,
		},
	}, nil
}

func (s *Server) EditShopItem(ctx context.Context, input *shopproto.EditShopItemRequest) (*shopproto.ShopItemResponse, error) {
	shopItem := &models.ShopItem{
		Id:     input.EditedItem.Id,
		ShopId: input.EditedItem.ShopId,
		Name:   input.EditedItem.Name,
		Image:  input.EditedItem.Image,
		Price:  input.EditedItem.Price,
	}
	editedShopItem, err := s.repo.EditShopItem(shopItem)
	if err != nil {
		return nil, status.Error(postgres.InferCodeFromError(err), errors.Wrap(err, "can't edit shop item").Error())
	}
	return &shopproto.ShopItemResponse{
		Item: &shopproto.ShopItem{
			Id:     editedShopItem.Id,
			ShopId: editedShopItem.ShopId,
			Name:   editedShopItem.Name,
			Image:  editedShopItem.Image,
			Price:  editedShopItem.Price,
		},
	}, nil
}

func (s *Server) DeleteShopItem(ctx context.Context, input *shopproto.DeleteShopItemRequest) (*shopproto.DeleteResponse, error) {
	rowsAffected, err := s.repo.DeleteShopItem(input.Id, input.ShopId)
	if err != nil {
		return nil, status.Error(postgres.InferCodeFromError(err), errors.Wrap(err, "can't delete shop item").Error())
	}
	return &shopproto.DeleteResponse{
		RowsAffected: rowsAffected,
	}, nil
}
