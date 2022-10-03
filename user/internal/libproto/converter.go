package libproto

import "github.com/GrabItYourself/giys-backend/lib/postgres/models"

func ConvertUserToProto(user *models.User) *User {
	return &User{
		Id:       user.ID,
		Role:     user.Role,
		Email:    user.Email,
		GoogleId: user.GoogleID,
	}
}
