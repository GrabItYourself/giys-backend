package repository

import (
	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
)

type Repository struct {
	pg  *gorm.DB
	rdb *redis.Client
}

func New(pg *gorm.DB, rdb *redis.Client) *Repository {
	return &Repository{
		pg:  pg,
		rdb: rdb,
	}
}
