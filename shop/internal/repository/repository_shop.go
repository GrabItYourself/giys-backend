package repository

import (
	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
	"gorm.io/gorm/clause"
)

func (r *Repository) CreateShop(shop *models.Shop) error {
	err := r.pg.Create(shop).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetShopById(id int32) (*models.Shop, error) {
	var shop models.Shop
	err := r.pg.Where("id = ?", id).Take(&shop).Error
	if err != nil {
		return nil, err
	}
	return &shop, nil
}

func (r *Repository) EditShop(shop *models.Shop) (*models.Shop, error) {
	err := r.pg.Model(&shop).Clauses(clause.Returning{}).Updates(shop).Error
	if err != nil {
		return nil, err
	}
	return shop, nil
}

func (r *Repository) DeleteShop(id int32) (int32, error) {
	result := r.pg.Delete(&models.Shop{}, id)
	if result.Error != nil {
		return 0, result.Error
	}
	return int32(result.RowsAffected), nil
}
