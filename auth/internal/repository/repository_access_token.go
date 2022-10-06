package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/GrabItYourself/giys-backend/auth/internal/types/accesstoken"
	"github.com/go-redis/redis/v9"
	"github.com/pkg/errors"
)

type AccessTokenKey struct {
	Token string
}

func (k *AccessTokenKey) Key() string {
	return fmt.Sprintf("access_token:%s", k.Token)
}

func (r *Repository) GetAccessToken(ctx context.Context, k *AccessTokenKey) (*accesstoken.AccessToken, error) {
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
	accessToken, err := accesstoken.FromMap(k.Key(), hash)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse AccessToken from hash")
	}
	return accessToken, nil
}

func (r *Repository) SetAccessToken(ctx context.Context, k *AccessTokenKey, accessToken *accesstoken.AccessToken) error {
	_, err := r.rdb.HSet(ctx, k.Key(), map[string]interface{}{
		"user_id": accessToken.UserId,
		"role":    accessToken.Role,
	}).Result()
	if err != nil {
		return errors.Wrapf(err, "HSet failed for key '%s'", k.Key())
	}
	_, err = r.rdb.Expire(ctx, k.Key(), 1*time.Hour).Result()
	if err != nil {
		return errors.Wrapf(err, "Expire failed for key '%s'", k.Key())
	}
	return nil
}

func (r *Repository) DeleteAccessToken(ctx context.Context, k *AccessTokenKey) error {
	_, err := r.rdb.Del(ctx, k.Key()).Result()
	if err != nil {
		return errors.Wrapf(err, "Del failed for key '%s'", k.Key())
	}
	return nil
}
