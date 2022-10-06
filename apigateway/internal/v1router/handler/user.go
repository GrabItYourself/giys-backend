package v1handler

import (
	"context"

	"github.com/GrabItYourself/giys-backend/lib/authutils"
	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
	"github.com/GrabItYourself/giys-backend/user/pkg/userproto"
	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"
)

func (h *Handler) HandleUserMe(ctx context.Context, userId string, role models.RoleEnum) (*userproto.User, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, authutils.UserHeader, userId, authutils.RoleHeader, string(role))
	res, err := h.UserClient.Me(ctx, &userproto.MeReq{})
	if err != nil {
		return nil, errors.Wrap(err, "Failed to request GRPC MeReq")
	}
	return res.User, nil
}
