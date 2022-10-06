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
