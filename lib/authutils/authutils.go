package authutils

import (
	"github.com/GrabItYourself/giys-backend/lib/proto/libproto"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

const userHeader = "Auth-User-Id"
const roleHeader = "Auth-User-Role"

// ExtractUserFromContext extracts userId and role from context
func ExtractUserFromContext(ctx *fiber.Ctx) (string, libproto.UserRole, error) {
	userId := ctx.Get(userHeader)
	if userId == "" {
		return "", libproto.UserRole_NULL, errors.Errorf("header %s is empty", userHeader)
	}
	role := ctx.Get(roleHeader)
	if role == "" {
		return "", libproto.UserRole_NULL, errors.Errorf("header %s is empty", roleHeader)
	}
	return userId, libproto.UserRole(libproto.UserRole_value[role]), nil
}
