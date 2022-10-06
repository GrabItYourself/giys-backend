package repository

import (
	"context"
	"fmt"

	"github.com/GrabItYourself/giys-backend/auth/internal/types/tokenmapping"
	"github.com/go-redis/redis/v9"
	"github.com/pkg/errors"
)

type TokenMappingKey struct {
	UserId string
}

func (k *TokenMappingKey) Key() string {
	return fmt.Sprintf("token_mapping:%s", k.UserId)
}

func (r *Repository) GetTokenMapping(ctx context.Context, k *TokenMappingKey) (*tokenmapping.TokenMapping, error) {
	exists, err := r.rdb.Exists(ctx, k.Key()).Result()
	if err != nil {
		return nil, errors.Wrapf(err, "Exists failed for key '%s'", k.Key())
	}
	if exists == 0 {
		return nil, errors.Wrapf(redis.Nil, "key '%s' not found", k.Key())
	}
	hash, err := r.rdb.HGetAll(ctx, k.Key()).Result()
	if err != nil {
		return nil, errors.Wrapf(err, "HGetAll failed for key '%s'", k.Key())
	}
	tokenMapping, err := tokenmapping.FromMap(k.Key(), hash)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse TokenMapping from hash")
	}
	return tokenMapping, nil
}

func (r *Repository) SetTokenMapping(ctx context.Context, k *TokenMappingKey, tokenMapping *tokenmapping.TokenMapping) error {
	_, err := r.rdb.HSet(ctx, k.Key(), map[string]interface{}{
		"access_token":  tokenMapping.AccessToken,
		"refresh_token": tokenMapping.RefreshToken,
	}).Result()
	if err != nil {
		return errors.Wrapf(err, "HSet failed for key '%s'", k.Key())
	}
	return nil
}
