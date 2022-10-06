package v1handler

import (
	"github.com/GrabItYourself/giys-backend/auth/pkg/authutils"
	"github.com/GrabItYourself/giys-backend/user/pkg/userproto"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

func (h *Handler) HandleUserMe(c *fiber.Ctx) (*userproto.User, error) {
	identity, ok := c.Locals(authutils.IdentityKey).(*authutils.Identity)
	if !ok {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "identity not found in context")
	}
	ctx, err := authutils.EmbedIdentityToContext(c.Context(), identity)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "can't embed identity to grpc context").Error())
	}

	res, err := h.Grpc.User.Me(ctx, &userproto.MeReq{})
	// TODO: infer error code from grpc error
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "Failed to request GRPC MeReq").Error())
	}
	return res.User, nil
}
