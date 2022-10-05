package repository

import (
	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
)

func (r *Repository) GetShopById(id string) (*models.Shop, error) {
	var shop models.Shop
	err := r.pg.Where("id = ?", id).Take(&shop).Error
	if err != nil {
		return nil, err
	}
	return &shop, nil
}

func (r *Repository) UpdateOmiseRecipientId(id string, omiseRecipientId string) error {
	err := r.pg.Model(&models.Shop{}).Where("id = ?", id).Update("omise_recipient_id", omiseRecipientId).Error
	if err != nil {
		return err
	}
	return nil
}
