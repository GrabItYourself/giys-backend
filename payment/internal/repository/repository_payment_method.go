package repository

import (
	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
)

func (r *Repository) GetPaymentMethodById(id int64) (*models.PaymentMethod, error) {
	var paymentMethod models.PaymentMethod
	err := r.pg.Where("id = ?", id).Take(&paymentMethod).Error
	if err != nil {
		return nil, err
	}
	return &paymentMethod, nil
}

func (r *Repository) GetMyPaymentMethods(userId string) (*[]models.PaymentMethod, error) {
	var paymentMethods []models.PaymentMethod
	err := r.pg.Where("user_id = ?", userId).Find(&paymentMethods).Error
	if err != nil {
		return nil, err
	}
	return &paymentMethods, nil
}

func (r *Repository) CreatePaymentMethod(paymentMethod *models.PaymentMethod) error {
	err := r.pg.Create(paymentMethod).Error
	if err != nil {
		return err
	}
	return nil
}
