package repository

import (
	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
)

func (r *Repository) UpdateOmiseRecipientId(id string, omiseRecipientId string) error {
	err := r.pg.Model(&models.Shop{}).Where("id = ?", id).Update("omise_recipient_id", omiseRecipientId).Error
	if err != nil {
		return err
	}
	return nil
}
