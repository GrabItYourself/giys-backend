package authutils

import (
	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

const userHeader = "Auth-User-Id"
const roleHeader = "Auth-User-Role"

// ExtractUserFromFiberContext extracts userId and role from context
func ExtractUserFromFiberContext(ctx *fiber.Ctx) (string, models.RoleEnum, error) {
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

// ExtractUserFromGrpcContext extracts userId and role from metadata in gRPC context
func ExtractUserFromGrpcContext(ctx context.Context) (string, models.RoleEnum, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", "", errors.New("can't get metadata from context")
	}
	userId := md.Get(userHeader)
	if len(userId) == 0 {
		return "", "", errors.Errorf("header %s is empty", userHeader)
	} else if len(userId) > 1 {
		return "", "", errors.Errorf("header %s has more than one value", userHeader)
	}
	role := md.Get(roleHeader)
	if len(role) == 0 {
		return "", "", errors.Errorf("header %s is empty", roleHeader)
	} else if len(role) > 1 {
		return "", "", errors.Errorf("header %s has more than one value", roleHeader)
	}
	return userId[0], models.RoleEnum(role[0]), nil
}
