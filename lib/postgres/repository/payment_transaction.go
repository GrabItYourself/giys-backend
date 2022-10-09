package repository

import (
	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
)

func (r *Repository) GetPaymentTransactionById(id string) (*models.PaymentTransaction, error) {
	var paymentTransaction models.PaymentTransaction
	err := r.pg.Where("id = ?", id).Take(&paymentTransaction).Error
	if err != nil {
		return nil, err
	}
	return &paymentTransaction, nil
}

func (r *Repository) CreatePaymentTransaction(paymentTransaction *models.PaymentTransaction) error {
	err := r.pg.Create(paymentTransaction).Error
	if err != nil {
		return err
	}
	return nil
}
