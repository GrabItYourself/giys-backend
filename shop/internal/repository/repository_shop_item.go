package repository

import (
	"gorm.io/gorm/clause"

	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
)

func (r *Repository) CreateShopItem(shopItem *models.ShopItem) error {
	err := r.pg.Create(shopItem).Error
	if err != nil {
		return err
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

func (r *Repository) GetShopItemById(id int32, shopId int32) (*models.ShopItem, error) {
	var shopItem models.ShopItem
	err := r.pg.Where(&models.ShopItem{Id: id, ShopId: shopId}).Take(&shopItem).Error
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

func (r *Repository) DeleteShopItem(id int32, shopId int32) (int32, error) {
	result := r.pg.Where(&models.ShopItem{Id: id, ShopId: shopId}).Delete(&models.ShopItem{})
	if result.Error != nil {
		return 0, result.Error
	}
	return int32(result.RowsAffected), nil
}
