package v1handler

import (
	"context"

	"github.com/GrabItYourself/giys-backend/user/libproto"
	"github.com/pkg/errors"
)

func (h *Handler) HandleUserMe(ctx context.Context) (*libproto.User, error) {
	res, err := h.UserClient.Me(ctx, &libproto.MeReq{})
	if err != nil {
		return nil, errors.Wrap(err, "Failed to request GRPC MeReq")
	}
	return res.User, nil
}
