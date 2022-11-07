package repository

import (
	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
)

func (r *Repository) GetOrderById(id int32, shopId int32) (*models.Order, error) {
	var order models.Order
	if err := r.pg.Where("id = ? AND shop_id = ?", id, shopId).Preload("Items.ShopItem").Find(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}
