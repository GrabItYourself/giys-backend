package repository

import "github.com/GrabItYourself/giys-backend/lib/postgres/models"

func (r *Repository) CreateShop(shop *models.Shop) error {
	err := r.pg.Create(shop).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetShopById(id string) (*models.Shop, error) {
	var shop models.Shop
	err := r.pg.Where("id = ?", id).Take(&shop).Error
	if err != nil {
		return nil, err
	}
	return &shop, nil
}

func (r *Repository) EditShop(shop *models.Shop) error {
	err := r.pg.Model(&shop).Updates(shop).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteShop(id string) error {
	err := r.pg.Delete(&models.Shop{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
