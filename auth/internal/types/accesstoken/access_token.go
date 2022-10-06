package accesstoken

import (
	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type AccessToken struct {
	Token  string
	UserId string
	Role   models.RoleEnum
}

func FromMap(token string, m map[string]string) (*AccessToken, error) {
	accessToken := &AccessToken{Token: token}
	if userId, ok := m["user_id"]; !ok {
		return nil, errors.New("user_id not found in hash")
	} else {
		accessToken.UserId = userId
	}
	if role, ok := m["role"]; !ok {
		return nil, errors.New("role not found in hash")
	} else {
		accessToken.Role = models.RoleEnum(role)
	}
	return accessToken, nil
}

func New(userId string, role models.RoleEnum) *AccessToken {
	newToken := &AccessToken{
		Token:  uuid.New().String(),
		UserId: userId,
		Role:   role,
	}
	return newToken
}
