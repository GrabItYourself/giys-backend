package authutils

import (
	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

const (
	UserHeader = "Auth-User-Id"
	RoleHeader = "Auth-User-Role"
)

// ExtractUserFromGrpcContext extracts userId and role from metadata in gRPC context
func ExtractUserFromGrpcContext(ctx context.Context) (string, models.RoleEnum, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", "", errors.New("can't get metadata from context")
	}
	userId := md.Get(UserHeader)
	if len(userId) == 0 {
		return "", "", errors.Errorf("header %s is empty", UserHeader)
	} else if len(userId) > 1 {
		return "", "", errors.Errorf("header %s has more than one value", UserHeader)
	}
	role := md.Get(RoleHeader)
	if len(role) == 0 {
		return "", "", errors.Errorf("header %s is empty", RoleHeader)
	} else if len(role) > 1 {
		return "", "", errors.Errorf("header %s has more than one value", RoleHeader)
	}
	return userId[0], models.RoleEnum(role[0]), nil
}
