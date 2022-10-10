package repository

import (
	"gorm.io/gorm/clause"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
)

func (r *Repository) CreateShopItem(shopItem *models.ShopItem) error {
	var shopExists models.Shop
	err1 := r.pg.Where("id = ?", shopItem.ShopID).First(&shopExists).Error
	if err1 != nil {
		return status.Error(codes.InvalidArgument, "specified shop doesn't exist")
	}
	err2 := r.pg.Create(shopItem).Error
	if err2 != nil {
		return err2
	}
	return nil
}

func (r *Repository) GetAllShopItems(shopId int32) (*[]models.ShopItem, error) {
	var shopItems []models.ShopItem
	err := r.pg.Where("shop_id = ?", shopId).Find(&shopItems).Error
	if err != nil {
		return nil, err
	}
	return &shopItems, nil
}

func (r *Repository) GetShopItemById(id int32) (*models.ShopItem, error) {
	var shopItem models.ShopItem
	err := r.pg.Where("id = ?", id).Take(&shopItem).Error
	if err != nil {
		return nil, err
	}
	return &shopItem, nil
}

func (r *Repository) EditShopItem(shopItem *models.ShopItem) (*models.ShopItem, error) {
	err := r.pg.Model(&shopItem).Clauses(clause.Returning{}).Updates(shopItem).Error
	if err != nil {
		return nil, err
	}
	return shopItem, nil
}

func (r *Repository) DeleteShopItem(id int32) (int32, error) {
	result := r.pg.Delete(&models.ShopItem{}, id)
	if result.Error != nil {
		return 0, result.Error
	}
	return int32(result.RowsAffected), nil
}
