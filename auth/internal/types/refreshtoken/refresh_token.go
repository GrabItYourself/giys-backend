package refreshtoken

import (
	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type RefreshToken struct {
	Token  string
	UserId string
	Role   models.RoleEnum
}

func FromMap(token string, m map[string]string) (*RefreshToken, error) {
	refreshToken := &RefreshToken{Token: token}
	if userId, ok := m["user_id"]; !ok {
		return nil, errors.New("user_id not found in hash")
	} else {
		refreshToken.UserId = userId
	}
	if role, ok := m["role"]; !ok {
		return nil, errors.New("role not found in hash")
	} else {
		refreshToken.Role = models.RoleEnum(role)
	}
	return refreshToken, nil
}

func New(userId string, role models.RoleEnum) *RefreshToken {
	newToken := &RefreshToken{
		Token:  uuid.New().String(),
		UserId: userId,
		Role:   role,
	}
	return newToken
}
