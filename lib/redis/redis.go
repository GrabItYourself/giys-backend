package redis

import (
	"context"

	"github.com/GrabItYourself/giys-backend/lib/logger"
	"github.com/go-redis/redis/v9"
	"github.com/pkg/errors"
)

type Config struct {
	Address  string `mapstructure:"address"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"` // optional
}

func New(ctx context.Context, config *Config) (*redis.Client, error) {
	if config.Address == "" {
		config.Address = "localhost:6379"
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Address,
		Password: config.Password,
		DB:       config.Database,
	})
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		return nil, errors.Wrap(err, "Cannot connect to redis")
	}
	logger.Info("Postgres connected to " + config.Address)
	return redisClient, nil
}
