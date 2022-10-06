package repository

import (
	"gorm.io/gorm"
)

type Repository struct {
	pg *gorm.DB
}

func New(pg *gorm.DB) *Repository {
	return &Repository{
		pg: pg,
	}
}
