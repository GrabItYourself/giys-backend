package repository

import (
	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
)

func (r *Repository) CreatePaymentMethod(paymentMethod *models.PaymentMethod) error {
	err := r.pg.Create(paymentMethod).Error
	if err != nil {
		return err
	}
	return nil
}
