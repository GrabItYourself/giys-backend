package repository

import (
	"github.com/go-redis/redis/v9"
)

type Repository struct {
	rdb *redis.Client
}

func New(rdb *redis.Client) *Repository {
	return &Repository{
		rdb: rdb,
	}
}
