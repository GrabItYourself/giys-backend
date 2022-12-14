package repository

import (
	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
)

func (r *Repository) GetUserById(id string) (*models.User, error) {
	var user models.User
	err := r.pg.Where("id = ?", id).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.pg.Where("email = ?", email).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) GetUserByGoogleId(googleId string) (*models.User, error) {
	var user models.User
	err := r.pg.Where("google_id = ?", googleId).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) CreateUser(user *models.User) error {
	err := r.pg.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdateOmiseCustomerId(id string, omiseCustomerId string) error {
	err := r.pg.Model(&models.User{}).Where("id = ?", id).Update("omise_customer_id", omiseCustomerId).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdateDefaultPaymentMethodId(id string, defaultPaymentMethodId int32) error {
	err := r.pg.Model(&models.User{}).Where("id = ?", id).Update("default_payment_method_id", defaultPaymentMethodId).Error
	if err != nil {
		return err
	}
	return nil
}
