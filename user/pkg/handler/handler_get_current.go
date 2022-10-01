package handler

import (
	"github.com/GrabItYourself/giys-backend/lib/authutils"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// TODO: centralize error handling
func (h *handler) GetCurrentUser(ctx *fiber.Ctx) error {
	userId, _, err := authutils.ExtractUserFromContext(ctx)
	if err != nil {
		return ctx.Status(401).JSON(fiber.Map{"error": err.Error()})
	}
	user, err := h.repo.GetUserById(userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(404).JSON(fiber.Map{"error": "user not found"})
		}
		return ctx.Status(500).JSON(fiber.Map{"error": errors.Wrap(err, "can't get user by id").Error()})
	}
	return ctx.JSON(user)
}
