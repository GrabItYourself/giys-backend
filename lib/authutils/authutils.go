package authutils

import (
	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

const userHeader = "Auth-User-Id"
const roleHeader = "Auth-User-Role"

// ExtractUserFromContext extracts userId and role from context
func ExtractUserFromContext(ctx *fiber.Ctx) (string, models.RoleEnum, error) {
	userId := ctx.Get(userHeader)
	if userId == "" {
		return "", "", errors.Errorf("header %s is empty", userHeader)
	}
	role := ctx.Get(roleHeader)
	if role == "" {
		return "", "", errors.Errorf("header %s is empty", roleHeader)
	}
	return userId, models.RoleEnum(role), nil
}
