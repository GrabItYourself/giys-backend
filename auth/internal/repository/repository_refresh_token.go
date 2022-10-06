package repository

import (
	"context"
	"fmt"

	"github.com/GrabItYourself/giys-backend/auth/internal/types/refreshtoken"
	"github.com/go-redis/redis/v9"
	"github.com/pkg/errors"
)

type RefreshTokenKey struct {
	Token string
}

func (k *RefreshTokenKey) Key() string {
	return fmt.Sprintf("refresh_token:%s", k.Token)
}

func (r *Repository) GetRefreshToken(ctx context.Context, k *RefreshTokenKey) (*refreshtoken.RefreshToken, error) {
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
	refreshToken, err := refreshtoken.FromMap(k.Key(), hash)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse RefreshToken from hash")
	}
	return refreshToken, nil
}

func (r *Repository) SetRefreshToken(ctx context.Context, k *RefreshTokenKey, refreshToken *refreshtoken.RefreshToken) error {
	_, err := r.rdb.HSet(ctx, k.Key(), map[string]interface{}{
		"user_id": refreshToken.UserId,
		"role":    refreshToken.Role,
	}).Result()
	if err != nil {
		return errors.Wrapf(err, "HSet failed for key '%s'", k.Key())
	}
	return nil
}

func (r *Repository) DeleteRefreshToken(ctx context.Context, k *RefreshTokenKey) error {
	_, err := r.rdb.Del(ctx, k.Key()).Result()
	if err != nil {
		return errors.Wrapf(err, "Del failed for key '%s'", k.Key())
	}
	return nil
}
