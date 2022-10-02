package handler

import (
	"github.com/GrabItYourself/giys-backend/user/pkg/repository"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	repo *repository.Repository
}

func NewHandler(repo *repository.Repository) Handler {
	return &handler{
		repo: repo,
	}
}

type Handler interface {
	GetCurrentUser(ctx *fiber.Ctx) error
}
