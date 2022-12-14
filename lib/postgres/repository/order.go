package repository

import (
	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *Repository) GetOrderById(id int32, shopId int32, userId string) (*models.Order, error) {
	var order models.Order
	if err := r.pg.Where("id = ? AND shop_id = ? AND user_id = ?", id, shopId, userId).Preload("Items").Preload("Shop").Find(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *Repository) GetMyOrders(userId string) ([]models.Order, error) {
	var order []models.Order
	if err := r.pg.Where("user_id = ?", userId).Preload("Items").Preload("Shop").Find(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (r *Repository) GetShopOrders(shopId int32) ([]models.Order, error) {
	var order []models.Order
	if err := r.pg.Where("shop_id = ?", shopId).Preload("Items").Preload("Shop").Find(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (r *Repository) CreateOrder(order *models.Order) error {
	if err := r.pg.Create(order).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdateOrder(order *models.Order) error {
	if err := r.pg.Session(&gorm.Session{FullSaveAssociations: true}).Save(order).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdateOrderStatus(orderId int32, shopId int32, status models.OrderStatus) (*models.Order, error) {
	var order models.Order
	if err := r.pg.Model(&order).Clauses(clause.Returning{}).Where("id = ? AND shop_id = ?", orderId, shopId).Update("status", status).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *Repository) DeleteOrder(shopId int32, orderId int32, userId string) error {
	if err := r.pg.Where("id = ? AND shop_id = ? AND user_id = ?", orderId, shopId, userId).Delete(&models.Order{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetOrderWithShopItemsById(id int32, shopId int32) (*models.Order, error) {
	var order models.Order
	if err := r.pg.Where("id = ? AND shop_id = ?", id, shopId).Preload("Items.ShopItem").Find(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}
