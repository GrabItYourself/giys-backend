package repository

import (
	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
)

func (r *Repository) GetOrderById(id string) (*models.Order, error) {
	var order *models.Order
	if err := r.pg.Where("id = ?", id).First(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (r *Repository) CreateOrder(order *models.Order) error {
	if err := r.pg.Create(&models.Order{
		UserId: order.UserId,
		ShopId: order.ShopId,
		Status: order.Status,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdateOrder(order *models.Order) error {
	if err := r.pg.Save(order).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdateOrderStatus(orderId int32, shopId int32, status models.OrderStatus) (*models.Order, error) {
	var order models.Order
	if err := r.pg.Model(&order).Where("id = ? AND shop_id = ?", orderId, shopId).Update("status", status).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *Repository) DeleteOrder(id string) error {
	if err := r.pg.Delete(&models.Order{}, id).Error; err != nil {
		return err
	}
	return nil
}
