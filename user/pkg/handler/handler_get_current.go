package handler

import (
	"github.com/GrabItYourself/giys-backend/lib/authutils"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (h *handler) GetCurrentUser(ctx *fiber.Ctx) error {
	userId, _, err := authutils.ExtractUserFromContext(ctx)
	if err != nil {
		return fiber.NewError(401, errors.Wrap(err, "can't extract user from context").Error())
	}
	user, err := h.repo.GetUserById(userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(404, errors.Wrap(err, "can't get user").Error())
		}
		return fiber.NewError(500, errors.Wrap(err, "can't get user").Error())
	}
	return ctx.JSON(user)
}
