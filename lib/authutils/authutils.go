package authutils

import (
	"github.com/GrabItYourself/giys-backend/lib/proto/libproto"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

const userHeader = "Auth-User-Id"
const roleHeader = "Auth-User-Role"

// ExtractUserFromFiberContext extracts userId and role from context
func ExtractUserFromFiberContext(ctx *fiber.Ctx) (string, libproto.UserRole, error) {
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

// ExtractUserFromGrpcContext extracts userId and role from metadata in gRPC context
func ExtractUserFromGrpcContext(ctx context.Context) (string, libproto.UserRole, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", libproto.UserRole_NULL, errors.New("can't get metadata from context")
	}
	userId := md.Get(userHeader)
	if len(userId) == 0 {
		return "", libproto.UserRole_NULL, errors.Errorf("header %s is empty", userHeader)
	} else if len(userId) > 1 {
		return "", libproto.UserRole_NULL, errors.Errorf("header %s has more than one value", userHeader)
	}
	role := md.Get(roleHeader)
	if len(role) == 0 {
		return "", libproto.UserRole_NULL, errors.Errorf("header %s is empty", roleHeader)
	} else if len(role) > 1 {
		return "", libproto.UserRole_NULL, errors.Errorf("header %s has more than one value", roleHeader)
	}
	return userId[0], libproto.UserRole(libproto.UserRole_value[role[0]]), nil
}
