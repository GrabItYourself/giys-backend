package tokenmapping

import (
	"github.com/pkg/errors"
)

type TokenMapping struct {
	UserId       string
	AccessToken  string
	RefreshToken string
}

func FromMap(userId string, m map[string]string) (*TokenMapping, error) {
	mapping := &TokenMapping{UserId: userId}
	if accessToken, ok := m["access_token"]; !ok {
		return nil, errors.New("access_token not found in hash")
	} else {
		mapping.AccessToken = accessToken
	}
	if refreshToken, ok := m["refresh_token"]; !ok {
		return nil, errors.New("refresh_token not found in hash")
	} else {
		mapping.RefreshToken = refreshToken
	}
	return mapping, nil
}
